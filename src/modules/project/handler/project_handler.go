package handler

import interfaces "task-management-be/src/modules/project/interface"

type ProjectHandler struct {
	Services interfaces.IProjectHandler
}

func (ph *ProjectHandler) NewProjectHandler(services interfaces.IProjectHandler) *ProjectHandler {
	return &ProjectHandler{
		Services: services,
	}
}
