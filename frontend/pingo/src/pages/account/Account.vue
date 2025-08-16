<template>
  <div class="min-h-screen transition-colors duration-300" 
       :style="{ backgroundColor: isDark ? '#000000' : '#ffffff' }">
    
    <!-- WebGL Background -->
    <div class="fixed inset-0 pointer-events-none overflow-hidden z-0">
      <WebGLBackground 
        :hue-shift="isDark ? 240 : 200"
        :noise-intensity="isDark ? 0.02 : 0.008"
        :scanline-intensity="isDark ? 0.08 : 0.03"
        :speed="0.15"
        :scanline-frequency="isDark ? 0.3 : 0.15"
        :warp-amount="0.06"
        :resolution-scale="0.8"
        class="opacity-30"
      />
    </div>

    <!-- Main Content -->
    <main class="relative z-10 min-h-screen pt-20">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        
        <!-- Hero Header -->
        <div class="text-center mb-16 animate-fade-in">
          <h1 class="text-5xl sm:text-6xl lg:text-7xl font-bold tracking-tight mb-6">
            <span :style="{ color: isDark ? '#ffffff' : '#000000' }">Your</span>
            <br />
            <span class="text-transparent bg-gradient-to-r from-blue-600 via-purple-600 to-pink-600 bg-clip-text">
              Account
            </span>
          </h1>
          <p class="text-xl sm:text-2xl max-w-2xl mx-auto leading-relaxed"
             :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">
            Manage your files, settings, and account preferences
          </p>
        </div>

        <!-- Profile Section -->
        <div class="max-w-4xl mx-auto mb-12 animate-fade-in-delay">
          <div class="relative">
            <div class="absolute inset-0 bg-gradient-to-r from-blue-600/20 to-purple-600/20 rounded-3xl blur-3xl"></div>
            <div class="relative backdrop-blur-xl border rounded-3xl p-8 sm:p-12"
                 :style="{
                   backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
                   borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)'
                 }">
              
              <div class="flex flex-col sm:flex-row items-center space-y-6 sm:space-y-0 sm:space-x-8">
                <!-- Avatar Section -->
                <div class="relative group">
                  <div class="absolute inset-0 bg-gradient-to-r from-blue-500 to-purple-500 rounded-3xl blur-lg opacity-75 group-hover:opacity-100 transition-opacity duration-300"></div>
                  <div class="relative w-32 h-32 rounded-3xl overflow-hidden border-2 bg-white/10 backdrop-blur-sm"
                       :style="{ borderColor: isDark ? 'rgba(255,255,255,0.2)' : 'rgba(0,0,0,0.2)' }">
                    <img v-if="user?.avatar" 
                         :src="`http://localhost:8080${user.avatar}`" 
                         :alt="user.username"
                         class="w-full h-full object-cover transition-transform duration-300 group-hover:scale-110" 
                         @error="handleAvatarError" />
                    <div v-else class="w-full h-full bg-gradient-to-br from-blue-500 via-purple-500 to-pink-500 flex items-center justify-center">
                      <svg class="w-16 h-16 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
                      </svg>
                    </div>
                  </div>
                  
                  <!-- Edit Avatar Button -->
                  <button 
                    @click="openAvatarUpload"
                    class="absolute -bottom-2 -right-2 w-10 h-10 bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700 text-white rounded-full flex items-center justify-center shadow-lg transition-all duration-200 transform hover:scale-110"
                    title="Change avatar"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 9a2 2 0 012-2h.93a2 2 0 001.664-.89l.812-1.22A2 2 0 0110.07 4h3.86a2 2 0 011.664.89l.812 1.22A2 2 0 0018.07 7H19a2 2 0 012 2v9a2 2 0 01-2 2H5a2 2 0 01-2-2V9z"></path>
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 13a3 3 0 11-6 0 3 3 0 016 0z"></path>
                    </svg>
                  </button>
                  <input 
                    ref="avatarInput" 
                    type="file" 
                    accept=".png,.jpg,.jpeg,.gif,image/png,image/jpeg,image/gif" 
                    @change="handleAvatarUpload" 
                    class="hidden"
                  />
                </div>

                <!-- User Info -->
                <div class="flex-1 text-center sm:text-left">
                  <h2 class="text-4xl font-bold mb-4"
                      :style="{ color: isDark ? '#ffffff' : '#000000' }">
                    {{ user?.username }}
                  </h2>
                  <p class="text-xl mb-6 opacity-80"
                     :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">
                    {{ user?.email }}
                  </p>
                  
                  <!-- User Badges -->
                  <div class="flex flex-wrap justify-center sm:justify-start gap-3 mb-6">
                    <div class="px-4 py-2 rounded-xl backdrop-blur-sm border"
                         :style="{
                           backgroundColor: isDark ? 'rgba(59, 130, 246, 0.2)' : 'rgba(59, 130, 246, 0.1)',
                           borderColor: 'rgba(59, 130, 246, 0.3)',
                           color: '#3b82f6'
                         }">
                      <div class="flex items-center space-x-2">
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                        </svg>
                        <span class="text-sm font-medium">Member since {{ formatDate(user?.created_at) }}</span>
                      </div>
                    </div>
                    
                    <div v-if="isAdmin" class="px-4 py-2 rounded-xl backdrop-blur-sm border"
                         :style="{
                           backgroundColor: isDark ? 'rgba(147, 51, 234, 0.2)' : 'rgba(147, 51, 234, 0.1)',
                           borderColor: 'rgba(147, 51, 234, 0.3)',
                           color: '#9333ea'
                         }">
                      <div class="flex items-center space-x-2">
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                        </svg>
                        <span class="text-sm font-medium">Administrator</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Stats Cards -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-12 animate-fade-in-delay">
          <div class="relative group">
            <div class="absolute inset-0 bg-gradient-to-r from-blue-600/20 to-cyan-600/20 rounded-2xl blur-lg opacity-75 group-hover:opacity-100 transition-opacity duration-300"></div>
            <div class="relative backdrop-blur-xl border rounded-2xl p-6"
                 :style="{
                   backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
                   borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)'
                 }">
              <div class="flex items-center space-x-4">
                <div class="w-12 h-12 rounded-xl bg-gradient-to-r from-blue-600 to-cyan-600 flex items-center justify-center">
                  <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M9 19l3 3m0 0l3-3m-3 3V10"></path>
                  </svg>
                </div>
                <div>
                  <p class="text-sm font-medium opacity-70"
                     :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">
                    Total Uploads
                  </p>
                  <p class="text-2xl font-bold"
                     :style="{ color: isDark ? '#ffffff' : '#000000' }">
                    {{ uploads.length }}
                  </p>
                </div>
              </div>
            </div>
          </div>

          <div class="relative group">
            <div class="absolute inset-0 bg-gradient-to-r from-purple-600/20 to-pink-600/20 rounded-2xl blur-lg opacity-75 group-hover:opacity-100 transition-opacity duration-300"></div>
            <div class="relative backdrop-blur-xl border rounded-2xl p-6"
                 :style="{
                   backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
                   borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)'
                 }">
              <div class="flex items-center space-x-4">
                <div class="w-12 h-12 rounded-xl bg-gradient-to-r from-purple-600 to-pink-600 flex items-center justify-center">
                  <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
                  </svg>
                </div>
                <div>
                  <p class="text-sm font-medium opacity-70"
                     :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">
                    Total Size
                  </p>
                  <p class="text-2xl font-bold"
                     :style="{ color: isDark ? '#ffffff' : '#000000' }">
                    {{ formatTotalSize() }}
                  </p>
                </div>
              </div>
            </div>
          </div>

          <div class="relative group">
            <div class="absolute inset-0 bg-gradient-to-r from-green-600/20 to-emerald-600/20 rounded-2xl blur-lg opacity-75 group-hover:opacity-100 transition-opacity duration-300"></div>
            <div class="relative backdrop-blur-xl border rounded-2xl p-6"
                 :style="{
                   backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
                   borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)'
                 }">
              <div class="flex items-center space-x-4">
                <div class="w-12 h-12 rounded-xl bg-gradient-to-r from-green-600 to-emerald-600 flex items-center justify-center">
                  <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
                  </svg>
                </div>
                <div>
                  <p class="text-sm font-medium opacity-70"
                     :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">
                    Active Files
                  </p>
                  <p class="text-2xl font-bold"
                     :style="{ color: isDark ? '#ffffff' : '#000000' }">
                    {{ activeUploads.length }}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Navigation Tabs -->
        <div class="max-w-4xl mx-auto mb-8 animate-fade-in-delay-2">
          <div class="relative">
            <div class="absolute inset-0 bg-gradient-to-r from-blue-600/10 to-purple-600/10 rounded-2xl blur-xl"></div>
            <div class="relative backdrop-blur-xl border rounded-2xl p-2"
                 :style="{
                   backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
                   borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)'
                 }">
              <nav class="flex space-x-2 overflow-x-auto">
                <button
                  v-for="tab in tabs"
                  :key="tab.id"
                  @click="activeTab = tab.id"
                  class="flex items-center space-x-2 px-6 py-3 rounded-xl font-medium transition-all duration-200 whitespace-nowrap"
                  :class="[
                    activeTab === tab.id
                      ? 'bg-gradient-to-r from-blue-600 to-purple-600 text-white shadow-lg'
                      : 'hover:bg-white/10'
                  ]"
                  :style="activeTab !== tab.id ? { color: isDark ? '#a1a1aa' : '#71717a' } : {}"
                >
                  <component :is="tab.icon" class="w-5 h-5" />
                  <span>{{ tab.name }}</span>
                </button>
              </nav>
            </div>
          </div>
        </div>

        <!-- Tab Content -->
        <div class="max-w-6xl mx-auto animate-fade-in-delay-2">
          <div class="relative">
            <div class="absolute inset-0 bg-gradient-to-r from-blue-600/20 to-purple-600/20 rounded-3xl blur-3xl"></div>
            <div class="relative backdrop-blur-xl border rounded-3xl p-8"
                 :style="{
                   backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
                   borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)'
                 }">
              
              <!-- Uploads Tab -->
              <UploadsTab v-if="activeTab === 'uploads'" 
                          :uploads="uploads" 
                          :isLoading="isLoading"
                          :showExpirationDropdown="showExpirationDropdown"
                          @copyToClipboard="handleCopyToClipboard"
                          @toggleAvailability="handleToggleAvailability"
                          @changeExpiration="handleChangeExpiration"
                          @toggleExpirationDropdown="handleToggleExpirationDropdown"
                          @deleteUpload="handleDeleteUpload" />
              
              <!-- Settings Tab -->
              <SettingsTab v-else-if="activeTab === 'settings'" 
                           :isAdmin="isAdmin"
                           @settingsUpdated="fetchUser" />
              
              <!-- Reverse Share Tab -->
              <ReverseShareTab v-else-if="activeTab === 'reverse'" 
                               :reverseTokens="reverseTokens"
                               :isLoading="isLoading"
                               @copyToClipboard="handleCopyToClipboard"
                               @createReverseToken="handleCreateReverseToken"
                               @deleteReverseToken="handleDeleteReverseToken" />
              
              <!-- Statistics Tab -->
              <StatisticsTab v-else-if="activeTab === 'statistics'" 
                             :adminStats="adminStats"
                             :isLoading="isLoading" />
              
              <!-- Admin Tabs -->
              <UsersTab v-else-if="activeTab === 'users' && isAdmin" 
                        :isDark="isDark" />
              
              <AdminSettings v-else-if="activeTab === 'admin' && isAdmin" 
                             :isDark="isDark" />
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- Success/Error Messages -->
    <div v-if="message" 
         class="fixed top-24 right-4 z-50 max-w-sm animate-fade-in">
      <div class="backdrop-blur-xl border rounded-2xl p-4 shadow-2xl"
           :style="{
             backgroundColor: message.type === 'success' 
               ? (isDark ? 'rgba(34, 197, 94, 0.2)' : 'rgba(34, 197, 94, 0.1)')
               : (isDark ? 'rgba(239, 68, 68, 0.2)' : 'rgba(239, 68, 68, 0.1)'),
             borderColor: message.type === 'success' ? 'rgba(34, 197, 94, 0.3)' : 'rgba(239, 68, 68, 0.3)',
             color: message.type === 'success' ? '#22c55e' : '#ef4444'
           }">
        <div class="flex items-center space-x-3">
          <div v-if="message.type === 'success'" class="w-5 h-5 rounded-full bg-green-500 flex items-center justify-center flex-shrink-0">
            <svg class="w-3 h-3 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
            </svg>
          </div>
          <div v-else class="w-5 h-5 rounded-full bg-red-500 flex items-center justify-center flex-shrink-0">
            <svg class="w-3 h-3 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </div>
          <p class="font-medium">{{ message.text }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useAuth } from '../../composables/useAuth'
