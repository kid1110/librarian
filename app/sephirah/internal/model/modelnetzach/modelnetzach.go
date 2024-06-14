package modelnetzach

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type NotifyFlow struct {
	ID          model.InternalID
	Name        string
	Description string
	Sources     []*NotifyFlowSource
	Targets     []*NotifyFlowTarget
	Status      NotifyFlowStatus
}

type NotifyFlowSource struct {
	FeedConfigID model.InternalID
	Filter       *NotifyFilter
}

type NotifyFlowTarget struct {
	TargetID  model.InternalID
	Filter    *NotifyFilter
	ChannelID string
}

type NotifyFlowStatus int

const (
	NotifyFlowStatusUnspecified NotifyFlowStatus = iota
	NotifyFlowStatusActive
	NotifyFlowStatusSuspend
)

type NotifyTarget struct {
	ID          model.InternalID
	Name        string
	Description string
	Destination string
	Status      NotifyTargetStatus
	Token       string
}

type NotifyTargetStatus int

const (
	NotifyTargetStatusUnspecified NotifyTargetStatus = iota
	NotifyTargetStatusActive
	NotifyTargetStatusSuspend
)

type NotifyFilter struct {
	ExcludeKeywords []string
	IncludeKeywords []string
}

type SystemNotification struct {
	ID         model.InternalID
	Type       SystemNotificationType
	Level      SystemNotificationLevel
	Status     SystemNotificationStatus
	Title      string
	Content    string
	CreateTime time.Time
}

type SystemNotificationType int

const (
	SystemNotificationTypeUnspecified SystemNotificationType = iota
	SystemNotificationTypeSystem
	SystemNotificationTypeUser
)

type SystemNotificationLevel int

const (
	SystemNotificationLevelUnspecified SystemNotificationLevel = iota
	SystemNotificationLevelOngoing
	SystemNotificationLevelError
	SystemNotificationLevelWarning
	SystemNotificationLevelInfo
)

type SystemNotificationStatus int

const (
	SystemNotificationStatusUnspecified SystemNotificationStatus = iota
	SystemNotificationStatusUnread
	SystemNotificationStatusRead
	SystemNotificationStatusDismissed
)
