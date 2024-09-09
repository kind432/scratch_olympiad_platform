package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"

	"github.com/robboworld/scratch_olympiad_platform/internal/consts"
	"github.com/robboworld/scratch_olympiad_platform/internal/models"
	"github.com/robboworld/scratch_olympiad_platform/pkg/utils"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateApplication is the resolver for the CreateApplication field.
func (r *mutationResolver) CreateApplication(ctx context.Context, input models.NewApplication) (*models.ApplicationHTTP, error) {
	ginContext, err := utils.GinContextFromContext(ctx)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	application := models.ApplicationCore{
		AuthorID:                      ginContext.Value(consts.KeyId).(uint),
		Nomination:                    input.Nomination,
		AlgorithmicTaskLink:           utils.StringPointerToString(input.AlgorithmicTaskLink),
		AlgorithmicTaskFile:           utils.StringPointerToString(input.AlgorithmicTaskFile),
		CreativeTaskFile:              utils.StringPointerToString(input.CreativeTaskFile),
		CreativeTaskLink:              utils.StringPointerToString(input.CreativeTaskLink),
		EngineeringTaskFile:           utils.StringPointerToString(input.EngineeringTaskFile),
		EngineeringTaskCloudLink:      utils.StringPointerToString(input.EngineeringTaskCloudLink),
		EngineeringTaskVideo:          utils.StringPointerToString(input.EngineeringTaskVideo),
		EngineeringTaskVideoCloudLink: utils.StringPointerToString(input.EngineeringTaskVideoCloudLink),
		Note:                          utils.StringPointerToString(input.Note),
	}
	newApplication, err := r.applicationService.CreateApplication(application)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}

	applicationHttp := models.ApplicationHTTP{}
	applicationHttp.FromCore(newApplication)
	return &applicationHttp, nil
}
