// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldClientIP holds the string denoting the client_ip field in the database.
	FieldClientIP = "client_ip"
	// FieldClientPort holds the string denoting the client_port field in the database.
	FieldClientPort = "client_port"
	// FieldSalt holds the string denoting the salt field in the database.
	FieldSalt = "salt"
	// FieldLoginTime holds the string denoting the login_time field in the database.
	FieldLoginTime = "login_time"
	// FieldHeartbeatTime holds the string denoting the heartbeat_time field in the database.
	FieldHeartbeatTime = "heartbeat_time"
	// FieldLogoutTime holds the string denoting the logout_time field in the database.
	FieldLogoutTime = "logout_time"
	// FieldIsLogout holds the string denoting the is_logout field in the database.
	FieldIsLogout = "is_logout"
	// FieldDeviceInfo holds the string denoting the device_info field in the database.
	FieldDeviceInfo = "device_info"
	// EdgeFormMsg holds the string denoting the form_msg edge name in mutations.
	EdgeFormMsg = "form_msg"
	// EdgeTargetMsg holds the string denoting the target_msg edge name in mutations.
	EdgeTargetMsg = "target_msg"
	// EdgeContactOwner holds the string denoting the contact_owner edge name in mutations.
	EdgeContactOwner = "contact_owner"
	// EdgeContactTarget holds the string denoting the contact_target edge name in mutations.
	EdgeContactTarget = "contact_target"
	// EdgeGroupRelationUser holds the string denoting the group_relation_user edge name in mutations.
	EdgeGroupRelationUser = "group_relation_user"
	// EdgeGroupOwner holds the string denoting the group_owner edge name in mutations.
	EdgeGroupOwner = "group_owner"
	// EdgeGroupMsg holds the string denoting the group_msg edge name in mutations.
	EdgeGroupMsg = "group_msg"
	// EdgeFiles holds the string denoting the files edge name in mutations.
	EdgeFiles = "files"
	// Table holds the table name of the user in the database.
	Table = "t_user"
	// FormMsgTable is the table that holds the form_msg relation/edge.
	FormMsgTable = "t_message"
	// FormMsgInverseTable is the table name for the Message entity.
	// It exists in this package in order to avoid circular dependency with the "message" package.
	FormMsgInverseTable = "t_message"
	// FormMsgColumn is the table column denoting the form_msg relation/edge.
	FormMsgColumn = "form_id"
	// TargetMsgTable is the table that holds the target_msg relation/edge.
	TargetMsgTable = "t_message"
	// TargetMsgInverseTable is the table name for the Message entity.
	// It exists in this package in order to avoid circular dependency with the "message" package.
	TargetMsgInverseTable = "t_message"
	// TargetMsgColumn is the table column denoting the target_msg relation/edge.
	TargetMsgColumn = "target_id"
	// ContactOwnerTable is the table that holds the contact_owner relation/edge.
	ContactOwnerTable = "t_contact"
	// ContactOwnerInverseTable is the table name for the Contact entity.
	// It exists in this package in order to avoid circular dependency with the "contact" package.
	ContactOwnerInverseTable = "t_contact"
	// ContactOwnerColumn is the table column denoting the contact_owner relation/edge.
	ContactOwnerColumn = "owner_id"
	// ContactTargetTable is the table that holds the contact_target relation/edge.
	ContactTargetTable = "t_contact"
	// ContactTargetInverseTable is the table name for the Contact entity.
	// It exists in this package in order to avoid circular dependency with the "contact" package.
	ContactTargetInverseTable = "t_contact"
	// ContactTargetColumn is the table column denoting the contact_target relation/edge.
	ContactTargetColumn = "target_id"
	// GroupRelationUserTable is the table that holds the group_relation_user relation/edge.
	GroupRelationUserTable = "t_group_relation"
	// GroupRelationUserInverseTable is the table name for the GroupRelation entity.
	// It exists in this package in order to avoid circular dependency with the "grouprelation" package.
	GroupRelationUserInverseTable = "t_group_relation"
	// GroupRelationUserColumn is the table column denoting the group_relation_user relation/edge.
	GroupRelationUserColumn = "user_id"
	// GroupOwnerTable is the table that holds the group_owner relation/edge.
	GroupOwnerTable = "t_group"
	// GroupOwnerInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupOwnerInverseTable = "t_group"
	// GroupOwnerColumn is the table column denoting the group_owner relation/edge.
	GroupOwnerColumn = "user_id"
	// GroupMsgTable is the table that holds the group_msg relation/edge.
	GroupMsgTable = "t_group_msg"
	// GroupMsgInverseTable is the table name for the GroupMsg entity.
	// It exists in this package in order to avoid circular dependency with the "groupmsg" package.
	GroupMsgInverseTable = "t_group_msg"
	// GroupMsgColumn is the table column denoting the group_msg relation/edge.
	GroupMsgColumn = "owner_id"
	// FilesTable is the table that holds the files relation/edge.
	FilesTable = "t_file"
	// FilesInverseTable is the table name for the File entity.
	// It exists in this package in order to avoid circular dependency with the "file" package.
	FilesInverseTable = "t_file"
	// FilesColumn is the table column denoting the files relation/edge.
	FilesColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldPassword,
	FieldAvatar,
	FieldPhone,
	FieldEmail,
	FieldClientIP,
	FieldClientPort,
	FieldSalt,
	FieldLoginTime,
	FieldHeartbeatTime,
	FieldLogoutTime,
	FieldIsLogout,
	FieldDeviceInfo,
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
	// DefaultAvatar holds the default value on creation for the "avatar" field.
	DefaultAvatar string
	// DefaultPhone holds the default value on creation for the "phone" field.
	DefaultPhone string
	// DefaultEmail holds the default value on creation for the "email" field.
	DefaultEmail string
	// DefaultClientIP holds the default value on creation for the "client_ip" field.
	DefaultClientIP string
	// DefaultClientPort holds the default value on creation for the "client_port" field.
	DefaultClientPort string
	// DefaultSalt holds the default value on creation for the "salt" field.
	DefaultSalt string
	// DefaultIsLogout holds the default value on creation for the "is_logout" field.
	DefaultIsLogout bool
	// DefaultDeviceInfo holds the default value on creation for the "device_info" field.
	DefaultDeviceInfo string
)

