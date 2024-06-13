// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccessTokensColumns holds the columns for the "access_tokens" table.
	AccessTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "name", Type: field.TypeString},
		{Name: "token", Type: field.TypeString},
		{Name: "token_hash", Type: field.TypeString, Unique: true},
		{Name: "token_salt", Type: field.TypeString},
		{Name: "token_last_eight", Type: field.TypeString},
		{Name: "scope", Type: field.TypeString},
		{Name: "has_recent_activity", Type: field.TypeString},
		{Name: "has_used", Type: field.TypeString},
		{Name: "create_time", Type: field.TypeInt64, Default: 1718295427},
	}
	// AccessTokensTable holds the schema information for the "access_tokens" table.
	AccessTokensTable = &schema.Table{
		Name:       "access_tokens",
		Columns:    AccessTokensColumns,
		PrimaryKey: []*schema.Column{AccessTokensColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "accesstoken_user_id",
				Unique:  false,
				Columns: []*schema.Column{AccessTokensColumns[1]},
			},
			{
				Name:    "accesstoken_token_last_eight",
				Unique:  false,
				Columns: []*schema.Column{AccessTokensColumns[6]},
			},
			{
				Name:    "accesstoken_create_time",
				Unique:  false,
				Columns: []*schema.Column{AccessTokensColumns[10]},
			},
		},
	}
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "account", Type: field.TypeString},
		{Name: "type", Type: field.TypeUint8},
		{Name: "desc", Type: field.TypeString},
		{Name: "is_private", Type: field.TypeBool, Default: true},
		{Name: "is_activated", Type: field.TypeBool},
		{Name: "is_primary", Type: field.TypeBool, Default: false},
		{Name: "create_time", Type: field.TypeInt64, Default: 1718295427},
		{Name: "update_time", Type: field.TypeInt64},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "accounts_user_id",
				Unique:  false,
				Columns: []*schema.Column{AccountsColumns[1]},
			},
			{
				Name:    "accounts_account_type",
				Unique:  true,
				Columns: []*schema.Column{AccountsColumns[2], AccountsColumns[3]},
			},
			{
				Name:    "accounts_create_time",
				Unique:  false,
				Columns: []*schema.Column{AccountsColumns[8]},
			},
			{
				Name:    "accounts_update_time",
				Unique:  false,
				Columns: []*schema.Column{AccountsColumns[9]},
			},
		},
	}
	// MemberRoleRelatedPermissionsColumns holds the columns for the "member_role_related_permissions" table.
	MemberRoleRelatedPermissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "role_id", Type: field.TypeUUID},
		{Name: "permission_group_id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeInt64, Default: 1718295427},
		{Name: "update_time", Type: field.TypeInt64},
	}
	// MemberRoleRelatedPermissionsTable holds the schema information for the "member_role_related_permissions" table.
	MemberRoleRelatedPermissionsTable = &schema.Table{
		Name:       "member_role_related_permissions",
		Columns:    MemberRoleRelatedPermissionsColumns,
		PrimaryKey: []*schema.Column{MemberRoleRelatedPermissionsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "memberrolerelatedpermission_role_id_permission_group_id",
				Unique:  true,
				Columns: []*schema.Column{MemberRoleRelatedPermissionsColumns[1], MemberRoleRelatedPermissionsColumns[2]},
			},
			{
				Name:    "memberrolerelatedpermission_create_time",
				Unique:  false,
				Columns: []*schema.Column{MemberRoleRelatedPermissionsColumns[3]},
			},
			{
				Name:    "memberrolerelatedpermission_update_time",
				Unique:  false,
				Columns: []*schema.Column{MemberRoleRelatedPermissionsColumns[4]},
			},
		},
	}
	// PermissionGroupsColumns holds the columns for the "permission_groups" table.
	PermissionGroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "permission_name", Type: field.TypeString, Size: 32, Default: ""},
		{Name: "ioc", Type: field.TypeString, Size: 254, Default: ""},
		{Name: "sort", Type: field.TypeInt32, Default: 0},
		{Name: "left", Type: field.TypeInt32, Default: 0},
		{Name: "right", Type: field.TypeInt32, Default: 0},
		{Name: "state", Type: field.TypeInt, Default: 1},
		{Name: "create_time", Type: field.TypeInt64, Default: 1718295427},
		{Name: "update_time", Type: field.TypeInt64},
	}
	// PermissionGroupsTable holds the schema information for the "permission_groups" table.
	PermissionGroupsTable = &schema.Table{
		Name:       "permission_groups",
		Columns:    PermissionGroupsColumns,
		PrimaryKey: []*schema.Column{PermissionGroupsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "permissiongroup_left_right",
				Unique:  false,
				Columns: []*schema.Column{PermissionGroupsColumns[4], PermissionGroupsColumns[5]},
			},
			{
				Name:    "permissiongroup_state_sort_permission_name",
				Unique:  false,
				Columns: []*schema.Column{PermissionGroupsColumns[6], PermissionGroupsColumns[3], PermissionGroupsColumns[1]},
			},
			{
				Name:    "permissiongroup_create_time",
				Unique:  false,
				Columns: []*schema.Column{PermissionGroupsColumns[7]},
			},
			{
				Name:    "permissiongroup_update_time",
				Unique:  false,
				Columns: []*schema.Column{PermissionGroupsColumns[8]},
			},
		},
	}
	// PermissionRelatedRoutersColumns holds the columns for the "permission_related_routers" table.
	PermissionRelatedRoutersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "router_id", Type: field.TypeUUID},
		{Name: "permission_group_id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeInt64, Default: 1718295427},
		{Name: "update_time", Type: field.TypeInt64},
	}
	// PermissionRelatedRoutersTable holds the schema information for the "permission_related_routers" table.
	PermissionRelatedRoutersTable = &schema.Table{
		Name:       "permission_related_routers",
		Columns:    PermissionRelatedRoutersColumns,
		PrimaryKey: []*schema.Column{PermissionRelatedRoutersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "permissionrelatedrouter_permission_group_id_router_id",
				Unique:  false,
				Columns: []*schema.Column{PermissionRelatedRoutersColumns[2], PermissionRelatedRoutersColumns[1]},
			},
			{
				Name:    "permissionrelatedrouter_create_time",
				Unique:  false,
				Columns: []*schema.Column{PermissionRelatedRoutersColumns[3]},
			},
			{
				Name:    "permissionrelatedrouter_update_time",
				Unique:  false,
				Columns: []*schema.Column{PermissionRelatedRoutersColumns[4]},
			},
		},
	}
	// RoutersColumns holds the columns for the "routers" table.
	RoutersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "route_name", Type: field.TypeString, Size: 32, Default: ""},
		{Name: "route", Type: field.TypeString, Size: 254, Default: ""},
		{Name: "description", Type: field.TypeString, Size: 254, Default: ""},
		{Name: "state", Type: field.TypeInt, Default: 1},
		{Name: "create_time", Type: field.TypeInt64, Default: 1718295427},
		{Name: "update_time", Type: field.TypeInt64},
	}
	// RoutersTable holds the schema information for the "routers" table.
	RoutersTable = &schema.Table{
		Name:       "routers",
		Columns:    RoutersColumns,
		PrimaryKey: []*schema.Column{RoutersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "router_state_route_name",
				Unique:  false,
				Columns: []*schema.Column{RoutersColumns[4], RoutersColumns[1]},
			},
			{
				Name:    "router_state_route",
				Unique:  false,
				Columns: []*schema.Column{RoutersColumns[4], RoutersColumns[2]},
			},
			{
				Name:    "router_create_time",
				Unique:  false,
				Columns: []*schema.Column{RoutersColumns[5]},
			},
			{
				Name:    "router_update_time",
				Unique:  false,
				Columns: []*schema.Column{RoutersColumns[6]},
			},
		},
	}
	// SourcesColumns holds the columns for the "sources" table.
	SourcesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "type", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString},
		{Name: "is_active", Type: field.TypeBool},
		{Name: "is_sync_enabled", Type: field.TypeBool},
		{Name: "cfg", Type: field.TypeJSON},
		{Name: "create_time", Type: field.TypeInt64, Default: 1718295427},
		{Name: "update_time", Type: field.TypeInt64},
	}
	// SourcesTable holds the schema information for the "sources" table.
	SourcesTable = &schema.Table{
		Name:       "sources",
		Columns:    SourcesColumns,
		PrimaryKey: []*schema.Column{SourcesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "source_is_active_is_sync_enabled",
				Unique:  false,
				Columns: []*schema.Column{SourcesColumns[3], SourcesColumns[4]},
			},
			{
				Name:    "source_create_time",
				Unique:  false,
				Columns: []*schema.Column{SourcesColumns[6]},
			},
			{
				Name:    "source_update_time",
				Unique:  false,
				Columns: []*schema.Column{SourcesColumns[7]},
			},
			{
				Name:    "source_name",
				Unique:  true,
				Columns: []*schema.Column{SourcesColumns[2]},
			},
			{
				Name:    "source_create_time",
				Unique:  false,
				Columns: []*schema.Column{SourcesColumns[6]},
			},
			{
				Name:    "source_update_time",
				Unique:  false,
				Columns: []*schema.Column{SourcesColumns[7]},
			},
		},
	}
	// SourceDataColumns holds the columns for the "source_data" table.
	SourceDataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "type", Type: field.TypeString, Size: 32, Default: ""},
		{Name: "sub_type", Type: field.TypeString, Size: 32, Default: ""},
		{Name: "info", Type: field.TypeString, Default: ""},
		{Name: "snapshot", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "create_time", Type: field.TypeInt64, Default: 1718295427},
		{Name: "update_time", Type: field.TypeInt64},
	}
	// SourceDataTable holds the schema information for the "source_data" table.
	SourceDataTable = &schema.Table{
		Name:       "source_data",
		Columns:    SourceDataColumns,
		PrimaryKey: []*schema.Column{SourceDataColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "sourcedata_user_id",
				Unique:  false,
				Columns: []*schema.Column{SourceDataColumns[1]},
			},
			{
				Name:    "sourcedata_type_sub_type",
				Unique:  false,
				Columns: []*schema.Column{SourceDataColumns[2], SourceDataColumns[3]},
			},
			{
				Name:    "sourcedata_create_time",
				Unique:  false,
				Columns: []*schema.Column{SourceDataColumns[6]},
			},
			{
				Name:    "sourcedata_update_time",
				Unique:  false,
				Columns: []*schema.Column{SourceDataColumns[7]},
			},
		},
	}
	// TwoFactorsColumns holds the columns for the "two_factors" table.
	TwoFactorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "user_id", Type: field.TypeInt64, Unique: true},
		{Name: "secret", Type: field.TypeString},
		{Name: "scratch_salt", Type: field.TypeString},
		{Name: "scratch_hash", Type: field.TypeString},
		{Name: "last_used_passcode", Type: field.TypeString},
		{Name: "create_time", Type: field.TypeInt64, Default: 1718295427},
		{Name: "update_time", Type: field.TypeInt64},
	}
	// TwoFactorsTable holds the schema information for the "two_factors" table.
	TwoFactorsTable = &schema.Table{
		Name:       "two_factors",
		Columns:    TwoFactorsColumns,
		PrimaryKey: []*schema.Column{TwoFactorsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "twofactor_create_time",
				Unique:  false,
				Columns: []*schema.Column{TwoFactorsColumns[6]},
			},
			{
				Name:    "twofactor_update_time",
				Unique:  false,
				Columns: []*schema.Column{TwoFactorsColumns[7]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "avatar", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Size: 128},
		{Name: "name", Type: field.TypeString, Size: 32},
		{Name: "lower_name", Type: field.TypeString, Size: 32},
		{Name: "full_name", Type: field.TypeString, Size: 32},
		{Name: "passwd_salt", Type: field.TypeString, Size: 120, Default: "argon2"},
		{Name: "passwd_hash_algo", Type: field.TypeString, Size: 32},
		{Name: "passwd", Type: field.TypeString, Size: 32},
		{Name: "language", Type: field.TypeString, Size: 32},
		{Name: "login_name", Type: field.TypeString},
		{Name: "login_source", Type: field.TypeInt64, Default: 0},
		{Name: "login_type", Type: field.TypeInt},
		{Name: "is_restricted", Type: field.TypeBool, Default: false},
		{Name: "is_active", Type: field.TypeBool, Comment: "true is activated", Default: false},
		{Name: "prohibit_login", Type: field.TypeBool, Comment: "is web login", Default: false},
		{Name: "create_time", Type: field.TypeInt64, Default: 1718295427},
		{Name: "update_time", Type: field.TypeInt64},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Comment:    "member",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_email",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[2]},
			},
			{
				Name:    "user_create_time",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[16]},
			},
			{
				Name:    "user_update_time",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[17]},
			},
		},
	}
	// UserAuthSourcesColumns holds the columns for the "user_auth_sources" table.
	UserAuthSourcesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "token", Type: field.TypeString, Size: 254},
		{Name: "channel", Type: field.TypeString, Size: 64},
		{Name: "device", Type: field.TypeString, Size: 64},
		{Name: "device_detail", Type: field.TypeString, Size: 64},
		{Name: "client_ip", Type: field.TypeString, Size: 32},
		{Name: "remote_ip", Type: field.TypeString, Size: 32},
		{Name: "snapshot", Type: field.TypeString, Size: 254},
		{Name: "login_name", Type: field.TypeString},
		{Name: "login_source", Type: field.TypeInt, Default: 0},
		{Name: "login_type", Type: field.TypeInt, Default: 0},
		{Name: "create_time", Type: field.TypeInt64, Default: 1718295427},
		{Name: "update_time", Type: field.TypeInt64},
	}
	// UserAuthSourcesTable holds the schema information for the "user_auth_sources" table.
	UserAuthSourcesTable = &schema.Table{
		Name:       "user_auth_sources",
		Columns:    UserAuthSourcesColumns,
		PrimaryKey: []*schema.Column{UserAuthSourcesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "userauthsource_user_id_channel_device",
				Unique:  false,
				Columns: []*schema.Column{UserAuthSourcesColumns[1], UserAuthSourcesColumns[3], UserAuthSourcesColumns[4]},
			},
			{
				Name:    "userauthsource_create_time",
				Unique:  false,
				Columns: []*schema.Column{UserAuthSourcesColumns[12]},
			},
			{
				Name:    "userauthsource_update_time",
				Unique:  false,
				Columns: []*schema.Column{UserAuthSourcesColumns[13]},
			},
		},
	}
	// UserRelatedRolesColumns holds the columns for the "user_related_roles" table.
	UserRelatedRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "role_id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeInt64, Default: 1718295427},
		{Name: "update_time", Type: field.TypeInt64},
	}
	// UserRelatedRolesTable holds the schema information for the "user_related_roles" table.
	UserRelatedRolesTable = &schema.Table{
		Name:       "user_related_roles",
		Columns:    UserRelatedRolesColumns,
		PrimaryKey: []*schema.Column{UserRelatedRolesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "userrelatedrole_user_id_role_id",
				Unique:  true,
				Columns: []*schema.Column{UserRelatedRolesColumns[1], UserRelatedRolesColumns[2]},
			},
			{
				Name:    "userrelatedrole_create_time",
				Unique:  false,
				Columns: []*schema.Column{UserRelatedRolesColumns[3]},
			},
			{
				Name:    "userrelatedrole_update_time",
				Unique:  false,
				Columns: []*schema.Column{UserRelatedRolesColumns[4]},
			},
		},
	}
	// UserRolesColumns holds the columns for the "user_roles" table.
	UserRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "role_name", Type: field.TypeString, Size: 64},
		{Name: "state", Type: field.TypeInt, Default: 1},
		{Name: "create_time", Type: field.TypeInt64, Default: 1718295427},
		{Name: "update_time", Type: field.TypeInt64},
	}
	// UserRolesTable holds the schema information for the "user_roles" table.
	UserRolesTable = &schema.Table{
		Name:       "user_roles",
		Columns:    UserRolesColumns,
		PrimaryKey: []*schema.Column{UserRolesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "userrole_state_role_name",
				Unique:  true,
				Columns: []*schema.Column{UserRolesColumns[2], UserRolesColumns[1]},
			},
			{
				Name:    "userrole_create_time",
				Unique:  false,
				Columns: []*schema.Column{UserRolesColumns[3]},
			},
			{
				Name:    "userrole_update_time",
				Unique:  false,
				Columns: []*schema.Column{UserRolesColumns[4]},
			},
		},
	}
	// WakatimesColumns holds the columns for the "wakatimes" table.
	WakatimesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "key", Type: field.TypeString},
		{Name: "api", Type: field.TypeString},
		{Name: "state", Type: field.TypeString},
		{Name: "create_time", Type: field.TypeInt64, Default: 1718295427},
		{Name: "update_time", Type: field.TypeInt64},
	}
	// WakatimesTable holds the schema information for the "wakatimes" table.
	WakatimesTable = &schema.Table{
		Name:       "wakatimes",
		Columns:    WakatimesColumns,
		PrimaryKey: []*schema.Column{WakatimesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "wakatime_user_id",
				Unique:  false,
				Columns: []*schema.Column{WakatimesColumns[1]},
			},
			{
				Name:    "wakatime_key",
				Unique:  false,
				Columns: []*schema.Column{WakatimesColumns[2]},
			},
			{
				Name:    "wakatime_create_time",
				Unique:  false,
				Columns: []*schema.Column{WakatimesColumns[5]},
			},
			{
				Name:    "wakatime_update_time",
				Unique:  false,
				Columns: []*schema.Column{WakatimesColumns[6]},
			},
		},
	}
	// WakatimeCategoriesColumns holds the columns for the "wakatime_categories" table.
	WakatimeCategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "wakatime_id", Type: field.TypeInt64, Comment: "wakatime id"},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "name", Type: field.TypeString, Comment: "名称", Default: ""},
		{Name: "total_seconds", Type: field.TypeInt64, Comment: "总时长(秒", Default: 0},
		{Name: "create_time", Type: field.TypeInt64, Default: 1718295427},
		{Name: "update_time", Type: field.TypeInt64},
	}
	// WakatimeCategoriesTable holds the schema information for the "wakatime_categories" table.
	WakatimeCategoriesTable = &schema.Table{
		Name:       "wakatime_categories",
		Comment:    "每日category统计表",
		Columns:    WakatimeCategoriesColumns,
		PrimaryKey: []*schema.Column{WakatimeCategoriesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "wakatimecategory_wakatime_id",
				Unique:  false,
				Columns: []*schema.Column{WakatimeCategoriesColumns[1]},
			},
			{
				Name:    "wakatimecategory_user_id_name",
				Unique:  false,
				Columns: []*schema.Column{WakatimeCategoriesColumns[2], WakatimeCategoriesColumns[3]},
			},
		},
	}
	// WakatimeDependenciesColumns holds the columns for the "wakatime_dependencies" table.
	WakatimeDependenciesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "wakatime_id", Type: field.TypeInt64, Comment: "wakatime id"},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "name", Type: field.TypeString, Comment: "名称", Default: ""},
		{Name: "total_seconds", Type: field.TypeInt64, Comment: "总时长(秒", Default: 0},
		{Name: "create_time", Type: field.TypeInt64, Default: 1718295427},
		{Name: "update_time", Type: field.TypeInt64},
	}
	// WakatimeDependenciesTable holds the schema information for the "wakatime_dependencies" table.
	WakatimeDependenciesTable = &schema.Table{
		Name:       "wakatime_dependencies",
		Comment:    "每日dependency统计表",
		Columns:    WakatimeDependenciesColumns,
		PrimaryKey: []*schema.Column{WakatimeDependenciesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "wakatimedependency_wakatime_id",
				Unique:  false,
				Columns: []*schema.Column{WakatimeDependenciesColumns[1]},
			},
			{
				Name:    "wakatimedependency_user_id_name",
				Unique:  false,
				Columns: []*schema.Column{WakatimeDependenciesColumns[2], WakatimeDependenciesColumns[3]},
			},
			{
				Name:    "wakatimedependency_create_time",
				Unique:  false,
				Columns: []*schema.Column{WakatimeDependenciesColumns[5]},
			},
			{
				Name:    "wakatimedependency_update_time",
				Unique:  false,
				Columns: []*schema.Column{WakatimeDependenciesColumns[6]},
			},
		},
	}
	// WakatimeDurationsColumns holds the columns for the "wakatime_durations" table.
	WakatimeDurationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
	}
	// WakatimeDurationsTable holds the schema information for the "wakatime_durations" table.
	WakatimeDurationsTable = &schema.Table{
		Name:       "wakatime_durations",
		Columns:    WakatimeDurationsColumns,
		PrimaryKey: []*schema.Column{WakatimeDurationsColumns[0]},
	}
	// WakatimeEditorsColumns holds the columns for the "wakatime_editors" table.
	WakatimeEditorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
	}
	// WakatimeEditorsTable holds the schema information for the "wakatime_editors" table.
	WakatimeEditorsTable = &schema.Table{
		Name:       "wakatime_editors",
		Columns:    WakatimeEditorsColumns,
		PrimaryKey: []*schema.Column{WakatimeEditorsColumns[0]},
	}
	// WakatimeEntitiesColumns holds the columns for the "wakatime_entities" table.
	WakatimeEntitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
	}
	// WakatimeEntitiesTable holds the schema information for the "wakatime_entities" table.
	WakatimeEntitiesTable = &schema.Table{
		Name:       "wakatime_entities",
		Columns:    WakatimeEntitiesColumns,
		PrimaryKey: []*schema.Column{WakatimeEntitiesColumns[0]},
	}
	// WakatimeGrandTotalsColumns holds the columns for the "wakatime_grand_totals" table.
	WakatimeGrandTotalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
	}
	// WakatimeGrandTotalsTable holds the schema information for the "wakatime_grand_totals" table.
	WakatimeGrandTotalsTable = &schema.Table{
		Name:       "wakatime_grand_totals",
		Columns:    WakatimeGrandTotalsColumns,
		PrimaryKey: []*schema.Column{WakatimeGrandTotalsColumns[0]},
	}
	// WakatimeHeartBeatsColumns holds the columns for the "wakatime_heart_beats" table.
	WakatimeHeartBeatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
	}
	// WakatimeHeartBeatsTable holds the schema information for the "wakatime_heart_beats" table.
	WakatimeHeartBeatsTable = &schema.Table{
		Name:       "wakatime_heart_beats",
		Columns:    WakatimeHeartBeatsColumns,
		PrimaryKey: []*schema.Column{WakatimeHeartBeatsColumns[0]},
	}
	// WakatimeLanguagesColumns holds the columns for the "wakatime_languages" table.
	WakatimeLanguagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
	}
	// WakatimeLanguagesTable holds the schema information for the "wakatime_languages" table.
	WakatimeLanguagesTable = &schema.Table{
		Name:       "wakatime_languages",
		Columns:    WakatimeLanguagesColumns,
		PrimaryKey: []*schema.Column{WakatimeLanguagesColumns[0]},
	}
	// WakatimeProjectsColumns holds the columns for the "wakatime_projects" table.
	WakatimeProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
	}
	// WakatimeProjectsTable holds the schema information for the "wakatime_projects" table.
	WakatimeProjectsTable = &schema.Table{
		Name:       "wakatime_projects",
		Columns:    WakatimeProjectsColumns,
		PrimaryKey: []*schema.Column{WakatimeProjectsColumns[0]},
	}
	// WakatimeProjectDurationsColumns holds the columns for the "wakatime_project_durations" table.
	WakatimeProjectDurationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
	}
	// WakatimeProjectDurationsTable holds the schema information for the "wakatime_project_durations" table.
	WakatimeProjectDurationsTable = &schema.Table{
		Name:       "wakatime_project_durations",
		Columns:    WakatimeProjectDurationsColumns,
		PrimaryKey: []*schema.Column{WakatimeProjectDurationsColumns[0]},
	}
	// WakatimeProjectInfosColumns holds the columns for the "wakatime_project_infos" table.
	WakatimeProjectInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
	}
	// WakatimeProjectInfosTable holds the schema information for the "wakatime_project_infos" table.
	WakatimeProjectInfosTable = &schema.Table{
		Name:       "wakatime_project_infos",
		Columns:    WakatimeProjectInfosColumns,
		PrimaryKey: []*schema.Column{WakatimeProjectInfosColumns[0]},
	}
	// WakatimeSystemsColumns holds the columns for the "wakatime_systems" table.
	WakatimeSystemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
	}
	// WakatimeSystemsTable holds the schema information for the "wakatime_systems" table.
	WakatimeSystemsTable = &schema.Table{
		Name:       "wakatime_systems",
		Columns:    WakatimeSystemsColumns,
		PrimaryKey: []*schema.Column{WakatimeSystemsColumns[0]},
	}
	// WebAuthnCredentialsColumns holds the columns for the "web_authn_credentials" table.
	WebAuthnCredentialsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "lower_name", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeInt64, Unique: true},
		{Name: "credential_id", Type: field.TypeBytes, Size: 1024},
		{Name: "public_key", Type: field.TypeBytes},
		{Name: "attestation_type", Type: field.TypeString},
		{Name: "aaguid", Type: field.TypeBytes},
		{Name: "sign_count", Type: field.TypeUint32},
		{Name: "clone_warning", Type: field.TypeBool},
		{Name: "create_time", Type: field.TypeInt64, Default: 1718295427},
		{Name: "update_time", Type: field.TypeInt64},
	}
	// WebAuthnCredentialsTable holds the schema information for the "web_authn_credentials" table.
	WebAuthnCredentialsTable = &schema.Table{
		Name:       "web_authn_credentials",
		Columns:    WebAuthnCredentialsColumns,
		PrimaryKey: []*schema.Column{WebAuthnCredentialsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "webauthncredential_create_time",
				Unique:  false,
				Columns: []*schema.Column{WebAuthnCredentialsColumns[10]},
			},
			{
				Name:    "webauthncredential_update_time",
				Unique:  false,
				Columns: []*schema.Column{WebAuthnCredentialsColumns[11]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccessTokensTable,
		AccountsTable,
		MemberRoleRelatedPermissionsTable,
		PermissionGroupsTable,
		PermissionRelatedRoutersTable,
		RoutersTable,
		SourcesTable,
		SourceDataTable,
		TwoFactorsTable,
		UsersTable,
		UserAuthSourcesTable,
		UserRelatedRolesTable,
		UserRolesTable,
		WakatimesTable,
		WakatimeCategoriesTable,
		WakatimeDependenciesTable,
		WakatimeDurationsTable,
		WakatimeEditorsTable,
		WakatimeEntitiesTable,
		WakatimeGrandTotalsTable,
		WakatimeHeartBeatsTable,
		WakatimeLanguagesTable,
		WakatimeProjectsTable,
		WakatimeProjectDurationsTable,
		WakatimeProjectInfosTable,
		WakatimeSystemsTable,
		WebAuthnCredentialsTable,
	}
)

func init() {
	UsersTable.Annotation = &entsql.Annotation{}
	WakatimeCategoriesTable.Annotation = &entsql.Annotation{}
	WakatimeDependenciesTable.Annotation = &entsql.Annotation{}
}
