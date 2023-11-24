package gin_response

import "github.com/gin-gonic/gin"

type Response struct {
	Msg  map[string]string      `json:"msg"`
	Data interface{} `json:"data"`
}

func CreateMsg(key,value string) map[string]string {
	var msg = make(map[string]string)
	msg[key] = value
	return msg
}

func SuccessResponse(c *gin.Context, statusCode int, msg map[string]string, data interface{}) {
	c.JSON(statusCode, Response{
		Msg:  msg,
		Data: data,
	})

	return
}

func ErrorResponse(c *gin.Context, statusCode int, msg map[string]string) {
	c.JSON(statusCode, Response{
		Msg:  msg,
		Data: nil,
	})

	return
}
