package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"go-zero-chat/ent/schema/mixins"
)

// GroupRelation holds the schema definition for the GroupRelation entity.
type GroupRelation struct {
	ent.Schema
}

// Mixin of the Contact.
func (GroupRelation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.TimeMixin{},
	}
}

// Fields of the GroupRelation.
// 人员关系
func (GroupRelation) Fields() []ent.Field {
	return []ent.Field{
		field.Int("group_id").Comment("群id").Annotations(entsql.WithComments(true)),
		field.Int("user_id").Comment("群成员id").Annotations(entsql.WithComments(true)),
		field.Int("type").Default(1).Comment("对应的类型 1普通成员 2管理员").Annotations(entsql.WithComments(true)),
		field.String("desc").Default("").Comment("描述相关").Annotations(entsql.WithComments(true)),
	}
}

// Edges of the GroupRelation.
func (GroupRelation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", Group.Type).Ref("group_relation_group").Unique().Field("group_id").Required(),
		edge.From("user", User.Type).Ref("group_relation_user").Unique().Field("user_id").Required(),
	}
}

// Annotations of the GroupRelation.
func (GroupRelation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_group_relation"},
	}
}
