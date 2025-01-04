# Go Lang Tutorial - API Example (PostgreSQL Included)

This tutorial provides a simple guide to interacting with a Go Lang API connected to a PostgreSQL database. Using this project, I aim to practice and understand key concepts of API development and database integration. Below are the steps to interact with the API, including retrieving, adding, updating, and deleting items. Each section includes a description and the necessary `curl` command for testing.

### 2. **Get Item by ID**
- **Endpoint**: `GET /api/stock/{id}`
- **Description**: Retrieves details of a specific stock item by its ID.

```bash
curl -X GET http://localhost:8080/api/stock/<id>
```

Replace `<id>` with the ID of the item you want to retrieve.

---

### 3. **Add a New Item**
- **Endpoint**: `POST /api/newstock`
- **Description**: Adds a new stock item to the database.

```bash
curl -X POST http://localhost:8080/api/newstock \
-H "Content-Type: application/json" \
-d '{
  "name": "facebook",
  "price": 15,
  "company": "twitter"
}'
```

---

### 4. **Update an Item**
- **Endpoint**: `PUT /api/stock/{id}`
- **Description**: Updates details of an existing stock item by its ID.

```bash
curl -X PUT http://localhost:8080/api/stock/<id> \
-H "Content-Type: application/json" \
-d '{
  "name": "google",
  "price": 20,
  "company": "alphabet"
}'
```

Replace `<id>` with the ID of the item you want to update.

---

### 5. **Delete an Item**
- **Endpoint**: `DELETE /api/deletestock/{id}`
- **Description**: Deletes a stock item from the database by its ID.

```bash
curl -X DELETE http://localhost:8080/api/deletestock/<id>
```

Replace `<id>` with the ID of the item you want to delete.

---

### 6. **Get All Items**
- **Endpoint**: `GET /api/stock`
- **Description**: Retrieves all stock items from the database.

```bash
curl -X GET http://localhost:8080/api/stock
```

---

This API integrates PostgreSQL to store and manage data efficiently. The database connection is established using the `github.com/lib/pq` driver and managed via Go's `database/sql` package. Happy coding!