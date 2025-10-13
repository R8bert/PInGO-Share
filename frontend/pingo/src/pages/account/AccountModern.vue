<template>
  <!-- Windows XP Desktop Background -->
  <div class="xp-account-page" :class="{ 'theme-dark': isDark }">
    <div class="xp-account-container">
      
      <!-- XP Window - User Profile -->
      <div class="xp-window user-profile-window">
        <!-- Title Bar -->
        <div class="xp-title-bar">
          <div class="title-bar-left">
            <svg class="title-icon" width="16" height="16" viewBox="0 0 16 16" fill="none">
              <circle cx="8" cy="5" r="3" fill="white"/>
              <path d="M2 14C2 10 4 8 8 8C12 8 14 10 14 14" fill="white"/>
            </svg>
            <span class="title-text">User Account - {{ user?.username }}</span>
          </div>
          <div class="title-bar-buttons">
            <button class="title-btn minimize" title="Minimize">
              <span></span>
            </button>
            <button class="title-btn maximize" title="Maximize">
              <svg width="10" height="10" viewBox="0 0 10 10">
                <rect x="1" y="1" width="8" height="8" fill="none" stroke="currentColor"/>
              </svg>
            </button>
            <button class="title-btn close" title="Close" @click="$router.push('/')">
              <svg width="10" height="10" viewBox="0 0 10 10">
                <path d="M1 1L9 9M9 1L1 9" stroke="currentColor" stroke-width="2"/>
              </svg>
            </button>
          </div>
        </div>

        <!-- Window Content -->
        <div class="xp-window-content">
          <!-- User Info Header -->
          <div class="xp-user-header">
            <div class="user-avatar-section">
              <div class="avatar-frame">
                <img 
                  v-if="user?.avatar" 
                  :src="getAssetUrl(user.avatar)" 
                  :alt="user.username"
                  class="avatar-image"
                  @error="handleAvatarError"
                />
                <UserIcon v-else class="avatar-placeholder" />
                <button 
                  @click="openAvatarUpload"
                  class="xp-button avatar-change-btn"
                  title="Change avatar"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 9a2 2 0 012-2h.93a2 2 0 001.664-.89l.812-1.22A2 2 0 0110.07 4h3.86a2 2 0 011.664.89l.812 1.22A2 2 0 0018.07 7H19a2 2 0 012 2v9a2 2 0 01-2 2H5a2 2 0 01-2-2V9z"></path>
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 13a3 3 0 11-6 0 3 3 0 016 0z"></path>
                  </svg>
                  Change
                </button>
                <input 
                  ref="avatarInput" 
                  type="file" 
                  accept=".png,.jpg,.jpeg,.gif,image/png,image/jpeg,image/gif" 
                  @change="handleAvatarUpload" 
                  class="hidden"
                />
              </div>
              <div class="user-info">
                <h2 class="username">{{ user?.username }}</h2>
                <div class="user-meta">
                  <span class="meta-item">
                    <CalendarIcon class="meta-icon" />
                    Member since {{ formatDate(user?.created_at) }}
                  </span>
                  <span v-if="isAdmin" class="meta-item admin-badge">
                    <ShieldCheckIcon class="meta-icon" />
                    Administrator
                  </span>
                </div>
              </div>
            </div>
          </div>

          <!-- Stats Group Box -->
          <div class="xp-group-box">
            <div class="group-box-title">Account Statistics</div>
            <div class="stats-grid">
              <!-- Total Uploads -->
              <div class="stat-item">
                <div class="stat-icon uploads">
                  <CloudArrowUpIcon class="icon" />
                </div>
                <div class="stat-info">
                  <div class="stat-label">Total Uploads</div>
                  <div class="stat-value">{{ uploads.length }}</div>
                </div>
              </div>

              <!-- Total Size -->
              <div class="stat-item">
                <div class="stat-icon size">
                  <ArchiveBoxIcon class="icon" />
                </div>
                <div class="stat-info">
                  <div class="stat-label">Total Size</div>
                  <div class="stat-value">{{ formatTotalSize() }}</div>
                </div>
              </div>

              <!-- Active Files -->
              <div class="stat-item">
                <div class="stat-icon active">
                  <ClockIcon class="icon" />
                </div>
                <div class="stat-info">
                  <div class="stat-label">Active Files</div>
                  <div class="stat-value">{{ activeUploads.length }}</div>
                </div>
              </div>
            </div>
          </div>

          <!-- XP Tab Control -->
          <div class="xp-tab-control">
            <!-- Tab Headers -->
            <div class="xp-tab-headers">
              <button
                @click="activeTab = 'uploads'"
                class="xp-tab"
                :class="{ 'active': activeTab === 'uploads' }"
              >
                <FolderIcon class="tab-icon" />
                <span>My Uploads</span>
              </button>
              <button
                @click="activeTab = 'reverse'"
                class="xp-tab"
                :class="{ 'active': activeTab === 'reverse' }"
              >
                <ShareIcon class="tab-icon" />
                <span>Reverse Share</span>
              </button>
              <button
                v-if="isAdmin"
                @click="activeTab = 'statistics'"
                class="xp-tab"
                :class="{ 'active': activeTab === 'statistics' }"
              >
                <ChartBarIcon class="tab-icon" />
                <span>Statistics</span>
              </button>
              <button
                v-if="isAdmin"
                @click="activeTab = 'users'"
                class="xp-tab"
                :class="{ 'active': activeTab === 'users' }"
              >
                <UsersIcon class="tab-icon" />
                <span>Users</span>
              </button>
              <button
                v-if="isAdmin"
                @click="activeTab = 'settings'"
                class="xp-tab"
                :class="{ 'active': activeTab === 'settings' }"
              >
                <CogIcon class="tab-icon" />
                <span>Settings</span>
              </button>
            </div>

            <!-- Tab Content Area -->
            <div class="xp-tab-content">
              <!-- Uploads Tab -->
              <div v-if="activeTab === 'uploads'">
                <UploadsTab
                  :uploads="uploads"
                  :is-loading="isLoading"
                  :show-expiration-dropdown="showExpirationDropdown"
                  @copy-to-clipboard="copyToClipboard"
                  @toggle-availability="toggleAvailability"
                  @change-expiration="changeExpiration"
                  @toggle-expiration-dropdown="toggleExpirationDropdown"
                  @delete-upload="handleDeleteUpload"
                />
              </div>

              <!-- Reverse Share Tab -->
              <div v-if="activeTab === 'reverse'">
                <ReverseShareTab
                  :reverse-tokens="reverseTokens"
                  :is-loading="isLoading"
                  @create-reverse-token="createReverseToken"
                  @delete-reverse-token="deleteReverseToken"
                  @copy-to-clipboard="copyToClipboard"
                />
              </div>

              <!-- Statistics Tab -->
              <div v-if="activeTab === 'statistics' && isAdmin">
                <StatisticsTab
                  :admin-stats="adminStats"
                  :is-loading="isLoadingAdminSettings"
                />
              </div>

              <!-- Users Tab -->
              <div v-if="activeTab === 'users' && isAdmin">
                <UsersTab
                  :admin-users="adminUsers"
                  :is-loading="isLoadingAdminSettings"
                  @toggle-user-block="toggleUserBlock"
                />
              </div>

              <!-- Settings Tab -->
              <div v-if="activeTab === 'settings' && isAdmin">
                <SettingsTab
                  :is-admin="isAdmin"
                  @settings-updated="fetchAdminData"
                />
              </div>
            </div> <!-- Close xp-tab-content -->
          </div> <!-- Close xp-tab-control -->
        </div> <!-- Close xp-window-content -->
      </div> <!-- Close xp-window -->

    <!-- XP Window - Success Toast -->
    <div
      v-if="showCopySuccess"
      class="xp-toast"
    >
      <div class="toast-content">
        <ClipboardDocumentIcon class="toast-icon" />
        {{ successMessage }}
      </div>
    </div>

    <!-- Avatar Upload Dialog - XP Style -->
  <div v-if="showAvatarDialog" 
       class="xp-modal-overlay"
       @click="closeAvatarDialog">
    <div class="xp-dialog">
      <!-- Title Bar -->
      <div class="xp-title-bar">
        <div class="title-bar-left">
          <svg class="title-icon" width="16" height="16" viewBox="0 0 24 24" fill="none">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 9a2 2 0 012-2h.93a2 2 0 001.664-.89l.812-1.22A2 2 0 0110.07 4h3.86a2 2 0 011.664.89l.812 1.22A2 2 0 0018.07 7H19a2 2 0 012 2v9a2 2 0 01-2 2H5a2 2 0 01-2-2V9z"></path>
          </svg>
          <span class="title-text">Choose Avatar</span>
        </div>
        <div class="title-bar-buttons">
          <button class="title-btn close" @click="closeAvatarDialog">
            <svg width="10" height="10" viewBox="0 0 10 10">
              <path d="M1 1L9 9M9 1L1 9" stroke="currentColor" stroke-width="2"/>
            </svg>
          </button>
        </div>
      </div>

      <!-- Dialog content -->
      <div class="xp-dialog-content" @click.stop>
        <div class="dialog-message">
          <h2>Choose your avatar</h2>
          <p>Select a PNG, JPG or GIF image</p>
        </div>

        <div class="dialog-buttons">
          <button @click="openAvatarUpload" class="xp-button primary">
            Browse Files
          </button>
          <button @click="closeAvatarDialog" class="xp-button">
            Cancel
          </button>
        </div>
      </div>
    </div>
  </div>

  <!-- Delete Confirmation Dialog - XP Style -->
  <div v-if="showDeleteDialog" 
       class="xp-modal-overlay"
       @click="cancelDelete">
    <div class="xp-dialog small">
      <!-- Title Bar -->
      <div class="xp-title-bar">
        <div class="title-bar-left">
          <svg class="title-icon" width="16" height="16" viewBox="0 0 16 16" fill="red">
            <path d="M8 1L1 15H15L8 1Z"/>
            <text x="8" y="12" text-anchor="middle" fill="white" font-size="10" font-weight="bold">!</text>
          </svg>
          <span class="title-text">Confirm Delete</span>
        </div>
        <div class="title-bar-buttons">
          <button class="title-btn close" @click="cancelDelete">
            <svg width="10" height="10" viewBox="0 0 10 10">
              <path d="M1 1L9 9M9 1L1 9" stroke="currentColor" stroke-width="2"/>
            </svg>
          </button>
        </div>
      </div>

      <!-- Dialog content -->
      <div class="xp-dialog-content" @click.stop>
        <div class="dialog-message">
          <h2>Delete Upload?</h2>
          <p>Are you sure you want to delete this upload? This action cannot be undone.</p>
        </div>

        <div class="dialog-buttons">
          <button @click="confirmDelete" class="xp-button danger">
            Yes, Delete
          </button>
          <button @click="cancelDelete" class="xp-button">
            Cancel
          </button>
        </div>
      </div>
    </div>
  </div>
  
  </div> <!-- Close xp-account-container -->
  </div> <!-- Close xp-account-page -->
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch, onBeforeUnmount } from 'vue'
import { useAuth } from '../../composables/useAuth'
import { useTheme } from '../../composables/useTheme'
import { getApiUrl, getAssetUrl } from '../../utils/apiUtils'
import UploadsTab from './components/UploadsTab.vue'
import ReverseShareTab from './components/ReverseShareTab.vue'
import StatisticsTab from './components/StatisticsTab.vue'
import UsersTab from './components/UsersTab.vue'
import SettingsTab from './components/SettingsTab.vue'
import {
  UserIcon,
  CloudArrowUpIcon,
  ArchiveBoxIcon,
  ClockIcon,
  FolderIcon,
  CogIcon,
  ClipboardDocumentIcon,
  CalendarIcon,
  ShieldCheckIcon,
  ShareIcon,
  ChartBarIcon,
  UsersIcon
} from '@heroicons/vue/24/outline'


