package v1

import (
	"net/http"
	"self_go_gin/domains/admin/service"
	"self_go_gin/common/common_msg_id"
	"self_go_gin/gin_application/handler"
	"self_go_gin/util/gin_response"
	"self_go_gin/gin_application/api/v1/admin/request"
	"self_go_gin/gin_application/api/v1/admin/response"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
)

// @Summary  Create Admins
// @Description Create Admins
// @Tags Admins
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param request body swagger_docs.DocAdminsCreate true "request body"
// @Success 200 {string} json "{"msg": {"success": "success"},"data": {}}"
// @Failure 400 {string} json "{"msg": {"fail": "帳密錯誤"},"data": null}"
// @Router /api/v1/auth/admins [post]
func CreateAdmin(ctx *gin.Context) {
	var data request.CreateAdminRequest
	if err := ctx.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		check := handler.ValidCheckAndTrans(ctx, err)
		if check {
			gin_response.ErrorResponse(ctx, http.StatusBadRequest, "request_parameter_validation_failed", common_msg_id.Fail, nil)
			return
		}
		// 非validator.ValidationErrors類型錯誤直接傳回
		zap.L().Error("\n Api CreateAdmin() 失敗(ShouldBindBodyWith fail) : " + err.Error())
		gin_response.ErrorResponse(ctx, http.StatusNotFound, "invalid_request_parameters", common_msg_id.Fail, nil)
		return
	}

	adminService := service.NewAdminService()
	admin, err := adminService.CreateAdmin(data)
	ok, err := handler.HandlerError(ctx, err)
	if !ok {
		zap.L().Error("\n Api CreateUser() \n " + err.Error())
		return
	}

	respData := response.CreateAdminResponse{
		AdminId: admin.ID,
		Account: admin.Account,
	}
	gin_response.SuccessResponse(ctx, http.StatusOK, "", respData, common_msg_id.Success)
}

// @Summary  Admin Login
// @Description Admin Login
// @Tags Admins
// @Accept json
// @Produce json
// @Param request body swagger_docs.DocAdminsLogin true "request body"
// @Success 200 {string}  "成功"
// @Failure 400 {string}  "失敗"
// @Failure 401 {string}  "Unauthorized"
// @Router /api/v1/admins/login [post]
func AdminLogin(ctx *gin.Context) {
	var data request.AdminLoginRequest

	if err := ctx.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		check := handler.ValidCheckAndTrans(ctx, err)
		if check {
			gin_response.ErrorResponse(ctx, http.StatusBadRequest, "request_parameter_validation_failed", common_msg_id.Fail, nil)
			return
		}
		// 非validator.ValidationErrors類型錯誤直接傳回
		zap.L().Error("\n Api AdminLogin() 失敗(ShouldBindBodyWith fail) : " + err.Error())
		gin_response.ErrorResponse(ctx, http.StatusNotFound, "invalid_request_parameters", common_msg_id.Fail, nil)
		return
	}

	adminService := service.NewAdminService()
	jwtToken, err := adminService.CheckLogin(data)
	ok, err := handler.HandlerError(ctx, err)
	if !ok {
		zap.L().Error("\n Api AdminLogin() \n " + err.Error())
		return
	}
	gin_response.SuccessResponse(ctx, http.StatusOK, "", gin_response.CreateMsgData("jwt_token", *jwtToken), common_msg_id.Success)

}

// @Summary Get Admins By ID
// @Description Get Admins By ID
// @Tags Admins
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param filterAdminsId path string true "filterAdminsId"
// @Success 200 {string}  "成功"
// @Failure 400 {string}  "失敗"
// @Failure 401 {string}  "Unauthorized"
// @Router /api/v1/auth/admins/{filterAdminsId} [get]
func GetAdminsById(ctx *gin.Context) {
	var data request.GetAdminsByIDRequest

	admin_id, _ := ctx.Get("adminId")
	zap.S().Info("admin_id :", admin_id)
	data.FilterAdminsId = ctx.Param("filterAdminsId")
	zap.S().Info("FilterAdminsId :", data)

	gin_response.SuccessResponse(ctx, http.StatusOK, "", nil, common_msg_id.Success)
}
