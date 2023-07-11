package interfaces

import (
	"context"
	"task-management-be/src/domain"
)

type IProjectService interface {
	Create(ctx context.Context, request domain.ProjectRequest) (domain.ProjectResponse, error)
}
