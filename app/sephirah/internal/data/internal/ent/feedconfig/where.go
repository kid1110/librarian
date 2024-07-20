// Code generated by ent, DO NOT EDIT.

package feedconfig

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/model"
)

// ID filters vertices based on their ID field.
func ID(id model.InternalID) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id model.InternalID) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id model.InternalID) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...model.InternalID) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...model.InternalID) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id model.InternalID) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id model.InternalID) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id model.InternalID) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id model.InternalID) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLTE(FieldID, id))
}

// UserFeedConfig applies equality check predicate on the "user_feed_config" field. It's identical to UserFeedConfigEQ.
func UserFeedConfig(v model.InternalID) predicate.FeedConfig {
	vc := int64(v)
	return predicate.FeedConfig(sql.FieldEQ(FieldUserFeedConfig, vc))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldDescription, v))
}

// Category applies equality check predicate on the "category" field. It's identical to CategoryEQ.
func Category(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldCategory, v))
}

// PullInterval applies equality check predicate on the "pull_interval" field. It's identical to PullIntervalEQ.
func PullInterval(v time.Duration) predicate.FeedConfig {
	vc := int64(v)
	return predicate.FeedConfig(sql.FieldEQ(FieldPullInterval, vc))
}

// HideItems applies equality check predicate on the "hide_items" field. It's identical to HideItemsEQ.
func HideItems(v bool) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldHideItems, v))
}

// LatestPullAt applies equality check predicate on the "latest_pull_at" field. It's identical to LatestPullAtEQ.
func LatestPullAt(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldLatestPullAt, v))
}

// LatestPullMessage applies equality check predicate on the "latest_pull_message" field. It's identical to LatestPullMessageEQ.
func LatestPullMessage(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldLatestPullMessage, v))
}

// NextPullBeginAt applies equality check predicate on the "next_pull_begin_at" field. It's identical to NextPullBeginAtEQ.
func NextPullBeginAt(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldNextPullBeginAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldCreatedAt, v))
}

// UserFeedConfigEQ applies the EQ predicate on the "user_feed_config" field.
func UserFeedConfigEQ(v model.InternalID) predicate.FeedConfig {
	vc := int64(v)
	return predicate.FeedConfig(sql.FieldEQ(FieldUserFeedConfig, vc))
}

// UserFeedConfigNEQ applies the NEQ predicate on the "user_feed_config" field.
func UserFeedConfigNEQ(v model.InternalID) predicate.FeedConfig {
	vc := int64(v)
	return predicate.FeedConfig(sql.FieldNEQ(FieldUserFeedConfig, vc))
}

// UserFeedConfigIn applies the In predicate on the "user_feed_config" field.
func UserFeedConfigIn(vs ...model.InternalID) predicate.FeedConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.FeedConfig(sql.FieldIn(FieldUserFeedConfig, v...))
}

// UserFeedConfigNotIn applies the NotIn predicate on the "user_feed_config" field.
func UserFeedConfigNotIn(vs ...model.InternalID) predicate.FeedConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.FeedConfig(sql.FieldNotIn(FieldUserFeedConfig, v...))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldContainsFold(FieldDescription, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNotIn(FieldStatus, vs...))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNotIn(FieldCategory, vs...))
}

// CategoryGT applies the GT predicate on the "category" field.
func CategoryGT(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGT(FieldCategory, v))
}

// CategoryGTE applies the GTE predicate on the "category" field.
func CategoryGTE(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGTE(FieldCategory, v))
}

// CategoryLT applies the LT predicate on the "category" field.
func CategoryLT(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLT(FieldCategory, v))
}

// CategoryLTE applies the LTE predicate on the "category" field.
func CategoryLTE(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLTE(FieldCategory, v))
}

// CategoryContains applies the Contains predicate on the "category" field.
func CategoryContains(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldContains(FieldCategory, v))
}

// CategoryHasPrefix applies the HasPrefix predicate on the "category" field.
func CategoryHasPrefix(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldHasPrefix(FieldCategory, v))
}

// CategoryHasSuffix applies the HasSuffix predicate on the "category" field.
func CategoryHasSuffix(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldHasSuffix(FieldCategory, v))
}

// CategoryEqualFold applies the EqualFold predicate on the "category" field.
func CategoryEqualFold(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEqualFold(FieldCategory, v))
}

// CategoryContainsFold applies the ContainsFold predicate on the "category" field.
func CategoryContainsFold(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldContainsFold(FieldCategory, v))
}

// PullIntervalEQ applies the EQ predicate on the "pull_interval" field.
func PullIntervalEQ(v time.Duration) predicate.FeedConfig {
	vc := int64(v)
	return predicate.FeedConfig(sql.FieldEQ(FieldPullInterval, vc))
}

// PullIntervalNEQ applies the NEQ predicate on the "pull_interval" field.
func PullIntervalNEQ(v time.Duration) predicate.FeedConfig {
	vc := int64(v)
	return predicate.FeedConfig(sql.FieldNEQ(FieldPullInterval, vc))
}

