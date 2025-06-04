package service

import (
	"context"

	"github.com/mauricio-pagarme/go-grpc-example/internal/database"
	"github.com/mauricio-pagarme/go-grpc-example/internal/pb"
)

type TransactionTypeService struct {
	pb.UnimplementedTransactionTypeServiceServer
	TransactionTypeDB database.TransactionType
}

func NewTransactionTypeService(TransactionTypeDB database.TransactionType) *TransactionTypeService {
	return &TransactionTypeService{
		TransactionTypeDB: TransactionTypeDB,
	}
}

func (u *TransactionTypeService) CreateTransactionType(ctx context.Context, input *pb.CreateTransactionTypeRequest) (*pb.TransactionTypeResponse, error) {
	TransactionType, err := u.TransactionTypeDB.Create(input.Name, input.Description)
	if err != nil {
		return nil, err
	}

	TransactionTypeResponse := &pb.TransactionType{
		Id:          TransactionType.ID,
		Name:        TransactionType.Name,
		Description: TransactionType.Description,
	}

	return &pb.TransactionTypeResponse{
		TransactionType: TransactionTypeResponse,
	}, nil

}
