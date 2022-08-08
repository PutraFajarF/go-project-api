package repository

import (
	"math"

	"github.com/PutraFajarF/go-project-api/dto"
	"github.com/PutraFajarF/go-project-api/entity"
	"github.com/jinzhu/gorm"
)

type BookRepository interface {
	InsertBook(b entity.Book) entity.Book
	UpdateBook(b entity.Book) entity.Book
	DeleteBook(b entity.Book)
	AllBook(dto.PaginationDTO) dto.PaginationDTO
	FindBookByID(bookID uint64) entity.Book
}

type bookConnection struct {
	connection *gorm.DB
}

// NewBookRepository creates an instance BookRepository
func NewBookRepository(dbConn *gorm.DB) BookRepository {
	return &bookConnection{
		connection: dbConn,
	}
}

func (db *bookConnection) InsertBook(b entity.Book) entity.Book {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *bookConnection) UpdateBook(b entity.Book) entity.Book {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *bookConnection) DeleteBook(b entity.Book) {
	db.connection.Delete(&b)
}

func (db *bookConnection) FindBookByID(bookID uint64) entity.Book {
	var book entity.Book
	db.connection.Preload("User").Find(&book, bookID)
	return book
}

func (db *bookConnection) AllBook(p dto.PaginationDTO) dto.PaginationDTO {
	var pag dto.PaginationDTO
	// pagination init attribute
	totalRow, totalPages, fromRow, toRow := 0, 0, 0, 0
	offset := p.Page * p.Limit

	var books []entity.Book
	var book entity.Book
	err := db.connection.Limit(p.Limit).Offset(offset).Preload("User").Find(&books).Error
	if err != nil {
		return pag
	}
	p.Rows = books

	countTotalRows := int64(totalRow)
	errs := db.connection.Model(book).Count(&countTotalRows).Error
	if errs != nil {
		return p
	}

	p.TotalRows = countTotalRows

	totalPages = int(math.Ceil(float64(countTotalRows)/float64(p.Limit))) - 1

	if p.Page == 0 {
		fromRow = 1
		toRow = p.Limit
	} else {
		if p.Page <= totalPages {
			fromRow = p.Page*p.Limit + 1
			toRow = (p.Page + 1) * p.Limit
		}
	}

	if toRow > totalRow {
		toRow = totalRow
	}

	p.FromRow = fromRow
	p.ToRow = toRow
	return p
}
