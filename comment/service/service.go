package service

import (
	"comment/pb"
	"context"
)

type CommentService struct{}

func NewCommentService() *CommentService {
	return &CommentService{}
}

func (service *CommentService) ShowComment(ctx context.Context, req *pb.ShowCommentRequest) (*pb.ShowCommentResponse, error) {
	return &pb.ShowCommentResponse{
		CommentId: 15,
	}, nil
}

func (service *CommentService) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	return &pb.CreateCommentResponse{
		CommentId: 19,
	}, nil
}
