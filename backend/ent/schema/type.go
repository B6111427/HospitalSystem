package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Type holds the schema definition for the Type entity.
type Type struct {
	ent.Schema
}

// Fields of the Type.
func (Type) Fields() []ent.Field {
	return []ent.Field{
		field.String("type_name").Unique(),
	}
}

// Edges of the Type.
func (Type) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("med_type", Medicine.Type).
			StorageKey(edge.Column("type_id")),
	}
}
