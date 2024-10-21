package handler

import (
	"net/http"

	"github.com/mechatron-x/atehere/internal/httpserver/handler/request"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/response"
	"github.com/mechatron-x/atehere/internal/usermanagement/service"
)

type (
	UserSignUp struct {
		userService service.User
	}
)

func NewUserSignUp(userService service.User) UserSignUp {
	return UserSignUp{
		userService: userService,
	}
}

func (u UserSignUp) Pattern() string {
	return "POST /api/v1/user/auth"
}

func (u UserSignUp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reqBody := &request.SignUpUser{}
	err := request.Decode(r, w, reqBody)
	if err != nil {
		return
	}

	user, err := u.userService.SignUp(reqBody.Email, reqBody.Password, reqBody.FullName, reqBody.BirthDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userResponse := response.SignUpUser{
		ID:        user.ID().String(),
		Email:     reqBody.Email,
		FullName:  user.FullName().String(),
		BirthDate: user.BirthDate().String(),
	}

	response.Encode(w, userResponse)
}
