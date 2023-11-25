package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"go-zero-chat/ent/schema/mixins"
)

// GroupMsg holds the schema definition for the GroupMsg entity.
type GroupMsg struct {
	ent.Schema
}

// Mixin of the GroupMsg.
func (GroupMsg) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.TimeMixin{},
	}
}

// Fields of the GroupMsg.
// 群信息
func (GroupMsg) Fields() []ent.Field {
	return []ent.Field{
		field.Int("owner_id").Comment("发送者id").Annotations(entsql.WithComments(true)),
		field.Int("target_id").Comment("群id").Annotations(entsql.WithComments(true)),
		field.Int("media").Comment("消息类型 1文字 2表情包 3图片 4音频 5视频").Annotations(entsql.WithComments(true)),
		field.String("content").Default("").Comment("消息内容").Annotations(entsql.WithComments(true)),
		field.String("pic").Default("").Comment("图片相关").Annotations(entsql.WithComments(true)),
		field.String("url").Default("").Comment("url相关").Annotations(entsql.WithComments(true)),
		field.String("desc").Default("").Comment("描述相关").Annotations(entsql.WithComments(true)),
		field.Int("amount").Comment("其他数字统计").Annotations(entsql.WithComments(true)),
		field.Int("type").Comment("类型").Annotations(entsql.WithComments(true)),
	}
}

// Edges of the GroupMsg.
func (GroupMsg) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("group_msg").Unique().Field("owner_id").Required(),
		edge.From("group", Group.Type).Ref("target_msg").Unique().Field("target_id").Required(),
	}
}

// Annotations of the GroupMsg.
func (GroupMsg) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_group_msg"},
	}
}
