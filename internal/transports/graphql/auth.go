package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"
	"time"

	"github.com/robboworld/scratch_olympiad_platform/internal/consts"
	"github.com/robboworld/scratch_olympiad_platform/internal/models"
	"github.com/robboworld/scratch_olympiad_platform/pkg/utils"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// SignUp is the resolver for the SignUp field.
func (r *mutationResolver) SignUp(ctx context.Context, input models.SignUp) (*models.Response, error) {
	// middlename не обязательное поле и может быть nil
	newUser := models.UserCore{
		Email:          input.Email,
		Password:       input.Password,
		Firstname:      input.Firstname,
		Lastname:       input.Lastname,
		Middlename:     utils.StringPointerToString(input.Middlename),
		Nickname:       input.Nickname,
		Role:           models.RoleStudent,
		IsActive:       false,
		ActivationLink: utils.GetHashString(time.Now().String()),
	}
	err := r.authService.SignUp(newUser)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return &models.Response{Ok: false}, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	return &models.Response{Ok: true}, nil
}

// SignIn is the resolver for the SignIn field.
func (r *mutationResolver) SignIn(ctx context.Context, input models.SignIn) (*models.SignInResponse, error) {
	tokens, err := r.authService.SignIn(input.Email, input.Password)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return &models.SignInResponse{}, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	return &models.SignInResponse{
		AccessToken:  tokens.Access,
		RefreshToken: tokens.Refresh,
	}, nil
}

// RefreshToken is the resolver for the RefreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, refreshToken string) (*models.SignInResponse, error) {
	accessToken, err := r.authService.Refresh(refreshToken)
	if err != nil {
		r.loggers.Err.Printf("%s", err)
		return &models.SignInResponse{}, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	return &models.SignInResponse{
		AccessToken: accessToken,
	}, nil
}

// ConfirmActivation is the resolver for the ConfirmActivation field.
func (r *mutationResolver) ConfirmActivation(ctx context.Context, activationLink string) (*models.SignInResponse, error) {
	tokens, err := r.authService.ConfirmActivation(activationLink)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return &models.SignInResponse{}, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	return &models.SignInResponse{
		AccessToken:  tokens.Access,
		RefreshToken: tokens.Refresh,
	}, nil
}

// Me is the resolver for the Me field.
func (r *queryResolver) Me(ctx context.Context) (*models.UserHTTP, error) {
	ginContext, err := utils.GinContextFromContext(ctx)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	user, err := r.userService.GetUserById(ginContext.Value(consts.KeyId).(uint), ginContext.Value(consts.KeyRole).(models.Role))
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	userHttp := models.UserHTTP{}
	userHttp.FromCore(user)
	return &userHttp, nil
}
