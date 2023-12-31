// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-zero-chat/ent/group"
	"go-zero-chat/ent/groupmsg"
	"go-zero-chat/ent/predicate"
	"go-zero-chat/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GroupMsgUpdate is the builder for updating GroupMsg entities.
type GroupMsgUpdate struct {
	config
	hooks    []Hook
	mutation *GroupMsgMutation
}

// Where appends a list predicates to the GroupMsgUpdate builder.
func (gmu *GroupMsgUpdate) Where(ps ...predicate.GroupMsg) *GroupMsgUpdate {
	gmu.mutation.Where(ps...)
	return gmu
}

// SetUpdatedAt sets the "updated_at" field.
func (gmu *GroupMsgUpdate) SetUpdatedAt(t time.Time) *GroupMsgUpdate {
	gmu.mutation.SetUpdatedAt(t)
	return gmu
}

// SetDeletedAt sets the "deleted_at" field.
func (gmu *GroupMsgUpdate) SetDeletedAt(t time.Time) *GroupMsgUpdate {
	gmu.mutation.SetDeletedAt(t)
	return gmu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gmu *GroupMsgUpdate) SetNillableDeletedAt(t *time.Time) *GroupMsgUpdate {
	if t != nil {
		gmu.SetDeletedAt(*t)
	}
	return gmu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (gmu *GroupMsgUpdate) ClearDeletedAt() *GroupMsgUpdate {
	gmu.mutation.ClearDeletedAt()
	return gmu
}

// SetOwnerID sets the "owner_id" field.
func (gmu *GroupMsgUpdate) SetOwnerID(i int) *GroupMsgUpdate {
	gmu.mutation.SetOwnerID(i)
	return gmu
}

// SetTargetID sets the "target_id" field.
func (gmu *GroupMsgUpdate) SetTargetID(i int) *GroupMsgUpdate {
	gmu.mutation.SetTargetID(i)
	return gmu
}

// SetMedia sets the "media" field.
func (gmu *GroupMsgUpdate) SetMedia(i int) *GroupMsgUpdate {
	gmu.mutation.ResetMedia()
	gmu.mutation.SetMedia(i)
	return gmu
}

// AddMedia adds i to the "media" field.
func (gmu *GroupMsgUpdate) AddMedia(i int) *GroupMsgUpdate {
	gmu.mutation.AddMedia(i)
	return gmu
}

// SetContent sets the "content" field.
func (gmu *GroupMsgUpdate) SetContent(s string) *GroupMsgUpdate {
	gmu.mutation.SetContent(s)
	return gmu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (gmu *GroupMsgUpdate) SetNillableContent(s *string) *GroupMsgUpdate {
	if s != nil {
		gmu.SetContent(*s)
	}
	return gmu
}

// SetPic sets the "pic" field.
func (gmu *GroupMsgUpdate) SetPic(s string) *GroupMsgUpdate {
	gmu.mutation.SetPic(s)
	return gmu
}

// SetNillablePic sets the "pic" field if the given value is not nil.
func (gmu *GroupMsgUpdate) SetNillablePic(s *string) *GroupMsgUpdate {
	if s != nil {
		gmu.SetPic(*s)
	}
	return gmu
}

// SetURL sets the "url" field.
func (gmu *GroupMsgUpdate) SetURL(s string) *GroupMsgUpdate {
	gmu.mutation.SetURL(s)
	return gmu
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (gmu *GroupMsgUpdate) SetNillableURL(s *string) *GroupMsgUpdate {
	if s != nil {
		gmu.SetURL(*s)
	}
	return gmu
}

// SetDesc sets the "desc" field.
func (gmu *GroupMsgUpdate) SetDesc(s string) *GroupMsgUpdate {
	gmu.mutation.SetDesc(s)
	return gmu
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (gmu *GroupMsgUpdate) SetNillableDesc(s *string) *GroupMsgUpdate {
	if s != nil {
		gmu.SetDesc(*s)
	}
	return gmu
}

// SetAmount sets the "amount" field.
func (gmu *GroupMsgUpdate) SetAmount(i int) *GroupMsgUpdate {
	gmu.mutation.ResetAmount()
	gmu.mutation.SetAmount(i)
	return gmu
}

// AddAmount adds i to the "amount" field.
func (gmu *GroupMsgUpdate) AddAmount(i int) *GroupMsgUpdate {
	gmu.mutation.AddAmount(i)
	return gmu
}

// SetType sets the "type" field.
func (gmu *GroupMsgUpdate) SetType(i int) *GroupMsgUpdate {
	gmu.mutation.ResetType()
	gmu.mutation.SetType(i)
	return gmu
}

// AddType adds i to the "type" field.
func (gmu *GroupMsgUpdate) AddType(i int) *GroupMsgUpdate {
	gmu.mutation.AddType(i)
	return gmu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (gmu *GroupMsgUpdate) SetUserID(id int) *GroupMsgUpdate {
	gmu.mutation.SetUserID(id)
	return gmu
}

// SetUser sets the "user" edge to the User entity.
func (gmu *GroupMsgUpdate) SetUser(u *User) *GroupMsgUpdate {
	return gmu.SetUserID(u.ID)
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (gmu *GroupMsgUpdate) SetGroupID(id int) *GroupMsgUpdate {
	gmu.mutation.SetGroupID(id)
	return gmu
}

// SetGroup sets the "group" edge to the Group entity.
func (gmu *GroupMsgUpdate) SetGroup(g *Group) *GroupMsgUpdate {
	return gmu.SetGroupID(g.ID)
}

// Mutation returns the GroupMsgMutation object of the builder.
func (gmu *GroupMsgUpdate) Mutation() *GroupMsgMutation {
	return gmu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (gmu *GroupMsgUpdate) ClearUser() *GroupMsgUpdate {
	gmu.mutation.ClearUser()
	return gmu
}

// ClearGroup clears the "group" edge to the Group entity.
func (gmu *GroupMsgUpdate) ClearGroup() *GroupMsgUpdate {
	gmu.mutation.ClearGroup()
	return gmu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gmu *GroupMsgUpdate) Save(ctx context.Context) (int, error) {
	gmu.defaults()
	return withHooks(ctx, gmu.sqlSave, gmu.mutation, gmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gmu *GroupMsgUpdate) SaveX(ctx context.Context) int {
	affected, err := gmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gmu *GroupMsgUpdate) Exec(ctx context.Context) error {
	_, err := gmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmu *GroupMsgUpdate) ExecX(ctx context.Context) {
	if err := gmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gmu *GroupMsgUpdate) defaults() {
	if _, ok := gmu.mutation.UpdatedAt(); !ok {
		v := groupmsg.UpdateDefaultUpdatedAt()
		gmu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gmu *GroupMsgUpdate) check() error {
	if _, ok := gmu.mutation.UserID(); gmu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GroupMsg.user"`)
	}
	if _, ok := gmu.mutation.GroupID(); gmu.mutation.GroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GroupMsg.group"`)
	}
	return nil
}

func (gmu *GroupMsgUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gmu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(groupmsg.Table, groupmsg.Columns, sqlgraph.NewFieldSpec(groupmsg.FieldID, field.TypeInt))
	if ps := gmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gmu.mutation.UpdatedAt(); ok {
		_spec.SetField(groupmsg.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := gmu.mutation.DeletedAt(); ok {
		_spec.SetField(groupmsg.FieldDeletedAt, field.TypeTime, value)
	}
	if gmu.mutation.DeletedAtCleared() {
		_spec.ClearField(groupmsg.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := gmu.mutation.Media(); ok {
		_spec.SetField(groupmsg.FieldMedia, field.TypeInt, value)
	}
	if value, ok := gmu.mutation.AddedMedia(); ok {
		_spec.AddField(groupmsg.FieldMedia, field.TypeInt, value)
	}
	if value, ok := gmu.mutation.Content(); ok {
		_spec.SetField(groupmsg.FieldContent, field.TypeString, value)
	}
	if value, ok := gmu.mutation.Pic(); ok {
		_spec.SetField(groupmsg.FieldPic, field.TypeString, value)
	}
	if value, ok := gmu.mutation.URL(); ok {
		_spec.SetField(groupmsg.FieldURL, field.TypeString, value)
	}
	if value, ok := gmu.mutation.Desc(); ok {
		_spec.SetField(groupmsg.FieldDesc, field.TypeString, value)
	}
	if value, ok := gmu.mutation.Amount(); ok {
		_spec.SetField(groupmsg.FieldAmount, field.TypeInt, value)
	}
	if value, ok := gmu.mutation.AddedAmount(); ok {
		_spec.AddField(groupmsg.FieldAmount, field.TypeInt, value)
	}
	if value, ok := gmu.mutation.GetType(); ok {
		_spec.SetField(groupmsg.FieldType, field.TypeInt, value)
	}
	if value, ok := gmu.mutation.AddedType(); ok {
		_spec.AddField(groupmsg.FieldType, field.TypeInt, value)
	}
	if gmu.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gmu.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gmu.mutation.GroupCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gmu.mutation.GroupIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupmsg.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gmu.mutation.done = true
	return n, nil
}

// GroupMsgUpdateOne is the builder for updating a single GroupMsg entity.
type GroupMsgUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GroupMsgMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (gmuo *GroupMsgUpdateOne) SetUpdatedAt(t time.Time) *GroupMsgUpdateOne {
	gmuo.mutation.SetUpdatedAt(t)
	return gmuo
}

// SetDeletedAt sets the "deleted_at" field.
func (gmuo *GroupMsgUpdateOne) SetDeletedAt(t time.Time) *GroupMsgUpdateOne {
	gmuo.mutation.SetDeletedAt(t)
	return gmuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gmuo *GroupMsgUpdateOne) SetNillableDeletedAt(t *time.Time) *GroupMsgUpdateOne {
	if t != nil {
		gmuo.SetDeletedAt(*t)
	}
	return gmuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (gmuo *GroupMsgUpdateOne) ClearDeletedAt() *GroupMsgUpdateOne {
	gmuo.mutation.ClearDeletedAt()
	return gmuo
}

// SetOwnerID sets the "owner_id" field.
func (gmuo *GroupMsgUpdateOne) SetOwnerID(i int) *GroupMsgUpdateOne {
	gmuo.mutation.SetOwnerID(i)
	return gmuo
}

// SetTargetID sets the "target_id" field.
func (gmuo *GroupMsgUpdateOne) SetTargetID(i int) *GroupMsgUpdateOne {
	gmuo.mutation.SetTargetID(i)
	return gmuo
}

// SetMedia sets the "media" field.
func (gmuo *GroupMsgUpdateOne) SetMedia(i int) *GroupMsgUpdateOne {
	gmuo.mutation.ResetMedia()
	gmuo.mutation.SetMedia(i)
	return gmuo
}

// AddMedia adds i to the "media" field.
func (gmuo *GroupMsgUpdateOne) AddMedia(i int) *GroupMsgUpdateOne {
	gmuo.mutation.AddMedia(i)
	return gmuo
}

// SetContent sets the "content" field.
func (gmuo *GroupMsgUpdateOne) SetContent(s string) *GroupMsgUpdateOne {
	gmuo.mutation.SetContent(s)
	return gmuo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (gmuo *GroupMsgUpdateOne) SetNillableContent(s *string) *GroupMsgUpdateOne {
	if s != nil {
		gmuo.SetContent(*s)
	}
	return gmuo
}

// SetPic sets the "pic" field.
func (gmuo *GroupMsgUpdateOne) SetPic(s string) *GroupMsgUpdateOne {
	gmuo.mutation.SetPic(s)
	return gmuo
}

// SetNillablePic sets the "pic" field if the given value is not nil.
func (gmuo *GroupMsgUpdateOne) SetNillablePic(s *string) *GroupMsgUpdateOne {
	if s != nil {
		gmuo.SetPic(*s)
	}
	return gmuo
}

// SetURL sets the "url" field.
func (gmuo *GroupMsgUpdateOne) SetURL(s string) *GroupMsgUpdateOne {
	gmuo.mutation.SetURL(s)
	return gmuo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (gmuo *GroupMsgUpdateOne) SetNillableURL(s *string) *GroupMsgUpdateOne {
	if s != nil {
		gmuo.SetURL(*s)
	}
	return gmuo
}

// SetDesc sets the "desc" field.
func (gmuo *GroupMsgUpdateOne) SetDesc(s string) *GroupMsgUpdateOne {
	gmuo.mutation.SetDesc(s)
	return gmuo
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (gmuo *GroupMsgUpdateOne) SetNillableDesc(s *string) *GroupMsgUpdateOne {
	if s != nil {
		gmuo.SetDesc(*s)
	}
	return gmuo
}

// SetAmount sets the "amount" field.
func (gmuo *GroupMsgUpdateOne) SetAmount(i int) *GroupMsgUpdateOne {
	gmuo.mutation.ResetAmount()
	gmuo.mutation.SetAmount(i)
	return gmuo
}

// AddAmount adds i to the "amount" field.
func (gmuo *GroupMsgUpdateOne) AddAmount(i int) *GroupMsgUpdateOne {
	gmuo.mutation.AddAmount(i)
	return gmuo
}

// SetType sets the "type" field.
func (gmuo *GroupMsgUpdateOne) SetType(i int) *GroupMsgUpdateOne {
	gmuo.mutation.ResetType()
	gmuo.mutation.SetType(i)
	return gmuo
}

// AddType adds i to the "type" field.
func (gmuo *GroupMsgUpdateOne) AddType(i int) *GroupMsgUpdateOne {
	gmuo.mutation.AddType(i)
	return gmuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (gmuo *GroupMsgUpdateOne) SetUserID(id int) *GroupMsgUpdateOne {
	gmuo.mutation.SetUserID(id)
	return gmuo
}

// SetUser sets the "user" edge to the User entity.
func (gmuo *GroupMsgUpdateOne) SetUser(u *User) *GroupMsgUpdateOne {
	return gmuo.SetUserID(u.ID)
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (gmuo *GroupMsgUpdateOne) SetGroupID(id int) *GroupMsgUpdateOne {
	gmuo.mutation.SetGroupID(id)
	return gmuo
}

// SetGroup sets the "group" edge to the Group entity.
func (gmuo *GroupMsgUpdateOne) SetGroup(g *Group) *GroupMsgUpdateOne {
	return gmuo.SetGroupID(g.ID)
}

// Mutation returns the GroupMsgMutation object of the builder.
func (gmuo *GroupMsgUpdateOne) Mutation() *GroupMsgMutation {
	return gmuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (gmuo *GroupMsgUpdateOne) ClearUser() *GroupMsgUpdateOne {
	gmuo.mutation.ClearUser()
	return gmuo
}

// ClearGroup clears the "group" edge to the Group entity.
func (gmuo *GroupMsgUpdateOne) ClearGroup() *GroupMsgUpdateOne {
	gmuo.mutation.ClearGroup()
	return gmuo
}

// Where appends a list predicates to the GroupMsgUpdate builder.
func (gmuo *GroupMsgUpdateOne) Where(ps ...predicate.GroupMsg) *GroupMsgUpdateOne {
	gmuo.mutation.Where(ps...)
	return gmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gmuo *GroupMsgUpdateOne) Select(field string, fields ...string) *GroupMsgUpdateOne {
	gmuo.fields = append([]string{field}, fields...)
	return gmuo
}

// Save executes the query and returns the updated GroupMsg entity.
func (gmuo *GroupMsgUpdateOne) Save(ctx context.Context) (*GroupMsg, error) {
	gmuo.defaults()
	return withHooks(ctx, gmuo.sqlSave, gmuo.mutation, gmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gmuo *GroupMsgUpdateOne) SaveX(ctx context.Context) *GroupMsg {
	node, err := gmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gmuo *GroupMsgUpdateOne) Exec(ctx context.Context) error {
	_, err := gmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmuo *GroupMsgUpdateOne) ExecX(ctx context.Context) {
	if err := gmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gmuo *GroupMsgUpdateOne) defaults() {
	if _, ok := gmuo.mutation.UpdatedAt(); !ok {
		v := groupmsg.UpdateDefaultUpdatedAt()
		gmuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gmuo *GroupMsgUpdateOne) check() error {
	if _, ok := gmuo.mutation.UserID(); gmuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GroupMsg.user"`)
	}
	if _, ok := gmuo.mutation.GroupID(); gmuo.mutation.GroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GroupMsg.group"`)
	}
	return nil
}

func (gmuo *GroupMsgUpdateOne) sqlSave(ctx context.Context) (_node *GroupMsg, err error) {
	if err := gmuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(groupmsg.Table, groupmsg.Columns, sqlgraph.NewFieldSpec(groupmsg.FieldID, field.TypeInt))
	id, ok := gmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GroupMsg.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, groupmsg.FieldID)
		for _, f := range fields {
			if !groupmsg.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != groupmsg.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gmuo.mutation.UpdatedAt(); ok {
		_spec.SetField(groupmsg.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := gmuo.mutation.DeletedAt(); ok {
		_spec.SetField(groupmsg.FieldDeletedAt, field.TypeTime, value)
	}
	if gmuo.mutation.DeletedAtCleared() {
		_spec.ClearField(groupmsg.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := gmuo.mutation.Media(); ok {
		_spec.SetField(groupmsg.FieldMedia, field.TypeInt, value)
	}
	if value, ok := gmuo.mutation.AddedMedia(); ok {
		_spec.AddField(groupmsg.FieldMedia, field.TypeInt, value)
	}
	if value, ok := gmuo.mutation.Content(); ok {
		_spec.SetField(groupmsg.FieldContent, field.TypeString, value)
	}
	if value, ok := gmuo.mutation.Pic(); ok {
		_spec.SetField(groupmsg.FieldPic, field.TypeString, value)
	}
	if value, ok := gmuo.mutation.URL(); ok {
		_spec.SetField(groupmsg.FieldURL, field.TypeString, value)
	}
	if value, ok := gmuo.mutation.Desc(); ok {
		_spec.SetField(groupmsg.FieldDesc, field.TypeString, value)
	}
	if value, ok := gmuo.mutation.Amount(); ok {
		_spec.SetField(groupmsg.FieldAmount, field.TypeInt, value)
	}
	if value, ok := gmuo.mutation.AddedAmount(); ok {
		_spec.AddField(groupmsg.FieldAmount, field.TypeInt, value)
	}
	if value, ok := gmuo.mutation.GetType(); ok {
		_spec.SetField(groupmsg.FieldType, field.TypeInt, value)
	}
	if value, ok := gmuo.mutation.AddedType(); ok {
		_spec.AddField(groupmsg.FieldType, field.TypeInt, value)
	}
	if gmuo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gmuo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gmuo.mutation.GroupCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gmuo.mutation.GroupIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &GroupMsg{config: gmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupmsg.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gmuo.mutation.done = true
	return _node, nil
}
