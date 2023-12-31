// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-zero-chat/ent/group"
	"go-zero-chat/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Group is the model entity for the Group schema.
type Group struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 创建时间
	CreatedAt time.Time `json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// 群名
	Name string `json:"name,omitempty"`
	// 所有者
	UserID int `json:"user_id,omitempty"`
	// 图片
	Img string `json:"img,omitempty"`
	// 描述相关
	Desc string `json:"desc,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupQuery when eager-loading is set.
	Edges        GroupEdges `json:"edges"`
	selectValues sql.SelectValues
}

// GroupEdges holds the relations/edges for other nodes in the graph.
type GroupEdges struct {
	// OwnerUser holds the value of the owner_user edge.
	OwnerUser *User `json:"owner_user,omitempty"`
	// GroupRelationGroup holds the value of the group_relation_group edge.
	GroupRelationGroup []*GroupRelation `json:"group_relation_group,omitempty"`
	// TargetMsg holds the value of the target_msg edge.
	TargetMsg []*GroupMsg `json:"target_msg,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// OwnerUserOrErr returns the OwnerUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupEdges) OwnerUserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.OwnerUser == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.OwnerUser, nil
	}
	return nil, &NotLoadedError{edge: "owner_user"}
}

// GroupRelationGroupOrErr returns the GroupRelationGroup value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) GroupRelationGroupOrErr() ([]*GroupRelation, error) {
	if e.loadedTypes[1] {
		return e.GroupRelationGroup, nil
	}
	return nil, &NotLoadedError{edge: "group_relation_group"}
}

// TargetMsgOrErr returns the TargetMsg value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) TargetMsgOrErr() ([]*GroupMsg, error) {
	if e.loadedTypes[2] {
		return e.TargetMsg, nil
	}
	return nil, &NotLoadedError{edge: "target_msg"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Group) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case group.FieldID, group.FieldUserID:
			values[i] = new(sql.NullInt64)
		case group.FieldName, group.FieldImg, group.FieldDesc:
			values[i] = new(sql.NullString)
		case group.FieldCreatedAt, group.FieldUpdatedAt, group.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Group fields.
func (gr *Group) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case group.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gr.ID = int(value.Int64)
		case group.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gr.CreatedAt = value.Time
			}
		case group.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gr.UpdatedAt = value.Time
			}
		case group.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				gr.DeletedAt = new(time.Time)
				*gr.DeletedAt = value.Time
			}
		case group.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				gr.Name = value.String
			}
		case group.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				gr.UserID = int(value.Int64)
			}
		case group.FieldImg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field img", values[i])
			} else if value.Valid {
				gr.Img = value.String
			}
		case group.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				gr.Desc = value.String
			}
		default:
			gr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Group.
// This includes values selected through modifiers, order, etc.
func (gr *Group) Value(name string) (ent.Value, error) {
	return gr.selectValues.Get(name)
}

// QueryOwnerUser queries the "owner_user" edge of the Group entity.
func (gr *Group) QueryOwnerUser() *UserQuery {
	return NewGroupClient(gr.config).QueryOwnerUser(gr)
}

// QueryGroupRelationGroup queries the "group_relation_group" edge of the Group entity.
func (gr *Group) QueryGroupRelationGroup() *GroupRelationQuery {
	return NewGroupClient(gr.config).QueryGroupRelationGroup(gr)
}

// QueryTargetMsg queries the "target_msg" edge of the Group entity.
func (gr *Group) QueryTargetMsg() *GroupMsgQuery {
	return NewGroupClient(gr.config).QueryTargetMsg(gr)
}

// Update returns a builder for updating this Group.
// Note that you need to call Group.Unwrap() before calling this method if this Group
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *Group) Update() *GroupUpdateOne {
	return NewGroupClient(gr.config).UpdateOne(gr)
}

// Unwrap unwraps the Group entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gr *Group) Unwrap() *Group {
	_tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Group is not a transactional entity")
	}
	gr.config.driver = _tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *Group) String() string {
	var builder strings.Builder
	builder.WriteString("Group(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(gr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(gr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := gr.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(gr.Name)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", gr.UserID))
	builder.WriteString(", ")
	builder.WriteString("img=")
	builder.WriteString(gr.Img)
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(gr.Desc)
	builder.WriteByte(')')
	return builder.String()
}

// Groups is a parsable slice of Group.
type Groups []*Group
