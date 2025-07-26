<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-indigo-50 via-purple-50 to-blue-50 px-4 py-8">
    <div class="w-full max-w-2xl bg-white shadow-2xl rounded-xl p-6">
      <!-- Titlu și tab-uri -->
      <div class="mb-6">
        <h1 class="text-2xl font-bold text-center text-gray-900 tracking-tight mb-4">Settings</h1>
        <div class="border-b border-gray-200">
          <nav class="flex space-x-6">
            <button
              v-for="tab in tabs"
              :key="tab"
              @click="activeTab = tab"
              :class="['text-sm font-medium px-4 py-2 rounded-t-md', activeTab === tab ? 'text-indigo-600 border-b-2 border-indigo-600' : 'text-gray-500 hover:text-gray-700']"
            >
              {{ tab }}
            </button>
          </nav>
        </div>
      </div>

      <!-- Conținut tab-uri -->
      <div class="space-y-6">
        <!-- Tab Branding & UI -->
        <div v-if="activeTab === 'Branding & UI'">
          <h2 class="text-sm font-semibold text-gray-700 border-b border-gray-200 pb-2">Branding & UI</h2>
          <div class="space-y-4 mt-4">
            <div>
              <label for="theme" class="block text-sm font-medium text-gray-700">Theme</label>
              <select
                id="theme"
                v-model="settings.theme"
                class="mt-1 w-full border border-gray-300 rounded-md px-3 py-2 text-sm text-gray-700 bg-white focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500"
              >
                <option value="light">Light</option>
                <option value="dark">Dark</option>
                <option value="system">System Default</option>
              </select>
            </div>
            <div>
              <label for="logo" class="block text-sm font-medium text-gray-700">Logo</label>
              <input
                id="logo"
                type="file"
                accept="image/*"
                @change="handleLogoUpload"
                class="mt-1 w-full border border-gray-300 rounded-md px-3 py-2 text-sm text-gray-700 bg-white focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500"
              />
            </div>
            <div>
              <label for="backgroundImage" class="block text-sm font-medium text-gray-700">Background Image</label>
              <input
                id="backgroundImage"
                type="file"
                accept="image/*"
                @change="handleBackgroundUpload"
                class="mt-1 w-full border border-gray-300 rounded-md px-3 py-2 text-sm text-gray-700 bg-white focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500"
              />
            </div>
            <div>
              <label for="maxUploadSize" class="block text-sm font-medium text-gray-700">Max Upload Size</label>
              <select
                id="maxUploadSize"
                v-model="settings.maxUploadSize"
                class="mt-1 w-full border border-gray-300 rounded-md px-3 py-2 text-sm text-gray-700 bg-white focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500"
              >
                <option value="1048576">1 MB</option> <!-- 1 MB = 1024 * 1024 bytes -->
                <option value="10485760">10 MB</option> <!-- 10 MB -->
                <option value="52428800">50 MB</option> <!-- 50 MB -->
                <option value="104857600">100 MB</option> <!-- 100 MB -->
                <option value="1073741824">1 GB</option> <!-- 1 GB -->
              </select>
            </div>
            <button
              @click="saveSettings"
              class="w-full bg-gradient-to-r from-indigo-600 to-purple-600 text-white py-2 rounded-md text-sm font-medium hover:from-indigo-700 hover:to-purple-700"
            >
              Save Settings
            </button>
          </div>
        </div>

        <!-- Tab Users -->
        <div v-if="activeTab === 'Users'">
          <h2 class="text-sm font-semibold text-gray-700 border-b border-gray-200 pb-2">Users</h2>
          <div class="space-y-4 mt-4">
            <div v-for="user in users" :key="user.id" class="flex items-center justify-between p-3 bg-gray-50 rounded-md">
              <span class="text-sm text-gray-700">{{ user.name }}</span>
              <button @click="removeUser(user.id)" class="text-red-600 text-xs hover:text-red-800">Remove</button>
            </div>
            <button
              @click="inviteUser"
              class="w-full bg-gradient-to-r from-indigo-600 to-purple-600 text-white py-2 rounded-md text-sm font-medium hover:from-indigo-700 hover:to-purple-700"
            >
              Invite User
            </button>
          </div>
        </div>

        <!-- Tab My Profile -->
        <div v-if="activeTab === 'My Profile'">
          <h2 class="text-sm font-semibold text-gray-700 border-b border-gray-200 pb-2">My Profile</h2>
          <div class="space-y-4 mt-4">
            <div>
              <label for="email" class="block text-sm font-medium text-gray-700">Email</label>
              <input
                id="email"
                type="email"
                v-model="profile.email"
                class="mt-1 w-full border border-gray-300 rounded-md px-3 py-2 text-sm text-gray-700 bg-white focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500"
                placeholder="your.email@example.com"
              />
            </div>
            <div>
              <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
              <input
                id="password"
                type="password"
                v-model="profile.password"
                class="mt-1 w-full border border-gray-300 rounded-md px-3 py-2 text-sm text-gray-700 bg-white focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500"
                placeholder="••••••••"
              />
            </div>
            <button
              @click="saveProfile"
              class="w-full bg-gradient-to-r from-indigo-600 to-purple-600 text-white py-2 rounded-md text-sm font-medium hover:from-indigo-700 hover:to-purple-700"
            >
              Save Profile
            </button>
          </div>
        </div>
      </div>

      <!-- Mesaj de succes/eroare -->
      <div v-if="message" class="text-center text-sm mt-4 animate-slide-up">
        <p :class="message.type === 'success' ? 'text-green-600' : 'text-red-600'">{{ message.text }}</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import axios from 'axios'

