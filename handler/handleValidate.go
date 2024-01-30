package handler

import (
	"api/common/common_msg_id"
	"api/initialize"
	"api/util/gin_response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidCheckAndTrans(context *gin.Context, err error) (ok bool) {
	validTrans, ok := initialize.ErrorValidateCheckAndTrans(err)
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
