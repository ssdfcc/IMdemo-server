package handler

import (
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"strings"
)

func dirHandler(pattern, dirPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		handler := http.StripPrefix(pattern, http.FileServer(http.Dir(dirPath)))
		handler.ServeHTTP(w, req)

	}
}

// 添加静态资源路由
func RegisterStaticHandlers(engine *rest.Server) {
	//这里注册
	dirLevel := []string{":1", ":2", ":3", ":4", ":5", ":6", ":7", ":8"}
	pattern := "/uploads/"
	dirPath := "./uploads/"
	for i := 1; i < len(dirLevel); i++ {
		path := pattern + strings.Join(dirLevel[:i], "/")
		// 添加进路由
		engine.AddRoute(
			rest.Route{
				Method:  http.MethodGet,
				Path:    path,
				Handler: dirHandler(pattern, dirPath),
			})
	}
}
