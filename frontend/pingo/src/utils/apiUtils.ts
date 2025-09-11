// Get the API base URL
export const getApiBaseUrl = () => {
  // In production (when served from same domain), API calls should go to /api/
  if (window.location.hostname !== 'localhost' && window.location.hostname !== '127.0.0.1') {
    console.log("Production mode detected for hostname:", window.location.hostname)
    return '/api'
  }
  // In development, use localhost
  return 'http://localhost:8080'
}

// Utility function to get the correct API base URL
export const getApiUrl = (path: string = '') => {
  // If path starts with /, remove it to avoid double slashes
  const cleanPath = path.startsWith('/') ? path.slice(1) : path
  
  const baseUrl = getApiBaseUrl()
  return `${baseUrl}/${cleanPath}`
}

// Get full URL for assets (avatars, files, etc.)
export const getAssetUrl = (path: string) => {
  // If path starts with /, remove it to avoid double slashes
  const cleanPath = path.startsWith('/') ? path.slice(1) : path
  
  const baseUrl = getApiBaseUrl()
  return `${baseUrl}/${cleanPath}`
}

// Get the current domain's base URL
export const getBaseUrl = () => {
  return `${window.location.protocol}//${window.location.host}`
}