import { useTheme } from '../../composables/useTheme'
import axios from 'axios'
import WebGLBackground from '../../components/WebGLBackground.vue'

// Import tab components
import UploadsTab from './components/UploadsTab.vue'
import SettingsTab from './components/SettingsTab.vue'
import ReverseShareTab from './components/ReverseShareTab.vue'
import StatisticsTab from './components/StatisticsTab.vue'
import UsersTab from './components/UsersTab.vue'
import AdminSettings from './components/AdminSettings.vue'

// Icons for tabs
import {
  CloudArrowUpIcon,
  CogIcon,
  ArrowPathIcon,
  ChartBarIcon,
  UsersIcon,
  ShieldCheckIcon
} from '@heroicons/vue/24/outline'

const { user, isAuthenticated, isAdmin } = useAuth()
const { isDark } = useTheme()

// State
const activeTab = ref('uploads')
const uploads = ref<any[]>([])
const message = ref<{ text: string; type: 'success' | 'error' } | null>(null)
const avatarInput = ref<HTMLInputElement | null>(null)
const isLoading = ref(false)
const showExpirationDropdown = ref<Record<string, boolean>>({})
const reverseTokens = ref<any[]>([])
const adminStats = ref({
  totalUsers: 0,
  totalUploads: 0,
  totalStorage: 0
})

