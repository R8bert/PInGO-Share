<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between">
      <div>
        <h3 class="text-2xl font-bold transition-colors duration-300"
            :style="{ color: isDark ? '#f9fafb' : '#111827' }">Reverse Share Tokens</h3>
        <p class="text-lg opacity-80 transition-colors duration-300"
           :style="{ color: isDark ? '#9ca3af' : '#4b5563' }">Allow others to upload files to your account without registering</p>
      </div>
      <button
        @click="showCreateToken = true"
        class="relative overflow-hidden group px-6 py-3 bg-gradient-to-r from-green-600 to-emerald-600 hover:from-green-700 hover:to-emerald-700 text-white font-semibold rounded-2xl transition-all duration-300 transform hover:scale-105 shadow-lg hover:shadow-2xl"
      >
        <div class="absolute inset-0 bg-gradient-to-r from-white/20 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
        <div class="relative flex items-center">
          <IconPlus class="w-5 h-5 mr-2" />
          Create Token
        </div>
      </button>
    </div>

    <!-- Create Token Form -->
    <div v-if="showCreateToken" class="relative group">
      <div class="absolute inset-0 bg-gradient-to-r from-green-600/20 to-emerald-600/20 rounded-2xl blur-xl group-hover:blur-2xl transition-all duration-300"></div>
      <div class="relative backdrop-blur-xl border rounded-2xl p-6 transition-all duration-300"
           :style="{ 
             backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
             borderColor: isDark ? 'rgba(34, 197, 94, 0.3)' : 'rgba(34, 197, 94, 0.2)'
           }">
        <h4 class="text-xl font-semibold mb-6 transition-colors duration-300"
            :style="{ color: isDark ? '#f9fafb' : '#111827' }">Create New Token</h4>
        <form @submit.prevent="createReverseToken" class="space-y-6">
          <div>
            <label class="block text-sm font-medium mb-2 transition-colors duration-300"
                   :style="{ color: isDark ? '#e5e7eb' : '#374151' }">Token Name</label>
            <input
              v-model="newToken.name"
              type="text"
              required
              placeholder="e.g., Client Uploads"
              class="w-full px-4 py-3 border rounded-xl focus:ring-2 focus:ring-green-500 focus:border-green-500 transition-all duration-300 backdrop-blur-sm"
              :style="{ 
                backgroundColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(255,255,255,0.8)',
                borderColor: isDark ? 'rgba(34, 197, 94, 0.3)' : 'rgba(34, 197, 94, 0.2)',
                color: isDark ? '#f9fafb' : '#111827'
              }"
            />
          </div>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium mb-2 transition-colors duration-300"
                     :style="{ color: isDark ? '#e5e7eb' : '#374151' }">Max Uses</label>
              <select
                v-model.number="newToken.max_uses"
                class="w-full px-4 py-3 border rounded-xl focus:ring-2 focus:ring-green-500 focus:border-green-500 transition-all duration-300 backdrop-blur-sm"
                :style="{ 
                  backgroundColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(255,255,255,0.8)',
                  borderColor: isDark ? 'rgba(34, 197, 94, 0.3)' : 'rgba(34, 197, 94, 0.2)',
                  color: isDark ? '#f9fafb' : '#111827'
                }"
              >
                <option :value="-1">Unlimited</option>
                <option :value="1">1 use</option>
                <option :value="5">5 uses</option>
                <option :value="10">10 uses</option>
                <option :value="25">25 uses</option>
                <option :value="100">100 uses</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium mb-2 transition-colors duration-300"
                     :style="{ color: isDark ? '#e5e7eb' : '#374151' }">Expires In</label>
              <select
                v-model="newToken.expires_in"
                class="w-full px-4 py-3 border rounded-xl focus:ring-2 focus:ring-green-500 focus:border-green-500 transition-all duration-300 backdrop-blur-sm"
                :style="{ 
                  backgroundColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(255,255,255,0.8)',
                  borderColor: isDark ? 'rgba(34, 197, 94, 0.3)' : 'rgba(34, 197, 94, 0.2)',
                  color: isDark ? '#f9fafb' : '#111827'
                }"
              >
                <option value="">Never</option>
                <option value="7days">7 days</option>
                <option value="1month">1 month</option>
                <option value="6months">6 months</option>
                <option value="1year">1 year</option>
              </select>
            </div>
          </div>
          <div class="flex space-x-3">
            <button
              type="submit"
              :disabled="isLoading"
              class="relative overflow-hidden group px-6 py-3 bg-gradient-to-r from-green-600 to-emerald-600 hover:from-green-700 hover:to-emerald-700 text-white font-semibold rounded-xl transition-all duration-300 disabled:opacity-50 disabled:cursor-not-allowed transform hover:scale-105"
            >
              <div class="absolute inset-0 bg-gradient-to-r from-white/20 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
              <span class="relative">Create Token</span>
            </button>
            <button
              type="button"
              @click="showCreateToken = false; resetNewToken()"
              class="px-6 py-3 backdrop-blur-sm border rounded-xl transition-all duration-300 hover:scale-105"
              :style="{
                backgroundColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(255,255,255,0.8)',
                borderColor: isDark ? 'rgba(255,255,255,0.2)' : 'rgba(0,0,0,0.1)',
                color: isDark ? '#f9fafb' : '#111827'
              }"
            >
              Cancel
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Tokens List -->
    <div v-if="reverseTokens.length === 0" class="text-center py-20">
      <div class="relative inline-block mb-6">
        <div class="w-20 h-20 rounded-2xl bg-gradient-to-r from-green-500 to-emerald-500 flex items-center justify-center">
          <IconShare class="w-10 h-10 text-white" />
        </div>
      </div>
      <h3 class="text-2xl font-bold mb-4 transition-colors duration-300"
          :style="{ color: isDark ? '#f9fafb' : '#111827' }">No reverse share tokens</h3>
      <p class="text-lg mb-8 opacity-80 transition-colors duration-300"
         :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">Create a token to allow others to upload files to your account.</p>
    </div>

    <div v-else class="space-y-6">
      <div
        v-for="token in reverseTokens"
        :key="token.id"
        class="group relative">
        
        <!-- Card Background Glow -->
        <div class="absolute inset-0 bg-gradient-to-r from-green-600/20 to-emerald-600/20 opacity-0 group-hover:opacity-100 transition-opacity duration-300 rounded-2xl blur-xl"></div>
        
        <!-- Main Card -->
        <div class="relative backdrop-blur-xl border rounded-2xl p-6 transition-all duration-300"
             :style="{ 
               backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
               borderColor: isDark ? 'rgba(34, 197, 94, 0.3)' : 'rgba(34, 197, 94, 0.2)'
             }">
          <div class="flex items-start justify-between">
            <div class="flex-1">
              <div class="flex items-center space-x-3 mb-4">
                <div class="w-10 h-10 rounded-xl bg-gradient-to-r from-green-500 to-emerald-500 flex items-center justify-center">
                  <IconShare class="w-5 h-5 text-white" />
                </div>
                <div>
                  <h4 class="text-lg font-semibold transition-colors duration-300"
                      :style="{ color: isDark ? '#f9fafb' : '#111827' }">{{ token.name }}</h4>
                  <div class="flex items-center space-x-4 text-sm opacity-80"
                       :style="{ color: isDark ? '#9ca3af' : '#4b5563' }">
                    <span class="flex items-center space-x-1">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v4"></path>
                      </svg>
                      <span>{{ token.used_count }}{{ token.max_uses !== -1 ? ` / ${token.max_uses}` : ' (unlimited)' }} uses</span>
                    </span>
                    <span v-if="token.expires_at" class="flex items-center space-x-1">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                      </svg>
                      <span>Expires {{ formatDate(token.expires_at) }}</span>
                    </span>
                    <span v-else class="flex items-center space-x-1">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
                      </svg>
                      <span>Never expires</span>
                    </span>
                  </div>
                </div>
              </div>
              
              <!-- URL Display -->
              <div class="p-4 rounded-xl backdrop-blur-sm border"
                   :style="{ 
                     backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.6)',
                     borderColor: isDark ? 'rgba(34, 197, 94, 0.2)' : 'rgba(34, 197, 94, 0.1)'
                   }">
                <div class="flex items-center justify-between">
                  <div class="flex-1 mr-4">
                    <p class="text-xs font-medium mb-1 opacity-70"
                       :style="{ color: isDark ? '#9ca3af' : '#4b5563' }">Upload URL</p>
                    <p class="font-mono text-sm break-all"
                       :style="{ color: isDark ? '#e5e7eb' : '#111827' }">
                      {{ getBaseUrl() }}/reverse-upload/{{ token.token }}
                    </p>
                  </div>
                  <button
                    @click="$emit('copyToClipboard', `${getBaseUrl()}/reverse-upload/${token.token}`)"
                    class="p-2 rounded-lg backdrop-blur-sm border transition-all duration-200 hover:scale-105"
                    :style="{
                      backgroundColor: isDark ? 'rgba(34, 197, 94, 0.1)' : 'rgba(34, 197, 94, 0.05)',
                      borderColor: 'rgba(34, 197, 94, 0.2)',
                      color: '#22c55e'
                    }"
                    title="Copy upload URL">
                    <IconClipboardDocument class="w-4 h-4" />
                  </button>
                </div>
              </div>
            </div>
            
            <!-- Delete Button -->
            <button
              @click="$emit('deleteReverseToken', token.id)"
              class="p-2 rounded-lg backdrop-blur-sm border transition-all duration-200 hover:scale-105 ml-4"
              :style="{
                backgroundColor: isDark ? 'rgba(239, 68, 68, 0.1)' : 'rgba(239, 68, 68, 0.05)',
                borderColor: 'rgba(239, 68, 68, 0.2)',
                color: '#ef4444'
              }"
              title="Delete token">
              <IconTrash class="w-5 h-5" />
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useTheme } from '../../../composables/useTheme'
import IconPlus from '~icons/solar/add-circle-bold'
import IconShare from '~icons/solar/share-bold'
import IconClipboardDocument from '~icons/solar/clipboard-bold'
import IconTrash from '~icons/solar/trash-bin-minimalistic-bold'

