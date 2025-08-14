# 🔑 KeyShoes Blog 🚀

Welcome to KeyShoes, a modern and stylish blog platform! 📝 This project provides a space for you to share your thoughts, ideas, and stories with the world.

## ✨ Features

-   **Create & Manage Posts**: Easily write, edit, and delete your blog posts.
-   **User Authentication**: Secure user registration and login system.
-   **Responsive Design**: A clean and beautiful interface that looks great on any device.

## 🛠️ Tech Stack

-   **Frontend**: [React](https://react.dev/) with [Vite](https://vitejs.dev/)
-   **Backend**: [Go](https://go.dev/)
-   **Database**: [PostgreSQL](https://www.postgresql.org/)
-   **Containerization**: [Docker](https://www.docker.com/)

## 🚀 Getting Started

Follow these instructions to get the project up and running on your local machine.

### Prerequisites

Make sure you have the following installed:

-   [Docker](https://docs.docker.com/get-docker/)
-   [Docker Compose](https://docs.docker.com/compose/install/)

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/keyshoes.git
cd keyshoes
```

### 2. Set Up Environment Variables

Create a `.env` file in the root directory by copying the example file:

```bash
cp .env.example .env
```

Now, open the `.env` file and replace the placeholder values with your own credentials.

### 3. Build and Run with Docker

To build and run the entire application, use Docker Compose:

```bash
docker-compose up --build
```

This command will:

-   Build the Docker images for the frontend and backend services.
-   Start the containers.
-   Initialize the database with the required tables from `migrations/migrations.sql`.

Once everything is running, you can access the application:

-   **Frontend**: [http://localhost:8081](http://localhost:8081)
-   **Backend**: [http://localhost:8080](http://localhost:8080)

## 💻 Frontend Development

If you want to work on the frontend separately, you can run it locally without Docker.

### 1. Navigate to the `web` Directory

```bash
cd web
```

### 2. Install Dependencies

```bash
npm install
```

### 3. Run the Development Server

```bash
npm run dev
```

This will start the Vite development server, and you can view the frontend at [http://localhost:5173](http://localhost:5173).

---

Happy blogging! 🎉
