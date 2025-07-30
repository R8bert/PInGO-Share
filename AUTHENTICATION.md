# PinGO File Transfer - Authentication System

## Overview

PinGO now includes a complete user authentication system with account management and upload history tracking.

## Features Added

### Backend (Go + PostgreSQL)
- **User Registration & Login**: Secure user authentication with JWT tokens
- **Password Security**: Bcrypt hashing for password storage
- **Session Management**: JWT tokens with cookie support
- **Upload Tracking**: All uploads are linked to user accounts with expiration dates
- **Database Integration**: PostgreSQL with users and uploads tables
- **API Security**: All upload endpoints now require authentication

### Frontend (Vue.js + TypeScript)
- **Authentication Pages**: Login and registration forms
- **Account Dashboard**: User profile and upload history management
- **Route Guards**: Protected routes requiring authentication
- **Upload History**: View, manage, and delete previous uploads
- **Responsive Design**: Mobile-friendly interface

## Database Configuration

The system connects to PostgreSQL with these settings:
- **Host**: 192.168.31.180
- **Port**: 5555
- **Database**: pingo_db
- **User**: pingo
- **Password**: pingo_filetrasnfer

### Database Tables

#### Users Table
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

#### Uploads Table
```sql
CREATE TABLE uploads (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    upload_id VARCHAR(255) UNIQUE NOT NULL,
    files TEXT NOT NULL,
    total_size BIGINT NOT NULL,
    email VARCHAR(255),
    download_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP + INTERVAL '7 days'
);
```

## API Endpoints

### Authentication
- `POST /register` - Create new user account
- `POST /login` - User login
- `POST /logout` - User logout
- `GET /me` - Get current user info

### File Management (Protected)
- `POST /upload` - Upload files (requires auth)
- `GET /uploads` - Get user's upload history
- `DELETE /uploads/:id` - Delete user's upload

### Public Endpoints
- `GET /download/:id` - Download files (public)
- `GET /files/:id` - Get file metadata (public)
- `GET /settings` - Get app settings
- `POST /settings/save` - Save app settings

## Security Features

1. **JWT Authentication**: Secure token-based authentication
2. **Password Hashing**: Bcrypt with salt rounds
3. **HTTP-only Cookies**: Secure token storage
4. **CORS Protection**: Configured for frontend origin
5. **Route Protection**: All upload functionality requires authentication
6. **File Cleanup**: Automatic cleanup of expired uploads

## Usage

### Running the Application

1. **Start Backend**:
   ```bash
   cd backend
   go build -o pingo-backend main.go
   ./pingo-backend
   ```
   Server runs on `http://localhost:8081`

2. **Start Frontend**:
   ```bash
   cd frontend/pingo
   npm run dev
   ```
   Frontend runs on `http://localhost:5173`

### User Workflow

1. **Registration**: New users must create an account
2. **Login**: Existing users sign in with email/password
3. **Upload Files**: Only authenticated users can upload files
4. **Manage Uploads**: Users can view their upload history in the account page
5. **Share Files**: Generate shareable links for downloads
6. **Automatic Cleanup**: Files expire after 7 days and are automatically removed

## Development Notes

- JWT secret should be changed in production
- Database connection uses environment-specific settings
- File uploads are limited by the maxUploadSize setting
- All user sessions expire after 24 hours
- Upload history shows file status (active/expired)

## Technology Stack

- **Backend**: Go, Gin, PostgreSQL, JWT, Bcrypt
- **Frontend**: Vue 3, TypeScript, Tailwind CSS, Axios
- **Database**: PostgreSQL with automatic table creation
- **Authentication**: JWT tokens with cookie support