// Computed
const activeUploads = computed(() => 
  uploads.value.filter(upload => upload.is_available !== false)
)

// Tab configuration
const tabs = computed(() => {
  const baseTabs = [
    { id: 'uploads', name: 'My Uploads', icon: CloudArrowUpIcon },
    { id: 'reverse', name: 'Reverse Share', icon: ArrowPathIcon },
    { id: 'statistics', name: 'Statistics', icon: ChartBarIcon },
    { id: 'settings', name: 'Settings', icon: CogIcon }
  ]
  
  if (isAdmin.value) {
    baseTabs.push(
      { id: 'users', name: 'Users', icon: UsersIcon },
      { id: 'admin', name: 'Admin', icon: ShieldCheckIcon }
    )
  }
  
  return baseTabs
})

// Methods
const fetchUser = async () => {
  // User data is already available through useAuth
}

const fetchUploads = async () => {
  try {
    const response = await axios.get('/uploads', {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('token')}`
      }
    })
    uploads.value = response.data
  } catch (error) {
    console.error('Error fetching uploads:', error)
    showMessage('Failed to fetch uploads', 'error')
  }
}

const formatDate = (dateString: string | undefined) => {
  if (!dateString) return 'Unknown'
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

const formatTotalSize = () => {
  const totalBytes = uploads.value.reduce((sum, upload) => sum + (upload.size || 0), 0)
  return formatFileSize(totalBytes)
}

const formatFileSize = (bytes: number) => {
  if (bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(1)) + ' ' + sizes[i]
}

const openAvatarUpload = () => {
  avatarInput.value?.click()
}

const handleAvatarUpload = async (event: Event) => {
  const target = event.target as HTMLInputElement
  const file = target.files?.[0]
  
  if (!file) return
  
  // Validate file type
  const allowedTypes = ['image/png', 'image/jpeg', 'image/jpg', 'image/gif']
  if (!allowedTypes.includes(file.type)) {
    showMessage('Please select a PNG, JPG, or GIF image', 'error')
    return
  }
  
  // Validate file size (max 5MB)
  if (file.size > 5 * 1024 * 1024) {
    showMessage('Image must be smaller than 5MB', 'error')
    return
  }
  
  try {
    const formData = new FormData()
    formData.append('avatar', file)
    
    await axios.post('/upload-avatar', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
        Authorization: `Bearer ${localStorage.getItem('token')}`
      }
    })
    
    showMessage('Avatar updated successfully!', 'success')
    // Trigger a re-fetch of user data
    window.location.reload()
  } catch (error) {
    console.error('Error uploading avatar:', error)
    showMessage('Failed to update avatar', 'error')
  }
}

const handleAvatarError = (event: Event) => {
  const target = event.target as HTMLImageElement
  target.style.display = 'none'
}

const showMessage = (text: string, type: 'success' | 'error') => {
  message.value = { text, type }
  setTimeout(() => {
    message.value = null
  }, 5000)
}

// Event handlers for child components
const handleCopyToClipboard = async (text: string) => {
  try {
    await navigator.clipboard.writeText(text)
    showMessage('Copied to clipboard!', 'success')
  } catch (error) {
    showMessage('Failed to copy to clipboard', 'error')
  }
}

const handleToggleAvailability = async (uploadId: string, currentAvailability: boolean) => {
  try {
    await axios.patch(`/uploads/${uploadId}/availability`, 
      { is_available: !currentAvailability },
      {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('token')}`
        }
      }
    )
    
    showMessage(`Upload ${!currentAvailability ? 'enabled' : 'disabled'} successfully`, 'success')
    fetchUploads()
  } catch (error) {
    console.error('Error toggling availability:', error)
    showMessage('Failed to update availability', 'error')
  }
}

