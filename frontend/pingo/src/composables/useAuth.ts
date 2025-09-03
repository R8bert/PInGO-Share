import { ref, computed } from 'vue'
import axios from 'axios'

interface User {
  id: number
  username: string
  email: string
  is_admin?: boolean
  avatar?: string
  created_at?: string
}

interface AdminUser extends User {
  isBlocked: boolean
  uploadCount?: number
  storageUsed?: number
}

interface Upload {
  id: number
  upload_id: string
  files: string
  total_size: number
  email: string
  download_url: string
  created_at: string
  expires_at: string | null
  is_available: boolean
  is_reverse: boolean
  reverse_token?: string
  is_deleted: boolean
  deleted_at: string | null
}

const user = ref<User | null>(null)
const token = ref<string | null>(null)
const isLoading = ref(false)
const uploads = ref<Upload[]>([])
const adminUsers = ref<AdminUser[]>([])
const isInitializing = ref(false)
const initPromise = ref<Promise<boolean> | null>(null)

// Setup axios defaults - automatically detect API base URL
const getApiBaseUrl = () => {
  // In production (when served from same domain), use relative paths
  if (window.location.hostname !== 'localhost' && window.location.hostname !== '127.0.0.1') {
    return '/api'
  }
  // In development, use localhost
  return 'http://localhost:8080'
}

axios.defaults.baseURL = getApiBaseUrl()
axios.defaults.withCredentials = true

// Add request interceptor to include token
axios.interceptors.request.use(
  (config) => {
    if (token.value) {
      config.headers.Authorization = `Bearer ${token.value}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// Add response interceptor to handle auth errors
axios.interceptors.response.use(
  (response) => response,
  (error) => {
    // Only clear auth state for specific 401 errors from auth endpoints
    if (error.response?.status === 401 && 
        (error.config?.url?.includes('/me') || 
         error.config?.url?.includes('/login') || 
         error.config?.url?.includes('/logout'))) {
      // Add a small delay to prevent race conditions during rapid refreshes
      setTimeout(() => {
        if (!isInitializing.value) {
          user.value = null
          token.value = null
          localStorage.removeItem('auth_token')
        }
      }, 100)
    }
    return Promise.reject(error)
  }
)

export const useAuth = () => {
  const isAuthenticated = computed(() => !!user.value)
  const isAdmin = computed(() => user.value?.is_admin === true)

  const loadTokenFromStorage = () => {
    const savedToken = localStorage.getItem('auth_token')
    if (savedToken) {
      token.value = savedToken
    }
  }

  const saveTokenToStorage = (newToken: string) => {
    localStorage.setItem('auth_token', newToken)
    token.value = newToken
  }

  const clearAuth = () => {
    user.value = null
    token.value = null
    uploads.value = []
    localStorage.removeItem('auth_token')
  }

  const register = async (username: string, email: string, password: string) => {
    try {
      isLoading.value = true
      const response = await axios.post('/register', {
        username,
        email,
        password
      })
      
      const { token: newToken, user: userData } = response.data
      saveTokenToStorage(newToken)
      user.value = userData
      
      return { success: true, message: 'Registration successful!' }
    } catch (error: any) {
      const message = error.response?.data?.error || 'Registration failed'
      return { success: false, message }
    } finally {
      isLoading.value = false
    }
  }

  const login = async (email: string, password: string) => {
    try {
      isLoading.value = true
      const response = await axios.post('/login', {
        email,
        password
      })
      
      const { token: newToken, user: userData } = response.data
      saveTokenToStorage(newToken)
      user.value = userData
      
      return { success: true, message: 'Login successful!' }
    } catch (error: any) {
      const message = error.response?.data?.error || 'Login failed'
      return { success: false, message }
    } finally {
      isLoading.value = false
    }
  }

  const logout = async () => {
    try {
      await axios.post('/logout')
    } catch (error) {
      console.error('Logout error:', error)
    } finally {
      clearAuth()
    }
  }

  const fetchCurrentUser = async () => {
    // If already initializing, return the existing promise
    if (initPromise.value) {
      return initPromise.value
    }

    // If no token, don't make a request
    if (!token.value) {
      return false
    }

    // Create a new initialization promise
    initPromise.value = (async () => {
      try {
        isInitializing.value = true
        const response = await axios.get('/me')
        user.value = response.data.user
        return true
      } catch (error: any) {
        // Only clear auth if it's a definitive 401 (unauthorized)
        if (error.response?.status === 401) {
          clearAuth()
        }
        return false
      } finally {
        isInitializing.value = false
        initPromise.value = null
      }
    })()

    return initPromise.value
  }

  const fetchUploads = async () => {
    try {
      const response = await axios.get('/uploads')
      uploads.value = response.data.uploads || []
      return { success: true }
    } catch (error: any) {
      const message = error.response?.data?.error || 'Failed to fetch uploads'
      return { success: false, message }
    }
  }

  const deleteUpload = async (uploadId: string) => {
    try {
      console.log('Deleting upload:', uploadId)
      await axios.delete(`/uploads/${uploadId}`)
      console.log('Upload deleted successfully, refreshing uploads list')
      
      // Refresh the entire uploads list to get updated data including deletion_reason
      await fetchUploads()
      
      return { success: true, message: 'Upload deleted successfully' }
    } catch (error: any) {
      const message = error.response?.data?.error || 'Failed to delete upload'
      return { success: false, message }
    }
  }

  const saveAdminSettings = async (formData: FormData) => {
    try {
      isLoading.value = true
      const response = await axios.post('/admin/settings', formData, {
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      })
      return response.data
    } catch (error: any) {
      throw new Error(error.response?.data?.error || 'Failed to save settings')
    } finally {
      isLoading.value = false
    }
  }

  const getSettings = async () => {
    try {
      const response = await axios.get('/settings')
      return response.data
    } catch (error: any) {
      throw new Error(error.response?.data?.error || 'Failed to fetch settings')
    }
  }

  const fetchAdminUsers = async () => {
    try {
      isLoading.value = true
      const response = await axios.get('/admin/users')
      adminUsers.value = response.data || []
      return { success: true }
    } catch (error: any) {
      const message = error.response?.data?.error || 'Failed to fetch admin users'
      return { success: false, message }
    } finally {
      isLoading.value = false
    }
  }

  const toggleUserBlock = async (userId: number, isBlocked: boolean) => {
    try {
      await axios.post(`/admin/users/${userId}/block`, { blocked: isBlocked })
      // Update local state
      const userIndex = adminUsers.value.findIndex(u => u.id === userId)
      if (userIndex !== -1) {
        adminUsers.value[userIndex].isBlocked = isBlocked
      }
      return { success: true, message: `User ${isBlocked ? 'blocked' : 'unblocked'} successfully` }
    } catch (error: any) {
      const message = error.response?.data?.error || 'Failed to update user status'
      return { success: false, message }
    }
  }

  // Initialize auth state
  loadTokenFromStorage()

  return {
    user: computed(() => user.value),
    token: computed(() => token.value),
    uploads: computed(() => uploads.value),
    adminUsers: computed(() => adminUsers.value),
    isAuthenticated,
    isAdmin,
    isLoading: computed(() => isLoading.value),
    register,
    login,
    logout,
    fetchCurrentUser,
    fetchUploads,
    deleteUpload,
    saveAdminSettings,
    getSettings,
    fetchAdminUsers,
    toggleUserBlock
  }
}
