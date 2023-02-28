// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "platform", Type: field.TypeEnum, Enums: []string{"steam"}},
		{Name: "platform_account_id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "profile_url", Type: field.TypeString},
		{Name: "avatar_url", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_account", Type: field.TypeInt64, Nullable: true},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "accounts_users_account",
				Columns:    []*schema.Column{AccountsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "account_platform_platform_account_id",
				Unique:  true,
				Columns: []*schema.Column{AccountsColumns[1], AccountsColumns[2]},
			},
		},
	}
	// AppsColumns holds the columns for the "apps" table.
	AppsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "source", Type: field.TypeEnum, Enums: []string{"internal", "steam"}},
		{Name: "source_app_id", Type: field.TypeString},
		{Name: "source_url", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"game"}},
		{Name: "short_description", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Size: 2147483647},
		{Name: "image_url", Type: field.TypeString},
		{Name: "release_date", Type: field.TypeString},
		{Name: "developer", Type: field.TypeString},
		{Name: "publisher", Type: field.TypeString},
		{Name: "version", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// AppsTable holds the schema information for the "apps" table.
	AppsTable = &schema.Table{
		Name:       "apps",
		Columns:    AppsColumns,
		PrimaryKey: []*schema.Column{AppsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "app_source_source_app_id",
				Unique:  true,
				Columns: []*schema.Column{AppsColumns[1], AppsColumns[2]},
			},
		},
	}
	// AppPackagesColumns holds the columns for the "app_packages" table.
	AppPackagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "source", Type: field.TypeEnum, Enums: []string{"manual", "sentinel"}},
		{Name: "source_id", Type: field.TypeInt64},
		{Name: "source_package_id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Size: 2147483647},
		{Name: "binary_name", Type: field.TypeString},
		{Name: "binary_size", Type: field.TypeInt64},
		{Name: "binary_public_url", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "app_app_package", Type: field.TypeInt64, Nullable: true},
	}
	// AppPackagesTable holds the schema information for the "app_packages" table.
	AppPackagesTable = &schema.Table{
		Name:       "app_packages",
		Columns:    AppPackagesColumns,
		PrimaryKey: []*schema.Column{AppPackagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "app_packages_apps_app_package",
				Columns:    []*schema.Column{AppPackagesColumns[11]},
				RefColumns: []*schema.Column{AppsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "apppackage_source_source_id_source_package_id",
				Unique:  true,
				Columns: []*schema.Column{AppPackagesColumns[1], AppPackagesColumns[2], AppPackagesColumns[3]},
			},
		},
	}
	// FeedsColumns holds the columns for the "feeds" table.
	FeedsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "title", Type: field.TypeString},
		{Name: "link", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "language", Type: field.TypeString},
		{Name: "authors", Type: field.TypeJSON},
		{Name: "image", Type: field.TypeJSON},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "feed_config_feed", Type: field.TypeInt64, Unique: true},
	}
	// FeedsTable holds the schema information for the "feeds" table.
	FeedsTable = &schema.Table{
		Name:       "feeds",
		Columns:    FeedsColumns,
		PrimaryKey: []*schema.Column{FeedsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "feeds_feed_configs_feed",
				Columns:    []*schema.Column{FeedsColumns[9]},
				RefColumns: []*schema.Column{FeedConfigsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// FeedConfigsColumns holds the columns for the "feed_configs" table.
	FeedConfigsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "feed_url", Type: field.TypeString},
		{Name: "author_account", Type: field.TypeInt64},
		{Name: "source", Type: field.TypeEnum, Enums: []string{"common"}},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"active", "suspend"}},
		{Name: "pull_interval", Type: field.TypeInt64},
		{Name: "next_pull_begin_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_feed_config", Type: field.TypeInt64},
	}
	// FeedConfigsTable holds the schema information for the "feed_configs" table.
	FeedConfigsTable = &schema.Table{
		Name:       "feed_configs",
		Columns:    FeedConfigsColumns,
		PrimaryKey: []*schema.Column{FeedConfigsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "feed_configs_users_feed_config",
				Columns:    []*schema.Column{FeedConfigsColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// FeedItemsColumns holds the columns for the "feed_items" table.
	FeedItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "title", Type: field.TypeString},
		{Name: "authors", Type: field.TypeJSON},
		{Name: "description", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "guid", Type: field.TypeString},
		{Name: "link", Type: field.TypeString},
		{Name: "image", Type: field.TypeJSON},
		{Name: "published", Type: field.TypeString},
		{Name: "published_parsed", Type: field.TypeTime},
		{Name: "updated", Type: field.TypeString},
		{Name: "updated_parsed", Type: field.TypeTime},
		{Name: "enclosure", Type: field.TypeJSON},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "feed_item", Type: field.TypeInt64},
	}
	// FeedItemsTable holds the schema information for the "feed_items" table.
	FeedItemsTable = &schema.Table{
		Name:       "feed_items",
		Columns:    FeedItemsColumns,
		PrimaryKey: []*schema.Column{FeedItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "feed_items_feeds_item",
				Columns:    []*schema.Column{FeedItemsColumns[15]},
				RefColumns: []*schema.Column{FeedsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"active", "blocked"}},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"admin", "normal", "sentinel"}},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_create", Type: field.TypeInt64, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_users_create",
				Columns:    []*schema.Column{UsersColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserAppColumns holds the columns for the "user_app" table.
	UserAppColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "app_id", Type: field.TypeInt64},
	}
	// UserAppTable holds the schema information for the "user_app" table.
	UserAppTable = &schema.Table{
		Name:       "user_app",
		Columns:    UserAppColumns,
		PrimaryKey: []*schema.Column{UserAppColumns[0], UserAppColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_app_user_id",
				Columns:    []*schema.Column{UserAppColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_app_app_id",
				Columns:    []*schema.Column{UserAppColumns[1]},
				RefColumns: []*schema.Column{AppsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		AppsTable,
		AppPackagesTable,
		FeedsTable,
		FeedConfigsTable,
		FeedItemsTable,
		UsersTable,
		UserAppTable,
	}
)

func init() {
	AccountsTable.ForeignKeys[0].RefTable = UsersTable
	AppPackagesTable.ForeignKeys[0].RefTable = AppsTable
	FeedsTable.ForeignKeys[0].RefTable = FeedConfigsTable
	FeedConfigsTable.ForeignKeys[0].RefTable = UsersTable
	FeedItemsTable.ForeignKeys[0].RefTable = FeedsTable
	UsersTable.ForeignKeys[0].RefTable = UsersTable
	UserAppTable.ForeignKeys[0].RefTable = UsersTable
	UserAppTable.ForeignKeys[1].RefTable = AppsTable
}
