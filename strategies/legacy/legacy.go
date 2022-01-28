package legacy

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/coralproject/coral-importer/common"
	"github.com/coralproject/coral-importer/common/coral"
	"github.com/coralproject/coral-importer/internal/utility"
	easyjson "github.com/mailru/easyjson"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var collections = []string{
	"actions",
	"assets",
	"comments",
	"settings",
	"users",
}

// PreferredPerspectiveModel is the stored preferred perspective model that
// should be used to copy over the perspective settings.
var PreferredPerspectiveModel string

// validateCollectionFilesExist will ensure that all the collection files that
// we expect to be in the input directory actually exist.
func validateCollectionFilesExist(input string) error {
	for _, collection := range collections {
		filePath := filepath.Join(input, fmt.Sprintf("%s.json", collection))
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return errors.Errorf("%s does not exist", filePath)
		}
	}

	return nil
}

type CommentRef struct {
	Status       string
	StoryID      string
	ParentID     string
	ActionCounts map[string]int
}

type StoryRef struct {
	ActionCounts map[string]int
	StatusCounts coral.CommentStatusCounts
}

type UserRef struct {
	StatusCounts coral.CommentStatusCounts
}

// Import will handle a data import task for importing comments into Coral from
// a legacy export.
func Import(c *cli.Context) error {
	// Copy over the preferredPerspectiveModel from the flags.
	PreferredPerspectiveModel = c.String("preferredPerspectiveModel")

	// tenantID is the ID of the Tenant that we are importing these documents
	// for.
	tenantID := c.String("tenantID")

	// siteID is the ID of the Site that we're importing records for.
	siteID := c.String("siteID")

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

	// Load all the comment actions from the actions.json file.
	actionsFileName := filepath.Join(input, "actions.json")
	commentsFileName := filepath.Join(input, "comments.json")

	comments := make(map[string]*CommentRef)
	if err := utility.ReadJSON(commentsFileName, func(line int, data []byte) error {
		var in Comment
		if err := easyjson.Unmarshal(data, &in); err != nil {
			logrus.WithField("line", line).Error(err)

			return errors.Wrap(err, "could not parse a comment")
		}

		// Check the input to ensure we're validated.
		if err := common.Check(&in); err != nil {
			logrus.WithError(err).WithField("line", line).Warn("validation failed for input user")

			return nil
		}

		ref := &CommentRef{
			Status:       TranslateCommentStatus(in.Status),
			StoryID:      in.AssetID,
			ActionCounts: map[string]int{},
		}

		if in.ParentID != nil {
			ref.ParentID = *in.ParentID
		}

		comments[in.ID] = ref

		return nil
	}); err != nil {
		return errors.Wrap(err, "could not process comment map")
	}

	logrus.WithField("comments", len(comments)).Debug("finished loading comments into map")

	commentActionsOutputFilename := filepath.Join(output, "commentActions.json")
	commentActionsWriter, err := utility.NewJSONWriter(commentActionsOutputFilename)
	if err != nil {
		return errors.Wrap(err, "could not create commentActionsWriter")
	}
	defer commentActionsWriter.Close()

	stories := make(map[string]*StoryRef)
	if err := utility.ReadJSON(actionsFileName, func(line int, data []byte) error {
		// Parse the Action from the file.
		var in Action
		if err := easyjson.Unmarshal(data, &in); err != nil {
			logrus.WithField("line", line).Error(err)

			return errors.Wrap(err, "could not parse an action")
		}

		// Ignore the action if it's not a comment action.
		if in.ItemType != "COMMENTS" {
			logrus.WithField("line", line).Warn("skipping non-comment flag")

			return nil
		}

		// Check the input to ensure we're validated.
		if err := common.Check(&in); err != nil {
			return errors.Wrap(err, "checking failed input Action")
		}

		// Translate the action to a comment action.
		action := TranslateCommentAction(tenantID, siteID, &in)
		ref, ok := comments[action.CommentID]
		if !ok {
			return nil
		}
		action.StoryID = ref.StoryID

		if err := commentActionsWriter.Write(action); err != nil {
			return errors.Wrap(err, "couldn't write out commentAction")
		}

		story, ok := stories[ref.StoryID]
		if !ok {
			story = &StoryRef{
				ActionCounts: map[string]int{},
			}
			stories[ref.StoryID] = story
		}

		ref.ActionCounts[action.ActionType]++
		story.ActionCounts[action.ActionType]++
		if action.ActionType == "FLAG" {
			ref.ActionCounts[action.ActionType+"__"+action.Reason]++
			story.ActionCounts[action.ActionType+"__"+action.Reason]++
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "could not process comment actions")
	}

	logrus.Debug("finished writing out comment actions")

	startedReconstructionAt := time.Now()
	logrus.Debug("reconstructing families based on parent id map")

	// Reconstruct all the family relationships from the parentID map.
	re := common.NewReconstructor()
	for commentID, comment := range comments {
		re.AddIDs(commentID, comment.ParentID)
	}

	logrus.WithField("took", time.Since(startedReconstructionAt).String()).Debug("finished family reconstruction")

	// Load all the comments in from the comments.json file.
	commentsOutputFilename := filepath.Join(output, "comments.json")
	commentsWriter, err := utility.NewJSONWriter(commentsOutputFilename)
	if err != nil {
		return errors.Wrap(err, "could not create comments writer")
	}
	defer commentsWriter.Close()

	users := make(map[string]*UserRef)
	if err := utility.ReadJSON(commentsFileName, func(line int, data []byte) error {
		// Parse the Comment from the file.
		var in Comment
		if err := easyjson.Unmarshal(data, &in); err != nil {
			logrus.WithField("line", line).Error(err)

			return errors.Wrap(err, "could not parse an comment")
		}

		// Check the input to ensure we're validated.
		if err := common.Check(&in); err != nil {
			return errors.Wrap(err, "checking failed input Action")
		}

		comment := TranslateComment(tenantID, siteID, &in)

		ref, ok := comments[comment.ID]
		if !ok {
			return errors.New("could not find comment ref")
		}

		// Associate the action count data.
		comment.ActionCounts = ref.ActionCounts
		if comment.DeletedAt == nil {
			comment.Revisions[len(comment.Revisions)-1].ActionCounts = ref.ActionCounts
		}

		// Add reconstructed data.
		comment.ChildIDs = re.GetChildren(comment.ID)
		comment.ChildCount = len(comment.ChildIDs)
		comment.AncestorIDs = re.GetAncestors(comment.ID)

		user, ok := users[comment.AuthorID]
		if !ok {
			user = &UserRef{}
			users[comment.AuthorID] = user
		}
		user.StatusCounts.Increment(comment.Status, 1)

		story, ok := stories[ref.StoryID]
		if !ok {
			story = &StoryRef{
				ActionCounts: map[string]int{},
			}
			stories[ref.StoryID] = story
		}
		story.StatusCounts.Increment(comment.Status, 1)

		if err := commentsWriter.Write(comment); err != nil {
			return errors.Wrap(err, "couldn't write out comment")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "could not read comments json")
	}

	// Walk across all the comment's status maps so we can determine how many
	// comments should be in the reported queue in each story.
	reportedMap := make(map[string]int)
	for _, comment := range comments {
		if comment.Status != "NONE" {
			// If the status is not none, then it can't be reported!
			continue
		}

		flagCount := comment.ActionCounts["FLAG"]
		if flagCount > 0 {
			// Increment the storyID.
			reportedMap[comment.StoryID]++
		}
	}

	// Process the stories now.
	storiesOutputFilename := filepath.Join(output, "stories.json")
	storiesWriter, err := utility.NewJSONWriter(storiesOutputFilename)
	if err != nil {
		return errors.Wrap(err, "could not create stories writer")
	}
	defer storiesWriter.Close()

	assetsFileName := filepath.Join(input, "assets.json")
	if err := utility.ReadJSON(assetsFileName, func(line int, data []byte) error {
		// Parse the asset from the file.
		var in Asset
		if err := easyjson.Unmarshal(data, &in); err != nil {
			logrus.WithField("line", line).Error(err)

			return errors.Wrap(err, "could not parse an asset")
		}

		story := TranslateAsset(tenantID, siteID, &in)

		// Get the status counts for this story.
		ref := stories[story.ID]

		story.CommentCounts.Status = ref.StatusCounts

		// Get the action counts for this story.
		story.CommentCounts.Action = ref.ActionCounts

		// Begin updating the story's moderation queue counts.
		story.CommentCounts.ModerationQueue.Total += story.CommentCounts.Status.None
		story.CommentCounts.ModerationQueue.Total += story.CommentCounts.Status.SystemWithheld
		story.CommentCounts.ModerationQueue.Total += story.CommentCounts.Status.Premod
		story.CommentCounts.ModerationQueue.Queues.Unmoderated += story.CommentCounts.Status.None
		story.CommentCounts.ModerationQueue.Queues.Unmoderated += story.CommentCounts.Status.SystemWithheld
		story.CommentCounts.ModerationQueue.Queues.Unmoderated += story.CommentCounts.Status.Premod
		story.CommentCounts.ModerationQueue.Queues.Pending += story.CommentCounts.Status.Premod
		story.CommentCounts.ModerationQueue.Queues.Pending += story.CommentCounts.Status.SystemWithheld

		// Update the reported queue counts based on the reported map.
		story.CommentCounts.ModerationQueue.Total += reportedMap[story.ID]
		story.CommentCounts.ModerationQueue.Queues.Reported += reportedMap[story.ID]

		if err := storiesWriter.Write(story); err != nil {
			return errors.Wrap(err, "couldn't write out story")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "could not process stories")
	}

	// Write out all the users to ${output}/users.json.
	usersOutputFilename := filepath.Join(output, "users.json")
	usersWriter, err := utility.NewJSONWriter(usersOutputFilename)
	if err != nil {
		return errors.Wrap(err, "could not create users writer")
	}
	defer usersWriter.Close()

	usersFileName := filepath.Join(input, "users.json")
	if err := utility.ReadJSON(usersFileName, func(line int, data []byte) error {
		// Parse the user from the file.
		var in User
		if err := easyjson.Unmarshal(data, &in); err != nil {
			logrus.WithField("line", line).Error(err)

			return errors.Wrap(err, "could not parse an user")
		}

		user := TranslateUser(tenantID, &in)

		// Get the status counts for this story.
		ref := users[user.ID]

		user.CommentCounts.Status = ref.StatusCounts

		if err := usersWriter.Write(user); err != nil {
			return errors.Wrap(err, "couldn't write out user")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "could not process users")
	}

	// Mark when we finished.
	finished := time.Now()
	logrus.WithField("took", finished.Sub(started).String()).Info("finished processing")

	return nil
}
