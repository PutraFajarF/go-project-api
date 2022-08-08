package service

import (
	"reflect"
	"testing"

	"github.com/PutraFajarF/go-project-api/dto"
	"github.com/PutraFajarF/go-project-api/entity"
	"github.com/PutraFajarF/go-project-api/repository"
)

func TestNewBookService(t *testing.T) {
	type args struct {
		bookRepo repository.BookRepository
	}
	tests := []struct {
		name string
		args args
		want BookService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBookService(tt.args.bookRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBookService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookService_Insert(t *testing.T) {
	type args struct {
		b dto.BookCreateDTO
	}
	tests := []struct {
		name    string
		service *bookService
		args    args
		want    entity.Book
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.service.Insert(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bookService.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookService_Update(t *testing.T) {
	type args struct {
		b dto.BookUpdateDTO
	}
	tests := []struct {
		name    string
		service *bookService
		args    args
		want    entity.Book
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.service.Update(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bookService.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookService_Delete(t *testing.T) {
	type args struct {
		b entity.Book
	}
	tests := []struct {
		name    string
		service *bookService
		args    args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.service.Delete(tt.args.b)
		})
	}
}

func Test_bookService_All(t *testing.T) {
	type args struct {
		p dto.PaginationDTO
	}
	tests := []struct {
		name    string
		service *bookService
		args    args
		want    dto.PaginationDTO
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.service.All(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bookService.All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookService_FindByID(t *testing.T) {
	type args struct {
		bookID uint64
	}
	tests := []struct {
		name    string
		service *bookService
		args    args
		want    entity.Book
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.service.FindByID(tt.args.bookID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bookService.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookService_IsAllowedToEdit(t *testing.T) {
	type args struct {
		userID string
		bookID uint64
	}
	tests := []struct {
		name    string
		service *bookService
		args    args
		want    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.service.IsAllowedToEdit(tt.args.userID, tt.args.bookID); got != tt.want {
				t.Errorf("bookService.IsAllowedToEdit() = %v, want %v", got, tt.want)
			}
		})
	}
}
