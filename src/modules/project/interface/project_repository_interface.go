package interfaces

import (
	"context"
	"task-management-be/src/domain"
)

type IProjectRepository interface {
	Create(ctx context.Context, request domain.ProjectRequest) (domain.Project, error)
}
