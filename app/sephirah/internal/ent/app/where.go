// Code generated by ent, DO NOT EDIT.

package app

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/model"
)

// ID filters vertices based on their ID field.
func ID(id model.InternalID) predicate.App {
	return predicate.App(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id model.InternalID) predicate.App {
	return predicate.App(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id model.InternalID) predicate.App {
	return predicate.App(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...model.InternalID) predicate.App {
	return predicate.App(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...model.InternalID) predicate.App {
	return predicate.App(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id model.InternalID) predicate.App {
	return predicate.App(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id model.InternalID) predicate.App {
	return predicate.App(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id model.InternalID) predicate.App {
	return predicate.App(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id model.InternalID) predicate.App {
	return predicate.App(sql.FieldLTE(FieldID, id))
}

// SourceAppID applies equality check predicate on the "source_app_id" field. It's identical to SourceAppIDEQ.
func SourceAppID(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldSourceAppID, v))
}

// SourceURL applies equality check predicate on the "source_url" field. It's identical to SourceURLEQ.
func SourceURL(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldSourceURL, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldName, v))
}

// ShortDescription applies equality check predicate on the "short_description" field. It's identical to ShortDescriptionEQ.
func ShortDescription(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldShortDescription, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldDescription, v))
}

// ImageURL applies equality check predicate on the "image_url" field. It's identical to ImageURLEQ.
func ImageURL(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldImageURL, v))
}

// ReleaseDate applies equality check predicate on the "release_date" field. It's identical to ReleaseDateEQ.
func ReleaseDate(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldReleaseDate, v))
}

// Developer applies equality check predicate on the "developer" field. It's identical to DeveloperEQ.
func Developer(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldDeveloper, v))
}

// Publisher applies equality check predicate on the "publisher" field. It's identical to PublisherEQ.
func Publisher(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldPublisher, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldVersion, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.App {
	return predicate.App(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.App {
	return predicate.App(sql.FieldEQ(FieldCreatedAt, v))
}

// SourceEQ applies the EQ predicate on the "source" field.
func SourceEQ(v Source) predicate.App {
	return predicate.App(sql.FieldEQ(FieldSource, v))
}

// SourceNEQ applies the NEQ predicate on the "source" field.
func SourceNEQ(v Source) predicate.App {
	return predicate.App(sql.FieldNEQ(FieldSource, v))
}

// SourceIn applies the In predicate on the "source" field.
func SourceIn(vs ...Source) predicate.App {
	return predicate.App(sql.FieldIn(FieldSource, vs...))
}

// SourceNotIn applies the NotIn predicate on the "source" field.
func SourceNotIn(vs ...Source) predicate.App {
	return predicate.App(sql.FieldNotIn(FieldSource, vs...))
}

// SourceAppIDEQ applies the EQ predicate on the "source_app_id" field.
func SourceAppIDEQ(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldSourceAppID, v))
}

// SourceAppIDNEQ applies the NEQ predicate on the "source_app_id" field.
func SourceAppIDNEQ(v string) predicate.App {
	return predicate.App(sql.FieldNEQ(FieldSourceAppID, v))
}

// SourceAppIDIn applies the In predicate on the "source_app_id" field.
func SourceAppIDIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldIn(FieldSourceAppID, vs...))
}

// SourceAppIDNotIn applies the NotIn predicate on the "source_app_id" field.
func SourceAppIDNotIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldNotIn(FieldSourceAppID, vs...))
}

// SourceAppIDGT applies the GT predicate on the "source_app_id" field.
func SourceAppIDGT(v string) predicate.App {
	return predicate.App(sql.FieldGT(FieldSourceAppID, v))
}

// SourceAppIDGTE applies the GTE predicate on the "source_app_id" field.
func SourceAppIDGTE(v string) predicate.App {
	return predicate.App(sql.FieldGTE(FieldSourceAppID, v))
}

// SourceAppIDLT applies the LT predicate on the "source_app_id" field.
func SourceAppIDLT(v string) predicate.App {
	return predicate.App(sql.FieldLT(FieldSourceAppID, v))
}

