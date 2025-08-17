<template>
  <div class="space-y-8">
    <div class="text-center">
      <h2 class="text-2xl font-bold mb-2 transition-colors duration-300"
          :style="{ color: isDark ? '#f9fafb' : '#111827' }">Admin Settings</h2>
      <p class="transition-colors duration-300"
         :style="{ color: isDark ? '#9ca3af' : '#4b5563' }">Configure global application settings</p>
    </div>

    <form @submit.prevent="saveSettings" class="space-y-8">
      <!-- Maximum Validity Setting -->
      <div class="relative group">
        <div class="absolute inset-0 bg-gradient-to-r from-purple-600/20 to-blue-600/20 rounded-2xl blur-xl group-hover:blur-2xl transition-all duration-300"></div>
        <div class="relative backdrop-blur-xl border rounded-2xl p-6 transition-all duration-300"
             :style="{
               backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
               borderColor: isDark ? 'rgba(147, 51, 234, 0.3)' : 'rgba(147, 51, 234, 0.2)'
             }">
          <h3 class="text-lg font-semibold mb-4 flex items-center transition-colors duration-300"
              :style="{ color: isDark ? '#f3e8ff' : '#111827' }">
            <ClockIcon class="w-5 h-5 mr-2"
                       :style="{ color: isDark ? '#a855f7' : '#7c3aed' }" />
            Maximum File Validity
          </h3>
          <div class="grid grid-cols-1 sm:grid-cols-5 gap-3">
            <label
              v-for="option in validityOptions"
              :key="option.value"
              :class="[
                'relative flex items-center justify-center px-4 py-3 rounded-xl border-2 cursor-pointer transition-all backdrop-blur-sm',
                settings.maxValidity === option.value
                  ? 'border-purple-400 bg-gradient-to-r from-purple-600 to-blue-600 text-white transform scale-105'
                  : (isDark 
                      ? 'border-gray-600/50 bg-gray-700/30 text-gray-200 hover:border-purple-400/50 hover:bg-purple-600/10' 
                      : 'border-gray-200/50 bg-white/70 text-gray-700 hover:border-purple-300/50 hover:bg-purple-50/80'
                    )
              ]"
            >
              <input
                type="radio"
                v-model="settings.maxValidity"
                :value="option.value"
                class="sr-only"
              />
              <div class="text-center">
                <div class="font-medium">{{ option.label }}</div>
                <div class="text-xs opacity-75">{{ option.description }}</div>
              </div>
            </label>
          </div>
          <p class="text-sm mt-3 transition-colors duration-300 opacity-80"
             :style="{ color: isDark ? '#d1d5db' : '#4b5563' }">
            Users will not be able to set file expiration longer than this maximum period.
          </p>
        </div>
      </div>

      <!-- User Registration Setting -->
      <div class="relative group">
        <div class="absolute inset-0 bg-gradient-to-r from-amber-600/20 to-orange-600/20 rounded-2xl blur-xl group-hover:blur-2xl transition-all duration-300"></div>
        <div class="relative backdrop-blur-xl border rounded-2xl p-6 transition-all duration-300"
             :style="{
               backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
               borderColor: isDark ? 'rgba(245, 158, 11, 0.3)' : 'rgba(245, 158, 11, 0.2)'
             }">
          <h3 class="text-lg font-semibold mb-4 flex items-center transition-colors duration-300"
              :style="{ color: isDark ? '#fef3c7' : '#111827' }">
            <UserGroupIcon class="w-5 h-5 mr-2"
                           :style="{ color: isDark ? '#fbbf24' : '#d97706' }" />
            User Registration
          </h3>
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm font-medium transition-colors duration-300"
                 :style="{ color: isDark ? '#fef3c7' : '#111827' }">Allow New User Registration</p>
              <p class="text-sm transition-colors duration-300 opacity-80"
                 :style="{ color: isDark ? '#d1d5db' : '#4b5563' }">When disabled, only existing users and admins can log in</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input
                type="checkbox"
                v-model="settings.allowRegistration"
                class="sr-only peer"
              />
              <div 
                :class="[
                  'w-11 h-6 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-orange-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[\'\'] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-gradient-to-r peer-checked:from-orange-500 peer-checked:to-amber-500',
                  isDark ? 'bg-gray-600/50 after:border-gray-400' : 'bg-gray-200/50 after:border-gray-300'
                ]"
              ></div>
            </label>
          </div>
        </div>
      </div>

      <!-- File Expiration Behavior Setting -->
      <div class="relative group">
        <div class="absolute inset-0 bg-gradient-to-r from-red-600/20 to-orange-600/20 rounded-2xl blur-xl group-hover:blur-2xl transition-all duration-300"></div>
        <div class="relative backdrop-blur-xl border rounded-2xl p-6 transition-all duration-300"
             :style="{
               backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
               borderColor: isDark ? 'rgba(234, 88, 12, 0.3)' : 'rgba(234, 88, 12, 0.2)'
             }">
          <h3 class="text-lg font-semibold mb-4 flex items-center transition-colors duration-300"
              :style="{ color: isDark ? '#fff7ed' : '#111827' }">
            <ClockIcon class="w-5 h-5 mr-2"
                       :style="{ color: isDark ? '#fb923c' : '#ea580c' }" />
            File Expiration Behavior
          </h3>
          <p class="text-sm mb-4 transition-colors duration-300 opacity-80"
             :style="{ color: isDark ? '#fed7aa' : '#6b7280' }">
            Choose what happens to files when they expire
          </p>
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <label
              :class="[
                'relative flex flex-col p-4 rounded-xl border-2 cursor-pointer transition-all backdrop-blur-sm',
                settings.expirationAction === 'unavailable'
                  ? 'border-orange-400 bg-gradient-to-r from-orange-500 to-red-500 text-white transform scale-105'
                  : (isDark 
                      ? 'border-gray-600/50 bg-gray-700/30 text-gray-200 hover:border-orange-400/50 hover:bg-orange-600/10' 
                      : 'border-gray-200/50 bg-white/70 text-gray-700 hover:border-orange-300/50 hover:bg-orange-50/80'
                    )
              ]"
            >
              <input
                type="radio"
                v-model="settings.expirationAction"
                value="unavailable"
                class="sr-only"
              />
              <div class="flex items-center mb-2">
                <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.878 9.878L3 3m6.878 6.878L12 12m-3.122-3.122L9.878 9.878"></path>
                </svg>
                <span class="font-medium">Make Unavailable</span>
              </div>
              <p class="text-sm opacity-90">Files become inaccessible but remain on the server</p>
            </label>
            
            <label
              :class="[
                'relative flex flex-col p-4 rounded-xl border-2 cursor-pointer transition-all backdrop-blur-sm',
                settings.expirationAction === 'delete'
                  ? 'border-orange-400 bg-gradient-to-r from-orange-500 to-red-500 text-white transform scale-105'
                  : (isDark 
                      ? 'border-gray-600/50 bg-gray-700/30 text-gray-200 hover:border-orange-400/50 hover:bg-orange-600/10' 
                      : 'border-gray-200/50 bg-white/70 text-gray-700 hover:border-orange-300/50 hover:bg-orange-50/80'
                    )
              ]"
            >
              <input
                type="radio"
                v-model="settings.expirationAction"
                value="delete"
                class="sr-only"
              />
              <div class="flex items-center mb-2">
                <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                </svg>
                <span class="font-medium">Delete Files</span>
              </div>
              <p class="text-sm opacity-90">Files are permanently removed from the server</p>
            </label>
          </div>
          <div class="mt-4 p-3 rounded-xl border backdrop-blur-sm transition-colors duration-300"
               :style="{
                 backgroundColor: isDark ? 'rgba(251, 146, 60, 0.1)' : 'rgba(234, 88, 12, 0.1)',
                 borderColor: isDark ? 'rgba(251, 146, 60, 0.3)' : 'rgba(234, 88, 12, 0.2)'
               }">
            <p class="text-sm transition-colors duration-300"
               :style="{ color: isDark ? '#fed7aa' : '#9a3412' }">
              <strong>Note:</strong> This setting affects how the system handles files after they expire.
              The cleanup process runs every hour.
            </p>
          </div>
        </div>
      </div>

      <!-- Upload Settings -->
      <div class="relative group">
        <div class="absolute inset-0 bg-gradient-to-r from-green-600/20 to-emerald-600/20 rounded-2xl blur-xl group-hover:blur-2xl transition-all duration-300"></div>
        <div class="relative backdrop-blur-xl border rounded-2xl p-6 transition-all duration-300"
             :style="{
               backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
               borderColor: isDark ? 'rgba(16, 185, 129, 0.3)' : 'rgba(16, 185, 129, 0.2)'
             }">
          <h3 class="text-lg font-semibold mb-4 flex items-center transition-colors duration-300"
              :style="{ color: isDark ? '#d1fae5' : '#111827' }">
            <CloudArrowUpIcon class="w-5 h-5 mr-2"
                              :style="{ color: isDark ? '#34d399' : '#059669' }" />
            Upload Settings
          </h3>
          <div class="max-w-md">
            <div>
              <label class="block text-sm font-medium mb-2 transition-colors duration-300"
                     :style="{ color: isDark ? '#d1fae5' : '#374151' }">
                Maximum Upload Size
              </label>
              <select
                v-model="settings.maxUploadSize"
                class="w-full px-3 py-2 border rounded-xl focus:ring-2 focus:ring-green-500 focus:border-green-500 transition-all duration-300 backdrop-blur-sm"
                :style="{
                  backgroundColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(255,255,255,0.8)',
                  borderColor: isDark ? 'rgba(16, 185, 129, 0.3)' : 'rgba(16, 185, 129, 0.2)',
                  color: isDark ? '#f9fafb' : '#111827'
                }"
              >
                <option value="1048576">1 MB</option>
                <option value="5242880">5 MB</option>
                <option value="10485760">10 MB</option>
                <option value="26214400">25 MB</option>
                <option value="52428800">50 MB</option>
                <option value="104857600">100 MB</option>
                <option value="262144000">250 MB</option>
                <option value="524288000">500 MB</option>
                <option value="1073741824">1 GB</option>
                <option value="2147483648">2 GB</option>
                <option value="5368709120">5 GB</option>
                <option value="10737418240">10 GB</option>
                <option value="-1">Unlimited</option>
              </select>
              <p class="text-sm mt-1 transition-colors duration-300 opacity-80"
                 :style="{ color: isDark ? '#d1d5db' : '#4b5563' }">
                Set the maximum file size users can upload. Choose "Unlimited" to allow any size.
              </p>
            </div>
          </div>
        </div>
      </div>

      <!-- Branding Settings -->
      <div class="relative group">
        <div class="absolute inset-0 bg-gradient-to-r from-red-600/20 to-pink-600/20 rounded-2xl blur-xl group-hover:blur-2xl transition-all duration-300"></div>
        <div class="relative backdrop-blur-xl border rounded-2xl p-6 transition-all duration-300"
             :style="{
               backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
               borderColor: isDark ? 'rgba(239, 68, 68, 0.3)' : 'rgba(239, 68, 68, 0.2)'
             }">
          <h3 class="text-lg font-semibold mb-4 flex items-center transition-colors duration-300"
              :style="{ color: isDark ? '#fef2f2' : '#111827' }">
            <PhotoIcon class="w-5 h-5 mr-2"
                       :style="{ color: isDark ? '#f87171' : '#dc2626' }" />
            Branding
          </h3>
          <div class="space-y-6">
            <!-- Application Title -->
            <div>
              <label class="block text-sm font-medium mb-2 transition-colors duration-300"
                     :style="{ color: isDark ? '#fef2f2' : '#374151' }">
                Application Title
              </label>
              <input
                type="text"
                v-model="settings.navbarTitle"
                placeholder="PinGO"
                class="w-full px-4 py-2 border rounded-xl focus:ring-2 focus:ring-red-500 focus:border-red-500 transition-all duration-300 backdrop-blur-sm"
                :style="{
                  backgroundColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(255,255,255,0.8)',
                  borderColor: isDark ? 'rgba(239, 68, 68, 0.3)' : 'rgba(239, 68, 68, 0.2)',
                  color: isDark ? '#f9fafb' : '#111827'
                }"
              />
              <p class="text-sm mt-1 transition-colors duration-300 opacity-80"
                 :style="{ color: isDark ? '#d1d5db' : '#4b5563' }">
                This title will appear next to the logo in the navigation bar.
              </p>
            </div>
            
            <!-- Logo and Background Images -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium mb-2 transition-colors duration-300"
                       :style="{ color: isDark ? '#fef2f2' : '#374151' }">
                  Logo
                </label>
                <div class="flex items-center space-x-4">
                  <img
                    v-if="settings.logo"
                    :src="`http://localhost:8080${settings.logo}`"
                    alt="Current logo"
                    class="w-12 h-12 object-contain rounded-xl border backdrop-blur-sm"
                    :style="{ borderColor: isDark ? 'rgba(239, 68, 68, 0.3)' : 'rgba(239, 68, 68, 0.2)' }"
                  />
                  <input
                    type="file"
                    ref="logoInput"
                    @change="handleLogoUpload"
                    accept="image/*"
                    class="block w-full text-sm file:mr-4 file:py-2 file:px-4 file:rounded-xl file:border-0 file:bg-gradient-to-r file:from-red-500 file:to-pink-500 file:text-white hover:file:from-red-600 hover:file:to-pink-600 transition-all duration-300"
                    :style="{ color: isDark ? '#d1d5db' : '#6b7280' }"
                  />
                </div>
              </div>
              <div>
                <label class="block text-sm font-medium mb-2 transition-colors duration-300"
                       :style="{ color: isDark ? '#fef2f2' : '#374151' }">
                  Background Image
                </label>
                <div class="flex items-center space-x-4">
                  <img
                    v-if="settings.backgroundImage"
                    :src="`http://localhost:8080${settings.backgroundImage}`"
                    alt="Current background"
                    class="w-12 h-12 object-cover rounded-xl border backdrop-blur-sm"
                    :style="{ borderColor: isDark ? 'rgba(239, 68, 68, 0.3)' : 'rgba(239, 68, 68, 0.2)' }"
                  />
                  <input
                    type="file"
                    ref="backgroundInput"
                    @change="handleBackgroundUpload"
                    accept="image/*"
                    class="block w-full text-sm file:mr-4 file:py-2 file:px-4 file:rounded-xl file:border-0 file:bg-gradient-to-r file:from-red-500 file:to-pink-500 file:text-white hover:file:from-red-600 hover:file:to-pink-600 transition-all duration-300"
                    :style="{ color: isDark ? '#d1d5db' : '#6b7280' }"
                  />
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Color Customization Settings -->
      <div class="relative group">
        <div class="absolute inset-0 bg-gradient-to-r from-pink-600/20 to-red-600/20 rounded-2xl blur-xl group-hover:blur-2xl transition-all duration-300"></div>
        <div class="relative backdrop-blur-xl border rounded-2xl p-6 transition-all duration-300"
             :style="{
               backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
               borderColor: isDark ? 'rgba(236, 72, 153, 0.3)' : 'rgba(236, 72, 153, 0.2)'
             }">
          <h3 class="text-lg font-semibold mb-4 flex items-center transition-colors duration-300"
              :style="{ color: isDark ? '#fdf2f8' : '#111827' }">
            <SwatchIcon class="w-5 h-5 mr-2"
                        :style="{ color: isDark ? '#ec4899' : '#be185d' }" />
            Color Customization
          </h3>
          
          <div class="space-y-6">
            <!-- Website Color -->
            <div>
              <ColorPicker
                v-model="settings.websiteColor"
                label="Website Color (WebGL Background)"
                input-id="website-color"
              />
              <p class="text-sm mt-2 transition-colors duration-300 opacity-80"
                 :style="{ color: isDark ? '#d1d5db' : '#4b5563' }">
                Primary color used for the WebGL background animation on home and download pages.
              </p>
            </div>

            <!-- UI Gradient Colors -->
            <div class="space-y-4">
              <div class="flex items-center justify-between">
                <div>
                  <h4 class="font-medium transition-colors duration-300"
                      :style="{ color: isDark ? '#fdf2f8' : '#111827' }">
                    UI Gradient Colors
                  </h4>
                  <p class="text-sm transition-colors duration-300 opacity-80"
                     :style="{ color: isDark ? '#d1d5db' : '#4b5563' }">
                    Colors used for buttons and UI elements gradients.
                  </p>
                </div>
                <button
                  @click="showGradientDialog = true"
                  type="button"
                  class="px-4 py-2 bg-gradient-to-r from-pink-600 to-purple-600 text-white font-medium rounded-xl transition-all duration-200 hover:scale-105"
                >
                  Configure Gradient
                </button>
              </div>

              <!-- Gradient Preview -->
              <div class="h-16 rounded-xl border-2 border-white shadow-lg relative overflow-hidden"
                   :style="{ background: currentGradient }">
                <div class="absolute inset-0 flex items-center justify-center">
                  <span class="text-white font-semibold drop-shadow-lg">
                    Current UI Gradient
                  </span>
                </div>
              </div>

              <!-- Individual Color Pickers (for fine-tuning) -->
              <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
                <ColorPicker
                  v-model="settings.gradientColor1"
                  label="Color 1"
                  input-id="gradient-color-1"
                  :show-presets="false"
                />
                <ColorPicker
                  v-model="settings.gradientColor2"
                  label="Color 2"
                  input-id="gradient-color-2"
                  :show-presets="false"
                />
                <ColorPicker
                  v-model="settings.gradientColor3"
                  label="Color 3"
                  input-id="gradient-color-3"
                  :show-presets="false"
                />
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Save Button -->
      <div class="flex justify-end">
        <button
          type="submit"
          :disabled="isLoading"
          class="relative overflow-hidden group px-8 py-4 bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700 text-white font-semibold rounded-2xl transition-all duration-300 disabled:opacity-50 disabled:cursor-not-allowed shadow-lg hover:shadow-2xl transform hover:scale-105"
        >
          <div class="absolute inset-0 bg-gradient-to-r from-white/20 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
          <div class="relative flex items-center">
            <ScaleIcon v-if="!isLoading" class="w-5 h-5 mr-2" />
            <div v-else class="w-5 h-5 mr-2 animate-spin rounded-full border-2 border-white border-t-transparent"></div>
            {{ isLoading ? 'Saving...' : 'Save Settings' }}
          </div>
        </button>
      </div>
    </form>

    <!-- Success Message -->
    <div
      v-if="showSuccess"
      class="fixed bottom-4 right-4 z-50 transition-all duration-300"
    >
      <div class="relative">
        <div class="absolute inset-0 bg-gradient-to-r from-green-600/30 to-emerald-600/30 rounded-2xl blur-xl"></div>
        <div class="relative backdrop-blur-xl border px-6 py-3 rounded-2xl"
             :style="{
               backgroundColor: isDark ? 'rgba(34, 197, 94, 0.2)' : 'rgba(34, 197, 94, 0.1)',
               borderColor: 'rgba(34, 197, 94, 0.3)',
               color: '#22c55e'
             }">
          <div class="flex items-center">
            <CheckCircleIcon class="w-5 h-5 mr-2" />
            <span class="font-medium">Settings saved successfully!</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Gradient Configuration Dialog -->
    <GradientDialog
      :is-open="showGradientDialog"
      title="Configure UI Gradient Colors"
      :initial-colors="{
        color1: settings.gradientColor1,
        color2: settings.gradientColor2,
        color3: settings.gradientColor3
      }"
      @close="showGradientDialog = false"
      @apply="applyGradientColors"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useAuth } from '../../../composables/useAuth'
