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
      <div class="rounded-xl p-6 border transition-colors duration-300"
           :style="{
             background: isDark 
               ? 'linear-gradient(135deg, #581c87 0%, #7c2d92 100%)' 
               : 'linear-gradient(135deg, #fdf4ff 0%, #fae8ff 100%)',
             borderColor: isDark ? '#7c3aed' : '#d8b4fe'
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
              'relative flex items-center justify-center px-4 py-3 rounded-lg border-2 cursor-pointer transition-all',
              settings.maxValidity === option.value
                ? (isDark ? 'border-purple-400 bg-purple-600 text-white' : 'border-purple-500 bg-purple-500 text-white')
                : (isDark 
                    ? 'border-gray-600 bg-gray-700 text-gray-200 hover:border-purple-400' 
                    : 'border-gray-200 bg-white text-gray-700 hover:border-purple-300'
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
        <p class="text-sm mt-3 transition-colors duration-300"
           :style="{ color: isDark ? '#d1d5db' : '#4b5563' }">
          Users will not be able to set file expiration longer than this maximum period.
        </p>
      </div>

      <!-- Theme Setting -->
      <div class="rounded-xl p-6 border transition-colors duration-300"
           :style="{
             background: isDark 
               ? 'linear-gradient(135deg, #1e3a8a 0%, #1e40af 100%)' 
               : 'linear-gradient(135deg, #dbeafe 0%, #bfdbfe 100%)',
             borderColor: isDark ? '#3b82f6' : '#93c5fd'
           }">
        <h3 class="text-lg font-semibold mb-4 flex items-center transition-colors duration-300"
            :style="{ color: isDark ? '#dbeafe' : '#111827' }">
          <SwatchIcon class="w-5 h-5 mr-2"
                      :style="{ color: isDark ? '#60a5fa' : '#2563eb' }" />
          Theme
        </h3>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <label :class="[
            'relative flex items-center p-4 rounded-lg border-2 cursor-pointer transition-all',
            settings.theme === 'light'
              ? (isDark ? 'border-blue-400 bg-blue-600 text-white' : 'border-blue-500 bg-blue-500 text-white')
              : (isDark 
                  ? 'border-gray-600 bg-gray-700 text-gray-200 hover:border-blue-400' 
                  : 'border-gray-200 bg-white text-gray-700 hover:border-blue-300'
                )
          ]">
            <input
              type="radio"
              v-model="settings.theme"
              value="light"
              class="sr-only"
            />
            <SunIcon class="w-5 h-5 mr-3" />
            <span class="font-medium">Light Theme</span>
          </label>
          <label :class="[
            'relative flex items-center p-4 rounded-lg border-2 cursor-pointer transition-all',
            settings.theme === 'dark'
              ? (isDark ? 'border-blue-400 bg-blue-600 text-white' : 'border-blue-500 bg-blue-500 text-white')
              : (isDark 
                  ? 'border-gray-600 bg-gray-700 text-gray-200 hover:border-blue-400' 
                  : 'border-gray-200 bg-white text-gray-700 hover:border-blue-300'
                )
          ]">
            <input
              type="radio"
              v-model="settings.theme"
              value="dark"
              class="sr-only"
            />
            <MoonIcon class="w-5 h-5 mr-3" />
            <span class="font-medium">Dark Theme</span>
          </label>
        </div>
      </div>

      <!-- User Registration Setting -->
      <div class="rounded-xl p-6 border transition-colors duration-300"
           :style="{
             background: isDark 
               ? 'linear-gradient(135deg, #92400e 0%, #b45309 100%)' 
               : 'linear-gradient(135deg, #fef3c7 0%, #fde68a 100%)',
             borderColor: isDark ? '#f59e0b' : '#fbbf24'
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
            <p class="text-sm transition-colors duration-300"
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
                'w-11 h-6 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-orange-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[\'\'] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-orange-600',
                isDark ? 'bg-gray-600 after:border-gray-400' : 'bg-gray-200 after:border-gray-300'
              ]"
            ></div>
          </label>
        </div>
      </div>

      <!-- Upload Settings -->
      <div class="rounded-xl p-6 border transition-colors duration-300"
           :style="{
             background: isDark 
               ? 'linear-gradient(135deg, #065f46 0%, #047857 100%)' 
               : 'linear-gradient(135deg, #d1fae5 0%, #a7f3d0 100%)',
             borderColor: isDark ? '#10b981' : '#6ee7b7'
           }">
        <h3 class="text-lg font-semibold mb-4 flex items-center transition-colors duration-300"
            :style="{ color: isDark ? '#d1fae5' : '#111827' }">
          <CloudArrowUpIcon class="w-5 h-5 mr-2"
                            :style="{ color: isDark ? '#34d399' : '#059669' }" />
          Upload Settings
        </h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium mb-2 transition-colors duration-300"
                   :style="{ color: isDark ? '#d1fae5' : '#374151' }">
              Maximum Upload Size
            </label>
            <select
              v-model="settings.maxUploadSize"
              class="w-full px-3 py-2 border rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500 transition-colors duration-300"
              :style="{
                backgroundColor: isDark ? '#374151' : '#ffffff',
                borderColor: isDark ? '#6b7280' : '#d1d5db',
                color: isDark ? '#f9fafb' : '#111827'
              }"
            >
              <option value="10485760">10 MB</option>
              <option value="52428800">50 MB</option>
              <option value="104857600">100 MB</option>
              <option value="524288000">500 MB</option>
              <option value="1073741824">1 GB</option>
              <option value="2147483648">2 GB</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium mb-2 transition-colors duration-300"
                   :style="{ color: isDark ? '#d1fae5' : '#374151' }">
              Upload Box Transparency ({{ settings.uploadBoxTransparency }}%)
            </label>
            <input
              type="range"
              v-model="settings.uploadBoxTransparency"
              min="0"
              max="100"
              class="w-full h-2 rounded-lg appearance-none cursor-pointer slider transition-colors duration-300"
              :style="{
                backgroundColor: isDark ? '#6b7280' : '#e5e7eb'
              }"
            />
          </div>
        </div>
      </div>

      <!-- Branding Settings -->
      <div class="rounded-xl p-6 border transition-colors duration-300"
           :style="{
             background: isDark 
               ? 'linear-gradient(135deg, #991b1b 0%, #dc2626 100%)' 
               : 'linear-gradient(135deg, #fef2f2 0%, #fee2e2 100%)',
             borderColor: isDark ? '#ef4444' : '#fca5a5'
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
              class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-orange-500 transition-colors duration-300"
              :style="{
                backgroundColor: isDark ? '#374151' : '#ffffff',
                borderColor: isDark ? '#6b7280' : '#d1d5db',
                color: isDark ? '#f9fafb' : '#111827'
              }"
            />
            <p class="text-sm mt-1 transition-colors duration-300"
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
                  class="w-12 h-12 object-contain rounded-lg border"
                  :style="{ borderColor: isDark ? '#6b7280' : '#d1d5db' }"
                />
                <input
                  type="file"
                  ref="logoInput"
                  @change="handleLogoUpload"
                  accept="image/*"
                  class="block w-full text-sm file:mr-4 file:py-2 file:px-4 file:rounded-lg file:border-0 file:bg-orange-500 file:text-white hover:file:bg-orange-600 transition-colors duration-300"
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
                  class="w-12 h-12 object-cover rounded-lg border"
                  :style="{ borderColor: isDark ? '#6b7280' : '#d1d5db' }"
                />
                <input
                  type="file"
                  ref="backgroundInput"
                  @change="handleBackgroundUpload"
                  accept="image/*"
                  class="block w-full text-sm file:mr-4 file:py-2 file:px-4 file:rounded-lg file:border-0 file:bg-orange-500 file:text-white hover:file:bg-orange-600 transition-colors duration-300"
                  :style="{ color: isDark ? '#d1d5db' : '#6b7280' }"
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
          class="inline-flex items-center px-6 py-3 bg-gradient-to-r from-blue-600 to-indigo-600 hover:from-blue-700 hover:to-indigo-700 text-white font-medium rounded-lg transition-all disabled:opacity-50 disabled:cursor-not-allowed shadow-lg hover:shadow-xl"
        >
          <ScaleIcon v-if="!isLoading" class="w-5 h-5 mr-2" />
          <div v-else class="w-5 h-5 mr-2 animate-spin rounded-full border-2 border-white border-t-transparent"></div>
          {{ isLoading ? 'Saving...' : 'Save Settings' }}
        </button>
      </div>
    </form>

    <!-- Success Message -->
    <div
      v-if="showSuccess"
      class="fixed bottom-4 right-4 bg-green-500 text-white px-6 py-3 rounded-lg shadow-lg z-50 transition-all"
    >
      <div class="flex items-center">
        <CheckCircleIcon class="w-5 h-5 mr-2" />
        Settings saved successfully!
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useAuth } from '../../../composables/useAuth'
import { useTheme } from '../../../composables/useTheme'
import {
  ClockIcon,
  SwatchIcon,
  CloudArrowUpIcon,
  PhotoIcon,
  ScaleIcon,
  CheckCircleIcon,
  SunIcon,
  MoonIcon,
  UserGroupIcon
} from '@heroicons/vue/24/outline'

