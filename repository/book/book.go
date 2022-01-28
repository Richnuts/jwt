package book

import (
	"database/sql"
	"fmt"
	"sirclo/entities"
	"time"
)

type BookRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (br *BookRepository) GetAllBook() ([]entities.Book, error) {
	var books []entities.Book
	result, err := br.db.Query("select id, tittle, author, status genre from books")
	if err != nil {
		return books, err
	}
	defer result.Close()
	for result.Next() {
		var book entities.Book
		err := result.Scan(&book.Id, &book.Tittle, &book.Author, &book.Status)
		if err != nil {
			return books, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (br *BookRepository) GetBook(id int) (entities.Book, error) {
	var book entities.Book
	result, err := br.db.Query("select id, tittle, author, status from books where id = ?", id)
	if err != nil {
		return book, err
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&book.Id, &book.Tittle, &book.Author, &book.Status)
		if err != nil {
			return book, fmt.Errorf("book not found")
		}
	}
	return book, nil
}

func (br *BookRepository) CreateBook(book entities.Book) error {
	result, err := br.db.Exec("INSERT INTO books(tittle, author, status) VALUES(?,?,?)", book.Tittle, book.Author, book.Status)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("book not created")
	}
	return nil
}

func (br *BookRepository) DeleteBook(id int) error {
	result, err := br.db.Exec("delete from books where id = ?", id)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("book not found")
	}
	return nil
}

func (br *BookRepository) EditBook(book entities.Book, id int) error {
	result, err := br.db.Exec("UPDATE books SET tittle= ?, author= ?, status= ? WHERE id = ?", book.Tittle, book.Author, book.Status, id)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("book not found")
	}
	return nil
}

// 1 adalah avaible (true)
// 0 adalah tidak tersedia (false)

func (br *BookRepository) RentBook(bookId int, userId int) error {
	var status bool
	// mengecheck status apakah bisa dipinjam
	checkStatus, err_status := br.db.Query("select status from books where id = ?", bookId)
	if err_status != nil {
		return err_status
	}
	for checkStatus.Next() {
		err := checkStatus.Scan(&status)
		if err != nil {
			return err
		}
	}
	if status == false {
		return fmt.Errorf("buku tidak tersedia")
	}
	// buat status sewa di table rent
	result, err := br.db.Exec("INSERT INTO rents(book_id, user_id, start_date) VALUES(?,?,?)", bookId, userId, time.Now())
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("rent not created")
	}
	// ubah status di table books
	result_rent, err := br.db.Exec("UPDATE books SET status= ? WHERE id = ?", false, bookId)
	if err != nil {
		return err
	}
	mengubah2, _ := result_rent.RowsAffected()
	if mengubah2 == 0 {
		return fmt.Errorf("book not found")
	}
	return nil
}

func (br *BookRepository) ReturnBook(bookId int, userId int) error {
	var status bool
	var rentId int
	// mengecheck status peminjaman
	checkStatus, err_status := br.db.Query("select id, status from rents where book_id = ? && user_id = ?", bookId, userId)
	if err_status != nil {
		return err_status
	}
	for checkStatus.Next() {
		err := checkStatus.Scan(&rentId, &status)
		if err != nil {
			return err
		}
	}
	if status == false {
		return fmt.Errorf("buku tidak sedang dipinjam / tidak ada penyewaan di id buku tersebut")
	}
	// kalo true
	// edit status sewa kita tutup
	// buat status sewa di table rent
	result, err := br.db.Exec("UPDATE rents SET status= ?, end_date = ? WHERE id = ?", 0, time.Now(), rentId)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("rent not created")
	}
	// ubah status di table books
	result_rent, err := br.db.Exec("UPDATE books SET status= ? WHERE id = ?", true, bookId)
	if err != nil {
		return err
	}
	mengubah2, _ := result_rent.RowsAffected()
	if mengubah2 == 0 {
		return fmt.Errorf("book not found")
	}
	return nil
}
