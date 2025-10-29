# PInGO-Share
<img src="https://github.com/user-attachments/assets/c996ea3a-bbb5-4ab9-8f1b-fd7f565fa1f9" alt="DeWatermark ai image" width="300">
--
Self-hosted file sharing platform that gives you complete control over your data. Upload, share, and manage files on your own infrastructure without relying on third-party services.

## Why PInGO-Share?

Take back control of your file sharing. Whether you're a privacy-conscious individual, a small business, or just tired of file size limits and subscription fees, PInGO-Share puts you in charge.

**Key Features:**
- No file size restrictions (limited only by your storage)
- Complete privacy - your files stay on your server
- No monthly fees or subscriptions
- Clean, modern web interface
- Secure user authentication
- Easy deployment with Docker

## Getting Started

### Quick Setup

Clone this repository and run:

```bash
docker-compose up -d
```

Your file sharing platform will be available at `http://localhost:3000`

### Production Deployment

For a production setup on your server:

```bash
./deploy-github-registry.sh
```

This script handles everything - Docker installation, image building, and service deployment.

## Configuration

Before deploying to production, edit the `.env` file:

```bash
DB_PASSWORD=choose_a_strong_password
JWT_SECRET=your_secret_key_minimum_32_characters
ALLOWED_ORIGINS=https://your-domain.com
```

## Self-Hosting Benefits

- **Privacy**: Your files never leave your infrastructure
- **Cost-effective**: No per-user or storage fees
- **Customizable**: Modify the platform to fit your needs
- **Reliable**: No dependence on external services
- **Scalable**: Grows with your storage requirements

## Architecture

Built with modern, reliable technologies:

- Vue.js frontend for a responsive user experience
- Go backend for high performance and reliability
- PostgreSQL for robust data management
- Docker for easy deployment and scaling

## Deployment Options

**Development**: Use `docker-compose.yml` for local testing

**Production**: Use `docker-compose.github.yml` with the deployment script

**Advanced**: Customize the Docker configurations for your specific needs

## Support

This is an open-source project. Feel free to contribute improvements or report issues.

## License

Licensed under the terms included in the LICENSE file.
