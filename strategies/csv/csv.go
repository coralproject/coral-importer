package csv

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"gitlab.com/coralproject/coral-importer/common"
	"gitlab.com/coralproject/coral-importer/common/coral"
	"gitlab.com/coralproject/coral-importer/common/pipeline"
)

// Import will perform the actual import process for the CSV strategy.
func Import(c *cli.Context) error {
	// tenantID is the ID of the Tenant that we are importing these documents
	// for.
	tenantID := c.String("tenantID")

	// output is the name of the folder where we are placing our outputted dumps
	// ready for MongoDB import.
	output := c.String("output")

	// input is the name of the folder where we are loading out collections
	// from the MongoDB export.
	input := c.String("input")

	// Validate that the collection files we expect exist in the input folder.
	if err := validateCollectionFilesExist(input); err != nil {
		return err
	}

	// Mark when we started.
	started := time.Now()

	// Write out all the comments to ${output}/comments.csv.
	commentsFileName := filepath.Join(input, "comments.csv")
	commentMap, err := pipeline.NewMapAggregator(
		pipeline.MergeTaskAggregatorOutputPipelines(
			pipeline.FanAggregatingProcessor(
				pipeline.NewCSVFileReader(commentsFileName, CommentColumns),
				ProcessCommentMap(),
			),
		),
	)
	if err != nil {
		logrus.WithError(err).Error("could not process comments")
		return err
	}

	logrus.WithField("comments", len(commentMap["storyID"])).WithField("children", len(commentMap["parentID"])).Debug("finished loading comments into map")

	startedReconstructionAt := time.Now()
	logrus.Debug("reconstructing families based on parent id map")

	// Reconstruct all the family relationships from the parentID map.
	reconstructor := common.NewReconstructor()
	for commentID, parentIDs := range commentMap["parentID"] {
		if len(parentIDs) == 1 {
			reconstructor.AddIDs(commentID, parentIDs[0])
		} else {
			reconstructor.AddIDs(commentID, "")
		}
	}

	logrus.WithField("took", time.Since(startedReconstructionAt).String()).Debug("finished family reconstruction")

	// Delete the reference to the parentID map that we don't need any more.
	delete(commentMap, "parentID")

	// Load all the comment statuses by reading the comments.json file again.
	statusCounts, err := pipeline.NewSummer(
		pipeline.MergeTaskSummerOutputPipelines(
			pipeline.FanSummerProcessor(
				pipeline.NewCSVFileReader(commentsFileName, CommentColumns),
				ProcessCommentStatusMap(),
			),
		),
	)
	if err != nil {
		logrus.WithError(err).Error("could not process status counts")
		return err
	}

	if err := pipeline.NewFileWriter(
		output,
		pipeline.MergeTaskWriterOutputPipelines(
			pipeline.FanWritingProcessors(
				pipeline.NewCSVFileReader(commentsFileName, CommentColumns),
				ProcessComments(tenantID, reconstructor),
			),
		),
	); err != nil {
		logrus.WithError(err).Error("could not process comments")
		return err
	}

	// Write out all the users to ${output}/users.csv.
	usersFileName := filepath.Join(input, "users.csv")
	if err := pipeline.NewFileWriter(
		output,
		pipeline.MergeTaskWriterOutputPipelines(
			pipeline.FanWritingProcessors(
				pipeline.NewCSVFileReader(usersFileName, UserColumns),
				ProcessUsers(tenantID),
			),
		),
	); err != nil {
		logrus.WithError(err).Error("could not process users")
		return err
	}

	// Write out all the stories to ${output}/stories.csv.
	storiesFileName := filepath.Join(input, "stories.csv")
	if err := pipeline.NewFileWriter(
		output,
		pipeline.MergeTaskWriterOutputPipelines(
			pipeline.FanWritingProcessors(
				pipeline.NewCSVFileReader(storiesFileName, StoryColumns),
				ProcessStories(tenantID, statusCounts),
			),
		),
	); err != nil {
		logrus.WithError(err).Error("could not process users")
		return err
	}

	// Mark when we finished.
	finished := time.Now()
	logrus.WithField("took", finished.Sub(started).String()).Info("finished processing")

	return nil
}

// IsHeaderRow will return true when the row contains the first field value as
// "id".
func IsHeaderRow(input *pipeline.TaskReaderInput) bool {
	if strings.ToLower(input.Fields[0]) == "id" {
		return true
	}

	return false
}

