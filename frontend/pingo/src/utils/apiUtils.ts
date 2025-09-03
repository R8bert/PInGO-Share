// Utility function to get the correct API base URL
export const getApiUrl = (path: string = '') => {
  // If path starts with /, remove it to avoid double slashes
  const cleanPath = path.startsWith('/') ? path.slice(1) : path
  
  // Since axios has baseURL configured, just return the path
  // axios will automatically prepend the correct base URL
  return `/${cleanPath}`
}

// Get full URL for assets (avatars, files, etc.)
export const getAssetUrl = (path: string) => {
  // If path starts with /, remove it to avoid double slashes
  const cleanPath = path.startsWith('/') ? path.slice(1) : path
  
  // Since axios has baseURL configured, just return the path
  // axios will automatically prepend the correct base URL
  return `/${cleanPath}`
}

// Get the current domain's base URL
export const getBaseUrl = () => {
  return `${window.location.protocol}//${window.location.host}`
}
