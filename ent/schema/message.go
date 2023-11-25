package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"go-zero-chat/ent/schema/mixins"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Mixin of the Message.
func (Message) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.TimeMixin{},
	}
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Int("type").Comment("发送类型 1私聊 2群聊 3广播").Annotations(entsql.WithComments(true)),
		field.Int("media").Comment("消息类型 1文字 2表情包 3图片 4音频 5视频").Annotations(entsql.WithComments(true)),
		field.String("content").Default("").Comment("消息内容").Annotations(entsql.WithComments(true)),
		field.String("pic").Default("").Comment("图片相关").Annotations(entsql.WithComments(true)),
		field.String("url").Default("").Comment("url相关").Annotations(entsql.WithComments(true)),
		field.String("desc").Default("").Comment("描述相关").Annotations(entsql.WithComments(true)),
		field.Int("amount").Comment("其他数字统计").Annotations(entsql.WithComments(true)),
		field.Int("form_id").Comment("发送者id").Annotations(entsql.WithComments(true)),
		field.Int("target_id").Comment("接收者id").Annotations(entsql.WithComments(true)),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sender", User.Type).Ref("form_msg").Unique().Field("form_id").Required(),
		edge.From("recipient", User.Type).Ref("target_msg").Unique().Field("target_id").Required(),
	}
}

// Annotations of the Message.
func (Message) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_message"},
	}
}
