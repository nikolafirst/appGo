package v1

import "appGo/pkg/pb"

type usersClient interface {
	pb.UserServiceClient
}

type linksClient interface {
	pb.LinkServiceClient
}
