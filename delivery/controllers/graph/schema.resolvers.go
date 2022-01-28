package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"sirclo/entities"
	"sirclo/entities/model"
	"sirclo/util/graph/generated"
)

func (r *mutationResolver) RentBook(ctx context.Context, id *int) (*model.SuccessResponse, error) {
	dataLogin := ctx.Value("EchoContextKey")
	if dataLogin == nil {
		return nil, errors.New("unauthorized")
	} else {
		convId := ctx.Value("EchoContextKey")
		fmt.Println("id user", convId)
	}
	userId := dataLogin.(int)
	bookId := *id
	err := r.bookRepo.RentBook(bookId, userId)
	if err != nil {
		return nil, errors.New("buku tidak tersedia untuk disewa")
	}
	var response model.SuccessResponse
	response.Code = 200
	response.Message = "berhasil menyewa buku"
	return &response, nil
}

func (r *mutationResolver) ReturnBook(ctx context.Context, id *int) (*model.SuccessResponse, error) {
	dataLogin := ctx.Value("EchoContextKey")
	if dataLogin == nil {
		return nil, errors.New("unauthorized")
	} else {
		convId := ctx.Value("EchoContextKey")
		fmt.Println("id user", convId)
	}
	userId := dataLogin.(int)
	bookId := *id
	err := r.bookRepo.ReturnBook(bookId, userId)
	if err != nil {
		return nil, errors.New("buku tidak sedang disewa")
	}
	var response model.SuccessResponse
	response.Code = 200
	response.Message = "berhasil mengembalikan buku"
	return &response, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.SuccessResponse, error) {
	var user entities.User
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password
	fmt.Println(user)
	err := r.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	var response model.SuccessResponse
	response.Code = 200
	response.Message = "berhasil membuat user"
	return &response, nil
}

func (r *mutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.SuccessResponse, error) {
	dataLogin := ctx.Value("EchoContextKey")
	if dataLogin == nil {
		return nil, errors.New("unauthorized")
	} else {
		convId := ctx.Value("EchoContextKey")
		fmt.Println("id user", convId)
	}
	var book entities.Book
	book.Tittle = input.Tittle
	book.Author = input.Author
	book.Status = input.Status
	fmt.Println(book)
	err := r.bookRepo.CreateBook(book)
	if err != nil {
		return nil, err
	}
	var response model.SuccessResponse
	response.Code = 200
	response.Message = "berhasil membuat book"
	return &response, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*model.SuccessResponse, error) {
	dataLogin := ctx.Value("EchoContextKey")
	if dataLogin == nil {
		return nil, errors.New("unauthorized")
	} else {
		convId := ctx.Value("EchoContextKey")
		fmt.Println("id user", convId)
	}
	fmt.Println("id = ", id)
	err := r.userRepo.DeleteUser(id)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	var response model.SuccessResponse
	response.Code = 200
	response.Message = "berhasil menghapus user"
	return &response, nil
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id int) (*model.SuccessResponse, error) {
	dataLogin := ctx.Value("EchoContextKey")
	if dataLogin == nil {
		return nil, errors.New("unauthorized")
	} else {
		convId := ctx.Value("EchoContextKey")
		fmt.Println("id user", convId)
	}
	err := r.bookRepo.DeleteBook(id)
	if err != nil {
		return nil, fmt.Errorf("book not found")
	}
	var response model.SuccessResponse
	response.Code = 200
	response.Message = "berhasil menghapus book"
	return &response, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id int, edit model.NewUser) (*model.SuccessResponse, error) {
	dataLogin := ctx.Value("EchoContextKey")
	if dataLogin == nil {
		return nil, errors.New("unauthorized")
	} else {
		convId := ctx.Value("EchoContextKey")
		fmt.Println("id user", convId)
	}
	var user entities.User
	user.Name = edit.Name
	user.Email = edit.Email
	user.Password = edit.Password
	err := r.userRepo.EditUser(user, id)
	if err != nil {
		return nil, err
	}
	var response model.SuccessResponse
	response.Code = 200
	response.Message = "berhasil mengedit user"
	return &response, nil
}

