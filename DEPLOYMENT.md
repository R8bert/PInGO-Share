# Server Deployment Guide for PInGO-Share

## Option 1: Deploy with Docker Registry (Recommended)

### Step 1: Build and Push Images

1. **Login to your Docker registry:**
   ```bash
   docker login your-registry.com
   ```

2. **Build and push images:**
   ```bash
   ./build-and-push.sh your-registry.com/your-username latest
   ```

### Step 2: Deploy on Server

1. **Copy files to server:**
   ```bash
   scp docker-compose.production.yml .env.production.example your-server:/opt/pingo-share/
   ```

2. **Setup environment on server:**
   ```bash
   ssh your-server
   cd /opt/pingo-share
   cp .env.production.example .env
   nano .env  # Edit with your values
   ```

3. **Deploy:**
   ```bash
   docker-compose -f docker-compose.production.yml up -d
   ```

## Option 2: Deploy with Portainer Stack

### Step 1: Prepare Images

1. **Build images locally:**
   ```bash
   docker build -t pingo-share-backend:latest -f ./backend/Dockerfile .
   docker build -t pingo-share-frontend:latest ./frontend/pingo/
   ```

2. **Save images to tar files:**
   ```bash
   docker save pingo-share-backend:latest | gzip > pingo-share-backend.tar.gz
   docker save pingo-share-frontend:latest | gzip > pingo-share-frontend.tar.gz
   ```

3. **Transfer to server and load:**
   ```bash
   scp *.tar.gz your-server:/tmp/
   ssh your-server
   docker load < /tmp/pingo-share-backend.tar.gz
   docker load < /tmp/pingo-share-frontend.tar.gz
   ```

### Step 2: Create Portainer Stack

1. **Login to Portainer** (usually at `https://your-server:9443`)

2. **Go to Stacks → Add Stack**

3. **Upload `portainer-stack.yml`** or copy its contents

4. **Edit environment variables** in the stack:
   - `JWT_SECRET`: Generate a secure secret
   - `DB_PASSWORD`: Set a secure database password
   - `ALLOWED_ORIGINS`: Set your domain
   - Port mappings if needed

5. **Deploy the stack**

## Option 3: Direct Server Build

### Copy entire project to server:
```bash
rsync -av --exclude=node_modules --exclude=.git . your-server:/opt/pingo-share/
```

### Build and run on server:
```bash
ssh your-server
cd /opt/pingo-share
cp .env.production.example .env
nano .env  # Edit with your values
docker-compose -f docker-compose.production.yml up -d --build
```

## Environment Variables to Configure

```bash
# Database
DB_USER=user
DB_PASSWORD=your-secure-password
DB_NAME=filesharing

# Security
JWT_SECRET=your-jwt-secret-min-32-chars

# Domain/CORS
ALLOWED_ORIGINS=https://your-domain.com
FRONTEND_PORT=80

# Registry (if using)
REGISTRY_URL=your-registry.com/username
TAG=latest
```

## Security Considerations

1. **Use HTTPS:** Set up reverse proxy (nginx/traefik) with SSL
2. **Firewall:** Only expose ports 80/443
3. **Strong passwords:** Use secure database passwords
4. **JWT Secret:** Use a strong, random JWT secret (32+ characters)
5. **Updates:** Regularly update images for security patches

## Monitoring

### Check service status:
```bash
docker-compose -f docker-compose.production.yml ps
```

### View logs:
```bash
docker-compose -f docker-compose.production.yml logs -f
```

### Backup database:
```bash
docker-compose -f docker-compose.production.yml exec postgres_pingo pg_dump -U user filesharing > backup.sql
```

## Troubleshooting

### Common issues:
1. **Port conflicts:** Change FRONTEND_PORT in .env
2. **Database connection:** Check DB_PASSWORD and DB_USER
3. **CORS errors:** Verify ALLOWED_ORIGINS includes your domain
4. **File uploads:** Ensure upload volumes are writable

### Debug commands:
```bash
# Check container status
docker ps

# Check logs
docker logs container-name

# Enter container
docker exec -it container-name /bin/sh
```
