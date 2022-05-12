// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PlatformsColumns holds the columns for the "platforms" table.
	PlatformsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "platform", Type: field.TypeString},
		{Name: "platform_auth_url", Type: field.TypeString},
		{Name: "logo_url", Type: field.TypeString},
		{Name: "platform_app_key", Type: field.TypeString},
		{Name: "platform_app_secret", Type: field.TypeString},
		{Name: "redirect_url", Type: field.TypeString},
	}
	// PlatformsTable holds the schema information for the "platforms" table.
	PlatformsTable = &schema.Table{
		Name:       "platforms",
		Columns:    PlatformsColumns,
		PrimaryKey: []*schema.Column{PlatformsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PlatformsTable,
	}
)

func init() {
}