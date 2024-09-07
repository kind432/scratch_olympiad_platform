package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"
	"net/http"
	"strconv"

	"github.com/skinnykaen/rpa_clone/internal/consts"
	"github.com/skinnykaen/rpa_clone/internal/models"
	"github.com/skinnykaen/rpa_clone/pkg/utils"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateProjectPage is the resolver for the CreateProjectPage field.
func (r *mutationResolver) CreateProjectPage(ctx context.Context) (*models.ProjectPageHTTP, error) {
	ginContext, err := utils.GinContextFromContext(ctx)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	newProjectPage, err := r.projectPageService.CreateProjectPage(ginContext.Value(consts.KeyId).(uint))
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	projectPageHttp := models.ProjectPageHTTP{}
	projectPageHttp.FromCore(newProjectPage)
	return &projectPageHttp, nil
}

// UpdateProjectPage is the resolver for the UpdateProjectPage field.
func (r *mutationResolver) UpdateProjectPage(ctx context.Context, input models.UpdateProjectPage) (*models.ProjectPageHTTP, error) {
	ginContext, err := utils.GinContextFromContext(ctx)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	atoi, err := strconv.Atoi(input.ID)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": utils.ResponseError{
					Code:    http.StatusBadRequest,
					Message: consts.ErrAtoi,
				},
			},
		}
	}
	projectPage := models.ProjectPageCore{
		ID:          uint(atoi),
		Title:       input.Title,
		Instruction: input.Instruction,
		Notes:       input.Notes,
		IsShared:    input.IsShared,
	}
	updatedProjectPage, err := r.projectPageService.UpdateProjectPage(projectPage, ginContext.Value(consts.KeyId).(uint))
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	projectPageHttp := models.ProjectPageHTTP{}
	projectPageHttp.FromCore(updatedProjectPage)
	return &projectPageHttp, nil
}

// DeleteProjectPage is the resolver for the DeleteProjectPage field.
func (r *mutationResolver) DeleteProjectPage(ctx context.Context, id string) (*models.Response, error) {
	ginContext, err := utils.GinContextFromContext(ctx)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	atoi, err := strconv.Atoi(id)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": utils.ResponseError{
					Code:    http.StatusBadRequest,
					Message: consts.ErrAtoi,
				},
			},
		}
	}
	if err := r.projectPageService.DeleteProjectPage(uint(atoi), ginContext.Value(consts.KeyId).(uint)); err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	return &models.Response{Ok: true}, nil
}

// SetIsBanned is the resolver for the SetIsBanned field.
func (r *mutationResolver) SetIsBanned(ctx context.Context, projectPageID string, isBanned bool) (*models.Response, error) {
	atoi, err := strconv.Atoi(projectPageID)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": utils.ResponseError{
					Code:    http.StatusBadRequest,
					Message: consts.ErrAtoi,
				},
			},
		}
	}
	if err := r.projectPageService.SetIsBanned(uint(atoi), isBanned); err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	return &models.Response{Ok: true}, nil
}

// GetProjectPageByID is the resolver for the GetProjectPageById field.
func (r *queryResolver) GetProjectPageByID(ctx context.Context, id string) (*models.ProjectPageHTTP, error) {
	ginContext, err := utils.GinContextFromContext(ctx)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	atoi, err := strconv.Atoi(id)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": utils.ResponseError{
					Code:    http.StatusBadRequest,
					Message: consts.ErrAtoi,
				},
			},
		}
	}
	project, err := r.projectPageService.GetProjectPageById(uint(atoi), ginContext.Value(consts.KeyId).(uint), ginContext.Value(consts.KeyRole).(models.Role))
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	projectPageHttp := models.ProjectPageHTTP{}
	projectPageHttp.FromCore(project)
	return &projectPageHttp, nil
}

// GetAllProjectPagesByAuthorID is the resolver for the GetAllProjectPagesByAuthorId field.
func (r *queryResolver) GetAllProjectPagesByAuthorID(ctx context.Context, id string, page *int, pageSize *int) (*models.ProjectPageHTTPList, error) {
	atoi, err := strconv.Atoi(id)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": utils.ResponseError{
					Code:    http.StatusBadRequest,
					Message: consts.ErrAtoi,
				},
			},
		}
	}
	projects, countRows, err := r.projectPageService.GetProjectsPageByAuthorId(uint(atoi), page, pageSize)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	return &models.ProjectPageHTTPList{
		ProjectPages: models.FromProjectPagesCore(projects),
		CountRows:    int(countRows),
	}, nil
}

// GetAllProjectPagesByAccessToken is the resolver for the GetAllProjectPagesByAccessToken field.
func (r *queryResolver) GetAllProjectPagesByAccessToken(ctx context.Context, page *int, pageSize *int) (*models.ProjectPageHTTPList, error) {
	ginContext, err := utils.GinContextFromContext(ctx)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	projects, countRows, err := r.projectPageService.GetAllProjectPages(page, pageSize, ginContext.Value(consts.KeyId).(uint), ginContext.Value(consts.KeyRole).(models.Role))
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	return &models.ProjectPageHTTPList{
		ProjectPages: models.FromProjectPagesCore(projects),
		CountRows:    int(countRows),
	}, nil
}
