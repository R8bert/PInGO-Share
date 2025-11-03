// Get the API base URL
export const getApiBaseUrl = () => {
  // In production (when served from same domain), API calls should go directly to the endpoint
  if (window.location.hostname !== 'localhost' && window.location.hostname !== '127.0.0.1') {
    console.log("Production mode detected for hostname:", window.location.hostname)
    return ''
  }
  // In development, use localhost
  return 'http://localhost:8080'
}

// Utility function to get the correct API base URL
export const getApiUrl = (path: string = '') => {
  // If path starts with /, remove it to avoid double slashes
  const cleanPath = path.startsWith('/') ? path.slice(1) : path
  
  // Always prepend /api to the path (backend routes are under /api scope)
  const apiPath = `api/${cleanPath}`
  
  // In production, use relative path
  if (window.location.hostname !== 'localhost' && window.location.hostname !== '127.0.0.1') {
    const fullUrl = `/${apiPath}`
    console.log(`API URL constructed: ${fullUrl}`)
    return fullUrl
  }
  
  // In development, use full base URL with /api prefix
  const baseUrl = getApiBaseUrl()
  const fullUrl = `${baseUrl}/${apiPath}`
  console.log(`API URL constructed: ${fullUrl}`)
  return fullUrl
}

// Get full URL for assets (avatars, files, etc.)
export const getAssetUrl = (path: string) => {
  // If path starts with /, remove it to avoid double slashes
  const cleanPath = path.startsWith('/') ? path.slice(1) : path
  
  // In production, assets also go through /api proxy
  if (window.location.hostname !== 'localhost' && window.location.hostname !== '127.0.0.1') {
    return `/api/${cleanPath}`
  }
  
  // In development, use full base URL
  const baseUrl = getApiBaseUrl()
  return `${baseUrl}/${cleanPath}`
}

// Get the current domain's base URL
export const getBaseUrl = () => {
  return `${window.location.protocol}//${window.location.host}`
}
