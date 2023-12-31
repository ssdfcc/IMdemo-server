// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-zero-chat/ent/group"
	"go-zero-chat/ent/groupmsg"
	"go-zero-chat/ent/grouprelation"
	"go-zero-chat/ent/predicate"
	"go-zero-chat/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GroupUpdate is the builder for updating Group entities.
type GroupUpdate struct {
	config
	hooks    []Hook
	mutation *GroupMutation
}

// Where appends a list predicates to the GroupUpdate builder.
func (gu *GroupUpdate) Where(ps ...predicate.Group) *GroupUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetUpdatedAt sets the "updated_at" field.
func (gu *GroupUpdate) SetUpdatedAt(t time.Time) *GroupUpdate {
	gu.mutation.SetUpdatedAt(t)
	return gu
}

// SetDeletedAt sets the "deleted_at" field.
func (gu *GroupUpdate) SetDeletedAt(t time.Time) *GroupUpdate {
	gu.mutation.SetDeletedAt(t)
	return gu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableDeletedAt(t *time.Time) *GroupUpdate {
	if t != nil {
		gu.SetDeletedAt(*t)
	}
	return gu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (gu *GroupUpdate) ClearDeletedAt() *GroupUpdate {
	gu.mutation.ClearDeletedAt()
	return gu
}

// SetName sets the "name" field.
func (gu *GroupUpdate) SetName(s string) *GroupUpdate {
	gu.mutation.SetName(s)
	return gu
}

// SetUserID sets the "user_id" field.
func (gu *GroupUpdate) SetUserID(i int) *GroupUpdate {
	gu.mutation.SetUserID(i)
	return gu
}

// SetImg sets the "img" field.
func (gu *GroupUpdate) SetImg(s string) *GroupUpdate {
	gu.mutation.SetImg(s)
	return gu
}

// SetNillableImg sets the "img" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableImg(s *string) *GroupUpdate {
	if s != nil {
		gu.SetImg(*s)
	}
	return gu
}

// SetDesc sets the "desc" field.
func (gu *GroupUpdate) SetDesc(s string) *GroupUpdate {
	gu.mutation.SetDesc(s)
	return gu
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableDesc(s *string) *GroupUpdate {
	if s != nil {
		gu.SetDesc(*s)
	}
	return gu
}

// SetOwnerUserID sets the "owner_user" edge to the User entity by ID.
func (gu *GroupUpdate) SetOwnerUserID(id int) *GroupUpdate {
	gu.mutation.SetOwnerUserID(id)
	return gu
}

// SetOwnerUser sets the "owner_user" edge to the User entity.
func (gu *GroupUpdate) SetOwnerUser(u *User) *GroupUpdate {
	return gu.SetOwnerUserID(u.ID)
}

// AddGroupRelationGroupIDs adds the "group_relation_group" edge to the GroupRelation entity by IDs.
func (gu *GroupUpdate) AddGroupRelationGroupIDs(ids ...int) *GroupUpdate {
	gu.mutation.AddGroupRelationGroupIDs(ids...)
	return gu
}

// AddGroupRelationGroup adds the "group_relation_group" edges to the GroupRelation entity.
func (gu *GroupUpdate) AddGroupRelationGroup(g ...*GroupRelation) *GroupUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gu.AddGroupRelationGroupIDs(ids...)
}

// AddTargetMsgIDs adds the "target_msg" edge to the GroupMsg entity by IDs.
func (gu *GroupUpdate) AddTargetMsgIDs(ids ...int) *GroupUpdate {
	gu.mutation.AddTargetMsgIDs(ids...)
	return gu
}

// AddTargetMsg adds the "target_msg" edges to the GroupMsg entity.
func (gu *GroupUpdate) AddTargetMsg(g ...*GroupMsg) *GroupUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gu.AddTargetMsgIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (gu *GroupUpdate) Mutation() *GroupMutation {
	return gu.mutation
}

// ClearOwnerUser clears the "owner_user" edge to the User entity.
func (gu *GroupUpdate) ClearOwnerUser() *GroupUpdate {
	gu.mutation.ClearOwnerUser()
	return gu
}

// ClearGroupRelationGroup clears all "group_relation_group" edges to the GroupRelation entity.
func (gu *GroupUpdate) ClearGroupRelationGroup() *GroupUpdate {
	gu.mutation.ClearGroupRelationGroup()
	return gu
}

// RemoveGroupRelationGroupIDs removes the "group_relation_group" edge to GroupRelation entities by IDs.
func (gu *GroupUpdate) RemoveGroupRelationGroupIDs(ids ...int) *GroupUpdate {
	gu.mutation.RemoveGroupRelationGroupIDs(ids...)
	return gu
}

// RemoveGroupRelationGroup removes "group_relation_group" edges to GroupRelation entities.
func (gu *GroupUpdate) RemoveGroupRelationGroup(g ...*GroupRelation) *GroupUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gu.RemoveGroupRelationGroupIDs(ids...)
}