// PullIntervalIn applies the In predicate on the "pull_interval" field.
func PullIntervalIn(vs ...time.Duration) predicate.FeedConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.FeedConfig(sql.FieldIn(FieldPullInterval, v...))
}

// PullIntervalNotIn applies the NotIn predicate on the "pull_interval" field.
func PullIntervalNotIn(vs ...time.Duration) predicate.FeedConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.FeedConfig(sql.FieldNotIn(FieldPullInterval, v...))
}

// PullIntervalGT applies the GT predicate on the "pull_interval" field.
func PullIntervalGT(v time.Duration) predicate.FeedConfig {
	vc := int64(v)
	return predicate.FeedConfig(sql.FieldGT(FieldPullInterval, vc))
}

// PullIntervalGTE applies the GTE predicate on the "pull_interval" field.
func PullIntervalGTE(v time.Duration) predicate.FeedConfig {
	vc := int64(v)
	return predicate.FeedConfig(sql.FieldGTE(FieldPullInterval, vc))
}

// PullIntervalLT applies the LT predicate on the "pull_interval" field.
func PullIntervalLT(v time.Duration) predicate.FeedConfig {
	vc := int64(v)
	return predicate.FeedConfig(sql.FieldLT(FieldPullInterval, vc))
}

// PullIntervalLTE applies the LTE predicate on the "pull_interval" field.
func PullIntervalLTE(v time.Duration) predicate.FeedConfig {
	vc := int64(v)
	return predicate.FeedConfig(sql.FieldLTE(FieldPullInterval, vc))
}

// HideItemsEQ applies the EQ predicate on the "hide_items" field.
func HideItemsEQ(v bool) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldHideItems, v))
}

// HideItemsNEQ applies the NEQ predicate on the "hide_items" field.
func HideItemsNEQ(v bool) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNEQ(FieldHideItems, v))
}

// LatestPullAtEQ applies the EQ predicate on the "latest_pull_at" field.
func LatestPullAtEQ(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldLatestPullAt, v))
}

// LatestPullAtNEQ applies the NEQ predicate on the "latest_pull_at" field.
func LatestPullAtNEQ(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNEQ(FieldLatestPullAt, v))
}

// LatestPullAtIn applies the In predicate on the "latest_pull_at" field.
func LatestPullAtIn(vs ...time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldIn(FieldLatestPullAt, vs...))
}

// LatestPullAtNotIn applies the NotIn predicate on the "latest_pull_at" field.
func LatestPullAtNotIn(vs ...time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNotIn(FieldLatestPullAt, vs...))
}

// LatestPullAtGT applies the GT predicate on the "latest_pull_at" field.
func LatestPullAtGT(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGT(FieldLatestPullAt, v))
}

// LatestPullAtGTE applies the GTE predicate on the "latest_pull_at" field.
func LatestPullAtGTE(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGTE(FieldLatestPullAt, v))
}

// LatestPullAtLT applies the LT predicate on the "latest_pull_at" field.
func LatestPullAtLT(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLT(FieldLatestPullAt, v))
}

// LatestPullAtLTE applies the LTE predicate on the "latest_pull_at" field.
func LatestPullAtLTE(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLTE(FieldLatestPullAt, v))
}

// LatestPullStatusEQ applies the EQ predicate on the "latest_pull_status" field.
func LatestPullStatusEQ(v LatestPullStatus) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldLatestPullStatus, v))
}

// LatestPullStatusNEQ applies the NEQ predicate on the "latest_pull_status" field.
func LatestPullStatusNEQ(v LatestPullStatus) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNEQ(FieldLatestPullStatus, v))
}

// LatestPullStatusIn applies the In predicate on the "latest_pull_status" field.
func LatestPullStatusIn(vs ...LatestPullStatus) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldIn(FieldLatestPullStatus, vs...))
}

// LatestPullStatusNotIn applies the NotIn predicate on the "latest_pull_status" field.
func LatestPullStatusNotIn(vs ...LatestPullStatus) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNotIn(FieldLatestPullStatus, vs...))
}

// LatestPullMessageEQ applies the EQ predicate on the "latest_pull_message" field.
func LatestPullMessageEQ(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldLatestPullMessage, v))
}

// LatestPullMessageNEQ applies the NEQ predicate on the "latest_pull_message" field.
func LatestPullMessageNEQ(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNEQ(FieldLatestPullMessage, v))
}

// LatestPullMessageIn applies the In predicate on the "latest_pull_message" field.
func LatestPullMessageIn(vs ...string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldIn(FieldLatestPullMessage, vs...))
}

// LatestPullMessageNotIn applies the NotIn predicate on the "latest_pull_message" field.
func LatestPullMessageNotIn(vs ...string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNotIn(FieldLatestPullMessage, vs...))
}

// LatestPullMessageGT applies the GT predicate on the "latest_pull_message" field.
func LatestPullMessageGT(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGT(FieldLatestPullMessage, v))
}

// LatestPullMessageGTE applies the GTE predicate on the "latest_pull_message" field.
func LatestPullMessageGTE(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGTE(FieldLatestPullMessage, v))
}

// LatestPullMessageLT applies the LT predicate on the "latest_pull_message" field.
func LatestPullMessageLT(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLT(FieldLatestPullMessage, v))
}