// ProcessCommentMap will collect maps based on the comment data.
func ProcessCommentMap() pipeline.AggregatingProcessor {
	return func(writer pipeline.AggregationWriter, input *pipeline.TaskReaderInput) error {
		// Ensure we skip the line if it's a header line.
		if IsHeaderRow(input) {
			return nil
		}

		c, err := ParseComment(input.Fields)
		if err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"line":    input.Line,
				"comment": input.Fields,
			}).Warn("failed to process comment")
			return nil
		}

		writer("status", c.ID, TranslateCommentStatus(c.Status))
		writer("storyID", c.ID, c.StoryID)
		writer("parentID", c.ID, c.ParentID)

		return nil
	}
}

// ProcessCommentStatusMap will link up comment statuses with the story ID.
func ProcessCommentStatusMap() pipeline.SummerProcessor {
	return func(writer pipeline.SummerWriter, input *pipeline.TaskReaderInput) error {
		// Ensure we skip the line if it's a header line.
		if IsHeaderRow(input) {
			return nil
		}

		c, err := ParseComment(input.Fields)
		if err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"line":    input.Line,
				"comment": input.Fields,
			}).Warn("failed to process comment")
			return nil
		}

		// Add the status to the map referencing the story id.
		writer(c.StoryID, TranslateCommentStatus(c.Status), 1)

		return nil
	}
}

// TranslateCommentStatus will convert the status that is expected as a part of
// the CSV import to the correct Coral status implementing a safe fallback.
func TranslateCommentStatus(status string) string {
	switch strings.ToUpper(status) {
	case "APPROVED":
		return "APPROVED"
	case "REJECTED":
		return "REJECTED"
	case "NONE":
		fallthrough
	default:
		return "NONE"
	}
}

// ProcessComments will parse each comment from the CSV input and emit comments
// to be created.
func ProcessComments(tenantID string, r *common.Reconstructor) pipeline.WritingProcessor {
	// Do this once for each unique policy, and use the policy for the life of the program
	// Policy creation/editing is not safe to use in multiple goroutines
	var p = bluemonday.UGCPolicy()

	return func(write pipeline.CollectionWriter, input *pipeline.TaskReaderInput) error {
		// Ensure we skip the line if it's a header line.
		if IsHeaderRow(input) {
			return nil
		}

		c, err := ParseComment(input.Fields)
		if err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"line":    input.Line,
				"comment": input.Fields,
			}).Warn("failed to process comment")
			return nil
		}

		createdAt, err := time.Parse(time.RFC3339, c.CreatedAt)
		if err != nil {
			return errors.Wrap(err, "could not parse created_at")
		}

		comment := coral.NewComment(tenantID)
		comment.ID = c.ID
		comment.AuthorID = c.AuthorID
		comment.StoryID = c.StoryID
		comment.CreatedAt.Time = createdAt

		body := coral.HTML(p.Sanitize(c.Body))

		revision := coral.Revision{
			ID:           comment.ID,
			Body:         body,
			Metadata:     coral.RevisionMetadata{},
			ActionCounts: map[string]int{},
			CreatedAt: coral.Time{
				Time: createdAt,
			},
		}
		comment.Revisions = []coral.Revision{
			revision,
		}
		comment.ParentID = c.ParentID

		// ID of the parent is the same as the revision ID.
		comment.ParentRevisionID = comment.ParentID

		comment.Status = TranslateCommentStatus(c.Status)

		// Add reconstructed data.
		comment.ChildIDs = r.GetChildren(comment.ID)
		comment.ChildCount = len(comment.ChildIDs)
		comment.AncestorIDs = r.GetAncestors(comment.ID)

		if err := common.Check(comment); err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"line":    input.Line,
				"comment": input.Fields,
			}).Warn("failed to process comment")
			return nil
		}

		if err := write("comments", comment); err != nil {
			return errors.Wrap(err, "couldn't write out comment")
		}

		return nil
	}
}

