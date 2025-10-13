<template>
  <!-- Windows XP Style Navbar -->
  <div class="xp-navbar-container">
    <!-- Top bar (XP Blue style) -->
    <div class="xp-top-bar">
      <!-- Logo Section - Left -->
      <router-link 
        to="/" 
        class="xp-logo-section"
      >
        <div class="xp-logo-icon">
          <img 
            v-if="logoPath" 
            :src="getAssetUrl(logoPath)" 
            class="logo-image" 
            alt="Logo" 
            @error="handleImageError"
          />
          <CloudArrowUpIcon 
            v-else 
            class="logo-icon-fallback"
          />
        </div>
        <span class="xp-title">
          {{ navbarTitle }}
        </span>
      </router-link>

      <!-- Start Button (XP Style) -->
      <button class="xp-start-button" @click="toggleStartMenu" title="Click here to begin">
        <!-- Windows XP Logo (4 colored panes) -->
        <svg width="22" height="22" viewBox="0 0 24 24" fill="none" class="xp-start-logo">
          <!-- Red pane (top-left) -->
          <path d="M3 3 L11 3 L11 11 L3 11 Z" fill="#f25022"/>
          <!-- Green pane (top-right) -->
          <path d="M13 3 L21 3 L21 11 L13 11 Z" fill="#7fba00"/>
          <!-- Blue pane (bottom-left) -->
          <path d="M3 13 L11 13 L11 21 L3 21 Z" fill="#00a4ef"/>
          <!-- Yellow pane (bottom-right) -->
          <path d="M13 13 L21 13 L21 21 L13 21 Z" fill="#ffb900"/>
        </svg>
        <span>start</span>
      </button>

      <div class="xp-taskbar-divider"></div>

      <!-- Taskbar Items (Open Windows) - Space for dynamic content -->
      <div class="xp-taskbar-items">
        <slot name="taskbar-items">
          <!-- This will be filled by parent components with open windows -->
        </slot>
      </div>

      <!-- Right Section -->
      <div class="xp-right-section">
        <!-- Theme Toggle -->
        <div class="xp-theme-toggle hidden md:block">
          <ThemeToggle />
        </div>

        <!-- System Tray Icons -->
        <div class="xp-system-tray hidden md:flex">
          <!-- Volume Icon -->
          <svg width="14" height="14" viewBox="0 0 16 16" fill="white" class="tray-icon" title="Volume">
            <path d="M2 6H5L9 2V14L5 10H2V6Z"/>
            <path d="M11 4C12 5 13 6 13 8C13 10 12 11 11 12" stroke="white" fill="none"/>
          </svg>
          
          <!-- Network Icon -->
          <svg width="14" height="14" viewBox="0 0 16 16" fill="white" class="tray-icon" title="Network">
            <rect x="2" y="10" width="4" height="4"/>
            <rect x="6" y="7" width="4" height="7"/>
            <rect x="10" y="4" width="4" height="10"/>
          </svg>
          
          <!-- Clock -->
          <span class="xp-clock">{{ currentTime }}</span>
        </div>
        
        <!-- GitHub Link -->
        <a 
          href="https://github.com/R8bert/PInGO-Share" 
          target="_blank" 
          rel="noopener noreferrer"
          class="xp-button xp-icon-button"
          title="View GitHub Repository"
        >
          <svg 
            class="button-icon" 
            fill="currentColor" 
            viewBox="0 0 24 24"
          >
            <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
          </svg>
        </a>

        <!-- Account Section (Desktop) -->
        <div v-if="isAuthenticated" class="xp-user-section hidden md:flex">
          <!-- User Info -->
          <div class="xp-user-info">
            <div class="xp-avatar">
              <img 
                v-if="user?.avatar" 
                :src="getAssetUrl(user.avatar)" 
                :alt="user.username"
                class="avatar-image"
                @error="handleAvatarError"
              />
              <UserIcon v-else class="avatar-icon" />
            </div>
            <span class="user-name">{{ user?.username }}</span>
          </div>
          
          <!-- Logout Button -->
          <button
            @click="handleLogout"
            class="xp-button xp-logout-button"
            title="Sign out"
          >
            <ArrowRightOnRectangleIcon class="button-icon" />
          </button>
        </div>

        <!-- Sign In Button (Desktop, not authenticated) -->
        <router-link
          v-else
          to="/auth"
          class="xp-button xp-signin-button hidden md:flex"
        >
          <ArrowRightOnRectangleIcon class="button-icon" />
          <span>Sign In</span>
        </router-link>

        <!-- Mobile Menu Toggle -->
        <button
          @click="toggleMenu"
          class="xp-button xp-menu-toggle md:hidden"
        >
          <Bars3Icon v-if="!isMenuOpen" class="button-icon" />
          <XMarkIcon v-else class="button-icon" />
        </button>
      </div>
    </div>

    <!-- Mobile Menu -->
    <transition name="xp-menu">
      <div 
        v-if="isMenuOpen" 
        class="xp-mobile-menu md:hidden"
        @click="toggleMenu"
      >
        <div class="xp-mobile-panel" @click.stop>
          <template v-if="isAuthenticated">
            <!-- User Profile -->
            <div class="xp-mobile-user">
              <div class="xp-avatar large">
                <img 
                  v-if="user?.avatar" 
                  :src="getAssetUrl(user.avatar)" 
                  :alt="user.username"
                  class="avatar-image"
                  @error="handleAvatarError"
                />
                <UserIcon v-else class="avatar-icon" />
              </div>
              <div class="user-details">
                <p class="user-name-large">{{ user?.username }}</p>
                <p class="user-email">{{ user?.email }}</p>
              </div>
            </div>

            <!-- Navigation Links -->
            <div class="xp-mobile-links">
              <router-link
                to="/"
                @click="toggleMenu"
                class="xp-mobile-link"
                :class="{ 'active': $route.path === '/' }"
              >
                <ArrowUpTrayIcon class="link-icon" />
                <span>Upload Files</span>
              </router-link>
              
              <router-link
                to="/account"
                @click="toggleMenu"
                class="xp-mobile-link"
                :class="{ 'active': $route.path === '/account' }"
              >
                <UserIcon class="link-icon" />
                <span>Account</span>
              </router-link>
              
              <a
                href="https://github.com/R8bert/PInGO-Share"
                target="_blank"
                rel="noopener noreferrer"
                @click="toggleMenu"
                class="xp-mobile-link"
              >
                <svg class="link-icon" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M12 0C5.374 0 0 5.373 0 12 0 17.302 3.438 21.8 8.207 23.387c.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23A11.509 11.509 0 0112 5.803c1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576C20.566 21.797 24 17.3 24 12c0-6.627-5.373-12-12-12z"/>
                </svg>
                <span>GitHub</span>
              </a>
            </div>

            <!-- Bottom Actions -->
            <div class="xp-mobile-actions">
              <ThemeToggle />
              
              <button
                @click="handleLogout"
                class="xp-button xp-logout-full"
              >
                <ArrowRightOnRectangleIcon class="button-icon" />
                <span>Sign Out</span>
              </button>
            </div>
          </template>
          
          <!-- Non-authenticated mobile menu -->
          <template v-else>
            <div class="xp-mobile-auth">
              <ThemeToggle />
              
              <a
                href="https://github.com/R8bert/PInGO-Share"
                target="_blank"
                rel="noopener noreferrer"
                @click="toggleMenu"
                class="xp-button xp-github-button"
              >
                <svg class="button-icon" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M12 0C5.374 0 0 5.373 0 12 0 17.302 3.438 21.8 8.207 23.387c.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23A11.509 11.509 0 0112 5.803c1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576C20.566 21.797 24 17.3 24 12c0-6.627-5.373-12-12-12z"/>
                </svg>
                <span>GitHub</span>
              </a>
              
              <router-link
                to="/auth"
                @click="toggleMenu"
                class="xp-button xp-signin-button"
              >
                <ArrowRightOnRectangleIcon class="button-icon" />
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
import { useRouter } from 'vue-router'
import { ArrowUpTrayIcon, UserIcon, CloudArrowUpIcon, Bars3Icon, XMarkIcon, ArrowRightOnRectangleIcon } from '@heroicons/vue/24/outline'
import { getApiUrl, getAssetUrl } from '../../utils/apiUtils'
import ThemeToggle from './ThemeToggle.vue'
import { useStartMenu } from '../../composables/useStartMenu'

