package v1

import (
	"fmt"
	"net/http"
	"self_go_gin/api/v1/request"

	"self_go_gin/common/common_msg_id"
	"self_go_gin/handler"
	"self_go_gin/service"
	"self_go_gin/util/gin_response"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
)

// @Summary  Create Users
// @Description Create Users
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param request body swagger_docs.DocUsersCreate true "request body"
// @Success 200 {string} json "{"msg": {"success": "success"},"data": {}}"
// @Failure 400 {string} json "{"msg": {"fail": "帳密錯誤"},"data": null}"
// @Router /api/v1/auth/users [post]
func CreateUser(ctx *gin.Context) {
	fmt.Println("CreateUser")
	var data request.CreateUserRequest
	// var respData response.CreateUserResponse

	if err := ctx.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		check := handler.ValidCheckAndTrans(ctx, err)
		if check {
			gin_response.ErrorResponse(ctx, http.StatusBadRequest, "request_parameter_validation_failed", common_msg_id.Fail, nil)
			return
		}
		// 非validator.ValidationErrors類型錯誤直接傳回
		zap.L().Error("\n Api CreateUser() 失敗(ShouldBindBodyWith fail) : " + err.Error())
		gin_response.ErrorResponse(ctx, http.StatusNotFound, "invalid_request_parameters", common_msg_id.Fail, nil)
		return
	}

	userService := service.NewUserService()
	_, err := userService.CreateUser(data)
	ok, err := handler.HandlerError(ctx, err)
	if !ok {
		zap.L().Error("\n Api CreateUser() \n " + err.Error())
		return
	}
	gin_response.SuccessResponse(ctx, http.StatusOK, "", nil, common_msg_id.Success)
}

// @Summary  User Login
// @Description User Login
// @Tags Users
// @Accept json
// @Produce json
// @Param request body swagger_docs.DocUsersLogin true "request body"
// @Success 200 {string}  "成功"
// @Failure 400 {string}  "失敗"
// @Failure 401 {string}  "Unauthorized"
// @Router /api/v1/users/login [post]
func UserLogin(ctx *gin.Context) {
	var data request.UserLoginRequest
	if err := ctx.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		check := handler.ValidCheckAndTrans(ctx, err)
		if check {
			gin_response.ErrorResponse(ctx, http.StatusBadRequest, "request_parameter_validation_failed", common_msg_id.Fail, nil)
			return
		}
		// 非validator.ValidationErrors類型錯誤直接傳回
		zap.L().Error("\n Api UserLogin() 失敗(ShouldBindBodyWith fail) : " + err.Error())
		gin_response.ErrorResponse(ctx, http.StatusNotFound, "invalid_request_parameters", common_msg_id.Fail, nil)
		return
	}

	userService := service.NewUserService()
	jwtToken, err := userService.CheckLogin(data)
	ok, err := handler.HandlerError(ctx, err)
	if !ok {
		zap.L().Error("\n Api UserLogin() \n " + err.Error())
		return
	}
	gin_response.SuccessResponse(ctx, http.StatusOK, "", gin_response.CreateMsgData("jwt_token", *jwtToken), common_msg_id.Success)
}

// @Summary Get Users By ID
// @Description Get Users By ID
// @Tags Users
// @Accept json
// @Produce json
// @Security JwtTokenAuth
// @Param filterUsersId path string true "filterUsersId"
// @Success 200 {string}  "成功"
// @Failure 400 {string}  "失敗"
// @Failure 401 {string}  "Unauthorized"
// @Router /api/v1/auth/users/{filterUsersId} [get]
func GetUsersById(ctx *gin.Context) {
	var data request.GetUsersByIDRequest
	usersId, ok := ctx.Get("usersId")
	if ok {
		zap.S().Info("usersId :", usersId)
	}
	data.FilterUsersId = ctx.Param("filterUsersId")
	zap.S().Info("filterUsersId :", data)

	gin_response.SuccessResponse(ctx, http.StatusOK, "", nil, common_msg_id.Success)

}
