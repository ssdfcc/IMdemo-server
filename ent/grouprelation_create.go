// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-zero-chat/ent/group"
	"go-zero-chat/ent/grouprelation"
	"go-zero-chat/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GroupRelationCreate is the builder for creating a GroupRelation entity.
type GroupRelationCreate struct {
	config
	mutation *GroupRelationMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (grc *GroupRelationCreate) SetCreatedAt(t time.Time) *GroupRelationCreate {
	grc.mutation.SetCreatedAt(t)
	return grc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (grc *GroupRelationCreate) SetNillableCreatedAt(t *time.Time) *GroupRelationCreate {
	if t != nil {
		grc.SetCreatedAt(*t)
	}
	return grc
}

// SetUpdatedAt sets the "updated_at" field.
func (grc *GroupRelationCreate) SetUpdatedAt(t time.Time) *GroupRelationCreate {
	grc.mutation.SetUpdatedAt(t)
	return grc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (grc *GroupRelationCreate) SetNillableUpdatedAt(t *time.Time) *GroupRelationCreate {
	if t != nil {
		grc.SetUpdatedAt(*t)
	}
	return grc
}

// SetDeletedAt sets the "deleted_at" field.
func (grc *GroupRelationCreate) SetDeletedAt(t time.Time) *GroupRelationCreate {
	grc.mutation.SetDeletedAt(t)
	return grc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (grc *GroupRelationCreate) SetNillableDeletedAt(t *time.Time) *GroupRelationCreate {
	if t != nil {
		grc.SetDeletedAt(*t)
	}
	return grc
}

// SetGroupID sets the "group_id" field.
func (grc *GroupRelationCreate) SetGroupID(i int) *GroupRelationCreate {
	grc.mutation.SetGroupID(i)
	return grc
}

// SetUserID sets the "user_id" field.
func (grc *GroupRelationCreate) SetUserID(i int) *GroupRelationCreate {
	grc.mutation.SetUserID(i)
	return grc
}

// SetType sets the "type" field.
func (grc *GroupRelationCreate) SetType(i int) *GroupRelationCreate {
	grc.mutation.SetType(i)
	return grc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (grc *GroupRelationCreate) SetNillableType(i *int) *GroupRelationCreate {
	if i != nil {
		grc.SetType(*i)
	}
	return grc
}

// SetDesc sets the "desc" field.
func (grc *GroupRelationCreate) SetDesc(s string) *GroupRelationCreate {
	grc.mutation.SetDesc(s)
	return grc
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (grc *GroupRelationCreate) SetNillableDesc(s *string) *GroupRelationCreate {
	if s != nil {
		grc.SetDesc(*s)
	}
	return grc
}

// SetGroup sets the "group" edge to the Group entity.
func (grc *GroupRelationCreate) SetGroup(g *Group) *GroupRelationCreate {
	return grc.SetGroupID(g.ID)
}

// SetUser sets the "user" edge to the User entity.
func (grc *GroupRelationCreate) SetUser(u *User) *GroupRelationCreate {
	return grc.SetUserID(u.ID)
}

// Mutation returns the GroupRelationMutation object of the builder.
func (grc *GroupRelationCreate) Mutation() *GroupRelationMutation {
	return grc.mutation
}

// Save creates the GroupRelation in the database.
func (grc *GroupRelationCreate) Save(ctx context.Context) (*GroupRelation, error) {
	grc.defaults()
	return withHooks(ctx, grc.sqlSave, grc.mutation, grc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (grc *GroupRelationCreate) SaveX(ctx context.Context) *GroupRelation {
	v, err := grc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (grc *GroupRelationCreate) Exec(ctx context.Context) error {
	_, err := grc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (grc *GroupRelationCreate) ExecX(ctx context.Context) {
	if err := grc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (grc *GroupRelationCreate) defaults() {
	if _, ok := grc.mutation.CreatedAt(); !ok {
		v := grouprelation.DefaultCreatedAt()
		grc.mutation.SetCreatedAt(v)
	}
	if _, ok := grc.mutation.UpdatedAt(); !ok {
		v := grouprelation.DefaultUpdatedAt()
		grc.mutation.SetUpdatedAt(v)
	}
	if _, ok := grc.mutation.GetType(); !ok {
		v := grouprelation.DefaultType
		grc.mutation.SetType(v)
	}
	if _, ok := grc.mutation.Desc(); !ok {
		v := grouprelation.DefaultDesc
		grc.mutation.SetDesc(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (grc *GroupRelationCreate) check() error {
	if _, ok := grc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "GroupRelation.created_at"`)}
	}
	if _, ok := grc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "GroupRelation.updated_at"`)}
	}
	if _, ok := grc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group_id", err: errors.New(`ent: missing required field "GroupRelation.group_id"`)}
	}
	if _, ok := grc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "GroupRelation.user_id"`)}
	}
	if _, ok := grc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "GroupRelation.type"`)}
	}
	if _, ok := grc.mutation.Desc(); !ok {
		return &ValidationError{Name: "desc", err: errors.New(`ent: missing required field "GroupRelation.desc"`)}
	}
	if _, ok := grc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group", err: errors.New(`ent: missing required edge "GroupRelation.group"`)}
	}
	if _, ok := grc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "GroupRelation.user"`)}
	}
	return nil
}

func (grc *GroupRelationCreate) sqlSave(ctx context.Context) (*GroupRelation, error) {
	if err := grc.check(); err != nil {
		return nil, err
	}
	_node, _spec := grc.createSpec()
	if err := sqlgraph.CreateNode(ctx, grc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	grc.mutation.id = &_node.ID
	grc.mutation.done = true
	return _node, nil
}

func (grc *GroupRelationCreate) createSpec() (*GroupRelation, *sqlgraph.CreateSpec) {
	var (
		_node = &GroupRelation{config: grc.config}
		_spec = sqlgraph.NewCreateSpec(grouprelation.Table, sqlgraph.NewFieldSpec(grouprelation.FieldID, field.TypeInt))
	)
	if value, ok := grc.mutation.CreatedAt(); ok {
		_spec.SetField(grouprelation.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := grc.mutation.UpdatedAt(); ok {
		_spec.SetField(grouprelation.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := grc.mutation.DeletedAt(); ok {
		_spec.SetField(grouprelation.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := grc.mutation.GetType(); ok {
		_spec.SetField(grouprelation.FieldType, field.TypeInt, value)
		_node.Type = value
	}
	if value, ok := grc.mutation.Desc(); ok {
		_spec.SetField(grouprelation.FieldDesc, field.TypeString, value)
		_node.Desc = value
	}
	if nodes := grc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   grouprelation.GroupTable,
			Columns: []string{grouprelation.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.GroupID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := grc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   grouprelation.UserTable,
			Columns: []string{grouprelation.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GroupRelationCreateBulk is the builder for creating many GroupRelation entities in bulk.
type GroupRelationCreateBulk struct {
	config
	err      error
	builders []*GroupRelationCreate
}

// Save creates the GroupRelation entities in the database.
func (grcb *GroupRelationCreateBulk) Save(ctx context.Context) ([]*GroupRelation, error) {
	if grcb.err != nil {
		return nil, grcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(grcb.builders))
	nodes := make([]*GroupRelation, len(grcb.builders))
	mutators := make([]Mutator, len(grcb.builders))
	for i := range grcb.builders {
		func(i int, root context.Context) {
			builder := grcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupRelationMutation)
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
					_, err = mutators[i+1].Mutate(root, grcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, grcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, grcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (grcb *GroupRelationCreateBulk) SaveX(ctx context.Context) []*GroupRelation {
	v, err := grcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (grcb *GroupRelationCreateBulk) Exec(ctx context.Context) error {
	_, err := grcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (grcb *GroupRelationCreateBulk) ExecX(ctx context.Context) {
	if err := grcb.Exec(ctx); err != nil {
		panic(err)
	}
}
