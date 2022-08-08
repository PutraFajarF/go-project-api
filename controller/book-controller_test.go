package controller

import (
	"reflect"
	"testing"

	"github.com/PutraFajarF/go-project-api/service"
	"github.com/gin-gonic/gin"
)

func TestNewBookController(t *testing.T) {
	type args struct {
		bookServ service.BookService
		jwtServ  service.JWTService
	}
	tests := []struct {
		name string
		args args
		want BookController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBookController(tt.args.bookServ, tt.args.jwtServ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBookController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookController_GetAllBooks(t *testing.T) {
	type args struct {
		context *gin.Context
	}
	tests := []struct {
		name string
		c    *bookController
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.GetAllBooks(tt.args.context)
		})
	}
}

func Test_bookController_GetBooksById(t *testing.T) {
	type args struct {
		context *gin.Context
	}
	tests := []struct {
		name string
		c    *bookController
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.GetBooksById(tt.args.context)
		})
	}
}

func Test_bookController_CreateBooks(t *testing.T) {
	type args struct {
		context *gin.Context
	}
	tests := []struct {
		name string
		c    *bookController
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.CreateBooks(tt.args.context)
		})
	}
}

func Test_bookController_UpdateBooks(t *testing.T) {
	type args struct {
		context *gin.Context
	}
	tests := []struct {
		name string
		c    *bookController
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.UpdateBooks(tt.args.context)
		})
	}
}

func Test_bookController_DeleteBooks(t *testing.T) {
	type args struct {
		context *gin.Context
	}
	tests := []struct {
		name string
		c    *bookController
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.DeleteBooks(tt.args.context)
		})
	}
}

func Test_bookController_getUserIDByToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name string
		c    *bookController
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.getUserIDByToken(tt.args.token); got != tt.want {
				t.Errorf("bookController.getUserIDByToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
