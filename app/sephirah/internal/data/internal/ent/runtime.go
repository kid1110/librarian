// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/account"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/app"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/apppackage"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/apppackageruntime"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feed"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feeditem"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/file"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/image"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflow"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflowsource"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflowtarget"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifytarget"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/schema"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescUpdatedAt is the schema descriptor for updated_at field.
	accountDescUpdatedAt := accountFields[6].Descriptor()
	// account.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	account.DefaultUpdatedAt = accountDescUpdatedAt.Default.(func() time.Time)
	// account.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	account.UpdateDefaultUpdatedAt = accountDescUpdatedAt.UpdateDefault.(func() time.Time)
	// accountDescCreatedAt is the schema descriptor for created_at field.
	accountDescCreatedAt := accountFields[7].Descriptor()
	// account.DefaultCreatedAt holds the default value on creation for the created_at field.
	account.DefaultCreatedAt = accountDescCreatedAt.Default.(func() time.Time)
	appFields := schema.App{}.Fields()
	_ = appFields
	// appDescUpdatedAt is the schema descriptor for updated_at field.
	appDescUpdatedAt := appFields[15].Descriptor()
	// app.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	app.DefaultUpdatedAt = appDescUpdatedAt.Default.(func() time.Time)
	// app.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	app.UpdateDefaultUpdatedAt = appDescUpdatedAt.UpdateDefault.(func() time.Time)
	// appDescCreatedAt is the schema descriptor for created_at field.
	appDescCreatedAt := appFields[16].Descriptor()
	// app.DefaultCreatedAt holds the default value on creation for the created_at field.
	app.DefaultCreatedAt = appDescCreatedAt.Default.(func() time.Time)
	apppackageFields := schema.AppPackage{}.Fields()
	_ = apppackageFields
	// apppackageDescUpdatedAt is the schema descriptor for updated_at field.
	apppackageDescUpdatedAt := apppackageFields[10].Descriptor()
	// apppackage.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	apppackage.DefaultUpdatedAt = apppackageDescUpdatedAt.Default.(func() time.Time)
	// apppackage.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	apppackage.UpdateDefaultUpdatedAt = apppackageDescUpdatedAt.UpdateDefault.(func() time.Time)
	// apppackageDescCreatedAt is the schema descriptor for created_at field.
	apppackageDescCreatedAt := apppackageFields[11].Descriptor()
	// apppackage.DefaultCreatedAt holds the default value on creation for the created_at field.
	apppackage.DefaultCreatedAt = apppackageDescCreatedAt.Default.(func() time.Time)
	apppackageruntimeFields := schema.AppPackageRunTime{}.Fields()
	_ = apppackageruntimeFields
	// apppackageruntimeDescUpdatedAt is the schema descriptor for updated_at field.
	apppackageruntimeDescUpdatedAt := apppackageruntimeFields[4].Descriptor()
	// apppackageruntime.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	apppackageruntime.DefaultUpdatedAt = apppackageruntimeDescUpdatedAt.Default.(func() time.Time)
	// apppackageruntime.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	apppackageruntime.UpdateDefaultUpdatedAt = apppackageruntimeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// apppackageruntimeDescCreatedAt is the schema descriptor for created_at field.
	apppackageruntimeDescCreatedAt := apppackageruntimeFields[5].Descriptor()
	// apppackageruntime.DefaultCreatedAt holds the default value on creation for the created_at field.
	apppackageruntime.DefaultCreatedAt = apppackageruntimeDescCreatedAt.Default.(func() time.Time)
	feedFields := schema.Feed{}.Fields()
	_ = feedFields
	// feedDescUpdatedAt is the schema descriptor for updated_at field.
	feedDescUpdatedAt := feedFields[7].Descriptor()
	// feed.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	feed.DefaultUpdatedAt = feedDescUpdatedAt.Default.(func() time.Time)
	// feed.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	feed.UpdateDefaultUpdatedAt = feedDescUpdatedAt.UpdateDefault.(func() time.Time)
	// feedDescCreatedAt is the schema descriptor for created_at field.
	feedDescCreatedAt := feedFields[8].Descriptor()
	// feed.DefaultCreatedAt holds the default value on creation for the created_at field.
	feed.DefaultCreatedAt = feedDescCreatedAt.Default.(func() time.Time)
	feedconfigFields := schema.FeedConfig{}.Fields()
	_ = feedconfigFields
	// feedconfigDescHideItems is the schema descriptor for hide_items field.
	feedconfigDescHideItems := feedconfigFields[9].Descriptor()
	// feedconfig.DefaultHideItems holds the default value on creation for the hide_items field.
	feedconfig.DefaultHideItems = feedconfigDescHideItems.Default.(bool)
	// feedconfigDescLatestPullAt is the schema descriptor for latest_pull_at field.
	feedconfigDescLatestPullAt := feedconfigFields[10].Descriptor()
	// feedconfig.DefaultLatestPullAt holds the default value on creation for the latest_pull_at field.
	feedconfig.DefaultLatestPullAt = feedconfigDescLatestPullAt.Default.(time.Time)
	// feedconfigDescNextPullBeginAt is the schema descriptor for next_pull_begin_at field.
	feedconfigDescNextPullBeginAt := feedconfigFields[11].Descriptor()
	// feedconfig.DefaultNextPullBeginAt holds the default value on creation for the next_pull_begin_at field.
	feedconfig.DefaultNextPullBeginAt = feedconfigDescNextPullBeginAt.Default.(time.Time)
	// feedconfigDescUpdatedAt is the schema descriptor for updated_at field.
	feedconfigDescUpdatedAt := feedconfigFields[12].Descriptor()
	// feedconfig.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	feedconfig.DefaultUpdatedAt = feedconfigDescUpdatedAt.Default.(func() time.Time)
	// feedconfig.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	feedconfig.UpdateDefaultUpdatedAt = feedconfigDescUpdatedAt.UpdateDefault.(func() time.Time)
	// feedconfigDescCreatedAt is the schema descriptor for created_at field.
	feedconfigDescCreatedAt := feedconfigFields[13].Descriptor()
	// feedconfig.DefaultCreatedAt holds the default value on creation for the created_at field.
	feedconfig.DefaultCreatedAt = feedconfigDescCreatedAt.Default.(func() time.Time)
	feeditemFields := schema.FeedItem{}.Fields()
	_ = feeditemFields
	// feeditemDescReadCount is the schema descriptor for read_count field.
	feeditemDescReadCount := feeditemFields[15].Descriptor()
	// feeditem.DefaultReadCount holds the default value on creation for the read_count field.
	feeditem.DefaultReadCount = feeditemDescReadCount.Default.(int64)
	// feeditemDescUpdatedAt is the schema descriptor for updated_at field.
	feeditemDescUpdatedAt := feeditemFields[18].Descriptor()
	// feeditem.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	feeditem.DefaultUpdatedAt = feeditemDescUpdatedAt.Default.(func() time.Time)
	// feeditem.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	feeditem.UpdateDefaultUpdatedAt = feeditemDescUpdatedAt.UpdateDefault.(func() time.Time)
	// feeditemDescCreatedAt is the schema descriptor for created_at field.
	feeditemDescCreatedAt := feeditemFields[19].Descriptor()
	// feeditem.DefaultCreatedAt holds the default value on creation for the created_at field.
	feeditem.DefaultCreatedAt = feeditemDescCreatedAt.Default.(func() time.Time)
	fileFields := schema.File{}.Fields()
	_ = fileFields
	// fileDescUpdatedAt is the schema descriptor for updated_at field.
	fileDescUpdatedAt := fileFields[5].Descriptor()
	// file.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	file.DefaultUpdatedAt = fileDescUpdatedAt.Default.(func() time.Time)
	// file.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	file.UpdateDefaultUpdatedAt = fileDescUpdatedAt.UpdateDefault.(func() time.Time)
	// fileDescCreatedAt is the schema descriptor for created_at field.
	fileDescCreatedAt := fileFields[6].Descriptor()
	// file.DefaultCreatedAt holds the default value on creation for the created_at field.
	file.DefaultCreatedAt = fileDescCreatedAt.Default.(func() time.Time)
	imageFields := schema.Image{}.Fields()
	_ = imageFields
	// imageDescUpdatedAt is the schema descriptor for updated_at field.
	imageDescUpdatedAt := imageFields[4].Descriptor()
	// image.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	image.DefaultUpdatedAt = imageDescUpdatedAt.Default.(func() time.Time)
	// image.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	image.UpdateDefaultUpdatedAt = imageDescUpdatedAt.UpdateDefault.(func() time.Time)
	// imageDescCreatedAt is the schema descriptor for created_at field.
	imageDescCreatedAt := imageFields[5].Descriptor()
	// image.DefaultCreatedAt holds the default value on creation for the created_at field.
	image.DefaultCreatedAt = imageDescCreatedAt.Default.(func() time.Time)
	notifyflowFields := schema.NotifyFlow{}.Fields()
	_ = notifyflowFields
	// notifyflowDescUpdatedAt is the schema descriptor for updated_at field.
	notifyflowDescUpdatedAt := notifyflowFields[4].Descriptor()
	// notifyflow.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notifyflow.DefaultUpdatedAt = notifyflowDescUpdatedAt.Default.(func() time.Time)
	// notifyflow.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notifyflow.UpdateDefaultUpdatedAt = notifyflowDescUpdatedAt.UpdateDefault.(func() time.Time)
	// notifyflowDescCreatedAt is the schema descriptor for created_at field.
	notifyflowDescCreatedAt := notifyflowFields[5].Descriptor()
	// notifyflow.DefaultCreatedAt holds the default value on creation for the created_at field.
	notifyflow.DefaultCreatedAt = notifyflowDescCreatedAt.Default.(func() time.Time)
	notifyflowsourceFields := schema.NotifyFlowSource{}.Fields()
	_ = notifyflowsourceFields
	// notifyflowsourceDescUpdatedAt is the schema descriptor for updated_at field.
	notifyflowsourceDescUpdatedAt := notifyflowsourceFields[4].Descriptor()
	// notifyflowsource.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notifyflowsource.DefaultUpdatedAt = notifyflowsourceDescUpdatedAt.Default.(func() time.Time)
	// notifyflowsource.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notifyflowsource.UpdateDefaultUpdatedAt = notifyflowsourceDescUpdatedAt.UpdateDefault.(func() time.Time)
	// notifyflowsourceDescCreatedAt is the schema descriptor for created_at field.
	notifyflowsourceDescCreatedAt := notifyflowsourceFields[5].Descriptor()
	// notifyflowsource.DefaultCreatedAt holds the default value on creation for the created_at field.
	notifyflowsource.DefaultCreatedAt = notifyflowsourceDescCreatedAt.Default.(func() time.Time)
	notifyflowtargetFields := schema.NotifyFlowTarget{}.Fields()
	_ = notifyflowtargetFields
	// notifyflowtargetDescUpdatedAt is the schema descriptor for updated_at field.
	notifyflowtargetDescUpdatedAt := notifyflowtargetFields[5].Descriptor()
	// notifyflowtarget.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notifyflowtarget.DefaultUpdatedAt = notifyflowtargetDescUpdatedAt.Default.(func() time.Time)
	// notifyflowtarget.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notifyflowtarget.UpdateDefaultUpdatedAt = notifyflowtargetDescUpdatedAt.UpdateDefault.(func() time.Time)
	// notifyflowtargetDescCreatedAt is the schema descriptor for created_at field.
	notifyflowtargetDescCreatedAt := notifyflowtargetFields[6].Descriptor()
	// notifyflowtarget.DefaultCreatedAt holds the default value on creation for the created_at field.
	notifyflowtarget.DefaultCreatedAt = notifyflowtargetDescCreatedAt.Default.(func() time.Time)
	notifytargetFields := schema.NotifyTarget{}.Fields()
	_ = notifytargetFields
	// notifytargetDescUpdatedAt is the schema descriptor for updated_at field.
	notifytargetDescUpdatedAt := notifytargetFields[6].Descriptor()
	// notifytarget.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notifytarget.DefaultUpdatedAt = notifytargetDescUpdatedAt.Default.(func() time.Time)
	// notifytarget.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notifytarget.UpdateDefaultUpdatedAt = notifytargetDescUpdatedAt.UpdateDefault.(func() time.Time)
	// notifytargetDescCreatedAt is the schema descriptor for created_at field.
	notifytargetDescCreatedAt := notifytargetFields[7].Descriptor()
	// notifytarget.DefaultCreatedAt holds the default value on creation for the created_at field.
	notifytarget.DefaultCreatedAt = notifytargetDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[5].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[6].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
}