const { user, uploads, isAuthenticated, isAdmin, isLoading, fetchCurrentUser, fetchUploads, deleteUpload } = useAuth()
const { isDark } = useTheme()

// Enhanced delete function with proper error handling
const handleDeleteUpload = async (uploadId: string) => {
  uploadToDelete.value = uploadId
  showDeleteDialog.value = true
}

const confirmDelete = async () => {
  if (!uploadToDelete.value) return

  try {
    console.log('Attempting to delete upload:', uploadToDelete.value)
    const result = await deleteUpload(uploadToDelete.value)
    
    if (result.success) {
      console.log('Upload deleted successfully')
      // Show success message
      successMessage.value = 'Upload deleted successfully!'
      showCopySuccess.value = true
      setTimeout(() => {
        showCopySuccess.value = false
        successMessage.value = 'Link copied to clipboard!'
      }, 2000)
    } else {
      console.error('Failed to delete upload:', result.message)
      alert(result.message || 'Failed to delete upload')
    }
  } catch (error) {
    console.error('Error deleting upload:', error)
    alert('Failed to delete upload. Please try again.')
  } finally {
    showDeleteDialog.value = false
    uploadToDelete.value = null
  }
}

const cancelDelete = () => {
  showDeleteDialog.value = false
  uploadToDelete.value = null
}

