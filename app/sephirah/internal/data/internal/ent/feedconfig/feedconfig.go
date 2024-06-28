// Code generated by ent, DO NOT EDIT.

package feedconfig

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the feedconfig type in the database.
	Label = "feed_config"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserFeedConfig holds the string denoting the user_feed_config field in the database.
	FieldUserFeedConfig = "user_feed_config"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldFeedURL holds the string denoting the feed_url field in the database.
	FieldFeedURL = "feed_url"
	// FieldAuthorAccount holds the string denoting the author_account field in the database.
	FieldAuthorAccount = "author_account"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// FieldPullInterval holds the string denoting the pull_interval field in the database.
	FieldPullInterval = "pull_interval"
	// FieldHideItems holds the string denoting the hide_items field in the database.
	FieldHideItems = "hide_items"
	// FieldLatestPullAt holds the string denoting the latest_pull_at field in the database.
	FieldLatestPullAt = "latest_pull_at"
	// FieldLatestPullStatus holds the string denoting the latest_pull_status field in the database.
	FieldLatestPullStatus = "latest_pull_status"
	// FieldLatestPullMessage holds the string denoting the latest_pull_message field in the database.
	FieldLatestPullMessage = "latest_pull_message"
	// FieldNextPullBeginAt holds the string denoting the next_pull_begin_at field in the database.
	FieldNextPullBeginAt = "next_pull_begin_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeFeed holds the string denoting the feed edge name in mutations.
	EdgeFeed = "feed"
	// EdgeNotifySource holds the string denoting the notify_source edge name in mutations.
	EdgeNotifySource = "notify_source"
	// EdgeFeedActionSet holds the string denoting the feed_action_set edge name in mutations.
	EdgeFeedActionSet = "feed_action_set"
	// EdgeFeedConfigAction holds the string denoting the feed_config_action edge name in mutations.
	EdgeFeedConfigAction = "feed_config_action"
	// Table holds the table name of the feedconfig in the database.
	Table = "feed_configs"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "feed_configs"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_feed_config"
	// FeedTable is the table that holds the feed relation/edge.
	FeedTable = "feeds"
	// FeedInverseTable is the table name for the Feed entity.
	// It exists in this package in order to avoid circular dependency with the "feed" package.
	FeedInverseTable = "feeds"
	// FeedColumn is the table column denoting the feed relation/edge.
	FeedColumn = "feed_config_feed"
	// NotifySourceTable is the table that holds the notify_source relation/edge.
	NotifySourceTable = "notify_sources"
	// NotifySourceInverseTable is the table name for the NotifySource entity.
	// It exists in this package in order to avoid circular dependency with the "notifysource" package.
	NotifySourceInverseTable = "notify_sources"
	// NotifySourceColumn is the table column denoting the notify_source relation/edge.
	NotifySourceColumn = "feed_config_id"
	// FeedActionSetTable is the table that holds the feed_action_set relation/edge. The primary key declared below.
	FeedActionSetTable = "feed_config_actions"
	// FeedActionSetInverseTable is the table name for the FeedActionSet entity.
	// It exists in this package in order to avoid circular dependency with the "feedactionset" package.
	FeedActionSetInverseTable = "feed_action_sets"
	// FeedConfigActionTable is the table that holds the feed_config_action relation/edge.
	FeedConfigActionTable = "feed_config_actions"
	// FeedConfigActionInverseTable is the table name for the FeedConfigAction entity.
	// It exists in this package in order to avoid circular dependency with the "feedconfigaction" package.
	FeedConfigActionInverseTable = "feed_config_actions"
	// FeedConfigActionColumn is the table column denoting the feed_config_action relation/edge.
	FeedConfigActionColumn = "feed_config_id"
)

// Columns holds all SQL columns for feedconfig fields.
var Columns = []string{
	FieldID,
	FieldUserFeedConfig,
	FieldName,
	FieldFeedURL,
	FieldAuthorAccount,
	FieldSource,
	FieldStatus,
	FieldCategory,
	FieldPullInterval,
	FieldHideItems,
	FieldLatestPullAt,
	FieldLatestPullStatus,
	FieldLatestPullMessage,
	FieldNextPullBeginAt,
	FieldUpdatedAt,
	FieldCreatedAt,
}

