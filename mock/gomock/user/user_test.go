package user

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sudosapan/mock/mocks"
	"github.com/sudosapan/mock/user"
)

func TestUserWithGoMock(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &user.User{Doer: mockDoer}

	// Expect Do to be called once with 1 and "abc" as parameters, and return nil from the mocked call.
	mockDoer.EXPECT().Do(1, "abc").Return(nil).Times(1)

	testUser.Use()
}
