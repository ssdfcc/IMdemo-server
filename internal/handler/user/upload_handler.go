package user

import (
	"errors"
	"go-zero-chat/ent/file"
	"go-zero-chat/utils"
	"net/http"
	"path"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
	"go-zero-chat/internal/logic/user"
	"go-zero-chat/internal/svc"
	"go-zero-chat/internal/types"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}
		// 获取上传的文件（FormData）
		_, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}
		// 判断文件在数据库中是否已经存在
		hash := utils.Md5(strconv.FormatInt(fileHeader.Size, 10))
		rp, _ := svcCtx.DB.File.Query().Where(file.Hash(hash)).First(r.Context())
		if rp != nil && rp.ID > 0 {
			// 文件已经存在，直接返回信息
			xhttp.JsonBaseResponseCtx(r.Context(), w, &types.UploadResp{
				ID:   rp.ID,
				Name: rp.Name,
				Ext:  rp.Ext,
				Size: rp.Size,
				Path: rp.Path,
			})
			return
		}
		// 文件写入本地
		filepath, filename, err := utils.UploadFile(fileHeader, svcCtx.Config.Local.Path)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New("保存文件失败"))
			return
		}
		// 往 logic 传递 request
		req.Name = filename
		req.Ext = path.Ext(fileHeader.Filename)
		req.Size = fileHeader.Size
		req.Hash = hash
		req.Path = filepath
		l := user.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
