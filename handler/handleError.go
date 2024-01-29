package handler

import (
	"api/common/common_msg_id"
	"api/util/gin_response"
	"api/util/mysql_manager"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func HandlerError(context *gin.Context, msgId common_msg_id.MsgId, errMsg string, err error, statusCode int,errData interface{}) {
	if mysql_manager.MysqlErrCode(err) == mysql_manager.DuplicateEntryCode {
		gin_response.ErrorResponse(context, http.StatusBadRequest, "", common_msg_id.Duplicate_Entry, nil)
		return
	}
	if err == gorm.ErrRecordNotFound {
		gin_response.ErrorResponse(context, http.StatusBadRequest, "", common_msg_id.No_Content, nil)
		return
	}
	if err != nil {
		zap.L().Error(errMsg + " : " + err.Error())
	} else {
		zap.L().Error(errMsg)
	}

	gin_response.ErrorResponse(context, statusCode, "", msgId, errData)
	return
}
