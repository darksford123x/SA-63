package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// RepairInvoice holds the schema definition for the RepairInvoice entity.
type RepairInvoice struct {
	ent.Schema
}

// Fields of the RepairInvoice.
func (RepairInvoice) Fields() []ent.Field {
	return []ent.Field{
		field.Int("RepairInvoice_ID").Unique(),
		field.Int("Status_ID").Unique(),
		field.Int("Device_ID").Unique(),
		field.Int("Symptom_ID").Unique(),
	}
}

// Edges of the RepairInvoice.
func (RepairInvoice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("device", Device.Type).
			Ref("devices").
			Unique().
			Required(),

		edge.From("statusinvoice", Status.Type).
			Ref("statusInvoice"),

		edge.From("symptom", Symptom.Type).
			Ref("symptoms"),
	}
}
