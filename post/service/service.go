// package service

// import (
// 	"context"

// 	"ponica/post/gen/api"
// )

// type PostService struct{}

// func NewPostService() PostService {
// 	return PostService{}
// }

// func (service *PostService) ShowPost(ctx context.Context, req *api.ShowPostRequest) (*api.ShowPostResponse, error) {
// 	return &api.ShowPostResponse{
// 		Post: &api.Post{
// 			Id:      3,
// 			Title:   "title",
// 			Message: "message",
// 		},
// 	}, nil
// }

// func (service *PostService) CreatePost(ctx context.Context, req *api.CreatePostRequest) (*api.CreatePostResponse, error) {
// 	return &api.CreatePostResponse{
// 		Id: 4,
// 	}, nil
// }
