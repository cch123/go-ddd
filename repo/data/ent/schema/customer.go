// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Customer struct {
	ent.Schema
}

func (Customer) Fields() []ent.Field {
	return []ent.Field{field.Int("id"), field.String("addr"), field.Int32("age"), field.String("name"), field.Time("created_at"), field.Time("updated_at")}
}
func (Customer) Edges() []ent.Edge {
	return nil
}
func (Customer) Annotations() []schema.Annotation {
	return nil
}
