# 🐳 PInGO-Share Portainer Stack Deployment Guide

## Step-by-Step Portainer Deployment

### Prerequisites
- Portainer installed and running on your server
- Access to Portainer web interface (usually `https://your-server:9443`)

### Step 1: Login to Portainer
1. Open your web browser
2. Navigate to your Portainer instance: `https://your-server-ip:9443`
3. Login with your Portainer credentials

### Step 2: Create the Stack
1. **Navigate to Stacks**:
   - Click on "Stacks" in the left sidebar
   - Click "Add stack" button

2. **Configure the Stack**:
   - **Name**: `pingo-share`
   - **Build method**: Choose "Web editor"

3. **Copy the Stack Configuration**:
   Copy and paste the entire content from `portainer-stack.yml`:

```yaml
services:
  postgres_pingo:
    image: postgres:15-alpine
    environment:
      POSTGRES_DB: filesharing
      POSTGRES_USER: "user"
      POSTGRES_PASSWORD: "your-secure-database-password-here"
      POSTGRES_INITDB_ARGS: "--auth-host=scram-sha-256"
    volumes:
      - postgres_data:/var/lib/postgresql/data
    restart: unless-stopped
    networks:
      - pingo-network
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U user"]
      interval: 30s
      timeout: 10s
      retries: 3
    security_opt:
      - no-new-privileges:true
    read_only: true
    tmpfs:
      - /tmp
      - /var/run/postgresql

  backend:
    image: ghcr.io/r8bert/pingo-share-backend:latest
    environment:
      - JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
      - DB_HOST=postgres_pingo
      - DB_PORT=5432
      - DB_USER=user
      - DB_PASSWORD=your-secure-database-password-here
      - DB_NAME=filesharing
      - DB_SSLMODE=disable
      - ALLOWED_ORIGINS=https://your-domain.com
      - GIN_MODE=release
    volumes:
      - uploads:/uploads
      - avatars:/avatars
      - backgrounds:/backgrounds
      - logos:/logos
    depends_on:
      postgres_pingo:
        condition: service_healthy
    restart: unless-stopped
    networks:
      - pingo-network
    security_opt:
      - no-new-privileges:true

  frontend:
    image: ghcr.io/r8bert/pingo-share-frontend:latest
    ports:
      - "80:80"
    depends_on:
      - backend
    restart: unless-stopped
    networks:
      - pingo-network
    security_opt:
      - no-new-privileges:true

networks:
  pingo-network:
    driver: bridge

volumes:
  postgres_data:
  uploads:
  avatars:
  backgrounds:
  logos:
```

### Step 3: Customize Environment Variables

**IMPORTANT**: Before deploying, change these values in the stack editor:

1. **Database Password**:
   ```yaml
   POSTGRES_PASSWORD: "MySecurePassword123!"
   DB_PASSWORD: "MySecurePassword123!"
   ```

2. **JWT Secret** (must be 32+ characters):
   ```yaml
   JWT_SECRET: "MyVerySecureJWTSecretKey32CharsMin"
   ```

3. **Domain/CORS**:
   ```yaml
   ALLOWED_ORIGINS: "https://yourdomain.com,https://www.yourdomain.com"
   ```

4. **Port** (optional - change if port 80 is in use):
   ```yaml
   ports:
     - "8080:80"  # Change 80 to 8080 if needed
   ```

### Step 4: Deploy the Stack
1. Scroll down in the stack editor
2. Click "Deploy the stack"
3. Wait for Portainer to pull images and start containers

### Step 5: Monitor Deployment
1. **Check Stack Status**:
   - Go to "Stacks" → Click on "pingo-share"
   - All services should show "running"

2. **Check Container Logs**:
   - Click on individual containers to view logs
   - Look for any errors in backend/frontend logs

### Step 6: Access Your Application
- **Default**: `http://your-server-ip`
- **Custom Port**: `http://your-server-ip:8080` (if you changed the port)

## Troubleshooting

### Images Not Found
If Portainer can't pull the images:
1. The images are public, so no authentication needed
2. Check your internet connection
3. Try pulling manually: `docker pull ghcr.io/r8bert/pingo-share-backend:latest`

### Port Conflicts
If port 80 is already in use:
1. Edit the stack
2. Change `"80:80"` to `"8080:80"`
3. Update the stack

### Database Connection Issues
1. Check that both `POSTGRES_PASSWORD` and `DB_PASSWORD` match
2. Verify database container is healthy
3. Check backend logs for connection errors

### CORS Errors
1. Update `ALLOWED_ORIGINS` to include your domain
2. Include both `http` and `https` versions
3. Add `http://your-server-ip` for direct IP access

## Stack Management

### Update Application
1. Go to Stacks → pingo-share
2. Click "Editor"
3. No changes needed to YAML
4. Click "Update the stack"
5. Portainer will pull latest images and restart

### View Logs
1. Go to Containers
2. Click on container name
3. Click "Logs" tab

### Backup Data
Your data is stored in Docker volumes:
- `postgres_data` - Database
- `uploads` - User uploaded files
- `avatars`, `backgrounds`, `logos` - App assets

## Success Checklist
- ✅ Stack deployed without errors
- ✅ All 3 containers showing "running"
- ✅ Frontend accessible at http://your-server-ip
- ✅ No errors in container logs
- ✅ Can access the application interface

Your PInGO-Share platform is now running! 🎉
