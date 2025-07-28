<template>
  <div class="min-h-screen relative overflow-hidden" :style="{ '--button-from': buttonFromColor, '--button-to': buttonToColor }">
    <!-- Dynamic background -->
    <div class="absolute inset-0 z-0">
      <img 
        ref="backgroundImage"
        :src="`http://localhost:8080${backgroundPath}`"
        class="w-full h-full object-cover"
        alt="Background"
        @error="handleImageError"
        @load="onBackgroundImageLoad"
      />
    </div>

    <!-- Main content -->
    <div class="relative z-10 min-h-screen px-4 py-8">
      <!-- Logo and back navigation -->
      <div class="absolute top-4 left-4 z-20">
        <router-link to="/" class="flex items-center space-x-2 group">
          <img v-if="logoPath" :src="logoPath ? `http://localhost:8080${logoPath}` : ''" class="h-8 w-8" alt="Logo" @error="handleImageError" />
          <CloudArrowUpIcon v-else class="h-8 w-8 text-blue-600" />
          <span class="text-2xl font-extrabold text-gray-900 group-hover:text-blue-600 transition-colors duration-200">PinGO File Transfer</span>
        </router-link>
      </div>

      <!-- Netflix/Spotify Style Container -->
      <div class="flex items-start justify-center min-h-screen pt-20">
        <div class="w-full max-w-5xl mx-auto">
          <!-- Modern Header -->
          <div class="mb-12">
            <div class="flex items-end space-x-4 mb-6">
              <h1 class="text-3xl font-bold text-white tracking-tight">Settings</h1>
              <div class="w-1.5 h-1.5 bg-blue-500 rounded-full mb-2 animate-pulse"></div>
            </div>
            <p class="text-gray-300 text-base font-light max-w-2xl">
              Customize your workspace and personalize your experience
            </p>
            
            <!-- Modern Tab Navigation -->
            <div class="mt-8">
              <div class="bg-gray-800/50 backdrop-blur-sm rounded-xl p-1.5 inline-flex space-x-1 border border-gray-700/50">
                <button
                  v-for="tab in tabs"
                  :key="tab"
                  @click="activeTab = tab"
                  :class="[
                    'px-4 py-2 rounded-lg text-sm font-medium transition-all duration-300',
                    activeTab === tab 
                      ? 'bg-white text-gray-900 shadow-md' 
                      : 'text-gray-300 hover:text-white hover:bg-gray-700/50'
                  ]"
                >
                  {{ tab }}
                </button>
              </div>
            </div>
          </div>

          <!-- Tab Content -->
          <div class="pb-20">
            <!-- Branding & UI Tab -->
            <div v-if="activeTab === 'Branding & UI'" class="space-y-8">
              
              <!-- General Settings Section -->
              <div class="space-y-4">
                <div class="flex items-center space-x-2 mb-6">
                  <div class="w-0.5 h-4 bg-gradient-to-b from-blue-500 to-purple-600 rounded-full"></div>
                  <h2 class="text-lg font-semibold text-white">General</h2>
                </div>
                
                <div class="grid grid-cols-1 lg:grid-cols-2 gap-4">
                  
                  <!-- Theme Card -->
                  <div class="group bg-gradient-to-br from-gray-800/80 to-gray-900/80 backdrop-blur-sm rounded-xl p-5 border border-gray-700/50 hover:border-gray-600/50 transition-all duration-300 hover:shadow-xl hover:shadow-blue-500/5">
                    <div class="flex items-start justify-between mb-4">
                      <div>
                        <h3 class="text-white font-medium text-base mb-1">Theme</h3>
                        <p class="text-gray-400 text-xs leading-relaxed">Choose your preferred appearance mode</p>
                      </div>
                      <div class="w-2 h-2 bg-blue-500 rounded-full opacity-60 group-hover:opacity-100 transition-opacity"></div>
                    </div>
                    <select
                      v-model="settings.theme"
                      class="w-full bg-gray-700/50 border border-gray-600/50 rounded-lg px-3 py-2 text-white text-sm focus:outline-none focus:ring-1 focus:ring-blue-500/50 focus:border-blue-500/50 transition-all duration-200 backdrop-blur-sm"
                    >
                      <option value="light">☀️ Light Mode</option>
                      <option value="dark">🌙 Dark Mode</option>
                      <option value="system">⚡ System Default</option>
                    </select>
                  </div>

                  <!-- Upload Limit Card -->
                  <div class="group bg-gradient-to-br from-gray-800/80 to-gray-900/80 backdrop-blur-sm rounded-xl p-5 border border-gray-700/50 hover:border-gray-600/50 transition-all duration-300 hover:shadow-xl hover:shadow-purple-500/5">
                    <div class="flex items-start justify-between mb-4">
                      <div>
                        <h3 class="text-white font-medium text-base mb-1">Upload Limit</h3>
                        <p class="text-gray-400 text-xs leading-relaxed">Set maximum file size for uploads</p>
                      </div>
                      <div class="w-2 h-2 bg-purple-500 rounded-full opacity-60 group-hover:opacity-100 transition-opacity"></div>
                    </div>
                    <select
                      v-model="settings.maxUploadSize"
                      class="w-full bg-gray-700/50 border border-gray-600/50 rounded-lg px-3 py-2 text-white text-sm focus:outline-none focus:ring-1 focus:ring-purple-500/50 focus:border-purple-500/50 transition-all duration-200 backdrop-blur-sm"
                    >
                      <option value="1048576">1 MB</option>
                      <option value="10485760">10 MB</option>
                      <option value="52428800">50 MB</option>
                      <option value="104857600">100 MB</option>
                      <option value="1073741824">1 GB</option>
                    </select>
                  </div>
                </div>
              </div>

              <!-- Customization Section -->
              <div class="space-y-4">
                <div class="flex items-center space-x-2 mb-6">
                  <div class="w-0.5 h-4 bg-gradient-to-b from-purple-500 to-pink-600 rounded-full"></div>
                  <h2 class="text-lg font-semibold text-white">Customization</h2>
                </div>
                
                <div class="grid grid-cols-1 lg:grid-cols-2 gap-4">
                  
                  <!-- Logo Upload Card -->
                  <div class="group bg-gradient-to-br from-gray-800/80 to-gray-900/80 backdrop-blur-sm rounded-xl p-5 border border-gray-700/50 hover:border-gray-600/50 transition-all duration-300 hover:shadow-xl hover:shadow-green-500/5">
                    <div class="flex items-start justify-between mb-4">
                      <div>
                        <h3 class="text-white font-medium text-base mb-1">Brand Logo</h3>
                        <p class="text-gray-400 text-xs leading-relaxed">Upload your company or personal logo</p>
                      </div>
                      <div class="w-2 h-2 bg-green-500 rounded-full opacity-60 group-hover:opacity-100 transition-opacity"></div>
                    </div>
                    <div class="relative">
                      <input
                        type="file"
                        accept="image/*"
                        @change="handleLogoUpload"
                        class="w-full text-xs text-gray-300 file:mr-3 file:py-2 file:px-4 file:rounded-lg file:border-0 file:bg-gradient-to-r file:from-green-600 file:to-emerald-600 file:text-white hover:file:from-green-500 hover:file:to-emerald-500 file:font-medium file:transition-all file:duration-200 file:cursor-pointer cursor-pointer file:shadow-md hover:file:shadow-lg file:text-xs"
                      />
                    </div>
                  </div>

                  <!-- Background Upload Card -->
                  <div class="group bg-gradient-to-br from-gray-800/80 to-gray-900/80 backdrop-blur-sm rounded-xl p-5 border border-gray-700/50 hover:border-gray-600/50 transition-all duration-300 hover:shadow-xl hover:shadow-orange-500/5">
                    <div class="flex items-start justify-between mb-4">
                      <div>
                        <h3 class="text-white font-medium text-base mb-1">Background Image</h3>
                        <p class="text-gray-400 text-xs leading-relaxed">Set a custom background for your workspace</p>
                      </div>
                      <div class="w-2 h-2 bg-orange-500 rounded-full opacity-60 group-hover:opacity-100 transition-opacity"></div>
                    </div>
                    <div class="relative">
                      <input
                        type="file"
                        accept="image/*"
                        @change="handleBackgroundUpload"
                        class="w-full text-xs text-gray-300 file:mr-3 file:py-2 file:px-4 file:rounded-lg file:border-0 file:bg-gradient-to-r file:from-orange-600 file:to-red-600 file:text-white hover:file:from-orange-500 hover:file:to-red-500 file:font-medium file:transition-all file:duration-200 file:cursor-pointer cursor-pointer file:shadow-md hover:file:shadow-lg file:text-xs"
                      />
                    </div>
                  </div>
                </div>
              </div>

              <!-- Visual Effects Section -->
              <div class="space-y-4">
                <div class="flex items-center space-x-2 mb-6">
                  <div class="w-0.5 h-4 bg-gradient-to-b from-cyan-500 to-blue-600 rounded-full"></div>
                  <h2 class="text-lg font-semibold text-white">Visual Effects</h2>
                </div>
                
                <div class="grid grid-cols-1 lg:grid-cols-2 gap-4">
                  
                  <!-- Transparency Card -->
                  <div class="group bg-gradient-to-br from-gray-800/80 to-gray-900/80 backdrop-blur-sm rounded-xl p-5 border border-gray-700/50 hover:border-gray-600/50 transition-all duration-300 hover:shadow-xl hover:shadow-cyan-500/5">
                    <div class="flex items-start justify-between mb-5">
                      <div>
                        <h3 class="text-white font-medium text-base mb-1">Transparency</h3>
                        <p class="text-gray-400 text-xs leading-relaxed">Fine-tune background opacity levels</p>
                      </div>
                      <div class="w-2 h-2 bg-cyan-500 rounded-full opacity-60 group-hover:opacity-100 transition-opacity"></div>
                    </div>
                    <div class="space-y-4">
                      <input
                        type="range"
                        min="0"
                        max="100"
                        step="5"
                        v-model="settings.uploadBoxTransparency"
                        class="w-full compact-slider"
                      />
                      <div class="flex justify-between items-center text-xs">
                        <span class="text-gray-500 font-medium">Opaque</span>
                        <div class="bg-gradient-to-r from-cyan-500 to-blue-600 px-3 py-1 rounded-lg">
                          <span class="text-white font-semibold text-xs">{{ settings.uploadBoxTransparency }}%</span>
                        </div>
                        <span class="text-gray-500 font-medium">Transparent</span>
                      </div>
                    </div>
                  </div>

                  <!-- Blur Effect Card -->
                  <div class="group bg-gradient-to-br from-gray-800/80 to-gray-900/80 backdrop-blur-sm rounded-xl p-5 border border-gray-700/50 hover:border-gray-600/50 transition-all duration-300 hover:shadow-xl hover:shadow-indigo-500/5">
                    <div class="flex items-start justify-between mb-5">
                      <div>
                        <h3 class="text-white font-medium text-base mb-1">Blur Effect</h3>
                        <p class="text-gray-400 text-xs leading-relaxed">Adjust glass morphism effect intensity</p>
                      </div>
                      <div class="w-2 h-2 bg-indigo-500 rounded-full opacity-60 group-hover:opacity-100 transition-opacity"></div>
                    </div>
                    <div class="space-y-4">
                      <input
                        type="range"
                        min="0"
                        max="24"
                        step="1"
                        v-model="settings.blurIntensity"
                        class="w-full compact-slider"
                      />
                      <div class="flex justify-between items-center text-xs">
                        <span class="text-gray-500 font-medium">Sharp</span>
                        <div class="bg-gradient-to-r from-indigo-500 to-purple-600 px-3 py-1 rounded-lg">
                          <span class="text-white font-semibold text-xs">{{ getBlurLabel(settings.blurIntensity) }}</span>
                        </div>
                        <span class="text-gray-500 font-medium">Blurred</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <!-- Save Button -->
              <div class="flex justify-end pt-8">
                <button
                  @click="saveSettings"
                  class="group relative px-6 py-2.5 bg-gradient-to-r from-blue-600 via-purple-600 to-blue-700 text-white rounded-xl font-medium text-sm shadow-xl shadow-blue-500/20 hover:shadow-blue-500/30 transition-all duration-300 hover:scale-105 overflow-hidden"
                >
                  <div class="absolute inset-0 bg-gradient-to-r from-white/0 via-white/10 to-white/0 -translate-x-full group-hover:translate-x-full transition-transform duration-700"></div>
                  <span class="relative flex items-center space-x-2">
                    <span>Save Changes</span>
                    <div class="w-1.5 h-1.5 bg-white rounded-full opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
                  </span>
                </button>
              </div>
            </div>

            <!-- Users Tab -->
            <div v-if="activeTab === 'Users'" class="space-y-6">
              <div class="space-y-4">
                <div class="flex items-center space-x-2 mb-6">
                  <div class="w-0.5 h-4 bg-gradient-to-b from-emerald-500 to-teal-600 rounded-full"></div>
                  <h2 class="text-lg font-semibold text-white">User Management</h2>
                </div>
                
                <div class="bg-gradient-to-br from-gray-800/80 to-gray-900/80 backdrop-blur-sm rounded-xl p-5 border border-gray-700/50">
                  <div class="space-y-3">
                    <div v-for="user in users" :key="user.id" class="flex items-center justify-between p-4 bg-gray-700/50 rounded-lg border border-gray-600/30 hover:bg-gray-700/70 transition-all duration-200">
                      <span class="text-white font-medium text-sm">{{ user.name }}</span>
                      <button 
                        @click="removeUser(user.id)" 
                        class="px-4 py-2 bg-gradient-to-r from-red-600 to-red-700 hover:from-red-500 hover:to-red-600 text-white rounded-lg font-medium text-xs transition-all duration-200 shadow-md hover:shadow-lg hover:scale-105"
                      >
                        Remove
                      </button>
                    </div>
                  </div>
                  
                  <div class="mt-5">
                    <button
                      @click="inviteUser"
                      class="w-full px-4 py-2.5 bg-gradient-to-r from-emerald-600 to-teal-600 hover:from-emerald-500 hover:to-teal-500 text-white rounded-lg font-medium text-sm transition-all duration-200 shadow-md hover:shadow-lg hover:scale-105"
                    >
                      Invite New User
                    </button>
                  </div>
                </div>
              </div>
            </div>

            <!-- My Profile Tab -->
            <div v-if="activeTab === 'My Profile'" class="space-y-6">
              <div class="space-y-4">
                <div class="flex items-center space-x-2 mb-6">
                  <div class="w-0.5 h-4 bg-gradient-to-b from-pink-500 to-rose-600 rounded-full"></div>
                  <h2 class="text-lg font-semibold text-white">Profile Settings</h2>
                </div>
                
                <div class="bg-gradient-to-br from-gray-800/80 to-gray-900/80 backdrop-blur-sm rounded-xl p-5 border border-gray-700/50 space-y-5">
                  <div>
                    <label class="block text-white font-medium text-sm mb-2">Email Address</label>
                    <input
                      type="email"
                      v-model="profile.email"
                      class="w-full bg-gray-700/50 border border-gray-600/50 rounded-lg px-3 py-2.5 text-white text-sm focus:outline-none focus:ring-1 focus:ring-pink-500/50 focus:border-pink-500/50 transition-all duration-200 backdrop-blur-sm"
                      placeholder="your.email@example.com"
                    />
                  </div>
                  
                  <div>
                    <label class="block text-white font-medium text-sm mb-2">Password</label>
                    <input
                      type="password"
                      v-model="profile.password"
                      class="w-full bg-gray-700/50 border border-gray-600/50 rounded-lg px-3 py-2.5 text-white text-sm focus:outline-none focus:ring-1 focus:ring-pink-500/50 focus:border-pink-500/50 transition-all duration-200 backdrop-blur-sm"
                      placeholder="••••••••••••"
                    />
                  </div>
                  
                  <div class="pt-3">
                    <button
                      @click="saveProfile"
                      class="w-full px-4 py-2.5 bg-gradient-to-r from-pink-600 to-rose-600 hover:from-pink-500 hover:to-rose-500 text-white rounded-lg font-medium text-sm transition-all duration-200 shadow-md hover:shadow-lg hover:scale-105"
                    >
                      Save Profile Changes
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Success/Error Message -->
          <div v-if="message" class="fixed bottom-6 right-6 z-50">
            <div :class="[
              'px-5 py-3 rounded-xl font-medium text-sm shadow-xl transition-all duration-500 backdrop-blur-sm border transform',
              message.type === 'success' 
                ? 'bg-gradient-to-r from-emerald-600 to-green-600 text-white border-emerald-400/30 shadow-emerald-500/20' 
                : 'bg-gradient-to-r from-red-600 to-rose-600 text-white border-red-400/30 shadow-red-500/20'
            ]">
              <div class="flex items-center space-x-2">
                <div :class="[
                  'w-2 h-2 rounded-full animate-pulse',
                  message.type === 'success' ? 'bg-emerald-300' : 'bg-red-300'
                ]"></div>
                <span>{{ message.text }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import axios from 'axios'
import { CloudArrowUpIcon } from '@heroicons/vue/24/outline'

// Types
interface Settings {
  theme: 'light' | 'dark' | 'system'
  logo: File | null
  backgroundImage: File | null
  maxUploadSize: string | number
  uploadBoxTransparency: number
  blurIntensity: number
}

interface Profile {
  email: string
  password: string
}

interface User {
  id: string
  name: string
}

interface Message {
  text: string
  type: 'success' | 'error'
}

// State
const tabs = ['Branding & UI', 'Users', 'My Profile']
const activeTab = ref('Branding & UI')
const settings = reactive<Settings>({
  theme: 'light',
  logo: null,
  backgroundImage: null,
  maxUploadSize: '104857600',
  uploadBoxTransparency: 80,
  blurIntensity: 12,
})
const profile = reactive<Profile>({
  email: '',
  password: '',
})
const users = ref<User[]>([{ id: '1', name: 'User 1' }])
const message = ref<Message | null>(null)

// Design system variables
const buttonFromColor = ref<string>('#6366f1')
const buttonToColor = ref<string>('#8b5cf6')
const logoPath = ref<string | null>(null)
const backgroundPath = ref<string | null>('/api/backgrounds/default.jpg')
const backgroundImage = ref<HTMLImageElement | null>(null)

// Functions
const getBlurLabel = (intensity: number) => {
  if (intensity === 0) return 'None'
  if (intensity <= 2) return 'Very Light'
  if (intensity <= 6) return 'Light'
  if (intensity <= 10) return 'Medium'
  if (intensity <= 14) return 'Strong'
  if (intensity <= 18) return 'Heavy'
  if (intensity <= 22) return 'Intense'
  return 'Maximum'
}

const handleLogoUpload = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files && target.files[0]) {
    settings.logo = target.files[0]
  }
}