var (
	// FeedActionSetPrimaryKey and FeedActionSetColumn2 are the table columns denoting the
	// primary key for the feed_action_set relation (M2M).
	FeedActionSetPrimaryKey = []string{"feed_config_id", "feed_action_set_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultHideItems holds the default value on creation for the "hide_items" field.
	DefaultHideItems bool
	// DefaultLatestPullAt holds the default value on creation for the "latest_pull_at" field.
	DefaultLatestPullAt time.Time
	// DefaultNextPullBeginAt holds the default value on creation for the "next_pull_begin_at" field.
	DefaultNextPullBeginAt time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusActive  Status = "active"
	StatusSuspend Status = "suspend"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusActive, StatusSuspend:
		return nil
	default:
		return fmt.Errorf("feedconfig: invalid enum value for status field: %q", s)
	}
}

// LatestPullStatus defines the type for the "latest_pull_status" enum field.
type LatestPullStatus string

// LatestPullStatus values.
const (
	LatestPullStatusProcessing LatestPullStatus = "processing"
	LatestPullStatusSuccess    LatestPullStatus = "success"
	LatestPullStatusFailed     LatestPullStatus = "failed"
)

func (lps LatestPullStatus) String() string {
	return string(lps)
}

// LatestPullStatusValidator is a validator for the "latest_pull_status" field enum values. It is called by the builders before save.
func LatestPullStatusValidator(lps LatestPullStatus) error {
	switch lps {
	case LatestPullStatusProcessing, LatestPullStatusSuccess, LatestPullStatusFailed:
		return nil
	default:
		return fmt.Errorf("feedconfig: invalid enum value for latest_pull_status field: %q", lps)
	}
}

// OrderOption defines the ordering options for the FeedConfig queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserFeedConfig orders the results by the user_feed_config field.
func ByUserFeedConfig(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserFeedConfig, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByFeedURL orders the results by the feed_url field.
func ByFeedURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFeedURL, opts...).ToFunc()
}

// ByAuthorAccount orders the results by the author_account field.
func ByAuthorAccount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthorAccount, opts...).ToFunc()
}

// BySource orders the results by the source field.
func BySource(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSource, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByCategory orders the results by the category field.
func ByCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategory, opts...).ToFunc()
}

// ByPullInterval orders the results by the pull_interval field.
func ByPullInterval(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPullInterval, opts...).ToFunc()
}

// ByHideItems orders the results by the hide_items field.
func ByHideItems(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHideItems, opts...).ToFunc()
}

// ByLatestPullAt orders the results by the latest_pull_at field.
func ByLatestPullAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLatestPullAt, opts...).ToFunc()
}

// ByLatestPullStatus orders the results by the latest_pull_status field.
func ByLatestPullStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLatestPullStatus, opts...).ToFunc()
}

// ByLatestPullMessage orders the results by the latest_pull_message field.
func ByLatestPullMessage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLatestPullMessage, opts...).ToFunc()
}

// ByNextPullBeginAt orders the results by the next_pull_begin_at field.
func ByNextPullBeginAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNextPullBeginAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}

// ByFeedField orders the results by feed field.
func ByFeedField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFeedStep(), sql.OrderByField(field, opts...))
	}
}

// ByNotifySourceCount orders the results by notify_source count.
func ByNotifySourceCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNotifySourceStep(), opts...)
	}
}

// ByNotifySource orders the results by notify_source terms.
func ByNotifySource(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNotifySourceStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFeedActionSetCount orders the results by feed_action_set count.
func ByFeedActionSetCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFeedActionSetStep(), opts...)
	}
}

// ByFeedActionSet orders the results by feed_action_set terms.
func ByFeedActionSet(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFeedActionSetStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFeedConfigActionCount orders the results by feed_config_action count.
func ByFeedConfigActionCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFeedConfigActionStep(), opts...)
	}
}

// ByFeedConfigAction orders the results by feed_config_action terms.
func ByFeedConfigAction(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFeedConfigActionStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
func newFeedStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FeedInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, FeedTable, FeedColumn),
	)
}
func newNotifySourceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NotifySourceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, NotifySourceTable, NotifySourceColumn),
	)
}
func newFeedActionSetStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FeedActionSetInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, FeedActionSetTable, FeedActionSetPrimaryKey...),
	)
}
func newFeedConfigActionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FeedConfigActionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, FeedConfigActionTable, FeedConfigActionColumn),
	)
}