const { isAuthenticated, user, logout } = useAuth()
const router = useRouter()
const { toggleStartMenu } = useStartMenu()

const isMenuOpen = ref(false)
const logoPath = ref<string | null>('/logos/004571540fcfd318992384ba96ffb6ae07634920f70e009cf76cf9a1aac603ab.png')
const navbarTitle = ref('PinGO')
const currentTime = ref(new Date().toLocaleTimeString('en-US', { 
    hour: '2-digit', 
    minute: '2-digit',
    hour12: true 
}))

let clockInterval: number | undefined

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
  
  // Start clock update
  clockInterval = setInterval(() => {
    currentTime.value = new Date().toLocaleTimeString('en-US', { 
      hour: '2-digit', 
      minute: '2-digit',
      hour12: true 
    })
  }, 1000)
})

onUnmounted(() => {
  window.removeEventListener('settings-updated', handleSettingsUpdate)
  if (clockInterval) {
    clearInterval(clockInterval)
  }
})
</script>

<style scoped>
/* ===========================================
   WINDOWS XP NAVBAR THEME
   Luna-style navigation bar
   =========================================== */

.xp-navbar-container {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  z-index: 1000;
}

/* Taskbar - XP Blue Gradient */
.xp-top-bar {
  height: 30px;
  background: linear-gradient(to bottom,
    #0054e3 0%,
    #0054e3 100%
  );
  border-top: 1px solid #0064f3;
  border-bottom: 1px solid #003db3;
  box-shadow: 0 -1px 3px rgba(0, 0, 0, 0.2);
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 8px;
}

