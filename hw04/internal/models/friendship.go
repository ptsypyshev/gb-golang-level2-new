package models

// Model for users friendship.
type Friendship struct {
	SourceID int `json:"source_id"`
	TargetID int `json:"target_id"`
}
