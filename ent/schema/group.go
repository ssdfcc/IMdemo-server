package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"go-zero-chat/ent/schema/mixins"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Mixin of the Group.
func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.TimeMixin{},
	}
}

// Fields of the Group.
// 群
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("群名").Annotations(entsql.WithComments(true)),
		field.Int("user_id").Comment("所有者").Annotations(entsql.WithComments(true)),
		field.String("img").Default("uploads/file/default.png").Comment("图片").Annotations(entsql.WithComments(true)),
		field.String("desc").Default("").Comment("描述相关").Annotations(entsql.WithComments(true)),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner_user", User.Type).Ref("group_owner").Unique().Field("user_id").Required(),
		edge.To("group_relation_group", GroupRelation.Type),
		edge.To("target_msg", GroupMsg.Type),
	}
}

// Annotations of the Group.
func (Group) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_group"},
	}
}
