package handler

import (
	"net/http"
	"self_go_gin/common/common_msg_id"
	"self_go_gin/gin_application/validate_lang"
	"self_go_gin/util/gin_response"

	"github.com/gin-gonic/gin"
)

func ValidCheckAndTrans(context *gin.Context, err error) (ok bool) {
	validTrans, ok := validate_lang.ErrorValidateCheckAndTrans(err)
	if ok {
		var validTransMsgData = make(map[int]string)
		var i int
		for _, v := range validTrans {
			validTransMsgData[i] = v
			i++
		}

		gin_response.ErrorResponse(context, http.StatusBadRequest, "validate fail", common_msg_id.Fail, validTransMsgData)
		return ok
	}
	return false
}
