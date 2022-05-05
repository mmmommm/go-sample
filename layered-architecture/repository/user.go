package repository

import "github.com/mmmommm/go-sample/layered-architecture/domain"

type UserRepository interface {
	FindByID(ID string) (*domain.User, error)
	Store(user *domain.User) error
}

type userRepository struct {
	// sql.dbとか、userの取得に必要な情報
}

func (r *userRepository) FindByID(ID string) (*domain.User, error) {
	// 実際に接続する処理
	return nil, nil
}

func (r *userRepository) Store(user *domain.User) error {
	// 実際に接続する処理
	return nil
}

func ProvideUserRepository() UserRepository {
	return &userRepository{}
}