package usecase

import (
	"context"
	"testing"

	"github.com/mmmommm/go-sample/layered-architecture/domain"
	"github.com/mmmommm/go-sample/layered-architecture/repository"
)

func TestUserGetCase(t *testing.T) {
	testCase := []struct {
		mockUserRepo repository.UserRepository
		usecaseInput *UserGetCaseInput
		expect       *domain.User
		expectErr    error
	}{
		{
			// TODO: mockのやり方を少し考えないといけないかも
			mockUserRepo: repository.ProvideUserRepository(),
			usecaseInput: &UserGetCaseInput{},
			expect:       &domain.User{},
			expectErr:    nil,
		},
	}

	for _, v := range testCase {
		uc := ProvideUserGetCase(v.mockUserRepo)
		result, err := uc(context.Background(), v.usecaseInput)
		if result != v.expect {
			t.Errorf("")
		}
		if err != v.expectErr {
			t.Errorf("")
		}
	}
}