//go:generate easyjson -all commentActions.go
package coral

// CommentAction is the base Coral Comment Action that represents an action
// against a Comment.
type CommentAction struct {
	TenantID          string                 `json:"tenantID" validate:"required"`
	ID                string                 `json:"id" conform:"trim" validate:"required"`
	ActionType        string                 `json:"actionType" validate:"oneof=REACTION DONT_AGREE FLAG,required"`
	CommentID         string                 `json:"commentID" validate:"required"`
	CommentRevisionID string                 `json:"commentRevisionID" validate:"required"`
	Reason            string                 `json:"reason,omitempty" validate:"oneof=COMMENT_DETECTED_BANNED_WORD COMMENT_DETECTED_LINKS COMMENT_DETECTED_PREMOD_USER COMMENT_DETECTED_RECENT_HISTORY COMMENT_DETECTED_REPEAT_POST COMMENT_DETECTED_SPAM COMMENT_DETECTED_SUSPECT_WORD COMMENT_DETECTED_TOXIC COMMENT_REPORTED_OFFENSIVE COMMENT_REPORTED_OTHER COMMENT_REPORTED_SPAM"`
	AdditionalDetails string                 `json:"additionalDetails,omitempty"`
	StoryID           string                 `json:"storyID" validate:"required"`
	UserID            *string                `json:"userID" validate:"required"`
	CreatedAt         Time                   `json:"createdAt"`
	Metadata          map[string]interface{} `json:"metadata"`
	Imported          bool                   `json:"imported"`
}

// NewCommentAction will return an initialized CommentAction.
func NewCommentAction(tenantID string) *CommentAction {
	return &CommentAction{
		TenantID: tenantID,
		Metadata: map[string]interface{}{},
		Imported: true,
	}
}