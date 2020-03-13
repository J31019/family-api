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

// Register user.
func (s Service) Register(ctx context.Context, req model.RegisterRequest) (res model.RegisterResponse) {
	if strings.EqualFold(req.Email, "test@mail.ru") {
		res.Error = "user with provided email already exists"
		return
	}
	res.Result = true
	return
}

// Submit user registration.
func (s Service) Submit(ctx context.Context, req model.SubmitRequest) (res model.SubmitResponse) {
	if !strings.EqualFold(req.Code, "qwerty") {
		res.Error = "invalid code"
		return
	}
	res.Result = true
	res.Token = "Dhx0L5jSHaandVLAeKVBanyoOlbIQU"
	return
}
