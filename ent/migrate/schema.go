// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CoinsColumns holds the columns for the "coins" table.
	CoinsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "deal", Type: field.TypeString},
		{Name: "amount", Type: field.TypeInt64},
		{Name: "balance", Type: field.TypeInt64},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "updated_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "user_coins", Type: field.TypeInt},
	}
	// CoinsTable holds the schema information for the "coins" table.
	CoinsTable = &schema.Table{
		Name:       "coins",
		Columns:    CoinsColumns,
		PrimaryKey: []*schema.Column{CoinsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "coins_users_coins",
				Columns:    []*schema.Column{CoinsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "updated_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "comment_author", Type: field.TypeInt},
		{Name: "comment_to", Type: field.TypeInt, Nullable: true},
		{Name: "post_comments", Type: field.TypeInt},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comments_users_author",
				Columns:    []*schema.Column{CommentsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "comments_users_to",
				Columns:    []*schema.Column{CommentsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "comments_posts_comments",
				Columns:    []*schema.Column{CommentsColumns[6]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "intro", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "updated_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "updated_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "group_posts", Type: field.TypeInt, Nullable: true},
		{Name: "user_posts", Type: field.TypeInt},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_groups_posts",
				Columns:    []*schema.Column{PostsColumns[4]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "posts_users_posts",
				Columns:    []*schema.Column{PostsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "account", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "salt", Type: field.TypeString, Nullable: true},
		{Name: "avatar_url", Type: field.TypeString, Nullable: true},
		{Name: "intro", Type: field.TypeString, Nullable: true},
		{Name: "coin", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "updated_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserFavoritePostsColumns holds the columns for the "user_favorite_posts" table.
	UserFavoritePostsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "post_id", Type: field.TypeInt},
	}
	// UserFavoritePostsTable holds the schema information for the "user_favorite_posts" table.
	UserFavoritePostsTable = &schema.Table{
		Name:       "user_favorite_posts",
		Columns:    UserFavoritePostsColumns,
		PrimaryKey: []*schema.Column{UserFavoritePostsColumns[0], UserFavoritePostsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_favorite_posts_user_id",
				Columns:    []*schema.Column{UserFavoritePostsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_favorite_posts_post_id",
				Columns:    []*schema.Column{UserFavoritePostsColumns[1]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CoinsTable,
		CommentsTable,
		GroupsTable,
		PostsTable,
		UsersTable,
		UserFavoritePostsTable,
	}
)

func init() {
	CoinsTable.ForeignKeys[0].RefTable = UsersTable
	CommentsTable.ForeignKeys[0].RefTable = UsersTable
	CommentsTable.ForeignKeys[1].RefTable = UsersTable
	CommentsTable.ForeignKeys[2].RefTable = PostsTable
	PostsTable.ForeignKeys[0].RefTable = GroupsTable
	PostsTable.ForeignKeys[1].RefTable = UsersTable
	UserFavoritePostsTable.ForeignKeys[0].RefTable = UsersTable
	UserFavoritePostsTable.ForeignKeys[1].RefTable = PostsTable
}
