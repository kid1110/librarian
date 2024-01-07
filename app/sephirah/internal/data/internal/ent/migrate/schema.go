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
		{Name: "platform", Type: field.TypeString},
		{Name: "platform_account_id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "profile_url", Type: field.TypeString},
		{Name: "avatar_url", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_bind_account", Type: field.TypeInt64, Nullable: true},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "accounts_users_bind_account",
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
		{Name: "internal", Type: field.TypeBool},
		{Name: "source", Type: field.TypeString},
		{Name: "source_app_id", Type: field.TypeString},
		{Name: "source_url", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"unknown", "game"}},
		{Name: "short_description", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "icon_image_url", Type: field.TypeString, Nullable: true},
		{Name: "hero_image_url", Type: field.TypeString, Nullable: true},
		{Name: "release_date", Type: field.TypeString, Nullable: true},
		{Name: "developer", Type: field.TypeString, Nullable: true},
		{Name: "publisher", Type: field.TypeString, Nullable: true},
		{Name: "version", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "app_bind_external", Type: field.TypeInt64, Nullable: true},
	}
	// AppsTable holds the schema information for the "apps" table.
	AppsTable = &schema.Table{
		Name:       "apps",
		Columns:    AppsColumns,
		PrimaryKey: []*schema.Column{AppsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "apps_apps_bind_external",
				Columns:    []*schema.Column{AppsColumns[17]},
				RefColumns: []*schema.Column{AppsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "app_source_source_app_id",
				Unique:  true,
				Columns: []*schema.Column{AppsColumns[2], AppsColumns[3]},
			},
		},
	}
	// AppPackagesColumns holds the columns for the "app_packages" table.
	AppPackagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "source", Type: field.TypeEnum, Enums: []string{"manual", "sentinel"}},
		{Name: "source_id", Type: field.TypeInt64},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Size: 2147483647},
		{Name: "public", Type: field.TypeBool},
		{Name: "binary_name", Type: field.TypeString, Nullable: true},
		{Name: "binary_size_bytes", Type: field.TypeInt64, Nullable: true},
		{Name: "binary_public_url", Type: field.TypeString, Nullable: true},
		{Name: "binary_sha256", Type: field.TypeBytes, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "app_app_package", Type: field.TypeInt64, Nullable: true},
		{Name: "user_app_package", Type: field.TypeInt64},
	}
	// AppPackagesTable holds the schema information for the "app_packages" table.
	AppPackagesTable = &schema.Table{
		Name:       "app_packages",
		Columns:    AppPackagesColumns,
		PrimaryKey: []*schema.Column{AppPackagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "app_packages_apps_app_package",
				Columns:    []*schema.Column{AppPackagesColumns[12]},
				RefColumns: []*schema.Column{AppsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "app_packages_users_app_package",
				Columns:    []*schema.Column{AppPackagesColumns[13]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "apppackage_binary_sha256",
				Unique:  true,
				Columns: []*schema.Column{AppPackagesColumns[9]},
			},
		},
	}
	// AppPackageRunTimesColumns holds the columns for the "app_package_run_times" table.
	AppPackageRunTimesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "app_package_id", Type: field.TypeInt64},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "run_duration", Type: field.TypeInt64},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// AppPackageRunTimesTable holds the schema information for the "app_package_run_times" table.
	AppPackageRunTimesTable = &schema.Table{
		Name:       "app_package_run_times",
		Columns:    AppPackageRunTimesColumns,
		PrimaryKey: []*schema.Column{AppPackageRunTimesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "apppackageruntime_user_id_app_package_id",
				Unique:  true,
				Columns: []*schema.Column{AppPackageRunTimesColumns[1], AppPackageRunTimesColumns[2]},
			},
		},
	}
	// FeedsColumns holds the columns for the "feeds" table.
	FeedsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "title", Type: field.TypeString, Nullable: true},
		{Name: "link", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "language", Type: field.TypeString, Nullable: true},
		{Name: "authors", Type: field.TypeJSON, Nullable: true},
		{Name: "image", Type: field.TypeJSON, Nullable: true},
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
		{Name: "name", Type: field.TypeString},
		{Name: "feed_url", Type: field.TypeString},
		{Name: "author_account", Type: field.TypeInt64},
		{Name: "source", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"active", "suspend"}},
		{Name: "category", Type: field.TypeString},
		{Name: "pull_interval", Type: field.TypeInt64},
		{Name: "hide_items", Type: field.TypeBool, Default: false},
		{Name: "latest_pull_at", Type: field.TypeTime},
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
				Columns:    []*schema.Column{FeedConfigsColumns[13]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "feedconfig_user_feed_config_feed_url",
				Unique:  true,
				Columns: []*schema.Column{FeedConfigsColumns[13], FeedConfigsColumns[2]},
			},
			{
				Name:    "feedconfig_category",
				Unique:  false,
				Columns: []*schema.Column{FeedConfigsColumns[6]},
			},
		},
	}
	// FeedItemsColumns holds the columns for the "feed_items" table.
	FeedItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "title", Type: field.TypeString, Nullable: true},
		{Name: "authors", Type: field.TypeJSON, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "content", Type: field.TypeString, Nullable: true},
		{Name: "guid", Type: field.TypeString},
		{Name: "link", Type: field.TypeString, Nullable: true},
		{Name: "image", Type: field.TypeJSON, Nullable: true},
		{Name: "published", Type: field.TypeString, Nullable: true},
		{Name: "published_parsed", Type: field.TypeTime},
		{Name: "updated", Type: field.TypeString, Nullable: true},
		{Name: "updated_parsed", Type: field.TypeTime, Nullable: true},
		{Name: "enclosures", Type: field.TypeJSON, Nullable: true},
		{Name: "publish_platform", Type: field.TypeString, Nullable: true},
		{Name: "read_count", Type: field.TypeInt64, Default: 0},
		{Name: "digest_description", Type: field.TypeString, Nullable: true},
		{Name: "digest_images", Type: field.TypeJSON, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "feed_id", Type: field.TypeInt64},
	}
	// FeedItemsTable holds the schema information for the "feed_items" table.
	FeedItemsTable = &schema.Table{
		Name:       "feed_items",
		Columns:    FeedItemsColumns,
		PrimaryKey: []*schema.Column{FeedItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "feed_items_feeds_item",
				Columns:    []*schema.Column{FeedItemsColumns[19]},
				RefColumns: []*schema.Column{FeedsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "feeditem_feed_id_guid",
				Unique:  true,
				Columns: []*schema.Column{FeedItemsColumns[19], FeedItemsColumns[5]},
			},
			{
				Name:    "feeditem_publish_platform",
				Unique:  false,
				Columns: []*schema.Column{FeedItemsColumns[13]},
			},
		},
	}
	// FilesColumns holds the columns for the "files" table.
	FilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "name", Type: field.TypeString},
		{Name: "size", Type: field.TypeInt64},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"gebura_save", "chesed_image"}},
		{Name: "sha256", Type: field.TypeBytes},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_file", Type: field.TypeInt64, Nullable: true},
	}
	// FilesTable holds the schema information for the "files" table.
	FilesTable = &schema.Table{
		Name:       "files",
		Columns:    FilesColumns,
		PrimaryKey: []*schema.Column{FilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "files_users_file",
				Columns:    []*schema.Column{FilesColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ImagesColumns holds the columns for the "images" table.
	ImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"uploaded", "scanned"}},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "file_image", Type: field.TypeInt64, Unique: true, Nullable: true},
		{Name: "user_image", Type: field.TypeInt64},
	}
	// ImagesTable holds the schema information for the "images" table.
	ImagesTable = &schema.Table{
		Name:       "images",
		Columns:    ImagesColumns,
		PrimaryKey: []*schema.Column{ImagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "images_files_image",
				Columns:    []*schema.Column{ImagesColumns[6]},
				RefColumns: []*schema.Column{FilesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "images_users_image",
				Columns:    []*schema.Column{ImagesColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// NotifyFlowsColumns holds the columns for the "notify_flows" table.
	NotifyFlowsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"active", "suspend"}},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_notify_flow", Type: field.TypeInt64},
	}
	// NotifyFlowsTable holds the schema information for the "notify_flows" table.
	NotifyFlowsTable = &schema.Table{
		Name:       "notify_flows",
		Columns:    NotifyFlowsColumns,
		PrimaryKey: []*schema.Column{NotifyFlowsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "notify_flows_users_notify_flow",
				Columns:    []*schema.Column{NotifyFlowsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// NotifyFlowSourcesColumns holds the columns for the "notify_flow_sources" table.
	NotifyFlowSourcesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "filter_include_keywords", Type: field.TypeJSON},
		{Name: "filter_exclude_keywords", Type: field.TypeJSON},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "notify_flow_id", Type: field.TypeInt64},
		{Name: "notify_source_id", Type: field.TypeInt64},
	}
	// NotifyFlowSourcesTable holds the schema information for the "notify_flow_sources" table.
	NotifyFlowSourcesTable = &schema.Table{
		Name:       "notify_flow_sources",
		Columns:    NotifyFlowSourcesColumns,
		PrimaryKey: []*schema.Column{NotifyFlowSourcesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "notify_flow_sources_notify_flows_notify_flow",
				Columns:    []*schema.Column{NotifyFlowSourcesColumns[5]},
				RefColumns: []*schema.Column{NotifyFlowsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "notify_flow_sources_feed_configs_notify_source",
				Columns:    []*schema.Column{NotifyFlowSourcesColumns[6]},
				RefColumns: []*schema.Column{FeedConfigsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "notifyflowsource_notify_flow_id_notify_source_id",
				Unique:  true,
				Columns: []*schema.Column{NotifyFlowSourcesColumns[5], NotifyFlowSourcesColumns[6]},
			},
		},
	}
	// NotifyFlowTargetsColumns holds the columns for the "notify_flow_targets" table.
	NotifyFlowTargetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "channel_id", Type: field.TypeString},
		{Name: "filter_include_keywords", Type: field.TypeJSON},
		{Name: "filter_exclude_keywords", Type: field.TypeJSON},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "notify_flow_id", Type: field.TypeInt64},
		{Name: "notify_target_id", Type: field.TypeInt64},
	}
	// NotifyFlowTargetsTable holds the schema information for the "notify_flow_targets" table.
	NotifyFlowTargetsTable = &schema.Table{
		Name:       "notify_flow_targets",
		Columns:    NotifyFlowTargetsColumns,
		PrimaryKey: []*schema.Column{NotifyFlowTargetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "notify_flow_targets_notify_flows_notify_flow",
				Columns:    []*schema.Column{NotifyFlowTargetsColumns[6]},
				RefColumns: []*schema.Column{NotifyFlowsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "notify_flow_targets_notify_targets_notify_target",
				Columns:    []*schema.Column{NotifyFlowTargetsColumns[7]},
				RefColumns: []*schema.Column{NotifyTargetsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "notifyflowtarget_notify_flow_id_notify_target_id",
				Unique:  true,
				Columns: []*schema.Column{NotifyFlowTargetsColumns[6], NotifyFlowTargetsColumns[7]},
			},
		},
	}
	// NotifyTargetsColumns holds the columns for the "notify_targets" table.
	NotifyTargetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "token", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "destination", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"active", "suspend"}},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_notify_target", Type: field.TypeInt64},
	}
	// NotifyTargetsTable holds the schema information for the "notify_targets" table.
	NotifyTargetsTable = &schema.Table{
		Name:       "notify_targets",
		Columns:    NotifyTargetsColumns,
		PrimaryKey: []*schema.Column{NotifyTargetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "notify_targets_users_notify_target",
				Columns:    []*schema.Column{NotifyTargetsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
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
		{Name: "user_created_user", Type: field.TypeInt64, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_users_created_user",
				Columns:    []*schema.Column{UsersColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AccountPurchasedAppColumns holds the columns for the "account_purchased_app" table.
	AccountPurchasedAppColumns = []*schema.Column{
		{Name: "account_id", Type: field.TypeInt64},
		{Name: "app_id", Type: field.TypeInt64},
	}
	// AccountPurchasedAppTable holds the schema information for the "account_purchased_app" table.
	AccountPurchasedAppTable = &schema.Table{
		Name:       "account_purchased_app",
		Columns:    AccountPurchasedAppColumns,
		PrimaryKey: []*schema.Column{AccountPurchasedAppColumns[0], AccountPurchasedAppColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "account_purchased_app_account_id",
				Columns:    []*schema.Column{AccountPurchasedAppColumns[0]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "account_purchased_app_app_id",
				Columns:    []*schema.Column{AccountPurchasedAppColumns[1]},
				RefColumns: []*schema.Column{AppsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserPurchasedAppColumns holds the columns for the "user_purchased_app" table.
	UserPurchasedAppColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "app_id", Type: field.TypeInt64},
	}
	// UserPurchasedAppTable holds the schema information for the "user_purchased_app" table.
	UserPurchasedAppTable = &schema.Table{
		Name:       "user_purchased_app",
		Columns:    UserPurchasedAppColumns,
		PrimaryKey: []*schema.Column{UserPurchasedAppColumns[0], UserPurchasedAppColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_purchased_app_user_id",
				Columns:    []*schema.Column{UserPurchasedAppColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_purchased_app_app_id",
				Columns:    []*schema.Column{UserPurchasedAppColumns[1]},
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
		AppPackageRunTimesTable,
		FeedsTable,
		FeedConfigsTable,
		FeedItemsTable,
		FilesTable,
		ImagesTable,
		NotifyFlowsTable,
		NotifyFlowSourcesTable,
		NotifyFlowTargetsTable,
		NotifyTargetsTable,
		UsersTable,
		AccountPurchasedAppTable,
		UserPurchasedAppTable,
	}
)

func init() {
	AccountsTable.ForeignKeys[0].RefTable = UsersTable
	AppsTable.ForeignKeys[0].RefTable = AppsTable
	AppPackagesTable.ForeignKeys[0].RefTable = AppsTable
	AppPackagesTable.ForeignKeys[1].RefTable = UsersTable
	FeedsTable.ForeignKeys[0].RefTable = FeedConfigsTable
	FeedConfigsTable.ForeignKeys[0].RefTable = UsersTable
	FeedItemsTable.ForeignKeys[0].RefTable = FeedsTable
	FilesTable.ForeignKeys[0].RefTable = UsersTable
	ImagesTable.ForeignKeys[0].RefTable = FilesTable
	ImagesTable.ForeignKeys[1].RefTable = UsersTable
	NotifyFlowsTable.ForeignKeys[0].RefTable = UsersTable
	NotifyFlowSourcesTable.ForeignKeys[0].RefTable = NotifyFlowsTable
	NotifyFlowSourcesTable.ForeignKeys[1].RefTable = FeedConfigsTable
	NotifyFlowTargetsTable.ForeignKeys[0].RefTable = NotifyFlowsTable
	NotifyFlowTargetsTable.ForeignKeys[1].RefTable = NotifyTargetsTable
	NotifyTargetsTable.ForeignKeys[0].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = UsersTable
	AccountPurchasedAppTable.ForeignKeys[0].RefTable = AccountsTable
	AccountPurchasedAppTable.ForeignKeys[1].RefTable = AppsTable
	UserPurchasedAppTable.ForeignKeys[0].RefTable = UsersTable
	UserPurchasedAppTable.ForeignKeys[1].RefTable = AppsTable
}