// SourceAppIDLTE applies the LTE predicate on the "source_app_id" field.
func SourceAppIDLTE(v string) predicate.App {
	return predicate.App(sql.FieldLTE(FieldSourceAppID, v))
}

// SourceAppIDContains applies the Contains predicate on the "source_app_id" field.
func SourceAppIDContains(v string) predicate.App {
	return predicate.App(sql.FieldContains(FieldSourceAppID, v))
}

// SourceAppIDHasPrefix applies the HasPrefix predicate on the "source_app_id" field.
func SourceAppIDHasPrefix(v string) predicate.App {
	return predicate.App(sql.FieldHasPrefix(FieldSourceAppID, v))
}

// SourceAppIDHasSuffix applies the HasSuffix predicate on the "source_app_id" field.
func SourceAppIDHasSuffix(v string) predicate.App {
	return predicate.App(sql.FieldHasSuffix(FieldSourceAppID, v))
}

// SourceAppIDEqualFold applies the EqualFold predicate on the "source_app_id" field.
func SourceAppIDEqualFold(v string) predicate.App {
	return predicate.App(sql.FieldEqualFold(FieldSourceAppID, v))
}

// SourceAppIDContainsFold applies the ContainsFold predicate on the "source_app_id" field.
func SourceAppIDContainsFold(v string) predicate.App {
	return predicate.App(sql.FieldContainsFold(FieldSourceAppID, v))
}

// SourceURLEQ applies the EQ predicate on the "source_url" field.
func SourceURLEQ(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldSourceURL, v))
}

// SourceURLNEQ applies the NEQ predicate on the "source_url" field.
func SourceURLNEQ(v string) predicate.App {
	return predicate.App(sql.FieldNEQ(FieldSourceURL, v))
}

// SourceURLIn applies the In predicate on the "source_url" field.
func SourceURLIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldIn(FieldSourceURL, vs...))
}

// SourceURLNotIn applies the NotIn predicate on the "source_url" field.
func SourceURLNotIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldNotIn(FieldSourceURL, vs...))
}

// SourceURLGT applies the GT predicate on the "source_url" field.
func SourceURLGT(v string) predicate.App {
	return predicate.App(sql.FieldGT(FieldSourceURL, v))
}

// SourceURLGTE applies the GTE predicate on the "source_url" field.
func SourceURLGTE(v string) predicate.App {
	return predicate.App(sql.FieldGTE(FieldSourceURL, v))
}

// SourceURLLT applies the LT predicate on the "source_url" field.
func SourceURLLT(v string) predicate.App {
	return predicate.App(sql.FieldLT(FieldSourceURL, v))
}

// SourceURLLTE applies the LTE predicate on the "source_url" field.
func SourceURLLTE(v string) predicate.App {
	return predicate.App(sql.FieldLTE(FieldSourceURL, v))
}

// SourceURLContains applies the Contains predicate on the "source_url" field.
func SourceURLContains(v string) predicate.App {
	return predicate.App(sql.FieldContains(FieldSourceURL, v))
}

// SourceURLHasPrefix applies the HasPrefix predicate on the "source_url" field.
func SourceURLHasPrefix(v string) predicate.App {
	return predicate.App(sql.FieldHasPrefix(FieldSourceURL, v))
}

// SourceURLHasSuffix applies the HasSuffix predicate on the "source_url" field.
func SourceURLHasSuffix(v string) predicate.App {
	return predicate.App(sql.FieldHasSuffix(FieldSourceURL, v))
}

// SourceURLEqualFold applies the EqualFold predicate on the "source_url" field.
func SourceURLEqualFold(v string) predicate.App {
	return predicate.App(sql.FieldEqualFold(FieldSourceURL, v))
}

