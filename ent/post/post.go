// Code generated by ent, DO NOT EDIT.

package post

import (
	"time"
)

const (
	// Label holds the string label denoting the post type in the database.
	Label = "post"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// EdgeGroup holds the string denoting the group edge name in mutations.
	EdgeGroup = "group"
	// EdgeFavoriteUsers holds the string denoting the favorite_users edge name in mutations.
	EdgeFavoriteUsers = "favorite_users"
	// EdgeComments holds the string denoting the comments edge name in mutations.
	EdgeComments = "comments"
	// Table holds the table name of the post in the database.
	Table = "posts"
	// AuthorTable is the table that holds the author relation/edge.
	AuthorTable = "posts"
	// AuthorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AuthorInverseTable = "users"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "user_posts"
	// GroupTable is the table that holds the group relation/edge.
	GroupTable = "posts"
	// GroupInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupInverseTable = "groups"
	// GroupColumn is the table column denoting the group relation/edge.
	GroupColumn = "group_posts"
	// FavoriteUsersTable is the table that holds the favorite_users relation/edge. The primary key declared below.
	FavoriteUsersTable = "user_favorite_posts"
	// FavoriteUsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	FavoriteUsersInverseTable = "users"
	// CommentsTable is the table that holds the comments relation/edge.
	CommentsTable = "comments"
	// CommentsInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentsInverseTable = "comments"
	// CommentsColumn is the table column denoting the comments relation/edge.
	CommentsColumn = "post_comments"
)

// Columns holds all SQL columns for post fields.
var Columns = []string{
	FieldID,
	FieldContent,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "posts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"group_posts",
	"user_posts",
}

var (
	// FavoriteUsersPrimaryKey and FavoriteUsersColumn2 are the table columns denoting the
	// primary key for the favorite_users relation (M2M).
	FavoriteUsersPrimaryKey = []string{"user_id", "post_id"}
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
