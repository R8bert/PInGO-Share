<template>
  <div>
    <!-- Logo Section - Top Left -->
    <div class="fixed top-4 left-4 z-50">
      <router-link 
        to="/" 
        class="flex items-center space-x-3 px-4 py-2.5 rounded-2xl backdrop-blur-xl transition-all duration-300 hover:scale-105 group"
        :style="{ 
          backgroundColor: isDark ? 'rgba(255, 255, 255, 0.08)' : 'rgba(255, 255, 255, 0.9)',
          borderColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)',
          borderWidth: '1px',
          boxShadow: isDark 
            ? '0 8px 32px rgba(0, 0, 0, 0.3)' 
            : '0 8px 32px rgba(0, 0, 0, 0.1)'
        }"
      >
        <div class="w-10 h-10 rounded-xl flex items-center justify-center transition-colors duration-300"
             :style="{ backgroundColor: isDark ? 'rgba(59, 130, 246, 0.2)' : 'rgba(59, 130, 246, 0.1)' }">
          <img 
            v-if="logoPath" 
            :src="logoPath ? `http://localhost:8080${logoPath}` : ''" 
            class="w-7 h-7" 
            alt="Logo" 
            @error="handleImageError" 
          />
          <CloudArrowUpIcon 
            v-else 
            class="w-7 h-7 transition-colors duration-300"
            :style="{ color: isDark ? '#60a5fa' : '#3b82f6' }" 
          />
        </div>
        <span 
          class="text-xl font-bold tracking-tight transition-colors duration-300 group-hover:scale-105"
          :style="{ color: isDark ? '#ffffff' : '#1f2937' }"
        >
          {{ navbarTitle }}
        </span>
      </router-link>
    </div>

    <!-- Account Section - Top Right -->
    <div class="fixed top-4 right-4 z-50 hidden md:block">
      <template v-if="isAuthenticated">
        <!-- User Profile Section -->
        <div class="flex items-center space-x-3 px-4 py-2.5 rounded-2xl backdrop-blur-xl"
             :style="{ 
               backgroundColor: isDark ? 'rgba(255, 255, 255, 0.08)' : 'rgba(255, 255, 255, 0.9)',
               borderColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)',
               borderWidth: '1px',
               boxShadow: isDark 
                 ? '0 8px 32px rgba(0, 0, 0, 0.3)' 
                 : '0 8px 32px rgba(0, 0, 0, 0.1)'
             }">
          <!-- User Avatar -->
          <div class="w-8 h-8 rounded-xl overflow-hidden flex items-center justify-center"
               :style="{ backgroundColor: isDark ? 'rgba(59, 130, 246, 0.2)' : 'rgba(59, 130, 246, 0.1)' }">
            <img 
              v-if="user?.avatar" 
              :src="`http://localhost:8080${user.avatar}`" 
              :alt="user.username"
              class="w-full h-full object-cover"
              @error="handleAvatarError"
            />
            <UserIcon 
              v-else 
              class="w-4 h-4"
              :style="{ color: isDark ? '#60a5fa' : '#3b82f6' }" 
            />
          </div>
          
          <!-- Username -->
          <span 
            class="text-sm font-medium"
            :style="{ color: isDark ? '#e5e7eb' : '#374151' }"
          >
            {{ user?.username }}
          </span>
          
          <!-- Logout Button -->
          <button
            @click="handleLogout"
            class="flex items-center justify-center w-8 h-8 rounded-lg transition-all duration-300 hover:scale-110 group"
            :style="{ 
              backgroundColor: isDark ? 'rgba(239, 68, 68, 0.2)' : 'rgba(239, 68, 68, 0.1)',
              color: isDark ? '#f87171' : '#dc2626'
            }"
            title="Sign out"
          >
            <ArrowRightOnRectangleIcon class="w-4 h-4 transition-transform duration-300 group-hover:rotate-12" />
          </button>
        </div>
      </template>
      
      <!-- Non-authenticated state -->
      <template v-else>
        <router-link
          to="/auth"
          class="flex items-center space-x-2 px-6 py-2.5 rounded-2xl font-medium transition-all duration-300 hover:scale-105 group backdrop-blur-xl"
          :style="{
            backgroundColor: isDark ? 'rgba(59, 130, 246, 0.9)' : 'rgba(37, 99, 235, 0.9)',
            color: '#ffffff',
            boxShadow: '0 4px 20px rgba(59, 130, 246, 0.3)'
          }"
        >
          <ArrowRightOnRectangleIcon class="w-4 h-4 transition-transform duration-300 group-hover:translate-x-1" />
          <span>Sign In</span>
        </router-link>
      </template>
    </div>

    <!-- Mobile Menu Button -->
    <button
      @click="toggleMenu"
      class="md:hidden fixed top-4 right-4 z-50 flex items-center justify-center w-10 h-10 rounded-xl backdrop-blur-xl transition-all duration-300 hover:scale-110"
      :style="{ 
        backgroundColor: isDark ? 'rgba(255, 255, 255, 0.08)' : 'rgba(255, 255, 255, 0.9)',
        borderColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)',
        borderWidth: '1px',
        color: isDark ? '#e5e7eb' : '#374151',
        boxShadow: isDark 
          ? '0 8px 32px rgba(0, 0, 0, 0.3)' 
          : '0 8px 32px rgba(0, 0, 0, 0.1)'
      }"
    >
      <Bars3Icon v-if="!isMenuOpen" class="w-5 h-5" />
      <XMarkIcon v-else class="w-5 h-5" />
    </button>

    <!-- Floating Navigation Buttons (Desktop - Bottom Right) -->
    <div class="hidden md:flex fixed bottom-6 right-6 z-50">
      <div class="flex items-center space-x-2 px-3 py-2 rounded-xl backdrop-blur-xl"
           :style="{ 
             backgroundColor: isDark ? 'rgba(255, 255, 255, 0.08)' : 'rgba(255, 255, 255, 0.9)',
             borderColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)',
             borderWidth: '1px',
             boxShadow: isDark 
               ? '0 8px 32px rgba(0, 0, 0, 0.3)' 
               : '0 8px 32px rgba(0, 0, 0, 0.1)'
           }">
        <template v-if="isAuthenticated">
          <!-- Upload Button -->
          <router-link
            to="/"
            class="nav-button group flex items-center justify-center w-10 h-10 rounded-xl transition-all duration-300"
            :class="$route.path === '/' ? 'nav-button-active' : 'nav-button-inactive'"
            title="Upload Files"
          >
            <ArrowUpTrayIcon class="w-5 h-5 transition-transform duration-300 group-hover:scale-110" />
          </router-link>
          
          <!-- Account Button -->
          <router-link
            to="/account"
            class="nav-button group flex items-center justify-center w-10 h-10 rounded-xl transition-all duration-300"
            :class="$route.path === '/account' ? 'nav-button-active' : 'nav-button-inactive'"
            title="Account"
          >
            <UserIcon class="w-5 h-5 transition-transform duration-300 group-hover:scale-110" />
          </router-link>
        </template>
        
        <!-- Theme Toggle Button -->
        <ThemeToggle />
      </div>
    </div>

    <!-- Mobile Menu Overlay -->
    <transition
      enter-active-class="transition duration-200 ease-out"
      enter-from-class="opacity-0 scale-95"
      enter-to-class="opacity-100 scale-100"
      leave-active-class="transition duration-150 ease-in"
      leave-from-class="opacity-100 scale-100"
      leave-to-class="opacity-0 scale-95"
    >
      <div 
        v-if="isMenuOpen" 
        class="md:hidden fixed inset-0 z-40 bg-black bg-opacity-50 backdrop-blur-sm"
        @click="toggleMenu"
      >
        <div 
          class="absolute top-20 right-6 w-72 rounded-2xl backdrop-blur-xl p-6 space-y-4"
          :style="{ 
            backgroundColor: isDark ? 'rgba(31, 41, 55, 0.95)' : 'rgba(255, 255, 255, 0.95)',
            borderColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)',
            borderWidth: '1px',
            boxShadow: '0 20px 50px rgba(0, 0, 0, 0.3)'
          }"
          @click.stop
        >
          <template v-if="isAuthenticated">
            <!-- User Profile -->
            <div class="flex items-center space-x-3 p-3 rounded-xl"
                 :style="{ backgroundColor: isDark ? 'rgba(255, 255, 255, 0.05)' : 'rgba(0, 0, 0, 0.05)' }">
              <div class="w-10 h-10 rounded-xl overflow-hidden flex items-center justify-center"
                   :style="{ backgroundColor: isDark ? 'rgba(59, 130, 246, 0.2)' : 'rgba(59, 130, 246, 0.1)' }">
                <img 
                  v-if="user?.avatar" 
                  :src="`http://localhost:8080${user.avatar}`" 
                  :alt="user.username"
                  class="w-full h-full object-cover"
                  @error="handleAvatarError"
                />
                <UserIcon 
                  v-else 
                  class="w-5 h-5"
                  :style="{ color: isDark ? '#60a5fa' : '#3b82f6' }" 
                />
              </div>
              <div>
                <p class="font-medium" :style="{ color: isDark ? '#f9fafb' : '#111827' }">
                  {{ user?.username }}
                </p>
                <p class="text-sm" :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">
                  {{ user?.email }}
                </p>
              </div>
            </div>

            <!-- Navigation Links -->
            <div class="space-y-2">
              <router-link
                to="/"
                @click="toggleMenu"
                class="flex items-center space-x-3 w-full p-3 rounded-xl transition-all duration-300 hover:scale-[1.02]"
                :style="{ 
                  backgroundColor: $route.path === '/' 
                    ? (isDark ? 'rgba(59, 130, 246, 0.2)' : 'rgba(59, 130, 246, 0.1)')
                    : 'transparent',
                  color: $route.path === '/' 
                    ? (isDark ? '#60a5fa' : '#3b82f6')
                    : (isDark ? '#e5e7eb' : '#374151')
                }"
              >
                <ArrowUpTrayIcon class="w-5 h-5" />
                <span class="font-medium">Upload Files</span>
              </router-link>
              
              <router-link
                to="/account"
                @click="toggleMenu"
                class="flex items-center space-x-3 w-full p-3 rounded-xl transition-all duration-300 hover:scale-[1.02]"
                :style="{ 
                  backgroundColor: $route.path === '/account' 
                    ? (isDark ? 'rgba(59, 130, 246, 0.2)' : 'rgba(59, 130, 246, 0.1)')
                    : 'transparent',
                  color: $route.path === '/account' 
                    ? (isDark ? '#60a5fa' : '#3b82f6')
                    : (isDark ? '#e5e7eb' : '#374151')
                }"
              >
                <UserIcon class="w-5 h-5" />
                <span class="font-medium">Account</span>
              </router-link>
            </div>

            <!-- Bottom Actions -->
            <div class="pt-4 border-t space-y-3"
                 :style="{ borderColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)' }">
              <div class="flex justify-center">
                <ThemeToggle />
              </div>
              
              <button
                @click="handleLogout"
                class="flex items-center justify-center space-x-2 w-full p-3 rounded-xl transition-all duration-300 hover:scale-[1.02]"
                :style="{
                  backgroundColor: isDark ? 'rgba(239, 68, 68, 0.2)' : 'rgba(239, 68, 68, 0.1)',
                  color: isDark ? '#f87171' : '#dc2626'
                }"
              >
                <ArrowRightOnRectangleIcon class="w-4 h-4" />
                <span class="font-medium">Sign Out</span>
              </button>
            </div>
          </template>
          
          <!-- Non-authenticated mobile menu -->
          <template v-else>
            <div class="space-y-4">
              <div class="flex justify-center">
                <ThemeToggle />
              </div>
              
              <router-link
                to="/auth"
                @click="toggleMenu"
                class="flex items-center justify-center space-x-2 w-full p-4 rounded-xl font-medium transition-all duration-300 hover:scale-[1.02]"
                :style="{
                  backgroundColor: isDark ? 'rgba(59, 130, 246, 0.9)' : 'rgba(37, 99, 235, 0.9)',
                  color: '#ffffff'
                }"
              >
                <ArrowRightOnRectangleIcon class="w-5 h-5" />
                <span>Sign In</span>
              </router-link>
            </div>
          </template>
        </div>
      </div>
    </transition>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { useAuth } from '../../composables/useAuth'
