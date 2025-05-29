package handler

import (
	"errors"
	"fmt"
	"net/http"
	"self_go_gin/common/common_msg_id"
	"self_go_gin/util/gin_response"
	"self_go_gin/util/mysql_manager"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandlerError(context *gin.Context, err error) (bool, error) {
	switch {
	case mysql_manager.MysqlErrCode(err) == mysql_manager.DuplicateEntryCode:
		gin_response.ErrorResponse(context, http.StatusBadRequest, "", common_msg_id.Duplicate_Entry, nil)
		return false, fmt.Errorf("HandlerError() DuplicateEntryCode : %w", err)
	case errors.Is(err, gorm.ErrRecordNotFound):
		gin_response.ErrorResponse(context, http.StatusNotFound, "", common_msg_id.No_Content, nil)
		return false, fmt.Errorf("HandlerError() ErrRecordNotFound : %w", err)
	case errors.Is(err, ErrResourceExist):
		gin_response.ErrorResponse(context, http.StatusBadRequest, "", common_msg_id.Fail, nil)
		return false, fmt.Errorf("HandlerError() ErrResourceExist : %w", err)
	case err != nil:
		gin_response.ErrorResponse(context, http.StatusInternalServerError, "", common_msg_id.Fail, nil)
		return false, fmt.Errorf("HandlerError() : %w", err)
	default:
		return true, nil
	}
}
