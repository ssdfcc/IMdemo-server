package contact

import (
	"net/http"

	xhttp "github.com/zeromicro/x/http"
	"go-zero-chat/internal/logic/contact"
	"go-zero-chat/internal/svc"
)

func SearchGroupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := contact.NewSearchGroupLogic(r.Context(), svcCtx)
		resp, err := l.SearchGroup()
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