import { useTheme } from '../../composables/useTheme'
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

const fetchSettings = async () => {
  try {
    const response = await fetch('http://localhost:8080/settings')
    if (response.ok) {
      const settings = await response.json()
      navbarTitle.value = settings.navbarTitle || 'PinGO'
      logoPath.value = settings.logo || '/logos/004571540fcfd318992384ba96ffb6ae07634920f70e009cf76cf9a1aac603ab.png'
    }
  } catch (error) {
    console.error('Failed to fetch settings:', error)
  }
}

// Listen for settings updates
const handleSettingsUpdate = () => {
  fetchSettings()
}

const handleImageError = () => {
  console.error('Logo image failed to load:', logoPath.value)
  logoPath.value = null // Fallback to icon
}

const handleAvatarError = () => {
  console.error('Avatar image failed to load')
}

onMounted(() => {
  fetchSettings()
  window.addEventListener('settings-updated', handleSettingsUpdate)
})

onUnmounted(() => {
  window.removeEventListener('settings-updated', handleSettingsUpdate)
})
</script>

<style scoped>
/* Navigation button styles */
.nav-button-active {
  background-color: v-bind('isDark ? "rgba(59, 130, 246, 0.2)" : "rgba(59, 130, 246, 0.1)"');
  color: v-bind('isDark ? "#60a5fa" : "#3b82f6"');
  box-shadow: 0 0 0 1px v-bind('isDark ? "rgba(59, 130, 246, 0.3)" : "rgba(59, 130, 246, 0.2)"');
}