const activeTab = ref('uploads')
const showCreateToken = ref(false)
const reverseTokens = ref<ReverseToken[]>([])
const newToken = ref({
  name: '',
  max_uses: -1,
  expires_in: ''
})

// Success message for copy operations and other actions
const showCopySuccess = ref(false)
const successMessage = ref('Link copied to clipboard!')

// Delete confirmation dialog
const showDeleteDialog = ref(false)
const uploadToDelete = ref<string | null>(null)

// Admin functionality
const isLoadingAdminSettings = ref(true)
const adminStats = ref({
  totalUsers: 0,
  totalUploads: 0,
  totalStorage: 0
})
const quickSettings = ref({
  allowRegistration: true
})
const advancedSettings = ref({
  maxUploadSize: 100, // in MB
  maxValidity: '7days',
  theme: 'light',
  navbarTitle: 'PInGO Share'
})
const adminUsers = ref<any[]>([])
const avatarInput = ref<HTMLInputElement | null>(null)

// Expiration dropdown state
const showExpirationDropdown = ref<Record<string, boolean>>({})

interface ReverseToken {
  id: number
  token: string
  name: string
  used_count: number
  max_uses: number
  created_at: string
  expires_at: string | null
}

const activeUploads = computed(() => {
  return uploads.value.filter(upload => {
    if (!upload.expires_at) return true // Never expires
    return new Date(upload.expires_at) > new Date()
  })
})

const formatDate = (dateString: string | undefined) => {
  if (!dateString) return 'Unknown'
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

const formatFileSize = (bytes: number) => {
  if (bytes === 0) return '0 Bytes'
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

const formatTotalSize = () => {
  const totalBytes = uploads.value
    .filter(upload => !upload.is_deleted) // Only include non-deleted uploads
    .reduce((sum, upload) => sum + upload.total_size, 0)
  return formatFileSize(totalBytes)
}

const copyToClipboard = async (text: string) => {
  try {
    await navigator.clipboard.writeText(text)
    successMessage.value = 'Link copied to clipboard!'
    showCopySuccess.value = true
    setTimeout(() => {
      showCopySuccess.value = false
    }, 2000)
  } catch (err) {
    console.error('Failed to copy: ', err)
    // Fallback for older browsers
    const textArea = document.createElement('textarea')
    textArea.value = text
    document.body.appendChild(textArea)
    textArea.select()
    document.body.removeChild(textArea)
    alert('Failed to copy to clipboard. Please copy manually.')
  }
}

const toggleAvailability = async (uploadId: string, currentAvailability: boolean) => {
  try {
    const response = await fetch(getApiUrl(`/uploads/${uploadId}/availability`), {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token') || ''}`
      },
      credentials: 'include',
      body: JSON.stringify({ is_available: !currentAvailability })
    })

    if (response.ok) {
      // Refresh uploads to get updated data
      await fetchUploads()
    } else {
      throw new Error('Failed to update availability')
    }
  } catch (error) {
    console.error('Error toggling availability:', error)
    alert('Failed to update file availability')
  }
}

const fetchReverseTokens = async () => {
  try {
    const response = await fetch(getApiUrl('/reverse-tokens'), {
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token') || ''}`
      },
      credentials: 'include'
    })

    if (response.ok) {
      const data = await response.json()
      reverseTokens.value = data.tokens || []
    }
  } catch (error) {
    console.error('Error fetching reverse tokens:', error)
  }
}

