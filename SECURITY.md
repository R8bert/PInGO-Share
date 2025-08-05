# 🔒 PInGO-Share Security Guide

## 🚨 CRITICAL - Before Production Deployment

### 1. Environment Configuration

**MUST DO**: Create a `.env` file from `.env.example`:

```bash
cp .env.example .env
```

**Edit `.env` file and update these critical values:**

```env
# Generate a strong JWT secret (32+ characters)
JWT_SECRET=your-actual-super-secure-jwt-secret-key-here-minimum-32-chars

# Set strong database credentials
DB_PASSWORD=your-very-secure-database-password-123!

# Set your production domain
ALLOWED_ORIGINS=https://yourdomain.com,https://www.yourdomain.com

# For development only
# ALLOWED_ORIGINS=http://localhost:5173,http://localhost:3000
```

### 2. Generate Secure JWT Secret

Use one of these methods to generate a secure JWT secret:

```bash
# Method 1: Using OpenSSL
openssl rand -base64 64

# Method 2: Using Python
python3 -c "import secrets; print(secrets.token_urlsafe(64))"

# Method 3: Using Node.js
node -e "console.log(require('crypto').randomBytes(64).toString('base64'))"
```

### 3. Database Security

Update `docker-compose.yml` environment variables:

```yaml
environment:
  POSTGRES_PASSWORD: your-secure-database-password-here
  DB_PASSWORD: your-secure-database-password-here
```

## 🛡️ Security Features Implemented

### ✅ Authentication & Authorization
- JWT-based authentication with configurable secret
- Password hashing using bcrypt
- Admin middleware protection
- User blocking system
- Session management

### ✅ Input Validation & Sanitization
- SQL injection prevention using parameterized queries
- File type validation (MIME type + extension)
- Filename sanitization (directory traversal prevention)
- File size limits (per file and total)
- Request body validation

### ✅ Rate Limiting
- Login attempts: 10 per minute per IP
- Registration: 5 per minute per IP
- File uploads: 20 per minute per IP
- Memory-based rate limiting

### ✅ Security Headers
- `X-Content-Type-Options: nosniff`
- `X-Frame-Options: DENY`
- `X-XSS-Protection: 1; mode=block`
- `Strict-Transport-Security` (HTTPS)
- `Content-Security-Policy`

### ✅ CORS Protection
- Configurable allowed origins
- Credential support with secure origins
- Environment-based configuration

### ✅ File Upload Security
- Configurable allowed file types
- Individual and total size limits
- Secure file storage with UUID prefixes
- Directory traversal prevention

### ✅ Container Security
- Multi-stage Docker builds
- Non-root user execution
- Read-only containers where possible
- Network isolation
- Health checks
- Security options (`no-new-privileges`)

## 🚀 Production Deployment

### 1. Environment Setup

```bash
# Copy and configure environment
cp .env.example .env
# Edit .env with your production values

# Ensure uploads directory exists
mkdir -p uploads
chmod 755 uploads
```

### 2. Docker Deployment

```bash
# Build and start services
docker-compose up -d

# Check logs
docker-compose logs -f backend
docker-compose logs -f postgres_pingo

# Check health
docker-compose ps
```

### 3. Reverse Proxy Setup (Recommended)

Use Nginx or Caddy as a reverse proxy:

**Nginx configuration:**
```nginx
server {
    listen 443 ssl http2;
    server_name yourdomain.com;
    
    ssl_certificate /path/to/cert.pem;
    ssl_certificate_key /path/to/key.pem;
    
    location / {
        proxy_pass http://127.0.0.1:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

**Caddy configuration:**
```
yourdomain.com {
    reverse_proxy 127.0.0.1:8080
}
```

### 4. Firewall Configuration

```bash
# Allow only necessary ports
ufw allow 22    # SSH
ufw allow 80    # HTTP
ufw allow 443   # HTTPS
ufw deny 8080   # Block direct backend access
ufw deny 5432   # Block direct database access
ufw enable
```

## 🔍 Security Monitoring

### Log Monitoring
- Monitor authentication failures
- Track file upload patterns
- Watch for rate limit violations
- Database connection monitoring

### Regular Maintenance
- Update Docker images regularly
- Monitor security advisories
- Backup database regularly
- Review access logs

## ⚠️ Security Warnings

### Development vs Production
- Never use development settings in production
- Always use HTTPS in production
- Never expose database ports publicly
- Use strong, unique passwords

### File Upload Considerations
- Monitor upload directory size
- Implement virus scanning if needed
- Consider file retention policies
- Backup uploaded content

### Database Security
- Use connection pooling limits
- Monitor for unusual queries
- Regular security updates
- Encrypted backups

## 🆘 Incident Response

If you suspect a security breach:

1. **Immediate Actions:**
   - Change JWT secret immediately
   - Revoke all active sessions
   - Check access logs
   - Monitor system resources

2. **Investigation:**
   - Review authentication logs
   - Check file uploads
   - Analyze network traffic
   - Examine database queries

3. **Recovery:**
   - Update passwords
   - Apply security patches
   - Restore from clean backup if needed
   - Notify users if required

## 📞 Security Support

For security issues:
- Review this guide thoroughly
- Check application logs
- Update to latest version
- Follow security best practices

Remember: Security is an ongoing process, not a one-time setup!