/* Logo Section */
.xp-logo-section {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 2px 6px;
  background: transparent;
  border: none;
  text-decoration: none;
  transition: none;
}

.xp-logo-section:hover {
  background: rgba(255, 255, 255, 0.1);
}

.xp-logo-section:active {
  background: rgba(0, 0, 0, 0.1);
}

.xp-logo-icon {
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: transparent;
  border-radius: 0;
}

.logo-image {
  width: 18px;
  height: 18px;
  border-radius: 0;
}

.logo-icon-fallback {
  width: 16px;
  height: 16px;
  color: #ffffff;
}

.xp-title {
  font-size: 11px;
  font-weight: bold;
  color: #ffffff;
  text-shadow: none;
  letter-spacing: 0;
  font-family: 'Tahoma', 'Segoe UI', sans-serif;
}

/* Start Button (Authentic Windows XP Style) */
.xp-start-button {
  display: flex;
  align-items: center;
  gap: 5px;
  height: 30px;
  padding: 0 18px 0 6px;
  margin: 0;
  background: linear-gradient(to bottom,
    #3ece51 0%,
    #2dc73d 50%,
    #23b439 100%
  );
  border: 1px solid #16a228;
  border-right: 1px solid #0e721c;
  border-bottom: 1px solid #0e721c;
  border-radius: 0 8px 0 0;
  color: #ffffff;
  font-size: 11px;
  font-weight: bold;
  font-family: 'Tahoma', 'Segoe UI', sans-serif;
  cursor: pointer;
  box-shadow: 
    inset 1px 1px 0 rgba(255, 255, 255, 0.4),
    inset -1px -1px 0 rgba(0, 0, 0, 0.2);
  transition: none;
  position: relative;
}

.xp-start-button:hover {
  background: linear-gradient(to bottom,
    #4dde61 0%,
    #3dd74d 50%,
    #33c749 100%
  );
}

.xp-start-button:active {
  background: linear-gradient(to bottom,
    #2ebe41 0%,
    #24b737 50%,
    #1aa32d 100%
  );
  box-shadow: 
    inset 1px 1px 2px rgba(0, 0, 0, 0.3),
    inset -1px -1px 0 rgba(255, 255, 255, 0.1);
  padding-top: 1px;
}

