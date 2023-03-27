// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/app"
	"github.com/tuihub/librarian/internal/model"
)

// App is the model entity for the App schema.
type App struct {
	config `json:"-"`
	// ID of the ent.
	ID model.InternalID `json:"id,omitempty"`
	// Source holds the value of the "source" field.
	Source app.Source `json:"source,omitempty"`
	// SourceAppID holds the value of the "source_app_id" field.
	SourceAppID string `json:"source_app_id,omitempty"`
	// SourceURL holds the value of the "source_url" field.
	SourceURL string `json:"source_url,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type app.Type `json:"type,omitempty"`
	// ShortDescription holds the value of the "short_description" field.
	ShortDescription string `json:"short_description,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// ImageURL holds the value of the "image_url" field.
	ImageURL string `json:"image_url,omitempty"`
	// ReleaseDate holds the value of the "release_date" field.
	ReleaseDate string `json:"release_date,omitempty"`
	// Developer holds the value of the "developer" field.
	Developer string `json:"developer,omitempty"`
	// Publisher holds the value of the "publisher" field.
	Publisher string `json:"publisher,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AppQuery when eager-loading is set.
	Edges             AppEdges `json:"edges"`
	app_bind_external *model.InternalID
}

// AppEdges holds the relations/edges for other nodes in the graph.
type AppEdges struct {
	// PurchasedBy holds the value of the purchased_by edge.
	PurchasedBy []*User `json:"purchased_by,omitempty"`
	// AppPackage holds the value of the app_package edge.
	AppPackage []*AppPackage `json:"app_package,omitempty"`
	// BindInternal holds the value of the bind_internal edge.
	BindInternal *App `json:"bind_internal,omitempty"`
	// BindExternal holds the value of the bind_external edge.
	BindExternal []*App `json:"bind_external,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// PurchasedByOrErr returns the PurchasedBy value or an error if the edge
// was not loaded in eager-loading.
func (e AppEdges) PurchasedByOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.PurchasedBy, nil
	}
	return nil, &NotLoadedError{edge: "purchased_by"}
}

// AppPackageOrErr returns the AppPackage value or an error if the edge
// was not loaded in eager-loading.
func (e AppEdges) AppPackageOrErr() ([]*AppPackage, error) {
	if e.loadedTypes[1] {
		return e.AppPackage, nil
	}
	return nil, &NotLoadedError{edge: "app_package"}
}

// BindInternalOrErr returns the BindInternal value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AppEdges) BindInternalOrErr() (*App, error) {
	if e.loadedTypes[2] {
		if e.BindInternal == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: app.Label}
		}
		return e.BindInternal, nil
	}
	return nil, &NotLoadedError{edge: "bind_internal"}
}

// BindExternalOrErr returns the BindExternal value or an error if the edge
// was not loaded in eager-loading.
func (e AppEdges) BindExternalOrErr() ([]*App, error) {
	if e.loadedTypes[3] {
		return e.BindExternal, nil
	}
	return nil, &NotLoadedError{edge: "bind_external"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*App) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case app.FieldID:
			values[i] = new(sql.NullInt64)
		case app.FieldSource, app.FieldSourceAppID, app.FieldSourceURL, app.FieldName, app.FieldType, app.FieldShortDescription, app.FieldDescription, app.FieldImageURL, app.FieldReleaseDate, app.FieldDeveloper, app.FieldPublisher, app.FieldVersion:
			values[i] = new(sql.NullString)
		case app.FieldUpdatedAt, app.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case app.ForeignKeys[0]: // app_bind_external
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type App", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the App fields.
func (a *App) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case app.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				a.ID = model.InternalID(value.Int64)
			}
		case app.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				a.Source = app.Source(value.String)
			}
		case app.FieldSourceAppID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source_app_id", values[i])
			} else if value.Valid {
				a.SourceAppID = value.String
			}
		case app.FieldSourceURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source_url", values[i])
			} else if value.Valid {
				a.SourceURL = value.String
			}
		case app.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case app.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				a.Type = app.Type(value.String)
			}
		case app.FieldShortDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field short_description", values[i])
			} else if value.Valid {
				a.ShortDescription = value.String
			}
		case app.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				a.Description = value.String
			}
		case app.FieldImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image_url", values[i])
			} else if value.Valid {
				a.ImageURL = value.String
			}
		case app.FieldReleaseDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field release_date", values[i])
			} else if value.Valid {
				a.ReleaseDate = value.String
			}
		case app.FieldDeveloper:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field developer", values[i])
			} else if value.Valid {
				a.Developer = value.String
			}
		case app.FieldPublisher:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field publisher", values[i])
			} else if value.Valid {
				a.Publisher = value.String
			}
		case app.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				a.Version = value.String
			}
		case app.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case app.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case app.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field app_bind_external", values[i])
			} else if value.Valid {
				a.app_bind_external = new(model.InternalID)
				*a.app_bind_external = model.InternalID(value.Int64)
			}
		}
	}
	return nil
}

// QueryPurchasedBy queries the "purchased_by" edge of the App entity.
func (a *App) QueryPurchasedBy() *UserQuery {
	return NewAppClient(a.config).QueryPurchasedBy(a)
}

// QueryAppPackage queries the "app_package" edge of the App entity.
func (a *App) QueryAppPackage() *AppPackageQuery {
	return NewAppClient(a.config).QueryAppPackage(a)
}

// QueryBindInternal queries the "bind_internal" edge of the App entity.
func (a *App) QueryBindInternal() *AppQuery {
	return NewAppClient(a.config).QueryBindInternal(a)
}

// QueryBindExternal queries the "bind_external" edge of the App entity.
func (a *App) QueryBindExternal() *AppQuery {
	return NewAppClient(a.config).QueryBindExternal(a)
}

// Update returns a builder for updating this App.
// Note that you need to call App.Unwrap() before calling this method if this App
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *App) Update() *AppUpdateOne {
	return NewAppClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the App entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *App) Unwrap() *App {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: App is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *App) String() string {
	var builder strings.Builder
	builder.WriteString("App(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("source=")
	builder.WriteString(fmt.Sprintf("%v", a.Source))
	builder.WriteString(", ")
	builder.WriteString("source_app_id=")
	builder.WriteString(a.SourceAppID)
	builder.WriteString(", ")
	builder.WriteString("source_url=")
	builder.WriteString(a.SourceURL)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", a.Type))
	builder.WriteString(", ")
	builder.WriteString("short_description=")
	builder.WriteString(a.ShortDescription)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(a.Description)
	builder.WriteString(", ")
	builder.WriteString("image_url=")
	builder.WriteString(a.ImageURL)
	builder.WriteString(", ")
	builder.WriteString("release_date=")
	builder.WriteString(a.ReleaseDate)
	builder.WriteString(", ")
	builder.WriteString("developer=")
	builder.WriteString(a.Developer)
	builder.WriteString(", ")
	builder.WriteString("publisher=")
	builder.WriteString(a.Publisher)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(a.Version)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Apps is a parsable slice of App.
type Apps []*App
