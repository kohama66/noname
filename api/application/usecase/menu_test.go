package usecase

import (
	"context"
	"testing"

	"github.com/myapp/noname/api/application/usecase/requestmodel"
	"github.com/myapp/noname/api/infrastructure/mock"
	"github.com/myapp/noname/api/infrastructure/response"
)

func InitMenu() Menu {
	menu := mock.NewMenuRepo()
	responseMenu := response.NewMenu()
	user := mock.NewUserRepo()
	usecaseMenu := NewMenu(menu, responseMenu, user)
	return usecaseMenu
}

func TestSuccessDeleteToBeautician(t *testing.T) {
	m := InitMenu()
	ctx := context.Background()
	tests := []*requestmodel.BeauticianMenuDelete{
		{
			AuthID: "1",
			RandID: "rand",
		},
	}
	for _, tt := range tests {
		result := m.DeleteToBeautician(ctx, tt)
		if result != nil {
			t.Errorf("TestDeleteToBeautician is Error %v", result)
		}
	}
}

func TestFaileDeleteToBeautician(t *testing.T) {
	m := InitMenu()
	ctx := context.Background()
	tests := []*requestmodel.BeauticianMenuDelete{
		{
			AuthID: "2",
			RandID: "rand",
		},
	}
	for _, tt := range tests {
		result := m.DeleteToBeautician(ctx, tt)
		if result == nil {
			t.Errorf("TestDeleteToBeautician is Error %v", result)
		}
	}
}
