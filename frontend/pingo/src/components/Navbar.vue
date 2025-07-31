<template>
  <div>
    <!-- Logo positioned in top-left corner -->
    <div class="fixed top-4 left-4 z-50">
      <router-link to="/" class="flex items-center space-x-2 px-3 py-2 rounded-lg shadow-lg transition-all duration-300"
                   :style="{ backgroundColor: isDark ? '#1a1a1a' : '#ffffff' }">
        <img v-if="logoPath" :src="logoPath ? `http://localhost:8080${logoPath}` : ''" class="h-8 w-8" alt="Logo" @error="handleImageError" />
        <CloudArrowUpIcon v-else class="h-8 w-8 text-blue-600 dark:text-blue-400" />
        <span class="text-2xl font-extrabold transition-colors duration-300" 
              :style="{ color: isDark ? '#ffffff' : '#000000' }">{{ navbarTitle }}</span>
      </router-link>
    </div>

    <!-- User info in top-right corner for authenticated users -->
    <div v-if="isAuthenticated" class="fixed top-4 right-4 z-50 hidden sm:block">
      <div class="flex items-center space-x-3 shadow-lg rounded-full px-4 py-2 transition-colors duration-300" 
           :style="{ 
             backgroundColor: isDark ? '#1a1a1a' : '#ffffff', 
             color: isDark ? '#ffffff' : '#000000',
             borderColor: isDark ? '#4b5563' : '#e5e7eb',
             borderWidth: '1px'
           }">
        <div class="flex items-center space-x-2">
          <div class="w-8 h-8 rounded-full overflow-hidden flex items-center justify-center transition-colors duration-300"
               :style="{ backgroundColor: isDark ? '#1e3a8a' : '#dbeafe' }">
            <img 
              v-if="user?.avatar" 
              :src="`http://localhost:8080${user.avatar}`" 
              :alt="user.username"
              class="w-full h-full object-cover"
              @error="handleAvatarError"
            />
            <UserIcon v-else class="w-5 h-5 transition-colors duration-300"
                      :style="{ color: isDark ? '#60a5fa' : '#2563eb' }" />
          </div>
          <router-link 
            to="/account" 
            class="text-sm font-medium px-2 py-1 rounded transition-all duration-300"
            :style="{ 
              color: isDark ? '#ffffff' : '#000000',
              backgroundColor: isDark ? '#2a2a2a' : '#f8f9fa'
            }"
          >
            {{ user?.username }}
          </router-link>
        </div>
        <ThemeToggle />
        <button
          @click="handleLogout"
          class="text-sm transition-colors duration-300"
          :style="{ color: isDark ? '#9ca3af' : '#6b7280' }"
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
        class="flex items-center justify-center w-10 h-10 shadow-lg rounded-full transition-all duration-200"
        :style="{ 
          backgroundColor: isDark ? '#1a1a1a' : '#ffffff', 
          color: isDark ? '#ffffff' : '#666666',
          borderColor: isDark ? '#4b5563' : '#e5e7eb',
          borderWidth: '1px'
        }"
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
        class="flex items-center justify-center w-12 h-12 shadow-lg rounded-full transition-all duration-200 group"
        :style="{
          backgroundColor: isDark ? '#1f2937' : '#ffffff',
          borderColor: isDark ? '#4b5563' : '#e5e7eb',
          borderWidth: '1px',
          color: isDark ? '#d1d5db' : '#6b7280'
        }"
        active-class="text-blue-600 dark:text-blue-400 shadow-xl"
        title="Upload"
      >
        <ArrowUpTrayIcon class="h-6 w-6 group-hover:scale-110 transition-transform duration-200" />
      </router-link>
      <router-link
        to="/account"
        class="flex items-center justify-center w-12 h-12 shadow-lg rounded-full transition-all duration-200 group"
        :style="{
          backgroundColor: isDark ? '#1f2937' : '#ffffff',
          borderColor: isDark ? '#4b5563' : '#e5e7eb',
          borderWidth: '1px',
          color: isDark ? '#d1d5db' : '#6b7280'
        }"
        active-class="text-blue-600 dark:text-blue-400 shadow-xl"
        title="Account"
      >
        <UserIcon class="h-6 w-6 group-hover:scale-110 transition-transform duration-200" />
      </router-link>
    </div>

    <!-- Login button for non-authenticated users (desktop) -->
    <div v-else class="fixed bottom-4 right-4 sm:block hidden z-50">
      <router-link
        to="/auth"
        class="flex items-center justify-center px-6 py-3 font-medium rounded-full shadow-lg transition-all duration-200"
        :style="{
          backgroundColor: isDark ? '#2563eb' : '#3b82f6',
          color: '#ffffff'
        }"
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
      <div v-if="isMenuOpen" class="sm:hidden fixed bottom-4 right-4 rounded-md z-50 shadow-lg transition-colors duration-300"
           :style="{ 
             backgroundColor: isDark ? '#1f2937' : '#ffffff',
             borderColor: isDark ? '#4b5563' : '#e5e7eb',
             borderWidth: '1px'
           }">
        <div class="px-4 py-4 flex flex-col space-y-2">
          <!-- Authenticated user menu -->
          <template v-if="isAuthenticated">
            <!-- User info -->
            <div class="flex items-center space-x-2 px-3 py-2 rounded-lg mb-2 transition-colors duration-300"
                 :style="{ backgroundColor: isDark ? '#374151' : '#f9fafb' }">
              <div class="w-8 h-8 rounded-full overflow-hidden flex items-center justify-center transition-colors duration-300"
                   :style="{ backgroundColor: isDark ? '#1e3a8a' : '#dbeafe' }">
                <img 
                  v-if="user?.avatar" 
                  :src="`http://localhost:8080${user.avatar}`" 
                  :alt="user.username"
                  class="w-full h-full object-cover"
                  @error="handleAvatarError"
                />
                <UserIcon v-else class="w-5 h-5 transition-colors duration-300"
                          :style="{ color: isDark ? '#60a5fa' : '#2563eb' }" />
              </div>
              <router-link 
                to="/account" 
                class="text-sm font-medium px-2 py-1 rounded transition-all duration-300"
                :style="{ 
                  color: isDark ? '#ffffff' : '#000000',
                  backgroundColor: isDark ? '#2a2a2a' : '#f8f9fa'
                }"
                @click="toggleMenu"
              >
                {{ user?.username }}
              </router-link>
            </div>
            
            <!-- Theme Toggle for Mobile -->
            <div class="flex justify-center mb-2">
              <ThemeToggle />
            </div>
            
            <router-link
              to="/"
              class="flex items-center justify-center w-12 h-12 rounded-full transition-all duration-200 group"
              :style="{
                backgroundColor: isDark ? '#374151' : '#ffffff',
                borderColor: isDark ? '#4b5563' : '#e5e7eb',
                borderWidth: '1px',
                color: isDark ? '#d1d5db' : '#6b7280'
              }"
              active-class="text-blue-600 dark:text-blue-400 shadow-xl"
              title="Upload"
              @click="toggleMenu"
            >
              <ArrowUpTrayIcon class="h-6 w-6 group-hover:scale-110 transition-transform duration-200" />
            </router-link>
            <router-link
              to="/account"
              class="flex items-center justify-center w-12 h-12 shadow-lg rounded-full transition-all duration-200 group"
              :style="{
                backgroundColor: isDark ? '#374151' : '#ffffff',
                borderColor: isDark ? '#4b5563' : '#e5e7eb',
                borderWidth: '1px',
                color: isDark ? '#d1d5db' : '#6b7280'
              }"
              active-class="text-blue-600 dark:text-blue-400 shadow-xl"
              title="Account"
              @click="toggleMenu"
            >
              <UserIcon class="h-6 w-6 group-hover:scale-110 transition-transform duration-200" />
            </router-link>
            
            <!-- Logout button -->
            <button
              @click="handleLogout"
              class="flex items-center justify-center w-12 h-12 rounded-full transition-all duration-200"
              :style="{
                backgroundColor: isDark ? '#7f1d1d' : '#fef2f2',
                borderColor: isDark ? '#dc2626' : '#fecaca',
                borderWidth: '1px',
                color: isDark ? '#f87171' : '#dc2626'
              }"
              title="Sign out"
            >
              <ArrowRightOnRectangleIcon class="h-6 w-6" />
            </button>
          </template>
          
          <!-- Non-authenticated user menu -->
          <template v-else>
            <!-- Theme Toggle for Non-authenticated Mobile Users -->
            <div class="flex justify-center mb-2">
              <ThemeToggle />
            </div>
            
            <router-link
              to="/auth"
              class="flex items-center justify-center px-4 py-3 font-medium rounded-lg transition-all duration-200"
              :style="{
                backgroundColor: isDark ? '#2563eb' : '#3b82f6',
                color: '#ffffff'
              }"
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
import { ref } from 'vue'
import { useAuth } from '../composables/useAuth'
import { useTheme } from '../composables/useTheme'
import { useRouter } from 'vue-router'
import { ArrowUpTrayIcon, UserIcon, CloudArrowUpIcon, Bars3Icon, XMarkIcon, ArrowRightOnRectangleIcon } from '@heroicons/vue/24/outline'
// import axios from 'axios'
import ThemeToggle from './ThemeToggle.vue'

const { isAuthenticated, user, logout } = useAuth()
const { isDark } = useTheme()
const router = useRouter()

const isMenuOpen = ref(false)
const logoPath = ref<string | null>('/logos/004571540fcfd318992384ba96ffb6ae07634920f70e009cf76cf9a1aac603ab.png')
const navbarTitle = ref('PinGO')

const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value
}

const handleLogout = async () => {
  await logout()
  isMenuOpen.value = false
  router.push('/auth')
}


const handleImageError = () => {
  console.error('Logo image failed to load:', logoPath.value)
  logoPath.value = null // Fallback to icon
}

const handleAvatarError = () => {
  console.error('Avatar image failed to load')
}
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
  /* box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06); */
  transition: all 0.2s ease-in-out;
}

div:first-child > a:hover {
  /* box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05); */
  transform: translateY(-1px);
    /* Add a subtle scale effect */

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