.nav-button-inactive {
  color: v-bind('isDark ? "#9ca3af" : "#6b7280"');
}

.nav-button-inactive:hover {
  background-color: v-bind('isDark ? "rgba(255, 255, 255, 0.05)" : "rgba(0, 0, 0, 0.05)"');
  color: v-bind('isDark ? "#e5e7eb" : "#374151"');
}

/* Enhanced backdrop blur for better glassmorphism effect */
nav {
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
}

/* Smooth transitions for all interactive elements */
* {
  transition-property: color, background-color, border-color, text-decoration-color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
}

/* Custom scrollbar for mobile menu */
.mobile-menu::-webkit-scrollbar {
  width: 4px;
}

.mobile-menu::-webkit-scrollbar-track {
  background: transparent;
}

.mobile-menu::-webkit-scrollbar-thumb {
  background: v-bind('isDark ? "rgba(255, 255, 255, 0.2)" : "rgba(0, 0, 0, 0.2)"');
  border-radius: 2px;
}

/* Focus states for accessibility */
button:focus-visible,
a:focus-visible {
  outline: 2px solid v-bind('isDark ? "#60a5fa" : "#3b82f6"');
  outline-offset: 2px;
}

/* Hover effects */
.group:hover svg {
  transform: scale(1.1);
}

/* Active states */
.nav-button:active {
  transform: scale(0.95);
}

/* Mobile menu animations */
@media (max-width: 768px) {
  .mobile-menu {
    animation: slideIn 0.3s ease-out;
  }
}

@keyframes slideIn {
  from {
    opacity: 0;
    transform: translateY(-10px) scale(0.95);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

/* Ensure proper spacing on smaller screens */
@media (max-width: 640px) {
  nav {
    padding-left: 1rem;
    padding-right: 1rem;
  }
}
</style>
