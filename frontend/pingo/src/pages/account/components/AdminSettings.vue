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

      <!-- File Expiration Behavior Setting -->
      <div class="rounded-xl p-6 border transition-colors duration-300"
           :style="{
             background: isDark 
               ? 'linear-gradient(135deg, #7c2d12 0%, #9a3412 100%)' 
               : 'linear-gradient(135deg, #fff7ed 0%, #fed7aa 100%)',
             borderColor: isDark ? '#ea580c' : '#fb923c'
           }">
        <h3 class="text-lg font-semibold mb-4 flex items-center transition-colors duration-300"
            :style="{ color: isDark ? '#fff7ed' : '#111827' }">
          <ClockIcon class="w-5 h-5 mr-2"
                     :style="{ color: isDark ? '#fb923c' : '#ea580c' }" />
          File Expiration Behavior
        </h3>
        <p class="text-sm mb-4 transition-colors duration-300"
           :style="{ color: isDark ? '#fed7aa' : '#6b7280' }">
          Choose what happens to files when they expire
        </p>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <label
            :class="[
              'relative flex flex-col p-4 rounded-lg border-2 cursor-pointer transition-all',
              settings.expirationAction === 'unavailable'
                ? (isDark ? 'border-orange-400 bg-orange-600 text-white' : 'border-orange-500 bg-orange-500 text-white')
                : (isDark 
                    ? 'border-gray-600 bg-gray-700 text-gray-200 hover:border-orange-400' 
                    : 'border-gray-200 bg-white text-gray-700 hover:border-orange-300'
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
              'relative flex flex-col p-4 rounded-lg border-2 cursor-pointer transition-all',
              settings.expirationAction === 'delete'
                ? (isDark ? 'border-orange-400 bg-orange-600 text-white' : 'border-orange-500 bg-orange-500 text-white')
                : (isDark 
                    ? 'border-gray-600 bg-gray-700 text-gray-200 hover:border-orange-400' 
                    : 'border-gray-200 bg-white text-gray-700 hover:border-orange-300'
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
        <div class="mt-4 p-3 rounded-lg border transition-colors duration-300"
             :style="{
               backgroundColor: isDark ? 'rgba(251, 146, 60, 0.1)' : 'rgba(234, 88, 12, 0.1)',
               borderColor: isDark ? '#fb923c' : '#ea580c'
             }">
          <p class="text-sm transition-colors duration-300"
             :style="{ color: isDark ? '#fed7aa' : '#9a3412' }">
            <strong>Note:</strong> This setting affects how the system handles files after they expire.
            The cleanup process runs every hour.
          </p>
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
        <div class="max-w-md">
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
            <p class="text-sm mt-1 transition-colors duration-300"
               :style="{ color: isDark ? '#d1d5db' : '#4b5563' }">
              Set the maximum file size users can upload. Choose "Unlimited" to allow any size.
            </p>
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
  CloudArrowUpIcon,
  PhotoIcon,
  ScaleIcon,
  CheckCircleIcon,
  UserGroupIcon
} from '@heroicons/vue/24/outline'

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
  expirationAction: 'unavailable'
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
    
    formData.append('navbarTitle', settings.value.navbarTitle)
    formData.append('maxUploadSize', settings.value.maxUploadSize.toString())
    formData.append('blurIntensity', settings.value.blurIntensity.toString())
    formData.append('maxValidity', settings.value.maxValidity)
    formData.append('allowRegistration', settings.value.allowRegistration.toString())
    formData.append('expirationAction', settings.value.expirationAction)
    
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
      expirationAction: currentSettings.expirationAction || 'unavailable'
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