func (r *mutationResolver) UpdateBook(ctx context.Context, id int, edit model.NewBook) (*model.SuccessResponse, error) {
	dataLogin := ctx.Value("EchoContextKey")
	if dataLogin == nil {
		return nil, errors.New("unauthorized")
	} else {
		convId := ctx.Value("EchoContextKey")
		fmt.Println("id user", convId)
	}
	var book entities.Book
	book.Tittle = edit.Tittle
	book.Author = edit.Author
	book.Status = edit.Status
	err := r.bookRepo.EditBook(book, id)
	if err != nil {
		return nil, err
	}
	var response model.SuccessResponse
	response.Code = 200
	response.Message = "berhasil mengedit book"
	return &response, nil
}

func (r *queryResolver) Login(ctx context.Context, email string, password string) (*model.LoginResponse, error) {
	token, user, err := r.authRepo.Login(email, password)
	if err != nil {
		return nil, errors.New("yang bener inputnya cuk")
	}
	var hasil model.LoginResponse
	hasil.Code = 200
	hasil.Message = "selamat anda berhasil login"
	hasil.Token = token
	var penampung model.User
	penampung.ID = &user.Id
	penampung.Name = user.Name
	penampung.Email = user.Email
	penampung.Password = user.Password
	hasil.User = &penampung
	return &hasil, nil
}

func (r *queryResolver) Product(ctx context.Context) ([]*model.Product, error) {
	responseData, err := r.productRepo.GetProduct()
	fmt.Println(responseData)

	if err != nil {
		return nil, errors.New("not found")
	}

	productResponseData := []*model.Product{}

	for _, v := range responseData {
		var product model.Product
		var user model.User
		product.ID = &v.Id
		product.Price = v.Price
		user.ID = &v.User.Id
		user.Name = v.User.Name
		user.Email = v.User.Email
		user.Password = v.User.Password
		product.User = &user
		fmt.Println("product = ", product)
		productResponseData = append(productResponseData, &product)
	}

	return productResponseData, nil
}

func (r *queryResolver) User(ctx context.Context) ([]*model.User, error) {
	dataLogin := ctx.Value("EchoContextKey")
	if dataLogin == nil {
		return nil, errors.New("unauthorized")
	} else {
		convId := ctx.Value("EchoContextKey")
		fmt.Println("id user", convId)
	}

	responseData, err := r.userRepo.GetUsers()

	if err != nil {
		return nil, errors.New("not found")
	}

	userResponseData := []*model.User{}

	for _, v := range responseData {
		theId := int(v.Id)
		userResponseData = append(userResponseData, &model.User{ID: &theId, Name: v.Name, Email: v.Email, Password: v.Password})
	}

	return userResponseData, nil
}

func (r *queryResolver) Book(ctx context.Context) ([]*model.Book, error) {
	responseData, err := r.bookRepo.GetAllBook()

	if err != nil {
		return nil, errors.New("not found")
	}

	bookResponseData := []*model.Book{}

	for _, v := range responseData {
		theId := int(v.Id)
		bookResponseData = append(bookResponseData, &model.Book{ID: &theId, Tittle: v.Tittle, Author: v.Author, Status: v.Status})
	}

	return bookResponseData, nil
}

func (r *queryResolver) BookByID(ctx context.Context, id int) (*model.Book, error) {
	fmt.Println("hasil id inputan dari param = ", id)
	responseData, err := r.bookRepo.GetBook(id)

	if err != nil {
		return nil, errors.New("not found")
	}
	fmt.Println("hasil query dari repository = ", responseData)
	var bookResponseData model.Book
	bookResponseData.ID = &responseData.Id
	bookResponseData.Tittle = responseData.Tittle
	bookResponseData.Author = responseData.Author
	bookResponseData.Status = responseData.Status
	fmt.Println("hasil =", bookResponseData)

	return &bookResponseData, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