const { isDark } = useTheme()

const { saveAdminSettings, getSettings, isLoading } = useAuth()

const settings = ref({
  theme: 'light',
  logo: '',
  backgroundImage: '',
  navbarTitle: 'PinGO',
  maxUploadSize: 104857600,
  uploadBoxTransparency: 0,
  blurIntensity: 0,
  maxValidity: '7days',
  allowRegistration: true
})

const logoFile = ref<File | null>(null)
const backgroundFile = ref<File | null>(null)
const showSuccess = ref(false)

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
    
    formData.append('theme', settings.value.theme)
    formData.append('navbarTitle', settings.value.navbarTitle)
    formData.append('maxUploadSize', settings.value.maxUploadSize.toString())
    formData.append('uploadBoxTransparency', settings.value.uploadBoxTransparency.toString())
    formData.append('blurIntensity', settings.value.blurIntensity.toString())
    formData.append('maxValidity', settings.value.maxValidity)
    formData.append('allowRegistration', settings.value.allowRegistration.toString())
    
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
      theme: currentSettings.theme || 'light',
      logo: currentSettings.logo || '',
      backgroundImage: currentSettings.backgroundImage || '',
      navbarTitle: currentSettings.navbarTitle || 'PinGO',
      maxUploadSize: currentSettings.maxUploadSize || 104857600,
      uploadBoxTransparency: currentSettings.uploadBoxTransparency || 0,
      blurIntensity: currentSettings.blurIntensity || 0,
      maxValidity: currentSettings.maxValidity || '7days',
      allowRegistration: currentSettings.allowRegistration !== undefined ? currentSettings.allowRegistration : true
    }
  } catch (error) {
    console.error('Failed to load settings:', error)
  }
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
