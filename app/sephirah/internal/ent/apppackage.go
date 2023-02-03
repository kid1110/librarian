// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/apppackage"
)

// AppPackage is the model entity for the AppPackage schema.
type AppPackage struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// InternalID holds the value of the "internal_id" field.
	InternalID int64 `json:"internal_id,omitempty"`
	// Source holds the value of the "source" field.
	Source apppackage.Source `json:"source,omitempty"`
	// SourceID holds the value of the "source_id" field.
	SourceID int64 `json:"source_id,omitempty"`
	// SourcePackageID holds the value of the "source_package_id" field.
	SourcePackageID string `json:"source_package_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// BinaryName holds the value of the "binary_name" field.
	BinaryName string `json:"binary_name,omitempty"`
	// BinarySize holds the value of the "binary_size" field.
	BinarySize string `json:"binary_size,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppPackage) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case apppackage.FieldID, apppackage.FieldInternalID, apppackage.FieldSourceID:
			values[i] = new(sql.NullInt64)
		case apppackage.FieldSource, apppackage.FieldSourcePackageID, apppackage.FieldName, apppackage.FieldDescription, apppackage.FieldBinaryName, apppackage.FieldBinarySize:
			values[i] = new(sql.NullString)
		case apppackage.FieldUpdatedAt, apppackage.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppPackage", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppPackage fields.
func (ap *AppPackage) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case apppackage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ap.ID = int(value.Int64)
		case apppackage.FieldInternalID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field internal_id", values[i])
			} else if value.Valid {
				ap.InternalID = value.Int64
			}
		case apppackage.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				ap.Source = apppackage.Source(value.String)
			}
		case apppackage.FieldSourceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field source_id", values[i])
			} else if value.Valid {
				ap.SourceID = value.Int64
			}
		case apppackage.FieldSourcePackageID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source_package_id", values[i])
			} else if value.Valid {
				ap.SourcePackageID = value.String
			}
		case apppackage.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ap.Name = value.String
			}
		case apppackage.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ap.Description = value.String
			}
		case apppackage.FieldBinaryName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field binary_name", values[i])
			} else if value.Valid {
				ap.BinaryName = value.String
			}
		case apppackage.FieldBinarySize:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field binary_size", values[i])
			} else if value.Valid {
				ap.BinarySize = value.String
			}
		case apppackage.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ap.UpdatedAt = value.Time
			}
		case apppackage.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ap.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AppPackage.
// Note that you need to call AppPackage.Unwrap() before calling this method if this AppPackage
// was returned from a transaction, and the transaction was committed or rolled back.
func (ap *AppPackage) Update() *AppPackageUpdateOne {
	return NewAppPackageClient(ap.config).UpdateOne(ap)
}

// Unwrap unwraps the AppPackage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ap *AppPackage) Unwrap() *AppPackage {
	_tx, ok := ap.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppPackage is not a transactional entity")
	}
	ap.config.driver = _tx.drv
	return ap
}

// String implements the fmt.Stringer.
func (ap *AppPackage) String() string {
	var builder strings.Builder
	builder.WriteString("AppPackage(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ap.ID))
	builder.WriteString("internal_id=")
	builder.WriteString(fmt.Sprintf("%v", ap.InternalID))
	builder.WriteString(", ")
	builder.WriteString("source=")
	builder.WriteString(fmt.Sprintf("%v", ap.Source))
	builder.WriteString(", ")
	builder.WriteString("source_id=")
	builder.WriteString(fmt.Sprintf("%v", ap.SourceID))
	builder.WriteString(", ")
	builder.WriteString("source_package_id=")
	builder.WriteString(ap.SourcePackageID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ap.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(ap.Description)
	builder.WriteString(", ")
	builder.WriteString("binary_name=")
	builder.WriteString(ap.BinaryName)
	builder.WriteString(", ")
	builder.WriteString("binary_size=")
	builder.WriteString(ap.BinarySize)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ap.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ap.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// AppPackages is a parsable slice of AppPackage.
type AppPackages []*AppPackage

func (ap AppPackages) config(cfg config) {
	for _i := range ap {
		ap[_i].config = cfg
	}
}
