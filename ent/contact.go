// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-zero-chat/ent/contact"
	"go-zero-chat/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Contact is the model entity for the Contact schema.
type Contact struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 创建时间
	CreatedAt time.Time `json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// 谁的关系
	OwnerID int `json:"owner_id,omitempty"`
	// 对应谁
	TargetID int `json:"target_id,omitempty"`
	// 对应的类型 1好友
	Type int `json:"type,omitempty"`
	// 描述相关
	Desc string `json:"desc,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ContactQuery when eager-loading is set.
	Edges        ContactEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ContactEdges holds the relations/edges for other nodes in the graph.
type ContactEdges struct {
	// OwnerUser holds the value of the owner_user edge.
	OwnerUser *User `json:"owner_user,omitempty"`
	// TargetUser holds the value of the target_user edge.
	TargetUser *User `json:"target_user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OwnerUserOrErr returns the OwnerUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ContactEdges) OwnerUserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.OwnerUser == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.OwnerUser, nil
	}
	return nil, &NotLoadedError{edge: "owner_user"}
}

// TargetUserOrErr returns the TargetUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ContactEdges) TargetUserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.TargetUser == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.TargetUser, nil
	}
	return nil, &NotLoadedError{edge: "target_user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Contact) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case contact.FieldID, contact.FieldOwnerID, contact.FieldTargetID, contact.FieldType:
			values[i] = new(sql.NullInt64)
		case contact.FieldDesc:
			values[i] = new(sql.NullString)
		case contact.FieldCreatedAt, contact.FieldUpdatedAt, contact.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Contact fields.
func (c *Contact) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case contact.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case contact.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case contact.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case contact.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				c.DeletedAt = new(time.Time)
				*c.DeletedAt = value.Time
			}
		case contact.FieldOwnerID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				c.OwnerID = int(value.Int64)
			}
		case contact.FieldTargetID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field target_id", values[i])
			} else if value.Valid {
				c.TargetID = int(value.Int64)
			}
		case contact.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				c.Type = int(value.Int64)
			}
		case contact.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				c.Desc = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Contact.
// This includes values selected through modifiers, order, etc.
func (c *Contact) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryOwnerUser queries the "owner_user" edge of the Contact entity.
func (c *Contact) QueryOwnerUser() *UserQuery {
	return NewContactClient(c.config).QueryOwnerUser(c)
}

// QueryTargetUser queries the "target_user" edge of the Contact entity.
func (c *Contact) QueryTargetUser() *UserQuery {
	return NewContactClient(c.config).QueryTargetUser(c)
}

// Update returns a builder for updating this Contact.
// Note that you need to call Contact.Unwrap() before calling this method if this Contact
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Contact) Update() *ContactUpdateOne {
	return NewContactClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Contact entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Contact) Unwrap() *Contact {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Contact is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Contact) String() string {
	var builder strings.Builder
	builder.WriteString("Contact(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := c.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(fmt.Sprintf("%v", c.OwnerID))
	builder.WriteString(", ")
	builder.WriteString("target_id=")
	builder.WriteString(fmt.Sprintf("%v", c.TargetID))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", c.Type))
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(c.Desc)
	builder.WriteByte(')')
	return builder.String()
}

// Contacts is a parsable slice of Contact.
type Contacts []*Contact