const handleBackgroundUpload = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files && target.files[0]) {
    settings.backgroundImage = target.files[0]
  }
}

const handleImageError = () => {
  console.error('Image failed to load:', backgroundPath.value || logoPath.value)
  if (!backgroundPath.value) backgroundPath.value = 'https://4kwallpapers.com/images/wallpapers/jungle-tree-dark-3840x2160-22695.jpg'
  if (!logoPath.value) logoPath.value = null
}

const onBackgroundImageLoad = () => {
  console.log('Background image loaded successfully')
}

const saveSettings = async () => {
  const formData = new FormData()
  formData.append('theme', settings.theme)
  if (settings.logo) formData.append('logo', settings.logo)
  if (settings.backgroundImage) formData.append('backgroundImage', settings.backgroundImage)
  if (settings.maxUploadSize) formData.append('maxUploadSize', settings.maxUploadSize.toString())
  formData.append('uploadBoxTransparency', settings.uploadBoxTransparency.toString())
  formData.append('blurIntensity', settings.blurIntensity.toString())

  try {
    const response = await axios.post('http://localhost:8080/settings/save', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    })
    message.value = { text: 'Settings saved successfully!', type: 'success' }
    console.log('Saved settings response:', response.data)
    settings.logo = null
    settings.backgroundImage = null
    await loadSettings()
    
    // Clear message after 3 seconds
    setTimeout(() => {
      message.value = null
    }, 3000)
  } catch (error) {
    console.error('Error saving settings:', error)
    message.value = { text: 'Failed to save settings.', type: 'error' }
    
    // Clear message after 3 seconds
    setTimeout(() => {
      message.value = null
    }, 3000)
  }
}