const { isDark } = useTheme()

// Types
interface ReverseToken {
  id: number
  token: string
  name: string
  used_count: number
  max_uses: number
  created_at: string
  expires_at: string | null
}

// Props
interface Props {
  reverseTokens: ReverseToken[]
  isLoading: boolean
}

defineProps<Props>()

// Emits
const emit = defineEmits<{
  createReverseToken: [tokenData: any]
  deleteReverseToken: [tokenId: number]
  copyToClipboard: [text: string]
}>()

// Local state
const showCreateToken = ref(false)
const newToken = ref({
  name: '',
  max_uses: -1,
  expires_in: ''
})

// Helper functions
const formatDate = (dateString: string | null) => {
  if (!dateString) return 'Unknown'
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

const getBaseUrl = () => {
  return window.location.origin
}

const resetNewToken = () => {
  newToken.value = {
    name: '',
    max_uses: -1,
    expires_in: ''
  }
}

// Event handlers
const createReverseToken = () => {
  if (!newToken.value.name || newToken.value.name.trim() === '') {
    alert('Please enter a token name')
    return
  }

  const tokenData = {
    name: newToken.value.name.trim(),
    max_uses: parseInt(newToken.value.max_uses.toString()),
    expires_in: newToken.value.expires_in
  }

  emit('createReverseToken', tokenData)
  showCreateToken.value = false
  resetNewToken()
}
</script>
