package service

import (
	"fmt"
	"lab-go-gin-jwt/dto"
	"lab-go-gin-jwt/entity"
	"lab-go-gin-jwt/repository"
	"log"

	"github.com/mashingan/smapping"
)

type BookService interface {
	Insert(book dto.BookCreateDTO) entity.Book
	Update(book dto.BookUpdateDTO) entity.Book
	Delete(book entity.Book)
	FindAll() []entity.Book
	FindByID(bookID uint) entity.Book
	IsAllowedToEdit(userID string, bookID uint) bool
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepo,
	}
}

func (s *bookService) Insert(b dto.BookCreateDTO) entity.Book {
	book := entity.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Panicf("Failed map %v: ", err)
	}
	res := s.bookRepository.InsertBook(book)
	return res
}

func (s *bookService) Update(book dto.BookUpdateDTO) entity.Book {
	b := entity.Book{}
	err := smapping.FillStruct(&b, smapping.MapFields(&book))
	if err != nil {
		log.Panicf("Failed update %v: ", err)
	}
	res := s.bookRepository.UpdateBook(b)
	return res
}

func (s *bookService) Delete(book entity.Book) {
	s.bookRepository.DeleteBook(book)
}

func (s *bookService) FindAll() []entity.Book {
	return s.bookRepository.FindBooks()
}

func (s *bookService) FindByID(bookID uint) entity.Book {
	return s.bookRepository.FindBookByID(bookID)
}

func (s *bookService) IsAllowedToEdit(userID string, bookID uint) bool {
	b := s.bookRepository.FindBookByID(bookID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
