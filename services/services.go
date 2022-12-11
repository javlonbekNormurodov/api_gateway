package services

import (
	"github.com/javlonbekNormurodov/api_gateway/config"
	"github.com/javlonbekNormurodov/api_gateway/genproto/book_service"
	"github.com/javlonbekNormurodov/api_gateway/genproto/user_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	BookService() book_service.BookServiceClient
	UserService() user_service.UserServiceClient
}

type grpcClients struct {
	bookService book_service.BookServiceClient
	userService user_service.UserServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {

	connBookService, err := grpc.Dial(
		cfg.BookServiceHost+cfg.BookGRPCPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	connUserService, err := grpc.Dial(
		cfg.UserServiceHost+cfg.UserGRPCPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(10000648), grpc.MaxCallSendMsgSize(10000648)),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		bookService: book_service.NewBookServiceClient(connBookService),
		userService: user_service.NewUserServiceClient(connUserService),
	}, nil
}

func (g *grpcClients) BookService() book_service.BookServiceClient {
	return g.bookService
}

func (g *grpcClients) UserService() user_service.UserServiceClient {
	return g.userService
}
