<template>
  <div>
    <!-- Logo positioned in top-left corner -->
    <div class="fixed top-4 left-4 z-50">
      <router-link to="/" class="flex items-center space-x-2">
        <img v-if="logoPath" :src="logoPath ? `http://localhost:8080${logoPath}` : ''" class="h-8 w-8" alt="Logo" @error="handleImageError" />
        <CloudArrowUpIcon v-else class="h-8 w-8 text-blue-600" />
        <span class="text-2xl font-extrabold text-gray-900">PinGO File Transfer</span>
      </router-link>
    </div>

    <!-- User info in top-right corner for authenticated users -->
    <div v-if="isAuthenticated" class="fixed top-4 right-4 z-50 hidden sm:block">
      <div class="flex items-center space-x-3 bg-white shadow-lg rounded-full px-4 py-2">
        <div class="flex items-center space-x-2">
          <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
            <UserIcon class="w-5 h-5 text-blue-600" />
          </div>
          <span class="text-sm font-medium text-gray-900">{{ user?.username }}</span>
        </div>
        <button
          @click="handleLogout"
          class="text-sm text-gray-600 hover:text-red-600 transition-colors"
          title="Sign out"
        >
          <ArrowRightOnRectangleIcon class="w-5 h-5" />
        </button>
      </div>
    </div>

    <!-- Hamburger menu for mobile -->
    <div class="fixed top-4 right-4 sm:hidden z-50">
      <button
        @click="toggleMenu"
        class="flex items-center justify-center w-10 h-10 bg-white shadow-lg rounded-full text-gray-600 hover:text-blue-600 hover:shadow-xl transition-all duration-200"
        aria-label="Toggle menu"
      >
        <Bars3Icon v-if="!isMenuOpen" class="h-6 w-6" />
        <XMarkIcon v-else class="h-6 w-6" />
      </button>
    </div>

    <!-- Navigation buttons (desktop, bottom-right) -->
    <div v-if="isAuthenticated" class="fixed bottom-4 right-4 sm:flex sm:flex-row sm:space-x-2 hidden z-50">
      <router-link
        to="/"
        class="flex items-center justify-center w-12 h-12 bg-white shadow-lg rounded-full text-gray-600 hover:text-blue-600 hover:shadow-xl transition-all duration-200 group"
        active-class="text-blue-600 shadow-xl"
        title="Upload"
      >
        <ArrowUpTrayIcon class="h-6 w-6 group-hover:scale-110 transition-transform duration-200" />
      </router-link>
      <router-link
        to="/settings"
        class="flex items-center justify-center w-12 h-12 bg-white shadow-lg rounded-full text-gray-600 hover:text-blue-600 hover:shadow-xl transition-all duration-200 group"
        active-class="text-blue-600 shadow-xl"
        title="Settings"
      >
        <CogIcon class="h-6 w-6 group-hover:scale-110 transition-transform duration-200" />
      </router-link>
      <router-link
        to="/account"
        class="flex items-center justify-center w-12 h-12 bg-white shadow-lg rounded-full text-gray-600 hover:text-blue-600 hover:shadow-xl transition-all duration-200 group"
        active-class="text-blue-600 shadow-xl"
        title="Account"
      >
        <UserIcon class="h-6 w-6 group-hover:scale-110 transition-transform duration-200" />
      </router-link>
    </div>

    <!-- Login button for non-authenticated users (desktop) -->
    <div v-else class="fixed bottom-4 right-4 sm:block hidden z-50">
      <router-link
        to="/auth"
        class="flex items-center justify-center px-6 py-3 bg-blue-600 hover:bg-blue-700 text-white font-medium rounded-full shadow-lg hover:shadow-xl transition-all duration-200"
      >
        <ArrowRightOnRectangleIcon class="h-5 w-5 mr-2" />
        Sign In
      </router-link>
    </div>

    <!-- Mobile menu (dropdown) -->
    <transition
      enter-active-class="transition ease-out duration-200"
      enter-from-class="opacity-0 translate-y-1"
      enter-to-class="opacity-100 translate-y-0"
      leave-active-class="transition ease-in duration-150"
      leave-from-class="opacity-100 translate-y-0"
      leave-to-class="opacity-0 translate-y-1"
    >
      <div v-if="isMenuOpen" class="sm:hidden fixed bottom-4 right-4 bg-white shadow-md rounded-md z-50">
        <div class="px-4 py-4 flex flex-col space-y-2">
          <!-- Authenticated user menu -->
          <template v-if="isAuthenticated">
            <!-- User info -->
            <div class="flex items-center space-x-2 px-3 py-2 bg-gray-50 rounded-lg mb-2">
              <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
                <UserIcon class="w-5 h-5 text-blue-600" />
              </div>
              <span class="text-sm font-medium text-gray-900">{{ user?.username }}</span>
            </div>
            
            <router-link
              to="/"
              class="flex items-center justify-center w-12 h-12 bg-white shadow-lg rounded-full text-gray-600 hover:text-blue-600 hover:shadow-xl transition-all duration-200 group"
              active-class="text-blue-600 shadow-xl"
              title="Upload"
              @click="toggleMenu"
            >
              <ArrowUpTrayIcon class="h-6 w-6 group-hover:scale-110 transition-transform duration-200" />
            </router-link>
            <router-link
              to="/settings"
              class="flex items-center justify-center w-12 h-12 bg-white shadow-lg rounded-full text-gray-600 hover:text-blue-600 hover:shadow-xl transition-all duration-200 group"
              active-class="text-blue-600 shadow-xl"
              title="Settings"
              @click="toggleMenu"
            >
              <CogIcon class="h-6 w-6 group-hover:scale-110 transition-transform duration-200" />
            </router-link>
            <router-link
              to="/account"
              class="flex items-center justify-center w-12 h-12 bg-white shadow-lg rounded-full text-gray-600 hover:text-blue-600 hover:shadow-xl transition-all duration-200 group"
              active-class="text-blue-600 shadow-xl"
              title="Account"
              @click="toggleMenu"
            >
              <UserIcon class="h-6 w-6 group-hover:scale-110 transition-transform duration-200" />
            </router-link>
            
            <!-- Logout button -->
            <button
              @click="handleLogout"
              class="flex items-center justify-center w-12 h-12 bg-red-50 hover:bg-red-100 rounded-full text-red-600 hover:shadow-lg transition-all duration-200"
              title="Sign out"
            >
              <ArrowRightOnRectangleIcon class="h-6 w-6" />
            </button>
          </template>
          
          <!-- Non-authenticated user menu -->
          <template v-else>
            <router-link
              to="/auth"
              class="flex items-center justify-center px-4 py-3 bg-blue-600 hover:bg-blue-700 text-white font-medium rounded-lg transition-all duration-200"
              @click="toggleMenu"
            >
              <ArrowRightOnRectangleIcon class="h-5 w-5 mr-2" />
              Sign In
            </router-link>
          </template>
        </div>
      </div>
    </transition>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useAuth } from '../composables/useAuth'
