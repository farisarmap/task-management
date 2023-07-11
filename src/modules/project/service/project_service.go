package service

import interfaces "task-management-be/src/modules/project/interface"

type ProjectService struct {
	Repository interfaces.IProjectRepository
}

func NewProjectService(repository interfaces.IProjectRepository) *ProjectService {
	return &ProjectService{
		Repository: repository,
	}
}
