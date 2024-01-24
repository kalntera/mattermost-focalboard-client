package boards

import (
)

// NotificationHint provides a hint that a block has been modified and has subscribers that
// should be notified.
// swagger:model
type NotificationHint struct {
	// BlockType is the block type of the entity (e.g. board, card) that was updated
	// required: true
	BlockType BlockType `json:"block_type"`

	// BlockID is id of the entity that was updated
	// required: true
	BlockID string `json:"block_id"`

	// ModifiedByID is the id of the user who made the block change
	ModifiedByID string `json:"modified_by_id"`

	// CreatedAt is the timestamp this notification hint was created in miliseconds since the current epoch
	// required: true
	CreateAt int64 `json:"create_at"`

	// NotifyAt is the timestamp this notification should be scheduled in miliseconds since the current epoch
	// required: true
	NotifyAt int64 `json:"notify_at"`
}

type ErrInvalidNotificationHint struct {
	msg string
}

func (e ErrInvalidNotificationHint) Error() string {
	return e.msg
}
