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

// FeedURL applies equality check predicate on the "feed_url" field. It's identical to FeedURLEQ.
func FeedURL(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldFeedURL, v))
}

// AuthorAccount applies equality check predicate on the "author_account" field. It's identical to AuthorAccountEQ.
func AuthorAccount(v model.InternalID) predicate.FeedConfig {
	vc := int64(v)
	return predicate.FeedConfig(sql.FieldEQ(FieldAuthorAccount, vc))
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

// FeedURLEQ applies the EQ predicate on the "feed_url" field.
func FeedURLEQ(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldFeedURL, v))
}

// FeedURLNEQ applies the NEQ predicate on the "feed_url" field.
func FeedURLNEQ(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNEQ(FieldFeedURL, v))
}

// FeedURLIn applies the In predicate on the "feed_url" field.
func FeedURLIn(vs ...string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldIn(FieldFeedURL, vs...))
}

// FeedURLNotIn applies the NotIn predicate on the "feed_url" field.
func FeedURLNotIn(vs ...string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNotIn(FieldFeedURL, vs...))
}

// FeedURLGT applies the GT predicate on the "feed_url" field.
func FeedURLGT(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGT(FieldFeedURL, v))
}

// FeedURLGTE applies the GTE predicate on the "feed_url" field.
func FeedURLGTE(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGTE(FieldFeedURL, v))
}

// FeedURLLT applies the LT predicate on the "feed_url" field.
func FeedURLLT(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLT(FieldFeedURL, v))
}

// FeedURLLTE applies the LTE predicate on the "feed_url" field.
func FeedURLLTE(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLTE(FieldFeedURL, v))
}

// FeedURLContains applies the Contains predicate on the "feed_url" field.
func FeedURLContains(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldContains(FieldFeedURL, v))
}

// FeedURLHasPrefix applies the HasPrefix predicate on the "feed_url" field.
func FeedURLHasPrefix(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldHasPrefix(FieldFeedURL, v))
}

// FeedURLHasSuffix applies the HasSuffix predicate on the "feed_url" field.
func FeedURLHasSuffix(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldHasSuffix(FieldFeedURL, v))
}

// FeedURLEqualFold applies the EqualFold predicate on the "feed_url" field.
func FeedURLEqualFold(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEqualFold(FieldFeedURL, v))
}

// FeedURLContainsFold applies the ContainsFold predicate on the "feed_url" field.
func FeedURLContainsFold(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldContainsFold(FieldFeedURL, v))
}

// AuthorAccountEQ applies the EQ predicate on the "author_account" field.
func AuthorAccountEQ(v model.InternalID) predicate.FeedConfig {
	vc := int64(v)
	return predicate.FeedConfig(sql.FieldEQ(FieldAuthorAccount, vc))
}

// AuthorAccountNEQ applies the NEQ predicate on the "author_account" field.
func AuthorAccountNEQ(v model.InternalID) predicate.FeedConfig {
	vc := int64(v)
	return predicate.FeedConfig(sql.FieldNEQ(FieldAuthorAccount, vc))
}

// AuthorAccountIn applies the In predicate on the "author_account" field.
func AuthorAccountIn(vs ...model.InternalID) predicate.FeedConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.FeedConfig(sql.FieldIn(FieldAuthorAccount, v...))
}

// AuthorAccountNotIn applies the NotIn predicate on the "author_account" field.
func AuthorAccountNotIn(vs ...model.InternalID) predicate.FeedConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.FeedConfig(sql.FieldNotIn(FieldAuthorAccount, v...))
}

// AuthorAccountGT applies the GT predicate on the "author_account" field.
func AuthorAccountGT(v model.InternalID) predicate.FeedConfig {
	vc := int64(v)
	return predicate.FeedConfig(sql.FieldGT(FieldAuthorAccount, vc))
}

// AuthorAccountGTE applies the GTE predicate on the "author_account" field.
func AuthorAccountGTE(v model.InternalID) predicate.FeedConfig {
	vc := int64(v)
	return predicate.FeedConfig(sql.FieldGTE(FieldAuthorAccount, vc))
}

// AuthorAccountLT applies the LT predicate on the "author_account" field.
func AuthorAccountLT(v model.InternalID) predicate.FeedConfig {
	vc := int64(v)
	return predicate.FeedConfig(sql.FieldLT(FieldAuthorAccount, vc))
}

// AuthorAccountLTE applies the LTE predicate on the "author_account" field.
func AuthorAccountLTE(v model.InternalID) predicate.FeedConfig {
	vc := int64(v)
	return predicate.FeedConfig(sql.FieldLTE(FieldAuthorAccount, vc))
}

// SourceEQ applies the EQ predicate on the "source" field.
func SourceEQ(v Source) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldSource, v))
}

// SourceNEQ applies the NEQ predicate on the "source" field.
func SourceNEQ(v Source) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNEQ(FieldSource, v))
}

// SourceIn applies the In predicate on the "source" field.
func SourceIn(vs ...Source) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldIn(FieldSource, vs...))
}

// SourceNotIn applies the NotIn predicate on the "source" field.
func SourceNotIn(vs ...Source) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNotIn(FieldSource, vs...))
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

// HasNotifyFlow applies the HasEdge predicate on the "notify_flow" edge.
func HasNotifyFlow() predicate.FeedConfig {
	return predicate.FeedConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, NotifyFlowTable, NotifyFlowPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNotifyFlowWith applies the HasEdge predicate on the "notify_flow" edge with a given conditions (other predicates).
func HasNotifyFlowWith(preds ...predicate.NotifyFlow) predicate.FeedConfig {
	return predicate.FeedConfig(func(s *sql.Selector) {
		step := newNotifyFlowStep()
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