// LatestPullMessageLTE applies the LTE predicate on the "latest_pull_message" field.
func LatestPullMessageLTE(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLTE(FieldLatestPullMessage, v))
}

// LatestPullMessageContains applies the Contains predicate on the "latest_pull_message" field.
func LatestPullMessageContains(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldContains(FieldLatestPullMessage, v))
}

// LatestPullMessageHasPrefix applies the HasPrefix predicate on the "latest_pull_message" field.
func LatestPullMessageHasPrefix(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldHasPrefix(FieldLatestPullMessage, v))
}

// LatestPullMessageHasSuffix applies the HasSuffix predicate on the "latest_pull_message" field.
func LatestPullMessageHasSuffix(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldHasSuffix(FieldLatestPullMessage, v))
}

// LatestPullMessageEqualFold applies the EqualFold predicate on the "latest_pull_message" field.
func LatestPullMessageEqualFold(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEqualFold(FieldLatestPullMessage, v))
}

// LatestPullMessageContainsFold applies the ContainsFold predicate on the "latest_pull_message" field.
func LatestPullMessageContainsFold(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldContainsFold(FieldLatestPullMessage, v))
}

// NextPullBeginAtEQ applies the EQ predicate on the "next_pull_begin_at" field.
func NextPullBeginAtEQ(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldNextPullBeginAt, v))
}

// NextPullBeginAtNEQ applies the NEQ predicate on the "next_pull_begin_at" field.
func NextPullBeginAtNEQ(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNEQ(FieldNextPullBeginAt, v))
}

// NextPullBeginAtIn applies the In predicate on the "next_pull_begin_at" field.
func NextPullBeginAtIn(vs ...time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldIn(FieldNextPullBeginAt, vs...))
}

// NextPullBeginAtNotIn applies the NotIn predicate on the "next_pull_begin_at" field.
func NextPullBeginAtNotIn(vs ...time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNotIn(FieldNextPullBeginAt, vs...))
}

// NextPullBeginAtGT applies the GT predicate on the "next_pull_begin_at" field.
func NextPullBeginAtGT(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGT(FieldNextPullBeginAt, v))
}

// NextPullBeginAtGTE applies the GTE predicate on the "next_pull_begin_at" field.
func NextPullBeginAtGTE(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGTE(FieldNextPullBeginAt, v))
}

// NextPullBeginAtLT applies the LT predicate on the "next_pull_begin_at" field.
func NextPullBeginAtLT(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLT(FieldNextPullBeginAt, v))
}

// NextPullBeginAtLTE applies the LTE predicate on the "next_pull_begin_at" field.
func NextPullBeginAtLTE(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLTE(FieldNextPullBeginAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLTE(FieldCreatedAt, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.FeedConfig {
	return predicate.FeedConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.FeedConfig {
	return predicate.FeedConfig(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFeed applies the HasEdge predicate on the "feed" edge.
func HasFeed() predicate.FeedConfig {
	return predicate.FeedConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, FeedTable, FeedColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFeedWith applies the HasEdge predicate on the "feed" edge with a given conditions (other predicates).
func HasFeedWith(preds ...predicate.Feed) predicate.FeedConfig {
	return predicate.FeedConfig(func(s *sql.Selector) {
		step := newFeedStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNotifySource applies the HasEdge predicate on the "notify_source" edge.
func HasNotifySource() predicate.FeedConfig {
	return predicate.FeedConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NotifySourceTable, NotifySourceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNotifySourceWith applies the HasEdge predicate on the "notify_source" edge with a given conditions (other predicates).
func HasNotifySourceWith(preds ...predicate.NotifySource) predicate.FeedConfig {
	return predicate.FeedConfig(func(s *sql.Selector) {
		step := newNotifySourceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFeedActionSet applies the HasEdge predicate on the "feed_action_set" edge.
func HasFeedActionSet() predicate.FeedConfig {
	return predicate.FeedConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, FeedActionSetTable, FeedActionSetPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFeedActionSetWith applies the HasEdge predicate on the "feed_action_set" edge with a given conditions (other predicates).
func HasFeedActionSetWith(preds ...predicate.FeedActionSet) predicate.FeedConfig {
	return predicate.FeedConfig(func(s *sql.Selector) {
		step := newFeedActionSetStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFeedConfigAction applies the HasEdge predicate on the "feed_config_action" edge.
func HasFeedConfigAction() predicate.FeedConfig {
	return predicate.FeedConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, FeedConfigActionTable, FeedConfigActionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFeedConfigActionWith applies the HasEdge predicate on the "feed_config_action" edge with a given conditions (other predicates).
func HasFeedConfigActionWith(preds ...predicate.FeedConfigAction) predicate.FeedConfig {
	return predicate.FeedConfig(func(s *sql.Selector) {
		step := newFeedConfigActionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FeedConfig) predicate.FeedConfig {
	return predicate.FeedConfig(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FeedConfig) predicate.FeedConfig {
	return predicate.FeedConfig(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.FeedConfig) predicate.FeedConfig {
	return predicate.FeedConfig(sql.NotPredicates(p))
}
