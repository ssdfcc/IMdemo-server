// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-zero-chat/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 创建时间
	CreatedAt time.Time `json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// 用户名
	Name string `json:"name,omitempty"`
	// 用户密码
	Password string `json:"password,omitempty"`
	// 头像
	Avatar string `json:"avatar,omitempty"`
	// 电话号码
	Phone string `json:"phone,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty"`
	// 客户端ip
	ClientIP string `json:"client_ip,omitempty"`
	// 客户端ip
	ClientPort string `json:"client_port,omitempty"`
	// 加密盐
	Salt string `json:"salt,omitempty"`
	// 登录时间
	LoginTime *time.Time `json:"login_time,omitempty"`
	// 心跳时间
	HeartbeatTime *time.Time `json:"heartbeat_time,omitempty"`
	// 下线时间
	LogoutTime *time.Time `json:"logout_time,omitempty"`
	// 是否下线
	IsLogout bool `json:"is_logout,omitempty"`
	// 设备信息
	DeviceInfo string `json:"device_info,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// FormMsg holds the value of the form_msg edge.
	FormMsg []*Message `json:"form_msg,omitempty"`
	// TargetMsg holds the value of the target_msg edge.
	TargetMsg []*Message `json:"target_msg,omitempty"`
	// ContactOwner holds the value of the contact_owner edge.
	ContactOwner []*Contact `json:"contact_owner,omitempty"`
	// ContactTarget holds the value of the contact_target edge.
	ContactTarget []*Contact `json:"contact_target,omitempty"`
	// GroupRelationUser holds the value of the group_relation_user edge.
	GroupRelationUser []*GroupRelation `json:"group_relation_user,omitempty"`
	// GroupOwner holds the value of the group_owner edge.
	GroupOwner []*Group `json:"group_owner,omitempty"`
	// GroupMsg holds the value of the group_msg edge.
	GroupMsg []*GroupMsg `json:"group_msg,omitempty"`
	// Files holds the value of the files edge.
	Files []*File `json:"files,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
}

// FormMsgOrErr returns the FormMsg value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FormMsgOrErr() ([]*Message, error) {
	if e.loadedTypes[0] {
		return e.FormMsg, nil
	}
	return nil, &NotLoadedError{edge: "form_msg"}
}

// TargetMsgOrErr returns the TargetMsg value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TargetMsgOrErr() ([]*Message, error) {
	if e.loadedTypes[1] {
		return e.TargetMsg, nil
	}
	return nil, &NotLoadedError{edge: "target_msg"}
}

// ContactOwnerOrErr returns the ContactOwner value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ContactOwnerOrErr() ([]*Contact, error) {
	if e.loadedTypes[2] {
		return e.ContactOwner, nil
	}
	return nil, &NotLoadedError{edge: "contact_owner"}
}

// ContactTargetOrErr returns the ContactTarget value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ContactTargetOrErr() ([]*Contact, error) {
	if e.loadedTypes[3] {
		return e.ContactTarget, nil
	}
	return nil, &NotLoadedError{edge: "contact_target"}
}

// GroupRelationUserOrErr returns the GroupRelationUser value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) GroupRelationUserOrErr() ([]*GroupRelation, error) {
	if e.loadedTypes[4] {
		return e.GroupRelationUser, nil
	}
	return nil, &NotLoadedError{edge: "group_relation_user"}
}

// GroupOwnerOrErr returns the GroupOwner value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) GroupOwnerOrErr() ([]*Group, error) {
	if e.loadedTypes[5] {
		return e.GroupOwner, nil
	}
	return nil, &NotLoadedError{edge: "group_owner"}
}

// GroupMsgOrErr returns the GroupMsg value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) GroupMsgOrErr() ([]*GroupMsg, error) {
	if e.loadedTypes[6] {
		return e.GroupMsg, nil
	}
	return nil, &NotLoadedError{edge: "group_msg"}
}

// FilesOrErr returns the Files value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FilesOrErr() ([]*File, error) {
	if e.loadedTypes[7] {
		return e.Files, nil
	}
	return nil, &NotLoadedError{edge: "files"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldIsLogout:
			values[i] = new(sql.NullBool)
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldName, user.FieldPassword, user.FieldAvatar, user.FieldPhone, user.FieldEmail, user.FieldClientIP, user.FieldClientPort, user.FieldSalt, user.FieldDeviceInfo:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldDeletedAt, user.FieldLoginTime, user.FieldHeartbeatTime, user.FieldLogoutTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				u.DeletedAt = new(time.Time)
				*u.DeletedAt = value.Time
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				u.Avatar = value.String
			}
		case user.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				u.Phone = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldClientIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_ip", values[i])
			} else if value.Valid {
				u.ClientIP = value.String
			}
		case user.FieldClientPort:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_port", values[i])
			} else if value.Valid {
				u.ClientPort = value.String
			}
		case user.FieldSalt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field salt", values[i])
			} else if value.Valid {
				u.Salt = value.String
			}
		case user.FieldLoginTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field login_time", values[i])
			} else if value.Valid {
				u.LoginTime = new(time.Time)
				*u.LoginTime = value.Time
			}
		case user.FieldHeartbeatTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field heartbeat_time", values[i])
			} else if value.Valid {
				u.HeartbeatTime = new(time.Time)
				*u.HeartbeatTime = value.Time
			}
		case user.FieldLogoutTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field logout_time", values[i])
			} else if value.Valid {
				u.LogoutTime = new(time.Time)
				*u.LogoutTime = value.Time
			}
		case user.FieldIsLogout:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_logout", values[i])
			} else if value.Valid {
				u.IsLogout = value.Bool
			}
		case user.FieldDeviceInfo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device_info", values[i])
			} else if value.Valid {
				u.DeviceInfo = value.String
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryFormMsg queries the "form_msg" edge of the User entity.
func (u *User) QueryFormMsg() *MessageQuery {
	return NewUserClient(u.config).QueryFormMsg(u)
}

// QueryTargetMsg queries the "target_msg" edge of the User entity.
func (u *User) QueryTargetMsg() *MessageQuery {
	return NewUserClient(u.config).QueryTargetMsg(u)
}

// QueryContactOwner queries the "contact_owner" edge of the User entity.
func (u *User) QueryContactOwner() *ContactQuery {
	return NewUserClient(u.config).QueryContactOwner(u)
}

// QueryContactTarget queries the "contact_target" edge of the User entity.
func (u *User) QueryContactTarget() *ContactQuery {
	return NewUserClient(u.config).QueryContactTarget(u)
}

// QueryGroupRelationUser queries the "group_relation_user" edge of the User entity.
func (u *User) QueryGroupRelationUser() *GroupRelationQuery {
	return NewUserClient(u.config).QueryGroupRelationUser(u)
}

// QueryGroupOwner queries the "group_owner" edge of the User entity.
func (u *User) QueryGroupOwner() *GroupQuery {
	return NewUserClient(u.config).QueryGroupOwner(u)
}

// QueryGroupMsg queries the "group_msg" edge of the User entity.
func (u *User) QueryGroupMsg() *GroupMsgQuery {
	return NewUserClient(u.config).QueryGroupMsg(u)
}

// QueryFiles queries the "files" edge of the User entity.
func (u *User) QueryFiles() *FileQuery {
	return NewUserClient(u.config).QueryFiles(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := u.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(u.Password)
	builder.WriteString(", ")
	builder.WriteString("avatar=")
	builder.WriteString(u.Avatar)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(u.Phone)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("client_ip=")
	builder.WriteString(u.ClientIP)
	builder.WriteString(", ")
	builder.WriteString("client_port=")
	builder.WriteString(u.ClientPort)
	builder.WriteString(", ")
	builder.WriteString("salt=")
	builder.WriteString(u.Salt)
	builder.WriteString(", ")
	if v := u.LoginTime; v != nil {
		builder.WriteString("login_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := u.HeartbeatTime; v != nil {
		builder.WriteString("heartbeat_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := u.LogoutTime; v != nil {
		builder.WriteString("logout_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("is_logout=")
	builder.WriteString(fmt.Sprintf("%v", u.IsLogout))
	builder.WriteString(", ")
	builder.WriteString("device_info=")
	builder.WriteString(u.DeviceInfo)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User