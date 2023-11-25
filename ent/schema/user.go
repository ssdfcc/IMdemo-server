package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"go-zero-chat/ent/schema/mixins"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.TimeMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().Comment("用户名").Annotations(entsql.WithComments(true)),
		field.String("password").Comment("用户密码").Annotations(entsql.WithComments(true)),
		field.String("avatar").Default("uploads/file/default.png").Comment("头像").Annotations(entsql.WithComments(true)),
		field.String("phone").Default("").Comment("电话号码").Annotations(entsql.WithComments(true)),
		field.String("email").Default("").Comment("邮箱").Annotations(entsql.WithComments(true)),
		field.String("client_ip").Default("").Comment("客户端ip").Annotations(entsql.WithComments(true)),
		field.String("client_port").Default("").Comment("客户端ip").Annotations(entsql.WithComments(true)),
		field.String("salt").Default("").Comment("加密盐").Annotations(entsql.WithComments(true)),
		field.Time("login_time").Optional().Nillable().Comment("登录时间").Annotations(entsql.WithComments(true)),
		field.Time("heartbeat_time").Optional().Nillable().Comment("心跳时间").Annotations(entsql.WithComments(true)),
		field.Time("logout_time").Optional().Nillable().Comment("下线时间").Annotations(entsql.WithComments(true)),
		field.Bool("is_logout").Default(true).Comment("是否下线").Annotations(entsql.WithComments(true)),
		field.String("device_info").Default("").Comment("设备信息").Annotations(entsql.WithComments(true)),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// 消息
		edge.To("form_msg", Message.Type),
		edge.To("target_msg", Message.Type),
		// 人员关系
		edge.To("contact_owner", Contact.Type),
		edge.To("contact_target", Contact.Type),
		edge.To("group_relation_user", GroupRelation.Type),
		// 群
		edge.To("group_owner", Group.Type),
		edge.To("group_msg", GroupMsg.Type),
		// 文件
		edge.To("files", File.Type),
	}
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_user"},
	}
}
