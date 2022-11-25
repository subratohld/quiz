package controllers

import (
	"context"
	"net/http"

	cmnerr "github.com/subratohld/quiz/cmnlib/errors"
	"github.com/subratohld/quiz/usermanagement/internal/common/consts"
	"github.com/subratohld/quiz/usermanagement/internal/services"
	umpb "github.com/subratohld/quiz/usermanagement/internal/umproto"
)

type (
	Usermanagement interface {
		AddUser(context.Context, *umpb.AddUserRequest) (*umpb.AddUserResponse, error)
		GetUserByID(context.Context, *umpb.GetUserByIDRequest) (*umpb.GetUserByIDResponse, error)
	}

	userManagement struct {
		svcManager *services.ServiceManager
	}
)

var _ Usermanagement = (*userManagement)(nil)

func newUsermanagement(svcManager *services.ServiceManager) *userManagement {
	return &userManagement{
		svcManager: svcManager,
	}
}

func (u *userManagement) AddUser(ctx context.Context, req *umpb.AddUserRequest) (*umpb.AddUserResponse, error) {
	user := convertAddUserRequest(req)

	resp, err := u.svcManager.UMService.AddUser(ctx, user)
	if err != nil {
		return createAddUserResponse(consts.ErrAddUser, http.StatusInternalServerError, nil), nil
	}

	return createAddUserResponse(consts.UserAdded, http.StatusOK, resp), nil
}

func (u *userManagement) GetUserByID(ctx context.Context, req *umpb.GetUserByIDRequest) (*umpb.GetUserByIDResponse, error) {
	if len(req.GetAuthData().GetUserId()) == 0 {
		return createGetUserByIDResponse(consts.MissingUserID, http.StatusBadRequest, nil), nil
	}

	user, err := u.svcManager.UMService.GetUserByID(ctx, req.GetAuthData().GetUserId())
	if err != nil {
		msg, statusCode := consts.SomethingWrong, http.StatusInternalServerError

		switch e := err.(type) {
		case *cmnerr.NotFoundError:
			statusCode = e.StatusCode()
			msg = consts.ErrFindUser
		}

		return createGetUserByIDResponse(msg, statusCode, nil), nil
	}

	return createGetUserByIDResponse(consts.UserRetrieveSuccess, http.StatusOK, user), nil
}
