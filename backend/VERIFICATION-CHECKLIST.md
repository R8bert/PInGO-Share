# Rust Backend Verification Checklist

Use this checklist to verify that the Rust backend is working correctly.

## ✅ Setup Verification

- [ ] Rust installed (`rustc --version` works)
- [ ] Cargo installed (`cargo --version` works)
- [ ] `.env` file created with all required variables
- [ ] PostgreSQL is running and accessible
- [ ] JWT_SECRET is at least 32 characters

## ✅ Build Verification

- [ ] `cargo build` completes without errors
- [ ] `cargo build --release` creates optimized binary
- [ ] Binary exists at `target/release/pingo-share-backend`
- [ ] No compilation warnings (or acceptable warnings only)

## ✅ Server Startup

- [ ] Server starts without errors
- [ ] Logs show "Database connected successfully"
- [ ] Logs show "Database migrations completed"
- [ ] Logs show "Starting server on 0.0.0.0:8080"
- [ ] Server responds to requests

## ✅ API Endpoints - Public

### Settings
- [ ] `GET /api/settings` returns settings JSON
- [ ] Response includes theme, max_upload_size, etc.

## ✅ API Endpoints - Authentication

### Registration
- [ ] `POST /api/register` creates new user
- [ ] Returns token and user info
- [ ] Password is hashed in database
- [ ] First user is automatically admin
- [ ] Subsequent users are not admin
- [ ] Registration can be disabled via settings

### Login
- [ ] `POST /api/login` with valid credentials works
- [ ] Returns token and user info
- [ ] Invalid credentials returns 401
- [ ] Token is set as cookie

### Logout
- [ ] `POST /api/logout` clears auth cookie
- [ ] Returns success message

### Profile
- [ ] `GET /api/me` with valid token returns user info
- [ ] Without token returns 401
- [ ] With invalid token returns 401

### Avatar
- [ ] `POST /api/avatar` with image uploads successfully
- [ ] Avatar path is returned
- [ ] Avatar file exists on disk
- [ ] Old avatar is deleted when new one uploaded
- [ ] Invalid file types are rejected

## ✅ API Endpoints - File Operations

### Upload
- [ ] `POST /api/upload` with valid token uploads file(s)
- [ ] Returns download_url, files list, count
- [ ] Files are saved to uploads directory
- [ ] File size limits are enforced
- [ ] File type restrictions work
- [ ] Validity period is set correctly
- [ ] Email is optional and saved
- [ ] Blocked users cannot upload

### Get Uploads
- [ ] `GET /api/uploads` returns user's uploads
- [ ] Shows all upload details correctly
- [ ] Returns empty array for new users
- [ ] Deleted uploads show is_deleted flag

### Delete Upload
- [ ] `DELETE /api/uploads/{id}` soft deletes upload
- [ ] Physical files are deleted from disk
- [ ] Deletion log is created
- [ ] Cannot delete already deleted upload
- [ ] Cannot delete another user's upload

### Toggle Availability
- [ ] `PUT /api/uploads/{id}/availability` updates status
- [ ] is_available flag changes
- [ ] Unavailable files cannot be downloaded by others
- [ ] Owner can still download unavailable files

### Update Expiration
- [ ] `PUT /api/uploads/{id}/expiration` updates expires_at
- [ ] New expiration date is calculated correctly
- [ ] Cannot exceed max_validity setting
- [ ] Blocked users cannot update

## ✅ API Endpoints - Download

### Download Files
- [ ] `GET /api/download/{id}` downloads files
- [ ] Single file downloads directly
- [ ] Multiple files download as ZIP
- [ ] ZIP contains all files with correct names
- [ ] Unavailable files return 410 (owner can access)
- [ ] Invalid ID returns 404

### Individual File
- [ ] `GET /api/file/{id}/{filename}` downloads specific file
- [ ] Correct file is served
- [ ] Content-Type is set correctly
- [ ] Path traversal attempts are blocked

### File Metadata
- [ ] `GET /api/files/{id}` returns metadata
- [ ] Shows uploader info (username, avatar)
- [ ] Lists all files with sizes
- [ ] Shows expiration date
- [ ] Unavailable to non-owners returns 410

## ✅ API Endpoints - Reverse Sharing

### Create Token
- [ ] `POST /api/reverse-tokens` creates token
- [ ] Returns token, name, max_uses
- [ ] Token is saved to database
- [ ] Expiration is set correctly (if specified)
- [ ] Max uses defaults to -1 (unlimited)

### List Tokens
- [ ] `GET /api/reverse-tokens` returns user's tokens
- [ ] Shows used_count correctly
- [ ] Shows all token details

### Delete Token
- [ ] `DELETE /api/reverse-tokens/{id}` deletes token
- [ ] Token is removed from database
- [ ] Cannot delete another user's token

### Upload via Token
- [ ] `POST /api/reverse-upload/{token}` uploads files
- [ ] Files are saved to uploader's account
- [ ] Token usage count increments
- [ ] Expired tokens are rejected
- [ ] Max uses limit is enforced
- [ ] Invalid tokens return 404

## ✅ API Endpoints - Admin

