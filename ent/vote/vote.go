// Code generated by ent, DO NOT EDIT.

package vote

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the vote type in the database.
	Label = "vote"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSelection holds the string denoting the selection field in the database.
	FieldSelection = "selection"
	// FieldRank holds the string denoting the rank field in the database.
	FieldRank = "rank"
	// FieldVoteID holds the string denoting the vote_id field in the database.
	FieldVoteID = "vote_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the vote in the database.
	Table = "votes"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "votes"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_votes"
)

// Columns holds all SQL columns for vote fields.
var Columns = []string{
	FieldID,
	FieldSelection,
	FieldRank,
	FieldVoteID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "votes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_votes",
}

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
	// SelectionValidator is a validator for the "selection" field. It is called by the builders before save.
	SelectionValidator func(string) error
	// RankValidator is a validator for the "rank" field. It is called by the builders before save.
	RankValidator func(int) error
	// VoteIDValidator is a validator for the "vote_id" field. It is called by the builders before save.
	VoteIDValidator func(string) error
)

// OrderOption defines the ordering options for the Vote queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySelection orders the results by the selection field.
func BySelection(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSelection, opts...).ToFunc()
}

// ByRank orders the results by the rank field.
func ByRank(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRank, opts...).ToFunc()
}

// ByVoteID orders the results by the vote_id field.
func ByVoteID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVoteID, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
