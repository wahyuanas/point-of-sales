package server

import (
	"context"
	"log"

	"github.com/kelseyhightower/envconfig"
	objectvalue "github.com/wahyuanas/point-of-sale/account/api/object-value"
	"github.com/wahyuanas/point-of-sale/account/delivery/grpc/proto/pb"
	"github.com/wahyuanas/point-of-sale/account/repository"
	"github.com/wahyuanas/point-of-sale/account/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

type grpcServer struct {
	pb.UnimplementedAccountServiceServer
	service service.AccountService
}

func NewGRPCServer() *grpc.Server {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var accRepo repository.AccountRepository
	accRepo, err = repository.ImplAccountRepository(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	svc := service.ImplAccountService(accRepo)

	gsrv := grpc.NewServer()
	pb.RegisterAccountServiceServer(gsrv, &grpcServer{
		UnimplementedAccountServiceServer: pb.UnimplementedAccountServiceServer{},
		service:                           svc,
	})
	return gsrv
}

func (s *grpcServer) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	obj := objectvalue.SignUp{UserName: in.UserName, Name: in.Name, Password: in.Password, Email: in.Email, PhoneNumber: in.PhoneNumber}
	_, err := s.service.SignUp(&obj)
	resp := &pb.SignUpResponse{Response: &pb.CommonResponse{
		Status: true, Code: 200, Message: "SUKSES",
	}}

	return resp, err

}
func (s *grpcServer) SignIn(context.Context, *pb.SignInRequest) (*pb.SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (s *grpcServer) SignOut(context.Context, *pb.SignOutRequest) (*pb.SignOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}
func (s *grpcServer) Update(context.Context, *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (s *grpcServer) Delete(context.Context, *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (s *grpcServer) GetAccount(context.Context, *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (s *grpcServer) GetAccounts(context.Context, *pb.EmptyRequest) (*pb.GetAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccounts not implemented")
}