import { useTheme } from '../../../composables/useTheme'
import {
  ClockIcon,
  CloudArrowUpIcon,
  PhotoIcon,
  ScaleIcon,
  CheckCircleIcon,
  UserGroupIcon,
  SwatchIcon
} from '@heroicons/vue/24/outline'
import ColorPicker from '../../../components/ColorPicker.vue'
import GradientDialog from '../../../components/GradientDialog.vue'

const { isDark } = useTheme()

const { saveAdminSettings, getSettings, isLoading } = useAuth()

const settings = ref({
  logo: '',
  backgroundImage: '',
  navbarTitle: 'PinGO',
  maxUploadSize: 104857600,
  blurIntensity: 0,
  maxValidity: '7days',
  allowRegistration: true,
  expirationAction: 'unavailable',
  websiteColor: '#3b82f6',
  gradientColor1: '#3b82f6',
  gradientColor2: '#8b5cf6',
  gradientColor3: '#ec4899'
})

const logoFile = ref<File | null>(null)
const backgroundFile = ref<File | null>(null)
const showSuccess = ref(false)
const showGradientDialog = ref(false)

// Computed property for current gradient preview
const currentGradient = computed(() => {
  return `linear-gradient(135deg, ${settings.value.gradientColor1} 0%, ${settings.value.gradientColor2} 50%, ${settings.value.gradientColor3} 100%)`
})

