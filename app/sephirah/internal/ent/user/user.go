// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeAccount holds the string denoting the account edge name in mutations.
	EdgeAccount = "account"
	// EdgeApp holds the string denoting the app edge name in mutations.
	EdgeApp = "app"
	// EdgeFeedConfig holds the string denoting the feed_config edge name in mutations.
	EdgeFeedConfig = "feed_config"
	// EdgeCreator holds the string denoting the creator edge name in mutations.
	EdgeCreator = "creator"
	// EdgeCreate holds the string denoting the create edge name in mutations.
	EdgeCreate = "create"
	// Table holds the table name of the user in the database.
	Table = "users"
	// AccountTable is the table that holds the account relation/edge.
	AccountTable = "accounts"
	// AccountInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	AccountInverseTable = "accounts"
	// AccountColumn is the table column denoting the account relation/edge.
	AccountColumn = "user_account"
	// AppTable is the table that holds the app relation/edge. The primary key declared below.
	AppTable = "user_app"
	// AppInverseTable is the table name for the App entity.
	// It exists in this package in order to avoid circular dependency with the "app" package.
	AppInverseTable = "apps"
	// FeedConfigTable is the table that holds the feed_config relation/edge.
	FeedConfigTable = "feed_configs"
	// FeedConfigInverseTable is the table name for the FeedConfig entity.
	// It exists in this package in order to avoid circular dependency with the "feedconfig" package.
	FeedConfigInverseTable = "feed_configs"
	// FeedConfigColumn is the table column denoting the feed_config relation/edge.
	FeedConfigColumn = "user_feed_config"
	// CreatorTable is the table that holds the creator relation/edge.
	CreatorTable = "users"
	// CreatorColumn is the table column denoting the creator relation/edge.
	CreatorColumn = "user_create"
	// CreateTable is the table that holds the create relation/edge.
	CreateTable = "users"
	// CreateColumn is the table column denoting the create relation/edge.
	CreateColumn = "user_create"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldPassword,
	FieldStatus,
	FieldType,
	FieldUpdatedAt,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_create",
}

var (
	// AppPrimaryKey and AppColumn2 are the table columns denoting the
	// primary key for the app relation (M2M).
	AppPrimaryKey = []string{"user_id", "app_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
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
	StatusBlocked Status = "blocked"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusActive, StatusBlocked:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for status field: %q", s)
	}
}

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeAdmin Type = "admin"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeAdmin:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for type field: %q", _type)
	}
}
