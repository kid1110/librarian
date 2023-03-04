// Code generated by ent, DO NOT EDIT.

package feedconfig

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/predicate"
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

// FeedURL applies equality check predicate on the "feed_url" field. It's identical to FeedURLEQ.
func FeedURL(v string) predicate.FeedConfig {
	return predicate.FeedConfig(sql.FieldEQ(FieldFeedURL, v))
}

// AuthorAccount applies equality check predicate on the "author_account" field. It's identical to AuthorAccountEQ.
func AuthorAccount(v model.InternalID) predicate.FeedConfig {
	vc := int64(v)
	return predicate.FeedConfig(sql.FieldEQ(FieldAuthorAccount, vc))
}

// PullInterval applies equality check predicate on the "pull_interval" field. It's identical to PullIntervalEQ.
func PullInterval(v time.Duration) predicate.FeedConfig {
	vc := int64(v)
	return predicate.FeedConfig(sql.FieldEQ(FieldPullInterval, vc))
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

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.FeedConfig {
	return predicate.FeedConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.FeedConfig {
	return predicate.FeedConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
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
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FeedInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, FeedTable, FeedColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
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
