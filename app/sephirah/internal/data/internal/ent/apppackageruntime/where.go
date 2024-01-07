// Code generated by ent, DO NOT EDIT.

package apppackageruntime

import (
	"github.com/tuihub/librarian/model"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v model.InternalID) predicate.AppPackageRunTime {
	vc := int64(v)
	return predicate.AppPackageRunTime(sql.FieldEQ(FieldUserID, vc))
}

// AppPackageID applies equality check predicate on the "app_package_id" field. It's identical to AppPackageIDEQ.
func AppPackageID(v model.InternalID) predicate.AppPackageRunTime {
	vc := int64(v)
	return predicate.AppPackageRunTime(sql.FieldEQ(FieldAppPackageID, vc))
}

// StartTime applies equality check predicate on the "start_time" field. It's identical to StartTimeEQ.
func StartTime(v time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldEQ(FieldStartTime, v))
}

// RunDuration applies equality check predicate on the "run_duration" field. It's identical to RunDurationEQ.
func RunDuration(v time.Duration) predicate.AppPackageRunTime {
	vc := int64(v)
	return predicate.AppPackageRunTime(sql.FieldEQ(FieldRunDuration, vc))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldEQ(FieldCreatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v model.InternalID) predicate.AppPackageRunTime {
	vc := int64(v)
	return predicate.AppPackageRunTime(sql.FieldEQ(FieldUserID, vc))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v model.InternalID) predicate.AppPackageRunTime {
	vc := int64(v)
	return predicate.AppPackageRunTime(sql.FieldNEQ(FieldUserID, vc))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...model.InternalID) predicate.AppPackageRunTime {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.AppPackageRunTime(sql.FieldIn(FieldUserID, v...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...model.InternalID) predicate.AppPackageRunTime {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.AppPackageRunTime(sql.FieldNotIn(FieldUserID, v...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v model.InternalID) predicate.AppPackageRunTime {
	vc := int64(v)
	return predicate.AppPackageRunTime(sql.FieldGT(FieldUserID, vc))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v model.InternalID) predicate.AppPackageRunTime {
	vc := int64(v)
	return predicate.AppPackageRunTime(sql.FieldGTE(FieldUserID, vc))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v model.InternalID) predicate.AppPackageRunTime {
	vc := int64(v)
	return predicate.AppPackageRunTime(sql.FieldLT(FieldUserID, vc))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v model.InternalID) predicate.AppPackageRunTime {
	vc := int64(v)
	return predicate.AppPackageRunTime(sql.FieldLTE(FieldUserID, vc))
}

// AppPackageIDEQ applies the EQ predicate on the "app_package_id" field.
func AppPackageIDEQ(v model.InternalID) predicate.AppPackageRunTime {
	vc := int64(v)
	return predicate.AppPackageRunTime(sql.FieldEQ(FieldAppPackageID, vc))
}

// AppPackageIDNEQ applies the NEQ predicate on the "app_package_id" field.
func AppPackageIDNEQ(v model.InternalID) predicate.AppPackageRunTime {
	vc := int64(v)
	return predicate.AppPackageRunTime(sql.FieldNEQ(FieldAppPackageID, vc))
}

// AppPackageIDIn applies the In predicate on the "app_package_id" field.
func AppPackageIDIn(vs ...model.InternalID) predicate.AppPackageRunTime {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.AppPackageRunTime(sql.FieldIn(FieldAppPackageID, v...))
}

// AppPackageIDNotIn applies the NotIn predicate on the "app_package_id" field.
func AppPackageIDNotIn(vs ...model.InternalID) predicate.AppPackageRunTime {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.AppPackageRunTime(sql.FieldNotIn(FieldAppPackageID, v...))
}

// AppPackageIDGT applies the GT predicate on the "app_package_id" field.
func AppPackageIDGT(v model.InternalID) predicate.AppPackageRunTime {
	vc := int64(v)
	return predicate.AppPackageRunTime(sql.FieldGT(FieldAppPackageID, vc))
}

// AppPackageIDGTE applies the GTE predicate on the "app_package_id" field.
func AppPackageIDGTE(v model.InternalID) predicate.AppPackageRunTime {
	vc := int64(v)
	return predicate.AppPackageRunTime(sql.FieldGTE(FieldAppPackageID, vc))
}

// AppPackageIDLT applies the LT predicate on the "app_package_id" field.
func AppPackageIDLT(v model.InternalID) predicate.AppPackageRunTime {
	vc := int64(v)
	return predicate.AppPackageRunTime(sql.FieldLT(FieldAppPackageID, vc))
}

// AppPackageIDLTE applies the LTE predicate on the "app_package_id" field.
func AppPackageIDLTE(v model.InternalID) predicate.AppPackageRunTime {
	vc := int64(v)
	return predicate.AppPackageRunTime(sql.FieldLTE(FieldAppPackageID, vc))
}

// StartTimeEQ applies the EQ predicate on the "start_time" field.
func StartTimeEQ(v time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldEQ(FieldStartTime, v))
}

// StartTimeNEQ applies the NEQ predicate on the "start_time" field.
func StartTimeNEQ(v time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldNEQ(FieldStartTime, v))
}

// StartTimeIn applies the In predicate on the "start_time" field.
func StartTimeIn(vs ...time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldIn(FieldStartTime, vs...))
}

// StartTimeNotIn applies the NotIn predicate on the "start_time" field.
func StartTimeNotIn(vs ...time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldNotIn(FieldStartTime, vs...))
}

// StartTimeGT applies the GT predicate on the "start_time" field.
func StartTimeGT(v time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldGT(FieldStartTime, v))
}