.xp-start-logo {
  width: 22px;
  height: 22px;
  flex-shrink: 0;
  filter: drop-shadow(0 1px 1px rgba(0, 0, 0, 0.3));
}

.xp-start-button span {
  text-transform: lowercase;
  font-style: italic;
  text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.4);
  letter-spacing: 0.5px;
}

/* Taskbar Divider */
.xp-taskbar-divider {
  width: 1px;
  height: 20px;
  background: rgba(255, 255, 255, 0.2);
  margin: 0 4px;
}

/* Taskbar Items Container */
.xp-taskbar-items {
  flex: 1;
  display: flex;
  align-items: center;
  gap: 2px;
  min-width: 0;
}

/* System Tray */
.xp-system-tray {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 0 8px;
  height: 24px;
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.1);
  margin-right: 4px;
}

.tray-icon {
  width: 14px;
  height: 14px;
  cursor: pointer;
  opacity: 0.9;
  transition: opacity 0.15s;
}

.tray-icon:hover {
  opacity: 1;
}

.xp-clock {
  font-size: 11px;
  color: #ffffff;
  font-family: 'Tahoma', 'Segoe UI', sans-serif;
  font-weight: normal;
  white-space: nowrap;
  margin-left: 2px;
}

/* Right Section */
.xp-right-section {
  display: flex;
  align-items: center;
  gap: 4px;
}

/* Theme Toggle Container */
.xp-theme-toggle {
  display: flex;
  align-items: center;
  margin-right: 4px;
}

/* XP Button Base */
.xp-button {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 4px;
  padding: 2px 8px;
  min-height: 22px;
  background: transparent;
  border: 1px solid transparent;
  border-radius: 0;
  color: #ffffff;
  font-size: 11px;
  font-weight: normal;
  cursor: pointer;
  transition: none;
  text-decoration: none;
  text-shadow: none;
  box-shadow: none;
}

.xp-button:hover {
  background: rgba(255, 255, 255, 0.15);
  border-color: rgba(255, 255, 255, 0.3);
}

.xp-button:active {
  background: rgba(0, 0, 0, 0.15);
  border-color: rgba(0, 0, 0, 0.3);
}

.button-icon {
  width: 14px;
  height: 14px;
}

.xp-icon-button {
  width: 24px;
  height: 24px;
  padding: 4px;
}

/* User Section */
.xp-user-section {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 4px 8px 4px 4px;
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 4px;
}

.xp-user-info {
  display: flex;
  align-items: center;
  gap: 8px;
}

.xp-avatar {
  width: 32px;
  height: 32px;
  border-radius: 4px;
  overflow: hidden;
  background: rgba(59, 130, 246, 0.3);
  display: flex;
  align-items: center;
  justify-content: center;
}

.xp-avatar.large {
  width: 48px;
  height: 48px;
}

.avatar-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.avatar-icon {
  width: 18px;
  height: 18px;
  color: #60a5fa;
}

.user-name {
  font-size: 14px;
  font-weight: 500;
  color: #ffffff;
  text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.5);
  max-width: 120px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.xp-logout-button {
  width: 32px;
  height: 32px;
  padding: 6px;
  background: linear-gradient(to bottom,
    rgba(239, 68, 68, 0.4) 0%,
    rgba(220, 38, 38, 0.4) 100%
  );
}

.xp-logout-button:hover {
  background: linear-gradient(to bottom,
    rgba(239, 68, 68, 0.6) 0%,
    rgba(220, 38, 38, 0.6) 100%
  );
}

/* Sign In Button */
.xp-signin-button {
  background: linear-gradient(to bottom,
    rgba(59, 130, 246, 0.9) 0%,
    rgba(37, 99, 235, 0.9) 100%
  );
  font-weight: 600;
}

.xp-signin-button:hover {
  background: linear-gradient(to bottom,
    rgba(59, 130, 246, 1) 0%,
    rgba(37, 99, 235, 1) 100%
  );
}

/* Menu Toggle */
.xp-menu-toggle {
  width: 40px;
  height: 40px;
}

