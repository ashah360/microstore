package account

import "context"
import "github.com/segmentio/ksuid"

type Service interface {
	PostAccount(ctx context.Context, name string) (*Account, error)
	GetAccount(ctx context.Context, id string) (*Account, error)
	GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
}

type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type accountService struct {
	respository Respository
}

func NewService(r Repository) Service {
	return &accountService{r}
}

func (s *accountService) PostAccount(ctx context.Context, name string) (*Account, error) {
	a := &Account{
		Name: name,
		ID:   ksuid.New().String(),
	}
	if err := s.respository.PutAccount(ctx, *a); err != nil {
		return nil, err
	}
	return a, nil
}
