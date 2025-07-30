import { ref, computed } from 'vue'
import axios from 'axios'

interface User {
  id: number
  username: string
  email: string
  created_at?: string
}

interface Upload {
  id: number
  upload_id: string
  files: string
  total_size: number
  email: string
  download_url: string
  created_at: string
  expires_at: string
}

const user = ref<User | null>(null)
const token = ref<string | null>(null)
const isLoading = ref(false)
const uploads = ref<Upload[]>([])

// Setup axios defaults
axios.defaults.baseURL = 'http://localhost:8080'
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
    if (error.response?.status === 401) {
      // Clear auth state if unauthorized
      user.value = null
      token.value = null
      localStorage.removeItem('auth_token')
    }
    return Promise.reject(error)
  }
)

export const useAuth = () => {
  const isAuthenticated = computed(() => !!user.value)

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
    try {
      if (!token.value) return false
      
      const response = await axios.get('/me')
      user.value = response.data.user
      return true
    } catch (error) {
      clearAuth()
      return false
    }
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
      await axios.delete(`/uploads/${uploadId}`)
      uploads.value = uploads.value.filter(upload => upload.upload_id !== uploadId)
      return { success: true, message: 'Upload deleted successfully' }
    } catch (error: any) {
      const message = error.response?.data?.error || 'Failed to delete upload'
      return { success: false, message }
    }
  }

  // Initialize auth state
  loadTokenFromStorage()

  return {
    user: computed(() => user.value),
    token: computed(() => token.value),
    uploads: computed(() => uploads.value),
    isAuthenticated,
    isLoading: computed(() => isLoading.value),
    register,
    login,
    logout,
    fetchCurrentUser,
    fetchUploads,
    deleteUpload
  }
}