func ProcessStories(tenantID string, statusCounts map[string]map[string]int) pipeline.WritingProcessor {
	return func(write pipeline.CollectionWriter, input *pipeline.TaskReaderInput) error {
		// Ensure we skip the line if it's a header line.
		if IsHeaderRow(input) {
			return nil
		}

		s, err := ParseStory(input.Fields)
		if err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"line":  input.Line,
				"story": input.Fields,
			}).Warn("failed to process story")
			return nil
		}

		story := coral.NewStory(tenantID)
		story.ID = s.ID

		// Get the status counts for this story.
		storyStatusCounts, ok := statusCounts[story.ID]
		if ok {
			story.CommentCounts.Status.Approved = storyStatusCounts["APPROVED"]
			story.CommentCounts.Status.None = storyStatusCounts["NONE"]
			story.CommentCounts.Status.Rejected = storyStatusCounts["REJECTED"]

			// Begin updating the story's moderation queue counts.
			story.CommentCounts.ModerationQueue.Total += story.CommentCounts.Status.None
			story.CommentCounts.ModerationQueue.Total += story.CommentCounts.Status.Premod
			story.CommentCounts.ModerationQueue.Queues.Unmoderated += story.CommentCounts.Status.None
			story.CommentCounts.ModerationQueue.Queues.Unmoderated += story.CommentCounts.Status.Premod
		}

		story.URL = s.URL

		if s.Title != "" {
			story.Metadata.Title = s.Title
		}
		if s.Author != "" {
			story.Metadata.Author = s.Author
		}
		if s.PublishedAt != "" {
			publishedAt, err := time.Parse(time.RFC3339, s.PublishedAt)
			if err != nil {
				return errors.Wrap(err, "could not parse created_at")
			}

			story.Metadata.PublishedAt = &coral.Time{
				Time: publishedAt,
			}
		}
		if s.ClosedAt != "" {
			closedAt, err := time.Parse(time.RFC3339, s.ClosedAt)
			if err != nil {
				return errors.Wrap(err, "could not parse created_at")
			}

			story.ClosedAt = &coral.Time{
				Time: closedAt,
			}
		}

		if err := common.Check(story); err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"line":  input.Line,
				"story": input.Fields,
			}).Warn("failed to process story")
			return nil
		}

		if err := write("stories", story); err != nil {
			return errors.Wrap(err, "couldn't write out story")
		}

		return nil
	}
}

func ProcessUsers(tenantID string) pipeline.WritingProcessor {
	return func(write pipeline.CollectionWriter, input *pipeline.TaskReaderInput) error {
		// Ensure we skip the line if it's a header line.
		if IsHeaderRow(input) {
			return nil
		}

		u, err := ParseUser(input.Fields)
		if err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"line": input.Line,
				"user": input.Fields,
			}).Warn("failed to process user")
			return nil
		}

		// Parse the user from the file.
		user := coral.NewUser(tenantID)

		// id
		user.ID = u.ID

		// email
		user.Email = u.Email

		// created_at
		if u.CreatedAt != "" {
			createdAt, err := time.Parse(time.RFC3339, u.CreatedAt)
			if err != nil {
				return errors.Wrap(err, "could not parse created_at")
			}

			user.CreatedAt.Time = createdAt
		} else {
			user.CreatedAt.Time = time.Now()
		}

		// username
		user.Username = u.Username
		user.Status.UsernameStatus.History = []coral.UserUsernameStatusHistory{
			{
				ID:        uuid.NewV1().String(),
				Username:  user.Username,
				CreatedBy: user.ID,
				CreatedAt: user.CreatedAt,
			},
		}

		// role
		switch strings.ToUpper(u.Role) {
		case "ADMIN":
			user.Role = "ADMIN"
		case "MODERATOR":
			user.Role = "MODERATOR"
		case "COMMENTER":
			user.Role = "COMMENTER"
		default:
			user.Role = "COMMENTER"
		}

		// banned
		switch strings.ToLower(u.Banned) {
		case "true":
			user.Status.BanStatus.Active = true
			user.Status.BanStatus.History = []coral.UserBanStatusHistory{
				{
					ID:        uuid.NewV1().String(),
					Message:   "",
					Active:    true,
					CreatedAt: user.CreatedAt,
				},
			}
		case "false":
			fallthrough
		default:
			user.Status.BanStatus.Active = false
		}

		user.Profiles = []coral.UserProfile{
			{
				ID:           user.ID,
				Type:         "sso",
				LastIssuedAt: &user.CreatedAt,
			},
		}

		// Check the user.
		if err := common.Check(user); err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"line": input.Line,
				"user": input.Fields,
			}).Warn("failed to process user")
			return nil
		}

		if err := write("users", user); err != nil {
			return errors.Wrap(err, "couldn't write out user")
		}

		return nil
	}
}

// validateCollectionFilesExist will ensure that all the collection files that
// we expect to be in the input directory actually exist.
func validateCollectionFilesExist(input string) error {
	var collections = []string{
		"users",
		"stories",
		"comments",
	}

	for _, collection := range collections {
		filePath := filepath.Join(input, fmt.Sprintf("%s.csv", collection))
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return errors.Errorf("%s does not exist", filePath)
		}
	}

	return nil
}