// OrderOption defines the ordering options for the User queries.
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

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByAvatar orders the results by the avatar field.
func ByAvatar(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatar, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByClientIP orders the results by the client_ip field.
func ByClientIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientIP, opts...).ToFunc()
}

// ByClientPort orders the results by the client_port field.
func ByClientPort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientPort, opts...).ToFunc()
}

// BySalt orders the results by the salt field.
func BySalt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSalt, opts...).ToFunc()
}

// ByLoginTime orders the results by the login_time field.
func ByLoginTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLoginTime, opts...).ToFunc()
}

// ByHeartbeatTime orders the results by the heartbeat_time field.
func ByHeartbeatTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHeartbeatTime, opts...).ToFunc()
}

// ByLogoutTime orders the results by the logout_time field.
func ByLogoutTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLogoutTime, opts...).ToFunc()
}

// ByIsLogout orders the results by the is_logout field.
func ByIsLogout(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsLogout, opts...).ToFunc()
}

// ByDeviceInfo orders the results by the device_info field.
func ByDeviceInfo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeviceInfo, opts...).ToFunc()
}

// ByFormMsgCount orders the results by form_msg count.
func ByFormMsgCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFormMsgStep(), opts...)
	}
}

// ByFormMsg orders the results by form_msg terms.
func ByFormMsg(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFormMsgStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTargetMsgCount orders the results by target_msg count.
func ByTargetMsgCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTargetMsgStep(), opts...)
	}
}

// ByTargetMsg orders the results by target_msg terms.
func ByTargetMsg(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTargetMsgStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByContactOwnerCount orders the results by contact_owner count.
func ByContactOwnerCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newContactOwnerStep(), opts...)
	}
}

// ByContactOwner orders the results by contact_owner terms.
func ByContactOwner(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newContactOwnerStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByContactTargetCount orders the results by contact_target count.
func ByContactTargetCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newContactTargetStep(), opts...)
	}
}

// ByContactTarget orders the results by contact_target terms.
func ByContactTarget(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newContactTargetStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByGroupRelationUserCount orders the results by group_relation_user count.
func ByGroupRelationUserCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newGroupRelationUserStep(), opts...)
	}
}

// ByGroupRelationUser orders the results by group_relation_user terms.
func ByGroupRelationUser(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGroupRelationUserStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByGroupOwnerCount orders the results by group_owner count.
func ByGroupOwnerCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newGroupOwnerStep(), opts...)
	}
}

// ByGroupOwner orders the results by group_owner terms.
func ByGroupOwner(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGroupOwnerStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByGroupMsgCount orders the results by group_msg count.
func ByGroupMsgCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newGroupMsgStep(), opts...)
	}
}

// ByGroupMsg orders the results by group_msg terms.
func ByGroupMsg(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGroupMsgStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFilesCount orders the results by files count.
func ByFilesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFilesStep(), opts...)
	}
}

// ByFiles orders the results by files terms.
func ByFiles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFilesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newFormMsgStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FormMsgInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FormMsgTable, FormMsgColumn),
	)
}
func newTargetMsgStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TargetMsgInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TargetMsgTable, TargetMsgColumn),
	)
}
func newContactOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ContactOwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ContactOwnerTable, ContactOwnerColumn),
	)
}
func newContactTargetStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ContactTargetInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ContactTargetTable, ContactTargetColumn),
	)
}
func newGroupRelationUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GroupRelationUserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, GroupRelationUserTable, GroupRelationUserColumn),
	)
}
func newGroupOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GroupOwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, GroupOwnerTable, GroupOwnerColumn),
	)
}
func newGroupMsgStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GroupMsgInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, GroupMsgTable, GroupMsgColumn),
	)
}
func newFilesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FilesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FilesTable, FilesColumn),
	)
}