interface ReverseTokenData {
  name: string;
  max_uses: number;
  expires_in: string;
}

const createReverseToken = async (tokenData: ReverseTokenData) => {
  try {
    // Add validation before sending
    if (!tokenData.name || tokenData.name.trim() === '') {
      alert('Please enter a token name')
      return
    }

    // Ensure max_uses is sent as a number
    const payload = {
      name: tokenData.name.trim(),
      max_uses: parseInt(tokenData.max_uses?.toString() ?? '-1'),
      expires_in: tokenData.expires_in
    }

    console.log('Creating token with data:', payload)
    
    const response = await fetch(getApiUrl('/reverse-tokens'), {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token') || ''}`
      },
      credentials: 'include',
      body: JSON.stringify(payload)
    })

    const responseData = await response.json()
    console.log('Response:', responseData)

    if (response.ok) {
      await fetchReverseTokens()
      showCreateToken.value = false
      resetNewToken()
    } else {
      throw new Error(responseData.error || 'Failed to create token')
    }
  } catch (error) {
    console.error('Error creating reverse token:', error)
    const errorMessage = error instanceof Error ? error.message : 'Unknown error'
    alert(`Failed to create reverse share token: ${errorMessage}`)
  }
}

const deleteReverseToken = async (tokenId: number) => {
  if (!confirm('Are you sure you want to delete this token?')) return

  try {
    const response = await fetch(getApiUrl(`/reverse-tokens/${tokenId}`), {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token') || ''}`
      },
      credentials: 'include'
    })

    if (response.ok) {
      await fetchReverseTokens()
    } else {
      throw new Error('Failed to delete token')
    }
  } catch (error) {
    console.error('Error deleting reverse token:', error)
    alert('Failed to delete reverse share token')
  }
}

const resetNewToken = () => {
  newToken.value = {
    name: '',
    max_uses: -1,
    expires_in: ''
  }
}

// Admin functions
const toggleUserBlock = async (userId: number, isBlocked: boolean) => {
  try {
    const response = await fetch(getApiUrl(`/admin/users/${userId}/block`), {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token') || ''}`
      },
      body: JSON.stringify({
        blocked: isBlocked
      })
    })

    if (response.ok) {
      // Update local state
      const userIndex = adminUsers.value.findIndex(u => u.id === userId)
      if (userIndex !== -1) {
        adminUsers.value[userIndex].isBlocked = isBlocked
      }
    } else {
      console.error('Failed to toggle user block status')
    }
  } catch (error) {
    console.error('Error toggling user block:', error)
  }
}

// Change upload expiration
const changeExpiration = async (uploadId: string, validity: string) => {
  try {
    const url = getApiUrl(`/uploads/${uploadId}/expiration`)
    const token = localStorage.getItem('auth_token')
    
    console.log('=== EXPIRATION UPDATE DEBUG ===')
    console.log('Update URL:', url)
    console.log('Upload ID:', uploadId)
    console.log('Validity:', validity)
    console.log('Has token:', !!token)
    
    const response = await fetch(url, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token || ''}`
      },
      body: JSON.stringify({
        validity
      })
    })

    console.log('Expiration response status:', response.status)

    if (response.ok) {
      console.log('Expiration update successful')
      // Refresh uploads to get updated expiration dates
      await fetchUploads()
      // Close dropdown
      showExpirationDropdown.value[uploadId] = false
    } else {
      const errorText = await response.text()
      console.error('Failed to update expiration - Status:', response.status)
      console.error('Failed to update expiration - Response:', errorText)
      alert('Failed to update expiration date')
    }
  } catch (error) {
    console.error('Error updating expiration:', error)
    const errorMessage = error instanceof Error ? error.message : 'Unknown error'
    alert(`Failed to update expiration date: ${errorMessage}`)
  }
}

// Toggle expiration dropdown
const toggleExpirationDropdown = (uploadId: string) => {
  // Close all other dropdowns
  Object.keys(showExpirationDropdown.value).forEach(key => {
    if (key !== uploadId) {
      showExpirationDropdown.value[key] = false
    }
  })
  
  // Toggle the current dropdown
  showExpirationDropdown.value[uploadId] = !showExpirationDropdown.value[uploadId]
}

const fetchAdminData = async () => {
  if (!isAdmin.value) return
  
  try {
    isLoadingAdminSettings.value = true
    
    // Fetch admin statistics
    const statsResponse = await fetch(getApiUrl('/admin/stats'), {
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token') || ''}`
      }
    })
    
    if (statsResponse.ok) {
      const stats = await statsResponse.json()
      adminStats.value = {
        totalUsers: stats.totalUsers,
        totalUploads: stats.totalUploads,
        totalStorage: stats.storageUsed
      }
    }

    // Fetch quick settings from main settings endpoint
    const settingsResponse = await fetch(getApiUrl('/settings'))
    if (settingsResponse.ok) {
      const settings = await settingsResponse.json()
      quickSettings.value = {
        allowRegistration: settings.allowRegistration
      }
      advancedSettings.value = {
        maxUploadSize: Math.round(settings.maxUploadSize / (1024 * 1024)), // Convert bytes to MB
        maxValidity: settings.maxValidity,
        theme: settings.theme,
        navbarTitle: settings.navbarTitle || 'PInGO Share'
      }
    }
    
    // Fetch admin users
    const usersResponse = await fetch(getApiUrl('/admin/users'), {
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token') || ''}`
      }
    })
    
    if (usersResponse.ok) {
      const users = await usersResponse.json()
      adminUsers.value = users
    }
    
  } catch (error) {
    console.error('Error fetching admin data:', error)
  } finally {
    isLoadingAdminSettings.value = false
  }
}

