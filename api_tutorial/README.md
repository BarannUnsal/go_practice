# Go Lang Tutorial - API Example

This tutorial showcases my learning process in Go Lang and how to build and interact with a simple API

---

### 1. **List All Books**
- **Endpoint**: `GET /books`
- **Description**: Retrieves the list of all books in the collection.

```bash
curl -X GET http://localhost:8080/books
```

### 2. **Get Book by ID**
- **Endpoint**: `GET /books/:id`
- **Description**: Retrieves details of a specific book by its ID.

```bash
curl -X GET http://localhost:8080/books/<id>
```

Replace `<id>` with the ID of the book you want to retrieve.

---

### 3. **Add a New Book**
- **Endpoint**: `POST /books`
- **Description**: Adds a new book to the collection.

```bash
curl -X POST http://localhost:8080/books \
-H "Content-Type: application/json" \
-d '{
  "id": "4",
  "title": "New Book",
  "author": "Author Name",
  "quantity": 5
}'
```

---

### 4. **Checkout a Book**
- **Endpoint**: `PATCH /checkout`
- **Description**: Decreases the quantity of a book by one (borrow it).

```bash
curl -X PATCH http://localhost:8080/checkout?id=<id>
```

Replace `<id>` with the ID of the book you want to check out.

---

### 5. **Return a Book**
- **Endpoint**: `PATCH /return`
- **Description**: Increases the quantity of a book by one (return it).

```bash
curl -X PATCH http://localhost:8080/return?id=<id>
```

Replace `<id>` with the ID of the book you want to return.

