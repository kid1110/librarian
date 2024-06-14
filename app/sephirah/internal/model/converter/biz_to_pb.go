package converter

import (
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// goverter:converter
// goverter:output:file ./generated.go
// goverter:output:package github.com/tuihub/librarian/app/sephirah/internal/model/converter
// goverter:extend ToPBInternalID
// goverter:extend ToPBTime
// goverter:extend ToPBTimePtr
// goverter:extend ToPBDuration
type toPBConverter interface { //nolint:unused // used by generator
	ToPBTimeRange(*model.TimeRange) *librarian.TimeRange
	ToPBInternalIDList([]model.InternalID) []*librarian.InternalID
	ToPBServerFeatureSummary(*modeltiphereth.ServerFeatureSummary) *pb.ServerFeatureSummary

	// goverter:matchIgnoreCase
	// goverter:map ID DeviceId
	ToPBDeviceInfo(*modeltiphereth.DeviceInfo) *pb.DeviceInfo
	ToPBDeviceInfoList([]*modeltiphereth.DeviceInfo) []*pb.DeviceInfo
	// goverter:enum:unknown SystemType_SYSTEM_TYPE_UNSPECIFIED
	// goverter:enum:map SystemTypeUnspecified SystemType_SYSTEM_TYPE_UNSPECIFIED
	// goverter:enum:map SystemTypeIOS SystemType_SYSTEM_TYPE_IOS
	// goverter:enum:map SystemTypeAndroid SystemType_SYSTEM_TYPE_ANDROID
	// goverter:enum:map SystemTypeWeb SystemType_SYSTEM_TYPE_WEB
	// goverter:enum:map SystemTypeWindows SystemType_SYSTEM_TYPE_WINDOWS
	// goverter:enum:map SystemTypeMacOS SystemType_SYSTEM_TYPE_MACOS
	// goverter:enum:map SystemTypeLinux SystemType_SYSTEM_TYPE_LINUX
	ToPBSystemType(modeltiphereth.SystemType) pb.SystemType

	// goverter:matchIgnoreCase
	// goverter:map CreateAt CreateTime
	// goverter:map ExpireAt ExpireTime
	ToPBUserSession(*modeltiphereth.UserSession) *pb.UserSession
	ToPBUserSessionList([]*modeltiphereth.UserSession) []*pb.UserSession

	// goverter:matchIgnoreCase
	// goverter:ignore Password
	ToPBUser(*modeltiphereth.User) *pb.User
	ToPBUserList([]*modeltiphereth.User) []*pb.User
	// goverter:enum:unknown UserType_USER_TYPE_UNSPECIFIED
	// goverter:enum:map UserTypeUnspecified UserType_USER_TYPE_UNSPECIFIED
	// goverter:enum:map UserTypeAdmin UserType_USER_TYPE_ADMIN
	// goverter:enum:map UserTypeNormal UserType_USER_TYPE_NORMAL
	// goverter:enum:map UserTypeSentinel UserType_USER_TYPE_SENTINEL
	// goverter:enum:map UserTypePorter UserType_USER_TYPE_PORTER
	ToPBUserType(libauth.UserType) pb.UserType
	// goverter:enum:unknown UserStatus_USER_STATUS_UNSPECIFIED
	// goverter:enum:map UserStatusUnspecified UserStatus_USER_STATUS_UNSPECIFIED
	// goverter:enum:map UserStatusActive UserStatus_USER_STATUS_ACTIVE
	// goverter:enum:map UserStatusBlocked UserStatus_USER_STATUS_BLOCKED
	ToPBUserStatus(modeltiphereth.UserStatus) pb.UserStatus

	// goverter:matchIgnoreCase
	ToPBAccount(*modeltiphereth.Account) *librarian.Account
	ToPBAccountList([]*modeltiphereth.Account) []*librarian.Account

	// goverter:matchIgnoreCase
	// goverter:map Status | ToPBPorterStatus
	// goverter:ignore FeatureSummary
	ToPBPorter(*modeltiphereth.PorterInstance) *pb.Porter
	ToPBPorterList([]*modeltiphereth.PorterInstance) []*pb.Porter
	// goverter:enum:unknown PorterConnectionStatus_PORTER_CONNECTION_STATUS_UNSPECIFIED
	// goverter:enum:map PorterConnectionStatusUnspecified PorterConnectionStatus_PORTER_CONNECTION_STATUS_UNSPECIFIED
	// goverter:enum:map PorterConnectionStatusConnected PorterConnectionStatus_PORTER_CONNECTION_STATUS_CONNECTED
	// goverter:enum:map PorterConnectionStatusDisconnected PorterConnectionStatus_PORTER_CONNECTION_STATUS_DISCONNECTED
	// goverter:enum:map PorterConnectionStatusActive PorterConnectionStatus_PORTER_CONNECTION_STATUS_ACTIVE
	// goverter:enum:map PorterConnectionStatusActivationFailed PorterConnectionStatus_PORTER_CONNECTION_STATUS_ACTIVATION_FAILED
	ToPBPorterConnectionStatus(modeltiphereth.PorterConnectionStatus) pb.PorterConnectionStatus

	// goverter:matchIgnoreCase
	// goverter:ignore AltNames
	ToPBAppInfo(*modelgebura.AppInfo) *librarian.AppInfo
	// goverter:matchIgnoreCase
	// goverter:ignore ImageUrls
	ToPBAppInfoDetail(*modelgebura.AppInfoDetails) *librarian.AppInfoDetails
	ToPBAppInfoList([]*modelgebura.AppInfo) []*librarian.AppInfo
	// goverter:matchIgnoreCase
	// goverter:ignore AltNames
	ToPBAppInfoMixed(*modelgebura.AppInfoMixed) *librarian.AppInfoMixed
	ToPBAppInfoMixedList([]*modelgebura.AppInfoMixed) []*librarian.AppInfoMixed
	// goverter:enum:unknown AppType_APP_TYPE_UNSPECIFIED
	// goverter:enum:map AppTypeUnspecified AppType_APP_TYPE_UNSPECIFIED
	// goverter:enum:map AppTypeGame AppType_APP_TYPE_GAME
	ToPBAppType(modelgebura.AppType) librarian.AppType

	// goverter:matchIgnoreCase
	ToPBApp(*modelgebura.App) *pb.App
	ToPBAppList([]*modelgebura.App) []*pb.App
	// goverter:matchIgnoreCase
	// goverter:ignore Id
	// goverter:ignore TokenServerUrl
	// goverter:ignore Chunks
	ToPBAppBinary(*modelgebura.AppBinary) *pb.AppBinary

	// goverter:matchIgnoreCase
	ToPBAppInst(*modelgebura.AppInst) *pb.AppInst
	ToPBAppInstList([]*modelgebura.AppInst) []*pb.AppInst

	// goverter:matchIgnoreCase
	ToPBFeed(*modelfeed.Feed) *librarian.Feed
	// goverter:matchIgnoreCase
	ToPBFeedItem(*modelfeed.Item) *librarian.FeedItem
	ToPBFeedItemList([]*modelfeed.Item) []*librarian.FeedItem
	// goverter:matchIgnoreCase
	ToPBFeedImage(*modelfeed.Image) *librarian.FeedImage
	// goverter:matchIgnoreCase
	ToPBEnclosure(*modelfeed.Enclosure) *librarian.FeedEnclosure
	// goverter:matchIgnoreCase
	// goverter:map LatestPullStatus | ToPBFeedConfigPullStatus
	ToPBFeedConfig(*modelyesod.FeedConfig) *pb.FeedConfig
	// goverter:enum:unknown FeedConfigStatus_FEED_CONFIG_STATUS_UNSPECIFIED
	// goverter:enum:map FeedConfigStatusUnspecified FeedConfigStatus_FEED_CONFIG_STATUS_UNSPECIFIED
	// goverter:enum:map FeedConfigStatusActive FeedConfigStatus_FEED_CONFIG_STATUS_ACTIVE
	// goverter:enum:map FeedConfigStatusSuspend FeedConfigStatus_FEED_CONFIG_STATUS_SUSPEND
	ToPBFeedConfigStatus(modelyesod.FeedConfigStatus) pb.FeedConfigStatus
	// goverter:matchIgnoreCase
	// goverter:map FeedConfig Config
	ToPBFeedWithConfig(*modelyesod.FeedWithConfig) *pb.ListFeedConfigsResponse_FeedWithConfig
	ToPBFeedWithConfigList([]*modelyesod.FeedWithConfig) []*pb.ListFeedConfigsResponse_FeedWithConfig
	// goverter:matchIgnoreCase
	ToPBFeedItemDigest(*modelyesod.FeedItemDigest) *pb.FeedItemDigest
	ToPBFeedItemDigestList([]*modelyesod.FeedItemDigest) []*pb.FeedItemDigest

	// goverter:matchIgnoreCase
	ToPBFeedItemCollection(*modelyesod.FeedItemCollection) *pb.FeedItemCollection
	ToPBFeedItemCollectionList([]*modelyesod.FeedItemCollection) []*pb.FeedItemCollection

	// goverter:matchIgnoreCase
	ToPBNotifyTarget(*modelnetzach.NotifyTarget) *pb.NotifyTarget
	ToPBNotifyTargetList([]*modelnetzach.NotifyTarget) []*pb.NotifyTarget
	// goverter:enum:unknown NotifyTargetStatus_NOTIFY_TARGET_STATUS_UNSPECIFIED
	// goverter:enum:map NotifyTargetStatusUnspecified NotifyTargetStatus_NOTIFY_TARGET_STATUS_UNSPECIFIED
	// goverter:enum:map NotifyTargetStatusActive NotifyTargetStatus_NOTIFY_TARGET_STATUS_ACTIVE
	// goverter:enum:map NotifyTargetStatusSuspend NotifyTargetStatus_NOTIFY_TARGET_STATUS_SUSPEND
	ToPBNotifyTargetStatus(modelnetzach.NotifyTargetStatus) pb.NotifyTargetStatus

	// goverter:matchIgnoreCase
	ToPBNotifyFlow(*modelnetzach.NotifyFlow) *pb.NotifyFlow
	// goverter:enum:unknown NotifyFlowStatus_NOTIFY_FLOW_STATUS_UNSPECIFIED
	// goverter:enum:map NotifyFlowStatusUnspecified NotifyFlowStatus_NOTIFY_FLOW_STATUS_UNSPECIFIED
	// goverter:enum:map NotifyFlowStatusActive NotifyFlowStatus_NOTIFY_FLOW_STATUS_ACTIVE
	// goverter:enum:map NotifyFlowStatusSuspend NotifyFlowStatus_NOTIFY_FLOW_STATUS_SUSPEND
	ToPBNotifyFlowStatus(modelnetzach.NotifyFlowStatus) pb.NotifyFlowStatus
	// goverter:matchIgnoreCase
	// goverter:ignore Source
	ToPBNotifyFlowSource(*modelnetzach.NotifyFlowSource) *pb.NotifyFlowSource
	// goverter:matchIgnoreCase
	ToPBNotifyFlowTarget(*modelnetzach.NotifyFlowTarget) *pb.NotifyFlowTarget
	ToPBNotifyFlowList([]*modelnetzach.NotifyFlow) []*pb.NotifyFlow
}

func DurationPBToDuration(t *durationpb.Duration) time.Duration {
	if t == nil {
		return time.Duration(0)
	}
	return t.AsDuration()
}

func ToPBInternalID(id model.InternalID) *librarian.InternalID {
	return &librarian.InternalID{Id: int64(id)}
}

func ToPBTime(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

func ToPBTimePtr(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

func ToPBDuration(d time.Duration) *durationpb.Duration {
	return durationpb.New(d)
}

func ToPBPorterStatus(s modeltiphereth.PorterInstanceStatus) pb.UserStatus {
	switch s {
	case modeltiphereth.PorterInstanceStatusUnspecified:
		return pb.UserStatus_USER_STATUS_UNSPECIFIED
	case modeltiphereth.PorterInstanceStatusActive:
		return pb.UserStatus_USER_STATUS_ACTIVE
	case modeltiphereth.PorterInstanceStatusBlocked:
		return pb.UserStatus_USER_STATUS_BLOCKED
	default:
		return pb.UserStatus_USER_STATUS_UNSPECIFIED
	}
}

func ToPBFeedConfigPullStatus(s modelyesod.FeedConfigPullStatus) *pb.FeedConfigPullStatus {
	var status pb.FeedConfigPullStatus
	switch s {
	case modelyesod.FeedConfigPullStatusUnspecified:
		status = pb.FeedConfigPullStatus_FEED_CONFIG_PULL_STATUS_UNSPECIFIED
	case modelyesod.FeedConfigPullStatusProcessing:
		status = pb.FeedConfigPullStatus_FEED_CONFIG_PULL_STATUS_PROCESSING
	case modelyesod.FeedConfigPullStatusSuccess:
		status = pb.FeedConfigPullStatus_FEED_CONFIG_PULL_STATUS_SUCCESS
	case modelyesod.FeedConfigPullStatusFailed:
		status = pb.FeedConfigPullStatus_FEED_CONFIG_PULL_STATUS_FAILED
	default:
		status = pb.FeedConfigPullStatus_FEED_CONFIG_PULL_STATUS_UNSPECIFIED
	}
	return &status
}
