package v1

import "appGo/pkg/api/apiv1"

type serverInterface interface {
	apiv1.ServerInterface
}

var _ serverInterface = (*Handler)(nil)

func New(usersRepository usersClient, linksRepository linksClient) *Handler {
	return &Handler{UsersHandler: NewUsersHandler(usersRepository), LinksHandler: NewLinksHandler(linksRepository)}
}

type Handler struct {
	*UsersHandler
	*LinksHandler
}
