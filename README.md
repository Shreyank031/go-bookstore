# Go Bookstore

This is a RESTful API built with Go, Gorilla Mux router, and GORM (an ORM library for Golang). It provides functionality to manage books, including creating, retrieving, updating, and deleting book records. The API interacts with a database (configured using GORM) to store and retrieve book data.

## Features

- Get all books
- Get a book by ID
- Create a new book
- Delete a book by ID
- Update a book by ID

## Technologies Used

- Go
- Gorilla Mux (for routing)
- GORM (ORM library for database interaction)

## Installation

1. Clone the repository:

```
    https://github.com/Shreyank031/go-bookstore.git
```

2. Navigate to the project directory:

```
    cd go-bookstore
```

3. Download the project dependencies:

```
    go mod download 
```

4. Configure the database connection by modifying the `config.Connect()` function in the `pkg/config` package. Set the appropriate database connection string and other parameters based on your database setup.


## Usage

1. Start the server:

```
    go run main.go
```

## API Endpoints

### Get All Books

- Endpoint: `/books`
- Method: `GET`
- Description: Retrieves a list of all books.

### Create a New Book

- Endpoint: `/books`
- Method: `POST`
- Description: Creates a new book.
- Request Body:
  ```json
  {
    "name": "Book Name",
    "author": "Author Name",
    "publication": "Publication Name"
  }
### Get Book by ID

- Endpoint: `/books/{bookId}`
- Method: `GET`
- Description: Retrieves a book by its ID.
- Path Parameter: bookId (the ID of the book)

### Update a Book

- Endpoint: `/books/{bookId}`
- Method: `PUT`
- Description: Updates a book by its ID.
- Path Parameter: bookId (the ID of the book)
- Request Body:

```
{
  "name": "Updated Book Name",
  "author": "Updated Author Name",
  "publication": "Updated Publication Name"
}
```
### Delete a Book

- Endpoint: `/books/{bookId}`
- Method: `DELETE`
- Description: Deletes a book by its ID.
- Path Parameter: bookId (the ID of the book)
