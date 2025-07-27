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

      <!-- Settings container with glassmorphism -->
      <div class="flex items-center justify-center min-h-screen pt-16">
        <div :class="settingsContainerClasses" :style="settingsContainerStyle">
          <!-- Refined Modern Header -->
          <div class="mb-8 relative">
            <!-- Subtle floating particles -->
            <div class="absolute inset-0 overflow-hidden pointer-events-none">
              <div class="floating-particle-small" style="left: 15%; animation-delay: 0s;"></div>
              <div class="floating-particle-small" style="left: 35%; animation-delay: 3s;"></div>
              <div class="floating-particle-small" style="left: 65%; animation-delay: 1.5s;"></div>
              <div class="floating-particle-small" style="left: 85%; animation-delay: 4s;"></div>
            </div>
            
            <!-- Refined title -->
            <div class="text-center relative">
              <h1 class="text-3xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-blue-600 via-purple-600 to-pink-600 tracking-tight mb-2 relative z-10">
                Settings
              </h1>
              <p class="text-sm text-gray-600/70 font-medium">Customize your experience</p>
            </div>
            
            <!-- Sleek Tab Navigation -->
            <div class="mt-6 relative">
              <div class="bg-white/8 backdrop-blur-md rounded-2xl p-1 border border-white/15 shadow-lg">
                <div class="flex space-x-0.5">
                  <button
                    v-for="tab in tabs"
                    :key="tab"
                    @click="activeTab = tab"
                    :class="[
                      'relative flex-1 text-xs font-semibold px-4 py-2.5 rounded-xl transition-all duration-300 group overflow-hidden',
                      activeTab === tab 
                        ? 'bg-white text-gray-900 shadow-md transform scale-[1.02]' 
                        : 'text-gray-700 hover:text-gray-900 hover:bg-white/20'
                    ]"
                  >
                    <span class="relative z-10">{{ tab }}</span>
                    <div v-if="activeTab === tab" class="absolute bottom-0 left-1/2 transform -translate-x-1/2 w-4 h-0.5 bg-gradient-to-r from-blue-500 to-purple-600 rounded-full"></div>
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- Tab Content -->
          <div class="space-y-6">
            <!-- Refined Branding & UI Tab -->
            <div v-if="activeTab === 'Branding & UI'" class="space-y-6 animate-slide-in">
              <!-- Compact grid layout -->
              <div class="grid grid-cols-1 lg:grid-cols-2 gap-4">
                
                <!-- Theme Selection Card -->
                <div :class="compactCardClasses" :style="compactCardStyle" class="group hover:scale-[1.02]">
                  <div class="p-5 relative">
                    <!-- Minimal card header -->
                    <div class="flex items-center space-x-3 mb-4">
                      <div class="w-8 h-8 rounded-xl bg-gradient-to-br from-blue-500 to-purple-600 flex items-center justify-center">
                        <SwatchIcon class="w-4 h-4 text-white" />
                      </div>
                      <div>
                        <h3 class="text-sm font-semibold text-gray-800">Theme</h3>
                        <p class="text-xs text-gray-600">Appearance mode</p>
                      </div>
                    </div>
                    
                    <select
                      id="theme"
                      v-model="settings.theme"
                      :class="compactInputClasses"
                      :style="compactInputStyle"
                    >
                      <option value="light">☀️ Light</option>
                      <option value="dark">🌙 Dark</option>
                      <option value="system">⚡ Auto</option>
                    </select>
                  </div>
                </div>

                <!-- Upload Size Card -->
                <div :class="compactCardClasses" :style="compactCardStyle" class="group hover:scale-[1.02]">
                  <div class="p-5 relative">
                    <div class="flex items-center space-x-3 mb-4">
                      <div class="w-8 h-8 rounded-xl bg-gradient-to-br from-green-500 to-blue-600 flex items-center justify-center">
                        <DocumentArrowUpIcon class="w-4 h-4 text-white" />
                      </div>
                      <div>
                        <h3 class="text-sm font-semibold text-gray-800">Upload Limit</h3>
                        <p class="text-xs text-gray-600">Max file size</p>
                      </div>
                    </div>
                    
                    <select
                      id="maxUploadSize"
                      v-model="settings.maxUploadSize"
                      :class="compactInputClasses"
                      :style="compactInputStyle"
                    >
                      <option value="1048576">1 MB</option>
                      <option value="10485760">10 MB</option>
                      <option value="52428800">50 MB</option>
                      <option value="104857600">100 MB</option>
                      <option value="1073741824">1 GB</option>
                    </select>
                  </div>
                </div>

                <!-- Logo Upload Card -->
                <div :class="compactCardClasses" :style="compactCardStyle" class="group hover:scale-[1.02]">
                  <div class="p-5 relative">
                    <div class="flex items-center space-x-3 mb-4">
                      <div class="w-8 h-8 rounded-xl bg-gradient-to-br from-purple-500 to-pink-600 flex items-center justify-center">
                        <PhotoIcon class="w-4 h-4 text-white" />
                      </div>
                      <div>
                        <h3 class="text-sm font-semibold text-gray-800">Logo</h3>
                        <p class="text-xs text-gray-600">Brand image</p>
                      </div>
                    </div>
                    
                    <input
                      id="logo"
                      type="file"
                      accept="image/*"
                      @change="handleLogoUpload"
                      :class="compactFileInputClasses"
                      :style="compactInputStyle"
                    />
                  </div>
                </div>

                <!-- Background Image Card -->
                <div :class="compactCardClasses" :style="compactCardStyle" class="group hover:scale-[1.02]">
                  <div class="p-5 relative">
                    <div class="flex items-center space-x-3 mb-4">
                      <div class="w-8 h-8 rounded-xl bg-gradient-to-br from-orange-500 to-red-600 flex items-center justify-center">
                        <RectangleStackIcon class="w-4 h-4 text-white" />
                      </div>
                      <div>
                        <h3 class="text-sm font-semibold text-gray-800">Background</h3>
                        <p class="text-xs text-gray-600">Custom image</p>
                      </div>
                    </div>
                    
                    <input
                      id="backgroundImage"
                      type="file"
                      accept="image/*"
                      @change="handleBackgroundUpload"
                      :class="compactFileInputClasses"
                      :style="compactInputStyle"
                    />
                  </div>
                </div>
              </div>

              <!-- Refined Controls Section -->
              <div class="mt-8">
                <div class="text-center mb-6">
                  <h2 class="text-xl font-semibold text-transparent bg-clip-text bg-gradient-to-r from-cyan-600 to-purple-600 mb-1">
                    Visual Effects
                  </h2>
                  <p class="text-xs text-gray-600">Fine-tune appearance</p>
                </div>
                
                <div class="grid grid-cols-1 xl:grid-cols-2 gap-4">
                  <!-- Transparency Control -->
                  <div :class="compactCardClasses" :style="compactCardStyle" class="group hover:scale-[1.02]">
                    <div class="p-5">
                      <div class="flex items-center space-x-3 mb-6">
                        <div class="w-10 h-10 rounded-2xl bg-gradient-to-br from-cyan-500 to-blue-600 flex items-center justify-center">
                          <EyeIcon class="w-5 h-5 text-white" />
                        </div>
                        <div>
                          <h3 class="text-base font-semibold text-gray-800">Transparency</h3>
                          <p class="text-xs text-gray-600">Background opacity</p>
                        </div>
                      </div>
                      
                      <div class="space-y-4">
                        <div class="flex items-center space-x-4">
                          <input
                            id="uploadBoxTransparency"
                            type="range"
                            min="0"
                            max="100"
                            step="5"
                            v-model="settings.uploadBoxTransparency"
                            class="flex-1 compact-slider"
                          />
                          <div class="bg-white/80 backdrop-blur-md rounded-lg px-3 py-1.5 border border-white/30">
                            <span class="text-sm font-semibold text-gray-800">{{ settings.uploadBoxTransparency }}%</span>
                          </div>
                        </div>
                        <div class="flex justify-between text-xs text-gray-500">
                          <span>Opaque</span>
                          <span>Transparent</span>
                        </div>
                      </div>
                    </div>
                  </div>

                  <!-- Blur Intensity Control -->
                  <div :class="compactCardClasses" :style="compactCardStyle" class="group hover:scale-[1.02]">
                    <div class="p-5">
                      <div class="flex items-center space-x-3 mb-6">
                        <div class="w-10 h-10 rounded-2xl bg-gradient-to-br from-indigo-500 to-purple-600 flex items-center justify-center">
                          <SparklesIcon class="w-5 h-5 text-white" />
                        </div>
                        <div>
                          <h3 class="text-base font-semibold text-gray-800">Blur Effect</h3>
                          <p class="text-xs text-gray-600">Glass effect strength</p>
                        </div>
                      </div>
                      
                      <div class="space-y-4">
                        <div class="flex items-center space-x-4">
                          <input
                            id="blurIntensity"
                            type="range"
                            min="0"
                            max="24"
                            step="1"
                            v-model="settings.blurIntensity"
                            class="flex-1 compact-slider"
                          />
                          <div class="bg-white/80 backdrop-blur-md rounded-lg px-3 py-1.5 border border-white/30">
                            <span class="text-sm font-semibold text-gray-800">{{ getBlurLabel(settings.blurIntensity) }}</span>
                          </div>
                        </div>
                        <div class="flex justify-between text-xs text-gray-500">
                          <span>Sharp</span>
                          <span>Blurred</span>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <!-- Refined Save Button -->
              <div class="flex justify-center pt-6">
                <button
                  @click="saveSettings"
                  class="group relative px-8 py-3 bg-gradient-to-r from-blue-600 to-purple-600 text-white rounded-xl font-semibold text-sm hover:shadow-lg transform hover:-translate-y-0.5 transition-all duration-300 overflow-hidden"
                >
                  <div class="absolute inset-0 bg-gradient-to-r from-white/0 via-white/20 to-white/0 translate-x-[-100%] group-hover:translate-x-[100%] transition-transform duration-500"></div>
                  
                  <span class="relative z-10 flex items-center space-x-2">
                    <span>Save Settings</span>
                    <div class="w-1.5 h-1.5 bg-white rounded-full opacity-0 group-hover:opacity-100 animate-ping"></div>
                  </span>
                </button>
              </div>
            </div>

            <!-- Users Tab -->
            <div v-if="activeTab === 'Users'" class="space-y-6">
              <div :class="inputGroupClasses" :style="inputGroupStyle">
                <div class="p-6">
                  <h2 class="text-lg font-bold text-gray-800 mb-6 flex items-center space-x-2">
                    <div class="w-3 h-3 rounded-full bg-gradient-to-r from-blue-500 to-purple-500"></div>
                    <span>User Management</span>
                  </h2>
                  
                  <div class="space-y-4">
                    <div v-for="user in users" :key="user.id" class="flex items-center justify-between p-4 bg-white/60 backdrop-blur-sm rounded-xl border border-white/30">
                      <span class="font-semibold text-gray-800">{{ user.name }}</span>
                      <button @click="removeUser(user.id)" class="px-4 py-2 bg-red-500 hover:bg-red-600 text-white rounded-lg text-sm font-semibold transition-all duration-200 hover:scale-105">
                        Remove
                      </button>
                    </div>
                  </div>
                  
                  <div class="mt-6">
                    <button
                      @click="inviteUser"
                      class="w-full px-6 py-3 bg-gradient-to-r from-[var(--button-from)] to-[var(--button-to)] text-white rounded-xl font-semibold hover:opacity-90 transform hover:-translate-y-1 transition-all duration-300 shadow-lg"
                    >
                      Invite User
                    </button>
                  </div>
                </div>
              </div>
            </div>

            <!-- My Profile Tab -->
            <div v-if="activeTab === 'My Profile'" class="space-y-6">
              <div :class="inputGroupClasses" :style="inputGroupStyle">
                <div class="p-6">
                  <h2 class="text-lg font-bold text-gray-800 mb-6 flex items-center space-x-2">
                    <div class="w-3 h-3 rounded-full bg-gradient-to-r from-green-500 to-blue-500"></div>
                    <span>Profile Settings</span>
                  </h2>
                  
                  <div class="space-y-6">
                    <div>
                      <label for="email" class="block text-sm font-bold text-gray-800 mb-3 flex items-center space-x-2">
                        <div class="w-2 h-2 rounded-full bg-gradient-to-r from-blue-500 to-cyan-500"></div>
                        <span>Email</span>
                      </label>
                      <input
                        id="email"
                        type="email"
                        v-model="profile.email"
                        :class="inputClasses"
                        :style="inputStyle"
                        placeholder="your.email@example.com"
                      />
                    </div>
                    
                    <div>
                      <label for="password" class="block text-sm font-bold text-gray-800 mb-3 flex items-center space-x-2">
                        <div class="w-2 h-2 rounded-full bg-gradient-to-r from-purple-500 to-pink-500"></div>
                        <span>Password</span>
                      </label>
                      <input
                        id="password"
                        type="password"
                        v-model="profile.password"
                        :class="inputClasses"
                        :style="inputStyle"
                        placeholder="••••••••"
                      />
                    </div>
                    
                    <div class="pt-4">
                      <button
                        @click="saveProfile"
                        class="w-full px-6 py-3 bg-gradient-to-r from-[var(--button-from)] to-[var(--button-to)] text-white rounded-xl font-semibold hover:opacity-90 transform hover:-translate-y-1 transition-all duration-300 shadow-lg"
                      >
                        Save Profile
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Success/Error Message -->
          <div v-if="message" class="mt-8 text-center animate-slide-up">
            <div :class="[
              'inline-flex items-center px-6 py-3 rounded-xl font-semibold',
              message.type === 'success' 
                ? 'bg-green-500/20 text-green-800 border border-green-500/30' 
                : 'bg-red-500/20 text-red-800 border border-red-500/30'
            ]" :style="{ backdropFilter: 'blur(10px)' }">
              <div :class="[
                'w-2 h-2 rounded-full mr-3',
                message.type === 'success' ? 'bg-green-500' : 'bg-red-500'
              ]" />
              {{ message.text }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import axios from 'axios'
import { 
  CloudArrowUpIcon, 
  SwatchIcon, 
  DocumentArrowUpIcon, 
  PhotoIcon, 
  RectangleStackIcon,
  EyeIcon,
  SparklesIcon
} from '@heroicons/vue/24/outline'

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

// Computed properties for glassmorphism styling
const settingsContainerClasses = computed(() => {
  return 'w-full max-w-4xl bg-white/10 backdrop-blur-lg rounded-3xl border border-white/20 shadow-2xl p-8 relative overflow-hidden'
})

const settingsContainerStyle = computed(() => {
  const opacity = Math.max((100 - settings.uploadBoxTransparency) / 100, 0.15)
  const blurValue = settings.blurIntensity
  
  return {
    backgroundColor: `rgba(255, 255, 255, ${opacity})`,
    backdropFilter: blurValue > 0 ? `blur(${blurValue}px)` : 'none',
    '-webkit-backdrop-filter': blurValue > 0 ? `blur(${blurValue}px)` : 'none'
  } as any
})

const inputGroupClasses = computed(() => {
  return 'relative rounded-2xl shadow-xl border border-white/30 overflow-hidden bg-white/20 backdrop-blur-sm hover:bg-white/30 transition-all duration-300'
})

const inputGroupStyle = computed(() => {
  const opacity = Math.max((100 - settings.uploadBoxTransparency) / 100 * 0.3, 0.1)
  const blurValue = settings.blurIntensity * 0.5
  
  return {
    backgroundColor: `rgba(255, 255, 255, ${opacity})`,
    backdropFilter: blurValue > 0 ? `blur(${blurValue}px)` : 'none',
    '-webkit-backdrop-filter': blurValue > 0 ? `blur(${blurValue}px)` : 'none'
  } as any
})

const inputClasses = computed(() => {
  return 'w-full bg-white/70 backdrop-blur-sm border border-white/30 rounded-xl px-4 py-3 text-gray-800 font-medium placeholder-gray-500 focus:ring-2 focus:ring-blue-400 focus:border-blue-400 focus:bg-white/90 outline-none transition-all duration-200'
})

const inputStyle = computed(() => {
  return {
    backdropFilter: 'blur(8px)',
    '-webkit-backdrop-filter': 'blur(8px)'
  }
})

// Compact styling for refined modern look
const compactCardClasses = computed(() => {
  return 'bg-white/20 backdrop-blur-md border border-white/30 rounded-2xl shadow-lg hover:shadow-xl transform transition-all duration-300 hover:bg-white/30'
})

const compactCardStyle = computed(() => {
  return {
    background: 'linear-gradient(135deg, rgba(255,255,255,0.1) 0%, rgba(255,255,255,0.3) 100%)',
    backdropFilter: `blur(${settings.blurIntensity || 12}px)`,
    boxShadow: '0 8px 32px rgba(31, 38, 135, 0.15)',
  }
})

const compactInputClasses = computed(() => {
  return 'w-full px-3 py-2.5 bg-white/80 backdrop-blur-md border border-white/40 rounded-xl text-gray-800 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500/30 focus:border-blue-500/50 transition-all duration-300'
})

const compactInputStyle = computed(() => {
  return {
    background: 'linear-gradient(135deg, rgba(255,255,255,0.8) 0%, rgba(255,255,255,0.6) 100%)',
    backdropFilter: 'blur(8px)',
  }
})

const compactFileInputClasses = computed(() => {
  return 'w-full px-3 py-2.5 bg-white/80 backdrop-blur-md border border-white/40 rounded-xl text-gray-800 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500/30 focus:border-blue-500/50 transition-all duration-300 file:mr-3 file:py-1.5 file:px-3 file:rounded-lg file:border-0 file:text-xs file:font-semibold file:bg-gradient-to-r file:from-blue-500 file:to-purple-600 file:text-white hover:file:shadow-md hover:file:-translate-y-0.5 file:transition-all file:duration-300'
})

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
/* ==================== ULTRA-MODERN ANIMATIONS ==================== */

/* Floating particles */
.floating-particle {
  position: absolute;
  width: 6px;
  height: 6px;
  background: linear-gradient(45deg, #3b82f6, #8b5cf6, #ec4899);
  border-radius: 50%;
  animation: float-particle 8s ease-in-out infinite;
  opacity: 0.6;
  filter: blur(1px);
}

@keyframes float-particle {
  0%, 100% {
    transform: translateY(0px) rotate(0deg);
    opacity: 0.6;
  }
  25% {
    transform: translateY(-100px) rotate(90deg);
    opacity: 0.8;
  }
  50% {
    transform: translateY(-50px) rotate(180deg);
    opacity: 0.4;
  }
  75% {
    transform: translateY(-80px) rotate(270deg);
    opacity: 0.9;
  }
}

/* Advanced gradient animations */
@keyframes gradient-x {
  0%, 100% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
}

.animate-gradient-x {
  background-size: 200% 200%;
  animation: gradient-x 4s ease infinite;
}

/* Ultra-modern glow effects */
@keyframes glow-pulse {
  0%, 100% {
    opacity: 0.5;
    transform: scale(1);
  }
  50% {
    opacity: 0.8;
    transform: scale(1.05);
  }
}

.animate-glow-pulse {
  animation: glow-pulse 3s ease-in-out infinite;
}

@keyframes pulse-glow {
  0%, 100% {
    box-shadow: 0 0 20px rgba(59, 130, 246, 0.4);
  }
  50% {
    box-shadow: 0 0 40px rgba(139, 92, 246, 0.8);
  }
}

.animate-pulse-glow {
  animation: pulse-glow 2s ease-in-out infinite;
}

/* Slide in animations */
@keyframes slide-in {
  from {
    opacity: 0;
    transform: translateY(30px) scale(0.95);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

.animate-slide-in {
  animation: slide-in 0.8s cubic-bezier(0.25, 0.46, 0.45, 0.94);
}

/* Tab indicator expansion */
@keyframes expand {
  from {
    width: 0;
  }
  to {
    width: 2rem;
  }
}

.animate-expand {
  animation: expand 0.3s ease-out;
}

/* Custom slow animations */
@keyframes pulse-slow {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: 0.3;
  }
}

.animate-pulse-slow {
  animation: pulse-slow 3s ease-in-out infinite;
}

@keyframes spin-slow {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

.animate-spin-slow {
  animation: spin-slow 8s linear infinite;
}

/* Fade in out effect */
@keyframes fade-in-out {
  0%, 100% {
    opacity: 0.3;
  }
  50% {
    opacity: 0.8;
  }
}

.animate-fade-in-out {
  animation: fade-in-out 2s ease-in-out infinite;
}

/* Blur demo animation */
@keyframes blur-demo {
  0%, 100% {
    filter: blur(0px);
  }
  50% {
    filter: blur(8px);
  }
}

.animate-blur-demo {
  animation: blur-demo 3s ease-in-out infinite;
}

/* ==================== COMPACT MODERN SLIDER STYLING ==================== */

.compact-slider {
  appearance: none;
  height: 6px;
  background: linear-gradient(90deg, #e2e8f0 0%, #cbd5e1 100%);
  border-radius: 12px;
  outline: none;
  border: 1px solid rgba(255, 255, 255, 0.3);
  backdrop-filter: blur(8px);
  box-shadow: 
    inset 0 1px 2px rgba(0, 0, 0, 0.1),
    0 4px 12px -2px rgba(0, 0, 0, 0.1);
  transition: all 0.2s ease;
}

.compact-slider:hover {
  background: linear-gradient(90deg, #cbd5e1 0%, #94a3b8 100%);
  box-shadow: 
    inset 0 1px 2px rgba(0, 0, 0, 0.1),
    0 6px 16px -2px rgba(59, 130, 246, 0.2);
}

.compact-slider::-webkit-slider-thumb {
  appearance: none;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  background: linear-gradient(135deg, #3b82f6 0%, #8b5cf6 100%);
  cursor: pointer;
  border: 2px solid rgba(255, 255, 255, 0.9);
  box-shadow: 
    0 4px 12px rgba(59, 130, 246, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.3);
  transition: all 0.2s cubic-bezier(0.25, 0.46, 0.45, 0.94);
}

.compact-slider::-webkit-slider-thumb:hover {
  transform: scale(1.1);
  box-shadow: 
    0 6px 16px rgba(59, 130, 246, 0.6),
    inset 0 1px 0 rgba(255, 255, 255, 0.4);
  background: linear-gradient(135deg, #2563eb 0%, #7c3aed 100%);
}

.compact-slider::-webkit-slider-thumb:active {
  transform: scale(1.05);
  box-shadow: 
    0 3px 8px rgba(59, 130, 246, 0.8),
    inset 0 1px 0 rgba(255, 255, 255, 0.5);
}

.compact-slider::-moz-range-thumb {
  width: 20px;
  height: 20px;
  border-radius: 50%;
  background: linear-gradient(135deg, #3b82f6 0%, #8b5cf6 100%);
  cursor: pointer;
  border: 2px solid rgba(255, 255, 255, 0.9);
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.4);
  transition: all 0.2s cubic-bezier(0.25, 0.46, 0.45, 0.94);
}

.compact-slider::-moz-range-thumb:hover {
  transform: scale(1.1);
  box-shadow: 0 6px 16px rgba(59, 130, 246, 0.6);
}

/* ==================== ULTRA-MODERN SLIDER STYLING ==================== */

.ultra-modern-slider {
  appearance: none;
  height: 8px;
  background: linear-gradient(90deg, 
    rgba(59, 130, 246, 0.3) 0%, 
    rgba(139, 92, 246, 0.5) 50%, 
    rgba(236, 72, 153, 0.3) 100%);
  border-radius: 20px;
  outline: none;
  border: 2px solid rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(10px);
  box-shadow: 
    inset 0 2px 4px rgba(0, 0, 0, 0.1),
    0 8px 25px -5px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

.ultra-modern-slider:hover {
  background: linear-gradient(90deg, 
    rgba(59, 130, 246, 0.5) 0%, 
    rgba(139, 92, 246, 0.7) 50%, 
    rgba(236, 72, 153, 0.5) 100%);
  box-shadow: 
    inset 0 2px 4px rgba(0, 0, 0, 0.1),
    0 15px 35px -5px rgba(59, 130, 246, 0.3);
}

.ultra-modern-slider::-webkit-slider-thumb {
  appearance: none;
  width: 28px;
  height: 28px;
  border-radius: 50%;
  background: linear-gradient(135deg, #3b82f6, #8b5cf6, #ec4899);
  cursor: pointer;
  border: 3px solid rgba(255, 255, 255, 0.9);
  box-shadow: 
    0 8px 20px rgba(0, 0, 0, 0.2),
    0 0 0 4px rgba(59, 130, 246, 0.2),
    inset 0 1px 0 rgba(255, 255, 255, 0.3);
  transition: all 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94);
  position: relative;
}

.ultra-modern-slider::-webkit-slider-thumb:hover {
  transform: scale(1.2);
  box-shadow: 
    0 12px 25px rgba(0, 0, 0, 0.3),
    0 0 0 6px rgba(139, 92, 246, 0.3),
    inset 0 1px 0 rgba(255, 255, 255, 0.4);
  background: linear-gradient(135deg, #2563eb, #7c3aed, #db2777);
}

.ultra-modern-slider::-webkit-slider-thumb:active {
  transform: scale(1.1);
  box-shadow: 
    0 6px 15px rgba(0, 0, 0, 0.4),
    0 0 0 8px rgba(236, 72, 153, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.5);
}

.ultra-modern-slider::-moz-range-thumb {
  width: 28px;
  height: 28px;
  border-radius: 50%;
  background: linear-gradient(135deg, #3b82f6, #8b5cf6, #ec4899);
  cursor: pointer;
  border: 3px solid rgba(255, 255, 255, 0.9);
  box-shadow: 
    0 8px 20px rgba(0, 0, 0, 0.2),
    0 0 0 4px rgba(59, 130, 246, 0.2);
  transition: all 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94);
}

.ultra-modern-slider::-moz-range-thumb:hover {
  transform: scale(1.2);
  box-shadow: 
    0 12px 25px rgba(0, 0, 0, 0.3),
    0 0 0 6px rgba(139, 92, 246, 0.3);
}

/* ==================== LEGACY ANIMATIONS ==================== */

.animate-slide-up {
  animation: slideUp 0.5s cubic-bezier(0.25, 0.46, 0.45, 0.94);
}

@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Enhanced hover effects */
.group:hover .group-hover\:scale-105 {
  transform: scale(1.05);
}

.hover\:scale-102:hover {
  transform: scale(1.02);
}

/* 3D transformation effects */
.group:hover {
  transform-style: preserve-3d;
}

/* Enhanced button effects */
button {
  position: relative;
  overflow: hidden;
}

button:hover {
  animation: float 3s ease-in-out infinite;
}

@keyframes float {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-4px);
  }
}

/* Advanced glow effects for focus states */
input:focus, select:focus {
  animation: glow 2s ease-in-out infinite;
}

@keyframes glow {
  0%, 100% {
    box-shadow: 
      0 0 20px rgba(59, 130, 246, 0.5),
      inset 0 1px 0 rgba(255, 255, 255, 0.6);
  }
  50% {
    box-shadow: 
      0 0 30px rgba(139, 92, 246, 0.8),
      inset 0 1px 0 rgba(255, 255, 255, 0.8);
  }
}

/* Glass morphism enhancement */
.backdrop-blur-xl {
  backdrop-filter: blur(24px) saturate(180%);
  -webkit-backdrop-filter: blur(24px) saturate(180%);
}

/* Advanced shadow definitions */
.shadow-3xl {
  box-shadow: 
    0 35px 60px -12px rgba(0, 0, 0, 0.25),
    0 0 0 1px rgba(255, 255, 255, 0.1),
    inset 0 1px 0 rgba(255, 255, 255, 0.2);
}

/* Responsive optimizations */
@media (max-width: 1024px) {
  .floating-particle {
    display: none;
  }
  
  .ultra-modern-slider::-webkit-slider-thumb {
    width: 24px;
    height: 24px;
  }
  
  .ultra-modern-slider::-moz-range-thumb {
    width: 24px;
    height: 24px;
  }
}

@media (max-width: 768px) {
  .settings-container {
    margin: 1rem;
    padding: 1.5rem;
  }
  
  .animate-glow-pulse,
  .animate-pulse-glow {
    animation: none;
  }
  
  .group:hover {
    transform: none;
  }
}

/* Reduced motion support */
@media (prefers-reduced-motion: reduce) {
  * {
    animation-duration: 0.01ms !important;
    animation-iteration-count: 1 !important;
    transition-duration: 0.01ms !important;
  }
  
  .floating-particle {
    display: none;
  }
}
</style>