// Code generated by ent, DO NOT EDIT.

package feedconfig

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLTE(FieldID, id))
}

// InternalID applies equality check predicate on the "internal_id" field. It's identical to InternalIDEQ.
func InternalID(v int64) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldInternalID, v))
}

// FeedURL applies equality check predicate on the "feed_url" field. It's identical to FeedURLEQ.
func FeedURL(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldFeedURL, v))
}

// AuthorAccount applies equality check predicate on the "author_account" field. It's identical to AuthorAccountEQ.
func AuthorAccount(v int64) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldAuthorAccount, v))
}

// PullInterval applies equality check predicate on the "pull_interval" field. It's identical to PullIntervalEQ.
func PullInterval(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldPullInterval, v))
}

// LastPullAt applies equality check predicate on the "last_pull_at" field. It's identical to LastPullAtEQ.
func LastPullAt(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldLastPullAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldCreatedAt, v))
}

// InternalIDEQ applies the EQ predicate on the "internal_id" field.
func InternalIDEQ(v int64) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldInternalID, v))
}

// InternalIDNEQ applies the NEQ predicate on the "internal_id" field.
func InternalIDNEQ(v int64) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNEQ(FieldInternalID, v))
}

// InternalIDIn applies the In predicate on the "internal_id" field.
func InternalIDIn(vs ...int64) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldIn(FieldInternalID, vs...))
}

// InternalIDNotIn applies the NotIn predicate on the "internal_id" field.
func InternalIDNotIn(vs ...int64) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNotIn(FieldInternalID, vs...))
}

// InternalIDGT applies the GT predicate on the "internal_id" field.
func InternalIDGT(v int64) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGT(FieldInternalID, v))
}

// InternalIDGTE applies the GTE predicate on the "internal_id" field.
func InternalIDGTE(v int64) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGTE(FieldInternalID, v))
}

// InternalIDLT applies the LT predicate on the "internal_id" field.
func InternalIDLT(v int64) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLT(FieldInternalID, v))
}

// InternalIDLTE applies the LTE predicate on the "internal_id" field.
func InternalIDLTE(v int64) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLTE(FieldInternalID, v))
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
func AuthorAccountEQ(v int64) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldAuthorAccount, v))
}

// AuthorAccountNEQ applies the NEQ predicate on the "author_account" field.
func AuthorAccountNEQ(v int64) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNEQ(FieldAuthorAccount, v))
}

// AuthorAccountIn applies the In predicate on the "author_account" field.
func AuthorAccountIn(vs ...int64) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldIn(FieldAuthorAccount, vs...))
}

// AuthorAccountNotIn applies the NotIn predicate on the "author_account" field.
func AuthorAccountNotIn(vs ...int64) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNotIn(FieldAuthorAccount, vs...))
}

// AuthorAccountGT applies the GT predicate on the "author_account" field.
func AuthorAccountGT(v int64) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGT(FieldAuthorAccount, v))
}

// AuthorAccountGTE applies the GTE predicate on the "author_account" field.
func AuthorAccountGTE(v int64) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGTE(FieldAuthorAccount, v))
}

// AuthorAccountLT applies the LT predicate on the "author_account" field.
func AuthorAccountLT(v int64) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLT(FieldAuthorAccount, v))
}

// AuthorAccountLTE applies the LTE predicate on the "author_account" field.
func AuthorAccountLTE(v int64) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLTE(FieldAuthorAccount, v))
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

// PullIntervalEQ applies the EQ predicate on the "pull_interval" field.
func PullIntervalEQ(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldPullInterval, v))
}

// PullIntervalNEQ applies the NEQ predicate on the "pull_interval" field.
func PullIntervalNEQ(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNEQ(FieldPullInterval, v))
}

// PullIntervalIn applies the In predicate on the "pull_interval" field.
func PullIntervalIn(vs ...time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldIn(FieldPullInterval, vs...))
}

// PullIntervalNotIn applies the NotIn predicate on the "pull_interval" field.
func PullIntervalNotIn(vs ...time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNotIn(FieldPullInterval, vs...))
}

// PullIntervalGT applies the GT predicate on the "pull_interval" field.
func PullIntervalGT(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGT(FieldPullInterval, v))
}

// PullIntervalGTE applies the GTE predicate on the "pull_interval" field.
func PullIntervalGTE(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGTE(FieldPullInterval, v))
}

// PullIntervalLT applies the LT predicate on the "pull_interval" field.
func PullIntervalLT(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLT(FieldPullInterval, v))
}

// PullIntervalLTE applies the LTE predicate on the "pull_interval" field.
func PullIntervalLTE(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLTE(FieldPullInterval, v))
}

// LastPullAtEQ applies the EQ predicate on the "last_pull_at" field.
func LastPullAtEQ(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldLastPullAt, v))
}

// LastPullAtNEQ applies the NEQ predicate on the "last_pull_at" field.
func LastPullAtNEQ(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNEQ(FieldLastPullAt, v))
}

// LastPullAtIn applies the In predicate on the "last_pull_at" field.
func LastPullAtIn(vs ...time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldIn(FieldLastPullAt, vs...))
}

// LastPullAtNotIn applies the NotIn predicate on the "last_pull_at" field.
func LastPullAtNotIn(vs ...time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldNotIn(FieldLastPullAt, vs...))
}

// LastPullAtGT applies the GT predicate on the "last_pull_at" field.
func LastPullAtGT(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGT(FieldLastPullAt, v))
}

// LastPullAtGTE applies the GTE predicate on the "last_pull_at" field.
func LastPullAtGTE(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldGTE(FieldLastPullAt, v))
}

// LastPullAtLT applies the LT predicate on the "last_pull_at" field.
func LastPullAtLT(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLT(FieldLastPullAt, v))
}

// LastPullAtLTE applies the LTE predicate on the "last_pull_at" field.
func LastPullAtLTE(v time.Time) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldLTE(FieldLastPullAt, v))
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

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FeedConfig) predicate.FeedConfig {
	return predicate.FeedConfig(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FeedConfig) predicate.FeedConfig {
	return predicate.FeedConfig(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.FeedConfig) predicate.FeedConfig {
	return predicate.FeedConfig(func(s *sql.Selector) {
		p(s.Not())
	})
}
