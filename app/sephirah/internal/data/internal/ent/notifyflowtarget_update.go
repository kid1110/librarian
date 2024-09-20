// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflow"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflowtarget"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifytarget"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/model"
)

// NotifyFlowTargetUpdate is the builder for updating NotifyFlowTarget entities.
type NotifyFlowTargetUpdate struct {
	config
	hooks    []Hook
	mutation *NotifyFlowTargetMutation
}

// Where appends a list predicates to the NotifyFlowTargetUpdate builder.
func (nftu *NotifyFlowTargetUpdate) Where(ps ...predicate.NotifyFlowTarget) *NotifyFlowTargetUpdate {
	nftu.mutation.Where(ps...)
	return nftu
}

// SetNotifyFlowID sets the "notify_flow_id" field.
func (nftu *NotifyFlowTargetUpdate) SetNotifyFlowID(mi model.InternalID) *NotifyFlowTargetUpdate {
	nftu.mutation.SetNotifyFlowID(mi)
	return nftu
}

// SetNillableNotifyFlowID sets the "notify_flow_id" field if the given value is not nil.
func (nftu *NotifyFlowTargetUpdate) SetNillableNotifyFlowID(mi *model.InternalID) *NotifyFlowTargetUpdate {
	if mi != nil {
		nftu.SetNotifyFlowID(*mi)
	}
	return nftu
}

// SetNotifyTargetID sets the "notify_target_id" field.
func (nftu *NotifyFlowTargetUpdate) SetNotifyTargetID(mi model.InternalID) *NotifyFlowTargetUpdate {
	nftu.mutation.SetNotifyTargetID(mi)
	return nftu
}

// SetNillableNotifyTargetID sets the "notify_target_id" field if the given value is not nil.
func (nftu *NotifyFlowTargetUpdate) SetNillableNotifyTargetID(mi *model.InternalID) *NotifyFlowTargetUpdate {
	if mi != nil {
		nftu.SetNotifyTargetID(*mi)
	}
	return nftu
}

// SetFilterIncludeKeywords sets the "filter_include_keywords" field.
func (nftu *NotifyFlowTargetUpdate) SetFilterIncludeKeywords(s []string) *NotifyFlowTargetUpdate {
	nftu.mutation.SetFilterIncludeKeywords(s)
	return nftu
}

// AppendFilterIncludeKeywords appends s to the "filter_include_keywords" field.
func (nftu *NotifyFlowTargetUpdate) AppendFilterIncludeKeywords(s []string) *NotifyFlowTargetUpdate {
	nftu.mutation.AppendFilterIncludeKeywords(s)
	return nftu
}

// SetFilterExcludeKeywords sets the "filter_exclude_keywords" field.
func (nftu *NotifyFlowTargetUpdate) SetFilterExcludeKeywords(s []string) *NotifyFlowTargetUpdate {
	nftu.mutation.SetFilterExcludeKeywords(s)
	return nftu
}

// AppendFilterExcludeKeywords appends s to the "filter_exclude_keywords" field.
func (nftu *NotifyFlowTargetUpdate) AppendFilterExcludeKeywords(s []string) *NotifyFlowTargetUpdate {
	nftu.mutation.AppendFilterExcludeKeywords(s)
	return nftu
}

// SetUpdatedAt sets the "updated_at" field.
func (nftu *NotifyFlowTargetUpdate) SetUpdatedAt(t time.Time) *NotifyFlowTargetUpdate {
	nftu.mutation.SetUpdatedAt(t)
	return nftu
}

// SetCreatedAt sets the "created_at" field.
func (nftu *NotifyFlowTargetUpdate) SetCreatedAt(t time.Time) *NotifyFlowTargetUpdate {
	nftu.mutation.SetCreatedAt(t)
	return nftu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nftu *NotifyFlowTargetUpdate) SetNillableCreatedAt(t *time.Time) *NotifyFlowTargetUpdate {
	if t != nil {
		nftu.SetCreatedAt(*t)
	}
	return nftu
}

// SetNotifyFlow sets the "notify_flow" edge to the NotifyFlow entity.
func (nftu *NotifyFlowTargetUpdate) SetNotifyFlow(n *NotifyFlow) *NotifyFlowTargetUpdate {
	return nftu.SetNotifyFlowID(n.ID)
}

// SetNotifyTarget sets the "notify_target" edge to the NotifyTarget entity.
func (nftu *NotifyFlowTargetUpdate) SetNotifyTarget(n *NotifyTarget) *NotifyFlowTargetUpdate {
	return nftu.SetNotifyTargetID(n.ID)
}

