package service

import (
	"errors"
	"fmt"
	"self_go_gin/gin_application/api/v1/user/request"

	"self_go_gin/common/common_const"
	"self_go_gin/domains/user/entity/model"
	"self_go_gin/domains/user/repository"
	"self_go_gin/gin_application/handler"
	"self_go_gin/util/bcryptEncap"
	"self_go_gin/util/jwt_secret"

	"gorm.io/gorm"
)

type UserService struct {
	repo repository.UserRepositoryInterface
}

func NewUserService() *UserService {
	return &UserService{
		repo: repository.NewUserRepository(),
	}
}

func (s *UserService) CreateUser(req request.CreateUserRequest) (*model.User, error) {
	log_data := map[string]interface{}{
		"req": req,
	}
	_, err := s.repo.GetUsersByAccount(req.Account)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("UserService CreateUser() data: %s \n %w", log_data, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//密碼加密
		bcryptPassword, err := bcryptEncap.GenerateFromPassword(req.Password)
		if err != nil {
			return nil, fmt.Errorf("UserService CreateUser() bcrypt fail data: %s \n %w", log_data, err)
		}

		newUsers := &model.User{
			Account:  req.Account,
			Password: string(bcryptPassword),
		}

		user, err := s.repo.CreateUser(newUsers)
		if err != nil {
			return nil, fmt.Errorf("UserService CreateUser() data: %s \n %w", log_data, err)
		}

		return user, nil
	}

	return nil, fmt.Errorf("UserService CreateUser() resource exist data: %s \n %w", log_data, handler.ErrResourceExist)
}

func (s *UserService) CheckLogin(req request.UserLoginRequest) (*string, error) {
	log_data := map[string]interface{}{
		"req": req,
	}

	user, err := s.repo.GetUsersByAccount(req.Account)
	if err != nil {
		return nil, fmt.Errorf("UserService CheckLogin() data: %s \n %w", log_data, err)
	}

	//密碼驗證
	if err := bcryptEncap.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("UserService CheckLogin() CompareHashAndPassword() data : %+v \n %w", log_data, err) // 密碼錯誤
	}
	jwtToken, err := jwt_secret.GenerateToken(common_const.LoginUser, user.ID)
	if err != nil {
		return nil, fmt.Errorf("UserService CheckLogin() GenerateToken() data : %+v \n %w", log_data, err) // 密碼錯誤
	}

	return &jwtToken, nil

}
