// Get the API base URL
export const getApiBaseUrl = () => {
  // In production (when served from same domain), use relative paths
  if (window.location.hostname !== 'localhost' && window.location.hostname !== '127.0.0.1') {
    console.log("The hostname is www." + window.location.hostname)
    // return error code in console and return root path
    console.error("API base URL is not set for production")
    return '/'
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
