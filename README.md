# Song Library API 🎶

Welcome to the **Song Library API**! This is a RESTful API that allows you to retrieve detailed information about songs, including metadata, lyrics, and links to external platforms like YouTube. With a simple set of endpoints, you can access the song's release date, lyrics, and much more.

## 🌟 Features

- Search for songs by group (artist/band) and title.
- Retrieve detailed song information (release date, lyrics, and more).
- Access song lyrics with pagination support.
- Comprehensive API documentation via Swagger.
- Simple and easy-to-use interface for song retrieval.

## 🚀 Getting Started

### Prerequisites

To run the Song Library API, you'll need the following installed on your machine:

## 🛠 Technologies Used

- ![Go](https://habrastorage.org/r/w1560/getpro/habr/upload_files/ed9/0e8/0bd/ed90e80bdbf970bc1101a4ac66e3221a.png) **Go** (v1.16+): Main programming language for backend logic.  
  [GitHub - Go](https://github.com/golang/go)

- ![Gin](https://i.ytimg.com/vi/vDIAwtGU9LE/maxresdefault.jpg) **Gin Framework**: Web framework for handling HTTP requests.  
  [GitHub - Gin Framework](https://github.com/gin-gonic/gin)

- ![Swagger](https://avatars.dzeninfra.ru/get-zen_brief/6638195/pub_627f733875980b241b9c3dd2_627f733875980b241b9c3dd3/scale_1200) **Swagger**: For automatic generation of API documentation.  
  [GitHub - Swagger](https://github.com/swagger-api/swagger-core)

- ![PostgreSQL](https://d15shllkswkct0.cloudfront.net/wp-content/blogs.dir/1/files/2012/06/slonik-outline.jpg) **PostgreSQL** (if applicable): For database management (if you're storing song data).  
  [GitHub - PostgreSQL](https://github.com/postgres/postgres)

- ![Lumberjack](https://i.ytimg.com/vi/WsNK1ZjNiUI/maxresdefault.jpg) **Lumberjack**: For log rotation and management.  
  [GitHub - Lumberjack](https://github.com/natefinch/lumberjack)

- ![GORM](https://i.ytimg.com/vi/04XyLJF1TDQ/maxresdefault.jpg) **GORM**: ORM for Golang to interact with databases like PostgreSQL.  
  [GitHub - GORM](https://github.com/go-gorm/gorm)

### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/shyn1ck/song-library.git
    ```

2. Navigate into the project directory:
    ```bash
    cd song-library
    ```

3. Create a `.env` file in the root directory and add your PostgreSQL password as follows:
    ```env
    DB_PASSWORD=postgres
    ```

4. Initialize Swagger documentation:
    ```bash
    swag init -g cmd/main.go
    ```

5. Install project dependencies:
    ```bash
    go mod tidy
    ```

6. Run the application:
    ```bash
    go run cmd/main.go
    ```

7. Once the server is running, you can access the Swagger API documentation at:
   [http://localhost:8181/swagger/index.html#/](http://localhost:8181/swagger/index.html#/)


### Contact
If you have any questions or suggestions, feel free to reach out to me:

Telegram: [@parvizjon_hasanov](https://t.me/parvizjon_hasanov)


Project Idea:  
[Effective Mobile](https://effective-mobile.ru/)