const validityOptions = [
  { value: '7days', label: '7 Days', description: 'One week' },
  { value: '1month', label: '1 Month', description: '30 days' },
  { value: '6months', label: '6 Months', description: 'Half year' },
  { value: '1year', label: '1 Year', description: '12 months' },
  { value: 'never', label: 'Never', description: 'Permanent' }
]

const handleLogoUpload = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files && target.files[0]) {
    logoFile.value = target.files[0]
  }
}

const handleBackgroundUpload = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files && target.files[0]) {
    backgroundFile.value = target.files[0]
  }
}

const saveSettings = async () => {
  try {
    const formData = new FormData()
    
    formData.append('navbarTitle', settings.value.navbarTitle)
    formData.append('maxUploadSize', settings.value.maxUploadSize.toString())
    formData.append('blurIntensity', settings.value.blurIntensity.toString())
    formData.append('maxValidity', settings.value.maxValidity)
    formData.append('allowRegistration', settings.value.allowRegistration.toString())
    formData.append('expirationAction', settings.value.expirationAction)
    
    // Append color settings
    formData.append('websiteColor', settings.value.websiteColor)
    formData.append('gradientColor1', settings.value.gradientColor1)
    formData.append('gradientColor2', settings.value.gradientColor2)
    formData.append('gradientColor3', settings.value.gradientColor3)
    
    if (logoFile.value) {
      formData.append('logo', logoFile.value)
    }
    
    if (backgroundFile.value) {
      formData.append('backgroundImage', backgroundFile.value)
    }
    
    const result = await saveAdminSettings(formData)
    
    // Update local settings with response
    settings.value = { ...settings.value, ...result }
    
    // Show success message
    showSuccess.value = true
    setTimeout(() => {
      showSuccess.value = false
    }, 3000)
    
    // Dispatch event to notify navbar of settings update
    window.dispatchEvent(new CustomEvent('settings-updated'))
    
    // Reset file inputs
    logoFile.value = null
    backgroundFile.value = null
    
  } catch (error: any) {
    console.error('Failed to save settings:', error)
    alert('Failed to save settings: ' + error.message)
  }
}

