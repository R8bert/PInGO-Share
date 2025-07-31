<template>
  <div class="mb-6">
    <div class="flex items-center justify-between mb-4">
      <div>
        <h3 class="text-lg font-medium transition-colors duration-300"
            :style="{ color: isDark ? '#f9fafb' : '#111827' }">Reverse Share Tokens</h3>
        <p class="text-sm transition-colors duration-300"
           :style="{ color: isDark ? '#9ca3af' : '#4b5563' }">Allow others to upload files to your account without registering</p>
      </div>
      <button
        @click="showCreateToken = true"
        class="inline-flex items-center px-4 py-2 bg-green-600 hover:bg-green-700 text-white font-medium rounded-lg transition-colors"
      >
        <PlusIcon class="w-4 h-4 mr-2" />
        Create Token
      </button>
    </div>

    <!-- Create Token Form -->
    <div v-if="showCreateToken" 
         class="rounded-xl p-6 border mb-6 transition-colors duration-300"
         :style="{ 
           backgroundColor: isDark ? '#1f2937' : '#f9fafb',
           borderColor: isDark ? '#374151' : '#e5e7eb'
         }">
      <h4 class="text-lg font-medium mb-4 transition-colors duration-300"
          :style="{ color: isDark ? '#f9fafb' : '#111827' }">Create New Token</h4>
      <form @submit.prevent="createReverseToken" class="space-y-4">
        <div>
          <label class="block text-sm font-medium mb-2 transition-colors duration-300"
                 :style="{ color: isDark ? '#e5e7eb' : '#374151' }">Token Name</label>
          <input
            v-model="newToken.name"
            type="text"
            required
            placeholder="e.g., Client Uploads"
            class="w-full px-3 py-2 border rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500 transition-colors duration-300"
            :style="{ 
              backgroundColor: isDark ? '#374151' : '#ffffff',
              borderColor: isDark ? '#4b5563' : '#d1d5db',
              color: isDark ? '#f9fafb' : '#111827'
            }"
          />
        </div>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium mb-2 transition-colors duration-300"
                   :style="{ color: isDark ? '#e5e7eb' : '#374151' }">Max Uses</label>
            <select
              v-model.number="newToken.max_uses"
              class="w-full px-3 py-2 border rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500 transition-colors duration-300"
              :style="{ 
                backgroundColor: isDark ? '#374151' : '#ffffff',
                borderColor: isDark ? '#4b5563' : '#d1d5db',
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
              class="w-full px-3 py-2 border rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500 transition-colors duration-300"
              :style="{ 
                backgroundColor: isDark ? '#374151' : '#ffffff',
                borderColor: isDark ? '#4b5563' : '#d1d5db',
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
            class="inline-flex items-center px-4 py-2 bg-green-600 hover:bg-green-700 text-white font-medium rounded-lg transition-colors disabled:opacity-50"
          >
            Create Token
          </button>
          <button
            type="button"
            @click="showCreateToken = false; resetNewToken()"
            class="inline-flex items-center px-4 py-2 bg-gray-300 hover:bg-gray-400 text-gray-700 font-medium rounded-lg transition-colors"
          >
            Cancel
          </button>
        </div>
      </form>
    </div>

    <!-- Tokens List -->
    <div v-if="reverseTokens.length === 0" class="text-center py-12">
      <ShareIcon class="w-16 h-16 mx-auto mb-4 transition-colors duration-300"
                 :style="{ color: isDark ? '#4b5563' : '#d1d5db' }" />
      <h3 class="text-lg font-medium mb-2 transition-colors duration-300"
          :style="{ color: isDark ? '#f9fafb' : '#111827' }">No reverse share tokens</h3>
      <p class="mb-6 transition-colors duration-300"
         :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">Create a token to allow others to upload files to your account.</p>
    </div>

    <div v-else class="space-y-4">
      <div
        v-for="token in reverseTokens"
        :key="token.id"
        class="rounded-xl p-6 border transition-colors duration-300"
        :style="{ 
          backgroundColor: isDark ? '#1f2937' : '#f9fafb',
          borderColor: isDark ? '#374151' : '#e5e7eb'
        }"
      >
        <div class="flex items-start justify-between">
          <div class="flex-1">
            <h4 class="font-medium mb-2 transition-colors duration-300"
                :style="{ color: isDark ? '#f9fafb' : '#111827' }">{{ token.name }}</h4>
            <div class="space-y-2 text-sm transition-colors duration-300"
                 :style="{ color: isDark ? '#9ca3af' : '#4b5563' }">
              <div class="flex items-center space-x-4">
                <span>Uses: {{ token.used_count }}{{ token.max_uses !== -1 ? ` / ${token.max_uses}` : ' (unlimited)' }}</span>
                <span v-if="token.expires_at">Expires: {{ formatDate(token.expires_at) }}</span>
                <span v-else>Never expires</span>
              </div>
              <div class="rounded p-2 font-mono text-xs break-all transition-colors duration-300"
                   :style="{ 
                     backgroundColor: isDark ? '#374151' : '#f3f4f6',
                     color: isDark ? '#e5e7eb' : '#4b5563'
                   }">
                Upload URL: {{ getBaseUrl() }}/reverse-upload/{{ token.token }}
              </div>
            </div>
          </div>
          <div class="flex items-center space-x-2">
            <button
              @click="copyToClipboard(`${getBaseUrl()}/reverse-upload/${token.token}`)"
              class="p-2 text-gray-400 hover:text-blue-600 hover:bg-blue-50 rounded-lg transition-colors"
              title="Copy upload URL"
            >
              <ClipboardDocumentIcon class="w-5 h-5" />
            </button>
            <button
              @click="deleteReverseToken(token.id)"
              class="p-2 text-gray-400 hover:text-red-600 hover:bg-red-50 rounded-lg transition-colors"
              title="Delete token"
            >
              <TrashIcon class="w-5 h-5" />
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
import {
  ShareIcon,
  PlusIcon,
  ClipboardDocumentIcon,
  TrashIcon
} from '@heroicons/vue/24/outline'

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

const deleteReverseToken = (tokenId: number) => {
  if (!confirm('Are you sure you want to delete this token?')) return
  emit('deleteReverseToken', tokenId)
}

const copyToClipboard = (text: string) => {
  emit('copyToClipboard', text)
}
</script>
