import { ref, onMounted, onUnmounted } from 'vue'
import axios from 'axios'
import { getApiUrl } from '../utils/apiUtils'

interface Settings {
  id: number
  theme: string
  logo: string
  backgroundImage: string
  navbarTitle: string
  maxUploadSize: number
  uploadBoxTransparency: number
  blurIntensity: number
  maxValidity: string
  allowRegistration: boolean
  websiteColor: string
  gradientColor1: string
  gradientColor2: string
  gradientColor3: string
}

export function useSettings() {
  const settings = ref<Settings | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  const fetchSettings = async () => {
    loading.value = true
    error.value = null
    
    try {
      const response = await axios.get(getApiUrl('/settings'))
      settings.value = response.data
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to fetch settings'
      console.error('Error fetching settings:', err)
    } finally {
      loading.value = false
    }
  }

  // Listen for settings updates
  const handleSettingsUpdate = () => {
    fetchSettings()
  }

  onMounted(() => {
    fetchSettings()
    window.addEventListener('settings-updated', handleSettingsUpdate)
  })

  onUnmounted(() => {
    window.removeEventListener('settings-updated', handleSettingsUpdate)
  })

  return {
    settings,
    loading,
    error,
    fetchSettings
  }
}
