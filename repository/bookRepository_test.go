package repository

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/PutraFajarF/go-project-api/dto"
	"github.com/PutraFajarF/go-project-api/entity"
	"github.com/jinzhu/gorm"
)

func MockDb() (*gorm.DB, sqlmock.Sqlmock) {
	_, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal("can't connect to database", err)
	}

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", "localhost", "postgres", "cap_books_api", "ktl123", "5432")
	conn, _ := gorm.Open("postgres", dbUri)

	return conn, mock
}

func TestNewBookRepository(t *testing.T) {
	type args struct {
		dbConn *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want BookRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBookRepository(tt.args.dbConn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBookRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookConnection_InsertBook(t *testing.T) {
	type args struct {
		b entity.Book
	}
	tests := []struct {
		name string
		db   *bookConnection
		args args
		want entity.Book
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.InsertBook(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bookConnection.InsertBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookConnection_UpdateBook(t *testing.T) {
	type args struct {
		b entity.Book
	}
	tests := []struct {
		name string
		db   *bookConnection
		args args
		want entity.Book
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.UpdateBook(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bookConnection.UpdateBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookConnection_DeleteBook(t *testing.T) {
	type args struct {
		b entity.Book
	}
	tests := []struct {
		name string
		db   *bookConnection
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.db.DeleteBook(tt.args.b)
		})
	}
}

func Test_bookConnection_FindBookByID(t *testing.T) {
	type args struct {
		bookID uint64
	}
	tests := []struct {
		name string
		db   *bookConnection
		args args
		want entity.Book
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.FindBookByID(tt.args.bookID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bookConnection.FindBookByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookConnection_AllBook(t *testing.T) {
	type args struct {
		p dto.PaginationDTO
	}
	tests := []struct {
		name string
		db   *bookConnection
		args args
		want dto.PaginationDTO
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.AllBook(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bookConnection.AllBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookConnection_Allbook_ShouldReturnError(t *testing.T) {
	type args struct {
		dto.PaginationDTO
	}
	tests := []struct {
		name    string
		args    args
		want    []entity.Book
		wantErr error
	}{
		// TODO: Add test cases.
		{
			"Request Get Data Books Failed",
			args{},
			nil,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := MockDb()
			repo := NewBookRepository(db)

			mock.ExpectQuery(`select \* from books`).WillReturnError(errors.New(""))
			got1 := repo.AllBook(dto.PaginationDTO{})
			if !reflect.DeepEqual(got1, tt.wantErr) {
				t.Errorf("FoodRepositoryDB.FindAll() got1 = %v, want %v", got1, tt.wantErr)
			}
		})
	}
}
