#!/bin/bash

# GitHub Container Registry Deployment Script
# Run this script on your server to deploy PInGO-Share

set -e

echo "🚀 PInGO-Share GitHub Container Registry Deployment"
echo "=================================================="

# Configuration
DEPLOY_PATH="/opt/pingo-share"
REPO_URL="https://raw.githubusercontent.com/R8bert/PInGO-Share/main"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

print_step() {
    echo -e "${BLUE}📋 Step $1:${NC} $2"
}

print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

# Step 1: Check prerequisites
print_step "1" "Checking prerequisites..."

if ! command -v docker &> /dev/null; then
    print_warning "Docker not found. Installing Docker..."
    curl -fsSL https://get.docker.com -o get-docker.sh
    sudo sh get-docker.sh
    sudo usermod -aG docker $USER
    print_success "Docker installed. Please log out and log back in, then run this script again."
    exit 0
fi

if ! command -v docker-compose &> /dev/null; then
    print_warning "Docker Compose not found. Installing..."
    sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    sudo chmod +x /usr/local/bin/docker-compose
    print_success "Docker Compose installed"
fi

print_success "Prerequisites check passed"

# Step 2: Create deployment directory
print_step "2" "Setting up deployment directory..."

if [ ! -d "$DEPLOY_PATH" ]; then
    sudo mkdir -p $DEPLOY_PATH
    sudo chown $USER:$USER $DEPLOY_PATH
    print_success "Created deployment directory: $DEPLOY_PATH"
else
    print_success "Deployment directory already exists"
fi

cd $DEPLOY_PATH

# Step 3: Download deployment files
print_step "3" "Downloading deployment files..."

curl -s -O $REPO_URL/docker-compose.github.yml
curl -s -O $REPO_URL/.env.production.example

if [ ! -f ".env" ]; then
    cp .env.production.example .env
    print_success "Created .env file from template"
else
    print_warning ".env file already exists, skipping creation"
fi

print_success "Deployment files downloaded"

# Step 4: Configure environment
print_step "4" "Environment configuration"

if [ ! -s ".env" ] || grep -q "your-secure-password" .env; then
    echo ""
    print_warning "⚙️  Environment Configuration Required"
    echo "The .env file needs to be configured with your specific values."
    echo ""
    echo "Required changes:"
    echo "1. DB_PASSWORD: Set a secure database password"
    echo "2. JWT_SECRET: Set a secure JWT secret (32+ characters)"
    echo "3. ALLOWED_ORIGINS: Set your domain (e.g., https://yourdomain.com)"
    echo ""
    
    read -p "Do you want to edit the .env file now? (y/n): " edit_env
    
    if [[ $edit_env =~ ^[Yy]$ ]]; then
        nano .env
    else
        echo "Please edit the .env file before continuing:"
        echo "  nano $DEPLOY_PATH/.env"
        echo ""
        echo "Then run this script again or continue manually:"
        echo "  docker-compose -f docker-compose.github.yml up -d"
        exit 1
    fi
fi

# Step 5: GitHub authentication (if needed)
print_step "5" "GitHub Container Registry authentication"

echo ""
echo "Is your GitHub repository private? (y/n)"
read -p "If unsure, check: https://github.com/R8bert/PInGO-Share/settings > " is_private

if [[ $is_private =~ ^[Yy]$ ]]; then
    echo ""
    echo "For private repositories, you need a GitHub Personal Access Token."
    echo "1. Go to: https://github.com/settings/personal-access-tokens/tokens"
    echo "2. Create a token with 'read:packages' permission"
    echo ""
    read -p "Enter your GitHub username: " github_user
    read -s -p "Enter your GitHub token: " github_token
    echo ""
    
    echo "$github_token" | docker login ghcr.io -u "$github_user" --password-stdin
    
    if [ $? -eq 0 ]; then
        print_success "Successfully authenticated with GitHub Container Registry"
    else
        print_error "Authentication failed. Please check your token and try again."
        exit 1
    fi
else
    print_success "Public repository - no authentication needed"
