// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/predicate"
)

// FeedConfigDelete is the builder for deleting a FeedConfig entity.
type FeedConfigDelete struct {
	config
	hooks    []Hook
	mutation *FeedConfigMutation
}

// Where appends a list predicates to the FeedConfigDelete builder.
func (fcd *FeedConfigDelete) Where(ps ...predicate.FeedConfig) *FeedConfigDelete {
	fcd.mutation.Where(ps...)
	return fcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fcd *FeedConfigDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, FeedConfigMutation](ctx, fcd.sqlExec, fcd.mutation, fcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fcd *FeedConfigDelete) ExecX(ctx context.Context) int {
	n, err := fcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fcd *FeedConfigDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(feedconfig.Table, sqlgraph.NewFieldSpec(feedconfig.FieldID, field.TypeInt64))
	if ps := fcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fcd.mutation.done = true
	return affected, err
}

// FeedConfigDeleteOne is the builder for deleting a single FeedConfig entity.
type FeedConfigDeleteOne struct {
	fcd *FeedConfigDelete
}

// Where appends a list predicates to the FeedConfigDelete builder.
func (fcdo *FeedConfigDeleteOne) Where(ps ...predicate.FeedConfig) *FeedConfigDeleteOne {
	fcdo.fcd.mutation.Where(ps...)
	return fcdo
}

// Exec executes the deletion query.
func (fcdo *FeedConfigDeleteOne) Exec(ctx context.Context) error {
	n, err := fcdo.fcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{feedconfig.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fcdo *FeedConfigDeleteOne) ExecX(ctx context.Context) {
	if err := fcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
