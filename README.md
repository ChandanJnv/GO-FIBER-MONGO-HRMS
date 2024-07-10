# GO-FIBER-MONGO-HRMS

A simple Human Resource Management System built using Go, Fiber, and MongoDB.

## Features
- Employee management

## Technologies
- Go
- Fiber (web framework)
- MongoDB (database)

## Setup
1. **Clone the repository:**
    ```sh
    git clone https://github.com/ChandanJnv/GO-FIBER-MONGO-HRMS.git
    ```
2. **Navigate to the project directory:**
    ```sh
    cd GO-FIBER-MONGO-HRMS
    ```
3. **Install dependencies:**
    ```sh
    go mod tidy
    ```
4. **Run the application:**
    ```sh
    go run main.go
    ```

## Configuration
- Ensure MongoDB is installed and running.
- Update the MongoDB connection string if necessary.

## API Endpoints

### Employee Management
- **Get All Employees**
  - `GET /employee`
- **Create Employee**
  - `POST /employee`
  - **Request Body:**
    ```json
    {
      "name": "John Doe",
      "salary": 50000,
      "age": 30
    }
    ```
- **Update Employee**
  - `PUT /employee/:id`
  - **Request Body:**
    ```json
    {
      "name": "John Smith",
      "salary": 60000,
      "age": 31
    }
    ```
- **Delete Employee**
  - `DELETE /employee/:id`

## Usage
- Access the application at `http://localhost:3000`
- Use the provided API endpoints to manage employees.

## Contributing
Contributions are welcome! Please create an issue or pull request.

## License
This project is licensed under the MIT License.