// Mutation returns the NotifyFlowTargetMutation object of the builder.
func (nftu *NotifyFlowTargetUpdate) Mutation() *NotifyFlowTargetMutation {
	return nftu.mutation
}

// ClearNotifyFlow clears the "notify_flow" edge to the NotifyFlow entity.
func (nftu *NotifyFlowTargetUpdate) ClearNotifyFlow() *NotifyFlowTargetUpdate {
	nftu.mutation.ClearNotifyFlow()
	return nftu
}

// ClearNotifyTarget clears the "notify_target" edge to the NotifyTarget entity.
func (nftu *NotifyFlowTargetUpdate) ClearNotifyTarget() *NotifyFlowTargetUpdate {
	nftu.mutation.ClearNotifyTarget()
	return nftu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nftu *NotifyFlowTargetUpdate) Save(ctx context.Context) (int, error) {
	nftu.defaults()
	return withHooks(ctx, nftu.sqlSave, nftu.mutation, nftu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nftu *NotifyFlowTargetUpdate) SaveX(ctx context.Context) int {
	affected, err := nftu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nftu *NotifyFlowTargetUpdate) Exec(ctx context.Context) error {
	_, err := nftu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nftu *NotifyFlowTargetUpdate) ExecX(ctx context.Context) {
	if err := nftu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nftu *NotifyFlowTargetUpdate) defaults() {
	if _, ok := nftu.mutation.UpdatedAt(); !ok {
		v := notifyflowtarget.UpdateDefaultUpdatedAt()
		nftu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nftu *NotifyFlowTargetUpdate) check() error {
	if nftu.mutation.NotifyFlowCleared() && len(nftu.mutation.NotifyFlowIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "NotifyFlowTarget.notify_flow"`)
	}
	if nftu.mutation.NotifyTargetCleared() && len(nftu.mutation.NotifyTargetIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "NotifyFlowTarget.notify_target"`)
	}
	return nil
}

func (nftu *NotifyFlowTargetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := nftu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(notifyflowtarget.Table, notifyflowtarget.Columns, sqlgraph.NewFieldSpec(notifyflowtarget.FieldID, field.TypeInt))
	if ps := nftu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nftu.mutation.FilterIncludeKeywords(); ok {
		_spec.SetField(notifyflowtarget.FieldFilterIncludeKeywords, field.TypeJSON, value)
	}
	if value, ok := nftu.mutation.AppendedFilterIncludeKeywords(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, notifyflowtarget.FieldFilterIncludeKeywords, value)
		})
	}
	if value, ok := nftu.mutation.FilterExcludeKeywords(); ok {
		_spec.SetField(notifyflowtarget.FieldFilterExcludeKeywords, field.TypeJSON, value)
	}
	if value, ok := nftu.mutation.AppendedFilterExcludeKeywords(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, notifyflowtarget.FieldFilterExcludeKeywords, value)
		})
	}
	if value, ok := nftu.mutation.UpdatedAt(); ok {
		_spec.SetField(notifyflowtarget.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := nftu.mutation.CreatedAt(); ok {
		_spec.SetField(notifyflowtarget.FieldCreatedAt, field.TypeTime, value)
	}
	if nftu.mutation.NotifyFlowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   notifyflowtarget.NotifyFlowTable,
			Columns: []string{notifyflowtarget.NotifyFlowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflow.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nftu.mutation.NotifyFlowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   notifyflowtarget.NotifyFlowTable,
			Columns: []string{notifyflowtarget.NotifyFlowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflow.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nftu.mutation.NotifyTargetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   notifyflowtarget.NotifyTargetTable,
			Columns: []string{notifyflowtarget.NotifyTargetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifytarget.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nftu.mutation.NotifyTargetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   notifyflowtarget.NotifyTargetTable,
			Columns: []string{notifyflowtarget.NotifyTargetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifytarget.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nftu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notifyflowtarget.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nftu.mutation.done = true
	return n, nil
}

// NotifyFlowTargetUpdateOne is the builder for updating a single NotifyFlowTarget entity.
type NotifyFlowTargetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NotifyFlowTargetMutation
}

// SetNotifyFlowID sets the "notify_flow_id" field.
func (nftuo *NotifyFlowTargetUpdateOne) SetNotifyFlowID(mi model.InternalID) *NotifyFlowTargetUpdateOne {
	nftuo.mutation.SetNotifyFlowID(mi)
	return nftuo
}

// SetNillableNotifyFlowID sets the "notify_flow_id" field if the given value is not nil.
func (nftuo *NotifyFlowTargetUpdateOne) SetNillableNotifyFlowID(mi *model.InternalID) *NotifyFlowTargetUpdateOne {
	if mi != nil {
		nftuo.SetNotifyFlowID(*mi)
	}
	return nftuo
}

// SetNotifyTargetID sets the "notify_target_id" field.
func (nftuo *NotifyFlowTargetUpdateOne) SetNotifyTargetID(mi model.InternalID) *NotifyFlowTargetUpdateOne {
	nftuo.mutation.SetNotifyTargetID(mi)
	return nftuo
}

// SetNillableNotifyTargetID sets the "notify_target_id" field if the given value is not nil.
func (nftuo *NotifyFlowTargetUpdateOne) SetNillableNotifyTargetID(mi *model.InternalID) *NotifyFlowTargetUpdateOne {
	if mi != nil {
		nftuo.SetNotifyTargetID(*mi)
	}
	return nftuo
}

// SetFilterIncludeKeywords sets the "filter_include_keywords" field.
func (nftuo *NotifyFlowTargetUpdateOne) SetFilterIncludeKeywords(s []string) *NotifyFlowTargetUpdateOne {
	nftuo.mutation.SetFilterIncludeKeywords(s)
	return nftuo
}

// AppendFilterIncludeKeywords appends s to the "filter_include_keywords" field.
func (nftuo *NotifyFlowTargetUpdateOne) AppendFilterIncludeKeywords(s []string) *NotifyFlowTargetUpdateOne {
	nftuo.mutation.AppendFilterIncludeKeywords(s)
	return nftuo
}

// SetFilterExcludeKeywords sets the "filter_exclude_keywords" field.
func (nftuo *NotifyFlowTargetUpdateOne) SetFilterExcludeKeywords(s []string) *NotifyFlowTargetUpdateOne {
	nftuo.mutation.SetFilterExcludeKeywords(s)
	return nftuo
}

// AppendFilterExcludeKeywords appends s to the "filter_exclude_keywords" field.
func (nftuo *NotifyFlowTargetUpdateOne) AppendFilterExcludeKeywords(s []string) *NotifyFlowTargetUpdateOne {
	nftuo.mutation.AppendFilterExcludeKeywords(s)
	return nftuo
}

// SetUpdatedAt sets the "updated_at" field.
func (nftuo *NotifyFlowTargetUpdateOne) SetUpdatedAt(t time.Time) *NotifyFlowTargetUpdateOne {
	nftuo.mutation.SetUpdatedAt(t)
	return nftuo
}

// SetCreatedAt sets the "created_at" field.
func (nftuo *NotifyFlowTargetUpdateOne) SetCreatedAt(t time.Time) *NotifyFlowTargetUpdateOne {
	nftuo.mutation.SetCreatedAt(t)
	return nftuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nftuo *NotifyFlowTargetUpdateOne) SetNillableCreatedAt(t *time.Time) *NotifyFlowTargetUpdateOne {
	if t != nil {
		nftuo.SetCreatedAt(*t)
	}
	return nftuo
}

// SetNotifyFlow sets the "notify_flow" edge to the NotifyFlow entity.
func (nftuo *NotifyFlowTargetUpdateOne) SetNotifyFlow(n *NotifyFlow) *NotifyFlowTargetUpdateOne {
	return nftuo.SetNotifyFlowID(n.ID)
}

// SetNotifyTarget sets the "notify_target" edge to the NotifyTarget entity.
func (nftuo *NotifyFlowTargetUpdateOne) SetNotifyTarget(n *NotifyTarget) *NotifyFlowTargetUpdateOne {
	return nftuo.SetNotifyTargetID(n.ID)
}

// Mutation returns the NotifyFlowTargetMutation object of the builder.
func (nftuo *NotifyFlowTargetUpdateOne) Mutation() *NotifyFlowTargetMutation {
	return nftuo.mutation
}

// ClearNotifyFlow clears the "notify_flow" edge to the NotifyFlow entity.
func (nftuo *NotifyFlowTargetUpdateOne) ClearNotifyFlow() *NotifyFlowTargetUpdateOne {
	nftuo.mutation.ClearNotifyFlow()
	return nftuo
}

// ClearNotifyTarget clears the "notify_target" edge to the NotifyTarget entity.
func (nftuo *NotifyFlowTargetUpdateOne) ClearNotifyTarget() *NotifyFlowTargetUpdateOne {
	nftuo.mutation.ClearNotifyTarget()
	return nftuo
}

// Where appends a list predicates to the NotifyFlowTargetUpdate builder.
func (nftuo *NotifyFlowTargetUpdateOne) Where(ps ...predicate.NotifyFlowTarget) *NotifyFlowTargetUpdateOne {
	nftuo.mutation.Where(ps...)
	return nftuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nftuo *NotifyFlowTargetUpdateOne) Select(field string, fields ...string) *NotifyFlowTargetUpdateOne {
	nftuo.fields = append([]string{field}, fields...)
	return nftuo
}

// Save executes the query and returns the updated NotifyFlowTarget entity.
func (nftuo *NotifyFlowTargetUpdateOne) Save(ctx context.Context) (*NotifyFlowTarget, error) {
	nftuo.defaults()
	return withHooks(ctx, nftuo.sqlSave, nftuo.mutation, nftuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nftuo *NotifyFlowTargetUpdateOne) SaveX(ctx context.Context) *NotifyFlowTarget {
	node, err := nftuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nftuo *NotifyFlowTargetUpdateOne) Exec(ctx context.Context) error {
	_, err := nftuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nftuo *NotifyFlowTargetUpdateOne) ExecX(ctx context.Context) {
	if err := nftuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nftuo *NotifyFlowTargetUpdateOne) defaults() {
	if _, ok := nftuo.mutation.UpdatedAt(); !ok {
		v := notifyflowtarget.UpdateDefaultUpdatedAt()
		nftuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nftuo *NotifyFlowTargetUpdateOne) check() error {
	if nftuo.mutation.NotifyFlowCleared() && len(nftuo.mutation.NotifyFlowIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "NotifyFlowTarget.notify_flow"`)
	}
	if nftuo.mutation.NotifyTargetCleared() && len(nftuo.mutation.NotifyTargetIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "NotifyFlowTarget.notify_target"`)
	}
	return nil
}

func (nftuo *NotifyFlowTargetUpdateOne) sqlSave(ctx context.Context) (_node *NotifyFlowTarget, err error) {
	if err := nftuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(notifyflowtarget.Table, notifyflowtarget.Columns, sqlgraph.NewFieldSpec(notifyflowtarget.FieldID, field.TypeInt))
	id, ok := nftuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NotifyFlowTarget.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nftuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notifyflowtarget.FieldID)
		for _, f := range fields {
			if !notifyflowtarget.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != notifyflowtarget.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nftuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nftuo.mutation.FilterIncludeKeywords(); ok {
		_spec.SetField(notifyflowtarget.FieldFilterIncludeKeywords, field.TypeJSON, value)
	}
	if value, ok := nftuo.mutation.AppendedFilterIncludeKeywords(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, notifyflowtarget.FieldFilterIncludeKeywords, value)
		})
	}
	if value, ok := nftuo.mutation.FilterExcludeKeywords(); ok {
		_spec.SetField(notifyflowtarget.FieldFilterExcludeKeywords, field.TypeJSON, value)
	}
	if value, ok := nftuo.mutation.AppendedFilterExcludeKeywords(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, notifyflowtarget.FieldFilterExcludeKeywords, value)
		})
	}
	if value, ok := nftuo.mutation.UpdatedAt(); ok {
		_spec.SetField(notifyflowtarget.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := nftuo.mutation.CreatedAt(); ok {
		_spec.SetField(notifyflowtarget.FieldCreatedAt, field.TypeTime, value)
	}
	if nftuo.mutation.NotifyFlowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   notifyflowtarget.NotifyFlowTable,
			Columns: []string{notifyflowtarget.NotifyFlowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflow.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nftuo.mutation.NotifyFlowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   notifyflowtarget.NotifyFlowTable,
			Columns: []string{notifyflowtarget.NotifyFlowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflow.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nftuo.mutation.NotifyTargetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   notifyflowtarget.NotifyTargetTable,
			Columns: []string{notifyflowtarget.NotifyTargetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifytarget.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nftuo.mutation.NotifyTargetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   notifyflowtarget.NotifyTargetTable,
			Columns: []string{notifyflowtarget.NotifyTargetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifytarget.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &NotifyFlowTarget{config: nftuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nftuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notifyflowtarget.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nftuo.mutation.done = true
	return _node, nil
}
