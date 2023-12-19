package v1

import (
	"api/initialize"
	"api/model"
	"api/util/bcryptEncap"
	"api/util/gin_response"
	"api/util/jwt_secret"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func validCheckAndTrans(context *gin.Context, err error) (ok bool) {
	validTrans, ok := initialize.ErrorValidateCheckAndTrans(err)
	if ok {
		gin_response.ErrorResponse(context, http.StatusBadRequest, validTrans)
		return ok
	}
	return false
}

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
func CreateUser(context *gin.Context) {
	type receiveData struct {
		Account  string `json:"account" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	type responseData struct {
		UsersId uint   `json:"id"`
		Account string `json:"account"`
	}

	var data receiveData
	var newUsers model.Users
	var respData responseData

	if err := context.ShouldBindJSON(&data); err != nil {
		// check := validCheckAndTrans(context, err)
		// if check {
		// 	return
		// }
		// 非validator.ValidationErrors類型錯誤直接傳回
		zap.S().Error("CreateUser() User 建立失敗(BindJSON fail) :" + err.Error())
		gin_response.ErrorResponse(context, http.StatusBadRequest, gin_response.CreateMsg("fail", "BindJSON fail"))
		return
	}

	_, err := newUsers.GetUsersByAccount(data.Account)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		gin_response.ErrorResponse(context, http.StatusBadRequest, gin_response.CreateMsg("fail", "db GetUsersByAccount fail"))
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		//密碼加密
		bcryptPassword, err := bcryptEncap.GenerateFromPassword(data.Password)
		if err != nil {
			zap.S().Error("CreateUser() User 建立失敗(bcrypt fail) :" + err.Error())
			gin_response.ErrorResponse(context, http.StatusBadRequest, gin_response.CreateMsg("fail", "bcrypt GenerateFromPassword fail"))
			return
		}
		newUsers.Account = data.Account
		newUsers.Password = string(bcryptPassword)
		users, err := newUsers.CreateUser()
		if err != nil {
			zap.S().Error("CreateUser() User 建立失敗(db fail) :" + err.Error())
			gin_response.ErrorResponse(context, http.StatusBadRequest, gin_response.CreateMsg("fail", "db CreateUser fail"))
			return
		}
		respData = responseData{
			UsersId: users.ID,
			Account: users.Account,
		}
		gin_response.SuccessResponse(context, http.StatusOK, gin_response.CreateMsg("success", "success"), respData)
	} else {
		gin_response.ErrorResponse(context, http.StatusBadRequest, gin_response.CreateMsg("fail", "帳號已存在"))
	}

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
func UserLogin(context *gin.Context) {
	type ReceiveData struct {
		Account  string `json:"account" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var data ReceiveData
	var users model.Users
	if err := context.ShouldBindJSON(&data); err != nil {
		validCheckAndTrans(context, err)
		// 非validator.ValidationErrors類型錯誤直接傳回
		zap.S().Error("CreateUser() User 建立失敗(BindJSON fail) :" + err.Error())
		gin_response.ErrorResponse(context, http.StatusBadRequest, gin_response.CreateMsg("fail", "BindJSON fail"))
		return
	}

	user, err := users.GetUsersByAccount(data.Account)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		gin_response.ErrorResponse(context, http.StatusBadRequest, gin_response.CreateMsg("fail", "db GetUsersByAccount fail"))
		return
	}

	//密碼驗證
	if err := bcryptEncap.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		zap.S().Error("UserLogin() User bcrypt:", err.Error())
		gin_response.ErrorResponse(context, http.StatusBadRequest, gin_response.CreateMsg("fail", "帳密錯誤"))
		return
	}

	jwtToken, err := jwt_secret.GenerateToken(user.ID)
	if err != nil {
		gin_response.ErrorResponse(context, http.StatusBadRequest, gin_response.CreateMsg("fail", "jwt GenerateToken fail"))
		return
	}

	gin_response.SuccessResponse(context, http.StatusOK, gin_response.CreateMsg("jwt_token", jwtToken), nil)

}

// @Summary Get Users By ID
// @Description Get Users By ID
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param filterUsersId path string true "filterUsersId"
// @Success 200 {string}  "成功"
// @Failure 400 {string}  "失敗"
// @Failure 401 {string}  "Unauthorized"
// @Router /api/v1/auth/users/{filterUsersId} [get]
func GetUsersById(context *gin.Context) {
	type receiveData struct {
		FilterUsersId string `form:"filterUsersId" json:"filterUsersId" binding:"required"`
	}

	var data receiveData
	usersId, _ := context.Get("usersId")
	zap.S().Info("usersId :", usersId)
	data.FilterUsersId = context.Param("filterUsersId")
	zap.S().Info("filterUsersId :", data)
	gin_response.SuccessResponse(context, http.StatusOK, gin_response.CreateMsg("GetAdminsById test", "just test"), data)
}
