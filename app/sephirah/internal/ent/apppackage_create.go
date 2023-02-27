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
	"github.com/tuihub/librarian/app/sephirah/internal/ent/app"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/apppackage"
)

// AppPackageCreate is the builder for creating a AppPackage entity.
type AppPackageCreate struct {
	config
	mutation *AppPackageMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetSource sets the "source" field.
func (apc *AppPackageCreate) SetSource(a apppackage.Source) *AppPackageCreate {
	apc.mutation.SetSource(a)
	return apc
}

// SetSourceID sets the "source_id" field.
func (apc *AppPackageCreate) SetSourceID(i int64) *AppPackageCreate {
	apc.mutation.SetSourceID(i)
	return apc
}

// SetSourcePackageID sets the "source_package_id" field.
func (apc *AppPackageCreate) SetSourcePackageID(s string) *AppPackageCreate {
	apc.mutation.SetSourcePackageID(s)
	return apc
}

// SetName sets the "name" field.
func (apc *AppPackageCreate) SetName(s string) *AppPackageCreate {
	apc.mutation.SetName(s)
	return apc
}

// SetDescription sets the "description" field.
func (apc *AppPackageCreate) SetDescription(s string) *AppPackageCreate {
	apc.mutation.SetDescription(s)
	return apc
}

// SetBinaryName sets the "binary_name" field.
func (apc *AppPackageCreate) SetBinaryName(s string) *AppPackageCreate {
	apc.mutation.SetBinaryName(s)
	return apc
}

// SetBinarySize sets the "binary_size" field.
func (apc *AppPackageCreate) SetBinarySize(i int64) *AppPackageCreate {
	apc.mutation.SetBinarySize(i)
	return apc
}

// SetBinaryPublicURL sets the "binary_public_url" field.
func (apc *AppPackageCreate) SetBinaryPublicURL(s string) *AppPackageCreate {
	apc.mutation.SetBinaryPublicURL(s)
	return apc
}

// SetUpdatedAt sets the "updated_at" field.
func (apc *AppPackageCreate) SetUpdatedAt(t time.Time) *AppPackageCreate {
	apc.mutation.SetUpdatedAt(t)
	return apc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (apc *AppPackageCreate) SetNillableUpdatedAt(t *time.Time) *AppPackageCreate {
	if t != nil {
		apc.SetUpdatedAt(*t)
	}
	return apc
}

// SetCreatedAt sets the "created_at" field.
func (apc *AppPackageCreate) SetCreatedAt(t time.Time) *AppPackageCreate {
	apc.mutation.SetCreatedAt(t)
	return apc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (apc *AppPackageCreate) SetNillableCreatedAt(t *time.Time) *AppPackageCreate {
	if t != nil {
		apc.SetCreatedAt(*t)
	}
	return apc
}

// SetID sets the "id" field.
func (apc *AppPackageCreate) SetID(i int64) *AppPackageCreate {
	apc.mutation.SetID(i)
	return apc
}

// SetAppID sets the "app" edge to the App entity by ID.
func (apc *AppPackageCreate) SetAppID(id int64) *AppPackageCreate {
	apc.mutation.SetAppID(id)
	return apc
}

// SetNillableAppID sets the "app" edge to the App entity by ID if the given value is not nil.
func (apc *AppPackageCreate) SetNillableAppID(id *int64) *AppPackageCreate {
	if id != nil {
		apc = apc.SetAppID(*id)
	}
	return apc
}

// SetApp sets the "app" edge to the App entity.
func (apc *AppPackageCreate) SetApp(a *App) *AppPackageCreate {
	return apc.SetAppID(a.ID)
}

// Mutation returns the AppPackageMutation object of the builder.
func (apc *AppPackageCreate) Mutation() *AppPackageMutation {
	return apc.mutation
}

// Save creates the AppPackage in the database.
func (apc *AppPackageCreate) Save(ctx context.Context) (*AppPackage, error) {
	apc.defaults()
	return withHooks[*AppPackage, AppPackageMutation](ctx, apc.sqlSave, apc.mutation, apc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (apc *AppPackageCreate) SaveX(ctx context.Context) *AppPackage {
	v, err := apc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (apc *AppPackageCreate) Exec(ctx context.Context) error {
	_, err := apc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apc *AppPackageCreate) ExecX(ctx context.Context) {
	if err := apc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (apc *AppPackageCreate) defaults() {
	if _, ok := apc.mutation.UpdatedAt(); !ok {
		v := apppackage.DefaultUpdatedAt()
		apc.mutation.SetUpdatedAt(v)
	}
	if _, ok := apc.mutation.CreatedAt(); !ok {
		v := apppackage.DefaultCreatedAt()
		apc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (apc *AppPackageCreate) check() error {
	if _, ok := apc.mutation.Source(); !ok {
		return &ValidationError{Name: "source", err: errors.New(`ent: missing required field "AppPackage.source"`)}
	}
	if v, ok := apc.mutation.Source(); ok {
		if err := apppackage.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf(`ent: validator failed for field "AppPackage.source": %w`, err)}
		}
	}
	if _, ok := apc.mutation.SourceID(); !ok {
		return &ValidationError{Name: "source_id", err: errors.New(`ent: missing required field "AppPackage.source_id"`)}
	}
	if _, ok := apc.mutation.SourcePackageID(); !ok {
		return &ValidationError{Name: "source_package_id", err: errors.New(`ent: missing required field "AppPackage.source_package_id"`)}
	}
	if _, ok := apc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "AppPackage.name"`)}
	}
	if _, ok := apc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "AppPackage.description"`)}
	}
	if _, ok := apc.mutation.BinaryName(); !ok {
		return &ValidationError{Name: "binary_name", err: errors.New(`ent: missing required field "AppPackage.binary_name"`)}
	}
	if _, ok := apc.mutation.BinarySize(); !ok {
		return &ValidationError{Name: "binary_size", err: errors.New(`ent: missing required field "AppPackage.binary_size"`)}
	}
	if _, ok := apc.mutation.BinaryPublicURL(); !ok {
		return &ValidationError{Name: "binary_public_url", err: errors.New(`ent: missing required field "AppPackage.binary_public_url"`)}
	}
	if _, ok := apc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AppPackage.updated_at"`)}
	}
	if _, ok := apc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppPackage.created_at"`)}
	}
	return nil
}

func (apc *AppPackageCreate) sqlSave(ctx context.Context) (*AppPackage, error) {
	if err := apc.check(); err != nil {
		return nil, err
	}
	_node, _spec := apc.createSpec()
	if err := sqlgraph.CreateNode(ctx, apc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	apc.mutation.id = &_node.ID
	apc.mutation.done = true
	return _node, nil
}

func (apc *AppPackageCreate) createSpec() (*AppPackage, *sqlgraph.CreateSpec) {
	var (
		_node = &AppPackage{config: apc.config}
		_spec = sqlgraph.NewCreateSpec(apppackage.Table, sqlgraph.NewFieldSpec(apppackage.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = apc.conflict
	if id, ok := apc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := apc.mutation.Source(); ok {
		_spec.SetField(apppackage.FieldSource, field.TypeEnum, value)
		_node.Source = value
	}
	if value, ok := apc.mutation.SourceID(); ok {
		_spec.SetField(apppackage.FieldSourceID, field.TypeInt64, value)
		_node.SourceID = value
	}
	if value, ok := apc.mutation.SourcePackageID(); ok {
		_spec.SetField(apppackage.FieldSourcePackageID, field.TypeString, value)
		_node.SourcePackageID = value
	}
	if value, ok := apc.mutation.Name(); ok {
		_spec.SetField(apppackage.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := apc.mutation.Description(); ok {
		_spec.SetField(apppackage.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := apc.mutation.BinaryName(); ok {
		_spec.SetField(apppackage.FieldBinaryName, field.TypeString, value)
		_node.BinaryName = value
	}
	if value, ok := apc.mutation.BinarySize(); ok {
		_spec.SetField(apppackage.FieldBinarySize, field.TypeInt64, value)
		_node.BinarySize = value
	}
	if value, ok := apc.mutation.BinaryPublicURL(); ok {
		_spec.SetField(apppackage.FieldBinaryPublicURL, field.TypeString, value)
		_node.BinaryPublicURL = value
	}
	if value, ok := apc.mutation.UpdatedAt(); ok {
		_spec.SetField(apppackage.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := apc.mutation.CreatedAt(); ok {
		_spec.SetField(apppackage.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := apc.mutation.AppIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apppackage.AppTable,
			Columns: []string{apppackage.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: app.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.app_app_package = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppPackage.Create().
//		SetSource(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppPackageUpsert) {
//			SetSource(v+v).
//		}).
//		Exec(ctx)
func (apc *AppPackageCreate) OnConflict(opts ...sql.ConflictOption) *AppPackageUpsertOne {
	apc.conflict = opts
	return &AppPackageUpsertOne{
		create: apc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppPackage.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (apc *AppPackageCreate) OnConflictColumns(columns ...string) *AppPackageUpsertOne {
	apc.conflict = append(apc.conflict, sql.ConflictColumns(columns...))
	return &AppPackageUpsertOne{
		create: apc,
	}
}

type (
	// AppPackageUpsertOne is the builder for "upsert"-ing
	//  one AppPackage node.
	AppPackageUpsertOne struct {
		create *AppPackageCreate
	}

	// AppPackageUpsert is the "OnConflict" setter.
	AppPackageUpsert struct {
		*sql.UpdateSet
	}
)

// SetSource sets the "source" field.
func (u *AppPackageUpsert) SetSource(v apppackage.Source) *AppPackageUpsert {
	u.Set(apppackage.FieldSource, v)
	return u
}

// UpdateSource sets the "source" field to the value that was provided on create.
func (u *AppPackageUpsert) UpdateSource() *AppPackageUpsert {
	u.SetExcluded(apppackage.FieldSource)
	return u
}

// SetSourceID sets the "source_id" field.
func (u *AppPackageUpsert) SetSourceID(v int64) *AppPackageUpsert {
	u.Set(apppackage.FieldSourceID, v)
	return u
}

// UpdateSourceID sets the "source_id" field to the value that was provided on create.
func (u *AppPackageUpsert) UpdateSourceID() *AppPackageUpsert {
	u.SetExcluded(apppackage.FieldSourceID)
	return u
}

// AddSourceID adds v to the "source_id" field.
func (u *AppPackageUpsert) AddSourceID(v int64) *AppPackageUpsert {
	u.Add(apppackage.FieldSourceID, v)
	return u
}

// SetSourcePackageID sets the "source_package_id" field.
func (u *AppPackageUpsert) SetSourcePackageID(v string) *AppPackageUpsert {
	u.Set(apppackage.FieldSourcePackageID, v)
	return u
}

// UpdateSourcePackageID sets the "source_package_id" field to the value that was provided on create.
func (u *AppPackageUpsert) UpdateSourcePackageID() *AppPackageUpsert {
	u.SetExcluded(apppackage.FieldSourcePackageID)
	return u
}

// SetName sets the "name" field.
func (u *AppPackageUpsert) SetName(v string) *AppPackageUpsert {
	u.Set(apppackage.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppPackageUpsert) UpdateName() *AppPackageUpsert {
	u.SetExcluded(apppackage.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *AppPackageUpsert) SetDescription(v string) *AppPackageUpsert {
	u.Set(apppackage.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AppPackageUpsert) UpdateDescription() *AppPackageUpsert {
	u.SetExcluded(apppackage.FieldDescription)
	return u
}

// SetBinaryName sets the "binary_name" field.
func (u *AppPackageUpsert) SetBinaryName(v string) *AppPackageUpsert {
	u.Set(apppackage.FieldBinaryName, v)
	return u
}

// UpdateBinaryName sets the "binary_name" field to the value that was provided on create.
func (u *AppPackageUpsert) UpdateBinaryName() *AppPackageUpsert {
	u.SetExcluded(apppackage.FieldBinaryName)
	return u
}

// SetBinarySize sets the "binary_size" field.
func (u *AppPackageUpsert) SetBinarySize(v int64) *AppPackageUpsert {
	u.Set(apppackage.FieldBinarySize, v)
	return u
}

// UpdateBinarySize sets the "binary_size" field to the value that was provided on create.
func (u *AppPackageUpsert) UpdateBinarySize() *AppPackageUpsert {
	u.SetExcluded(apppackage.FieldBinarySize)
	return u
}

// AddBinarySize adds v to the "binary_size" field.
func (u *AppPackageUpsert) AddBinarySize(v int64) *AppPackageUpsert {
	u.Add(apppackage.FieldBinarySize, v)
	return u
}

// SetBinaryPublicURL sets the "binary_public_url" field.
func (u *AppPackageUpsert) SetBinaryPublicURL(v string) *AppPackageUpsert {
	u.Set(apppackage.FieldBinaryPublicURL, v)
	return u
}

// UpdateBinaryPublicURL sets the "binary_public_url" field to the value that was provided on create.
func (u *AppPackageUpsert) UpdateBinaryPublicURL() *AppPackageUpsert {
	u.SetExcluded(apppackage.FieldBinaryPublicURL)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppPackageUpsert) SetUpdatedAt(v time.Time) *AppPackageUpsert {
	u.Set(apppackage.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppPackageUpsert) UpdateUpdatedAt() *AppPackageUpsert {
	u.SetExcluded(apppackage.FieldUpdatedAt)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppPackageUpsert) SetCreatedAt(v time.Time) *AppPackageUpsert {
	u.Set(apppackage.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppPackageUpsert) UpdateCreatedAt() *AppPackageUpsert {
	u.SetExcluded(apppackage.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppPackage.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(apppackage.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AppPackageUpsertOne) UpdateNewValues() *AppPackageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(apppackage.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppPackage.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AppPackageUpsertOne) Ignore() *AppPackageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppPackageUpsertOne) DoNothing() *AppPackageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppPackageCreate.OnConflict
// documentation for more info.
func (u *AppPackageUpsertOne) Update(set func(*AppPackageUpsert)) *AppPackageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppPackageUpsert{UpdateSet: update})
	}))
	return u
}

// SetSource sets the "source" field.
func (u *AppPackageUpsertOne) SetSource(v apppackage.Source) *AppPackageUpsertOne {
	return u.Update(func(s *AppPackageUpsert) {
		s.SetSource(v)
	})
}

// UpdateSource sets the "source" field to the value that was provided on create.
func (u *AppPackageUpsertOne) UpdateSource() *AppPackageUpsertOne {
	return u.Update(func(s *AppPackageUpsert) {
		s.UpdateSource()
	})
}

// SetSourceID sets the "source_id" field.
func (u *AppPackageUpsertOne) SetSourceID(v int64) *AppPackageUpsertOne {
	return u.Update(func(s *AppPackageUpsert) {
		s.SetSourceID(v)
	})
}

// AddSourceID adds v to the "source_id" field.
func (u *AppPackageUpsertOne) AddSourceID(v int64) *AppPackageUpsertOne {
	return u.Update(func(s *AppPackageUpsert) {
		s.AddSourceID(v)
	})
}

// UpdateSourceID sets the "source_id" field to the value that was provided on create.
func (u *AppPackageUpsertOne) UpdateSourceID() *AppPackageUpsertOne {
	return u.Update(func(s *AppPackageUpsert) {
		s.UpdateSourceID()
	})
}

// SetSourcePackageID sets the "source_package_id" field.
func (u *AppPackageUpsertOne) SetSourcePackageID(v string) *AppPackageUpsertOne {
	return u.Update(func(s *AppPackageUpsert) {
		s.SetSourcePackageID(v)
	})
}

// UpdateSourcePackageID sets the "source_package_id" field to the value that was provided on create.
func (u *AppPackageUpsertOne) UpdateSourcePackageID() *AppPackageUpsertOne {
	return u.Update(func(s *AppPackageUpsert) {
		s.UpdateSourcePackageID()
	})
}

// SetName sets the "name" field.
func (u *AppPackageUpsertOne) SetName(v string) *AppPackageUpsertOne {
	return u.Update(func(s *AppPackageUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppPackageUpsertOne) UpdateName() *AppPackageUpsertOne {
	return u.Update(func(s *AppPackageUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *AppPackageUpsertOne) SetDescription(v string) *AppPackageUpsertOne {
	return u.Update(func(s *AppPackageUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AppPackageUpsertOne) UpdateDescription() *AppPackageUpsertOne {
	return u.Update(func(s *AppPackageUpsert) {
		s.UpdateDescription()
	})
}

// SetBinaryName sets the "binary_name" field.
func (u *AppPackageUpsertOne) SetBinaryName(v string) *AppPackageUpsertOne {
	return u.Update(func(s *AppPackageUpsert) {
		s.SetBinaryName(v)
	})
}

// UpdateBinaryName sets the "binary_name" field to the value that was provided on create.
func (u *AppPackageUpsertOne) UpdateBinaryName() *AppPackageUpsertOne {
	return u.Update(func(s *AppPackageUpsert) {
		s.UpdateBinaryName()
	})
}

// SetBinarySize sets the "binary_size" field.
func (u *AppPackageUpsertOne) SetBinarySize(v int64) *AppPackageUpsertOne {
	return u.Update(func(s *AppPackageUpsert) {
		s.SetBinarySize(v)
	})
}

// AddBinarySize adds v to the "binary_size" field.
func (u *AppPackageUpsertOne) AddBinarySize(v int64) *AppPackageUpsertOne {
	return u.Update(func(s *AppPackageUpsert) {
		s.AddBinarySize(v)
	})
}

// UpdateBinarySize sets the "binary_size" field to the value that was provided on create.
func (u *AppPackageUpsertOne) UpdateBinarySize() *AppPackageUpsertOne {
	return u.Update(func(s *AppPackageUpsert) {
		s.UpdateBinarySize()
	})
}

// SetBinaryPublicURL sets the "binary_public_url" field.
func (u *AppPackageUpsertOne) SetBinaryPublicURL(v string) *AppPackageUpsertOne {
	return u.Update(func(s *AppPackageUpsert) {
		s.SetBinaryPublicURL(v)
	})
}

// UpdateBinaryPublicURL sets the "binary_public_url" field to the value that was provided on create.
func (u *AppPackageUpsertOne) UpdateBinaryPublicURL() *AppPackageUpsertOne {
	return u.Update(func(s *AppPackageUpsert) {
		s.UpdateBinaryPublicURL()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppPackageUpsertOne) SetUpdatedAt(v time.Time) *AppPackageUpsertOne {
	return u.Update(func(s *AppPackageUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppPackageUpsertOne) UpdateUpdatedAt() *AppPackageUpsertOne {
	return u.Update(func(s *AppPackageUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *AppPackageUpsertOne) SetCreatedAt(v time.Time) *AppPackageUpsertOne {
	return u.Update(func(s *AppPackageUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppPackageUpsertOne) UpdateCreatedAt() *AppPackageUpsertOne {
	return u.Update(func(s *AppPackageUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *AppPackageUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppPackageCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppPackageUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppPackageUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppPackageUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppPackageCreateBulk is the builder for creating many AppPackage entities in bulk.
type AppPackageCreateBulk struct {
	config
	builders []*AppPackageCreate
	conflict []sql.ConflictOption
}

// Save creates the AppPackage entities in the database.
func (apcb *AppPackageCreateBulk) Save(ctx context.Context) ([]*AppPackage, error) {
	specs := make([]*sqlgraph.CreateSpec, len(apcb.builders))
	nodes := make([]*AppPackage, len(apcb.builders))
	mutators := make([]Mutator, len(apcb.builders))
	for i := range apcb.builders {
		func(i int, root context.Context) {
			builder := apcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppPackageMutation)
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
					_, err = mutators[i+1].Mutate(root, apcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = apcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, apcb.driver, spec); err != nil {
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
					nodes[i].ID = int64(id)
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
		if _, err := mutators[0].Mutate(ctx, apcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (apcb *AppPackageCreateBulk) SaveX(ctx context.Context) []*AppPackage {
	v, err := apcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (apcb *AppPackageCreateBulk) Exec(ctx context.Context) error {
	_, err := apcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apcb *AppPackageCreateBulk) ExecX(ctx context.Context) {
	if err := apcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppPackage.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppPackageUpsert) {
//			SetSource(v+v).
//		}).
//		Exec(ctx)
func (apcb *AppPackageCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppPackageUpsertBulk {
	apcb.conflict = opts
	return &AppPackageUpsertBulk{
		create: apcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppPackage.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (apcb *AppPackageCreateBulk) OnConflictColumns(columns ...string) *AppPackageUpsertBulk {
	apcb.conflict = append(apcb.conflict, sql.ConflictColumns(columns...))
	return &AppPackageUpsertBulk{
		create: apcb,
	}
}

// AppPackageUpsertBulk is the builder for "upsert"-ing
// a bulk of AppPackage nodes.
type AppPackageUpsertBulk struct {
	create *AppPackageCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppPackage.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(apppackage.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AppPackageUpsertBulk) UpdateNewValues() *AppPackageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(apppackage.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppPackage.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AppPackageUpsertBulk) Ignore() *AppPackageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppPackageUpsertBulk) DoNothing() *AppPackageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppPackageCreateBulk.OnConflict
// documentation for more info.
func (u *AppPackageUpsertBulk) Update(set func(*AppPackageUpsert)) *AppPackageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppPackageUpsert{UpdateSet: update})
	}))
	return u
}

// SetSource sets the "source" field.
func (u *AppPackageUpsertBulk) SetSource(v apppackage.Source) *AppPackageUpsertBulk {
	return u.Update(func(s *AppPackageUpsert) {
		s.SetSource(v)
	})
}

// UpdateSource sets the "source" field to the value that was provided on create.
func (u *AppPackageUpsertBulk) UpdateSource() *AppPackageUpsertBulk {
	return u.Update(func(s *AppPackageUpsert) {
		s.UpdateSource()
	})
}

// SetSourceID sets the "source_id" field.
func (u *AppPackageUpsertBulk) SetSourceID(v int64) *AppPackageUpsertBulk {
	return u.Update(func(s *AppPackageUpsert) {
		s.SetSourceID(v)
	})
}

// AddSourceID adds v to the "source_id" field.
func (u *AppPackageUpsertBulk) AddSourceID(v int64) *AppPackageUpsertBulk {
	return u.Update(func(s *AppPackageUpsert) {
		s.AddSourceID(v)
	})
}

// UpdateSourceID sets the "source_id" field to the value that was provided on create.
func (u *AppPackageUpsertBulk) UpdateSourceID() *AppPackageUpsertBulk {
	return u.Update(func(s *AppPackageUpsert) {
		s.UpdateSourceID()
	})
}

// SetSourcePackageID sets the "source_package_id" field.
func (u *AppPackageUpsertBulk) SetSourcePackageID(v string) *AppPackageUpsertBulk {
	return u.Update(func(s *AppPackageUpsert) {
		s.SetSourcePackageID(v)
	})
}

// UpdateSourcePackageID sets the "source_package_id" field to the value that was provided on create.
func (u *AppPackageUpsertBulk) UpdateSourcePackageID() *AppPackageUpsertBulk {
	return u.Update(func(s *AppPackageUpsert) {
		s.UpdateSourcePackageID()
	})
}

// SetName sets the "name" field.
func (u *AppPackageUpsertBulk) SetName(v string) *AppPackageUpsertBulk {
	return u.Update(func(s *AppPackageUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppPackageUpsertBulk) UpdateName() *AppPackageUpsertBulk {
	return u.Update(func(s *AppPackageUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *AppPackageUpsertBulk) SetDescription(v string) *AppPackageUpsertBulk {
	return u.Update(func(s *AppPackageUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AppPackageUpsertBulk) UpdateDescription() *AppPackageUpsertBulk {
	return u.Update(func(s *AppPackageUpsert) {
		s.UpdateDescription()
	})
}

// SetBinaryName sets the "binary_name" field.
func (u *AppPackageUpsertBulk) SetBinaryName(v string) *AppPackageUpsertBulk {
	return u.Update(func(s *AppPackageUpsert) {
		s.SetBinaryName(v)
	})
}

// UpdateBinaryName sets the "binary_name" field to the value that was provided on create.
func (u *AppPackageUpsertBulk) UpdateBinaryName() *AppPackageUpsertBulk {
	return u.Update(func(s *AppPackageUpsert) {
		s.UpdateBinaryName()
	})
}

// SetBinarySize sets the "binary_size" field.
func (u *AppPackageUpsertBulk) SetBinarySize(v int64) *AppPackageUpsertBulk {
	return u.Update(func(s *AppPackageUpsert) {
		s.SetBinarySize(v)
	})
}

// AddBinarySize adds v to the "binary_size" field.
func (u *AppPackageUpsertBulk) AddBinarySize(v int64) *AppPackageUpsertBulk {
	return u.Update(func(s *AppPackageUpsert) {
		s.AddBinarySize(v)
	})
}

// UpdateBinarySize sets the "binary_size" field to the value that was provided on create.
func (u *AppPackageUpsertBulk) UpdateBinarySize() *AppPackageUpsertBulk {
	return u.Update(func(s *AppPackageUpsert) {
		s.UpdateBinarySize()
	})
}

// SetBinaryPublicURL sets the "binary_public_url" field.
func (u *AppPackageUpsertBulk) SetBinaryPublicURL(v string) *AppPackageUpsertBulk {
	return u.Update(func(s *AppPackageUpsert) {
		s.SetBinaryPublicURL(v)
	})
}

// UpdateBinaryPublicURL sets the "binary_public_url" field to the value that was provided on create.
func (u *AppPackageUpsertBulk) UpdateBinaryPublicURL() *AppPackageUpsertBulk {
	return u.Update(func(s *AppPackageUpsert) {
		s.UpdateBinaryPublicURL()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppPackageUpsertBulk) SetUpdatedAt(v time.Time) *AppPackageUpsertBulk {
	return u.Update(func(s *AppPackageUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppPackageUpsertBulk) UpdateUpdatedAt() *AppPackageUpsertBulk {
	return u.Update(func(s *AppPackageUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *AppPackageUpsertBulk) SetCreatedAt(v time.Time) *AppPackageUpsertBulk {
	return u.Update(func(s *AppPackageUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppPackageUpsertBulk) UpdateCreatedAt() *AppPackageUpsertBulk {
	return u.Update(func(s *AppPackageUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *AppPackageUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppPackageCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppPackageCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppPackageUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