onMounted(async () => {
  if (isAuthenticated.value) {
    await fetchCurrentUser()
    await fetchUploads()
    await fetchReverseTokens()
    if (isAdmin.value) {
      await fetchAdminData()
    }
  }
})

onBeforeUnmount(() => {
  stopFileFormatAnimation()
})

// Watch for tab changes to load admin data when needed
watch(activeTab, async (newTab: string) => {
  if ((newTab === 'statistics' || newTab === 'users' || newTab === 'settings') && isAdmin.value && adminStats.value.totalUsers === 0) {
    await fetchAdminData()
  }
})

const showAvatarDialog = ref(false)
const currentFileFormat = ref(0)
const fileFormats = ['.PNG', '.JPG', '.JPEG', '.GIF']

// Avatar functions
const openAvatarUpload = () => {
  showAvatarDialog.value = true
  startFileFormatAnimation()
}

const closeAvatarDialog = () => {
  showAvatarDialog.value = false
  stopFileFormatAnimation()
}

// File format animation
let formatInterval: number | null = null

const startFileFormatAnimation = () => {
  currentFileFormat.value = 0
  formatInterval = setInterval(() => {
    currentFileFormat.value = (currentFileFormat.value + 1) % fileFormats.length
  }, 1200) // Change every 1.2 seconds
}

const stopFileFormatAnimation = () => {
  if (formatInterval) {
    clearInterval(formatInterval)
    formatInterval = null
  }
}

const handleAvatarError = () => {
  console.error('Failed to load avatar')
}

