// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package converter

import (
	ent "github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent"
	appinfo "github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/appinfo"
	deviceinfo "github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/deviceinfo"
	feedconfig "github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedconfig"
	image "github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/image"
	notifyflow "github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflow"
	notifytarget "github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifytarget"
	porterinstance "github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/porterinstance"
	user "github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	modelchesed "github.com/tuihub/librarian/app/sephirah/internal/model/modelchesed"
	modelgebura "github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	modelnetzach "github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	modeltiphereth "github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	modelyesod "github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	libauth "github.com/tuihub/librarian/internal/lib/libauth"
	model "github.com/tuihub/librarian/internal/model"
	modelfeed "github.com/tuihub/librarian/internal/model/modelfeed"
	"time"
)

type toBizConverterImpl struct{}

func (c *toBizConverterImpl) ToBizAccount(source *ent.Account) *modeltiphereth.Account {
	var pModeltipherethAccount *modeltiphereth.Account
	if source != nil {
		var modeltipherethAccount modeltiphereth.Account
		modeltipherethAccount.ID = model.InternalID((*source).ID)
		modeltipherethAccount.Platform = (*source).Platform
		modeltipherethAccount.PlatformAccountID = (*source).PlatformAccountID
		modeltipherethAccount.Name = (*source).Name
		modeltipherethAccount.ProfileURL = (*source).ProfileURL
		modeltipherethAccount.AvatarURL = (*source).AvatarURL
		modeltipherethAccount.LatestUpdateTime = TimeToTime((*source).UpdatedAt)
		pModeltipherethAccount = &modeltipherethAccount
	}
	return pModeltipherethAccount
}
func (c *toBizConverterImpl) ToBizAccountList(source []*ent.Account) []*modeltiphereth.Account {
	var pModeltipherethAccountList []*modeltiphereth.Account
	if source != nil {
		pModeltipherethAccountList = make([]*modeltiphereth.Account, len(source))
		for i := 0; i < len(source); i++ {
			pModeltipherethAccountList[i] = c.ToBizAccount(source[i])
		}
	}
	return pModeltipherethAccountList
}
func (c *toBizConverterImpl) ToBizApp(source *ent.App) *modelgebura.App {
	var pModelgeburaApp *modelgebura.App
	if source != nil {
		var modelgeburaApp modelgebura.App
		modelgeburaApp.ID = model.InternalID((*source).ID)
		modelgeburaApp.Name = (*source).Name
		modelgeburaApp.Description = (*source).Description
		modelgeburaApp.Public = (*source).Public
		pModelgeburaApp = &modelgeburaApp
	}
	return pModelgeburaApp
}
func (c *toBizConverterImpl) ToBizAppBinary(source ent.AppBinary) modelgebura.AppBinary {
	var modelgeburaAppBinary modelgebura.AppBinary
	modelgeburaAppBinary.Name = source.Name
	modelgeburaAppBinary.SizeBytes = source.SizeBytes
	modelgeburaAppBinary.PublicURL = source.PublicURL
	var byteList []uint8
	if source.Sha256 != nil {
		byteList = make([]uint8, len(source.Sha256))
		for i := 0; i < len(source.Sha256); i++ {
			byteList[i] = source.Sha256[i]
		}
	}
	modelgeburaAppBinary.Sha256 = byteList
	return modelgeburaAppBinary
}
func (c *toBizConverterImpl) ToBizAppInfo(source *ent.AppInfo) *modelgebura.AppInfo {
	var pModelgeburaAppInfo *modelgebura.AppInfo
	if source != nil {
		var modelgeburaAppInfo modelgebura.AppInfo
		modelgeburaAppInfo.ID = c.modelInternalIDToModelInternalID((*source).ID)
		modelgeburaAppInfo.Internal = (*source).Internal
		modelgeburaAppInfo.Source = (*source).Source
		modelgeburaAppInfo.SourceAppID = (*source).SourceAppID
		modelgeburaAppInfo.SourceURL = (*source).SourceURL
		modelgeburaAppInfo.Name = (*source).Name
		modelgeburaAppInfo.Type = c.ToBizAppType((*source).Type)
		modelgeburaAppInfo.ShortDescription = (*source).ShortDescription
		modelgeburaAppInfo.IconImageURL = (*source).IconImageURL
		modelgeburaAppInfo.BackgroundImageURL = (*source).BackgroundImageURL
		modelgeburaAppInfo.CoverImageURL = (*source).CoverImageURL
		modelgeburaAppInfo.Details = c.entAppInfoToPModelgeburaAppInfoDetails((*source))
		modelgeburaAppInfo.LatestUpdateTime = TimeToTime((*source).UpdatedAt)
		pModelgeburaAppInfo = &modelgeburaAppInfo
	}
	return pModelgeburaAppInfo
}
func (c *toBizConverterImpl) ToBizAppInfoList(source []*ent.AppInfo) []*modelgebura.AppInfo {
	var pModelgeburaAppInfoList []*modelgebura.AppInfo
	if source != nil {
		pModelgeburaAppInfoList = make([]*modelgebura.AppInfo, len(source))
		for i := 0; i < len(source); i++ {
			pModelgeburaAppInfoList[i] = c.ToBizAppInfo(source[i])
		}
	}
	return pModelgeburaAppInfoList
}
func (c *toBizConverterImpl) ToBizAppInst(source *ent.AppInst) *modelgebura.AppInst {
	var pModelgeburaAppInst *modelgebura.AppInst
	if source != nil {
		var modelgeburaAppInst modelgebura.AppInst
		modelgeburaAppInst.ID = c.modelInternalIDToModelInternalID((*source).ID)
		modelgeburaAppInst.AppID = c.modelInternalIDToModelInternalID((*source).AppID)
		modelgeburaAppInst.DeviceID = c.modelInternalIDToModelInternalID((*source).DeviceID)
		pModelgeburaAppInst = &modelgeburaAppInst
	}
	return pModelgeburaAppInst
}
func (c *toBizConverterImpl) ToBizAppInstList(source []*ent.AppInst) []*modelgebura.AppInst {
	var pModelgeburaAppInstList []*modelgebura.AppInst
	if source != nil {
		pModelgeburaAppInstList = make([]*modelgebura.AppInst, len(source))
		for i := 0; i < len(source); i++ {
			pModelgeburaAppInstList[i] = c.ToBizAppInst(source[i])
		}
	}
	return pModelgeburaAppInstList
}
func (c *toBizConverterImpl) ToBizAppList(source []*ent.App) []*modelgebura.App {
	var pModelgeburaAppList []*modelgebura.App
	if source != nil {
		pModelgeburaAppList = make([]*modelgebura.App, len(source))
		for i := 0; i < len(source); i++ {
			pModelgeburaAppList[i] = c.ToBizApp(source[i])
		}
	}
	return pModelgeburaAppList
}
func (c *toBizConverterImpl) ToBizAppType(source appinfo.Type) modelgebura.AppType {
	var modelgeburaAppType modelgebura.AppType
	switch source {
	case appinfo.TypeGame:
		modelgeburaAppType = modelgebura.AppTypeGame
	case appinfo.TypeUnknown:
		modelgeburaAppType = modelgebura.AppTypeUnspecified
	default:
		modelgeburaAppType = modelgebura.AppTypeUnspecified
	}
	return modelgeburaAppType
}
func (c *toBizConverterImpl) ToBizDeviceInfo(source *ent.DeviceInfo) *modeltiphereth.DeviceInfo {
	var pModeltipherethDeviceInfo *modeltiphereth.DeviceInfo
	if source != nil {
		var modeltipherethDeviceInfo modeltiphereth.DeviceInfo
		modeltipherethDeviceInfo.ID = c.modelInternalIDToModelInternalID((*source).ID)
		modeltipherethDeviceInfo.DeviceName = (*source).DeviceName
		modeltipherethDeviceInfo.SystemType = c.ToBizSystemType((*source).SystemType)
		modeltipherethDeviceInfo.SystemVersion = (*source).SystemVersion
		modeltipherethDeviceInfo.ClientName = (*source).ClientName
		modeltipherethDeviceInfo.ClientSourceCodeAddress = (*source).ClientSourceCodeAddress
		modeltipherethDeviceInfo.ClientVersion = (*source).ClientVersion
		pModeltipherethDeviceInfo = &modeltipherethDeviceInfo
	}
	return pModeltipherethDeviceInfo
}
func (c *toBizConverterImpl) ToBizDeviceInfoList(source []*ent.DeviceInfo) []*modeltiphereth.DeviceInfo {
	var pModeltipherethDeviceInfoList []*modeltiphereth.DeviceInfo
	if source != nil {
		pModeltipherethDeviceInfoList = make([]*modeltiphereth.DeviceInfo, len(source))
		for i := 0; i < len(source); i++ {
			pModeltipherethDeviceInfoList[i] = c.ToBizDeviceInfo(source[i])
		}
	}
	return pModeltipherethDeviceInfoList
}
func (c *toBizConverterImpl) ToBizFeed(source *ent.Feed) *modelfeed.Feed {
	var pModelfeedFeed *modelfeed.Feed
	if source != nil {
		var modelfeedFeed modelfeed.Feed
		modelfeedFeed.ID = c.modelInternalIDToModelInternalID((*source).ID)
		modelfeedFeed.Title = (*source).Title
		modelfeedFeed.Description = (*source).Description
		modelfeedFeed.Link = (*source).Link
		var pModelfeedPersonList []*modelfeed.Person
		if (*source).Authors != nil {
			pModelfeedPersonList = make([]*modelfeed.Person, len((*source).Authors))
			for i := 0; i < len((*source).Authors); i++ {
				pModelfeedPersonList[i] = c.pModelfeedPersonToPModelfeedPerson((*source).Authors[i])
			}
		}
		modelfeedFeed.Authors = pModelfeedPersonList
		modelfeedFeed.Language = (*source).Language
		modelfeedFeed.Image = c.pModelfeedImageToPModelfeedImage((*source).Image)
		pModelfeedFeed = &modelfeedFeed
	}
	return pModelfeedFeed
}
func (c *toBizConverterImpl) ToBizFeedConfig(source *ent.FeedConfig) *modelyesod.FeedConfig {
	var pModelyesodFeedConfig *modelyesod.FeedConfig
	if source != nil {
		var modelyesodFeedConfig modelyesod.FeedConfig
		modelyesodFeedConfig.ID = c.modelInternalIDToModelInternalID((*source).ID)
		modelyesodFeedConfig.Name = (*source).Name
		modelyesodFeedConfig.FeedURL = (*source).FeedURL
		modelyesodFeedConfig.Category = (*source).Category
		modelyesodFeedConfig.AuthorAccount = c.modelInternalIDToModelInternalID((*source).AuthorAccount)
		modelyesodFeedConfig.Source = (*source).Source
		modelyesodFeedConfig.Status = c.ToBizFeedConfigStatus((*source).Status)
		modelyesodFeedConfig.PullInterval = time.Duration((*source).PullInterval)
		modelyesodFeedConfig.LatestPullTime = TimeToTime((*source).LatestPullAt)
		modelyesodFeedConfig.LatestPullStatus = c.ToBizFeedConfigPullStatus((*source).LatestPullStatus)
		modelyesodFeedConfig.LatestPullMessage = (*source).LatestPullMessage
		modelyesodFeedConfig.HideItems = (*source).HideItems
		pModelyesodFeedConfig = &modelyesodFeedConfig
	}
	return pModelyesodFeedConfig
}
func (c *toBizConverterImpl) ToBizFeedConfigList(source []*ent.FeedConfig) []*modelyesod.FeedConfig {
	var pModelyesodFeedConfigList []*modelyesod.FeedConfig
	if source != nil {
		pModelyesodFeedConfigList = make([]*modelyesod.FeedConfig, len(source))
		for i := 0; i < len(source); i++ {
			pModelyesodFeedConfigList[i] = c.ToBizFeedConfig(source[i])
		}
	}
	return pModelyesodFeedConfigList
}
func (c *toBizConverterImpl) ToBizFeedConfigPullStatus(source feedconfig.LatestPullStatus) modelyesod.FeedConfigPullStatus {
	var modelyesodFeedConfigPullStatus modelyesod.FeedConfigPullStatus
	switch source {
	case feedconfig.LatestPullStatusFailed:
		modelyesodFeedConfigPullStatus = modelyesod.FeedConfigPullStatusFailed
	case feedconfig.LatestPullStatusProcessing:
		modelyesodFeedConfigPullStatus = modelyesod.FeedConfigPullStatusProcessing
	case feedconfig.LatestPullStatusSuccess:
		modelyesodFeedConfigPullStatus = modelyesod.FeedConfigPullStatusSuccess
	default:
		modelyesodFeedConfigPullStatus = modelyesod.FeedConfigPullStatusUnspecified
	}
	return modelyesodFeedConfigPullStatus
}
func (c *toBizConverterImpl) ToBizFeedConfigStatus(source feedconfig.Status) modelyesod.FeedConfigStatus {
	var modelyesodFeedConfigStatus modelyesod.FeedConfigStatus
	switch source {
	case feedconfig.StatusActive:
		modelyesodFeedConfigStatus = modelyesod.FeedConfigStatusActive
	case feedconfig.StatusSuspend:
		modelyesodFeedConfigStatus = modelyesod.FeedConfigStatusSuspend
	default:
		modelyesodFeedConfigStatus = modelyesod.FeedConfigStatusUnspecified
	}
	return modelyesodFeedConfigStatus
}
func (c *toBizConverterImpl) ToBizFeedItem(source *ent.FeedItem) *modelfeed.Item {
	var pModelfeedItem *modelfeed.Item
	if source != nil {
		var modelfeedItem modelfeed.Item
		modelfeedItem.ID = c.modelInternalIDToModelInternalID((*source).ID)
		modelfeedItem.Title = (*source).Title
		modelfeedItem.Description = (*source).Description
		modelfeedItem.Content = (*source).Content
		modelfeedItem.Link = (*source).Link
		modelfeedItem.Updated = (*source).Updated
		modelfeedItem.UpdatedParsed = TimeToTimePtr((*source).UpdatedParsed)
		modelfeedItem.Published = (*source).Published
		modelfeedItem.PublishedParsed = c.timeTimeToPTimeTime((*source).PublishedParsed)
		var pModelfeedPersonList []*modelfeed.Person
		if (*source).Authors != nil {
			pModelfeedPersonList = make([]*modelfeed.Person, len((*source).Authors))
			for i := 0; i < len((*source).Authors); i++ {
				pModelfeedPersonList[i] = c.pModelfeedPersonToPModelfeedPerson((*source).Authors[i])
			}
		}
		modelfeedItem.Authors = pModelfeedPersonList
		modelfeedItem.GUID = (*source).GUID
		modelfeedItem.Image = c.pModelfeedImageToPModelfeedImage((*source).Image)
		var pModelfeedEnclosureList []*modelfeed.Enclosure
		if (*source).Enclosures != nil {
			pModelfeedEnclosureList = make([]*modelfeed.Enclosure, len((*source).Enclosures))
			for j := 0; j < len((*source).Enclosures); j++ {
				pModelfeedEnclosureList[j] = c.pModelfeedEnclosureToPModelfeedEnclosure((*source).Enclosures[j])
			}
		}
		modelfeedItem.Enclosures = pModelfeedEnclosureList
		modelfeedItem.PublishPlatform = (*source).PublishPlatform
		modelfeedItem.ReadCount = (*source).ReadCount
		modelfeedItem.DigestDescription = (*source).DigestDescription
		var pModelfeedImageList []*modelfeed.Image
		if (*source).DigestImages != nil {
			pModelfeedImageList = make([]*modelfeed.Image, len((*source).DigestImages))
			for k := 0; k < len((*source).DigestImages); k++ {
				pModelfeedImageList[k] = c.pModelfeedImageToPModelfeedImage((*source).DigestImages[k])
			}
		}
		modelfeedItem.DigestImages = pModelfeedImageList
		pModelfeedItem = &modelfeedItem
	}
	return pModelfeedItem
}
func (c *toBizConverterImpl) ToBizFeedItemCollection(source *ent.FeedItemCollection) *modelyesod.FeedItemCollection {
	var pModelyesodFeedItemCollection *modelyesod.FeedItemCollection
	if source != nil {
		var modelyesodFeedItemCollection modelyesod.FeedItemCollection
		modelyesodFeedItemCollection.ID = c.modelInternalIDToModelInternalID((*source).ID)
		modelyesodFeedItemCollection.Name = (*source).Name
		modelyesodFeedItemCollection.Description = (*source).Description
		modelyesodFeedItemCollection.Category = (*source).Category
		pModelyesodFeedItemCollection = &modelyesodFeedItemCollection
	}
	return pModelyesodFeedItemCollection
}
func (c *toBizConverterImpl) ToBizFeedItemCollectionList(source []*ent.FeedItemCollection) []*modelyesod.FeedItemCollection {
	var pModelyesodFeedItemCollectionList []*modelyesod.FeedItemCollection
	if source != nil {
		pModelyesodFeedItemCollectionList = make([]*modelyesod.FeedItemCollection, len(source))
		for i := 0; i < len(source); i++ {
			pModelyesodFeedItemCollectionList[i] = c.ToBizFeedItemCollection(source[i])
		}
	}
	return pModelyesodFeedItemCollectionList
}
func (c *toBizConverterImpl) ToBizFeedItemList(source []*ent.FeedItem) []*modelfeed.Item {
	var pModelfeedItemList []*modelfeed.Item
	if source != nil {
		pModelfeedItemList = make([]*modelfeed.Item, len(source))
		for i := 0; i < len(source); i++ {
			pModelfeedItemList[i] = c.ToBizFeedItem(source[i])
		}
	}
	return pModelfeedItemList
}
func (c *toBizConverterImpl) ToBizImage(source *ent.Image) *modelchesed.Image {
	var pModelchesedImage *modelchesed.Image
	if source != nil {
		var modelchesedImage modelchesed.Image
		modelchesedImage.ID = c.modelInternalIDToModelInternalID((*source).ID)
		modelchesedImage.Name = (*source).Name
		modelchesedImage.Description = (*source).Description
		modelchesedImage.Status = c.ToBizImageStatus((*source).Status)
		pModelchesedImage = &modelchesedImage
	}
	return pModelchesedImage
}
func (c *toBizConverterImpl) ToBizImageList(source []*ent.Image) []*modelchesed.Image {
	var pModelchesedImageList []*modelchesed.Image
	if source != nil {
		pModelchesedImageList = make([]*modelchesed.Image, len(source))
		for i := 0; i < len(source); i++ {
			pModelchesedImageList[i] = c.ToBizImage(source[i])
		}
	}
	return pModelchesedImageList
}
func (c *toBizConverterImpl) ToBizImageStatus(source image.Status) modelchesed.ImageStatus {
	var modelchesedImageStatus modelchesed.ImageStatus
	switch source {
	case image.StatusScanned:
		modelchesedImageStatus = modelchesed.ImageStatusScanned
	case image.StatusUploaded:
		modelchesedImageStatus = modelchesed.ImageStatusUploaded
	default:
		modelchesedImageStatus = modelchesed.ImageStatusUnspecified
	}
	return modelchesedImageStatus
}
func (c *toBizConverterImpl) ToBizNotifyFlow(source *ent.NotifyFlow) *modelnetzach.NotifyFlow {
	var pModelnetzachNotifyFlow *modelnetzach.NotifyFlow
	if source != nil {
		var modelnetzachNotifyFlow modelnetzach.NotifyFlow
		modelnetzachNotifyFlow.ID = c.modelInternalIDToModelInternalID((*source).ID)
		modelnetzachNotifyFlow.Name = (*source).Name
		modelnetzachNotifyFlow.Description = (*source).Description
		modelnetzachNotifyFlow.Status = c.ToBizNotifyFlowStatus((*source).Status)
		pModelnetzachNotifyFlow = &modelnetzachNotifyFlow
	}
	return pModelnetzachNotifyFlow
}
func (c *toBizConverterImpl) ToBizNotifyFlowStatus(source notifyflow.Status) modelnetzach.NotifyFlowStatus {
	var modelnetzachNotifyFlowStatus modelnetzach.NotifyFlowStatus
	switch source {
	case notifyflow.StatusActive:
		modelnetzachNotifyFlowStatus = modelnetzach.NotifyFlowStatusActive
	case notifyflow.StatusSuspend:
		modelnetzachNotifyFlowStatus = modelnetzach.NotifyFlowStatusSuspend
	default:
		modelnetzachNotifyFlowStatus = modelnetzach.NotifyFlowStatusUnspecified
	}
	return modelnetzachNotifyFlowStatus
}
func (c *toBizConverterImpl) ToBizNotifyTarget(source *ent.NotifyTarget) *modelnetzach.NotifyTarget {
	var pModelnetzachNotifyTarget *modelnetzach.NotifyTarget
	if source != nil {
		var modelnetzachNotifyTarget modelnetzach.NotifyTarget
		modelnetzachNotifyTarget.ID = c.modelInternalIDToModelInternalID((*source).ID)
		modelnetzachNotifyTarget.Name = (*source).Name
		modelnetzachNotifyTarget.Description = (*source).Description
		modelnetzachNotifyTarget.Destination = (*source).Destination
		modelnetzachNotifyTarget.Status = c.ToBizNotifyTargetStatus((*source).Status)
		modelnetzachNotifyTarget.Token = (*source).Token
		pModelnetzachNotifyTarget = &modelnetzachNotifyTarget
	}
	return pModelnetzachNotifyTarget
}
func (c *toBizConverterImpl) ToBizNotifyTargetList(source []*ent.NotifyTarget) []*modelnetzach.NotifyTarget {
	var pModelnetzachNotifyTargetList []*modelnetzach.NotifyTarget
	if source != nil {
		pModelnetzachNotifyTargetList = make([]*modelnetzach.NotifyTarget, len(source))
		for i := 0; i < len(source); i++ {
			pModelnetzachNotifyTargetList[i] = c.ToBizNotifyTarget(source[i])
		}
	}
	return pModelnetzachNotifyTargetList
}
func (c *toBizConverterImpl) ToBizNotifyTargetStatus(source notifytarget.Status) modelnetzach.NotifyTargetStatus {
	var modelnetzachNotifyTargetStatus modelnetzach.NotifyTargetStatus
	switch source {
	case notifytarget.StatusActive:
		modelnetzachNotifyTargetStatus = modelnetzach.NotifyTargetStatusActive
	case notifytarget.StatusSuspend:
		modelnetzachNotifyTargetStatus = modelnetzach.NotifyTargetStatusSuspend
	default:
		modelnetzachNotifyTargetStatus = modelnetzach.NotifyTargetStatusUnspecified
	}
	return modelnetzachNotifyTargetStatus
}
func (c *toBizConverterImpl) ToBizPorter(source *ent.PorterInstance) *modeltiphereth.PorterInstance {
	var pModeltipherethPorterInstance *modeltiphereth.PorterInstance
	if source != nil {
		var modeltipherethPorterInstance modeltiphereth.PorterInstance
		modeltipherethPorterInstance.ID = c.modelInternalIDToModelInternalID((*source).ID)
		modeltipherethPorterInstance.Name = (*source).Name
		modeltipherethPorterInstance.Version = (*source).Version
		modeltipherethPorterInstance.GlobalName = (*source).GlobalName
		modeltipherethPorterInstance.Address = (*source).Address
		modeltipherethPorterInstance.FeatureSummary = c.pModeltipherethPorterFeatureSummaryToPModeltipherethPorterFeatureSummary((*source).FeatureSummary)
		modeltipherethPorterInstance.Status = c.ToBizPorterStatus((*source).Status)
		pModeltipherethPorterInstance = &modeltipherethPorterInstance
	}
	return pModeltipherethPorterInstance
}
func (c *toBizConverterImpl) ToBizPorterList(source []*ent.PorterInstance) []*modeltiphereth.PorterInstance {
	var pModeltipherethPorterInstanceList []*modeltiphereth.PorterInstance
	if source != nil {
		pModeltipherethPorterInstanceList = make([]*modeltiphereth.PorterInstance, len(source))
		for i := 0; i < len(source); i++ {
			pModeltipherethPorterInstanceList[i] = c.ToBizPorter(source[i])
		}
	}
	return pModeltipherethPorterInstanceList
}
func (c *toBizConverterImpl) ToBizPorterStatus(source porterinstance.Status) modeltiphereth.PorterInstanceStatus {
	var modeltipherethPorterInstanceStatus modeltiphereth.PorterInstanceStatus
	switch source {
	case porterinstance.StatusActive:
		modeltipherethPorterInstanceStatus = modeltiphereth.PorterInstanceStatusActive
	case porterinstance.StatusBlocked:
		modeltipherethPorterInstanceStatus = modeltiphereth.PorterInstanceStatusBlocked
	default:
		modeltipherethPorterInstanceStatus = modeltiphereth.PorterInstanceStatusUnspecified
	}
	return modeltipherethPorterInstanceStatus
}
func (c *toBizConverterImpl) ToBizSystemType(source deviceinfo.SystemType) modeltiphereth.SystemType {
	var modeltipherethSystemType modeltiphereth.SystemType
	switch source {
	case deviceinfo.SystemTypeAndroid:
		modeltipherethSystemType = modeltiphereth.SystemTypeAndroid
	case deviceinfo.SystemTypeIos:
		modeltipherethSystemType = modeltiphereth.SystemTypeIOS
	case deviceinfo.SystemTypeLinux:
		modeltipherethSystemType = modeltiphereth.SystemTypeLinux
	case deviceinfo.SystemTypeMacos:
		modeltipherethSystemType = modeltiphereth.SystemTypeMacOS
	case deviceinfo.SystemTypeUnknown:
		modeltipherethSystemType = modeltiphereth.SystemTypeUnspecified
	case deviceinfo.SystemTypeWeb:
		modeltipherethSystemType = modeltiphereth.SystemTypeWeb
	case deviceinfo.SystemTypeWindows:
		modeltipherethSystemType = modeltiphereth.SystemTypeWindows
	default:
		modeltipherethSystemType = modeltiphereth.SystemTypeUnspecified
	}
	return modeltipherethSystemType
}
func (c *toBizConverterImpl) ToBizUser(source *ent.User) *modeltiphereth.User {
	var pModeltipherethUser *modeltiphereth.User
	if source != nil {
		var modeltipherethUser modeltiphereth.User
		modeltipherethUser.ID = c.modelInternalIDToModelInternalID((*source).ID)
		modeltipherethUser.UserName = (*source).Username
		modeltipherethUser.Type = c.ToLibAuthUserType((*source).Type)
		modeltipherethUser.Status = c.ToBizUserStatus((*source).Status)
		pModeltipherethUser = &modeltipherethUser
	}
	return pModeltipherethUser
}
func (c *toBizConverterImpl) ToBizUserList(source []*ent.User) []*modeltiphereth.User {
	var pModeltipherethUserList []*modeltiphereth.User
	if source != nil {
		pModeltipherethUserList = make([]*modeltiphereth.User, len(source))
		for i := 0; i < len(source); i++ {
			pModeltipherethUserList[i] = c.ToBizUser(source[i])
		}
	}
	return pModeltipherethUserList
}
func (c *toBizConverterImpl) ToBizUserSession(source *ent.UserSession) *modeltiphereth.UserSession {
	var pModeltipherethUserSession *modeltiphereth.UserSession
	if source != nil {
		var modeltipherethUserSession modeltiphereth.UserSession
		modeltipherethUserSession.ID = c.modelInternalIDToModelInternalID((*source).ID)
		modeltipherethUserSession.UserID = c.modelInternalIDToModelInternalID((*source).UserID)
		modeltipherethUserSession.RefreshToken = (*source).RefreshToken
		modeltipherethUserSession.CreateAt = TimeToTime((*source).CreatedAt)
		modeltipherethUserSession.ExpireAt = TimeToTime((*source).ExpireAt)
		pModeltipherethUserSession = &modeltipherethUserSession
	}
	return pModeltipherethUserSession
}
func (c *toBizConverterImpl) ToBizUserSessionList(source []*ent.UserSession) []*modeltiphereth.UserSession {
	var pModeltipherethUserSessionList []*modeltiphereth.UserSession
	if source != nil {
		pModeltipherethUserSessionList = make([]*modeltiphereth.UserSession, len(source))
		for i := 0; i < len(source); i++ {
			pModeltipherethUserSessionList[i] = c.ToBizUserSession(source[i])
		}
	}
	return pModeltipherethUserSessionList
}
func (c *toBizConverterImpl) ToBizUserStatus(source user.Status) modeltiphereth.UserStatus {
	var modeltipherethUserStatus modeltiphereth.UserStatus
	switch source {
	case user.StatusActive:
		modeltipherethUserStatus = modeltiphereth.UserStatusActive
	case user.StatusBlocked:
		modeltipherethUserStatus = modeltiphereth.UserStatusBlocked
	default:
		modeltipherethUserStatus = modeltiphereth.UserStatusUnspecified
	}
	return modeltipherethUserStatus
}
func (c *toBizConverterImpl) ToLibAuthUserType(source user.Type) libauth.UserType {
	var libauthUserType libauth.UserType
	switch source {
	case user.TypeAdmin:
		libauthUserType = libauth.UserTypeAdmin
	case user.TypeNormal:
		libauthUserType = libauth.UserTypeNormal
	case user.TypeSentinel:
		libauthUserType = libauth.UserTypeSentinel
	default:
		libauthUserType = libauth.UserTypeUnspecified
	}
	return libauthUserType
}
func (c *toBizConverterImpl) entAppInfoToPModelgeburaAppInfoDetails(source ent.AppInfo) *modelgebura.AppInfoDetails {
	var modelgeburaAppInfoDetails modelgebura.AppInfoDetails
	modelgeburaAppInfoDetails.Description = source.Description
	modelgeburaAppInfoDetails.ReleaseDate = source.ReleaseDate
	modelgeburaAppInfoDetails.Developer = source.Developer
	modelgeburaAppInfoDetails.Publisher = source.Publisher
	modelgeburaAppInfoDetails.Version = source.Version
	return &modelgeburaAppInfoDetails
}
func (c *toBizConverterImpl) modelInternalIDToModelInternalID(source model.InternalID) model.InternalID {
	return model.InternalID(source)
}
func (c *toBizConverterImpl) pModelfeedEnclosureToPModelfeedEnclosure(source *modelfeed.Enclosure) *modelfeed.Enclosure {
	var pModelfeedEnclosure *modelfeed.Enclosure
	if source != nil {
		var modelfeedEnclosure modelfeed.Enclosure
		modelfeedEnclosure.URL = (*source).URL
		modelfeedEnclosure.Length = (*source).Length
		modelfeedEnclosure.Type = (*source).Type
		pModelfeedEnclosure = &modelfeedEnclosure
	}
	return pModelfeedEnclosure
}
func (c *toBizConverterImpl) pModelfeedImageToPModelfeedImage(source *modelfeed.Image) *modelfeed.Image {
	var pModelfeedImage *modelfeed.Image
	if source != nil {
		var modelfeedImage modelfeed.Image
		modelfeedImage.URL = (*source).URL
		modelfeedImage.Title = (*source).Title
		pModelfeedImage = &modelfeedImage
	}
	return pModelfeedImage
}
func (c *toBizConverterImpl) pModelfeedPersonToPModelfeedPerson(source *modelfeed.Person) *modelfeed.Person {
	var pModelfeedPerson *modelfeed.Person
	if source != nil {
		var modelfeedPerson modelfeed.Person
		modelfeedPerson.Name = (*source).Name
		modelfeedPerson.Email = (*source).Email
		pModelfeedPerson = &modelfeedPerson
	}
	return pModelfeedPerson
}
func (c *toBizConverterImpl) pModeltipherethPorterFeatureSummaryToPModeltipherethPorterFeatureSummary(source *modeltiphereth.PorterFeatureSummary) *modeltiphereth.PorterFeatureSummary {
	var pModeltipherethPorterFeatureSummary *modeltiphereth.PorterFeatureSummary
	if source != nil {
		var modeltipherethPorterFeatureSummary modeltiphereth.PorterFeatureSummary
		var pModeltipherethSupportedAccountList []*modeltiphereth.SupportedAccount
		if (*source).SupportedAccounts != nil {
			pModeltipherethSupportedAccountList = make([]*modeltiphereth.SupportedAccount, len((*source).SupportedAccounts))
			for i := 0; i < len((*source).SupportedAccounts); i++ {
				pModeltipherethSupportedAccountList[i] = c.pModeltipherethSupportedAccountToPModeltipherethSupportedAccount((*source).SupportedAccounts[i])
			}
		}
		modeltipherethPorterFeatureSummary.SupportedAccounts = pModeltipherethSupportedAccountList
		var stringList []string
		if (*source).SupportedAppInfoSources != nil {
			stringList = make([]string, len((*source).SupportedAppInfoSources))
			for j := 0; j < len((*source).SupportedAppInfoSources); j++ {
				stringList[j] = (*source).SupportedAppInfoSources[j]
			}
		}
		modeltipherethPorterFeatureSummary.SupportedAppInfoSources = stringList
		var stringList2 []string
		if (*source).SupportedFeedSources != nil {
			stringList2 = make([]string, len((*source).SupportedFeedSources))
			for k := 0; k < len((*source).SupportedFeedSources); k++ {
				stringList2[k] = (*source).SupportedFeedSources[k]
			}
		}
		modeltipherethPorterFeatureSummary.SupportedFeedSources = stringList2
		var stringList3 []string
		if (*source).SupportedNotifyDestinations != nil {
			stringList3 = make([]string, len((*source).SupportedNotifyDestinations))
			for l := 0; l < len((*source).SupportedNotifyDestinations); l++ {
				stringList3[l] = (*source).SupportedNotifyDestinations[l]
			}
		}
		modeltipherethPorterFeatureSummary.SupportedNotifyDestinations = stringList3
		pModeltipherethPorterFeatureSummary = &modeltipherethPorterFeatureSummary
	}
	return pModeltipherethPorterFeatureSummary
}
func (c *toBizConverterImpl) pModeltipherethSupportedAccountToPModeltipherethSupportedAccount(source *modeltiphereth.SupportedAccount) *modeltiphereth.SupportedAccount {
	var pModeltipherethSupportedAccount *modeltiphereth.SupportedAccount
	if source != nil {
		var modeltipherethSupportedAccount modeltiphereth.SupportedAccount
		modeltipherethSupportedAccount.Platform = (*source).Platform
		var modelAccountAppRelationTypeList []model.AccountAppRelationType
		if (*source).AppRelationTypes != nil {
			modelAccountAppRelationTypeList = make([]model.AccountAppRelationType, len((*source).AppRelationTypes))
			for i := 0; i < len((*source).AppRelationTypes); i++ {
				modelAccountAppRelationTypeList[i] = model.AccountAppRelationType((*source).AppRelationTypes[i])
			}
		}
		modeltipherethSupportedAccount.AppRelationTypes = modelAccountAppRelationTypeList
		pModeltipherethSupportedAccount = &modeltipherethSupportedAccount
	}
	return pModeltipherethSupportedAccount
}
func (c *toBizConverterImpl) timeTimeToPTimeTime(source time.Time) *time.Time {
	timeTime := TimeToTime(source)
	return &timeTime
}

