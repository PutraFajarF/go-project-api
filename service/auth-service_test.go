package service

import (
	"reflect"
	"testing"

	"github.com/PutraFajarF/go-project-api/dto"
	"github.com/PutraFajarF/go-project-api/entity"
	"github.com/PutraFajarF/go-project-api/repository"
)

func TestNewAuthService(t *testing.T) {
	type args struct {
		userRep repository.UserRepository
	}
	tests := []struct {
		name string
		args args
		want AuthService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthService(tt.args.userRep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authService_VerifyCredential(t *testing.T) {
	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name    string
		service *authService
		args    args
		want    interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.service.VerifyCredential(tt.args.email, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authService.VerifyCredential() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authService_CreateUser(t *testing.T) {
	type args struct {
		user dto.RegisterDTO
	}
	tests := []struct {
		name    string
		service *authService
		args    args
		want    entity.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.service.CreateUser(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authService.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authService_FindByEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		service *authService
		args    args
		want    entity.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.service.FindByEmail(tt.args.email); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authService.FindByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authService_IsDuplicateEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		service *authService
		args    args
		want    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.service.IsDuplicateEmail(tt.args.email); got != tt.want {
				t.Errorf("authService.IsDuplicateEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_comparePassword(t *testing.T) {
	type args struct {
		hashedPwd     string
		plainPassword []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := comparePassword(tt.args.hashedPwd, tt.args.plainPassword); got != tt.want {
				t.Errorf("comparePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
