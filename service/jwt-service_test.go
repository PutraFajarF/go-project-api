package service

import (
	"reflect"
	"testing"

	"github.com/golang-jwt/jwt/v4"
)

func TestNewJWTService(t *testing.T) {
	tests := []struct {
		name string
		want JWTService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJWTService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJWTService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSecretKey(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSecretKey(); got != tt.want {
				t.Errorf("getSecretKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jwtService_GenerateToken(t *testing.T) {
	type args struct {
		UserID string
	}
	tests := []struct {
		name string
		j    *jwtService
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.GenerateToken(tt.args.UserID); got != tt.want {
				t.Errorf("jwtService.GenerateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jwtService_ValidateToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		j       *jwtService
		args    args
		want    *jwt.Token
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.ValidateToken(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("jwtService.ValidateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jwtService.ValidateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
