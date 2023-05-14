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
	// FieldPullInterval holds the string denoting the pull_interval field in the database.
	FieldPullInterval = "pull_interval"
	// FieldLatestPullAt holds the string denoting the latest_pull_at field in the database.
	FieldLatestPullAt = "latest_pull_at"
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
	// EdgeNotifyFlow holds the string denoting the notify_flow edge name in mutations.
	EdgeNotifyFlow = "notify_flow"
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
	// NotifyFlowTable is the table that holds the notify_flow relation/edge. The primary key declared below.
	NotifyFlowTable = "feed_config_notify_flow"
	// NotifyFlowInverseTable is the table name for the NotifyFlow entity.
	// It exists in this package in order to avoid circular dependency with the "notifyflow" package.
	NotifyFlowInverseTable = "notify_flows"
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
	FieldPullInterval,
	FieldLatestPullAt,
	FieldNextPullBeginAt,
	FieldUpdatedAt,
	FieldCreatedAt,
}

var (
	// NotifyFlowPrimaryKey and NotifyFlowColumn2 are the table columns denoting the
	// primary key for the notify_flow relation (M2M).
	NotifyFlowPrimaryKey = []string{"feed_config_id", "notify_flow_id"}
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

// Source defines the type for the "source" enum field.
type Source string

// Source values.
const (
	SourceCommon Source = "common"
)

func (s Source) String() string {
	return string(s)
}

// SourceValidator is a validator for the "source" field enum values. It is called by the builders before save.
func SourceValidator(s Source) error {
	switch s {
	case SourceCommon:
		return nil
	default:
		return fmt.Errorf("feedconfig: invalid enum value for source field: %q", s)
	}
}

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

// ByPullInterval orders the results by the pull_interval field.
func ByPullInterval(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPullInterval, opts...).ToFunc()
}

// ByLatestPullAt orders the results by the latest_pull_at field.
func ByLatestPullAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLatestPullAt, opts...).ToFunc()
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

// ByNotifyFlowCount orders the results by notify_flow count.
func ByNotifyFlowCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNotifyFlowStep(), opts...)
	}
}

// ByNotifyFlow orders the results by notify_flow terms.
func ByNotifyFlow(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNotifyFlowStep(), append([]sql.OrderTerm{term}, terms...)...)
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
func newNotifyFlowStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NotifyFlowInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, NotifyFlowTable, NotifyFlowPrimaryKey...),
	)
}