// Tipuri
interface Settings {
  theme: 'light' | 'dark' | 'system'
  logo: File | null
  backgroundImage: File | null
  maxUploadSize: string | number // Allow string from select for v-model
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

// Starea
const tabs = ['Branding & UI', 'Users', 'My Profile']
const activeTab = ref('Branding & UI')
const settings = reactive<Settings>({
  theme: 'light',
  logo: null,
  backgroundImage: null,
  maxUploadSize: '104857600', // Default 100 MB
})
const profile = reactive<Profile>({
  email: '',
  password: '',
})
const users = ref<User[]>([{ id: '1', name: 'User 1' }])
const message = ref<Message | null>(null)

// Funcții
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

const saveSettings = async () => {
  const formData = new FormData()
  formData.append('theme', settings.theme)
  if (settings.logo) formData.append('logo', settings.logo)
  if (settings.backgroundImage) formData.append('backgroundImage', settings.backgroundImage)
  if (settings.maxUploadSize) formData.append('maxUploadSize', settings.maxUploadSize.toString()) // Convert to string for backend

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
  } catch (error) {
    console.error('Error saving settings:', error)
    message.value = { text: 'Failed to save settings.', type: 'error' }
  }
}

const saveProfile = async () => {
  try {
    await axios.post('http://localhost:8080/profile', profile)
    message.value = { text: 'Profile saved successfully!', type: 'success' }
  } catch (error) {
    console.error('Error saving profile:', error)
    message.value = { text: 'Failed to save profile.', type: 'error' }
  }
}

const inviteUser = async () => {
  try {
    const res = await axios.post('http://localhost:8080/users/invite')
    users.value.push({ id: res.data.id, name: res.data.name })
    message.value = { text: 'User invited successfully!', type: 'success' }
  } catch (error) {
    console.error('Error inviting user:', error)
    message.value = { text: 'Failed to invite user.', type: 'error' }
  }
}

const removeUser = (id: string) => {
  users.value = users.value.filter(user => user.id !== id)
}

const loadSettings = async () => {
  try {
    const response = await axios.get('http://localhost:8080/settings')
    console.log('Loaded settings:', response.data)
    settings.theme = response.data.theme || 'light'
    settings.maxUploadSize = response.data.maxUploadSize || '104857600' // Default 100 MB if not set
  } catch (error) {
    console.error('Error loading settings:', error)
  }
}

onMounted(() => {
  loadSettings()
})
</script>

<style scoped>
.animate-slide-up {
  animation: slideUp 0.3s ease-out;
}

@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>