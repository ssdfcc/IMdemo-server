// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-zero-chat/ent/message"
	"go-zero-chat/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageCreate is the builder for creating a Message entity.
type MessageCreate struct {
	config
	mutation *MessageMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mc *MessageCreate) SetCreatedAt(t time.Time) *MessageCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MessageCreate) SetNillableCreatedAt(t *time.Time) *MessageCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MessageCreate) SetUpdatedAt(t time.Time) *MessageCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MessageCreate) SetNillableUpdatedAt(t *time.Time) *MessageCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetDeletedAt sets the "deleted_at" field.
func (mc *MessageCreate) SetDeletedAt(t time.Time) *MessageCreate {
	mc.mutation.SetDeletedAt(t)
	return mc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mc *MessageCreate) SetNillableDeletedAt(t *time.Time) *MessageCreate {
	if t != nil {
		mc.SetDeletedAt(*t)
	}
	return mc
}

// SetType sets the "type" field.
func (mc *MessageCreate) SetType(i int) *MessageCreate {
	mc.mutation.SetType(i)
	return mc
}

// SetMedia sets the "media" field.
func (mc *MessageCreate) SetMedia(i int) *MessageCreate {
	mc.mutation.SetMedia(i)
	return mc
}

// SetContent sets the "content" field.
func (mc *MessageCreate) SetContent(s string) *MessageCreate {
	mc.mutation.SetContent(s)
	return mc
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (mc *MessageCreate) SetNillableContent(s *string) *MessageCreate {
	if s != nil {
		mc.SetContent(*s)
	}
	return mc
}

// SetPic sets the "pic" field.
func (mc *MessageCreate) SetPic(s string) *MessageCreate {
	mc.mutation.SetPic(s)
	return mc
}

// SetNillablePic sets the "pic" field if the given value is not nil.
func (mc *MessageCreate) SetNillablePic(s *string) *MessageCreate {
	if s != nil {
		mc.SetPic(*s)
	}
	return mc
}

// SetURL sets the "url" field.
func (mc *MessageCreate) SetURL(s string) *MessageCreate {
	mc.mutation.SetURL(s)
	return mc
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (mc *MessageCreate) SetNillableURL(s *string) *MessageCreate {
	if s != nil {
		mc.SetURL(*s)
	}
	return mc
}

// SetDesc sets the "desc" field.
func (mc *MessageCreate) SetDesc(s string) *MessageCreate {
	mc.mutation.SetDesc(s)
	return mc
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (mc *MessageCreate) SetNillableDesc(s *string) *MessageCreate {
	if s != nil {
		mc.SetDesc(*s)
	}
	return mc
}

// SetAmount sets the "amount" field.
func (mc *MessageCreate) SetAmount(i int) *MessageCreate {
	mc.mutation.SetAmount(i)
	return mc
}

// SetFormID sets the "form_id" field.
func (mc *MessageCreate) SetFormID(i int) *MessageCreate {
	mc.mutation.SetFormID(i)
	return mc
}

// SetTargetID sets the "target_id" field.
func (mc *MessageCreate) SetTargetID(i int) *MessageCreate {
	mc.mutation.SetTargetID(i)
	return mc
}

// SetSenderID sets the "sender" edge to the User entity by ID.
func (mc *MessageCreate) SetSenderID(id int) *MessageCreate {
	mc.mutation.SetSenderID(id)
	return mc
}

// SetSender sets the "sender" edge to the User entity.
func (mc *MessageCreate) SetSender(u *User) *MessageCreate {
	return mc.SetSenderID(u.ID)
}

// SetRecipientID sets the "recipient" edge to the User entity by ID.
func (mc *MessageCreate) SetRecipientID(id int) *MessageCreate {
	mc.mutation.SetRecipientID(id)
	return mc
}

// SetRecipient sets the "recipient" edge to the User entity.
func (mc *MessageCreate) SetRecipient(u *User) *MessageCreate {
	return mc.SetRecipientID(u.ID)
}

// Mutation returns the MessageMutation object of the builder.
func (mc *MessageCreate) Mutation() *MessageMutation {
	return mc.mutation
}

// Save creates the Message in the database.
func (mc *MessageCreate) Save(ctx context.Context) (*Message, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MessageCreate) SaveX(ctx context.Context) *Message {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MessageCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MessageCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MessageCreate) defaults() {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := message.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := message.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mc.mutation.Content(); !ok {
		v := message.DefaultContent
		mc.mutation.SetContent(v)
	}
	if _, ok := mc.mutation.Pic(); !ok {
		v := message.DefaultPic
		mc.mutation.SetPic(v)
	}
	if _, ok := mc.mutation.URL(); !ok {
		v := message.DefaultURL
		mc.mutation.SetURL(v)
	}
	if _, ok := mc.mutation.Desc(); !ok {
		v := message.DefaultDesc
		mc.mutation.SetDesc(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MessageCreate) check() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Message.created_at"`)}
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Message.updated_at"`)}
	}
	if _, ok := mc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Message.type"`)}
	}
	if _, ok := mc.mutation.Media(); !ok {
		return &ValidationError{Name: "media", err: errors.New(`ent: missing required field "Message.media"`)}
	}
	if _, ok := mc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Message.content"`)}
	}
	if _, ok := mc.mutation.Pic(); !ok {
		return &ValidationError{Name: "pic", err: errors.New(`ent: missing required field "Message.pic"`)}
	}
	if _, ok := mc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Message.url"`)}
	}
	if _, ok := mc.mutation.Desc(); !ok {
		return &ValidationError{Name: "desc", err: errors.New(`ent: missing required field "Message.desc"`)}
	}
	if _, ok := mc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "Message.amount"`)}
	}
	if _, ok := mc.mutation.FormID(); !ok {
		return &ValidationError{Name: "form_id", err: errors.New(`ent: missing required field "Message.form_id"`)}
	}
	if _, ok := mc.mutation.TargetID(); !ok {
		return &ValidationError{Name: "target_id", err: errors.New(`ent: missing required field "Message.target_id"`)}
	}
	if _, ok := mc.mutation.SenderID(); !ok {
		return &ValidationError{Name: "sender", err: errors.New(`ent: missing required edge "Message.sender"`)}
	}
	if _, ok := mc.mutation.RecipientID(); !ok {
		return &ValidationError{Name: "recipient", err: errors.New(`ent: missing required edge "Message.recipient"`)}
	}
	return nil
}

func (mc *MessageCreate) sqlSave(ctx context.Context) (*Message, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MessageCreate) createSpec() (*Message, *sqlgraph.CreateSpec) {
	var (
		_node = &Message{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(message.Table, sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt))
	)
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(message.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.SetField(message.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.DeletedAt(); ok {
		_spec.SetField(message.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := mc.mutation.GetType(); ok {
		_spec.SetField(message.FieldType, field.TypeInt, value)
		_node.Type = value
	}
	if value, ok := mc.mutation.Media(); ok {
		_spec.SetField(message.FieldMedia, field.TypeInt, value)
		_node.Media = value
	}
	if value, ok := mc.mutation.Content(); ok {
		_spec.SetField(message.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := mc.mutation.Pic(); ok {
		_spec.SetField(message.FieldPic, field.TypeString, value)
		_node.Pic = value
	}
	if value, ok := mc.mutation.URL(); ok {
		_spec.SetField(message.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := mc.mutation.Desc(); ok {
		_spec.SetField(message.FieldDesc, field.TypeString, value)
		_node.Desc = value
	}
	if value, ok := mc.mutation.Amount(); ok {
		_spec.SetField(message.FieldAmount, field.TypeInt, value)
		_node.Amount = value
	}
	if nodes := mc.mutation.SenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.SenderTable,
			Columns: []string{message.SenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.FormID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.RecipientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.RecipientTable,
			Columns: []string{message.RecipientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TargetID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MessageCreateBulk is the builder for creating many Message entities in bulk.
type MessageCreateBulk struct {
	config
	err      error
	builders []*MessageCreate
}

// Save creates the Message entities in the database.
func (mcb *MessageCreateBulk) Save(ctx context.Context) ([]*Message, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Message, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MessageCreateBulk) SaveX(ctx context.Context) []*Message {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MessageCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MessageCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
