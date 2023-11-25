package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"go-zero-chat/ent/schema/mixins"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Mixin of the File.
func (File) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.TimeMixin{},
	}
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.String("hash").Comment("哈希值").Annotations(entsql.WithComments(true)),
		field.String("name").Comment("文件名").Annotations(entsql.WithComments(true)),
		field.String("ext").Comment("文件格式").Annotations(entsql.WithComments(true)),
		field.Int64("size").Comment("文件大小").Annotations(entsql.WithComments(true)),
		field.String("path").Comment("文件路径").Annotations(entsql.WithComments(true)),
		field.Int("user_id").Comment("上传人id").Annotations(entsql.WithComments(true)),
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("files").Unique().Field("user_id").Required(),
	}
}

// Annotations of the File.
func (File) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_file"},
	}
}
