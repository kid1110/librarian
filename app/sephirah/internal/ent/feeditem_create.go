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
	"github.com/tuihub/librarian/app/sephirah/internal/ent/feed"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/feeditem"
	"github.com/tuihub/librarian/internal/model/modelfeed"
)

// FeedItemCreate is the builder for creating a FeedItem entity.
type FeedItemCreate struct {
	config
	mutation *FeedItemMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTitle sets the "title" field.
func (fic *FeedItemCreate) SetTitle(s string) *FeedItemCreate {
	fic.mutation.SetTitle(s)
	return fic
}

// SetAuthors sets the "authors" field.
func (fic *FeedItemCreate) SetAuthors(m []modelfeed.Person) *FeedItemCreate {
	fic.mutation.SetAuthors(m)
	return fic
}

// SetDescription sets the "description" field.
func (fic *FeedItemCreate) SetDescription(s string) *FeedItemCreate {
	fic.mutation.SetDescription(s)
	return fic
}

// SetContent sets the "content" field.
func (fic *FeedItemCreate) SetContent(s string) *FeedItemCreate {
	fic.mutation.SetContent(s)
	return fic
}

// SetGUID sets the "guid" field.
func (fic *FeedItemCreate) SetGUID(s string) *FeedItemCreate {
	fic.mutation.SetGUID(s)
	return fic
}

// SetLink sets the "link" field.
func (fic *FeedItemCreate) SetLink(s string) *FeedItemCreate {
	fic.mutation.SetLink(s)
	return fic
}

// SetImage sets the "image" field.
func (fic *FeedItemCreate) SetImage(m *modelfeed.Image) *FeedItemCreate {
	fic.mutation.SetImage(m)
	return fic
}

// SetPublished sets the "published" field.
func (fic *FeedItemCreate) SetPublished(s string) *FeedItemCreate {
	fic.mutation.SetPublished(s)
	return fic
}

// SetPublishedParsed sets the "published_parsed" field.
func (fic *FeedItemCreate) SetPublishedParsed(t time.Time) *FeedItemCreate {
	fic.mutation.SetPublishedParsed(t)
	return fic
}

// SetUpdated sets the "updated" field.
func (fic *FeedItemCreate) SetUpdated(s string) *FeedItemCreate {
	fic.mutation.SetUpdated(s)
	return fic
}

// SetUpdatedParsed sets the "updated_parsed" field.
func (fic *FeedItemCreate) SetUpdatedParsed(t time.Time) *FeedItemCreate {
	fic.mutation.SetUpdatedParsed(t)
	return fic
}

// SetEnclosure sets the "enclosure" field.
func (fic *FeedItemCreate) SetEnclosure(m []modelfeed.Enclosure) *FeedItemCreate {
	fic.mutation.SetEnclosure(m)
	return fic
}

// SetUpdatedAt sets the "updated_at" field.
func (fic *FeedItemCreate) SetUpdatedAt(t time.Time) *FeedItemCreate {
	fic.mutation.SetUpdatedAt(t)
	return fic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fic *FeedItemCreate) SetNillableUpdatedAt(t *time.Time) *FeedItemCreate {
	if t != nil {
		fic.SetUpdatedAt(*t)
	}
	return fic
}

// SetCreatedAt sets the "created_at" field.
func (fic *FeedItemCreate) SetCreatedAt(t time.Time) *FeedItemCreate {
	fic.mutation.SetCreatedAt(t)
	return fic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fic *FeedItemCreate) SetNillableCreatedAt(t *time.Time) *FeedItemCreate {
	if t != nil {
		fic.SetCreatedAt(*t)
	}
	return fic
}

// SetID sets the "id" field.
func (fic *FeedItemCreate) SetID(i int64) *FeedItemCreate {
	fic.mutation.SetID(i)
	return fic
}

// SetFeedID sets the "feed" edge to the Feed entity by ID.
func (fic *FeedItemCreate) SetFeedID(id int64) *FeedItemCreate {
	fic.mutation.SetFeedID(id)
	return fic
}

// SetFeed sets the "feed" edge to the Feed entity.
func (fic *FeedItemCreate) SetFeed(f *Feed) *FeedItemCreate {
	return fic.SetFeedID(f.ID)
}

// Mutation returns the FeedItemMutation object of the builder.
func (fic *FeedItemCreate) Mutation() *FeedItemMutation {
	return fic.mutation
}

// Save creates the FeedItem in the database.
func (fic *FeedItemCreate) Save(ctx context.Context) (*FeedItem, error) {
	fic.defaults()
	return withHooks[*FeedItem, FeedItemMutation](ctx, fic.sqlSave, fic.mutation, fic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fic *FeedItemCreate) SaveX(ctx context.Context) *FeedItem {
	v, err := fic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fic *FeedItemCreate) Exec(ctx context.Context) error {
	_, err := fic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fic *FeedItemCreate) ExecX(ctx context.Context) {
	if err := fic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fic *FeedItemCreate) defaults() {
	if _, ok := fic.mutation.UpdatedAt(); !ok {
		v := feeditem.DefaultUpdatedAt()
		fic.mutation.SetUpdatedAt(v)
	}
	if _, ok := fic.mutation.CreatedAt(); !ok {
		v := feeditem.DefaultCreatedAt()
		fic.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fic *FeedItemCreate) check() error {
	if _, ok := fic.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "FeedItem.title"`)}
	}
	if _, ok := fic.mutation.Authors(); !ok {
		return &ValidationError{Name: "authors", err: errors.New(`ent: missing required field "FeedItem.authors"`)}
	}
	if _, ok := fic.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "FeedItem.description"`)}
	}
	if _, ok := fic.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "FeedItem.content"`)}
	}
	if _, ok := fic.mutation.GUID(); !ok {
		return &ValidationError{Name: "guid", err: errors.New(`ent: missing required field "FeedItem.guid"`)}
	}
	if _, ok := fic.mutation.Link(); !ok {
		return &ValidationError{Name: "link", err: errors.New(`ent: missing required field "FeedItem.link"`)}
	}
	if _, ok := fic.mutation.Image(); !ok {
		return &ValidationError{Name: "image", err: errors.New(`ent: missing required field "FeedItem.image"`)}
	}
	if _, ok := fic.mutation.Published(); !ok {
		return &ValidationError{Name: "published", err: errors.New(`ent: missing required field "FeedItem.published"`)}
	}
	if _, ok := fic.mutation.PublishedParsed(); !ok {
		return &ValidationError{Name: "published_parsed", err: errors.New(`ent: missing required field "FeedItem.published_parsed"`)}
	}
	if _, ok := fic.mutation.Updated(); !ok {
		return &ValidationError{Name: "updated", err: errors.New(`ent: missing required field "FeedItem.updated"`)}
	}
	if _, ok := fic.mutation.UpdatedParsed(); !ok {
		return &ValidationError{Name: "updated_parsed", err: errors.New(`ent: missing required field "FeedItem.updated_parsed"`)}
	}
	if _, ok := fic.mutation.Enclosure(); !ok {
		return &ValidationError{Name: "enclosure", err: errors.New(`ent: missing required field "FeedItem.enclosure"`)}
	}
	if _, ok := fic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "FeedItem.updated_at"`)}
	}
	if _, ok := fic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "FeedItem.created_at"`)}
	}
	if _, ok := fic.mutation.FeedID(); !ok {
		return &ValidationError{Name: "feed", err: errors.New(`ent: missing required edge "FeedItem.feed"`)}
	}
	return nil
}

func (fic *FeedItemCreate) sqlSave(ctx context.Context) (*FeedItem, error) {
	if err := fic.check(); err != nil {
		return nil, err
	}
	_node, _spec := fic.createSpec()
	if err := sqlgraph.CreateNode(ctx, fic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	fic.mutation.id = &_node.ID
	fic.mutation.done = true
	return _node, nil
}

func (fic *FeedItemCreate) createSpec() (*FeedItem, *sqlgraph.CreateSpec) {
	var (
		_node = &FeedItem{config: fic.config}
		_spec = sqlgraph.NewCreateSpec(feeditem.Table, sqlgraph.NewFieldSpec(feeditem.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = fic.conflict
	if id, ok := fic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fic.mutation.Title(); ok {
		_spec.SetField(feeditem.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := fic.mutation.Authors(); ok {
		_spec.SetField(feeditem.FieldAuthors, field.TypeJSON, value)
		_node.Authors = value
	}
	if value, ok := fic.mutation.Description(); ok {
		_spec.SetField(feeditem.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := fic.mutation.Content(); ok {
		_spec.SetField(feeditem.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := fic.mutation.GUID(); ok {
		_spec.SetField(feeditem.FieldGUID, field.TypeString, value)
		_node.GUID = value
	}
	if value, ok := fic.mutation.Link(); ok {
		_spec.SetField(feeditem.FieldLink, field.TypeString, value)
		_node.Link = value
	}
	if value, ok := fic.mutation.Image(); ok {
		_spec.SetField(feeditem.FieldImage, field.TypeJSON, value)
		_node.Image = value
	}
	if value, ok := fic.mutation.Published(); ok {
		_spec.SetField(feeditem.FieldPublished, field.TypeString, value)
		_node.Published = value
	}
	if value, ok := fic.mutation.PublishedParsed(); ok {
		_spec.SetField(feeditem.FieldPublishedParsed, field.TypeTime, value)
		_node.PublishedParsed = value
	}
	if value, ok := fic.mutation.Updated(); ok {
		_spec.SetField(feeditem.FieldUpdated, field.TypeString, value)
		_node.Updated = value
	}
	if value, ok := fic.mutation.UpdatedParsed(); ok {
		_spec.SetField(feeditem.FieldUpdatedParsed, field.TypeTime, value)
		_node.UpdatedParsed = value
	}
	if value, ok := fic.mutation.Enclosure(); ok {
		_spec.SetField(feeditem.FieldEnclosure, field.TypeJSON, value)
		_node.Enclosure = value
	}
	if value, ok := fic.mutation.UpdatedAt(); ok {
		_spec.SetField(feeditem.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := fic.mutation.CreatedAt(); ok {
		_spec.SetField(feeditem.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := fic.mutation.FeedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   feeditem.FeedTable,
			Columns: []string{feeditem.FeedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: feed.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.feed_item = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FeedItem.Create().
//		SetTitle(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FeedItemUpsert) {
//			SetTitle(v+v).
//		}).
//		Exec(ctx)
func (fic *FeedItemCreate) OnConflict(opts ...sql.ConflictOption) *FeedItemUpsertOne {
	fic.conflict = opts
	return &FeedItemUpsertOne{
		create: fic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FeedItem.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fic *FeedItemCreate) OnConflictColumns(columns ...string) *FeedItemUpsertOne {
	fic.conflict = append(fic.conflict, sql.ConflictColumns(columns...))
	return &FeedItemUpsertOne{
		create: fic,
	}
}

type (
	// FeedItemUpsertOne is the builder for "upsert"-ing
	//  one FeedItem node.
	FeedItemUpsertOne struct {
		create *FeedItemCreate
	}

	// FeedItemUpsert is the "OnConflict" setter.
	FeedItemUpsert struct {
		*sql.UpdateSet
	}
)

// SetTitle sets the "title" field.
func (u *FeedItemUpsert) SetTitle(v string) *FeedItemUpsert {
	u.Set(feeditem.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *FeedItemUpsert) UpdateTitle() *FeedItemUpsert {
	u.SetExcluded(feeditem.FieldTitle)
	return u
}

// SetAuthors sets the "authors" field.
func (u *FeedItemUpsert) SetAuthors(v []modelfeed.Person) *FeedItemUpsert {
	u.Set(feeditem.FieldAuthors, v)
	return u
}

// UpdateAuthors sets the "authors" field to the value that was provided on create.
func (u *FeedItemUpsert) UpdateAuthors() *FeedItemUpsert {
	u.SetExcluded(feeditem.FieldAuthors)
	return u
}

// SetDescription sets the "description" field.
func (u *FeedItemUpsert) SetDescription(v string) *FeedItemUpsert {
	u.Set(feeditem.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *FeedItemUpsert) UpdateDescription() *FeedItemUpsert {
	u.SetExcluded(feeditem.FieldDescription)
	return u
}

// SetContent sets the "content" field.
func (u *FeedItemUpsert) SetContent(v string) *FeedItemUpsert {
	u.Set(feeditem.FieldContent, v)
	return u
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *FeedItemUpsert) UpdateContent() *FeedItemUpsert {
	u.SetExcluded(feeditem.FieldContent)
	return u
}

// SetGUID sets the "guid" field.
func (u *FeedItemUpsert) SetGUID(v string) *FeedItemUpsert {
	u.Set(feeditem.FieldGUID, v)
	return u
}

// UpdateGUID sets the "guid" field to the value that was provided on create.
func (u *FeedItemUpsert) UpdateGUID() *FeedItemUpsert {
	u.SetExcluded(feeditem.FieldGUID)
	return u
}

// SetLink sets the "link" field.
func (u *FeedItemUpsert) SetLink(v string) *FeedItemUpsert {
	u.Set(feeditem.FieldLink, v)
	return u
}

// UpdateLink sets the "link" field to the value that was provided on create.
func (u *FeedItemUpsert) UpdateLink() *FeedItemUpsert {
	u.SetExcluded(feeditem.FieldLink)
	return u
}

// SetImage sets the "image" field.
func (u *FeedItemUpsert) SetImage(v *modelfeed.Image) *FeedItemUpsert {
	u.Set(feeditem.FieldImage, v)
	return u
}

// UpdateImage sets the "image" field to the value that was provided on create.
func (u *FeedItemUpsert) UpdateImage() *FeedItemUpsert {
	u.SetExcluded(feeditem.FieldImage)
	return u
}

// SetPublished sets the "published" field.
func (u *FeedItemUpsert) SetPublished(v string) *FeedItemUpsert {
	u.Set(feeditem.FieldPublished, v)
	return u
}

// UpdatePublished sets the "published" field to the value that was provided on create.
func (u *FeedItemUpsert) UpdatePublished() *FeedItemUpsert {
	u.SetExcluded(feeditem.FieldPublished)
	return u
}

// SetPublishedParsed sets the "published_parsed" field.
func (u *FeedItemUpsert) SetPublishedParsed(v time.Time) *FeedItemUpsert {
	u.Set(feeditem.FieldPublishedParsed, v)
	return u
}

// UpdatePublishedParsed sets the "published_parsed" field to the value that was provided on create.
func (u *FeedItemUpsert) UpdatePublishedParsed() *FeedItemUpsert {
	u.SetExcluded(feeditem.FieldPublishedParsed)
	return u
}

// SetUpdated sets the "updated" field.
func (u *FeedItemUpsert) SetUpdated(v string) *FeedItemUpsert {
	u.Set(feeditem.FieldUpdated, v)
	return u
}

// UpdateUpdated sets the "updated" field to the value that was provided on create.
func (u *FeedItemUpsert) UpdateUpdated() *FeedItemUpsert {
	u.SetExcluded(feeditem.FieldUpdated)
	return u
}

// SetUpdatedParsed sets the "updated_parsed" field.
func (u *FeedItemUpsert) SetUpdatedParsed(v time.Time) *FeedItemUpsert {
	u.Set(feeditem.FieldUpdatedParsed, v)
	return u
}

// UpdateUpdatedParsed sets the "updated_parsed" field to the value that was provided on create.
func (u *FeedItemUpsert) UpdateUpdatedParsed() *FeedItemUpsert {
	u.SetExcluded(feeditem.FieldUpdatedParsed)
	return u
}

// SetEnclosure sets the "enclosure" field.
func (u *FeedItemUpsert) SetEnclosure(v []modelfeed.Enclosure) *FeedItemUpsert {
	u.Set(feeditem.FieldEnclosure, v)
	return u
}

// UpdateEnclosure sets the "enclosure" field to the value that was provided on create.
func (u *FeedItemUpsert) UpdateEnclosure() *FeedItemUpsert {
	u.SetExcluded(feeditem.FieldEnclosure)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FeedItemUpsert) SetUpdatedAt(v time.Time) *FeedItemUpsert {
	u.Set(feeditem.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FeedItemUpsert) UpdateUpdatedAt() *FeedItemUpsert {
	u.SetExcluded(feeditem.FieldUpdatedAt)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *FeedItemUpsert) SetCreatedAt(v time.Time) *FeedItemUpsert {
	u.Set(feeditem.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FeedItemUpsert) UpdateCreatedAt() *FeedItemUpsert {
	u.SetExcluded(feeditem.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.FeedItem.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(feeditem.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FeedItemUpsertOne) UpdateNewValues() *FeedItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(feeditem.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FeedItem.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FeedItemUpsertOne) Ignore() *FeedItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FeedItemUpsertOne) DoNothing() *FeedItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FeedItemCreate.OnConflict
// documentation for more info.
func (u *FeedItemUpsertOne) Update(set func(*FeedItemUpsert)) *FeedItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FeedItemUpsert{UpdateSet: update})
	}))
	return u
}

// SetTitle sets the "title" field.
func (u *FeedItemUpsertOne) SetTitle(v string) *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *FeedItemUpsertOne) UpdateTitle() *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdateTitle()
	})
}

// SetAuthors sets the "authors" field.
func (u *FeedItemUpsertOne) SetAuthors(v []modelfeed.Person) *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetAuthors(v)
	})
}

// UpdateAuthors sets the "authors" field to the value that was provided on create.
func (u *FeedItemUpsertOne) UpdateAuthors() *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdateAuthors()
	})
}

// SetDescription sets the "description" field.
func (u *FeedItemUpsertOne) SetDescription(v string) *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *FeedItemUpsertOne) UpdateDescription() *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdateDescription()
	})
}

// SetContent sets the "content" field.
func (u *FeedItemUpsertOne) SetContent(v string) *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *FeedItemUpsertOne) UpdateContent() *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdateContent()
	})
}

// SetGUID sets the "guid" field.
func (u *FeedItemUpsertOne) SetGUID(v string) *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetGUID(v)
	})
}

// UpdateGUID sets the "guid" field to the value that was provided on create.
func (u *FeedItemUpsertOne) UpdateGUID() *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdateGUID()
	})
}

// SetLink sets the "link" field.
func (u *FeedItemUpsertOne) SetLink(v string) *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetLink(v)
	})
}

// UpdateLink sets the "link" field to the value that was provided on create.
func (u *FeedItemUpsertOne) UpdateLink() *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdateLink()
	})
}

// SetImage sets the "image" field.
func (u *FeedItemUpsertOne) SetImage(v *modelfeed.Image) *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetImage(v)
	})
}

// UpdateImage sets the "image" field to the value that was provided on create.
func (u *FeedItemUpsertOne) UpdateImage() *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdateImage()
	})
}

// SetPublished sets the "published" field.
func (u *FeedItemUpsertOne) SetPublished(v string) *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetPublished(v)
	})
}

// UpdatePublished sets the "published" field to the value that was provided on create.
func (u *FeedItemUpsertOne) UpdatePublished() *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdatePublished()
	})
}

// SetPublishedParsed sets the "published_parsed" field.
func (u *FeedItemUpsertOne) SetPublishedParsed(v time.Time) *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetPublishedParsed(v)
	})
}

// UpdatePublishedParsed sets the "published_parsed" field to the value that was provided on create.
func (u *FeedItemUpsertOne) UpdatePublishedParsed() *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdatePublishedParsed()
	})
}

// SetUpdated sets the "updated" field.
func (u *FeedItemUpsertOne) SetUpdated(v string) *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetUpdated(v)
	})
}

// UpdateUpdated sets the "updated" field to the value that was provided on create.
func (u *FeedItemUpsertOne) UpdateUpdated() *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdateUpdated()
	})
}

// SetUpdatedParsed sets the "updated_parsed" field.
func (u *FeedItemUpsertOne) SetUpdatedParsed(v time.Time) *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetUpdatedParsed(v)
	})
}

// UpdateUpdatedParsed sets the "updated_parsed" field to the value that was provided on create.
func (u *FeedItemUpsertOne) UpdateUpdatedParsed() *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdateUpdatedParsed()
	})
}

// SetEnclosure sets the "enclosure" field.
func (u *FeedItemUpsertOne) SetEnclosure(v []modelfeed.Enclosure) *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetEnclosure(v)
	})
}

// UpdateEnclosure sets the "enclosure" field to the value that was provided on create.
func (u *FeedItemUpsertOne) UpdateEnclosure() *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdateEnclosure()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FeedItemUpsertOne) SetUpdatedAt(v time.Time) *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FeedItemUpsertOne) UpdateUpdatedAt() *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *FeedItemUpsertOne) SetCreatedAt(v time.Time) *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FeedItemUpsertOne) UpdateCreatedAt() *FeedItemUpsertOne {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *FeedItemUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FeedItemCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FeedItemUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FeedItemUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FeedItemUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FeedItemCreateBulk is the builder for creating many FeedItem entities in bulk.
type FeedItemCreateBulk struct {
	config
	builders []*FeedItemCreate
	conflict []sql.ConflictOption
}

// Save creates the FeedItem entities in the database.
func (ficb *FeedItemCreateBulk) Save(ctx context.Context) ([]*FeedItem, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ficb.builders))
	nodes := make([]*FeedItem, len(ficb.builders))
	mutators := make([]Mutator, len(ficb.builders))
	for i := range ficb.builders {
		func(i int, root context.Context) {
			builder := ficb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FeedItemMutation)
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
					_, err = mutators[i+1].Mutate(root, ficb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ficb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ficb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ficb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ficb *FeedItemCreateBulk) SaveX(ctx context.Context) []*FeedItem {
	v, err := ficb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ficb *FeedItemCreateBulk) Exec(ctx context.Context) error {
	_, err := ficb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ficb *FeedItemCreateBulk) ExecX(ctx context.Context) {
	if err := ficb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FeedItem.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FeedItemUpsert) {
//			SetTitle(v+v).
//		}).
//		Exec(ctx)
func (ficb *FeedItemCreateBulk) OnConflict(opts ...sql.ConflictOption) *FeedItemUpsertBulk {
	ficb.conflict = opts
	return &FeedItemUpsertBulk{
		create: ficb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FeedItem.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ficb *FeedItemCreateBulk) OnConflictColumns(columns ...string) *FeedItemUpsertBulk {
	ficb.conflict = append(ficb.conflict, sql.ConflictColumns(columns...))
	return &FeedItemUpsertBulk{
		create: ficb,
	}
}

// FeedItemUpsertBulk is the builder for "upsert"-ing
// a bulk of FeedItem nodes.
type FeedItemUpsertBulk struct {
	create *FeedItemCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.FeedItem.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(feeditem.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FeedItemUpsertBulk) UpdateNewValues() *FeedItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(feeditem.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FeedItem.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FeedItemUpsertBulk) Ignore() *FeedItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FeedItemUpsertBulk) DoNothing() *FeedItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FeedItemCreateBulk.OnConflict
// documentation for more info.
func (u *FeedItemUpsertBulk) Update(set func(*FeedItemUpsert)) *FeedItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FeedItemUpsert{UpdateSet: update})
	}))
	return u
}

// SetTitle sets the "title" field.
func (u *FeedItemUpsertBulk) SetTitle(v string) *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *FeedItemUpsertBulk) UpdateTitle() *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdateTitle()
	})
}

// SetAuthors sets the "authors" field.
func (u *FeedItemUpsertBulk) SetAuthors(v []modelfeed.Person) *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetAuthors(v)
	})
}

// UpdateAuthors sets the "authors" field to the value that was provided on create.
func (u *FeedItemUpsertBulk) UpdateAuthors() *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdateAuthors()
	})
}

// SetDescription sets the "description" field.
func (u *FeedItemUpsertBulk) SetDescription(v string) *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *FeedItemUpsertBulk) UpdateDescription() *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdateDescription()
	})
}

// SetContent sets the "content" field.
func (u *FeedItemUpsertBulk) SetContent(v string) *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *FeedItemUpsertBulk) UpdateContent() *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdateContent()
	})
}

// SetGUID sets the "guid" field.
func (u *FeedItemUpsertBulk) SetGUID(v string) *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetGUID(v)
	})
}

// UpdateGUID sets the "guid" field to the value that was provided on create.
func (u *FeedItemUpsertBulk) UpdateGUID() *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdateGUID()
	})
}

// SetLink sets the "link" field.
func (u *FeedItemUpsertBulk) SetLink(v string) *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetLink(v)
	})
}

// UpdateLink sets the "link" field to the value that was provided on create.
func (u *FeedItemUpsertBulk) UpdateLink() *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdateLink()
	})
}

// SetImage sets the "image" field.
func (u *FeedItemUpsertBulk) SetImage(v *modelfeed.Image) *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetImage(v)
	})
}

// UpdateImage sets the "image" field to the value that was provided on create.
func (u *FeedItemUpsertBulk) UpdateImage() *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdateImage()
	})
}

// SetPublished sets the "published" field.
func (u *FeedItemUpsertBulk) SetPublished(v string) *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetPublished(v)
	})
}

// UpdatePublished sets the "published" field to the value that was provided on create.
func (u *FeedItemUpsertBulk) UpdatePublished() *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdatePublished()
	})
}

// SetPublishedParsed sets the "published_parsed" field.
func (u *FeedItemUpsertBulk) SetPublishedParsed(v time.Time) *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetPublishedParsed(v)
	})
}

// UpdatePublishedParsed sets the "published_parsed" field to the value that was provided on create.
func (u *FeedItemUpsertBulk) UpdatePublishedParsed() *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdatePublishedParsed()
	})
}

// SetUpdated sets the "updated" field.
func (u *FeedItemUpsertBulk) SetUpdated(v string) *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetUpdated(v)
	})
}

// UpdateUpdated sets the "updated" field to the value that was provided on create.
func (u *FeedItemUpsertBulk) UpdateUpdated() *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdateUpdated()
	})
}

// SetUpdatedParsed sets the "updated_parsed" field.
func (u *FeedItemUpsertBulk) SetUpdatedParsed(v time.Time) *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetUpdatedParsed(v)
	})
}

// UpdateUpdatedParsed sets the "updated_parsed" field to the value that was provided on create.
func (u *FeedItemUpsertBulk) UpdateUpdatedParsed() *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdateUpdatedParsed()
	})
}

// SetEnclosure sets the "enclosure" field.
func (u *FeedItemUpsertBulk) SetEnclosure(v []modelfeed.Enclosure) *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetEnclosure(v)
	})
}

// UpdateEnclosure sets the "enclosure" field to the value that was provided on create.
func (u *FeedItemUpsertBulk) UpdateEnclosure() *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdateEnclosure()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FeedItemUpsertBulk) SetUpdatedAt(v time.Time) *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FeedItemUpsertBulk) UpdateUpdatedAt() *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *FeedItemUpsertBulk) SetCreatedAt(v time.Time) *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FeedItemUpsertBulk) UpdateCreatedAt() *FeedItemUpsertBulk {
	return u.Update(func(s *FeedItemUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *FeedItemUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FeedItemCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FeedItemCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FeedItemUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
