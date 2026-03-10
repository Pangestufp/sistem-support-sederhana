# Simple Support Ticket System

A simple support ticket management system that allows users to create tickets, attach files, and process approval workflows.

## Tech Stack

- **Backend**: Gin (Go)
- **Frontend**: React
- **Database**: MySQL
- **Containerization**: Docker

## Features

- Create support tickets
- Upload attachments
- Ticket workflow approval
- Ticket review and return
- Dashboard for assigned jobs
- Pagination for ticket browsing

## Project Structure
```
.
├── backend/
├── frontend/
└── docker-compose.yml
```

## Getting Started

### 1. Clone the repository

git clone https://github.com/Pangestufp/sistem-support-sederhana.git

### 2. Create .env files for both frontend and backend based on the provided .env.example.

### 3. Run with Docker

docker compose up --build

### 4. Import the schema from Database/schema.sql

## Notes
Some data are handled directly in the backend/database.
The frontend only focuses on the ticket submission and management features.
