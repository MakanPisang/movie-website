# MovieFlix 🎥

MovieFlix is a movie streaming website built with Golang, utilizing the Gin framework, PostgreSQL database, and REST API. It features secure user authentication using JWT, allowing users to browse, search, and manage their favorite movies.

## Features

- 🎬 **Browse Movies**: Discover movies by genre, popularity, or search for specific titles.
- 🔐 **User Authentication**: Secure login and registration using JWT.
- 📋 **Favorites Management**: Add movies to your favorites list.
- 🔍 **Search**: Quickly find movies by title or keyword.
- ⚡ **Fast & Scalable**: Built with Golang and the Gin framework for performance.

## Tech Stack

- **Backend**: Go (Gin framework)
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Tokens)
- **API**: RESTful API

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/movieflix.git
   cd movieflix
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up the database:

   - Create a PostgreSQL database.
   - Run the migrations located in the `migrations/` folder.

4. Configure environment variables:
   Create a `.env` file in the root directory and add the following:

   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=yourusername
   DB_PASSWORD=yourpassword
   DB_NAME=moviedb
   JWT_SECRET=your_jwt_secret
   ```

5. Run the server:

   ```bash
   go run main.go
   ```

6. Access the application:
   Open your browser and go to `http://localhost:8080`.

## API Endpoints

### Authentication

- **POST /auth/register**: Register a new user.
- **POST /auth/login**: Authenticate a user and return a JWT.

### Movies

- **GET /movies**: Get a list of movies.
- **GET /movies/:id**: Get details of a specific movie.

## Contact

Feel free to reach out if you have any questions or feedback:

- Email: dputro18@gmail.com
- GitHub: [MakanPisang](https://github.com/yourusername)

---

⭐ Don't forget to star the repository if you find this project helpful!