const loadSettings = async () => {
  try {
    const currentSettings = await getSettings()
    settings.value = {
      logo: currentSettings.logo || '',
      backgroundImage: currentSettings.backgroundImage || '',
      navbarTitle: currentSettings.navbarTitle || 'PinGO',
      maxUploadSize: currentSettings.maxUploadSize || 104857600,
      blurIntensity: currentSettings.blurIntensity || 0,
      maxValidity: currentSettings.maxValidity || '7days',
      allowRegistration: currentSettings.allowRegistration !== undefined ? currentSettings.allowRegistration : true,
      expirationAction: currentSettings.expirationAction || 'unavailable',
      websiteColor: currentSettings.websiteColor || '#3b82f6',
      gradientColor1: currentSettings.gradientColor1 || '#3b82f6',
      gradientColor2: currentSettings.gradientColor2 || '#8b5cf6',
      gradientColor3: currentSettings.gradientColor3 || '#ec4899'
    }
  } catch (error) {
    console.error('Failed to load settings:', error)
  }
}

// Apply gradient colors from dialog
const applyGradientColors = (colors: { color1: string; color2: string; color3: string }) => {
  settings.value.gradientColor1 = colors.color1
  settings.value.gradientColor2 = colors.color2
  settings.value.gradientColor3 = colors.color3
}

onMounted(() => {
  loadSettings()
})
</script>

<style scoped>
.slider::-webkit-slider-thumb {
  appearance: none;
  height: 20px;
  width: 20px;
  border-radius: 50%;
  background: #3b82f6;
  cursor: pointer;
  border: 2px solid #ffffff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

.slider::-moz-range-thumb {
  height: 20px;
  width: 20px;
  border-radius: 50%;
  background: #3b82f6;
  cursor: pointer;
  border: 2px solid #ffffff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

/* Dark mode slider track */
[data-theme="dark"] .slider::-webkit-slider-track {
  background: #4b5563;
}

[data-theme="dark"] .slider::-moz-range-track {
  background: #4b5563;
}

/* Dropdown/select styling for dark mode compatibility */
select option {
  background-color: inherit;
  color: inherit;
}
</style>
