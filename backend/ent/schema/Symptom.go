package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Symptom holds the schema definition for the Symptom entity.
type Symptom struct {
	ent.Schema
}

// Fields of the Symptom.
func (Symptom) Fields() []ent.Field {
	return []ent.Field{
		field.Int("Symptom_ID").Unique(),
		field.String("Symptom_name").NotEmpty(),
	}

}

// Edges of the Symptom.
func (Symptom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("symptoms", RepairInvoice.Type),
	}
}
