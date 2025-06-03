# Pokemon Integration Service

## 📋 Overview
A Go-based RESTful API service that integrates with the [PokeAPI](https://pokeapi.co/) to manage Pokemon data. This service allows you to fetch, store, and manage Pokemon information using a PostgreSQL database.

## ⭐ Features
- ✨ Integrate Pokemon data from PokeAPI
- 💾 Store Pokemon information in PostgreSQL database
- 🔄 RESTful API endpoints for Pokemon management
- 📚 Swagger documentation
- 🐳 Docker support for development environment
- 📝 Request logging middleware
- ⚠️ Error handling
- ⚙️ Environment configuration

## 🛠️ Technologies
- **Go** 1.24+
- **Gin** Web Framework
- **PostgreSQL** Database
- **Docker** & Docker Compose
- **Swagger** Documentation
- **Logrus** for logging

## 📌 Prerequisites
Before you begin, ensure you have:
- Go 1.24 or higher installed
- Docker and Docker Compose installed
- PostgreSQL (or use the provided Docker setup)

## 🚀 Quick Start
1. **Clone the repository**
```bash
git clone <repository-url>
cd pokemon-integration-service
```

2. **Set up environment variables**
```bash
cp .env.example .env
```

3. **Start PostgreSQL with Docker**
```bash
docker-compose up -d
```

4. **Run the application**
```bash
go run cmd/main.go
```

The server will start at `http://localhost:3000` 🎉

## 📖 API Documentation
### Swagger Documentation
Access the interactive API documentation at:
```
http://localhost:3000/swagger/index.html
```

### Available Endpoints

| Method | Endpoint                                  | Description                 |
| ------ | ----------------------------------------- | --------------------------- |
| GET    | `/api/v1/pokemons`                        | Get all integrated Pokemons |
| POST   | `/api/v1/pokemons/integrate/{name}`       | Integrate a new Pokemon     |
| DELETE | `/api/v1/pokemons/deletePokemonByID/{id}` | Delete a Pokemon            |

## 📁 Project Structure
```
.
├── cmd/
│   ├── main.go
│   └── docker-compose.yaml
├── internal/
│   ├── database/
│   ├── domain/
│   ├── handlers/
│   ├── middleware/
│   ├── repository/
│   └── services/
└── docs/
```

## ⚙️ Environment Variables
```env
URL_POKEAPI=https://pokeapi.co/api/v2/pokemon
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=your_database
```

## 📄 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

### 📝 Notes
- Documentation will be updated as the project evolves.