<template>
  <div>
    <!-- Left Glass Panel -->
    <div 
      class="fixed top-4 left-4 z-50 backdrop-blur-xl transition-all duration-300 hover:scale-105 animate-float-gentle"
      :style="glassPanelStyle"
    >
      <router-link 
        to="/" 
        class="flex items-center space-x-3 px-6 py-3 rounded-xl transition-all duration-300 group hover:scale-105"
        :style="{ 
          backgroundColor: 'transparent'
        }"
      >
        <div class="rounded-lg flex items-center justify-center transition-colors duration-300">
          <img 
            v-if="logoPath" 
            :src="getAssetUrl(logoPath)" 
            class="w-8 h-8 rounded-lg" 
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
          class="text-lg font-bold tracking-tight transition-colors duration-300"
          :style="{ color: isDark ? '#f8fafc' : '#0f172a' }"
        >
          {{ navbarTitle }}
        </span>
      </router-link>
    </div>    <!-- Right Glass Panel - Navigation & User Controls -->
    <div 
      class="fixed top-4 right-4 z-50 backdrop-blur-xl transition-all duration-300 animate-float-gentle-reverse"
      :style="glassPanelStyle"
    >
      <div class="flex items-center px-3 py-3 rounded-xl">
        <!-- Navigation Links (Desktop) -->
        <div class="hidden md:flex items-center space-x-2 mr-4">
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
            v-if="isAuthenticated"
            to="/account"
            class="nav-button group flex items-center justify-center w-10 h-10 rounded-xl transition-all duration-300"
            :class="$route.path === '/account' ? 'nav-button-active' : 'nav-button-inactive'"
            title="Account Settings"
          >
            <UserIcon class="w-5 h-5 transition-transform duration-300 group-hover:scale-110" />
          </router-link>
        </div>

        <!-- Separator for desktop -->
        <div 
          v-if="isAuthenticated" 
          class="hidden md:block w-px h-6 mx-2"
          :style="{ backgroundColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)' }"
        ></div>

        <!-- Simple Theme Toggle Button -->
        <button
          @click="toggleTheme"
          class="flex items-center justify-center w-10 h-10 rounded-xl transition-all duration-300 hover:scale-105 mr-3"
          :style="{ 
            backgroundColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)'
          }"
          :title="isDark ? 'Switch to Light Mode' : 'Switch to Dark Mode'"
        >
          <!-- Sun icon for dark mode (switch to light) -->
          <svg v-if="isDark" class="w-5 h-5 text-yellow-400 transition-all duration-300" fill="currentColor" viewBox="0 0 24 24">
            <path d="M12 2.25a.75.75 0 01.75.75v2.25a.75.75 0 01-1.5 0V3a.75.75 0 01.75-.75zM7.5 12a4.5 4.5 0 119 0 4.5 4.5 0 01-9 0zM18.894 6.166a.75.75 0 00-1.06-1.06l-1.591 1.59a.75.75 0 101.06 1.061l1.591-1.59zM21.75 12a.75.75 0 01-.75.75h-2.25a.75.75 0 010-1.5H21a.75.75 0 01.75.75zM17.834 18.894a.75.75 0 001.06-1.06l-1.59-1.591a.75.75 0 10-1.061 1.06l1.59 1.591zM12 18a.75.75 0 01.75.75V21a.75.75 0 01-1.5 0v-2.25A.75.75 0 0112 18zM7.758 17.303a.75.75 0 00-1.061-1.06l-1.591 1.59a.75.75 0 001.06 1.061l1.591-1.59zM6 12a.75.75 0 01-.75.75H3a.75.75 0 010-1.5h2.25A.75.75 0 016 12zM6.697 7.757a.75.75 0 001.06-1.06l-1.59-1.591a.75.75 0 00-1.061 1.06l1.59 1.591z"/>
          </svg>
          <!-- Moon icon for light mode (switch to dark) -->
          <svg v-else class="w-5 h-5 text-slate-600 transition-all duration-300" fill="currentColor" viewBox="0 0 24 24">
            <path fill-rule="evenodd" d="M9.528 1.718a.75.75 0 01.162.819A8.97 8.97 0 009 6a9 9 0 009 9 8.97 8.97 0 003.463-.69.75.75 0 01.981.98 10.503 10.503 0 01-9.694 6.46c-5.799 0-10.5-4.701-10.5-10.5 0-4.368 2.667-8.112 6.46-9.694a.75.75 0 01.818.162z" clip-rule="evenodd"/>
          </svg>
        </button>
        


        <!-- User Profile Section -->
        <template v-if="isAuthenticated">
          <div class="flex items-center space-x-3 px-3 py-2 rounded-xl"
               :style="{ 
                 backgroundColor: isDark ? 'rgba(255, 255, 255, 0.05)' : 'rgba(0, 0, 0, 0.05)'
               }">
            <!-- User Avatar -->
            <div class="w-8 h-8 rounded-xl overflow-hidden flex items-center justify-center"
                 :style="{ backgroundColor: isDark ? 'rgba(59, 130, 246, 0.2)' : 'rgba(59, 130, 246, 0.1)' }">
              <img 
                v-if="user?.avatar" 
                :src="getAssetUrl(user.avatar)" 
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
            
            <!-- Username (hidden on small screens) -->
            <span 
              class="text-sm font-medium truncate max-w-20 hidden lg:block"
              :style="{ color: isDark ? '#e2e8f0' : '#334155' }"
              :title="user?.username"
            >
              {{ user?.username }}
            </span>
            
            <!-- Logout Button -->
            <button
              @click="handleLogout"
              class="flex items-center justify-center w-7 h-7 rounded-lg transition-all duration-300 hover:scale-110 group"
              :style="{ 
                backgroundColor: isDark ? 'rgba(239, 68, 68, 0.2)' : 'rgba(239, 68, 68, 0.1)',
                color: isDark ? '#fca5a5' : '#dc2626'
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
            class="flex items-center space-x-2 px-6 py-2.5 rounded-xl font-medium transition-all duration-300 hover:scale-105 group"
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

        <!-- Mobile Menu Button -->
        <button
          @click="toggleMenu"
          class="md:hidden flex items-center justify-center w-10 h-10 rounded-xl transition-all duration-300 hover:scale-110 ml-2"
          :style="{ 
            backgroundColor: isDark ? 'rgba(255, 255, 255, 0.05)' : 'rgba(0, 0, 0, 0.05)',
            color: isDark ? '#e2e8f0' : '#334155'
          }"
        >
          <Bars3Icon v-if="!isMenuOpen" class="w-5 h-5" />
          <XMarkIcon v-else class="w-5 h-5" />
        </button>
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
        class="md:hidden fixed inset-0 z-40 bg-black/50 backdrop-blur-sm"
        @click="toggleMenu"
      >
        <div 
          class="absolute top-20 right-6 w-72 rounded-xl backdrop-blur-xl p-6 space-y-4"
          :style="{ 
            backgroundColor: isDark ? 'rgba(15, 23, 42, 0.95)' : 'rgba(255, 255, 255, 0.95)',
            borderColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)',
            boxShadow: '0 25px 50px rgba(0, 0, 0, 0.4)'
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
                  :src="getAssetUrl(user.avatar)" 
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
              <div class="flex-1 min-w-0">
                <p class="font-medium truncate" :style="{ color: isDark ? '#f8fafc' : '#0f172a' }" :title="user?.username">
                  {{ user?.username }}
                </p>
                <p class="text-sm truncate" :style="{ color: isDark ? '#94a3b8' : '#64748b' }" :title="user?.email">
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
                    : (isDark ? '#e2e8f0' : '#334155')
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
                    : (isDark ? '#e2e8f0' : '#334155')
                }"
              >
                <UserIcon class="w-5 h-5" />
                <span class="font-medium">Account</span>
              </router-link>
            </div>

            <!-- Bottom Actions -->
            <div class="pt-4 border-t space-y-3"
                 :style="{ borderColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)' }">
              <button
                @click="handleLogout"
                class="flex items-center justify-center space-x-2 w-full p-3 rounded-xl transition-all duration-300 hover:scale-[1.02]"
                :style="{
                  backgroundColor: isDark ? 'rgba(239, 68, 68, 0.2)' : 'rgba(239, 68, 68, 0.1)',
                  color: isDark ? '#fca5a5' : '#dc2626'
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
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useAuth } from '../../composables/useAuth'
import { useTheme } from '../../composables/useTheme'
import { useRouter } from 'vue-router'
import { ArrowUpTrayIcon, UserIcon, CloudArrowUpIcon, Bars3Icon, XMarkIcon, ArrowRightOnRectangleIcon } from '@heroicons/vue/24/outline'
import { getApiUrl, getAssetUrl } from '../../utils/apiUtils'

