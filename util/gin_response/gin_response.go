package gin_response

import (
	"self_go_gin/common/common_msg_id"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Result common_msg_id.MsgId `json:"result" binding:"required"`
	Msg    string              `json:"msg"`
	Data   interface{}         `json:"data"`
}

func CreateMsgData(key, value string) map[string]string {
	var msg = make(map[string]string)
	msg[key] = value
	return msg
}

func SuccessResponse(c *gin.Context, statusCode int, msg string, data interface{}, result common_msg_id.MsgId) {
	c.JSON(statusCode, Response{
		Result: result,
		Msg:    msg,
		Data:   data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, msg string, result common_msg_id.MsgId, errData interface{}) {
	c.JSON(statusCode, Response{
		Result: result,
		Msg:    msg,
		Data:   errData,
	})
}
