package contact

import (
	"net/http"

	xhttp "github.com/zeromicro/x/http"
	"go-zero-chat/internal/logic/contact"
	"go-zero-chat/internal/svc"
)

func SearchFriendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := contact.NewSearchFriendLogic(r.Context(), svcCtx)
		resp, err := l.SearchFriend()
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