const { isAuthenticated, user, logout } = useAuth()
const { isDark, toggleTheme } = useTheme()
const router = useRouter()

const isMenuOpen = ref(false)
const logoPath = ref<string | null>('/logos/004571540fcfd318992384ba96ffb6ae07634920f70e009cf76cf9a1aac603ab.png')
const navbarTitle = ref('PinGO')

// Simple transparent panel style - removing glassmorphism effects
const glassPanelStyle = computed(() => ({
  backgroundColor: isDark.value 
    ? 'rgba(15, 23, 42, 0.7)' 
    : 'rgba(255, 255, 255, 0.7)',
  backdropFilter: 'blur(20px)',
  WebkitBackdropFilter: 'blur(20px)',
}))

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
    const response = await fetch(getApiUrl('/settings'))
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
  background-color: v-bind('isDark ? "rgba(59, 130, 246, 0.25)" : "rgba(59, 130, 246, 0.15)"');
  color: v-bind('isDark ? "#60a5fa" : "#3b82f6"');
  box-shadow: 0 0 0 1px v-bind('isDark ? "rgba(59, 130, 246, 0.3)" : "rgba(59, 130, 246, 0.25)"');
}

.nav-button-inactive {
  color: v-bind('isDark ? "#94a3b8" : "#64748b"');
}

.nav-button-inactive:hover {
  background-color: v-bind('isDark ? "rgba(255, 255, 255, 0.05)" : "rgba(0, 0, 0, 0.05)"');
  color: v-bind('isDark ? "#e2e8f0" : "#334155"');
}

/* Simple backdrop blur for better performance */

/* Smooth transitions for all interactive elements */
* {
  transition-property: color, background-color, border-color, text-decoration-color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
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

/* Gentle floating animations */
@keyframes float-gentle {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-2px);
  }
}

@keyframes float-gentle-reverse {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(2px);
  }
}

.animate-float-gentle {
  animation: float-gentle 4s ease-in-out infinite;
}

.animate-float-gentle-reverse {
  animation: float-gentle-reverse 4s ease-in-out infinite;
  animation-delay: 2s;
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

/* Simple hover states */
.nav-button:hover {
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
}

/* Smooth morphing effects */
.transition-all {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.hover\:scale-105:hover {
  transform: scale(1.05);
}

.hover\:scale-110:hover {
  transform: scale(1.1);
}

/* Responsive design improvements */
@media (max-width: 640px) {
  .fixed.top-4.left-4 {
    left: 0.5rem;
  }
  
  .fixed.top-4.right-4 {
    right: 0.5rem;
  }
}
</style>
