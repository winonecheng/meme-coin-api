# meme-coin-api
A simple RESTful API built with Golang (Gin framework) for managing meme coins, containerized with Docker. This project implements CRUD operations and a "poke" feature to increment a meme coin's popularity score, using SQLite as the database.

## Project Structure
- `main.go`: Application entry point.
- `models/meme_coin.go`: Defines the `MemeCoin` data model.
- `handlers/meme_coin.go`: Contains API endpoint logic.
- `handlers/meme_coin_test.go`: Unit tests for the API endpoints.
- `db/db.go`: Manages database connection and operations with SQLite.
- `Dockerfile`: Docker configuration for building and running the app.
- `go.mod`: Go module dependency file.
- `go.sum`: Go module checksum file.

## Instructions for Running the Application

### Running Locally
1. **Install Go**: Ensure you have Go 1.24 or later installed on your system.
2. **Clone the Repository**: 
    ```bash
    git clone https://github.com/your-username/meme-coin-api.git
    cd meme-coin-api
    ```
3. **Install Dependencies**:
    ```bash
    go mod download
    ```
4. **Run the Application**:
    ```bash
    go run main.go
    ```
5. **Access the API**: The application will be available at `http://localhost:8080`.

### Running in a Docker Container
1. **Build the Docker Image**:
    ```bash
    docker build -t meme-coin-api .
    ```
2. **Run the Docker Container**:
    ```bash
    docker run -d -p 8080:8080 -v $(pwd)/meme_coins.db:/app/meme_coins.db meme-coin-api
    ```
    - `-d`: Runs the container in detached mode.
    - `-p 8080:8080`: Maps port 8080 on the host to 8080 in the container.
    - `-v $(pwd)/meme_coins.db:/app/meme_coins.db`: Mounts the SQLite database file to persist data outside the container.

3. **Access the API**: The application will be available at `http://localhost:8080`.

4. **Stop the container**:
    ```bash
    docker stop $(docker ps -q --filter ancestor=meme-coin-api)
    ```

## API Endpoints
- **POST** `/meme-coins`: Create a new meme coin.
- **GET** `/meme-coins/:id`: Retrieve a meme coin by ID.
- **PUT** `/meme-coins/:id`: Update a meme coin's description.
- **DELETE** `/meme-coins/:id`: Delete a meme coin by ID.
- **POST** `/meme-coins/:id/poke`: Increment the popularity score of a meme coin.