import { useRouter } from 'vue-router'
import { ArrowUpTrayIcon, CogIcon, UserIcon, CloudArrowUpIcon, Bars3Icon, XMarkIcon, ArrowRightOnRectangleIcon } from '@heroicons/vue/24/outline'
import axios from 'axios'

const { isAuthenticated, user, logout } = useAuth()
const router = useRouter()

const isMenuOpen = ref(false)
const logoPath = ref<string | null>(null)

const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value
}

const handleLogout = async () => {
  await logout()
  isMenuOpen.value = false
  router.push('/auth')
}

const loadSettings = async () => {
  try {
    const response = await axios.get('/settings')
    logoPath.value = response.data.logo || null
  } catch (error) {
    console.error('Error loading settings:', error)
  }
}

const handleImageError = () => {
  console.error('Logo image failed to load:', logoPath.value)
  logoPath.value = null // Fallback to icon
}

onMounted(() => {
  loadSettings()
})
</script>

<style scoped>
/* Logo styling */
div:first-child {
  font-family: 'Inter', sans-serif;
}

/* Logo backdrop for better visibility */
div:first-child > a {
  padding: 8px 12px;
  border-radius: 12px;
  background-color: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
  transition: all 0.2s ease-in-out;
}

div:first-child > a:hover {
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
  transform: translateY(-1px);
}

/* Navigation buttons hover effects */
.group:hover svg {
  transform: scale(1.1);
}

/* Add backdrop blur for better visibility */
.group {
  backdrop-filter: blur(10px);
  background-color: rgba(255, 255, 255, 0.95);
}

/* Mobile menu container */
.sm\:hidden .bg-white {
  background-color: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
}
</style>
