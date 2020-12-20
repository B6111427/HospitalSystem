package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Medicine holds the schema definition for the Medicine entity.
type Medicine struct {
	ent.Schema
}

// Fields of the Medicine.
func (Medicine) Fields() []ent.Field {
	return []ent.Field{
		field.String("medicine_name").Unique(),
	}
}

// Edges of the Medicine.
func (Medicine) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("in_this_box", Box.Type).
			StorageKey(edge.Column("med_id")),
		edge.From("type", Type.Type).
			Ref("med_type").
			Unique(),
	}
}
