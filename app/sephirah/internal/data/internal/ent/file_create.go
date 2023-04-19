// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/file"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/image"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
)

// FileCreate is the builder for creating a File entity.
type FileCreate struct {
	config
	mutation *FileMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (fc *FileCreate) SetName(s string) *FileCreate {
	fc.mutation.SetName(s)
	return fc
}

// SetSize sets the "size" field.
func (fc *FileCreate) SetSize(i int64) *FileCreate {
	fc.mutation.SetSize(i)
	return fc
}

// SetType sets the "type" field.
func (fc *FileCreate) SetType(f file.Type) *FileCreate {
	fc.mutation.SetType(f)
	return fc
}

// SetSha256 sets the "sha256" field.
func (fc *FileCreate) SetSha256(b []byte) *FileCreate {
	fc.mutation.SetSha256(b)
	return fc
}

// SetUpdatedAt sets the "updated_at" field.
func (fc *FileCreate) SetUpdatedAt(t time.Time) *FileCreate {
	fc.mutation.SetUpdatedAt(t)
	return fc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fc *FileCreate) SetNillableUpdatedAt(t *time.Time) *FileCreate {
	if t != nil {
		fc.SetUpdatedAt(*t)
	}
	return fc
}

// SetCreatedAt sets the "created_at" field.
func (fc *FileCreate) SetCreatedAt(t time.Time) *FileCreate {
	fc.mutation.SetCreatedAt(t)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FileCreate) SetNillableCreatedAt(t *time.Time) *FileCreate {
	if t != nil {
		fc.SetCreatedAt(*t)
	}
	return fc
}

// SetID sets the "id" field.
func (fc *FileCreate) SetID(mi model.InternalID) *FileCreate {
	fc.mutation.SetID(mi)
	return fc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (fc *FileCreate) SetOwnerID(id model.InternalID) *FileCreate {
	fc.mutation.SetOwnerID(id)
	return fc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (fc *FileCreate) SetNillableOwnerID(id *model.InternalID) *FileCreate {
	if id != nil {
		fc = fc.SetOwnerID(*id)
	}
	return fc
}

// SetOwner sets the "owner" edge to the User entity.
func (fc *FileCreate) SetOwner(u *User) *FileCreate {
	return fc.SetOwnerID(u.ID)
}

// SetImageID sets the "image" edge to the Image entity by ID.
func (fc *FileCreate) SetImageID(id model.InternalID) *FileCreate {
	fc.mutation.SetImageID(id)
	return fc
}

// SetNillableImageID sets the "image" edge to the Image entity by ID if the given value is not nil.
func (fc *FileCreate) SetNillableImageID(id *model.InternalID) *FileCreate {
	if id != nil {
		fc = fc.SetImageID(*id)
	}
	return fc
}

// SetImage sets the "image" edge to the Image entity.
func (fc *FileCreate) SetImage(i *Image) *FileCreate {
	return fc.SetImageID(i.ID)
}

// Mutation returns the FileMutation object of the builder.
func (fc *FileCreate) Mutation() *FileMutation {
	return fc.mutation
}

// Save creates the File in the database.
func (fc *FileCreate) Save(ctx context.Context) (*File, error) {
	fc.defaults()
	return withHooks[*File, FileMutation](ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FileCreate) SaveX(ctx context.Context) *File {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FileCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FileCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FileCreate) defaults() {
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		v := file.DefaultUpdatedAt()
		fc.mutation.SetUpdatedAt(v)
	}
	if _, ok := fc.mutation.CreatedAt(); !ok {
		v := file.DefaultCreatedAt()
		fc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FileCreate) check() error {
	if _, ok := fc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "File.name"`)}
	}
	if _, ok := fc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "File.size"`)}
	}
	if _, ok := fc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "File.type"`)}
	}
	if v, ok := fc.mutation.GetType(); ok {
		if err := file.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "File.type": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Sha256(); !ok {
		return &ValidationError{Name: "sha256", err: errors.New(`ent: missing required field "File.sha256"`)}
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "File.updated_at"`)}
	}
	if _, ok := fc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "File.created_at"`)}
	}
	return nil
}

func (fc *FileCreate) sqlSave(ctx context.Context) (*File, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = model.InternalID(id)
	}
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FileCreate) createSpec() (*File, *sqlgraph.CreateSpec) {
	var (
		_node = &File{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(file.Table, sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = fc.conflict
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fc.mutation.Name(); ok {
		_spec.SetField(file.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := fc.mutation.Size(); ok {
		_spec.SetField(file.FieldSize, field.TypeInt64, value)
		_node.Size = value
	}
	if value, ok := fc.mutation.GetType(); ok {
		_spec.SetField(file.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := fc.mutation.Sha256(); ok {
		_spec.SetField(file.FieldSha256, field.TypeBytes, value)
		_node.Sha256 = value
	}
	if value, ok := fc.mutation.UpdatedAt(); ok {
		_spec.SetField(file.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.SetField(file.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := fc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.OwnerTable,
			Columns: []string{file.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_file = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.ImageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   file.ImageTable,
			Columns: []string{file.ImageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt64),
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
//	client.File.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FileUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (fc *FileCreate) OnConflict(opts ...sql.ConflictOption) *FileUpsertOne {
	fc.conflict = opts
	return &FileUpsertOne{
		create: fc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fc *FileCreate) OnConflictColumns(columns ...string) *FileUpsertOne {
	fc.conflict = append(fc.conflict, sql.ConflictColumns(columns...))
	return &FileUpsertOne{
		create: fc,
	}
}

type (
	// FileUpsertOne is the builder for "upsert"-ing
	//  one File node.
	FileUpsertOne struct {
		create *FileCreate
	}

	// FileUpsert is the "OnConflict" setter.
	FileUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *FileUpsert) SetName(v string) *FileUpsert {
	u.Set(file.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FileUpsert) UpdateName() *FileUpsert {
	u.SetExcluded(file.FieldName)
	return u
}

// SetSize sets the "size" field.
func (u *FileUpsert) SetSize(v int64) *FileUpsert {
	u.Set(file.FieldSize, v)
	return u
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *FileUpsert) UpdateSize() *FileUpsert {
	u.SetExcluded(file.FieldSize)
	return u
}

// AddSize adds v to the "size" field.
func (u *FileUpsert) AddSize(v int64) *FileUpsert {
	u.Add(file.FieldSize, v)
	return u
}

// SetType sets the "type" field.
func (u *FileUpsert) SetType(v file.Type) *FileUpsert {
	u.Set(file.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *FileUpsert) UpdateType() *FileUpsert {
	u.SetExcluded(file.FieldType)
	return u
}

// SetSha256 sets the "sha256" field.
func (u *FileUpsert) SetSha256(v []byte) *FileUpsert {
	u.Set(file.FieldSha256, v)
	return u
}

// UpdateSha256 sets the "sha256" field to the value that was provided on create.
func (u *FileUpsert) UpdateSha256() *FileUpsert {
	u.SetExcluded(file.FieldSha256)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FileUpsert) SetUpdatedAt(v time.Time) *FileUpsert {
	u.Set(file.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FileUpsert) UpdateUpdatedAt() *FileUpsert {
	u.SetExcluded(file.FieldUpdatedAt)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *FileUpsert) SetCreatedAt(v time.Time) *FileUpsert {
	u.Set(file.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FileUpsert) UpdateCreatedAt() *FileUpsert {
	u.SetExcluded(file.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(file.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FileUpsertOne) UpdateNewValues() *FileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(file.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.File.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FileUpsertOne) Ignore() *FileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FileUpsertOne) DoNothing() *FileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FileCreate.OnConflict
// documentation for more info.
func (u *FileUpsertOne) Update(set func(*FileUpsert)) *FileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FileUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *FileUpsertOne) SetName(v string) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateName() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateName()
	})
}

// SetSize sets the "size" field.
func (u *FileUpsertOne) SetSize(v int64) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetSize(v)
	})
}

// AddSize adds v to the "size" field.
func (u *FileUpsertOne) AddSize(v int64) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.AddSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateSize() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateSize()
	})
}

// SetType sets the "type" field.
func (u *FileUpsertOne) SetType(v file.Type) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateType() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateType()
	})
}

// SetSha256 sets the "sha256" field.
func (u *FileUpsertOne) SetSha256(v []byte) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetSha256(v)
	})
}

// UpdateSha256 sets the "sha256" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateSha256() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateSha256()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FileUpsertOne) SetUpdatedAt(v time.Time) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateUpdatedAt() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *FileUpsertOne) SetCreatedAt(v time.Time) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateCreatedAt() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *FileUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FileCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FileUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FileUpsertOne) ID(ctx context.Context) (id model.InternalID, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FileUpsertOne) IDX(ctx context.Context) model.InternalID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FileCreateBulk is the builder for creating many File entities in bulk.
type FileCreateBulk struct {
	config
	builders []*FileCreate
	conflict []sql.ConflictOption
}

// Save creates the File entities in the database.
func (fcb *FileCreateBulk) Save(ctx context.Context) ([]*File, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*File, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FileMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = fcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FileCreateBulk) SaveX(ctx context.Context) []*File {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FileCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FileCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.File.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FileUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (fcb *FileCreateBulk) OnConflict(opts ...sql.ConflictOption) *FileUpsertBulk {
	fcb.conflict = opts
	return &FileUpsertBulk{
		create: fcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fcb *FileCreateBulk) OnConflictColumns(columns ...string) *FileUpsertBulk {
	fcb.conflict = append(fcb.conflict, sql.ConflictColumns(columns...))
	return &FileUpsertBulk{
		create: fcb,
	}
}

// FileUpsertBulk is the builder for "upsert"-ing
// a bulk of File nodes.
type FileUpsertBulk struct {
	create *FileCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(file.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FileUpsertBulk) UpdateNewValues() *FileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(file.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FileUpsertBulk) Ignore() *FileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FileUpsertBulk) DoNothing() *FileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FileCreateBulk.OnConflict
// documentation for more info.
func (u *FileUpsertBulk) Update(set func(*FileUpsert)) *FileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FileUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *FileUpsertBulk) SetName(v string) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateName() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateName()
	})
}

// SetSize sets the "size" field.
func (u *FileUpsertBulk) SetSize(v int64) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetSize(v)
	})
}

// AddSize adds v to the "size" field.
func (u *FileUpsertBulk) AddSize(v int64) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.AddSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateSize() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateSize()
	})
}

// SetType sets the "type" field.
func (u *FileUpsertBulk) SetType(v file.Type) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateType() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateType()
	})
}

// SetSha256 sets the "sha256" field.
func (u *FileUpsertBulk) SetSha256(v []byte) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetSha256(v)
	})
}

// UpdateSha256 sets the "sha256" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateSha256() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateSha256()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FileUpsertBulk) SetUpdatedAt(v time.Time) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateUpdatedAt() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *FileUpsertBulk) SetCreatedAt(v time.Time) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateCreatedAt() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *FileUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FileCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FileCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FileUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
