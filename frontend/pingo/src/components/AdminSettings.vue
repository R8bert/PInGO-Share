<template>
  <div class="space-y-8">
    <div class="text-center">
      <h2 class="text-2xl font-bold text-gray-900 mb-2">Admin Settings</h2>
      <p class="text-gray-600">Configure global application settings</p>
    </div>

    <form @submit.prevent="saveSettings" class="space-y-8">
      <!-- Maximum Validity Setting -->
      <div class="bg-gradient-to-r from-purple-50 to-pink-50 rounded-xl p-6 border border-purple-200">
        <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
          <ClockIcon class="w-5 h-5 mr-2 text-purple-600" />
          Maximum File Validity
        </h3>
        <div class="grid grid-cols-1 sm:grid-cols-5 gap-3">
          <label
            v-for="option in validityOptions"
            :key="option.value"
            :class="[
              'relative flex items-center justify-center px-4 py-3 rounded-lg border-2 cursor-pointer transition-all',
              settings.maxValidity === option.value
                ? 'border-purple-500 bg-purple-500 text-white'
                : 'border-gray-200 bg-white text-gray-700 hover:border-purple-300'
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
        <p class="text-sm text-gray-600 mt-3">
          Users will not be able to set file expiration longer than this maximum period.
        </p>
      </div>

      <!-- Theme Setting -->
      <div class="bg-gradient-to-r from-blue-50 to-indigo-50 rounded-xl p-6 border border-blue-200">
        <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
          <SwatchIcon class="w-5 h-5 mr-2 text-blue-600" />
          Theme
        </h3>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <label :class="[
            'relative flex items-center p-4 rounded-lg border-2 cursor-pointer transition-all',
            settings.theme === 'light'
              ? 'border-blue-500 bg-blue-500 text-white'
              : 'border-gray-200 bg-white text-gray-700 hover:border-blue-300'
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
              ? 'border-blue-500 bg-blue-500 text-white'
              : 'border-gray-200 bg-white text-gray-700 hover:border-blue-300'
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
      <div class="bg-gradient-to-r from-orange-50 to-yellow-50 rounded-xl p-6 border border-orange-200">
        <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
          <UserGroupIcon class="w-5 h-5 mr-2 text-orange-600" />
          User Registration
        </h3>
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-gray-900">Allow New User Registration</p>
            <p class="text-sm text-gray-600">When disabled, only existing users and admins can log in</p>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input
              type="checkbox"
              v-model="settings.allowRegistration"
              class="sr-only peer"
            />
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-orange-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-orange-600"></div>
          </label>
        </div>
      </div>

      <!-- Upload Settings -->
      <div class="bg-gradient-to-r from-green-50 to-emerald-50 rounded-xl p-6 border border-green-200">
        <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
          <CloudArrowUpIcon class="w-5 h-5 mr-2 text-green-600" />
          Upload Settings
        </h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Maximum Upload Size
            </label>
            <select
              v-model="settings.maxUploadSize"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500"
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
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Upload Box Transparency ({{ settings.uploadBoxTransparency }}%)
            </label>
            <input
              type="range"
              v-model="settings.uploadBoxTransparency"
              min="0"
              max="100"
              class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer slider"
            />
          </div>
        </div>
      </div>

      <!-- Logo Upload -->
      <div class="bg-gradient-to-r from-orange-50 to-red-50 rounded-xl p-6 border border-orange-200">
        <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
          <PhotoIcon class="w-5 h-5 mr-2 text-orange-600" />
          Branding
        </h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Logo
            </label>
            <div class="flex items-center space-x-4">
              <img
                v-if="settings.logo"
                :src="`http://localhost:8080${settings.logo}`"
                alt="Current logo"
                class="w-12 h-12 object-contain rounded-lg border border-gray-200"
              />
              <input
                type="file"
                ref="logoInput"
                @change="handleLogoUpload"
                accept="image/*"
                class="block w-full text-sm text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded-lg file:border-0 file:bg-orange-500 file:text-white hover:file:bg-orange-600"
              />
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Background Image
            </label>
            <div class="flex items-center space-x-4">
              <img
                v-if="settings.backgroundImage"
                :src="`http://localhost:8080${settings.backgroundImage}`"
                alt="Current background"
                class="w-12 h-12 object-cover rounded-lg border border-gray-200"
              />
              <input
                type="file"
                ref="backgroundInput"
                @change="handleBackgroundUpload"
                accept="image/*"
                class="block w-full text-sm text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded-lg file:border-0 file:bg-orange-500 file:text-white hover:file:bg-orange-600"
              />
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
import { useAuth } from '../composables/useAuth'
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

const { saveAdminSettings, getSettings, isLoading } = useAuth()

const settings = ref({
  theme: 'light',
  logo: '',
  backgroundImage: '',
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
}

.slider::-moz-range-thumb {
  height: 20px;
  width: 20px;
  border-radius: 50%;
  background: #3b82f6;
  cursor: pointer;
  border: none;
}
</style>
