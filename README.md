# 🚀 Job Application Tracker

  A full-stack Job Application Tracking system built with **Next.js (Frontend)**, **Go (Backend - MVC architecture)**, and **PostgreSQL (Neon DB)**.
  This project helps users track job applications across different hiring stages like Applied, Interviewing, Offer, and Rejected.

  ---

## 📌 Project Overview

  This application allows users to:

  - Add new job applications
  - View all applications in a structured dashboard
  - Edit application details
  - Delete applications
  - Filter applications by status
  - Search by company name or job title

  The system follows a clean **MVC architecture on the backend** and a modular component-based structure on the frontend.

  ---

## 🛠 Tech Stack Used

### Frontend
  - Next.js (React)
  - TypeScript
  - Tailwind CSS

### Backend
  - Go (Golang)
  - Gin Web Framework
  - MVC Architecture

### Database
  - PostgreSQL (Neon DB)

  ### Other Tools
  - REST API
  - Fetch API
  - Environment Variables (.env)

  ---

## 📦 Prerequisites

  Make sure you have installed:

  - Node.js (>= 18)
  - npm or yarn
  - Go (>= 1.20)
  - PostgreSQL (Neon account or connection string)
  - Git

  ---

## ⚙️ Installation Steps

### 1. Clone the repository

  ```bash
  git clone https://github.com/ArbinBhasimaOfficial/JobTracker.git
  cd JobTracker

### 2. Backend Setup

  ```bash
  cd server
  go mod tidy

  Create .env file

  ```bash
  DATABASE_URL=your_neon_postgres_url
  PORT=8080

  Run Backend

  ```bash
  go run main.go

  Backend runs on:

  ```bash
  http://localhost:8080

### 3 Frontend Setup (Next.js)

  ```bash
  cd client
  npm install

  Create .env.local:

  ```bash
  NEXT_PUBLIC_API_URL=http://localhost:8080

  Run frontend

  ```bash
  npm run dev

  frontend runs on:
  http://localhost:3000


### How to run in development mode

  Start backend

  ```bash
  cd server
  go run main.go


  Start frontend

  ```bash
  cd client
  npm run dev


### How to run test

  Currently, no automated tests are included in this version.
  Future improvements may include:

  - Unit tests for Go services
  - API integration tests
  - Frontend component testing with Jest/React Testing Library


### Environment Variables

  Backend

  ```bash
  DATABASE_URL=postgresql://user:password@host/dbname?sslmode=require
  PORT=8080

  Frontend

  ```bash
  NEXT_PUBLIC_API_URL=http://localhost:8080


### API Documentation

  Base URL

  ```bash
    http://localhost:8080

  Endpoints

  Get all applications
    GET /applications

  Supports:

    ?status=Applied
    ?search=google

  Get single application
    GET /applications/:id

  Create application
    POST /applications

  Update application
    PATCH /applications/:id

  Delete application
    DELETE /applications/:id



### Data Model

```bash
{
  "id": "uuid",
  "company_name": "Google",
  "job_title": "Backend Intern",
  "job_type": "Internship | Full-time | Part-time",
  "status": "Applied | Interviewing | Offer | Rejected",
  "applied_date": "2026-06-19",
  "notes": "Optional notes",
  "created_at": "timestamp",
  "updated_at": "timestamp"
}


### Deployment

  Frontend: Vercel
  Backend: Render / Fly.io
  Database: Neon PostgreSQL


### Future Improvements

  Authentication (JWT / OAuth)
  Pagination
  Drag-and-drop pipeline (Kanban board)
  Email reminders
  Analytics dashboard

### Author
  Arbin Bhasima


### License
  This project is for educational and internship submission purposes.
