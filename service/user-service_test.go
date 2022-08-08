package service

import (
	"reflect"
	"testing"

	"github.com/PutraFajarF/go-project-api/dto"
	"github.com/PutraFajarF/go-project-api/entity"
	"github.com/PutraFajarF/go-project-api/repository"
)

func TestNewUserService(t *testing.T) {
	type args struct {
		userRepo repository.UserRepository
	}
	tests := []struct {
		name string
		args args
		want UserService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.userRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_Update(t *testing.T) {
	type args struct {
		user dto.UserUpdateDTO
	}
	tests := []struct {
		name    string
		service *userService
		args    args
		want    entity.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.service.Update(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_Profile(t *testing.T) {
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		service *userService
		args    args
		want    entity.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.service.Profile(tt.args.userID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.Profile() = %v, want %v", got, tt.want)
			}
		})
	}
}