### Update Settings
- [ ] `POST /api/admin/settings` updates settings (admin only)
- [ ] Non-admin users get 403
- [ ] Theme updates correctly
- [ ] Max upload size updates
- [ ] Logo upload works
- [ ] Background upload works
- [ ] Color settings update
- [ ] Navbar title updates

### Get Statistics
- [ ] `GET /api/admin/stats` returns stats (admin only)
- [ ] Shows total_users count
- [ ] Shows total_uploads count
- [ ] Shows storage_used sum
- [ ] Non-admin users get 403

### List Users
- [ ] `GET /api/admin/users` returns all users (admin only)
- [ ] Shows upload_count per user
- [ ] Shows storage_used per user
- [ ] Shows last_activity
- [ ] Shows is_blocked status
- [ ] Non-admin users get 403

### Block/Unblock User
- [ ] `POST /api/admin/users/{id}/block` blocks user
- [ ] Blocked users cannot upload
- [ ] Cannot block admin users
- [ ] Unblocking works correctly

### Quick Settings
- [ ] `POST /api/admin/quick-settings` updates setting
- [ ] allowRegistration toggle works
- [ ] Other settings supported

## ✅ Security Features

### Authentication
- [ ] JWT tokens expire after 24 hours
- [ ] Invalid tokens are rejected
- [ ] Expired tokens are rejected
- [ ] Password hashing uses bcrypt

### File Security
- [ ] Executable files are blocked (.exe, .bat, etc.)
- [ ] Path traversal attempts fail
- [ ] Filename sanitization works
- [ ] Files are validated before saving

### Headers
- [ ] X-Content-Type-Options: nosniff
- [ ] X-Frame-Options: DENY
- [ ] X-XSS-Protection header present
- [ ] Strict-Transport-Security header present

### CORS
- [ ] CORS headers are set correctly
- [ ] Allowed origins are respected
- [ ] Credentials are supported

## ✅ Background Tasks

### Expired Files Cleanup
- [ ] Cleanup task runs every hour
- [ ] Expired files are handled per expiration_action
- [ ] "delete" action removes files and marks deleted
- [ ] "unavailable" action marks as unavailable
- [ ] Logs show cleanup activity

## ✅ Database

### Schema
- [ ] All tables created successfully
- [ ] Indexes are created
- [ ] Foreign keys work correctly
- [ ] Default settings are inserted

### Queries
- [ ] User queries work
- [ ] Upload queries work
- [ ] Token queries work
- [ ] Admin queries work
- [ ] Deletion log queries work

## ✅ Performance

### Response Times
- [ ] Settings endpoint < 50ms
- [ ] Login endpoint < 100ms
- [ ] Upload endpoint < 500ms (depends on file size)
- [ ] Download endpoint < 200ms

### Memory
- [ ] Idle memory < 100MB
- [ ] Memory stable under load
- [ ] No memory leaks detected

### Concurrency
- [ ] Handles multiple simultaneous requests
- [ ] Database connection pool works
- [ ] No race conditions

## ✅ Compatibility

### Go API Compatibility
- [ ] All Go endpoints have Rust equivalents
- [ ] Request formats are identical
- [ ] Response formats are identical
- [ ] Status codes match
- [ ] Error messages are similar

### Frontend Compatibility
- [ ] Frontend works without changes
- [ ] Login flow works
- [ ] File upload works
- [ ] File download works
- [ ] Admin panel works

## ✅ Docker

### Build
- [ ] Docker image builds successfully
- [ ] Image size is reasonable (~50-100MB)
- [ ] No build errors or warnings

### Run
- [ ] Container starts successfully
- [ ] Can connect to external PostgreSQL
- [ ] Environment variables work
- [ ] Volumes mount correctly
- [ ] Ports are exposed properly

## ✅ Documentation

- [ ] README-RUST.md is complete
- [ ] MIGRATION-GUIDE.md is clear
- [ ] GO-VS-RUST.md shows comparisons
- [ ] All code is well-commented
- [ ] API examples are provided

## ✅ Production Readiness

### Logging
- [ ] Info logs are helpful
- [ ] Error logs are detailed
- [ ] Log levels work (info, debug, error)
- [ ] No excessive logging

### Error Handling
- [ ] All errors return appropriate status codes
- [ ] Error messages are user-friendly
- [ ] Internal errors don't expose secrets
- [ ] Database errors are handled gracefully

### Graceful Shutdown
- [ ] Server shuts down cleanly on SIGTERM
- [ ] Active requests complete before shutdown
- [ ] Database connections close properly

### Configuration
- [ ] All config via environment variables
- [ ] Reasonable defaults provided
- [ ] Required variables validated on startup
- [ ] Invalid config fails fast with clear error

## 🎯 Final Verification

- [ ] All tests above pass
- [ ] No critical errors in logs
- [ ] Performance is acceptable
- [ ] Frontend works correctly
- [ ] Ready for production deployment

---

## Notes

Document any issues found:

```
Issue: _______________________________________
Status: [ ] Fixed  [ ] Documented  [ ] Investigating
Details: _____________________________________
```

---

**Verification completed by:** _______________  
**Date:** _______________  
**Version:** _______________  
**Result:** [ ] Pass  [ ] Fail  [ ] Needs work
