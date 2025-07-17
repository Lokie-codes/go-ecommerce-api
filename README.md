# Go E-Commerce API

A robust and scalable backend API for a modern e-commerce platform, built entirely in Go. This project provides the core functionalities for managing users, products, shopping carts, and orders.

## ✨ Features

  - **User Authentication**: Secure user registration and login using JWT (JSON Web Tokens).
  - **Product Management**: Full CRUD (Create, Read, Update, Delete) operations for products.
  - **Shopping Cart**: Add, view, update, and remove items from the user's cart.
  - **Order Processing**: Create orders from the shopping cart and view order history.
  - **RESTful Architecture**: Clean, predictable, and resource-oriented API endpoints.
  - **Scalable Design**: Built with a modular structure for easy extension and maintenance.

## 🛠️ Tech Stack

  - **Language**: [Go](https://golang.org/) (Golang)
  - **Web Framework**: [Gin](https://gin-gonic.com/)
  - **Database**: [PostgreSQL](https://www.postgresql.org/)
  - **ORM**: [GORM](https://gorm.io/)
  - **Authentication**: [golang-jwt/jwt](https://github.com/golang-jwt/jwt)
  - **Environment Variables**: [go-dotenv](https://github.com/joho/godotenv)
  - **Containerization**: [Docker](https://www.docker.com/) & Docker Compose

## 🚀 Getting Started

Follow these instructions to get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

You need to have the following software installed on your machine:

  - [Go](https://golang.org/doc/install) (version 1.21 or later)
  - [Docker](https://docs.docker.com/get-docker/)
  - [Docker Compose](https://docs.docker.com/compose/install/)

### Installation

1.  **Clone the repository:**

    ```sh
    git clone https://github.com/your-username/go-ecommerce-api.git
    cd go-ecommerce-api
    ```

2.  **Create an environment file:**
    Copy the example environment file and update it with your configuration.

    ```sh
    cp .env.example .env
    ```

3.  **Run the application with Docker Compose:**
    This is the recommended way to run the project locally. It will start the Go API server and a PostgreSQL database container.

    ```sh
    docker-compose up --build
    ```

The API server will be running on `http://localhost:8080`.

## ⚙️ Environment Variables

To run this project, you will need to add the following environment variables to your `.env` file:

```ini
# Server Configuration
API_PORT=8080

# Database Configuration
DB_HOST=db
DB_USER=yourusername
DB_PASSWORD=yourstrongpassword
DB_NAME=ecommerce_db
DB_PORT=5432

# JWT Configuration
JWT_SECRET_KEY=your-super-secret-key
```

## Endpoints API

Here are the main API endpoints available.

| Method | Endpoint                    | Description                       | Protected |
|--------|-----------------------------|-----------------------------------|-----------|
| `POST` | `/api/v1/register`          | Register a new user               | No        |
| `POST` | `/api/v1/login`             | Log in a user and get a JWT token | No        |
|        |                             |                                   |           |
| `GET`  | `/api/v1/products`          | Get a list of all products        | No        |
| `GET`  | `/api/v1/products/{id}`     | Get a single product by ID        | No        |
| `POST` | `/api/v1/products`          | Create a new product              | Yes (Admin) |
| `PUT`  | `/api/v1/products/{id}`     | Update a product                  | Yes (Admin) |
| `DELETE`| `/api/v1/products/{id}`     | Delete a product                  | Yes (Admin) |
|        |                             |                                   |           |
| `GET`  | `/api/v1/cart`              | Get the user's shopping cart      | Yes       |
| `POST` | `/api/v1/cart/add`          | Add an item to the cart           | Yes       |
| `DELETE`| `/api/v1/cart/remove/{itemId}` | Remove an item from the cart   | Yes       |
|        |                             |                                   |           |
| `POST` | `/api/v1/orders`            | Create an order from the cart     | Yes       |
| `GET`  | `/api/v1/orders`            | Get the user's order history      | Yes       |
| `GET`  | `/api/v1/orders/{id}`       | Get details of a specific order   | Yes       |

## 🧪 Running Tests

To run the test suite for the application, use the following command:

```sh
go test ./... -v
```

## 🤝 Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1.  Fork the Project
2.  Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3.  Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4.  Push to the Branch (`git push origin feature/AmazingFeature`)
5.  Open a Pull Request

## 📜 License

Distributed under the MIT License. See `LICENSE` for more information.
