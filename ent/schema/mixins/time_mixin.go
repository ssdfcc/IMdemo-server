package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"time"
)

// TimeMixin holds the schema definition for the TimeMixin entity.
type TimeMixin struct {
	mixin.Schema
}

// Fields of the TimeMixin.
func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Immutable().Default(time.Now).Comment("创建时间").Annotations(entsql.WithComments(true)),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间").Annotations(entsql.WithComments(true)),
		field.Time("deleted_at").Optional().Nillable().Comment("删除时间").Annotations(entsql.WithComments(true)),
	}
}

// Edges of the TimeMixin.
func (TimeMixin) Edges() []ent.Edge {
	return nil
}

// Indexes of the TimeMixin.
func (TimeMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
		index.Fields("deleted_at"),
	}
}