// ClearTargetMsg clears all "target_msg" edges to the GroupMsg entity.
func (gu *GroupUpdate) ClearTargetMsg() *GroupUpdate {
	gu.mutation.ClearTargetMsg()
	return gu
}

// RemoveTargetMsgIDs removes the "target_msg" edge to GroupMsg entities by IDs.
func (gu *GroupUpdate) RemoveTargetMsgIDs(ids ...int) *GroupUpdate {
	gu.mutation.RemoveTargetMsgIDs(ids...)
	return gu
}

// RemoveTargetMsg removes "target_msg" edges to GroupMsg entities.
func (gu *GroupUpdate) RemoveTargetMsg(g ...*GroupMsg) *GroupUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gu.RemoveTargetMsgIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GroupUpdate) Save(ctx context.Context) (int, error) {
	gu.defaults()
	return withHooks(ctx, gu.sqlSave, gu.mutation, gu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GroupUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GroupUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GroupUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gu *GroupUpdate) defaults() {
	if _, ok := gu.mutation.UpdatedAt(); !ok {
		v := group.UpdateDefaultUpdatedAt()
		gu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gu *GroupUpdate) check() error {
	if _, ok := gu.mutation.OwnerUserID(); gu.mutation.OwnerUserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Group.owner_user"`)
	}
	return nil
}

func (gu *GroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt))
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.UpdatedAt(); ok {
		_spec.SetField(group.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := gu.mutation.DeletedAt(); ok {
		_spec.SetField(group.FieldDeletedAt, field.TypeTime, value)
	}
	if gu.mutation.DeletedAtCleared() {
		_spec.ClearField(group.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.SetField(group.FieldName, field.TypeString, value)
	}
	if value, ok := gu.mutation.Img(); ok {
		_spec.SetField(group.FieldImg, field.TypeString, value)
	}
	if value, ok := gu.mutation.Desc(); ok {
		_spec.SetField(group.FieldDesc, field.TypeString, value)
	}
	if gu.mutation.OwnerUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.OwnerUserTable,
			Columns: []string{group.OwnerUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.OwnerUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.OwnerUserTable,
			Columns: []string{group.OwnerUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.GroupRelationGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.GroupRelationGroupTable,
			Columns: []string{group.GroupRelationGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grouprelation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedGroupRelationGroupIDs(); len(nodes) > 0 && !gu.mutation.GroupRelationGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.GroupRelationGroupTable,
			Columns: []string{group.GroupRelationGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grouprelation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.GroupRelationGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.GroupRelationGroupTable,
			Columns: []string{group.GroupRelationGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grouprelation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.TargetMsgCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.TargetMsgTable,
			Columns: []string{group.TargetMsgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupmsg.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedTargetMsgIDs(); len(nodes) > 0 && !gu.mutation.TargetMsgCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.TargetMsgTable,
			Columns: []string{group.TargetMsgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupmsg.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.TargetMsgIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.TargetMsgTable,
			Columns: []string{group.TargetMsgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupmsg.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gu.mutation.done = true
	return n, nil
}

// GroupUpdateOne is the builder for updating a single Group entity.
type GroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GroupMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (guo *GroupUpdateOne) SetUpdatedAt(t time.Time) *GroupUpdateOne {
	guo.mutation.SetUpdatedAt(t)
	return guo
}

// SetDeletedAt sets the "deleted_at" field.
func (guo *GroupUpdateOne) SetDeletedAt(t time.Time) *GroupUpdateOne {
	guo.mutation.SetDeletedAt(t)
	return guo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableDeletedAt(t *time.Time) *GroupUpdateOne {
	if t != nil {
		guo.SetDeletedAt(*t)
	}
	return guo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (guo *GroupUpdateOne) ClearDeletedAt() *GroupUpdateOne {
	guo.mutation.ClearDeletedAt()
	return guo
}

// SetName sets the "name" field.
func (guo *GroupUpdateOne) SetName(s string) *GroupUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// SetUserID sets the "user_id" field.
func (guo *GroupUpdateOne) SetUserID(i int) *GroupUpdateOne {
	guo.mutation.SetUserID(i)
	return guo
}

// SetImg sets the "img" field.
func (guo *GroupUpdateOne) SetImg(s string) *GroupUpdateOne {
	guo.mutation.SetImg(s)
	return guo
}

// SetNillableImg sets the "img" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableImg(s *string) *GroupUpdateOne {
	if s != nil {
		guo.SetImg(*s)
	}
	return guo
}

// SetDesc sets the "desc" field.
func (guo *GroupUpdateOne) SetDesc(s string) *GroupUpdateOne {
	guo.mutation.SetDesc(s)
	return guo
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableDesc(s *string) *GroupUpdateOne {
	if s != nil {
		guo.SetDesc(*s)
	}
	return guo
}

// SetOwnerUserID sets the "owner_user" edge to the User entity by ID.
func (guo *GroupUpdateOne) SetOwnerUserID(id int) *GroupUpdateOne {
	guo.mutation.SetOwnerUserID(id)
	return guo
}

// SetOwnerUser sets the "owner_user" edge to the User entity.
func (guo *GroupUpdateOne) SetOwnerUser(u *User) *GroupUpdateOne {
	return guo.SetOwnerUserID(u.ID)
}

// AddGroupRelationGroupIDs adds the "group_relation_group" edge to the GroupRelation entity by IDs.
func (guo *GroupUpdateOne) AddGroupRelationGroupIDs(ids ...int) *GroupUpdateOne {
	guo.mutation.AddGroupRelationGroupIDs(ids...)
	return guo
}

// AddGroupRelationGroup adds the "group_relation_group" edges to the GroupRelation entity.
func (guo *GroupUpdateOne) AddGroupRelationGroup(g ...*GroupRelation) *GroupUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return guo.AddGroupRelationGroupIDs(ids...)
}

// AddTargetMsgIDs adds the "target_msg" edge to the GroupMsg entity by IDs.
func (guo *GroupUpdateOne) AddTargetMsgIDs(ids ...int) *GroupUpdateOne {
	guo.mutation.AddTargetMsgIDs(ids...)
	return guo
}

// AddTargetMsg adds the "target_msg" edges to the GroupMsg entity.
func (guo *GroupUpdateOne) AddTargetMsg(g ...*GroupMsg) *GroupUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return guo.AddTargetMsgIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (guo *GroupUpdateOne) Mutation() *GroupMutation {
	return guo.mutation
}

// ClearOwnerUser clears the "owner_user" edge to the User entity.
func (guo *GroupUpdateOne) ClearOwnerUser() *GroupUpdateOne {
	guo.mutation.ClearOwnerUser()
	return guo
}

// ClearGroupRelationGroup clears all "group_relation_group" edges to the GroupRelation entity.
func (guo *GroupUpdateOne) ClearGroupRelationGroup() *GroupUpdateOne {
	guo.mutation.ClearGroupRelationGroup()
	return guo
}

// RemoveGroupRelationGroupIDs removes the "group_relation_group" edge to GroupRelation entities by IDs.
func (guo *GroupUpdateOne) RemoveGroupRelationGroupIDs(ids ...int) *GroupUpdateOne {
	guo.mutation.RemoveGroupRelationGroupIDs(ids...)
	return guo
}

// RemoveGroupRelationGroup removes "group_relation_group" edges to GroupRelation entities.
func (guo *GroupUpdateOne) RemoveGroupRelationGroup(g ...*GroupRelation) *GroupUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return guo.RemoveGroupRelationGroupIDs(ids...)
}

// ClearTargetMsg clears all "target_msg" edges to the GroupMsg entity.
func (guo *GroupUpdateOne) ClearTargetMsg() *GroupUpdateOne {
	guo.mutation.ClearTargetMsg()
	return guo
}

// RemoveTargetMsgIDs removes the "target_msg" edge to GroupMsg entities by IDs.
func (guo *GroupUpdateOne) RemoveTargetMsgIDs(ids ...int) *GroupUpdateOne {
	guo.mutation.RemoveTargetMsgIDs(ids...)
	return guo
}

// RemoveTargetMsg removes "target_msg" edges to GroupMsg entities.
func (guo *GroupUpdateOne) RemoveTargetMsg(g ...*GroupMsg) *GroupUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return guo.RemoveTargetMsgIDs(ids...)
}

// Where appends a list predicates to the GroupUpdate builder.
func (guo *GroupUpdateOne) Where(ps ...predicate.Group) *GroupUpdateOne {
	guo.mutation.Where(ps...)
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GroupUpdateOne) Select(field string, fields ...string) *GroupUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Group entity.
func (guo *GroupUpdateOne) Save(ctx context.Context) (*Group, error) {
	guo.defaults()
	return withHooks(ctx, guo.sqlSave, guo.mutation, guo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GroupUpdateOne) SaveX(ctx context.Context) *Group {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GroupUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GroupUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guo *GroupUpdateOne) defaults() {
	if _, ok := guo.mutation.UpdatedAt(); !ok {
		v := group.UpdateDefaultUpdatedAt()
		guo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (guo *GroupUpdateOne) check() error {
	if _, ok := guo.mutation.OwnerUserID(); guo.mutation.OwnerUserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Group.owner_user"`)
	}
	return nil
}

func (guo *GroupUpdateOne) sqlSave(ctx context.Context) (_node *Group, err error) {
	if err := guo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt))
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Group.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, group.FieldID)
		for _, f := range fields {
			if !group.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != group.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.UpdatedAt(); ok {
		_spec.SetField(group.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := guo.mutation.DeletedAt(); ok {
		_spec.SetField(group.FieldDeletedAt, field.TypeTime, value)
	}
	if guo.mutation.DeletedAtCleared() {
		_spec.ClearField(group.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := guo.mutation.Name(); ok {
		_spec.SetField(group.FieldName, field.TypeString, value)
	}
	if value, ok := guo.mutation.Img(); ok {
		_spec.SetField(group.FieldImg, field.TypeString, value)
	}
	if value, ok := guo.mutation.Desc(); ok {
		_spec.SetField(group.FieldDesc, field.TypeString, value)
	}
	if guo.mutation.OwnerUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.OwnerUserTable,
			Columns: []string{group.OwnerUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.OwnerUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.OwnerUserTable,
			Columns: []string{group.OwnerUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.GroupRelationGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.GroupRelationGroupTable,
			Columns: []string{group.GroupRelationGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grouprelation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedGroupRelationGroupIDs(); len(nodes) > 0 && !guo.mutation.GroupRelationGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.GroupRelationGroupTable,
			Columns: []string{group.GroupRelationGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grouprelation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.GroupRelationGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.GroupRelationGroupTable,
			Columns: []string{group.GroupRelationGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grouprelation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.TargetMsgCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.TargetMsgTable,
			Columns: []string{group.TargetMsgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupmsg.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedTargetMsgIDs(); len(nodes) > 0 && !guo.mutation.TargetMsgCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.TargetMsgTable,
			Columns: []string{group.TargetMsgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupmsg.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.TargetMsgIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.TargetMsgTable,
			Columns: []string{group.TargetMsgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupmsg.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Group{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	guo.mutation.done = true
	return _node, nil
}
