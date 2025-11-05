<template>
  <div>
    <!-- Loading skeleton -->
    <div v-if="isLoading" class="space-y-4">
      <div class="animate-pulse rounded-lg p-6 border transition-colors duration-300"
           :style="{ 
             backgroundColor: isDark ? '#111827' : '#ffffff',
             borderColor: isDark ? '#374151' : '#e5e7eb'
           }">
        <div class="h-6 rounded w-1/4 mb-4 transition-colors duration-300"
             :style="{ backgroundColor: isDark ? '#374151' : '#d1d5db' }"></div>
        <div class="space-y-3">
          <div class="h-4 rounded w-full transition-colors duration-300"
               :style="{ backgroundColor: isDark ? '#374151' : '#d1d5db' }"></div>
          <div class="h-4 rounded w-3/4 transition-colors duration-300"
               :style="{ backgroundColor: isDark ? '#374151' : '#d1d5db' }"></div>
        </div>
      </div>
    </div>
    
    <!-- User Management Content -->
    <div v-else class="space-y-6">
      <!-- Header with stats -->
      <div class="flex items-center justify-between">
        <h3 class="text-xl font-bold transition-colors duration-300"
            :style="{ color: isDark ? '#f9fafb' : '#111827' }">
          User Management
        </h3>
        <div class="flex items-center gap-4 text-sm">
          <div class="px-3 py-1.5 rounded-lg border transition-colors duration-300"
               :style="{ 
                 backgroundColor: isDark ? '#1f2937' : '#f9fafb',
                 borderColor: isDark ? '#374151' : '#e5e7eb',
                 color: isDark ? '#d1d5db' : '#6b7280'
               }">
            Total Users: {{ adminUsers.length }}
          </div>
          <div class="px-3 py-1.5 rounded-lg border transition-colors duration-300"
               :style="{ 
                 backgroundColor: isDark ? '#1f2937' : '#f9fafb',
                 borderColor: isDark ? '#374151' : '#e5e7eb',
                 color: isDark ? '#d1d5db' : '#6b7280'
               }">
            Active: {{ adminUsers.filter(u => !u.isBlocked).length }}
          </div>
        </div>
      </div>

      <!-- Users Grid -->
      <div class="grid gap-4">
        <div v-for="user in adminUsers" :key="user.id" 
             class="rounded-lg p-6 border hover:shadow-lg transition-all duration-300 hover:scale-[1.02]"
             :style="{ 
               backgroundColor: isDark ? '#111827' : '#ffffff',
               borderColor: isDark ? '#374151' : '#e5e7eb'
             }">
          <div class="flex items-center justify-between">
            <!-- User Info -->
            <div class="flex items-center space-x-4">
              <!-- Avatar -->
              <div class="relative">
                <div class="w-12 h-12 rounded-full overflow-hidden bg-gradient-to-r from-blue-500 to-indigo-600 flex items-center justify-center shadow-lg">
                  <img 
                    v-if="user.avatar" 
                    :src="getAssetUrl(user.avatar)" 
                    :alt="user.username"
                    class="w-full h-full object-cover"
                    @error="handleAvatarError"
                  />
                  <IconUser v-else class="w-6 h-6 text-white" />
                </div>
                <!-- Status indicator -->
                <div class="absolute -bottom-1 -right-1 w-4 h-4 rounded-full border-2 transition-colors duration-300"
                     :class="user.isBlocked 
                       ? 'bg-red-500 border-red-200' 
                       : 'bg-green-500 border-green-200'"
                     :style="{ borderColor: isDark ? '#111827' : '#ffffff' }">
                </div>
              </div>

              <!-- User Details -->
              <div>
                <h4 class="font-semibold text-base transition-colors duration-300"
                    :style="{ color: isDark ? '#f9fafb' : '#111827' }">
                  {{ user.username }}
                </h4>
                <p class="text-sm transition-colors duration-300"
                   :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">
                  {{ user.email }}
                </p>
                <div class="flex items-center gap-4 mt-1 text-xs">
                  <span class="flex items-center gap-1 transition-colors duration-300"
                        :style="{ color: isDark ? '#d1d5db' : '#4b5563' }">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"></path>
                    </svg>
                    {{ user.uploadCount || 0 }} uploads
                  </span>
                  <span class="flex items-center gap-1 transition-colors duration-300"
                        :style="{ color: isDark ? '#d1d5db' : '#4b5563' }">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 19a2 2 0 01-2-2V7a2 2 0 012-2h4l2 2h4a2 2 0 012 2v1M5 19h14a2 2 0 002-2v-5a2 2 0 00-2-2H9a2 2 0 00-2 2v5a2 2 0 01-2 2z"></path>
                    </svg>
                    {{ formatBytes(user.storageUsed || 0) }}
                  </span>
                </div>
              </div>
            </div>

            <!-- Actions -->
            <div class="flex items-center gap-3">
              <!-- Admin Status Badge -->
              <span class="inline-flex items-center px-3 py-1 rounded-full text-xs font-semibold transition-colors duration-300"
                    :style="user.is_admin 
                      ? { 
                          backgroundColor: isDark ? '#7c3aed' : '#ddd6fe',
                          color: isDark ? '#e9d5ff' : '#7c3aed'
                        }
                      : { 
                          backgroundColor: isDark ? '#374151' : '#f3f4f6',
                          color: isDark ? '#d1d5db' : '#6b7280'
                        }">
                {{ user.is_admin ? 'Admin' : 'User' }}
              </span>

              <!-- Status Badge -->
              <span class="inline-flex items-center px-3 py-1 rounded-full text-xs font-semibold transition-colors duration-300"
                    :style="user.isBlocked 
                      ? { 
                          backgroundColor: isDark ? '#7f1d1d' : '#fecaca',
                          color: isDark ? '#fecaca' : '#7f1d1d'
                        }
                      : { 
                          backgroundColor: isDark ? '#14532d' : '#dcfce7',
                          color: isDark ? '#bbf7d0' : '#14532d'
                        }">
                {{ user.isBlocked ? 'Blocked' : 'Active' }}
              </span>

              <!-- Action Buttons -->
              <div class="flex gap-2">
                <!-- Promote Button -->
                <button
                  v-if="!user.is_admin && currentUser?.id === 1"
                  @click="promoteUser(user.id, true)"
                  class="px-3 py-2 rounded-lg text-xs font-semibold transition-all duration-300 hover:scale-105 hover:shadow-md hover:brightness-110"
                  :style="{ 
                    backgroundColor: isDark ? '#7c3aed' : '#8b5cf6',
                    color: '#ffffff'
                  }"
                >
                  Promote to Admin
                </button>

                <!-- Demote Button -->
                <button
                  v-if="user.is_admin && currentUser?.id === 1"
                  @click="promoteUser(user.id, false)"
                  class="px-3 py-2 rounded-lg text-xs font-semibold transition-all duration-300 hover:scale-105 hover:shadow-md hover:brightness-110"
                  :style="{ 
                    backgroundColor: isDark ? '#f59e0b' : '#f97316',
                    color: '#ffffff'
                  }"
                >
                  Demote from Admin
                </button>

                <!-- Block/Unblock Button -->
                <button
                  @click="toggleUserBlock(user.id, !user.isBlocked)"
                  class="px-3 py-2 rounded-lg text-xs font-semibold transition-all duration-300 hover:scale-105 hover:shadow-md hover:brightness-110"
                  :style="user.isBlocked 
                    ? { 
                        backgroundColor: isDark ? '#059669' : '#10b981',
                        color: '#ffffff'
                      }
                    : { 
                        backgroundColor: isDark ? '#dc2626' : '#ef4444',
                        color: '#ffffff'
                      }"
                >
                  {{ user.isBlocked ? 'Unblock' : 'Block' }}
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-if="adminUsers.length === 0" class="text-center py-12">
        <IconUser class="w-16 h-16 mx-auto mb-4 transition-colors duration-300"
                  :style="{ color: isDark ? '#4b5563' : '#d1d5db' }" />
        <h3 class="text-lg font-medium mb-2 transition-colors duration-300"
            :style="{ color: isDark ? '#f9fafb' : '#111827' }">
          No users found
        </h3>
        <p class="transition-colors duration-300"
           :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">
          Users will appear here once they register.
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useAuth } from '../../../composables/useAuth'
import { useTheme } from '../../../composables/useTheme'
import { getAssetUrl } from '../../../utils/apiUtils'
import IconUser from '~icons/solar/user-bold'

const { adminUsers, fetchAdminUsers, toggleUserBlock, promoteUser, user: currentUser } = useAuth()
const { isDark } = useTheme()
const isLoading = ref(true)

onMounted(async () => {
  try {
    await fetchAdminUsers()
  } finally {
    isLoading.value = false
  }
})

const formatBytes = (bytes: number): string => {
  if (bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

const handleAvatarError = (event: Event) => {
  const target = event.target as HTMLImageElement
  target.style.display = 'none'
}
</script>
