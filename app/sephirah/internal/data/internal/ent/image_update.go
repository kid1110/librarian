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
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
)

// ImageUpdate is the builder for updating Image entities.
type ImageUpdate struct {
	config
	hooks    []Hook
	mutation *ImageMutation
}

// Where appends a list predicates to the ImageUpdate builder.
func (iu *ImageUpdate) Where(ps ...predicate.Image) *ImageUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetName sets the "name" field.
func (iu *ImageUpdate) SetName(s string) *ImageUpdate {
	iu.mutation.SetName(s)
	return iu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (iu *ImageUpdate) SetNillableName(s *string) *ImageUpdate {
	if s != nil {
		iu.SetName(*s)
	}
	return iu
}

// SetDescription sets the "description" field.
func (iu *ImageUpdate) SetDescription(s string) *ImageUpdate {
	iu.mutation.SetDescription(s)
	return iu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (iu *ImageUpdate) SetNillableDescription(s *string) *ImageUpdate {
	if s != nil {
		iu.SetDescription(*s)
	}
	return iu
}

// SetStatus sets the "status" field.
func (iu *ImageUpdate) SetStatus(i image.Status) *ImageUpdate {
	iu.mutation.SetStatus(i)
	return iu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (iu *ImageUpdate) SetNillableStatus(i *image.Status) *ImageUpdate {
	if i != nil {
		iu.SetStatus(*i)
	}
	return iu
}

// SetUpdatedAt sets the "updated_at" field.
func (iu *ImageUpdate) SetUpdatedAt(t time.Time) *ImageUpdate {
	iu.mutation.SetUpdatedAt(t)
	return iu
}

// SetCreatedAt sets the "created_at" field.
func (iu *ImageUpdate) SetCreatedAt(t time.Time) *ImageUpdate {
	iu.mutation.SetCreatedAt(t)
	return iu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (iu *ImageUpdate) SetNillableCreatedAt(t *time.Time) *ImageUpdate {
	if t != nil {
		iu.SetCreatedAt(*t)
	}
	return iu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (iu *ImageUpdate) SetOwnerID(id model.InternalID) *ImageUpdate {
	iu.mutation.SetOwnerID(id)
	return iu
}

// SetOwner sets the "owner" edge to the User entity.
func (iu *ImageUpdate) SetOwner(u *User) *ImageUpdate {
	return iu.SetOwnerID(u.ID)
}

// SetFileID sets the "file" edge to the File entity by ID.
func (iu *ImageUpdate) SetFileID(id model.InternalID) *ImageUpdate {
	iu.mutation.SetFileID(id)
	return iu
}

// SetNillableFileID sets the "file" edge to the File entity by ID if the given value is not nil.
func (iu *ImageUpdate) SetNillableFileID(id *model.InternalID) *ImageUpdate {
	if id != nil {
		iu = iu.SetFileID(*id)
	}
	return iu
}

// SetFile sets the "file" edge to the File entity.
func (iu *ImageUpdate) SetFile(f *File) *ImageUpdate {
	return iu.SetFileID(f.ID)
}

// Mutation returns the ImageMutation object of the builder.
func (iu *ImageUpdate) Mutation() *ImageMutation {
	return iu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (iu *ImageUpdate) ClearOwner() *ImageUpdate {
	iu.mutation.ClearOwner()
	return iu
}

// ClearFile clears the "file" edge to the File entity.
func (iu *ImageUpdate) ClearFile() *ImageUpdate {
	iu.mutation.ClearFile()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *ImageUpdate) Save(ctx context.Context) (int, error) {
	iu.defaults()
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ImageUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ImageUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ImageUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *ImageUpdate) defaults() {
	if _, ok := iu.mutation.UpdatedAt(); !ok {
		v := image.UpdateDefaultUpdatedAt()
		iu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *ImageUpdate) check() error {
	if v, ok := iu.mutation.Status(); ok {
		if err := image.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Image.status": %w`, err)}
		}
	}
	if iu.mutation.OwnerCleared() && len(iu.mutation.OwnerIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Image.owner"`)
	}
	return nil
}

func (iu *ImageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(image.Table, image.Columns, sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt64))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Name(); ok {
		_spec.SetField(image.FieldName, field.TypeString, value)
	}
	if value, ok := iu.mutation.Description(); ok {
		_spec.SetField(image.FieldDescription, field.TypeString, value)
	}
	if value, ok := iu.mutation.Status(); ok {
		_spec.SetField(image.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := iu.mutation.UpdatedAt(); ok {
		_spec.SetField(image.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := iu.mutation.CreatedAt(); ok {
		_spec.SetField(image.FieldCreatedAt, field.TypeTime, value)
	}
	if iu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.OwnerTable,
			Columns: []string{image.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.OwnerTable,
			Columns: []string{image.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.FileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   image.FileTable,
			Columns: []string{image.FileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.FileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   image.FileTable,
			Columns: []string{image.FileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{image.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// ImageUpdateOne is the builder for updating a single Image entity.
type ImageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ImageMutation
}

// SetName sets the "name" field.
func (iuo *ImageUpdateOne) SetName(s string) *ImageUpdateOne {
	iuo.mutation.SetName(s)
	return iuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (iuo *ImageUpdateOne) SetNillableName(s *string) *ImageUpdateOne {
	if s != nil {
		iuo.SetName(*s)
	}
	return iuo
}

// SetDescription sets the "description" field.
func (iuo *ImageUpdateOne) SetDescription(s string) *ImageUpdateOne {
	iuo.mutation.SetDescription(s)
	return iuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (iuo *ImageUpdateOne) SetNillableDescription(s *string) *ImageUpdateOne {
	if s != nil {
		iuo.SetDescription(*s)
	}
	return iuo
}

// SetStatus sets the "status" field.
func (iuo *ImageUpdateOne) SetStatus(i image.Status) *ImageUpdateOne {
	iuo.mutation.SetStatus(i)
	return iuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (iuo *ImageUpdateOne) SetNillableStatus(i *image.Status) *ImageUpdateOne {
	if i != nil {
		iuo.SetStatus(*i)
	}
	return iuo
}

// SetUpdatedAt sets the "updated_at" field.
func (iuo *ImageUpdateOne) SetUpdatedAt(t time.Time) *ImageUpdateOne {
	iuo.mutation.SetUpdatedAt(t)
	return iuo
}

// SetCreatedAt sets the "created_at" field.
func (iuo *ImageUpdateOne) SetCreatedAt(t time.Time) *ImageUpdateOne {
	iuo.mutation.SetCreatedAt(t)
	return iuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (iuo *ImageUpdateOne) SetNillableCreatedAt(t *time.Time) *ImageUpdateOne {
	if t != nil {
		iuo.SetCreatedAt(*t)
	}
	return iuo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (iuo *ImageUpdateOne) SetOwnerID(id model.InternalID) *ImageUpdateOne {
	iuo.mutation.SetOwnerID(id)
	return iuo
}

// SetOwner sets the "owner" edge to the User entity.
func (iuo *ImageUpdateOne) SetOwner(u *User) *ImageUpdateOne {
	return iuo.SetOwnerID(u.ID)
}

// SetFileID sets the "file" edge to the File entity by ID.
func (iuo *ImageUpdateOne) SetFileID(id model.InternalID) *ImageUpdateOne {
	iuo.mutation.SetFileID(id)
	return iuo
}

// SetNillableFileID sets the "file" edge to the File entity by ID if the given value is not nil.
func (iuo *ImageUpdateOne) SetNillableFileID(id *model.InternalID) *ImageUpdateOne {
	if id != nil {
		iuo = iuo.SetFileID(*id)
	}
	return iuo
}

// SetFile sets the "file" edge to the File entity.
func (iuo *ImageUpdateOne) SetFile(f *File) *ImageUpdateOne {
	return iuo.SetFileID(f.ID)
}

// Mutation returns the ImageMutation object of the builder.
func (iuo *ImageUpdateOne) Mutation() *ImageMutation {
	return iuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (iuo *ImageUpdateOne) ClearOwner() *ImageUpdateOne {
	iuo.mutation.ClearOwner()
	return iuo
}

// ClearFile clears the "file" edge to the File entity.
func (iuo *ImageUpdateOne) ClearFile() *ImageUpdateOne {
	iuo.mutation.ClearFile()
	return iuo
}

// Where appends a list predicates to the ImageUpdate builder.
func (iuo *ImageUpdateOne) Where(ps ...predicate.Image) *ImageUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *ImageUpdateOne) Select(field string, fields ...string) *ImageUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Image entity.
func (iuo *ImageUpdateOne) Save(ctx context.Context) (*Image, error) {
	iuo.defaults()
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ImageUpdateOne) SaveX(ctx context.Context) *Image {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *ImageUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ImageUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *ImageUpdateOne) defaults() {
	if _, ok := iuo.mutation.UpdatedAt(); !ok {
		v := image.UpdateDefaultUpdatedAt()
		iuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *ImageUpdateOne) check() error {
	if v, ok := iuo.mutation.Status(); ok {
		if err := image.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Image.status": %w`, err)}
		}
	}
	if iuo.mutation.OwnerCleared() && len(iuo.mutation.OwnerIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Image.owner"`)
	}
	return nil
}

func (iuo *ImageUpdateOne) sqlSave(ctx context.Context) (_node *Image, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(image.Table, image.Columns, sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt64))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Image.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, image.FieldID)
		for _, f := range fields {
			if !image.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != image.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.Name(); ok {
		_spec.SetField(image.FieldName, field.TypeString, value)
	}
	if value, ok := iuo.mutation.Description(); ok {
		_spec.SetField(image.FieldDescription, field.TypeString, value)
	}
	if value, ok := iuo.mutation.Status(); ok {
		_spec.SetField(image.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := iuo.mutation.UpdatedAt(); ok {
		_spec.SetField(image.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := iuo.mutation.CreatedAt(); ok {
		_spec.SetField(image.FieldCreatedAt, field.TypeTime, value)
	}
	if iuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.OwnerTable,
			Columns: []string{image.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.OwnerTable,
			Columns: []string{image.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.FileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   image.FileTable,
			Columns: []string{image.FileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.FileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   image.FileTable,
			Columns: []string{image.FileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Image{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{image.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
