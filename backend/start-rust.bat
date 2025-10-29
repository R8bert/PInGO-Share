@echo off
REM PInGO Share Rust Backend - Quick Start Script for Windows

echo.
echo 🦀 PInGO Share Rust Backend Setup
echo ==================================
echo.

REM Check if Rust is installed
where cargo >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo ❌ Rust is not installed!
    echo 📦 Please install Rust from: https://rustup.rs/
    echo.
    echo Download and run: https://win.rustup.rs/
    pause
    exit /b 1
)

echo ✅ Rust is installed
rustc --version
cargo --version
echo.

REM Check for .env file
if not exist ".env" (
    echo ⚠️  No .env file found!
    echo 📝 Creating .env file from template...
    (
        echo # Database Configuration
        echo DB_HOST=localhost
        echo DB_PORT=5432
        echo DB_USER=user
        echo DB_PASSWORD=pass
        echo DB_NAME=filesharing
        echo.
        echo # JWT Secret ^(MUST be at least 32 characters^)
        echo JWT_SECRET=please-change-this-to-a-secure-secret-at-least-32-chars-long
        echo.
        echo # Server Configuration
        echo SERVER_PORT=8080
        echo ALLOWED_ORIGINS=*
        echo.
        echo # Logging
        echo RUST_LOG=info
    ) > .env
    echo ✅ Created .env file
    echo ⚠️  Please update JWT_SECRET and database credentials in .env
    echo.
)

REM Build the project
echo 🔨 Building Rust backend...
echo    ^(First build may take a few minutes^)
cargo build --release

if %ERRORLEVEL% NEQ 0 (
    echo.
    echo ❌ Build failed!
    echo Please check the error messages above.
    pause
    exit /b 1
)

echo.
echo ✅ Build complete!
echo.
echo 🚀 Starting server...
echo    Server will be available at: http://localhost:8080
echo    Press Ctrl+C to stop
echo.

REM Run the server
cargo run --release

pause