const saveProfile = async () => {
  try {
    await axios.post('http://localhost:8080/profile', profile)
    message.value = { text: 'Profile saved successfully!', type: 'success' }
    
    // Clear message after 3 seconds
    setTimeout(() => {
      message.value = null
    }, 3000)
  } catch (error) {
    console.error('Error saving profile:', error)
    message.value = { text: 'Failed to save profile.', type: 'error' }
    
    // Clear message after 3 seconds
    setTimeout(() => {
      message.value = null
    }, 3000)
  }
}

const inviteUser = async () => {
  try {
    const res = await axios.post('http://localhost:8080/users/invite')
    users.value.push({ id: res.data.id, name: res.data.name })
    message.value = { text: 'User invited successfully!', type: 'success' }
    
    // Clear message after 3 seconds
    setTimeout(() => {
      message.value = null
    }, 3000)
  } catch (error) {
    console.error('Error inviting user:', error)
    message.value = { text: 'Failed to invite user.', type: 'error' }
    
    // Clear message after 3 seconds
    setTimeout(() => {
      message.value = null
    }, 3000)
  }
}

const removeUser = (id: string) => {
  users.value = users.value.filter(user => user.id !== id)
}

const loadSettings = async () => {
  try {
    const response = await axios.get('http://localhost:8080/settings')
    console.log('Loaded settings:', response.data)
    
    logoPath.value = response.data.logo || null
    backgroundPath.value = response.data.backgroundImage || null
    
    settings.theme = response.data.theme || 'light'
    settings.maxUploadSize = response.data.maxUploadSize || '104857600'
    settings.uploadBoxTransparency = response.data.uploadBoxTransparency !== undefined ? response.data.uploadBoxTransparency : 80
    settings.blurIntensity = response.data.blurIntensity !== undefined ? response.data.blurIntensity : 12
    
    console.log('Final transparency:', settings.uploadBoxTransparency, 'blur:', settings.blurIntensity)
  } catch (error) {
    console.error('Error loading settings:', error)
    // Set fallback background if loading fails
    if (!backgroundPath.value) {
      backgroundPath.value = 'https://4kwallpapers.com/images/wallpapers/jungle-tree-dark-3840x2160-22695.jpg'
    }
  }
}

