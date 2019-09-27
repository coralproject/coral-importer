//go:generate easyjson -all stories.go
package coral

type CommentModerationCountsPerQueue struct {
	Unmoderated int `json:"unmoderated"`
	Pending     int `json:"pending"`
	Reported    int `json:"reported"`
}

type CommentModerationQueueCounts struct {
	Total  int                             `json:"total"`
	Queues CommentModerationCountsPerQueue `json:"queues"`
}

func NewCommentModerationQueueCounts() CommentModerationQueueCounts {
	return CommentModerationQueueCounts{
		Queues: CommentModerationCountsPerQueue{},
	}
}

type CommentStatusCounts struct {
	Approved       int `json:"APPROVED"`
	None           int `json:"NONE"`
	Premod         int `json:"PREMOD"`
	Rejected       int `json:"REJECTED"`
	SystemWithheld int `json:"SYSTEM_WITHHELD"`
}

type StoryCommentCounts struct {
	Action          map[string]int               `json:"action"`
	Status          CommentStatusCounts          `json:"status"`
	ModerationQueue CommentModerationQueueCounts `json:"moderationQueue"`
}

func NewStoryCommentCounts() StoryCommentCounts {
	return StoryCommentCounts{
		Action:          map[string]int{},
		Status:          CommentStatusCounts{},
		ModerationQueue: NewCommentModerationQueueCounts(),
	}
}

type StorySettings struct{}

type StoryMetadata struct {
	Title string `json:"title,omitempty"`
}

// Story is thye base Coral Story that is used in Coral.
type Story struct {
	TenantID      string             `json:"tenantID" validate:"required"`
	ID            string             `json:"id" conform:"trim" validate:"required"`
	URL           string             `json:"url" validate:"required,url"`
	CommentCounts StoryCommentCounts `json:"commentCounts"`
	Settings      StorySettings      `json:"settings"`
	Metadata      StoryMetadata      `json:"metadata"`
	CreatedAt     Time               `json:"createdAt" validate:"required"`
	Imported      bool               `json:"imported"`
}

func (s *Story) IncrementCommentCounts(status string) {
	switch status {
	case "APPROVED":
		s.CommentCounts.Status.Approved++
		return
	case "REJECTED":
		s.CommentCounts.Status.Rejected++
		return
	case "NONE":
		s.CommentCounts.Status.None++
		s.CommentCounts.ModerationQueue.Total++
		s.CommentCounts.ModerationQueue.Queues.Unmoderated++
		break
	case "PREMOD":
		s.CommentCounts.Status.Premod++
		s.CommentCounts.ModerationQueue.Total++
		s.CommentCounts.ModerationQueue.Queues.Pending++
		s.CommentCounts.ModerationQueue.Queues.Unmoderated++
		break
	case "SYSTEM_WITHHELD":
		s.CommentCounts.Status.SystemWithheld++
		s.CommentCounts.ModerationQueue.Total++
		s.CommentCounts.ModerationQueue.Queues.Pending++
		s.CommentCounts.ModerationQueue.Queues.Unmoderated++
		break
	}
}

// NewStory will return an initalized Story.
func NewStory(tenantID string) *Story {
	return &Story{
		TenantID:      tenantID,
		CommentCounts: NewStoryCommentCounts(),
		Settings:      StorySettings{},
		Metadata:      StoryMetadata{},
		Imported:      true,
	}
}