/* Mobile Menu */
.xp-mobile-menu {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.5);
  z-index: 999;
  display: flex;
  align-items: flex-start;
  justify-content: flex-end;
  padding: 72px 16px 16px 16px;
}

.xp-mobile-panel {
  width: 100%;
  max-width: 320px;
  background: linear-gradient(to bottom,
    #f8f8f8 0%,
    #e8e8e8 100%
  );
  border: 2px solid #808080;
  border-radius: 6px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

/* Mobile User Section */
.xp-mobile-user {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px;
  background: rgba(59, 130, 246, 0.1);
  border: 1px solid rgba(59, 130, 246, 0.2);
  border-radius: 4px;
}

.user-details {
  flex: 1;
  min-width: 0;
}

.user-name-large {
  font-size: 16px;
  font-weight: 600;
  color: #000000;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.user-email {
  font-size: 13px;
  color: #666666;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

/* Mobile Links */
.xp-mobile-links {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.xp-mobile-link {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px 12px;
  background: linear-gradient(to bottom,
    #ffffff 0%,
    #f0f0f0 100%
  );
  border: 1px solid #c0c0c0;
  border-radius: 4px;
  color: #000000;
  font-size: 14px;
  font-weight: 500;
  text-decoration: none;
  transition: all 0.2s;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}

.xp-mobile-link:hover {
  background: linear-gradient(to bottom,
    #f0f8ff 0%,
    #e0f0ff 100%
  );
  border-color: #3b82f6;
}

.xp-mobile-link.active {
  background: linear-gradient(to bottom,
    #3b82f6 0%,
    #2563eb 100%
  );
  color: #ffffff;
  border-color: #1d4ed8;
}

.link-icon {
  width: 18px;
  height: 18px;
}

/* Mobile Actions */
.xp-mobile-actions {
  display: flex;
  flex-direction: column;
  gap: 12px;
  padding-top: 12px;
  border-top: 1px solid #c0c0c0;
}

.xp-logout-full {
  width: 100%;
  justify-content: center;
  background: linear-gradient(to bottom,
    rgba(239, 68, 68, 0.9) 0%,
    rgba(220, 38, 38, 0.9) 100%
  );
  font-weight: 600;
}

.xp-logout-full:hover {
  background: linear-gradient(to bottom,
    rgba(239, 68, 68, 1) 0%,
    rgba(220, 38, 38, 1) 100%
  );
}

/* Mobile Auth (Not authenticated) */
.xp-mobile-auth {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.xp-github-button {
  width: 100%;
  justify-content: center;
  background: linear-gradient(to bottom,
    #ffffff 0%,
    #f0f0f0 100%
  );
  color: #000000;
  border-color: #c0c0c0;
}

.xp-github-button:hover {
  background: linear-gradient(to bottom,
    #f8f8f8 0%,
    #e8e8e8 100%
  );
}

/* Menu Transition */
.xp-menu-enter-active,
.xp-menu-leave-active {
  transition: opacity 0.2s ease;
}

.xp-menu-enter-from,
.xp-menu-leave-to {
  opacity: 0;
}

.xp-menu-enter-active .xp-mobile-panel,
.xp-menu-leave-active .xp-mobile-panel {
  transition: transform 0.2s ease;
}

.xp-menu-enter-from .xp-mobile-panel,
.xp-menu-leave-to .xp-mobile-panel {
  transform: translateY(-20px);
}

/* Focus states for accessibility */
button:focus-visible,
a:focus-visible {
  outline: 2px solid #3b82f6;
  outline-offset: 2px;
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .xp-top-bar {
    height: 52px;
    padding: 0 12px;
  }
  
  .xp-title {
    font-size: 16px;
  }
}

@media (max-width: 640px) {
  .xp-logo-section {
    gap: 6px;
    padding: 4px 8px;
  }
  
  .xp-logo-icon {
    width: 28px;
    height: 28px;
  }
  
  .logo-image {
    width: 24px;
    height: 24px;
  }
  
  .xp-title {
    font-size: 15px;
  }
}
</style>