// StartTimeGTE applies the GTE predicate on the "start_time" field.
func StartTimeGTE(v time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldGTE(FieldStartTime, v))
}

// StartTimeLT applies the LT predicate on the "start_time" field.
func StartTimeLT(v time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldLT(FieldStartTime, v))
}

// StartTimeLTE applies the LTE predicate on the "start_time" field.
func StartTimeLTE(v time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldLTE(FieldStartTime, v))
}

// RunDurationEQ applies the EQ predicate on the "run_duration" field.
func RunDurationEQ(v time.Duration) predicate.AppPackageRunTime {
	vc := int64(v)
	return predicate.AppPackageRunTime(sql.FieldEQ(FieldRunDuration, vc))
}

// RunDurationNEQ applies the NEQ predicate on the "run_duration" field.
func RunDurationNEQ(v time.Duration) predicate.AppPackageRunTime {
	vc := int64(v)
	return predicate.AppPackageRunTime(sql.FieldNEQ(FieldRunDuration, vc))
}

// RunDurationIn applies the In predicate on the "run_duration" field.
func RunDurationIn(vs ...time.Duration) predicate.AppPackageRunTime {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.AppPackageRunTime(sql.FieldIn(FieldRunDuration, v...))
}

// RunDurationNotIn applies the NotIn predicate on the "run_duration" field.
func RunDurationNotIn(vs ...time.Duration) predicate.AppPackageRunTime {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.AppPackageRunTime(sql.FieldNotIn(FieldRunDuration, v...))
}

// RunDurationGT applies the GT predicate on the "run_duration" field.
func RunDurationGT(v time.Duration) predicate.AppPackageRunTime {
	vc := int64(v)
	return predicate.AppPackageRunTime(sql.FieldGT(FieldRunDuration, vc))
}

// RunDurationGTE applies the GTE predicate on the "run_duration" field.
func RunDurationGTE(v time.Duration) predicate.AppPackageRunTime {
	vc := int64(v)
	return predicate.AppPackageRunTime(sql.FieldGTE(FieldRunDuration, vc))
}

// RunDurationLT applies the LT predicate on the "run_duration" field.
func RunDurationLT(v time.Duration) predicate.AppPackageRunTime {
	vc := int64(v)
	return predicate.AppPackageRunTime(sql.FieldLT(FieldRunDuration, vc))
}

// RunDurationLTE applies the LTE predicate on the "run_duration" field.
func RunDurationLTE(v time.Duration) predicate.AppPackageRunTime {
	vc := int64(v)
	return predicate.AppPackageRunTime(sql.FieldLTE(FieldRunDuration, vc))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppPackageRunTime) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppPackageRunTime) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AppPackageRunTime) predicate.AppPackageRunTime {
	return predicate.AppPackageRunTime(sql.NotPredicates(p))
}
