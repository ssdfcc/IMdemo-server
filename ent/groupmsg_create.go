// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-zero-chat/ent/group"
	"go-zero-chat/ent/groupmsg"
	"go-zero-chat/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GroupMsgCreate is the builder for creating a GroupMsg entity.
type GroupMsgCreate struct {
	config
	mutation *GroupMsgMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (gmc *GroupMsgCreate) SetCreatedAt(t time.Time) *GroupMsgCreate {
	gmc.mutation.SetCreatedAt(t)
	return gmc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gmc *GroupMsgCreate) SetNillableCreatedAt(t *time.Time) *GroupMsgCreate {
	if t != nil {
		gmc.SetCreatedAt(*t)
	}
	return gmc
}

// SetUpdatedAt sets the "updated_at" field.
func (gmc *GroupMsgCreate) SetUpdatedAt(t time.Time) *GroupMsgCreate {
	gmc.mutation.SetUpdatedAt(t)
	return gmc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gmc *GroupMsgCreate) SetNillableUpdatedAt(t *time.Time) *GroupMsgCreate {
	if t != nil {
		gmc.SetUpdatedAt(*t)
	}
	return gmc
}

// SetDeletedAt sets the "deleted_at" field.
func (gmc *GroupMsgCreate) SetDeletedAt(t time.Time) *GroupMsgCreate {
	gmc.mutation.SetDeletedAt(t)
	return gmc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gmc *GroupMsgCreate) SetNillableDeletedAt(t *time.Time) *GroupMsgCreate {
	if t != nil {
		gmc.SetDeletedAt(*t)
	}
	return gmc
}

// SetOwnerID sets the "owner_id" field.
func (gmc *GroupMsgCreate) SetOwnerID(i int) *GroupMsgCreate {
	gmc.mutation.SetOwnerID(i)
	return gmc
}

// SetTargetID sets the "target_id" field.
func (gmc *GroupMsgCreate) SetTargetID(i int) *GroupMsgCreate {
	gmc.mutation.SetTargetID(i)
	return gmc
}

// SetMedia sets the "media" field.
func (gmc *GroupMsgCreate) SetMedia(i int) *GroupMsgCreate {
	gmc.mutation.SetMedia(i)
	return gmc
}

// SetContent sets the "content" field.
func (gmc *GroupMsgCreate) SetContent(s string) *GroupMsgCreate {
	gmc.mutation.SetContent(s)
	return gmc
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (gmc *GroupMsgCreate) SetNillableContent(s *string) *GroupMsgCreate {
	if s != nil {
		gmc.SetContent(*s)
	}
	return gmc
}

// SetPic sets the "pic" field.
func (gmc *GroupMsgCreate) SetPic(s string) *GroupMsgCreate {
	gmc.mutation.SetPic(s)
	return gmc
}

// SetNillablePic sets the "pic" field if the given value is not nil.
func (gmc *GroupMsgCreate) SetNillablePic(s *string) *GroupMsgCreate {
	if s != nil {
		gmc.SetPic(*s)
	}
	return gmc
}

// SetURL sets the "url" field.
func (gmc *GroupMsgCreate) SetURL(s string) *GroupMsgCreate {
	gmc.mutation.SetURL(s)
	return gmc
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (gmc *GroupMsgCreate) SetNillableURL(s *string) *GroupMsgCreate {
	if s != nil {
		gmc.SetURL(*s)
	}
	return gmc
}

// SetDesc sets the "desc" field.
func (gmc *GroupMsgCreate) SetDesc(s string) *GroupMsgCreate {
	gmc.mutation.SetDesc(s)
	return gmc
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (gmc *GroupMsgCreate) SetNillableDesc(s *string) *GroupMsgCreate {
	if s != nil {
		gmc.SetDesc(*s)
	}
	return gmc
}

// SetAmount sets the "amount" field.
func (gmc *GroupMsgCreate) SetAmount(i int) *GroupMsgCreate {
	gmc.mutation.SetAmount(i)
	return gmc
}

// SetType sets the "type" field.
func (gmc *GroupMsgCreate) SetType(i int) *GroupMsgCreate {
	gmc.mutation.SetType(i)
	return gmc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (gmc *GroupMsgCreate) SetUserID(id int) *GroupMsgCreate {
	gmc.mutation.SetUserID(id)
	return gmc
}

// SetUser sets the "user" edge to the User entity.
func (gmc *GroupMsgCreate) SetUser(u *User) *GroupMsgCreate {
	return gmc.SetUserID(u.ID)
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (gmc *GroupMsgCreate) SetGroupID(id int) *GroupMsgCreate {
	gmc.mutation.SetGroupID(id)
	return gmc
}

// SetGroup sets the "group" edge to the Group entity.
func (gmc *GroupMsgCreate) SetGroup(g *Group) *GroupMsgCreate {
	return gmc.SetGroupID(g.ID)
}

// Mutation returns the GroupMsgMutation object of the builder.
func (gmc *GroupMsgCreate) Mutation() *GroupMsgMutation {
	return gmc.mutation
}

// Save creates the GroupMsg in the database.
func (gmc *GroupMsgCreate) Save(ctx context.Context) (*GroupMsg, error) {
	gmc.defaults()
	return withHooks(ctx, gmc.sqlSave, gmc.mutation, gmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gmc *GroupMsgCreate) SaveX(ctx context.Context) *GroupMsg {
	v, err := gmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gmc *GroupMsgCreate) Exec(ctx context.Context) error {
	_, err := gmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmc *GroupMsgCreate) ExecX(ctx context.Context) {
	if err := gmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gmc *GroupMsgCreate) defaults() {
	if _, ok := gmc.mutation.CreatedAt(); !ok {
		v := groupmsg.DefaultCreatedAt()
		gmc.mutation.SetCreatedAt(v)
	}
	if _, ok := gmc.mutation.UpdatedAt(); !ok {
		v := groupmsg.DefaultUpdatedAt()
		gmc.mutation.SetUpdatedAt(v)
	}
	if _, ok := gmc.mutation.Content(); !ok {
		v := groupmsg.DefaultContent
		gmc.mutation.SetContent(v)
	}
	if _, ok := gmc.mutation.Pic(); !ok {
		v := groupmsg.DefaultPic
		gmc.mutation.SetPic(v)
	}
	if _, ok := gmc.mutation.URL(); !ok {
		v := groupmsg.DefaultURL
		gmc.mutation.SetURL(v)
	}
	if _, ok := gmc.mutation.Desc(); !ok {
		v := groupmsg.DefaultDesc
		gmc.mutation.SetDesc(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gmc *GroupMsgCreate) check() error {
	if _, ok := gmc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "GroupMsg.created_at"`)}
	}
	if _, ok := gmc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "GroupMsg.updated_at"`)}
	}
	if _, ok := gmc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`ent: missing required field "GroupMsg.owner_id"`)}
	}
	if _, ok := gmc.mutation.TargetID(); !ok {
		return &ValidationError{Name: "target_id", err: errors.New(`ent: missing required field "GroupMsg.target_id"`)}
	}
	if _, ok := gmc.mutation.Media(); !ok {
		return &ValidationError{Name: "media", err: errors.New(`ent: missing required field "GroupMsg.media"`)}
	}
	if _, ok := gmc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "GroupMsg.content"`)}
	}
	if _, ok := gmc.mutation.Pic(); !ok {
		return &ValidationError{Name: "pic", err: errors.New(`ent: missing required field "GroupMsg.pic"`)}
	}
	if _, ok := gmc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "GroupMsg.url"`)}
	}
	if _, ok := gmc.mutation.Desc(); !ok {
		return &ValidationError{Name: "desc", err: errors.New(`ent: missing required field "GroupMsg.desc"`)}
	}
	if _, ok := gmc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "GroupMsg.amount"`)}
	}
	if _, ok := gmc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "GroupMsg.type"`)}
	}
	if _, ok := gmc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "GroupMsg.user"`)}
	}
	if _, ok := gmc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group", err: errors.New(`ent: missing required edge "GroupMsg.group"`)}
	}
	return nil
}

func (gmc *GroupMsgCreate) sqlSave(ctx context.Context) (*GroupMsg, error) {
	if err := gmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	gmc.mutation.id = &_node.ID
	gmc.mutation.done = true
	return _node, nil
}

func (gmc *GroupMsgCreate) createSpec() (*GroupMsg, *sqlgraph.CreateSpec) {
	var (
		_node = &GroupMsg{config: gmc.config}
		_spec = sqlgraph.NewCreateSpec(groupmsg.Table, sqlgraph.NewFieldSpec(groupmsg.FieldID, field.TypeInt))
	)
	if value, ok := gmc.mutation.CreatedAt(); ok {
		_spec.SetField(groupmsg.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := gmc.mutation.UpdatedAt(); ok {
		_spec.SetField(groupmsg.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := gmc.mutation.DeletedAt(); ok {
		_spec.SetField(groupmsg.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := gmc.mutation.Media(); ok {
		_spec.SetField(groupmsg.FieldMedia, field.TypeInt, value)
		_node.Media = value
	}
	if value, ok := gmc.mutation.Content(); ok {
		_spec.SetField(groupmsg.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := gmc.mutation.Pic(); ok {
		_spec.SetField(groupmsg.FieldPic, field.TypeString, value)
		_node.Pic = value
	}
	if value, ok := gmc.mutation.URL(); ok {
		_spec.SetField(groupmsg.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := gmc.mutation.Desc(); ok {
		_spec.SetField(groupmsg.FieldDesc, field.TypeString, value)
		_node.Desc = value
	}
	if value, ok := gmc.mutation.Amount(); ok {
		_spec.SetField(groupmsg.FieldAmount, field.TypeInt, value)
		_node.Amount = value
	}
	if value, ok := gmc.mutation.GetType(); ok {
		_spec.SetField(groupmsg.FieldType, field.TypeInt, value)
		_node.Type = value
	}
	if nodes := gmc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmsg.UserTable,
			Columns: []string{groupmsg.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OwnerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gmc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmsg.GroupTable,
			Columns: []string{groupmsg.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
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

// GroupMsgCreateBulk is the builder for creating many GroupMsg entities in bulk.
type GroupMsgCreateBulk struct {
	config
	err      error
	builders []*GroupMsgCreate
}

// Save creates the GroupMsg entities in the database.
func (gmcb *GroupMsgCreateBulk) Save(ctx context.Context) ([]*GroupMsg, error) {
	if gmcb.err != nil {
		return nil, gmcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(gmcb.builders))
	nodes := make([]*GroupMsg, len(gmcb.builders))
	mutators := make([]Mutator, len(gmcb.builders))
	for i := range gmcb.builders {
		func(i int, root context.Context) {
			builder := gmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupMsgMutation)
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
					_, err = mutators[i+1].Mutate(root, gmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gmcb *GroupMsgCreateBulk) SaveX(ctx context.Context) []*GroupMsg {
	v, err := gmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gmcb *GroupMsgCreateBulk) Exec(ctx context.Context) error {
	_, err := gmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmcb *GroupMsgCreateBulk) ExecX(ctx context.Context) {
	if err := gmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
