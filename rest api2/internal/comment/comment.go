package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch the dta")
	ErrNotImplement    = errors.New("not yet")
)

// structure for our sservice
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// define all the method to operate the service
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	UpdateComment(context.Context, string) (Comment, error)
	DeleteComment(context.Context, string) error
}

type Service struct {
	Store Store
}

// return pointer to new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}

}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println(" retvie the comment")

	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplement
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplement
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplement
}
