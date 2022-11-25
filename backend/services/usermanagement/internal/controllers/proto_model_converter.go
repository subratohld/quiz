package controllers

import (
	cmnutil "github.com/subratohld/quiz/cmnlib/util"
	"github.com/subratohld/quiz/usermanagement/internal/models"
	umpb "github.com/subratohld/quiz/usermanagement/internal/umproto"
)

func convertAddUserRequest(req *umpb.AddUserRequest) *models.User {
	return &models.User{
		UserID:       cmnutil.GenerateUUID(),
		EmailID:      req.GetUser().GetEmailId(),
		MobileNumber: req.GetUser().GetMobileNumber(),
		FirstName:    req.GetUser().GetFirstName(),
		LastName:     req.GetUser().GetLastName(),
		Address:      req.GetUser().GetAddress(),
	}
}

func convertUser(user *models.User) *umpb.User {
	if user == nil {
		return nil
	}

	return &umpb.User{
		UserId:       user.UserID,
		EmailId:      user.EmailID,
		MobileNumber: user.MobileNumber,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Address:      user.Address,
	}
}

func createAddUserResponse(msg string, statusCode int, user *models.User) *umpb.AddUserResponse {
	return &umpb.AddUserResponse{
		Message:    msg,
		StatusCode: int32(statusCode),
		User:       convertUser(user),
	}
}

func createGetUserByIDResponse(msg string, statusCode int, user *models.User) *umpb.GetUserByIDResponse {
	return &umpb.GetUserByIDResponse{
		Message:    msg,
		StatusCode: int32(statusCode),
		User:       convertUser(user),
	}
}
