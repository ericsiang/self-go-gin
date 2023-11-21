package v1

import (
	"api/initialize"
	"api/model"
	"api/util/gin_response"
	"api/util/bcryptEncap"
	"api/util/jwt_secret"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"gorm.io/gorm"
)

func validCheckAndTrans(context *gin.Context, err error) {
	validTrans, ok := initialize.ErrorValidateCheckAndTrans(err)
	if ok {
		gin_response.ErrorResponse(context, http.StatusBadRequest, validTrans)
		return
	}
}

func CreateUser(context *gin.Context) {
	type receiveData struct {
		Account  string `json:"account" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var data receiveData
	var newUsers model.Users

	if err := context.ShouldBindJSON(&data); err != nil {
		validCheckAndTrans(context, err)
		// 非validator.ValidationErrors類型錯誤直接傳回
		zap.S().Error("CreateUser() User 建立失敗(BindJSON fail) :" + err.Error())
		gin_response.ErrorResponse(context, http.StatusBadRequest, gin_response.CreateMsg("fail", "BindJSON fail"))
		return
	}

	_, err := newUsers.GetUsersByAccount(data.Account)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		gin_response.ErrorResponse(context, http.StatusBadRequest, gin_response.CreateMsg("fail", "db GetUsersByAccount fail"))
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

		gin_response.SuccessResponse(context, http.StatusOK, gin_response.CreateMsg("success", "success"), users)
	} else {
		gin_response.ErrorResponse(context, http.StatusBadRequest, gin_response.CreateMsg("fail", "帳號已存在"))
	}

}

func UserLogin(context *gin.Context) {
	type receiveData struct {
		Account  string `json:"account" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var data receiveData
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
	}

	//密碼驗證
	if err := bcryptEncap.CompareHashAndPassword([]byte(user.Password),[]byte(data.Password)); err != nil {
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

func GetUsers(context *gin.Context) {
	type receiveData struct {
	}
	log.Println("GetUsers : ")
	// var data receiveData
	// var users model.Users

}

func GetUsersById(context *gin.Context) {
	type receiveData struct {
		FilterUsersId string `form:"filterUsersId"`
	}

	var data receiveData
	usersId ,_ :=context.Get("usersId")
	zap.S().Info("usersId :",usersId);
	data.FilterUsersId = context.Param("filterUsersId")
	zap.S().Info("filterUsersId :",data);
}
