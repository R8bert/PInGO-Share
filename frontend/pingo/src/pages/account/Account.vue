<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-200 via-purple-100 to-yellow-100 px-4">
    <div class="w-full max-w-lg bg-white shadow-2xl rounded-3xl p-8 space-y-8 animate-fade-in">
      <!-- Titlu -->
      <h1 class="text-4xl font-extrabold text-center text-gray-900 tracking-tight">
        Account
      </h1>
      <p class="text-center text-gray-600 text-sm">
        Manage your account details
      </p>

      <!-- Formular cont -->
      <form @submit.prevent="saveAccount" class="space-y-6">
        <!-- Nume utilizator -->
        <div>
          <label for="username" class="block text-gray-700 font-medium">
            Username
          </label>
          <input
            id="username"
            type="text"
            v-model="account.username"
            class="mt-2 w-full border border-gray-300 px-4 py-3 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 outline-none transition-all"
            placeholder="Enter your username"
          />
        </div>

        <!-- Email -->
        <div>
          <label for="email" class="block text-gray-700 font-medium">
            Email
          </label>
          <input
            id="email"
            type="email"
            v-model="account.email"
            class="mt-2 w-full border border-gray-300 px-4 py-3 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 outline-none transition-all"
            placeholder="Enter your email"
          />
        </div>

        <!-- Buton de salvare -->
        <button
          type="submit"
          :disabled="isSaving"
          class="w-full bg-gradient-to-r from-blue-600 to-blue-700 text-white py-3 rounded-lg font-semibold hover:from-blue-700 hover:to-blue-800 disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-300"
        >
          {{ isSaving ? 'Saving...' : 'Save Account' }}
        </button>
      </form>

      <!-- Mesaj de succes/eroare -->
      <div v-if="message" class="text-center">
        <p :class="message.type === 'success' ? 'text-green-600' : 'text-red-600'">
          {{ message.text }}
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import axios from 'axios'

// Definim tipurile pentru cont
interface Account {
  username: string
  email: string
}

interface Message {
  text: string
  type: 'success' | 'error'
}

// Starea formularului
const account = reactive<Account>({
  username: '',
  email: '',
})

const isSaving = ref(false)
const message = ref<Message | null>(null)

// Funcție pentru salvarea datelor contului
const saveAccount = async () => {
  isSaving.value = true
  message.value = null

  try {
    // Trimite datele contului către backend
    await axios.post('http://localhost:8080/account', account)
    message.value = { text: 'Account details saved successfully!', type: 'success' }
  } catch (error) {
    console.error('Error saving account:', error)
    message.value = { text: 'Failed to save account details. Please try again.', type: 'error' }
  } finally {
    isSaving.value = false
  }
}
</script>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.5s ease-in;
}

@keyframes fadeIn {
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