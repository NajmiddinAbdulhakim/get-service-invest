package repo

import (
	pb "github.com/NajmiddinAbdulhakim/iman/get-service/genproto"
)

type GetAPIStorage interface {
	GetInfoFromAPI() ([]*pb.Post, error)
}

type SQLStorage interface {
	CreatePost(posts []*pb.Post) error
}