const handleAvatarUpload = async (event: Event) => {
  const target = event.target as HTMLInputElement
  if (!target.files || target.files.length === 0) return

  const file = target.files[0]
  
  // Validate file type - only PNG, JPG, JPEG, and GIF
  const allowedTypes = ['image/jpeg', 'image/jpg', 'image/png', 'image/gif']
  if (!allowedTypes.includes(file.type)) {
    alert('Please select a valid image file (PNG, JPG, JPEG, or GIF only)')
    return
  }

  // Validate file size (5MB max)
//   if (file.size > 5 * 1024 * 1024) {
//     alert('File size must be less than 5MB')
//     return
//   }

  const formData = new FormData()
  formData.append('avatar', file)

  try {
    const url = getApiUrl('/avatar')
    const token = localStorage.getItem('auth_token')
    
    console.log('=== AVATAR UPLOAD DEBUG ===')
    console.log('Upload URL:', url)
    console.log('Has token:', !!token)
    console.log('Token preview:', token ? token.substring(0, 20) + '...' : 'none')
    console.log('File type:', file.type)
    console.log('File size:', file.size)
    
    const response = await fetch(url, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${token || ''}`
      },
      body: formData
    })

    console.log('Response status:', response.status)
    console.log('Response headers:', Object.fromEntries(response.headers.entries()))

    if (response.ok) {
      const result = await response.json()
      console.log('Upload successful:', result)
      // Update user avatar
      if (user.value) {
        user.value.avatar = result.avatar
      }
      // Reset input
      target.value = ''
      alert('Avatar uploaded successfully!')
    } else {
      const errorText = await response.text()
      console.error('Upload failed - Status:', response.status)
      console.error('Upload failed - Response:', errorText)
      
      let errorMessage = 'Failed to upload avatar'
      try {
        const errorJson = JSON.parse(errorText)
        errorMessage = errorJson.error || errorMessage
      } catch (e) {
        errorMessage = errorText || errorMessage
      }
      
      alert(`Failed to save avatar: ${errorMessage}`)
    }
  } catch (error) {
    console.error('Error uploading avatar:', error)
    const errorMessage = error instanceof Error ? error.message : 'Unknown error'
    alert(`Upload failed. Please try again. Error: ${errorMessage}`)
  }
}

</script>

<style scoped>
/* ============================================================================
   WINDOWS XP ACCOUNT PAGE - RETRO STYLING
   Authentic Windows XP Luna theme for user account management
   ============================================================================ */

/* Main Page Container */
.xp-account-page {
  min-height: 100vh;
  background: linear-gradient(to bottom, 
    #5a8fd5 0%,
    #6b9ddb 10%,
    #a4d3ee 25%,
    #b8e2f5 35%,
    #96c9dc 50%,
    #8ab9d0 60%,
    #76a9d5 75%,
    #6698c7 85%,
    #5a8fd5 100%
  );
  background-attachment: fixed;
  padding: 40px 20px;
  padding-bottom: 60px; /* Space for taskbar */
  font-family: 'Tahoma', 'Segoe UI', sans-serif;
}

.xp-account-page.theme-dark {
  background: linear-gradient(to bottom, 
    #1a1a2e 0%,
    #16213e 50%,
    #0f3460 100%
  );
}

.xp-account-container {
  max-width: 1200px;
  margin: 0 auto;
}

/* ============================================================================
   XP WINDOW STYLING
   ============================================================================ */

.xp-window {
  background: #ece9d8;
  border: 3px solid;
  border-color: #0054e3 #003db3 #003db3 #0054e3;
  border-radius: 8px 8px 0 0;
  box-shadow: 
    0 0 0 1px #87ceeb inset,
    4px 4px 12px rgba(0, 0, 0, 0.4);
  margin-bottom: 20px;
  overflow: hidden;
}

.xp-window.user-profile-window {
  min-height: 600px;
}

/* Title Bar */
.xp-title-bar {
  height: 30px;
  background: linear-gradient(to bottom,
    #0054e3 0%,
    #0054e3 100%
  );
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 6px 0 8px;
  border-bottom: 1px solid #003db3;
}

.title-bar-left {
  display: flex;
  align-items: center;
  gap: 6px;
}

.title-icon {
  width: 16px;
  height: 16px;
  flex-shrink: 0;
}

.title-text {
  color: #ffffff;
  font-size: 11px;
  font-weight: bold;
  font-family: 'Tahoma', sans-serif;
  text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.5);
}

.title-bar-buttons {
  display: flex;
  gap: 2px;
}

.title-btn {
  width: 21px;
  height: 21px;
  border: 1px solid;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: none;
}

.title-btn.minimize {
  background: linear-gradient(to bottom, #3d95ff 0%, #0054e3 100%);
  border-color: #87ceeb #003db3 #003db3 #87ceeb;
}

.title-btn.minimize span {
  width: 8px;
  height: 2px;
  background: #000;
  display: block;
}

.title-btn.maximize {
  background: linear-gradient(to bottom, #3d95ff 0%, #0054e3 100%);
  border-color: #87ceeb #003db3 #003db3 #87ceeb;
  color: #000;
}

.title-btn.close {
  background: linear-gradient(to bottom, #ff6b6b 0%, #ee5050 100%);
  border-color: #ff9999 #cc3333 #cc3333 #ff9999;
  color: #ffffff;
}

.title-btn:hover {
  filter: brightness(1.1);
}

.title-btn:active {
  filter: brightness(0.9);
}

/* Window Content */
.xp-window-content {
  padding: 16px;
  background: #ece9d8;
}

.theme-dark .xp-window {
  background: #2d2d2d;
  border-color: #1a4d8f #0d2a5c #0d2a5c #1a4d8f;
}

.theme-dark .xp-window-content {
  background: #2d2d2d;
}

/* ============================================================================
   USER HEADER SECTION
   ============================================================================ */

.xp-user-header {
  background: #ffffff;
  border: 2px solid;
  border-color: #7f9db9 #ffffff #ffffff #7f9db9;
  padding: 16px;
  margin-bottom: 16px;
}

.user-avatar-section {
  display: flex;
  align-items: center;
  gap: 20px;
}

.avatar-frame {
  position: relative;
  flex-shrink: 0;
}

.avatar-frame img,
.avatar-frame svg {
  width: 96px;
  height: 96px;
  border: 3px solid;
  border-color: #003db3 #87ceeb #87ceeb #003db3;
  box-shadow: inset 0 0 0 1px rgba(0, 0, 0, 0.1);
}

.avatar-placeholder {
  width: 96px;
  height: 96px;
  background: linear-gradient(135deg, #4a90e2 0%, #357abd 100%);
  color: white;
  padding: 20px;
}

.avatar-change-btn {
  margin-top: 8px;
  font-size: 10px;
  padding: 4px 12px !important;
}

.user-info {
  flex: 1;
}

.username {
  font-size: 18px;
  font-weight: bold;
  color: #003db3;
  margin: 0 0 8px 0;
  font-family: 'Tahoma', sans-serif;
}

.user-meta {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 11px;
  color: #000000;
}

.meta-icon {
  width: 14px;
  height: 14px;
}

.admin-badge {
  color: #c00000;
  font-weight: bold;
}

.theme-dark .xp-user-header {
  background: #1a1a1a;
  border-color: #444 #666 #666 #444;
}

.theme-dark .username {
  color: #4a9eff;
}

.theme-dark .meta-item {
  color: #cccccc;
}

/* ============================================================================
   GROUP BOX (STATS)
   ============================================================================ */

.xp-group-box {
  border: 2px solid;
  border-color: #ffffff #7f9db9 #7f9db9 #ffffff;
  padding: 16px;
  margin-bottom: 16px;
  position: relative;
  background: #f0f0f0;
}

.group-box-title {
  position: absolute;
  top: -10px;
  left: 12px;
  background: #ece9d8;
  padding: 0 6px;
  font-size: 11px;
  font-weight: bold;
  color: #000000;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 16px;
  margin-top: 8px;
}

.stat-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px;
  background: #ffffff;
  border: 1px solid;
  border-color: #7f9db9 #ffffff #ffffff #7f9db9;
}

.stat-icon {
  width: 48px;
  height: 48px;
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.stat-icon.uploads {
  background: linear-gradient(135deg, #4a90e2 0%, #357abd 100%);
}

.stat-icon.size {
  background: linear-gradient(135deg, #16a34a 0%, #15803d 100%);
}

.stat-icon.active {
  background: linear-gradient(135deg, #7c3aed 0%, #6d28d9 100%);
}

.stat-icon .icon {
  width: 24px;
  height: 24px;
  color: #ffffff;
}

.stat-info {
  flex: 1;
}

.stat-label {
  font-size: 10px;
  color: #666666;
  margin-bottom: 4px;
}

.stat-value {
  font-size: 20px;
  font-weight: bold;
  color: #000000;
}

.theme-dark .xp-group-box {
  background: #2d2d2d;
  border-color: #444 #666 #666 #444;
}

.theme-dark .group-box-title {
  background: #2d2d2d;
  color: #ffffff;
}

.theme-dark .stat-item {
  background: #1a1a1a;
  border-color: #444 #666 #666 #444;
}

.theme-dark .stat-label {
  color: #999999;
}

.theme-dark .stat-value {
  color: #ffffff;
}

/* ============================================================================
   XP TAB CONTROL
   ============================================================================ */

.xp-tab-control {
  border: 2px solid;
  border-color: #ffffff #7f9db9 #7f9db9 #ffffff;
  background: #ece9d8;
}

.xp-tab-headers {
  display: flex;
  gap: 2px;
  padding: 4px 4px 0 4px;
  background: #ece9d8;
  overflow-x: auto;
}

.xp-tab {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 6px 16px 8px;
  background: #d4d0c8;
  border: 2px solid;
  border-color: #ffffff #7f9db9 transparent #ffffff;
  border-radius: 4px 4px 0 0;
  font-size: 11px;
  font-family: 'Tahoma', sans-serif;
  color: #000000;
  cursor: pointer;
  transition: none;
  white-space: nowrap;
}

.xp-tab:hover {
  background: #e0ddd7;
}

.xp-tab.active {
  background: #ffffff;
  border-bottom-color: #ffffff;
  margin-bottom: -2px;
  padding-bottom: 10px;
  font-weight: bold;
}

.tab-icon {
  width: 16px;
  height: 16px;
  flex-shrink: 0;
}

.xp-tab-content {
  background: #ffffff;
  border-top: 2px solid #7f9db9;
  padding: 16px;
  min-height: 400px;
}

.theme-dark .xp-tab-control {
  background: #2d2d2d;
  border-color: #444 #666 #666 #444;
}

.xp-tab-headers {
  background: #2d2d2d;
}

.theme-dark .xp-tab {
  background: #1a1a1a;
  border-color: #444 #000 transparent #444;
  color: #cccccc;
}

.theme-dark .xp-tab:hover {
  background: #333333;
}

.theme-dark .xp-tab.active {
  background: #2d2d2d;
  border-bottom-color: #2d2d2d;
  color: #ffffff;
}

.theme-dark .xp-tab-content {
  background: #2d2d2d;
  border-top-color: #444;
}

/* ============================================================================
   XP BUTTONS
   ============================================================================ */

.xp-button {
  padding: 6px 16px;
  background: linear-gradient(to bottom,
    #ffffff 0%,
    #ece9d8 100%
  );
  border: 2px solid;
  border-color: #ffffff #7f9db9 #7f9db9 #ffffff;
  border-radius: 3px;
  font-size: 11px;
  font-family: 'Tahoma', sans-serif;
  font-weight: normal;
  color: #000000;
  cursor: pointer;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
  transition: none;
}

.xp-button:hover {
  background: linear-gradient(to bottom,
    #ffffef 0%,
    #f5f2e8 100%
  );
  border-color: #f5f2e8 #6f8dad #6f8dad #f5f2e8;
}

.xp-button:active {
  background: linear-gradient(to bottom,
    #d4d0c8 0%,
    #ece9d8 100%
  );
  border-color: #7f9db9 #ffffff #ffffff #7f9db9;
  box-shadow: inset 1px 1px 2px rgba(0, 0, 0, 0.2);
}

.xp-button.primary {
  background: linear-gradient(to bottom,
    #3d95ff 0%,
    #0054e3 100%
  );
  border-color: #87ceeb #003db3 #003db3 #87ceeb;
  color: #ffffff;
  font-weight: bold;
}

.xp-button.primary:hover {
  background: linear-gradient(to bottom,
    #5aa5ff 0%,
    #1a64f3 100%
  );
}

.xp-button.danger {
  background: linear-gradient(to bottom,
    #ff6b6b 0%,
    #ee5050 100%
  );
  border-color: #ff9999 #cc3333 #cc3333 #ff9999;
  color: #ffffff;
  font-weight: bold;
}

.theme-dark .xp-button {
  background: linear-gradient(to bottom, #3d3d3d 0%, #2d2d2d 100%);
  border-color: #555 #111 #111 #555;
  color: #cccccc;
}

.theme-dark .xp-button:hover {
  background: linear-gradient(to bottom, #4d4d4d 0%, #3d3d3d 100%);
}

/* ============================================================================
   TOAST NOTIFICATIONS
   ============================================================================ */

.xp-toast {
  position: fixed;
  bottom: 60px;
  right: 20px;
  z-index: 10000;
  animation: toast-slide-in 0.3s ease-out;
}

.toast-content {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 12px 20px;
  background: linear-gradient(to bottom,
    #3d95ff 0%,
    #0054e3 100%
  );
  border: 2px solid;
  border-color: #87ceeb #003db3 #003db3 #87ceeb;
  border-radius: 4px;
  color: #ffffff;
  font-size: 11px;
  font-weight: bold;
  box-shadow: 4px 4px 8px rgba(0, 0, 0, 0.3);
}

.toast-icon {
  width: 20px;
  height: 20px;
}

@keyframes toast-slide-in {
  from {
    transform: translateX(400px);
    opacity: 0;
  }
  to {
    transform: translateX(0);
    opacity: 1;
  }
}

/* ============================================================================
   XP MODAL & DIALOGS
   ============================================================================ */

.xp-modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 10001;
  padding: 20px;
  animation: fade-in 0.2s ease-out;
}

.xp-dialog {
  background: #ece9d8;
  border: 3px solid;
  border-color: #0054e3 #003db3 #003db3 #0054e3;
  border-radius: 8px 8px 0 0;
  box-shadow: 
    0 0 0 1px #87ceeb inset,
    4px 4px 16px rgba(0, 0, 0, 0.5);
  width: 100%;
  max-width: 450px;
  animation: dialog-bounce 0.3s ease-out;
}

.xp-dialog.small {
  max-width: 350px;
}

.xp-dialog-content {
  padding: 20px;
}

.dialog-message h2 {
  font-size: 14px;
  font-weight: bold;
  color: #003db3;
  margin: 0 0 12px 0;
}

.dialog-message p {
  font-size: 11px;
  color: #000000;
  margin: 0 0 20px 0;
  line-height: 1.5;
}

.dialog-buttons {
  display: flex;
  gap: 8px;
  justify-content: flex-end;
}

.dialog-buttons .xp-button {
  min-width: 80px;
}

@keyframes fade-in {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

@keyframes dialog-bounce {
  0% {
    transform: scale(0.8);
    opacity: 0;
  }
  50% {
    transform: scale(1.05);
  }
  100% {
    transform: scale(1);
    opacity: 1;
  }
}

.theme-dark .xp-dialog {
  background: #2d2d2d;
}

.theme-dark .xp-dialog-content {
  background: #2d2d2d;
}

.theme-dark .dialog-message h2 {
  color: #4a9eff;
}

.theme-dark .dialog-message p {
  color: #cccccc;
}

/* ============================================================================
   RESPONSIVE DESIGN
   ============================================================================ */

@media (max-width: 768px) {
  .xp-account-page {
    padding: 20px 10px 60px;
  }

  .xp-tab-headers {
    overflow-x: auto;
    -webkit-overflow-scrolling: touch;
  }

  .xp-tab {
    font-size: 10px;
    padding: 5px 12px 7px;
  }

  .stats-grid {
    grid-template-columns: 1fr;
  }

  .user-avatar-section {
    flex-direction: column;
    text-align: center;
  }

  .dialog-buttons {
    flex-direction: column;
  }

  .dialog-buttons .xp-button {
    width: 100%;
  }
}

/* Hide utility classes */
.hidden {
  display: none;
}

/* Dialog animations */
@keyframes dialog-backdrop {
  from { opacity: 0; }
  to { opacity: 1; }
}

@keyframes dialog-content {
  from { 
    opacity: 0; 
    transform: scale(0.95) translateY(-20px); 
  }
  to { 
    opacity: 1; 
    transform: scale(1) translateY(0); 
  }
}

@keyframes avatar-pulse {
  0%, 100% { 
    transform: scale(1); 
    box-shadow: 0 0 0 0 rgba(59, 130, 246, 0.4);
  }
  50% { 
    transform: scale(1.05); 
    box-shadow: 0 0 0 20px rgba(59, 130, 246, 0);
  }
}

@keyframes slide-up {
  from { 
    opacity: 0; 
    transform: translateY(20px); 
  }
  to { 
    opacity: 1; 
    transform: translateY(0); 
  }
}

@keyframes rainbow {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

@keyframes text-reveal-in {
  0% { 
    transform: translateY(100%); 
    opacity: 0; 
  }
  100% { 
    transform: translateY(0); 
    opacity: 1; 
  }
}

@keyframes text-reveal-out {
  0% { 
    transform: translateY(0); 
    opacity: 1; 
  }
  100% { 
    transform: translateY(-100%); 
    opacity: 0; 
  }
}

/* Animation classes */
.animate-dialog-backdrop {
  animation: dialog-backdrop 0.3s ease-out;
}

.animate-dialog-content {
  animation: dialog-content 0.4s ease-out;
}

.animate-avatar-pulse {
  animation: avatar-pulse 2s ease-in-out infinite;
}

.animate-slide-up {
  animation: slide-up 0.6s ease-out;
}

.animate-slide-up-delay {
  animation: slide-up 0.6s ease-out 0.2s both;
}

.animate-slide-up-delay-2 {
  animation: slide-up 0.6s ease-out 0.4s both;
}

.animate-slide-up-delay-3 {
  animation: slide-up 0.6s ease-out 0.6s both;
}

.animate-slide-up-delay-4 {
  animation: slide-up 0.6s ease-out 0.8s both;
}

.animate-rainbow {
  background-size: 400% 400%;
  animation: rainbow 3s ease-in-out infinite;
}

.animate-text-reveal-in {
  animation: text-reveal-in 0.7s ease-out forwards;
}

.animate-text-reveal-out {
  animation: text-reveal-out 0.7s ease-out forwards;
}

/* Apple-style backdrop blur */
.backdrop-blur-sm {
  backdrop-filter: blur(4px);
  -webkit-backdrop-filter: blur(4px);
}

/* Hide scrollbar for tab navigation */
.scrollbar-hide {
  -ms-overflow-style: none;
  scrollbar-width: none;
}

.scrollbar-hide::-webkit-scrollbar {
  display: none;
}
</style>
