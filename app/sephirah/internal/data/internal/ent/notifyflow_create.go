// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"github.com/tuihub/librarian/model"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflow"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflowsource"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflowtarget"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifytarget"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
)

// NotifyFlowCreate is the builder for creating a NotifyFlow entity.
type NotifyFlowCreate struct {
	config
	mutation *NotifyFlowMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (nfc *NotifyFlowCreate) SetName(s string) *NotifyFlowCreate {
	nfc.mutation.SetName(s)
	return nfc
}

// SetDescription sets the "description" field.
func (nfc *NotifyFlowCreate) SetDescription(s string) *NotifyFlowCreate {
	nfc.mutation.SetDescription(s)
	return nfc
}

// SetStatus sets the "status" field.
func (nfc *NotifyFlowCreate) SetStatus(n notifyflow.Status) *NotifyFlowCreate {
	nfc.mutation.SetStatus(n)
	return nfc
}

// SetUpdatedAt sets the "updated_at" field.
func (nfc *NotifyFlowCreate) SetUpdatedAt(t time.Time) *NotifyFlowCreate {
	nfc.mutation.SetUpdatedAt(t)
	return nfc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (nfc *NotifyFlowCreate) SetNillableUpdatedAt(t *time.Time) *NotifyFlowCreate {
	if t != nil {
		nfc.SetUpdatedAt(*t)
	}
	return nfc
}

// SetCreatedAt sets the "created_at" field.
func (nfc *NotifyFlowCreate) SetCreatedAt(t time.Time) *NotifyFlowCreate {
	nfc.mutation.SetCreatedAt(t)
	return nfc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nfc *NotifyFlowCreate) SetNillableCreatedAt(t *time.Time) *NotifyFlowCreate {
	if t != nil {
		nfc.SetCreatedAt(*t)
	}
	return nfc
}

// SetID sets the "id" field.
func (nfc *NotifyFlowCreate) SetID(mi model.InternalID) *NotifyFlowCreate {
	nfc.mutation.SetID(mi)
	return nfc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (nfc *NotifyFlowCreate) SetOwnerID(id model.InternalID) *NotifyFlowCreate {
	nfc.mutation.SetOwnerID(id)
	return nfc
}

// SetOwner sets the "owner" edge to the User entity.
func (nfc *NotifyFlowCreate) SetOwner(u *User) *NotifyFlowCreate {
	return nfc.SetOwnerID(u.ID)
}

// AddNotifyTargetIDs adds the "notify_target" edge to the NotifyTarget entity by IDs.
func (nfc *NotifyFlowCreate) AddNotifyTargetIDs(ids ...model.InternalID) *NotifyFlowCreate {
	nfc.mutation.AddNotifyTargetIDs(ids...)
	return nfc
}

// AddNotifyTarget adds the "notify_target" edges to the NotifyTarget entity.
func (nfc *NotifyFlowCreate) AddNotifyTarget(n ...*NotifyTarget) *NotifyFlowCreate {
	ids := make([]model.InternalID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nfc.AddNotifyTargetIDs(ids...)
}

// AddFeedConfigIDs adds the "feed_config" edge to the FeedConfig entity by IDs.
func (nfc *NotifyFlowCreate) AddFeedConfigIDs(ids ...model.InternalID) *NotifyFlowCreate {
	nfc.mutation.AddFeedConfigIDs(ids...)
	return nfc
}

// AddFeedConfig adds the "feed_config" edges to the FeedConfig entity.
func (nfc *NotifyFlowCreate) AddFeedConfig(f ...*FeedConfig) *NotifyFlowCreate {
	ids := make([]model.InternalID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return nfc.AddFeedConfigIDs(ids...)
}

// AddNotifyFlowTargetIDs adds the "notify_flow_target" edge to the NotifyFlowTarget entity by IDs.
func (nfc *NotifyFlowCreate) AddNotifyFlowTargetIDs(ids ...int) *NotifyFlowCreate {
	nfc.mutation.AddNotifyFlowTargetIDs(ids...)
	return nfc
}

// AddNotifyFlowTarget adds the "notify_flow_target" edges to the NotifyFlowTarget entity.
func (nfc *NotifyFlowCreate) AddNotifyFlowTarget(n ...*NotifyFlowTarget) *NotifyFlowCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nfc.AddNotifyFlowTargetIDs(ids...)
}

// AddNotifyFlowSourceIDs adds the "notify_flow_source" edge to the NotifyFlowSource entity by IDs.
func (nfc *NotifyFlowCreate) AddNotifyFlowSourceIDs(ids ...int) *NotifyFlowCreate {
	nfc.mutation.AddNotifyFlowSourceIDs(ids...)
	return nfc
}

// AddNotifyFlowSource adds the "notify_flow_source" edges to the NotifyFlowSource entity.
func (nfc *NotifyFlowCreate) AddNotifyFlowSource(n ...*NotifyFlowSource) *NotifyFlowCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nfc.AddNotifyFlowSourceIDs(ids...)
}

// Mutation returns the NotifyFlowMutation object of the builder.
func (nfc *NotifyFlowCreate) Mutation() *NotifyFlowMutation {
	return nfc.mutation
}

// Save creates the NotifyFlow in the database.
func (nfc *NotifyFlowCreate) Save(ctx context.Context) (*NotifyFlow, error) {
	nfc.defaults()
	return withHooks(ctx, nfc.sqlSave, nfc.mutation, nfc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nfc *NotifyFlowCreate) SaveX(ctx context.Context) *NotifyFlow {
	v, err := nfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nfc *NotifyFlowCreate) Exec(ctx context.Context) error {
	_, err := nfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nfc *NotifyFlowCreate) ExecX(ctx context.Context) {
	if err := nfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nfc *NotifyFlowCreate) defaults() {
	if _, ok := nfc.mutation.UpdatedAt(); !ok {
		v := notifyflow.DefaultUpdatedAt()
		nfc.mutation.SetUpdatedAt(v)
	}
	if _, ok := nfc.mutation.CreatedAt(); !ok {
		v := notifyflow.DefaultCreatedAt()
		nfc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nfc *NotifyFlowCreate) check() error {
	if _, ok := nfc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "NotifyFlow.name"`)}
	}
	if _, ok := nfc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "NotifyFlow.description"`)}
	}
	if _, ok := nfc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "NotifyFlow.status"`)}
	}
	if v, ok := nfc.mutation.Status(); ok {
		if err := notifyflow.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "NotifyFlow.status": %w`, err)}
		}
	}
	if _, ok := nfc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "NotifyFlow.updated_at"`)}
	}
	if _, ok := nfc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "NotifyFlow.created_at"`)}
	}
	if _, ok := nfc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "NotifyFlow.owner"`)}
	}
	return nil
}

func (nfc *NotifyFlowCreate) sqlSave(ctx context.Context) (*NotifyFlow, error) {
	if err := nfc.check(); err != nil {
		return nil, err
	}
	_node, _spec := nfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nfc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = model.InternalID(id)
	}
	nfc.mutation.id = &_node.ID
	nfc.mutation.done = true
	return _node, nil
}

func (nfc *NotifyFlowCreate) createSpec() (*NotifyFlow, *sqlgraph.CreateSpec) {
	var (
		_node = &NotifyFlow{config: nfc.config}
		_spec = sqlgraph.NewCreateSpec(notifyflow.Table, sqlgraph.NewFieldSpec(notifyflow.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = nfc.conflict
	if id, ok := nfc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := nfc.mutation.Name(); ok {
		_spec.SetField(notifyflow.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := nfc.mutation.Description(); ok {
		_spec.SetField(notifyflow.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := nfc.mutation.Status(); ok {
		_spec.SetField(notifyflow.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := nfc.mutation.UpdatedAt(); ok {
		_spec.SetField(notifyflow.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := nfc.mutation.CreatedAt(); ok {
		_spec.SetField(notifyflow.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := nfc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notifyflow.OwnerTable,
			Columns: []string{notifyflow.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_notify_flow = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nfc.mutation.NotifyTargetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   notifyflow.NotifyTargetTable,
			Columns: notifyflow.NotifyTargetPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifytarget.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &NotifyFlowTargetCreate{config: nfc.config, mutation: newNotifyFlowTargetMutation(nfc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nfc.mutation.FeedConfigIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notifyflow.FeedConfigTable,
			Columns: notifyflow.FeedConfigPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feedconfig.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &NotifyFlowSourceCreate{config: nfc.config, mutation: newNotifyFlowSourceMutation(nfc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nfc.mutation.NotifyFlowTargetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   notifyflow.NotifyFlowTargetTable,
			Columns: []string{notifyflow.NotifyFlowTargetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflowtarget.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nfc.mutation.NotifyFlowSourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   notifyflow.NotifyFlowSourceTable,
			Columns: []string{notifyflow.NotifyFlowSourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflowsource.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.NotifyFlow.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NotifyFlowUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (nfc *NotifyFlowCreate) OnConflict(opts ...sql.ConflictOption) *NotifyFlowUpsertOne {
	nfc.conflict = opts
	return &NotifyFlowUpsertOne{
		create: nfc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.NotifyFlow.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (nfc *NotifyFlowCreate) OnConflictColumns(columns ...string) *NotifyFlowUpsertOne {
	nfc.conflict = append(nfc.conflict, sql.ConflictColumns(columns...))
	return &NotifyFlowUpsertOne{
		create: nfc,
	}
}

type (
	// NotifyFlowUpsertOne is the builder for "upsert"-ing
	//  one NotifyFlow node.
	NotifyFlowUpsertOne struct {
		create *NotifyFlowCreate
	}

	// NotifyFlowUpsert is the "OnConflict" setter.
	NotifyFlowUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *NotifyFlowUpsert) SetName(v string) *NotifyFlowUpsert {
	u.Set(notifyflow.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *NotifyFlowUpsert) UpdateName() *NotifyFlowUpsert {
	u.SetExcluded(notifyflow.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *NotifyFlowUpsert) SetDescription(v string) *NotifyFlowUpsert {
	u.Set(notifyflow.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *NotifyFlowUpsert) UpdateDescription() *NotifyFlowUpsert {
	u.SetExcluded(notifyflow.FieldDescription)
	return u
}

// SetStatus sets the "status" field.
func (u *NotifyFlowUpsert) SetStatus(v notifyflow.Status) *NotifyFlowUpsert {
	u.Set(notifyflow.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *NotifyFlowUpsert) UpdateStatus() *NotifyFlowUpsert {
	u.SetExcluded(notifyflow.FieldStatus)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NotifyFlowUpsert) SetUpdatedAt(v time.Time) *NotifyFlowUpsert {
	u.Set(notifyflow.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NotifyFlowUpsert) UpdateUpdatedAt() *NotifyFlowUpsert {
	u.SetExcluded(notifyflow.FieldUpdatedAt)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *NotifyFlowUpsert) SetCreatedAt(v time.Time) *NotifyFlowUpsert {
	u.Set(notifyflow.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *NotifyFlowUpsert) UpdateCreatedAt() *NotifyFlowUpsert {
	u.SetExcluded(notifyflow.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.NotifyFlow.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(notifyflow.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *NotifyFlowUpsertOne) UpdateNewValues() *NotifyFlowUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(notifyflow.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.NotifyFlow.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *NotifyFlowUpsertOne) Ignore() *NotifyFlowUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NotifyFlowUpsertOne) DoNothing() *NotifyFlowUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NotifyFlowCreate.OnConflict
// documentation for more info.
func (u *NotifyFlowUpsertOne) Update(set func(*NotifyFlowUpsert)) *NotifyFlowUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NotifyFlowUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *NotifyFlowUpsertOne) SetName(v string) *NotifyFlowUpsertOne {
	return u.Update(func(s *NotifyFlowUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *NotifyFlowUpsertOne) UpdateName() *NotifyFlowUpsertOne {
	return u.Update(func(s *NotifyFlowUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *NotifyFlowUpsertOne) SetDescription(v string) *NotifyFlowUpsertOne {
	return u.Update(func(s *NotifyFlowUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *NotifyFlowUpsertOne) UpdateDescription() *NotifyFlowUpsertOne {
	return u.Update(func(s *NotifyFlowUpsert) {
		s.UpdateDescription()
	})
}

// SetStatus sets the "status" field.
func (u *NotifyFlowUpsertOne) SetStatus(v notifyflow.Status) *NotifyFlowUpsertOne {
	return u.Update(func(s *NotifyFlowUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *NotifyFlowUpsertOne) UpdateStatus() *NotifyFlowUpsertOne {
	return u.Update(func(s *NotifyFlowUpsert) {
		s.UpdateStatus()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NotifyFlowUpsertOne) SetUpdatedAt(v time.Time) *NotifyFlowUpsertOne {
	return u.Update(func(s *NotifyFlowUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NotifyFlowUpsertOne) UpdateUpdatedAt() *NotifyFlowUpsertOne {
	return u.Update(func(s *NotifyFlowUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *NotifyFlowUpsertOne) SetCreatedAt(v time.Time) *NotifyFlowUpsertOne {
	return u.Update(func(s *NotifyFlowUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *NotifyFlowUpsertOne) UpdateCreatedAt() *NotifyFlowUpsertOne {
	return u.Update(func(s *NotifyFlowUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *NotifyFlowUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NotifyFlowCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NotifyFlowUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *NotifyFlowUpsertOne) ID(ctx context.Context) (id model.InternalID, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *NotifyFlowUpsertOne) IDX(ctx context.Context) model.InternalID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// NotifyFlowCreateBulk is the builder for creating many NotifyFlow entities in bulk.
type NotifyFlowCreateBulk struct {
	config
	err      error
	builders []*NotifyFlowCreate
	conflict []sql.ConflictOption
}

// Save creates the NotifyFlow entities in the database.
func (nfcb *NotifyFlowCreateBulk) Save(ctx context.Context) ([]*NotifyFlow, error) {
	if nfcb.err != nil {
		return nil, nfcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(nfcb.builders))
	nodes := make([]*NotifyFlow, len(nfcb.builders))
	mutators := make([]Mutator, len(nfcb.builders))
	for i := range nfcb.builders {
		func(i int, root context.Context) {
			builder := nfcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NotifyFlowMutation)
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
					_, err = mutators[i+1].Mutate(root, nfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = nfcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, nfcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = model.InternalID(id)
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
		if _, err := mutators[0].Mutate(ctx, nfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (nfcb *NotifyFlowCreateBulk) SaveX(ctx context.Context) []*NotifyFlow {
	v, err := nfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nfcb *NotifyFlowCreateBulk) Exec(ctx context.Context) error {
	_, err := nfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nfcb *NotifyFlowCreateBulk) ExecX(ctx context.Context) {
	if err := nfcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.NotifyFlow.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NotifyFlowUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (nfcb *NotifyFlowCreateBulk) OnConflict(opts ...sql.ConflictOption) *NotifyFlowUpsertBulk {
	nfcb.conflict = opts
	return &NotifyFlowUpsertBulk{
		create: nfcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.NotifyFlow.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (nfcb *NotifyFlowCreateBulk) OnConflictColumns(columns ...string) *NotifyFlowUpsertBulk {
	nfcb.conflict = append(nfcb.conflict, sql.ConflictColumns(columns...))
	return &NotifyFlowUpsertBulk{
		create: nfcb,
	}
}

// NotifyFlowUpsertBulk is the builder for "upsert"-ing
// a bulk of NotifyFlow nodes.
type NotifyFlowUpsertBulk struct {
	create *NotifyFlowCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.NotifyFlow.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(notifyflow.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *NotifyFlowUpsertBulk) UpdateNewValues() *NotifyFlowUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(notifyflow.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.NotifyFlow.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *NotifyFlowUpsertBulk) Ignore() *NotifyFlowUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NotifyFlowUpsertBulk) DoNothing() *NotifyFlowUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NotifyFlowCreateBulk.OnConflict
// documentation for more info.
func (u *NotifyFlowUpsertBulk) Update(set func(*NotifyFlowUpsert)) *NotifyFlowUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NotifyFlowUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *NotifyFlowUpsertBulk) SetName(v string) *NotifyFlowUpsertBulk {
	return u.Update(func(s *NotifyFlowUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *NotifyFlowUpsertBulk) UpdateName() *NotifyFlowUpsertBulk {
	return u.Update(func(s *NotifyFlowUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *NotifyFlowUpsertBulk) SetDescription(v string) *NotifyFlowUpsertBulk {
	return u.Update(func(s *NotifyFlowUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *NotifyFlowUpsertBulk) UpdateDescription() *NotifyFlowUpsertBulk {
	return u.Update(func(s *NotifyFlowUpsert) {
		s.UpdateDescription()
	})
}

// SetStatus sets the "status" field.
func (u *NotifyFlowUpsertBulk) SetStatus(v notifyflow.Status) *NotifyFlowUpsertBulk {
	return u.Update(func(s *NotifyFlowUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *NotifyFlowUpsertBulk) UpdateStatus() *NotifyFlowUpsertBulk {
	return u.Update(func(s *NotifyFlowUpsert) {
		s.UpdateStatus()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NotifyFlowUpsertBulk) SetUpdatedAt(v time.Time) *NotifyFlowUpsertBulk {
	return u.Update(func(s *NotifyFlowUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NotifyFlowUpsertBulk) UpdateUpdatedAt() *NotifyFlowUpsertBulk {
	return u.Update(func(s *NotifyFlowUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *NotifyFlowUpsertBulk) SetCreatedAt(v time.Time) *NotifyFlowUpsertBulk {
	return u.Update(func(s *NotifyFlowUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *NotifyFlowUpsertBulk) UpdateCreatedAt() *NotifyFlowUpsertBulk {
	return u.Update(func(s *NotifyFlowUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *NotifyFlowUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the NotifyFlowCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NotifyFlowCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NotifyFlowUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