const handleChangeExpiration = async (uploadId: string, validity: string) => {
  try {
    await axios.patch(`/uploads/${uploadId}/expiration`, 
      { validity },
      {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('token')}`
        }
      }
    )
    
    showMessage('Expiration updated successfully', 'success')
    fetchUploads()
  } catch (error) {
    console.error('Error updating expiration:', error)
    showMessage('Failed to update expiration', 'error')
  }
}

const handleToggleExpirationDropdown = (uploadId: string) => {
  showExpirationDropdown.value[uploadId] = !showExpirationDropdown.value[uploadId]
}

const handleDeleteUpload = async (uploadId: string) => {
  try {
    await axios.delete(`/uploads/${uploadId}`, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('token')}`
      }
    })
    
    showMessage('Upload deleted successfully', 'success')
    fetchUploads()
  } catch (error) {
    console.error('Error deleting upload:', error)
    showMessage('Failed to delete upload', 'error')
  }
}

const handleCreateReverseToken = async (tokenData: any) => {
  try {
    await axios.post('/reverse-tokens', tokenData, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('token')}`
      }
    })
    
    showMessage('Reverse token created successfully', 'success')
    fetchReverseTokens()
  } catch (error) {
    console.error('Error creating reverse token:', error)
    showMessage('Failed to create reverse token', 'error')
  }
}

const handleDeleteReverseToken = async (tokenId: number) => {
  try {
    await axios.delete(`/reverse-tokens/${tokenId}`, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('token')}`
      }
    })
    
    showMessage('Reverse token deleted successfully', 'success')
    fetchReverseTokens()
  } catch (error) {
    console.error('Error deleting reverse token:', error)
    showMessage('Failed to delete reverse token', 'error')
  }
}

