// Code generated by ent, DO NOT EDIT.

package message

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the message type in the database.
	Label = "message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldMedia holds the string denoting the media field in the database.
	FieldMedia = "media"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldPic holds the string denoting the pic field in the database.
	FieldPic = "pic"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldFormID holds the string denoting the form_id field in the database.
	FieldFormID = "form_id"
	// FieldTargetID holds the string denoting the target_id field in the database.
	FieldTargetID = "target_id"
	// EdgeSender holds the string denoting the sender edge name in mutations.
	EdgeSender = "sender"
	// EdgeRecipient holds the string denoting the recipient edge name in mutations.
	EdgeRecipient = "recipient"
	// Table holds the table name of the message in the database.
	Table = "t_message"
	// SenderTable is the table that holds the sender relation/edge.
	SenderTable = "t_message"
	// SenderInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	SenderInverseTable = "t_user"
	// SenderColumn is the table column denoting the sender relation/edge.
	SenderColumn = "form_id"
	// RecipientTable is the table that holds the recipient relation/edge.
	RecipientTable = "t_message"
	// RecipientInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	RecipientInverseTable = "t_user"
	// RecipientColumn is the table column denoting the recipient relation/edge.
	RecipientColumn = "target_id"
)

// Columns holds all SQL columns for message fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldType,
	FieldMedia,
	FieldContent,
	FieldPic,
	FieldURL,
	FieldDesc,
	FieldAmount,
	FieldFormID,
	FieldTargetID,
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
	// DefaultContent holds the default value on creation for the "content" field.
	DefaultContent string
	// DefaultPic holds the default value on creation for the "pic" field.
	DefaultPic string
	// DefaultURL holds the default value on creation for the "url" field.
	DefaultURL string
	// DefaultDesc holds the default value on creation for the "desc" field.
	DefaultDesc string
)

// OrderOption defines the ordering options for the Message queries.
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

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByMedia orders the results by the media field.
func ByMedia(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMedia, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByPic orders the results by the pic field.
func ByPic(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPic, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByDesc orders the results by the desc field.
func ByDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDesc, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByFormID orders the results by the form_id field.
func ByFormID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFormID, opts...).ToFunc()
}

// ByTargetID orders the results by the target_id field.
func ByTargetID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTargetID, opts...).ToFunc()
}

// BySenderField orders the results by sender field.
func BySenderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSenderStep(), sql.OrderByField(field, opts...))
	}
}

// ByRecipientField orders the results by recipient field.
func ByRecipientField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRecipientStep(), sql.OrderByField(field, opts...))
	}
}
func newSenderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SenderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SenderTable, SenderColumn),
	)
}
func newRecipientStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RecipientInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, RecipientTable, RecipientColumn),
	)
}
