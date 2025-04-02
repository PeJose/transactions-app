# Cashflow Project

This repository contains a containerized solution for the ING assessment titled "Cashflow". It is divided into two main components: `cashflow-api` and `cashflow-next`.

## Table of Contents
- [Description](#description)
- [Setup Instructions](#setup-instructions)
- [Testing Endpoints](#testing-endpoints)
- [Components](#components)
  - [Cashflow API](#cashflow-api)
  - [Cashflow Next](#cashflow-next)

---

## Description

The project is designed to manage and analyze cash flow data. It includes a backend API (`cashflow-api`) for data handling and a frontend application (`cashflow-next`) for user interaction. Both components are containerized for ease of deployment.

---

## Setup Instructions

0. Install dependencies in root, cashflow-next(frontend) and cashflow-api(backend) folders.
   ```bash
   <!-- ROOT -->
   pnpm i
   <!-- API -->
   cd cashflow-api
   go mod tidy
   cd ..
   <!-- NEXT -->
   cd cashflow-next
   pnpm i
   ```

1. Build and run the containers using Docker Compose:
   ```bash
   docker-compose up --build
   ```
   Or if you have task installed:
   ```bash
   task dcu
   ```

2. Access the frontend at `http://localhost:3000` and the API at `http://localhost:8081`.

3. Login Information:
   - Use the `/register` route in Bruno to create a new user. This route is not available through the frontend.
   - Alternatively, use the email and password provided in the Bruno environment variables for login.

---

## Testing Endpoints

For manual testing of API endpoints, the project uses [Bruno](https://www.usebruno.com/), an open-source alternative to Postman. Bruno allows you to organize and execute API requests efficiently.

### How to Use Bruno:
1. Install Bruno by following the instructions on their [official website](https://www.usebruno.com/).
2. Import the API collection provided in the `Cashflow` directory in root.
3. Use Bruno's interface to test the endpoints manually. (important notice: change environment in right top corner to Dev)

---

## Components

### Cashflow API

The backend API is built with Go and provides endpoints for managing companies, users, and cash flow data.

#### Key Features:
- RESTful API for CRUD operations.
- Database integration using GORM.
- Authentication and token generation.

#### Included Packages:
The following packages are used in `cashflow-api` (from `go.mod`):
- `github.com/gofiber/fiber/v2`: Web framework for building APIs built on top fasthttp.
- `gorm.io/gorm`: ORM for database interactions.
- `gorm.io/driver/sqlite`: SQLite driver for GORM.
- `github.com/dgrijalva/jwt-go`: JWT for authentication.
- `github.com/go-playground/validator/v10`: Input validation.
- `github.com/joho/godotenv`: Loads environment variables from a `.env` file.
- `github.com/air`: Live reload for Go applications during development.

---

### Cashflow Next

The frontend application is built with Next.js(React) and provides a user-friendly interface for interacting with the API.

#### Key Features:
- Dynamic routing and server-side rendering.
- Integration with the backend API.
- Responsive design.

#### Included Packages:
The following key packages are used in `cashflow-next` (from `package.json`):
- `next`: React framework for server-side rendering.
- `react`: Library for building user interfaces.
- `@tanstack/react-query`: Data fetching and state management.
- `tailwindcss`: Utility-first CSS framework.
- `daisyui`: Tailwind CSS components.
- `typescript`: TypeScript language support.
- `recharts`: Charting library for React.
- `react-google-charts`: Google Charts integration for React. (used for map only)
- `biome`: Code formatting and linting. (This one is installed outside the app in root folder)

