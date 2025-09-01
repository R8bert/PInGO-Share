<template>
  <div class="space-y-8">
    <div class="text-center">
      <h2 class="text-2xl font-bold mb-2 text-gray-900 dark:text-gray-100">Admin Settings</h2>
      <p class="text-gray-600 dark:text-gray-400">Configure global application settings</p>
    </div>  <div class="space-y-8">
    <div class="text-center">
      <h2 class="text-2xl font-bold mb-2 transition-colors duration-300"
          :style="{ color: isDark ? '#f9fafb' : '#111827' }">Admin Settings</h2>
      <p class="transition-colors duration-300"
         :style="{ color: isDark ? '#9            <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>f' : '#4b5563' }">Configure global application settings</p>
    </div>

    <form @submit.prevent="saveSettings" class="space-y-6">
      <!-- Maximum Validity Setting -->
      <div class="bg-white/80 dark:bg-gray-800/80 backdrop-blur-sm border border-gray-200 dark:border-gray-700 rounded-xl p-6 transition-all duration-300">
        <h3 class="text-lg font-semibold mb-4 flex items-center text-gray-900 dark:text-gray-100">
          <ClockIcon class="w-5 h-5 mr-2 text-gray-600 dark:text-gray-400" />
          Maximum File Validity
        </h3>
        <div class="grid grid-cols-1 sm:grid-cols-5 gap-3">
          <label
            v-for="option in validityOptions"
            :key="option.value"
            :class="[
              'relative flex items-center justify-center px-4 py-3 rounded-lg border cursor-pointer transition-all',
              settings.maxValidity === option.value
                ? 'border-blue-500 bg-blue-50 dark:bg-blue-900/30 text-blue-900 dark:text-blue-100'
                : 'border-gray-200 dark:border-gray-600 bg-white dark:bg-gray-700/50 text-gray-700 dark:text-gray-300 hover:border-gray-300 dark:hover:border-gray-500'
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
        <p class="text-sm mt-3 text-gray-600 dark:text-gray-400">
          Users will not be able to set file expiration longer than this maximum period.
        </p>
      </div>

      <!-- User Registration Setting -->
      <div class="bg-white/80 dark:bg-gray-800/80 backdrop-blur-sm border border-gray-200 dark:border-gray-700 rounded-xl p-6 transition-all duration-300">
        <h3 class="text-lg font-semibold mb-4 flex items-center text-gray-900 dark:text-gray-100">
          <UserGroupIcon class="w-5 h-5 mr-2 text-gray-600 dark:text-gray-400" />
          User Registration
        </h3>
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-gray-900 dark:text-gray-100">Allow New User Registration</p>
            <p class="text-sm text-gray-600 dark:text-gray-400">When disabled, only existing users and admins can log in</p>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input
              type="checkbox"
              v-model="settings.allowRegistration"
              class="sr-only peer"
            />
            <div 
              class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 rounded-full peer dark:bg-gray-600 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600"
            ></div>
          </label>
        </div>
      </div>

      <!-- File Expiration Behavior Setting -->
      <div class="bg-white/80 dark:bg-gray-800/80 backdrop-blur-sm border border-gray-200 dark:border-gray-700 rounded-xl p-6 transition-all duration-300">
        <h3 class="text-lg font-semibold mb-4 flex items-center text-gray-900 dark:text-gray-100">
          <ClockIcon class="w-5 h-5 mr-2 text-gray-600 dark:text-gray-400" />
          File Expiration Behavior
        </h3>
        <p class="text-sm mb-4 text-gray-600 dark:text-gray-400">
          Choose what happens to files when they expire
        </p>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <label
            :class="[
              'relative flex flex-col p-4 rounded-lg border cursor-pointer transition-all',
              settings.expirationAction === 'unavailable'
                ? 'border-blue-500 bg-blue-50 dark:bg-blue-900/30 text-blue-900 dark:text-blue-100'
                : 'border-gray-200 dark:border-gray-600 bg-white dark:bg-gray-700/50 text-gray-700 dark:text-gray-300 hover:border-gray-300 dark:hover:border-gray-500'
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
            <p class="text-sm opacity-75">Files become inaccessible but remain on the server</p>
          </label>
          
          <label
            :class="[
              'relative flex flex-col p-4 rounded-lg border cursor-pointer transition-all',
              settings.expirationAction === 'delete'
                ? 'border-blue-500 bg-blue-50 dark:bg-blue-900/30 text-blue-900 dark:text-blue-100'
                : 'border-gray-200 dark:border-gray-600 bg-white dark:bg-gray-700/50 text-gray-700 dark:text-gray-300 hover:border-gray-300 dark:hover:border-gray-500'
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
            <p class="text-sm opacity-75">Files are permanently removed from the server</p>
          </label>
        </div>
        <div class="mt-4 p-3 rounded-lg border border-amber-200 dark:border-amber-800 bg-amber-50 dark:bg-amber-900/20">
          <p class="text-sm text-amber-800 dark:text-amber-200">
            <strong>Note:</strong> This setting affects how the system handles files after they expire.
            The cleanup process runs every hour.
          </p>
        </div>
      </div>

      <!-- Upload Settings -->
      <div class="bg-white/80 dark:bg-gray-800/80 backdrop-blur-sm border border-gray-200 dark:border-gray-700 rounded-xl p-6 transition-all duration-300">
        <h3 class="text-lg font-semibold mb-4 flex items-center text-gray-900 dark:text-gray-100">
          <CloudArrowUpIcon class="w-5 h-5 mr-2 text-gray-600 dark:text-gray-400" />
          Upload Settings
        </h3>
        <div class="max-w-md">
          <div>
            <label class="block text-sm font-medium mb-2 text-gray-700 dark:text-gray-300">
              Maximum Upload Size
            </label>
            <select
              v-model="settings.maxUploadSize"
              class="w-full px-3 py-2 border border-gray-200 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all duration-300"
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
            <p class="text-sm mt-1 text-gray-600 dark:text-gray-400">
              Set the maximum file size users can upload. Choose "Unlimited" to allow any size.
            </p>
          </div>
        </div>
      </div>

      <!-- Branding Settings -->
      <div class="bg-white/80 dark:bg-gray-800/80 backdrop-blur-sm border border-gray-200 dark:border-gray-700 rounded-xl p-6 transition-all duration-300">
        <h3 class="text-lg font-semibold mb-4 flex items-center text-gray-900 dark:text-gray-100">
          <PhotoIcon class="w-5 h-5 mr-2 text-gray-600 dark:text-gray-400" />
          Branding
        </h3>
        <div class="space-y-6">
          <!-- Application Title -->
          <div>
            <label class="block text-sm font-medium mb-2 text-gray-700 dark:text-gray-300">
              Application Title
            </label>
            <input
              type="text"
              v-model="settings.navbarTitle"
              placeholder="PinGO"
              class="w-full px-4 py-2 border border-gray-200 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all duration-300"
            />
            <p class="text-sm mt-1 text-gray-600 dark:text-gray-400">
              This title will appear next to the logo in the navigation bar.
            </p>
          </div>
          
          <!-- Logo and Background Images -->
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium mb-2 text-gray-700 dark:text-gray-300">
                Logo
              </label>
              <div class="flex items-center space-x-4">
                <img
                  v-if="settings.logo"
                  :src="`http://localhost:8080${settings.logo}`"
                  alt="Current logo"
                  class="w-12 h-12 object-contain rounded-lg border border-gray-200 dark:border-gray-600"
                />
                <input
                  type="file"
                  ref="logoInput"
                  @change="handleLogoUpload"
                  accept="image/*"
                  class="block w-full text-sm file:mr-4 file:py-2 file:px-4 file:rounded-lg file:border-0 file:bg-blue-600 file:text-white hover:file:bg-blue-700 transition-all duration-300 text-gray-600 dark:text-gray-400"
                />
              </div>
            </div>
            <div>
              <label class="block text-sm font-medium mb-2 text-gray-700 dark:text-gray-300">
                Background Image
              </label>
              <div class="flex items-center space-x-4">
                <img
                  v-if="settings.backgroundImage"
                  :src="`http://localhost:8080${settings.backgroundImage}`"
                  alt="Current background"
                  class="w-12 h-12 object-cover rounded-lg border border-gray-200 dark:border-gray-600"
                />
                <input
                  type="file"
                  ref="backgroundInput"
                  @change="handleBackgroundUpload"
                  accept="image/*"
                  class="block w-full text-sm file:mr-4 file:py-2 file:px-4 file:rounded-lg file:border-0 file:bg-blue-600 file:text-white hover:file:bg-blue-700 transition-all duration-300 text-gray-600 dark:text-gray-400"
                />
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Color Customization Settings -->
      <div class="bg-white/80 dark:bg-gray-800/80 backdrop-blur-sm border border-gray-200 dark:border-gray-700 rounded-xl p-6 transition-all duration-300">
        <h3 class="text-lg font-semibold mb-4 flex items-center text-gray-900 dark:text-gray-100">
          <SwatchIcon class="w-5 h-5 mr-2 text-gray-600 dark:text-gray-400" />
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
            <p class="text-sm mt-2 text-gray-600 dark:text-gray-400">
              Primary color used for the WebGL background animation on home and download pages.
            </p>
          </div>

          <!-- UI Gradient Colors -->
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <div>
                <h4 class="font-medium text-gray-900 dark:text-gray-100">
                  UI Gradient Colors
                </h4>
                <p class="text-sm text-gray-600 dark:text-gray-400">
                  Colors used for buttons and UI elements gradients.
                </p>
              </div>
              <button
                @click="showGradientDialog = true"
                type="button"
                class="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-colors duration-300"
              >
                Configure Colors
              </button>
            </div>
            
            <!-- Color Preview -->
            <div class="flex items-center space-x-2 flex-wrap gap-2">
              <div
                class="w-8 h-8 rounded-lg border border-gray-200 dark:border-gray-600"
                :style="{ backgroundColor: settings.gradientColor1 }"
                :title="settings.gradientColor1"
              ></div>
              <div
                class="w-8 h-8 rounded-lg border border-gray-200 dark:border-gray-600"
                :style="{ backgroundColor: settings.gradientColor2 }"
                :title="settings.gradientColor2"
              ></div>
              <div
                class="w-8 h-8 rounded-lg border border-gray-200 dark:border-gray-600"
                :style="{ backgroundColor: settings.gradientColor3 }"
                :title="settings.gradientColor3"
              ></div>
            </div>
          </div>
        </div>
      </div>

      <!-- Save Button -->
      <div class="flex justify-end pt-6">
        <button
          type="submit"
          :disabled="isLoading"
          class="px-6 py-3 bg-blue-600 hover:bg-blue-700 disabled:opacity-50 text-white font-medium rounded-lg transition-colors duration-300"
        >
          <span v-if="isLoading" class="flex items-center">
            <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 714 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            Saving...
          </span>
          <span v-else>Save Settings</span>
        </button>
      </div>
    </form>
          class="px-6 py-3 bg-blue-600 hover:bg-blue-700 disabled:opacity-50 text-white font-medium rounded-lg transition-colors duration-300"
        >
          <span v-if="isLoading" class="flex items-center">
            <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            Saving...
          </span>
          <span v-else>Save Settings</span>
        </button>
      </div>
    </form>

    <!-- Success Message -->
    <div
      v-if="showSuccess"
      class="fixed bottom-4 right-4 z-50 transition-all duration-300"
    >
      <div class="bg-green-100 dark:bg-green-900/20 border border-green-200 dark:border-green-800 px-6 py-3 rounded-lg shadow-lg">
        <div class="flex items-center text-green-800 dark:text-green-200">
          <CheckCircleIcon class="w-5 h-5 mr-2" />
          <span class="font-medium">Settings saved successfully!</span>
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
import { ref, onMounted } from 'vue'
import { useAuth } from '../../../composables/useAuth'
import { useTheme } from '../../../composables/useTheme'
import {
  ClockIcon,
  CloudArrowUpIcon,
  PhotoIcon,
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
