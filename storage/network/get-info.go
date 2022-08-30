package network

import (
	"encoding/json"
	"net/http"

	pb "github.com/NajmiddinAbdulhakim/iman/get-service/genproto"
)

type getRepo struct {
	link string
}

func New(link string) *getRepo {
	return &getRepo{link: link}
}

func (r *getRepo) GetInfoFromAPI() ([]*pb.Post, error) {
	var post []*pb.Post
	for {
		info, err := getInfo(r.link)
		if err != nil {
			return nil, err
		}
		r.link = info.Meta.Pagination.Links.Next

		post = append(post, info.Data...)
		if len(info.Data) == 0 {
			break
		}
	}
	return post, nil
}

func getInfo(url string) (pb.AllInfo, error) {
	var info pb.AllInfo
	response, err := http.Get(url)
	if err != nil {
		return pb.AllInfo{}, err
	}
	err = json.NewDecoder(response.Body).Decode(&info)
	if err != nil {
		return pb.AllInfo{}, err
	}
	return info, err
}
