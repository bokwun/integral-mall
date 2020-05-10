package logic

import "integral-mall/user/model"

type (
	UserLogic struct {
		userModel *model.UserModel
	}
	RegisterRequest struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
	}
	RegisterResponse struct {
	}
)

func NewUserLogic(userModel *model.UserModel) *UserLogic {
	return &UserLogic{
		userModel: userModel,
	}
}

//register
func (c *UserLogic) Register(r *RegisterRequest) (*RegisterResponse, error) {
	response := new(RegisterResponse)

	return response, nil
}