const fetchReverseTokens = async () => {
  try {
    const response = await axios.get('/reverse-tokens', {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('token')}`
      }
    })
    reverseTokens.value = response.data
  } catch (error) {
    console.error('Error fetching reverse tokens:', error)
  }
}

const fetchAdminStats = async () => {
  if (!isAdmin.value) return
  
  try {
    const response = await axios.get('/admin/stats', {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('token')}`
      }
    })
    adminStats.value = response.data
  } catch (error) {
    console.error('Error fetching admin stats:', error)
  }
}

// Lifecycle
onMounted(() => {
  if (!isAuthenticated.value) {
    window.location.href = '/auth'
    return
  }
  
  fetchUploads()
  fetchReverseTokens()
  fetchAdminStats()
})
</script>

<style scoped>
@keyframes fade-in {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}

@keyframes fade-in-delay {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}

@keyframes fade-in-delay-2 {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}

.animate-fade-in {
  animation: fade-in 0.8s ease-out;
}

.animate-fade-in-delay {
  animation: fade-in-delay 0.8s ease-out 0.2s both;
}

.animate-fade-in-delay-2 {
  animation: fade-in-delay-2 0.8s ease-out 0.4s both;
}

/* Custom scrollbar for nav */
nav::-webkit-scrollbar {
  height: 4px;
}

nav::-webkit-scrollbar-track {
  background: transparent;
}

nav::-webkit-scrollbar-thumb {
  background: rgba(255, 255, 255, 0.2);
  border-radius: 2px;
}

nav::-webkit-scrollbar-thumb:hover {
  background: rgba(255, 255, 255, 0.3);
}
</style>
