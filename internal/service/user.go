package service

import (
	"context"
	"database/sql"
	"errors"
	"io"

	"github.com/mauricio-pagarme/go-grpc-example/internal/database"
	"github.com/mauricio-pagarme/go-grpc-example/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	UserDB database.User
}

func NewUserService(userDB database.User) *UserService {
	return &UserService{
		UserDB: userDB,
	}
}

func (u *UserService) CreateUser(ctx context.Context, input *pb.CreateUserRequest) (*pb.UserResponse, error) {
	user, err := u.UserDB.Create(input.Name, input.Cnpj)
	if err != nil {
		return nil, err
	}

	userResponse := &pb.User{
		Id:   user.ID,
		Name: user.Name,
		Cnpj: user.Cnpj,
	}

	return &pb.UserResponse{
		User: userResponse,
	}, nil

}

func (u *UserService) ListUsers(context.Context, *pb.Blank) (*pb.UserListResponse, error) {
	users, err := u.UserDB.FindAll()
	if err != nil {
		return nil, err
	}

	var userList []*pb.User
	for _, user := range users {
		userList = append(userList, &pb.User{
			Id:   user.ID,
			Name: user.Name,
			Cnpj: user.Cnpj,
		})
	}

	return &pb.UserListResponse{
		Users: userList,
	}, nil
}

func (u *UserService) GetUser(cxt context.Context, input *pb.GetUserRequest) (*pb.UserResponse, error) {
	user, err := u.UserDB.FindById(input.Id)
	if err != nil {
		// Detecta erro específico (ajuste de acordo com sua implementação)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "Usuário com ID %d não encontrado", input.Id)
		}
		return nil, status.Errorf(codes.Internal, "Erro ao buscar usuário: %v", err)
	}

	userResponse := &pb.User{
		Id:   user.ID,
		Name: user.Name,
		Cnpj: user.Cnpj,
	}

	return &pb.UserResponse{
		User: userResponse,
	}, nil
}

func (u *UserService) CreateUserStream(stream grpc.ClientStreamingServer[pb.CreateUserRequest, pb.UserListResponse]) error {
	ur := &pb.UserListResponse{}
	for {
		user, err := stream.Recv() // com esse receive começamos a receber o stream de dados
		if err == io.EOF {
			return stream.SendAndClose(ur)
		}
		if err != nil {
			return err
		}

		userResult, err := u.UserDB.Create(user.Name, user.Cnpj)
		if err != nil {
			return err
		}

		ur.Users = append(ur.Users, &pb.User{
			Id:   userResult.ID,
			Name: userResult.Name,
			Cnpj: userResult.Cnpj,
		})
	}
}

func (u *UserService) CreateUserStreamBidirectional(stream grpc.BidiStreamingServer[pb.CreateUserRequest, pb.UserResponse]) error {
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		userResult, err := u.UserDB.Create(user.Name, user.Cnpj)

		if err != nil {
			return err
		}

		err = stream.Send(&pb.UserResponse{
			User: &pb.User{
				Id:   userResult.ID,
				Name: userResult.Name,
				Cnpj: userResult.Cnpj,
			},
		})

		if err != nil {
			return err
		}
	}
}
