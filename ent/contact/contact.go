// Code generated by ent, DO NOT EDIT.

package contact

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the contact type in the database.
	Label = "contact"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldOwnerID holds the string denoting the owner_id field in the database.
	FieldOwnerID = "owner_id"
	// FieldTargetID holds the string denoting the target_id field in the database.
	FieldTargetID = "target_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// EdgeOwnerUser holds the string denoting the owner_user edge name in mutations.
	EdgeOwnerUser = "owner_user"
	// EdgeTargetUser holds the string denoting the target_user edge name in mutations.
	EdgeTargetUser = "target_user"
	// Table holds the table name of the contact in the database.
	Table = "t_contact"
	// OwnerUserTable is the table that holds the owner_user relation/edge.
	OwnerUserTable = "t_contact"
	// OwnerUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerUserInverseTable = "t_user"
	// OwnerUserColumn is the table column denoting the owner_user relation/edge.
	OwnerUserColumn = "owner_id"
	// TargetUserTable is the table that holds the target_user relation/edge.
	TargetUserTable = "t_contact"
	// TargetUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	TargetUserInverseTable = "t_user"
	// TargetUserColumn is the table column denoting the target_user relation/edge.
	TargetUserColumn = "target_id"
)

// Columns holds all SQL columns for contact fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldOwnerID,
	FieldTargetID,
	FieldType,
	FieldDesc,
}

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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType int
	// DefaultDesc holds the default value on creation for the "desc" field.
	DefaultDesc string
)

// OrderOption defines the ordering options for the Contact queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByOwnerID orders the results by the owner_id field.
func ByOwnerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwnerID, opts...).ToFunc()
}

// ByTargetID orders the results by the target_id field.
func ByTargetID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTargetID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByDesc orders the results by the desc field.
func ByDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDesc, opts...).ToFunc()
}

// ByOwnerUserField orders the results by owner_user field.
func ByOwnerUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByTargetUserField orders the results by target_user field.
func ByTargetUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTargetUserStep(), sql.OrderByField(field, opts...))
	}
}
func newOwnerUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerUserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerUserTable, OwnerUserColumn),
	)
}
func newTargetUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TargetUserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TargetUserTable, TargetUserColumn),
	)
}
