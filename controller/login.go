package controller

import (
	"context"
	"strings"

	"github.com/MonkeyBuisness/family-api/model"
)

// Login user by email and password.
func (s Service) Login(ctx context.Context, req model.LoginRequest) (res model.LoginResponse) {
	if strings.EqualFold(req.Email, "test@mail.ru") && strings.EqualFold(req.Password, "12345") {
		res.Result = true
		res.Token = "Dhx0L5jSHaandVLAeKVBanyoOlbIQU"
		return
	}
	res.Error = "invalid login or password"
	return
}
