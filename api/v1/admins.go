package v1

import (
	"api/common/common_const"
	"api/common/common_msg_id"
	"api/handler"
	"api/model"
	"api/util/bcryptEncap"
	"api/util/gin_response"
	"api/util/jwt_secret"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
	"gorm.io/gorm"
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
func CreateAdmin(context *gin.Context) {
	type receiveData struct {
		Account  string `json:"account" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	type responseData struct {
		AdminId uint   `json:"id"`
		Account string `json:"account"`
	}

	var data receiveData
	var newAdmins model.Admins
	var respData responseData

	if err := context.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		check := handler.ValidCheckAndTrans(context, err)
		if check {
			return
		}
		// 非validator.ValidationErrors類型錯誤直接傳回
		handler.HandlerError(context, common_msg_id.Fail, "CreateUser() BindJSON fail", err, http.StatusBadRequest, nil)
		return
	}

	_, err := newAdmins.GetAdminsByAccount(data.Account)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		gin_response.ErrorResponse(context, http.StatusBadRequest, "db GetUsersByAccount fail", common_msg_id.Fail, err)
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		//密碼加密
		bcryptPassword, err := bcryptEncap.GenerateFromPassword(data.Password)
		if err != nil {
			handler.HandlerError(context, common_msg_id.Fail, "CreateUser() bcrypt fail", err, http.StatusBadRequest, nil)
			return
		}
		newAdmins.Account = data.Account
		newAdmins.Password = string(bcryptPassword)
		admins, err := newAdmins.CreateAdmin()
		if err != nil {
			handler.HandlerError(context, common_msg_id.Fail, "CreateUser() model CreateUser() fail", err, http.StatusBadRequest, nil)
			return
		}
		respData = responseData{
			AdminId: admins.ID,
			Account: admins.Account,
		}
		gin_response.SuccessResponse(context, http.StatusOK, "", respData, common_msg_id.Success)
	} else {
		gin_response.ErrorResponse(context, http.StatusBadRequest, "帳號已存在", common_msg_id.Fail, nil)
	}
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
func AdminLogin(context *gin.Context) {
	type ReceiveData struct {
		Account  string `json:"account" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var data ReceiveData
	var admins model.Admins
	if err := context.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		check := handler.ValidCheckAndTrans(context, err)
		if check {
			return
		}
		// 非validator.ValidationErrors類型錯誤直接傳回
		handler.HandlerError(context, common_msg_id.Fail, "CreateUser() User 建立失敗(BindJSON fail)", err, http.StatusBadRequest, nil)
		return
	}

	admin, err := admins.GetAdminsByAccount(data.Account)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		handler.HandlerError(context, common_msg_id.Fail, "AdminLogin() model GetAdminsByAccount fail", err, http.StatusBadRequest, nil)
		return
	}

	//密碼驗證
	if err := bcryptEncap.CompareHashAndPassword([]byte(admin.Password), []byte(data.Password)); err != nil {
		handler.HandlerError(context, common_msg_id.Fail, "UserLogin() User CompareHashAndPassword fail", err, http.StatusBadRequest, gin_response.CreateMsgData("fail", "帳密錯誤"))
		return
	}

	jwtToken, err := jwt_secret.GenerateToken(common_const.LoginAdmin, admin.ID)
	if err != nil {
		handler.HandlerError(context, common_msg_id.Fail, "UserLogin() User jwt GenerateToken fail", err, http.StatusBadRequest, nil)
		return
	}

	gin_response.SuccessResponse(context, http.StatusOK, "", gin_response.CreateMsgData("jwt_token", jwtToken), common_msg_id.Success)

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
func GetAdminsById(context *gin.Context) {
	type receiveData struct {
		FilterAdminsId string `form:"filterAdminsId" json:"filterAdminsId" binding:"required"`
	}

	var data receiveData
	admin_id, _ := context.Get("adminId")
	zap.S().Info("admin_id :", admin_id)
	data.FilterAdminsId = context.Param("filterAdminsId")
	zap.S().Info("FilterAdminsId :", data)

	gin_response.SuccessResponse(context, http.StatusOK, "", nil, common_msg_id.Success)
}
