// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package converter

import (
	ent "github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent"
	apppackage "github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/apppackage"
	feedconfig "github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedconfig"
	notifytarget "github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifytarget"
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
		modelgeburaApp.ID = c.modelInternalIDToModelInternalID((*source).ID)
		modelgeburaApp.Internal = (*source).Internal
		modelgeburaApp.Source = (*source).Source
		modelgeburaApp.SourceAppID = (*source).SourceAppID
		modelgeburaApp.SourceURL = (*source).SourceURL
		modelgeburaApp.Name = (*source).Name
		modelgeburaApp.Type = ToBizAppType((*source).Type)
		modelgeburaApp.ShortDescription = (*source).ShortDescription
		modelgeburaApp.IconImageURL = (*source).IconImageURL
		modelgeburaApp.BackgroundImageURL = (*source).BackgroundImageURL
		modelgeburaApp.CoverImageURL = (*source).CoverImageURL
		modelgeburaApp.Details = c.entAppToPModelgeburaAppDetails((*source))
		modelgeburaApp.LatestUpdateTime = TimeToTime((*source).UpdatedAt)
		pModelgeburaApp = &modelgeburaApp
	}
	return pModelgeburaApp
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
func (c *toBizConverterImpl) ToBizAppPackage(source *ent.AppPackage) *modelgebura.AppPackage {
	var pModelgeburaAppPackage *modelgebura.AppPackage
	if source != nil {
		var modelgeburaAppPackage modelgebura.AppPackage
		modelgeburaAppPackage.ID = c.modelInternalIDToModelInternalID((*source).ID)
		modelgeburaAppPackage.Source = ToBizAppPackageSource((*source).Source)
		modelgeburaAppPackage.SourceID = c.modelInternalIDToModelInternalID((*source).SourceID)
		modelgeburaAppPackage.Name = (*source).Name
		modelgeburaAppPackage.Description = (*source).Description
		modelgeburaAppPackage.Binary = c.entAppPackageToPModelgeburaAppPackageBinary((*source))
		modelgeburaAppPackage.Public = (*source).Public
		pModelgeburaAppPackage = &modelgeburaAppPackage
	}
	return pModelgeburaAppPackage
}
func (c *toBizConverterImpl) ToBizAppPackageBinary(source ent.AppPackage) modelgebura.AppPackageBinary {
	var modelgeburaAppPackageBinary modelgebura.AppPackageBinary
	modelgeburaAppPackageBinary.Name = source.BinaryName
	modelgeburaAppPackageBinary.SizeBytes = source.BinarySizeBytes
	modelgeburaAppPackageBinary.PublicURL = source.BinaryPublicURL
	var byteList []uint8
	if source.BinarySha256 != nil {
		byteList = make([]uint8, len(source.BinarySha256))
		for i := 0; i < len(source.BinarySha256); i++ {
			byteList[i] = source.BinarySha256[i]
		}
	}
	modelgeburaAppPackageBinary.Sha256 = byteList
	return modelgeburaAppPackageBinary
}
func (c *toBizConverterImpl) ToBizAppPackageList(source []*ent.AppPackage) []*modelgebura.AppPackage {
	var pModelgeburaAppPackageList []*modelgebura.AppPackage
	if source != nil {
		pModelgeburaAppPackageList = make([]*modelgebura.AppPackage, len(source))
		for i := 0; i < len(source); i++ {
			pModelgeburaAppPackageList[i] = c.ToBizAppPackage(source[i])
		}
	}
	return pModelgeburaAppPackageList
}
func (c *toBizConverterImpl) ToBizDeviceInfo(source *ent.DeviceInfo) *modeltiphereth.DeviceInfo {
	var pModeltipherethDeviceInfo *modeltiphereth.DeviceInfo
	if source != nil {
		var modeltipherethDeviceInfo modeltiphereth.DeviceInfo
		modeltipherethDeviceInfo.ID = c.modelInternalIDToModelInternalID((*source).ID)
		modeltipherethDeviceInfo.DeviceModel = (*source).DeviceModel
		modeltipherethDeviceInfo.SystemVersion = (*source).SystemVersion
		modeltipherethDeviceInfo.ClientName = (*source).ClientName
		modeltipherethDeviceInfo.ClientSourceCodeAddress = (*source).ClientSourceCodeAddress
		modeltipherethDeviceInfo.ClientVersion = (*source).ClientVersion
		pModeltipherethDeviceInfo = &modeltipherethDeviceInfo
	}
	return pModeltipherethDeviceInfo
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
		modelyesodFeedConfig.Status = ToBizFeedConfigStatus((*source).Status)
		modelyesodFeedConfig.PullInterval = time.Duration((*source).PullInterval)
		modelyesodFeedConfig.LatestUpdateTime = TimeToTime((*source).LatestPullAt)
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
		modelchesedImage.Status = ToBizImageStatus((*source).Status)
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
func (c *toBizConverterImpl) ToBizNotifyFlow(source *ent.NotifyFlow) *modelnetzach.NotifyFlow {
	var pModelnetzachNotifyFlow *modelnetzach.NotifyFlow
	if source != nil {
		var modelnetzachNotifyFlow modelnetzach.NotifyFlow
		modelnetzachNotifyFlow.ID = c.modelInternalIDToModelInternalID((*source).ID)
		modelnetzachNotifyFlow.Name = (*source).Name
		modelnetzachNotifyFlow.Description = (*source).Description
		modelnetzachNotifyFlow.Status = ToBizNotifyFlowStatus((*source).Status)
		pModelnetzachNotifyFlow = &modelnetzachNotifyFlow
	}
	return pModelnetzachNotifyFlow
}
func (c *toBizConverterImpl) ToBizNotifyTarget(source *ent.NotifyTarget) *modelnetzach.NotifyTarget {
	var pModelnetzachNotifyTarget *modelnetzach.NotifyTarget
	if source != nil {
		var modelnetzachNotifyTarget modelnetzach.NotifyTarget
		modelnetzachNotifyTarget.ID = c.modelInternalIDToModelInternalID((*source).ID)
		modelnetzachNotifyTarget.Name = (*source).Name
		modelnetzachNotifyTarget.Description = (*source).Description
		modelnetzachNotifyTarget.Destination = (*source).Destination
		modelnetzachNotifyTarget.Status = ToBizNotifyTargetStatus((*source).Status)
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
		modeltipherethPorterInstance.Status = ToBizPorterStatus((*source).Status)
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
func (c *toBizConverterImpl) ToBizUser(source *ent.User) *modeltiphereth.User {
	var pModeltipherethUser *modeltiphereth.User
	if source != nil {
		var modeltipherethUser modeltiphereth.User
		modeltipherethUser.ID = c.modelInternalIDToModelInternalID((*source).ID)
		modeltipherethUser.UserName = (*source).Username
		modeltipherethUser.Type = ToLibAuthUserType((*source).Type)
		modeltipherethUser.Status = ToBizUserStatus((*source).Status)
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
func (c *toBizConverterImpl) entAppPackageToPModelgeburaAppPackageBinary(source ent.AppPackage) *modelgebura.AppPackageBinary {
	modelgeburaAppPackageBinary := c.ToBizAppPackageBinary(source)
	return &modelgeburaAppPackageBinary
}
func (c *toBizConverterImpl) entAppToPModelgeburaAppDetails(source ent.App) *modelgebura.AppDetails {
	var modelgeburaAppDetails modelgebura.AppDetails
	modelgeburaAppDetails.Description = source.Description
	modelgeburaAppDetails.ReleaseDate = source.ReleaseDate
	modelgeburaAppDetails.Developer = source.Developer
	modelgeburaAppDetails.Publisher = source.Publisher
	modelgeburaAppDetails.Version = source.Version
	return &modelgeburaAppDetails
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
		if (*source).SupportedAppSources != nil {
			stringList = make([]string, len((*source).SupportedAppSources))
			for j := 0; j < len((*source).SupportedAppSources); j++ {
				stringList[j] = (*source).SupportedAppSources[j]
			}
		}
		modeltipherethPorterFeatureSummary.SupportedAppSources = stringList
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

func (c *toEntConverterImpl) ToEntApp(source modelgebura.App) ent.App {
	var entApp ent.App
	entApp.ID = model.InternalID(source.ID)
	entApp.Internal = source.Internal
	entApp.Source = source.Source
	entApp.SourceAppID = source.SourceAppID
	entApp.SourceURL = source.SourceURL
	entApp.Name = source.Name
	entApp.Type = ToEntAppType(source.Type)
	entApp.ShortDescription = source.ShortDescription
	var pString *string
	if source.Details != nil {
		pString = &source.Details.Description
	}
	var xstring string
	if pString != nil {
		xstring = *pString
	}
	entApp.Description = xstring
	entApp.IconImageURL = source.IconImageURL
	entApp.BackgroundImageURL = source.BackgroundImageURL
	entApp.CoverImageURL = source.CoverImageURL
	var pString2 *string
	if source.Details != nil {
		pString2 = &source.Details.ReleaseDate
	}
	var xstring2 string
	if pString2 != nil {
		xstring2 = *pString2
	}
	entApp.ReleaseDate = xstring2
	var pString3 *string
	if source.Details != nil {
		pString3 = &source.Details.Developer
	}
	var xstring3 string
	if pString3 != nil {
		xstring3 = *pString3
	}
	entApp.Developer = xstring3
	var pString4 *string
	if source.Details != nil {
		pString4 = &source.Details.Publisher
	}
	var xstring4 string
	if pString4 != nil {
		xstring4 = *pString4
	}
	entApp.Publisher = xstring4
	var pString5 *string
	if source.Details != nil {
		pString5 = &source.Details.Version
	}
	var xstring5 string
	if pString5 != nil {
		xstring5 = *pString5
	}
	entApp.Version = xstring5
	return entApp
}
func (c *toEntConverterImpl) ToEntAppPackageSourceList(source []modelgebura.AppPackageSource) []apppackage.Source {
	var apppackageSourceList []apppackage.Source
	if source != nil {
		apppackageSourceList = make([]apppackage.Source, len(source))
		for i := 0; i < len(source); i++ {
			apppackageSourceList[i] = ToEntAppPackageSource(source[i])
		}
	}
	return apppackageSourceList
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
