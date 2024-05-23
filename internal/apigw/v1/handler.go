package v1

import "appGo/pkg/api/apiv1"

type serverInterface interface {
	apiv1.ServerInterface
}

var _ serverInterface = (*Handler)(nil)

func New(usersRepository usersClient, linksRepository linksClient) *Handler {
	return &Handler{usersHandler: newUsersHandler(usersRepository), linksHandler: newLinksHandler(linksRepository)}
}

type Handler struct {
	*usersHandler
	*linksHandler
}