onMounted(() => {
  loadSettings()
})
</script>

<style scoped>
/* ==================== MODERN DESIGN SYSTEM ==================== */

/* Enhanced dark theme with depth */
.min-h-screen {
  background: linear-gradient(135deg, #0f0f23 0%, #1a1a2e 50%, #16213e 100%);
}

/* Compact modern sliders */
.compact-slider {
  appearance: none;
  height: 6px;
  background: linear-gradient(90deg, #374151 0%, #4b5563 50%, #6b7280 100%);
  border-radius: 8px;
  outline: none;
  transition: all 0.3s ease;
  border: 1px solid rgba(255, 255, 255, 0.1);
  box-shadow: inset 0 1px 2px rgba(0, 0, 0, 0.3);
}

.compact-slider:hover {
  background: linear-gradient(90deg, #4b5563 0%, #6b7280 50%, #9ca3af 100%);
  box-shadow: inset 0 1px 2px rgba(0, 0, 0, 0.3), 0 0 15px rgba(59, 130, 246, 0.2);
}

.compact-slider::-webkit-slider-thumb {
  appearance: none;
  width: 18px;
  height: 18px;
  border-radius: 50%;
  background: linear-gradient(135deg, #3b82f6 0%, #8b5cf6 50%, #ec4899 100%);
  cursor: pointer;
  border: 2px solid rgba(255, 255, 255, 0.9);
  box-shadow: 0 2px 8px rgba(59, 130, 246, 0.4), 0 0 0 2px rgba(59, 130, 246, 0.2);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.compact-slider::-webkit-slider-thumb:hover {
  transform: scale(1.15);
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.6), 0 0 0 4px rgba(139, 92, 246, 0.3);
  background: linear-gradient(135deg, #2563eb 0%, #7c3aed 50%, #db2777 100%);
}

.compact-slider::-webkit-slider-thumb:active {
  transform: scale(1.05);
  box-shadow: 0 2px 6px rgba(59, 130, 246, 0.8), 0 0 0 6px rgba(236, 72, 153, 0.4);
}

.compact-slider::-moz-range-thumb {
  width: 18px;
  height: 18px;
  border-radius: 50%;
  background: linear-gradient(135deg, #3b82f6 0%, #8b5cf6 50%, #ec4899 100%);
  cursor: pointer;
  border: 2px solid rgba(255, 255, 255, 0.9);
  box-shadow: 0 2px 8px rgba(59, 130, 246, 0.4);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.compact-slider::-moz-range-thumb:hover {
  transform: scale(1.15);
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.6);
}

/* Enhanced focus states with compact rings */
input:focus, select:focus {
  box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.3), 0 0 10px rgba(59, 130, 246, 0.15);
}

/* Smooth transitions for all interactive elements */
button, input, select, .group {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

/* Compact file input styling */
input[type="file"]::-webkit-file-upload-button {
  background: linear-gradient(135deg, #3b82f6 0%, #8b5cf6 100%);
  border: none;
  color: white;
  padding: 8px 16px;
  border-radius: 8px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 2px 6px rgba(59, 130, 246, 0.3);
  font-size: 12px;
}

input[type="file"]::-webkit-file-upload-button:hover {
  background: linear-gradient(135deg, #2563eb 0%, #7c3aed 100%);
  transform: translateY(-1px);
  box-shadow: 0 4px 10px rgba(59, 130, 246, 0.4);
}

/* Advanced backdrop blur effects */
.backdrop-blur-sm {
  backdrop-filter: blur(8px) saturate(180%);
  -webkit-backdrop-filter: blur(8px) saturate(180%);
}

/* Custom animations */
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.group:hover {
  transform: translateY(-1px);
}

/* Gradient text effects */
.bg-gradient-to-r {
  background-size: 200% 200%;
  animation: gradientShift 6s ease-in-out infinite;
}

@keyframes gradientShift {
  0%, 100% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
}

/* Compact shadow systems */
.shadow-xl {
  box-shadow: 
    0 15px 30px -6px rgba(0, 0, 0, 0.4),
    0 0 0 1px rgba(255, 255, 255, 0.05),
    inset 0 1px 0 rgba(255, 255, 255, 0.1);
}

/* Modern glass morphism cards */
.group {
  position: relative;
  overflow: hidden;
}

.group::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.1), transparent);
  transition: left 0.5s ease;
}

.group:hover::before {
  left: 100%;
}

/* Responsive optimizations */
@media (max-width: 1024px) {
  .grid-cols-1.lg\\:grid-cols-2 {
    grid-template-columns: 1fr;
  }
  
  .compact-slider::-webkit-slider-thumb {
    width: 16px;
    height: 16px;
  }
  
  .compact-slider::-moz-range-thumb {
    width: 16px;
    height: 16px;
  }
}

@media (max-width: 768px) {
  .text-3xl {
    font-size: 1.875rem;
  }
  
  .space-x-4 {
    gap: 0.75rem;
  }
  
  .p-5 {
    padding: 1rem;
  }
}

/* Reduced motion support */
@media (prefers-reduced-motion: reduce) {
  * {
    animation-duration: 0.01ms !important;
    animation-iteration-count: 1 !important;
    transition-duration: 0.01ms !important;
  }
  
  .group::before {
    display: none;
  }
  
  .group:hover {
    transform: none;
  }
}

/* High contrast mode support */
@media (prefers-contrast: high) {
  .bg-gradient-to-br {
    background: #1f2937;
    border: 2px solid #ffffff;
  }
  
  .text-gray-400 {
    color: #ffffff;
  }
  
  .border-gray-700 {
    border-color: #ffffff;
  }
}
</style>