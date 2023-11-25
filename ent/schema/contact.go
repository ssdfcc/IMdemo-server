package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"go-zero-chat/ent/schema/mixins"
)

// Contact holds the schema definition for the Contact entity.
type Contact struct {
	ent.Schema
}

// Mixin of the Contact.
func (Contact) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.TimeMixin{},
	}
}

// Fields of the Contact.
// 人员关系
func (Contact) Fields() []ent.Field {
	return []ent.Field{
		field.Int("owner_id").Comment("谁的关系").Annotations(entsql.WithComments(true)),
		field.Int("target_id").Comment("对应谁").Annotations(entsql.WithComments(true)),
		field.Int("type").Default(1).Comment("对应的类型 1好友").Annotations(entsql.WithComments(true)),
		field.String("desc").Default("").Comment("描述相关").Annotations(entsql.WithComments(true)),
	}
}

// Edges of the Contact.
func (Contact) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner_user", User.Type).Ref("contact_owner").Unique().Field("owner_id").Required(),
		edge.From("target_user", User.Type).Ref("contact_target").Unique().Field("target_id").Required(),
	}
}

// Annotations of the Contact.
func (Contact) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_contact"},
	}
}