fi

# Step 6: Deploy the application
print_step "6" "Deploying the application..."

# Stop existing containers if any
if [ -f "docker-compose.github.yml" ]; then
    docker-compose -f docker-compose.github.yml down 2>/dev/null || true
fi

# Check if images exist in registry
print_warning "Checking if images exist in GitHub Container Registry..."
if ! docker-compose -f docker-compose.github.yml pull 2>/dev/null; then
    print_warning "Images not found in GitHub Container Registry!"
    echo ""
    echo "This usually means:"
    echo "1. GitHub Actions are still building the images"
    echo "2. GitHub Actions failed to build"
    echo "3. This is the first deployment"
    echo ""
    echo "Let's check GitHub Actions status and build locally if needed..."
    echo ""
    
    read -p "Would you like to build images locally instead? (y/n): " build_local
    
    if [[ $build_local =~ ^[Yy]$ ]]; then
        print_warning "Switching to local build method..."
        
        # Download source code
        if [ ! -d "source" ]; then
            git clone https://github.com/R8bert/PInGO-Share.git source
        else
            cd source && git pull && cd ..
        fi
        
        # Build images locally
        print_warning "Building backend image..."
        docker build -t ghcr.io/r8bert/pingo-share-backend:latest -f source/backend/Dockerfile source/
        
        print_warning "Building frontend image..."
        docker build -t ghcr.io/r8bert/pingo-share-frontend:latest source/frontend/pingo/
        
        print_success "Images built locally"
    else
        echo ""
        echo "Please:"
        echo "1. Check GitHub Actions: https://github.com/R8bert/PInGO-Share/actions"
        echo "2. Wait for build to complete"
        echo "3. Run this script again"
        echo ""
        exit 1
    fi
else
    print_success "Images pulled from GitHub Container Registry"
fi

# Start services
print_warning "Starting services..."
docker-compose -f docker-compose.github.yml up -d

# Wait for services to start
echo "Waiting for services to start..."
sleep 10

# Step 7: Verify deployment
print_step "7" "Verifying deployment..."

# Check if containers are running
if docker-compose -f docker-compose.github.yml ps | grep -q "Up"; then
    print_success "Containers are running"
    
    # Show status
    echo ""
    echo "📊 Service Status:"
    docker-compose -f docker-compose.github.yml ps
    
    # Test frontend
    if curl -s -f http://localhost > /dev/null; then
        print_success "Frontend is responding"
        
        # Get server IP
        SERVER_IP=$(curl -s ifconfig.me 2>/dev/null || curl -s icanhazip.com 2>/dev/null || echo "your-server-ip")
        
        echo ""
        echo "🎉 Deployment successful!"
        echo "📱 Your application is available at:"
        echo "   🌐 http://$SERVER_IP"
        if [ -f ".env" ] && grep -q "FRONTEND_PORT=80" .env; then
            echo "   🌐 http://$SERVER_IP:80"
        fi
        
    else
        print_warning "Frontend might still be starting up"
        echo "Check status with: docker-compose -f docker-compose.github.yml logs -f"
    fi
    
else
    print_error "Some containers failed to start"
    echo "Check logs with: docker-compose -f docker-compose.github.yml logs"
    exit 1
fi

# Step 8: Provide management commands
echo ""
echo "🔧 Management Commands:"
echo "  📊 Check status:    docker-compose -f docker-compose.github.yml ps"
echo "  📝 View logs:       docker-compose -f docker-compose.github.yml logs -f"
echo "  🔄 Restart:         docker-compose -f docker-compose.github.yml restart"
echo "  🛑 Stop:            docker-compose -f docker-compose.github.yml down"
echo "  🔄 Update:          docker-compose -f docker-compose.github.yml pull && docker-compose -f docker-compose.github.yml up -d"

echo ""
echo "📚 Full documentation: https://github.com/R8bert/PInGO-Share/blob/main/GITHUB-REGISTRY-DEPLOY.md"
echo ""
print_success "Deployment completed! 🎊"
