package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Box holds the schema definition for the Box entity.
type Box struct {
	ent.Schema
}

// Fields of the Box.
func (Box) Fields() []ent.Field {
	return []ent.Field{
		field.String("barcode"),
		field.Int("quantity").Positive(),
		field.Time("expire_date"),
	}
}

// Edges of the Box.
func (Box) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("med", Medicine.Type).
			Ref("in_this_box").
			Unique(),
	}
}