type toEntConverterImpl struct{}

func (c *toEntConverterImpl) ToEntAppInfo(source modelgebura.AppInfo) ent.AppInfo {
	var entAppInfo ent.AppInfo
	entAppInfo.ID = model.InternalID(source.ID)
	entAppInfo.Internal = source.Internal
	entAppInfo.Source = source.Source
	entAppInfo.SourceAppID = source.SourceAppID
	entAppInfo.SourceURL = source.SourceURL
	entAppInfo.Name = source.Name
	entAppInfo.Type = ToEntAppType(source.Type)
	entAppInfo.ShortDescription = source.ShortDescription
	var pString *string
	if source.Details != nil {
		pString = &source.Details.Description
	}
	var xstring string
	if pString != nil {
		xstring = *pString
	}
	entAppInfo.Description = xstring
	entAppInfo.IconImageURL = source.IconImageURL
	entAppInfo.BackgroundImageURL = source.BackgroundImageURL
	entAppInfo.CoverImageURL = source.CoverImageURL
	var pString2 *string
	if source.Details != nil {
		pString2 = &source.Details.ReleaseDate
	}
	var xstring2 string
	if pString2 != nil {
		xstring2 = *pString2
	}
	entAppInfo.ReleaseDate = xstring2
	var pString3 *string
	if source.Details != nil {
		pString3 = &source.Details.Developer
	}
	var xstring3 string
	if pString3 != nil {
		xstring3 = *pString3
	}
	entAppInfo.Developer = xstring3
	var pString4 *string
	if source.Details != nil {
		pString4 = &source.Details.Publisher
	}
	var xstring4 string
	if pString4 != nil {
		xstring4 = *pString4
	}
	entAppInfo.Publisher = xstring4
	var pString5 *string
	if source.Details != nil {
		pString5 = &source.Details.Version
	}
	var xstring5 string
	if pString5 != nil {
		xstring5 = *pString5
	}
	entAppInfo.Version = xstring5
	return entAppInfo
}
func (c *toEntConverterImpl) ToEntFeedConfigStatusList(source []modelyesod.FeedConfigStatus) []feedconfig.Status {
	var feedconfigStatusList []feedconfig.Status
	if source != nil {
		feedconfigStatusList = make([]feedconfig.Status, len(source))
		for i := 0; i < len(source); i++ {
			feedconfigStatusList[i] = ToEntFeedConfigStatus(source[i])
		}
	}
	return feedconfigStatusList
}
func (c *toEntConverterImpl) ToEntNotifyTargetStatusList(source []modelnetzach.NotifyTargetStatus) []notifytarget.Status {
	var notifytargetStatusList []notifytarget.Status
	if source != nil {
		notifytargetStatusList = make([]notifytarget.Status, len(source))
		for i := 0; i < len(source); i++ {
			notifytargetStatusList[i] = ToEntNotifyTargetStatus(source[i])
		}
	}
	return notifytargetStatusList
}
func (c *toEntConverterImpl) ToEntUserStatusList(source []modeltiphereth.UserStatus) []user.Status {
	var userStatusList []user.Status
	if source != nil {
		userStatusList = make([]user.Status, len(source))
		for i := 0; i < len(source); i++ {
			userStatusList[i] = ToEntUserStatus(source[i])
		}
	}
	return userStatusList
}
func (c *toEntConverterImpl) ToEntUserTypeList(source []libauth.UserType) []user.Type {
	var userTypeList []user.Type
	if source != nil {
		userTypeList = make([]user.Type, len(source))
		for i := 0; i < len(source); i++ {
			userTypeList[i] = ToEntUserType(source[i])
		}
	}
	return userTypeList
}
