package service

import (
	"context"
	"log"

	pb "github.com/NajmiddinAbdulhakim/iman/get-service/genproto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NajmiddinAbdulhakim/iman/get-service/storage"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	repo storage.IStorage
}

func NewService(db *sqlx.DB, link string) *Service {
	return &Service{
		repo: storage.NewStoragePg(db, link),
	}
}

func (s *Service) GetInfoFromAPI(ctx context.Context, req *pb.Empty) (*pb.GetNetRes, error) {
	res, err := s.repo.GetFromAPI().GetInfoFromAPI()
	if err != nil {
		log.Println(`failed while getting info from net: `, err)
		return nil, status.Error(codes.Internal, `failed while getting info from net`)
	}

	return &pb.GetNetRes{Posts: res}, nil

}

func (s *Service) CreatePost(ctx context.Context, req *pb.CreatePostsReq) (*pb.CreatePostsRes, error) {
	err := s.repo.SQLMethods().CreatePost(req.Posts)
	if err != nil {
		log.Println(`failed while creating posts:`, err)
		return &pb.CreatePostsRes{Success: `failed`}, status.Error(codes.Internal,`failed while creating posts`)
	}
	return &pb.CreatePostsRes{Success: `successful`} , nil
}

func (s *Service) GetNetThenInsertDB(ctx context.Context, req *pb.Empty) (*pb.CreatePostsRes, error) {
	res, err := s.repo.GetFromAPI().GetInfoFromAPI()
	if err != nil {
		log.Println(`failed while getting from net: `, err)
		return &pb.CreatePostsRes{Success: `failed`}, status.Error(codes.Internal,`failed while getting from net`)

	}

	err = s.repo.SQLMethods().CreatePost(res)
	if err != nil {
		log.Println(`failed creating posts:`, err)
		return &pb.CreatePostsRes{Success: `failed`}, status.Error(codes.Internal,`failed creating posts`)
	}
	return &pb.CreatePostsRes{Success: `successful`} , nil
}  