// SourceURLContainsFold applies the ContainsFold predicate on the "source_url" field.
func SourceURLContainsFold(v string) predicate.App {
	return predicate.App(sql.FieldContainsFold(FieldSourceURL, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.App {
	return predicate.App(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.App {
	return predicate.App(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.App {
	return predicate.App(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.App {
	return predicate.App(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.App {
	return predicate.App(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.App {
	return predicate.App(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.App {
	return predicate.App(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.App {
	return predicate.App(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.App {
	return predicate.App(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.App {
	return predicate.App(sql.FieldContainsFold(FieldName, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.App {
	return predicate.App(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.App {
	return predicate.App(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.App {
	return predicate.App(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.App {
	return predicate.App(sql.FieldNotIn(FieldType, vs...))
}

// ShortDescriptionEQ applies the EQ predicate on the "short_description" field.
func ShortDescriptionEQ(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldShortDescription, v))
}

// ShortDescriptionNEQ applies the NEQ predicate on the "short_description" field.
func ShortDescriptionNEQ(v string) predicate.App {
	return predicate.App(sql.FieldNEQ(FieldShortDescription, v))
}

// ShortDescriptionIn applies the In predicate on the "short_description" field.
func ShortDescriptionIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldIn(FieldShortDescription, vs...))
}

// ShortDescriptionNotIn applies the NotIn predicate on the "short_description" field.
func ShortDescriptionNotIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldNotIn(FieldShortDescription, vs...))
}

// ShortDescriptionGT applies the GT predicate on the "short_description" field.
func ShortDescriptionGT(v string) predicate.App {
	return predicate.App(sql.FieldGT(FieldShortDescription, v))
}

// ShortDescriptionGTE applies the GTE predicate on the "short_description" field.
func ShortDescriptionGTE(v string) predicate.App {
	return predicate.App(sql.FieldGTE(FieldShortDescription, v))
}

// ShortDescriptionLT applies the LT predicate on the "short_description" field.
func ShortDescriptionLT(v string) predicate.App {
	return predicate.App(sql.FieldLT(FieldShortDescription, v))
}

// ShortDescriptionLTE applies the LTE predicate on the "short_description" field.
func ShortDescriptionLTE(v string) predicate.App {
	return predicate.App(sql.FieldLTE(FieldShortDescription, v))
}

// ShortDescriptionContains applies the Contains predicate on the "short_description" field.
func ShortDescriptionContains(v string) predicate.App {
	return predicate.App(sql.FieldContains(FieldShortDescription, v))
}

// ShortDescriptionHasPrefix applies the HasPrefix predicate on the "short_description" field.
func ShortDescriptionHasPrefix(v string) predicate.App {
	return predicate.App(sql.FieldHasPrefix(FieldShortDescription, v))
}

// ShortDescriptionHasSuffix applies the HasSuffix predicate on the "short_description" field.
func ShortDescriptionHasSuffix(v string) predicate.App {
	return predicate.App(sql.FieldHasSuffix(FieldShortDescription, v))
}

// ShortDescriptionEqualFold applies the EqualFold predicate on the "short_description" field.
func ShortDescriptionEqualFold(v string) predicate.App {
	return predicate.App(sql.FieldEqualFold(FieldShortDescription, v))
}

// ShortDescriptionContainsFold applies the ContainsFold predicate on the "short_description" field.
func ShortDescriptionContainsFold(v string) predicate.App {
	return predicate.App(sql.FieldContainsFold(FieldShortDescription, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.App {
	return predicate.App(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.App {
	return predicate.App(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.App {
	return predicate.App(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.App {
	return predicate.App(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.App {
	return predicate.App(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.App {
	return predicate.App(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.App {
	return predicate.App(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.App {
	return predicate.App(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.App {
	return predicate.App(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.App {
	return predicate.App(sql.FieldContainsFold(FieldDescription, v))
}

// ImageURLEQ applies the EQ predicate on the "image_url" field.
func ImageURLEQ(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldImageURL, v))
}

// ImageURLNEQ applies the NEQ predicate on the "image_url" field.
func ImageURLNEQ(v string) predicate.App {
	return predicate.App(sql.FieldNEQ(FieldImageURL, v))
}

// ImageURLIn applies the In predicate on the "image_url" field.
func ImageURLIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldIn(FieldImageURL, vs...))
}

// ImageURLNotIn applies the NotIn predicate on the "image_url" field.
func ImageURLNotIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldNotIn(FieldImageURL, vs...))
}

// ImageURLGT applies the GT predicate on the "image_url" field.
func ImageURLGT(v string) predicate.App {
	return predicate.App(sql.FieldGT(FieldImageURL, v))
}

// ImageURLGTE applies the GTE predicate on the "image_url" field.
func ImageURLGTE(v string) predicate.App {
	return predicate.App(sql.FieldGTE(FieldImageURL, v))
}

// ImageURLLT applies the LT predicate on the "image_url" field.
func ImageURLLT(v string) predicate.App {
	return predicate.App(sql.FieldLT(FieldImageURL, v))
}

// ImageURLLTE applies the LTE predicate on the "image_url" field.
func ImageURLLTE(v string) predicate.App {
	return predicate.App(sql.FieldLTE(FieldImageURL, v))
}

// ImageURLContains applies the Contains predicate on the "image_url" field.
func ImageURLContains(v string) predicate.App {
	return predicate.App(sql.FieldContains(FieldImageURL, v))
}

// ImageURLHasPrefix applies the HasPrefix predicate on the "image_url" field.
func ImageURLHasPrefix(v string) predicate.App {
	return predicate.App(sql.FieldHasPrefix(FieldImageURL, v))
}

// ImageURLHasSuffix applies the HasSuffix predicate on the "image_url" field.
func ImageURLHasSuffix(v string) predicate.App {
	return predicate.App(sql.FieldHasSuffix(FieldImageURL, v))
}

// ImageURLEqualFold applies the EqualFold predicate on the "image_url" field.
func ImageURLEqualFold(v string) predicate.App {
	return predicate.App(sql.FieldEqualFold(FieldImageURL, v))
}

// ImageURLContainsFold applies the ContainsFold predicate on the "image_url" field.
func ImageURLContainsFold(v string) predicate.App {
	return predicate.App(sql.FieldContainsFold(FieldImageURL, v))
}

// ReleaseDateEQ applies the EQ predicate on the "release_date" field.
func ReleaseDateEQ(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldReleaseDate, v))
}

// ReleaseDateNEQ applies the NEQ predicate on the "release_date" field.
func ReleaseDateNEQ(v string) predicate.App {
	return predicate.App(sql.FieldNEQ(FieldReleaseDate, v))
}

// ReleaseDateIn applies the In predicate on the "release_date" field.
func ReleaseDateIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldIn(FieldReleaseDate, vs...))
}

// ReleaseDateNotIn applies the NotIn predicate on the "release_date" field.
func ReleaseDateNotIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldNotIn(FieldReleaseDate, vs...))
}

// ReleaseDateGT applies the GT predicate on the "release_date" field.
func ReleaseDateGT(v string) predicate.App {
	return predicate.App(sql.FieldGT(FieldReleaseDate, v))
}

// ReleaseDateGTE applies the GTE predicate on the "release_date" field.
func ReleaseDateGTE(v string) predicate.App {
	return predicate.App(sql.FieldGTE(FieldReleaseDate, v))
}

// ReleaseDateLT applies the LT predicate on the "release_date" field.
func ReleaseDateLT(v string) predicate.App {
	return predicate.App(sql.FieldLT(FieldReleaseDate, v))
}

// ReleaseDateLTE applies the LTE predicate on the "release_date" field.
func ReleaseDateLTE(v string) predicate.App {
	return predicate.App(sql.FieldLTE(FieldReleaseDate, v))
}

// ReleaseDateContains applies the Contains predicate on the "release_date" field.
func ReleaseDateContains(v string) predicate.App {
	return predicate.App(sql.FieldContains(FieldReleaseDate, v))
}

// ReleaseDateHasPrefix applies the HasPrefix predicate on the "release_date" field.
func ReleaseDateHasPrefix(v string) predicate.App {
	return predicate.App(sql.FieldHasPrefix(FieldReleaseDate, v))
}

// ReleaseDateHasSuffix applies the HasSuffix predicate on the "release_date" field.
func ReleaseDateHasSuffix(v string) predicate.App {
	return predicate.App(sql.FieldHasSuffix(FieldReleaseDate, v))
}

// ReleaseDateEqualFold applies the EqualFold predicate on the "release_date" field.
func ReleaseDateEqualFold(v string) predicate.App {
	return predicate.App(sql.FieldEqualFold(FieldReleaseDate, v))
}

// ReleaseDateContainsFold applies the ContainsFold predicate on the "release_date" field.
func ReleaseDateContainsFold(v string) predicate.App {
	return predicate.App(sql.FieldContainsFold(FieldReleaseDate, v))
}

// DeveloperEQ applies the EQ predicate on the "developer" field.
func DeveloperEQ(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldDeveloper, v))
}

// DeveloperNEQ applies the NEQ predicate on the "developer" field.
func DeveloperNEQ(v string) predicate.App {
	return predicate.App(sql.FieldNEQ(FieldDeveloper, v))
}

// DeveloperIn applies the In predicate on the "developer" field.
func DeveloperIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldIn(FieldDeveloper, vs...))
}

// DeveloperNotIn applies the NotIn predicate on the "developer" field.
func DeveloperNotIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldNotIn(FieldDeveloper, vs...))
}

// DeveloperGT applies the GT predicate on the "developer" field.
func DeveloperGT(v string) predicate.App {
	return predicate.App(sql.FieldGT(FieldDeveloper, v))
}

// DeveloperGTE applies the GTE predicate on the "developer" field.
func DeveloperGTE(v string) predicate.App {
	return predicate.App(sql.FieldGTE(FieldDeveloper, v))
}

// DeveloperLT applies the LT predicate on the "developer" field.
func DeveloperLT(v string) predicate.App {
	return predicate.App(sql.FieldLT(FieldDeveloper, v))
}

// DeveloperLTE applies the LTE predicate on the "developer" field.
func DeveloperLTE(v string) predicate.App {
	return predicate.App(sql.FieldLTE(FieldDeveloper, v))
}

// DeveloperContains applies the Contains predicate on the "developer" field.
func DeveloperContains(v string) predicate.App {
	return predicate.App(sql.FieldContains(FieldDeveloper, v))
}

// DeveloperHasPrefix applies the HasPrefix predicate on the "developer" field.
func DeveloperHasPrefix(v string) predicate.App {
	return predicate.App(sql.FieldHasPrefix(FieldDeveloper, v))
}

// DeveloperHasSuffix applies the HasSuffix predicate on the "developer" field.
func DeveloperHasSuffix(v string) predicate.App {
	return predicate.App(sql.FieldHasSuffix(FieldDeveloper, v))
}

// DeveloperEqualFold applies the EqualFold predicate on the "developer" field.
func DeveloperEqualFold(v string) predicate.App {
	return predicate.App(sql.FieldEqualFold(FieldDeveloper, v))
}

// DeveloperContainsFold applies the ContainsFold predicate on the "developer" field.
func DeveloperContainsFold(v string) predicate.App {
	return predicate.App(sql.FieldContainsFold(FieldDeveloper, v))
}

// PublisherEQ applies the EQ predicate on the "publisher" field.
func PublisherEQ(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldPublisher, v))
}

// PublisherNEQ applies the NEQ predicate on the "publisher" field.
func PublisherNEQ(v string) predicate.App {
	return predicate.App(sql.FieldNEQ(FieldPublisher, v))
}

// PublisherIn applies the In predicate on the "publisher" field.
func PublisherIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldIn(FieldPublisher, vs...))
}

// PublisherNotIn applies the NotIn predicate on the "publisher" field.
func PublisherNotIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldNotIn(FieldPublisher, vs...))
}

// PublisherGT applies the GT predicate on the "publisher" field.
func PublisherGT(v string) predicate.App {
	return predicate.App(sql.FieldGT(FieldPublisher, v))
}

// PublisherGTE applies the GTE predicate on the "publisher" field.
func PublisherGTE(v string) predicate.App {
	return predicate.App(sql.FieldGTE(FieldPublisher, v))
}

// PublisherLT applies the LT predicate on the "publisher" field.
func PublisherLT(v string) predicate.App {
	return predicate.App(sql.FieldLT(FieldPublisher, v))
}

// PublisherLTE applies the LTE predicate on the "publisher" field.
func PublisherLTE(v string) predicate.App {
	return predicate.App(sql.FieldLTE(FieldPublisher, v))
}

// PublisherContains applies the Contains predicate on the "publisher" field.
func PublisherContains(v string) predicate.App {
	return predicate.App(sql.FieldContains(FieldPublisher, v))
}

// PublisherHasPrefix applies the HasPrefix predicate on the "publisher" field.
func PublisherHasPrefix(v string) predicate.App {
	return predicate.App(sql.FieldHasPrefix(FieldPublisher, v))
}

// PublisherHasSuffix applies the HasSuffix predicate on the "publisher" field.
func PublisherHasSuffix(v string) predicate.App {
	return predicate.App(sql.FieldHasSuffix(FieldPublisher, v))
}

// PublisherEqualFold applies the EqualFold predicate on the "publisher" field.
func PublisherEqualFold(v string) predicate.App {
	return predicate.App(sql.FieldEqualFold(FieldPublisher, v))
}

// PublisherContainsFold applies the ContainsFold predicate on the "publisher" field.
func PublisherContainsFold(v string) predicate.App {
	return predicate.App(sql.FieldContainsFold(FieldPublisher, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v string) predicate.App {
	return predicate.App(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v string) predicate.App {
	return predicate.App(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...string) predicate.App {
	return predicate.App(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v string) predicate.App {
	return predicate.App(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v string) predicate.App {
	return predicate.App(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v string) predicate.App {
	return predicate.App(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v string) predicate.App {
	return predicate.App(sql.FieldLTE(FieldVersion, v))
}

// VersionContains applies the Contains predicate on the "version" field.
func VersionContains(v string) predicate.App {
	return predicate.App(sql.FieldContains(FieldVersion, v))
}

// VersionHasPrefix applies the HasPrefix predicate on the "version" field.
func VersionHasPrefix(v string) predicate.App {
	return predicate.App(sql.FieldHasPrefix(FieldVersion, v))
}

// VersionHasSuffix applies the HasSuffix predicate on the "version" field.
func VersionHasSuffix(v string) predicate.App {
	return predicate.App(sql.FieldHasSuffix(FieldVersion, v))
}

// VersionEqualFold applies the EqualFold predicate on the "version" field.
func VersionEqualFold(v string) predicate.App {
	return predicate.App(sql.FieldEqualFold(FieldVersion, v))
}

// VersionContainsFold applies the ContainsFold predicate on the "version" field.
func VersionContainsFold(v string) predicate.App {
	return predicate.App(sql.FieldContainsFold(FieldVersion, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.App {
	return predicate.App(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.App {
	return predicate.App(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.App {
	return predicate.App(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.App {
	return predicate.App(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.App {
	return predicate.App(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.App {
	return predicate.App(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.App {
	return predicate.App(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.App {
	return predicate.App(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.App {
	return predicate.App(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.App {
	return predicate.App(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.App {
	return predicate.App(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.App {
	return predicate.App(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.App {
	return predicate.App(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.App {
	return predicate.App(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.App {
	return predicate.App(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.App {
	return predicate.App(sql.FieldLTE(FieldCreatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.App {
	return predicate.App(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UserTable, UserPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UserTable, UserPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAppPackage applies the HasEdge predicate on the "app_package" edge.
func HasAppPackage() predicate.App {
	return predicate.App(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AppPackageTable, AppPackageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAppPackageWith applies the HasEdge predicate on the "app_package" edge with a given conditions (other predicates).
func HasAppPackageWith(preds ...predicate.AppPackage) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AppPackageInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AppPackageTable, AppPackageColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.App) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.App) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.App) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		p(s.Not())
	})
}
