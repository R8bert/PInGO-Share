<template>
  <div>
    <!-- Loading skeleton -->
    <div v-if="isLoading" class="space-y-6">
      <div class="bg-gray-200 animate-pulse rounded-xl p-6 border">
        <div class="h-6 bg-gray-300 rounded w-1/4 mb-4"></div>
        <div class="space-y-3">
          <div class="h-4 bg-gray-300 rounded w-full"></div>
          <div class="h-4 bg-gray-300 rounded w-3/4"></div>
        </div>
      </div>
    </div>
    
    <!-- User Management Content -->
    <div v-else class="space-y-8">
      <div>
        <h3 class="text-xl font-bold mb-6 transition-colors duration-300"
            :style="{ color: isDark ? '#f9fafb' : '#111827' }">User Management</h3>
        <div class="rounded-xl p-6 border transition-colors duration-300"
             :style="{ 
               backgroundColor: isDark ? '#1a1a1a' : '#ffffff',
               borderColor: isDark ? '#374151' : '#e5e7eb'
             }">
          <div class="overflow-x-auto">
            <table class="min-w-full divide-y transition-colors duration-300"
                   :style="{ borderColor: isDark ? '#374151' : '#e5e7eb' }">
              <thead class="transition-colors duration-300"
                     :style="{ backgroundColor: isDark ? '#1f2937' : '#f9fafb' }">
                <tr>
                  <th class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider transition-colors duration-300"
                      :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">User</th>
                  <th class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider transition-colors duration-300"
                      :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">Uploads</th>
                  <th class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider transition-colors duration-300"
                      :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">Storage</th>
                  <th class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider transition-colors duration-300"
                      :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">Status</th>
                  <th class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider transition-colors duration-300"
                      :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">Actions</th>
                </tr>
              </thead>
              <tbody class="divide-y transition-colors duration-300"
                     :style="{ 
                       backgroundColor: isDark ? '#1a1a1a' : '#ffffff',
                       borderColor: isDark ? '#374151' : '#e5e7eb'
                     }">
                <tr v-for="user in adminUsers" :key="user.id" class="hover:bg-gray-50">
                  <td class="px-6 py-4 whitespace-nowrap">
                    <div class="flex items-center">
                      <div class="w-8 h-8 rounded-full overflow-hidden bg-gradient-to-r from-blue-500 to-indigo-600 flex items-center justify-center mr-3">
                        <img 
                          v-if="user.avatar" 
                          :src="`http://localhost:8080${user.avatar}`" 
                          :alt="user.username"
                          class="w-full h-full object-cover"
                          @error="handleAvatarError"
                        />
                        <UserIcon v-else class="w-4 h-4 text-white" />
                      </div>
                      <div>
                        <div class="text-sm font-medium text-gray-900">{{ user.username }}</div>
                        <div class="text-sm text-gray-500">{{ user.email }}</div>
                      </div>
                    </div>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ user.uploadCount || 0 }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatBytes(user.storageUsed || 0) }}</td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <span :class="[
                      'inline-flex px-2 py-1 text-xs font-semibold rounded-full',
                      user.isBlocked ? 'bg-red-100 text-red-800' : 'bg-green-100 text-green-800'
                    ]">
                      {{ user.isBlocked ? 'Blocked' : 'Active' }}
                    </span>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                    <button
                      @click="toggleUserBlock(user.id, !user.isBlocked)"
                      :class="[
                        'px-3 py-1 rounded-md text-sm font-medium transition-colors',
                        user.isBlocked 
                          ? 'bg-green-600 hover:bg-green-700 text-white' 
                          : 'bg-red-600 hover:bg-red-700 text-white'
                      ]"
                    >
                      {{ user.isBlocked ? 'Unblock' : 'Block' }}
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useTheme } from '../../../composables/useTheme'
import { UserIcon } from '@heroicons/vue/24/outline'

const { isDark } = useTheme()

// Props
interface Props {
  adminUsers: any[]
  isLoading: boolean
}

defineProps<Props>()

// Emits
const emit = defineEmits<{
  toggleUserBlock: [userId: number, isBlocked: boolean]
}>()

// Helper functions
const formatBytes = (bytes: number) => {
  if (bytes === 0) return '0 Bytes'
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

const handleAvatarError = () => {
  console.error('Failed to load avatar')
}

// Event handlers
const toggleUserBlock = (userId: number, isBlocked: boolean) => {
  emit('toggleUserBlock', userId, isBlocked)
}
</script>
