<template>
  <div :class="['min-h-screen flex items-center justify-center p-4', isDark ? 'bg-gray-900' : 'bg-gray-100']">
    <!-- Main Container -->
    <div :class="['w-full max-w-6xl rounded-3xl shadow-2xl overflow-hidden flex min-h-[600px]', isDark ? 'bg-gray-800' : 'bg-white']">
      <!-- Left Section - Blue Gradient Panel -->
      <div class="w-[45%] bg-gradient-to-br from-blue-600 via-indigo-600 to-cyan-500 relative flex flex-col justify-between p-12 text-white overflow-hidden">
        <!-- Subtle animated background elements -->
        <div class="absolute top-0 left-0 w-full h-full">
          <div class="absolute top-10 right-10 w-32 h-32 bg-white/5 rounded-full blur-xl animate-float"></div>
          <div class="absolute bottom-20 left-10 w-24 h-24 bg-cyan-400/10 rounded-full blur-lg animate-float" style="animation-delay: 1s;"></div>
          <div class="absolute top-1/2 right-1/4 w-16 h-16 bg-indigo-400/10 rounded-full blur-md animate-float" style="animation-delay: 2s;"></div>
        </div>

        <!-- Top Content -->
        <div class="space-y-6 relative z-10">
          <div class="text-sm font-light opacity-90 animate-fade-in">Join the community</div>
          <h1 class="text-3xl font-bold leading-tight animate-fade-in" style="animation-delay: 0.2s;">
            Start with RootDrop
          </h1>
        </div>

        <!-- Bottom Open Source Section -->
        <div class="space-y-4 relative z-10">
          <div class="text-sm font-light opacity-90 animate-fade-in" style="animation-delay: 0.4s;">Open Source Project</div>
          <div class="flex items-center justify-between opacity-80 animate-fade-in" style="animation-delay: 0.6s;">
            <div class="w-8 h-8 bg-white/20 rounded-lg flex items-center justify-center hover:bg-white/30 transition-all duration-300 hover:scale-110 hover:rotate-6 cursor-pointer group animate-icon-bounce" style="animation-delay: 0.8s;">
              <IconDiscord class="w-5 h-5 group-hover:scale-110 transition-transform duration-200" />
            </div>
            <div class="w-8 h-8 bg-white/20 rounded-lg flex items-center justify-center hover:bg-white/30 transition-all duration-300 hover:scale-110 hover:-rotate-6 cursor-pointer group animate-icon-bounce" style="animation-delay: 0.9s;">
              <IconInstagram class="w-5 h-5 group-hover:scale-110 transition-transform duration-200" />
            </div>
            <div class="w-8 h-8 bg-white/20 rounded-lg flex items-center justify-center hover:bg-white/30 transition-all duration-300 hover:scale-110 hover:rotate-6 cursor-pointer group animate-icon-bounce" style="animation-delay: 1s;">
              <IconSpotify class="w-5 h-5 group-hover:scale-110 transition-transform duration-200" />
            </div>
            <div class="w-8 h-8 bg-white/20 rounded-lg flex items-center justify-center hover:bg-white/30 transition-all duration-300 hover:scale-110 hover:-rotate-6 cursor-pointer group animate-icon-bounce" style="animation-delay: 1.1s;">
              <IconYoutube class="w-5 h-5 group-hover:scale-110 transition-transform duration-200" />
            </div>
            <div class="w-8 h-8 bg-white/20 rounded-lg flex items-center justify-center hover:bg-white/30 transition-all duration-300 hover:scale-110 hover:rotate-6 cursor-pointer group animate-icon-bounce" style="animation-delay: 1.2s;">
              <IconTwitter class="w-5 h-5 group-hover:scale-110 transition-transform duration-200" />
            </div>
          </div>
        </div>
      </div>

      <!-- Right Section - Register Form Panel -->
      <div :class="['w-[55%] flex items-center justify-center p-16 rounded-r-3xl', isDark ? 'bg-gray-800' : 'bg-white']">
        <div class="w-full max-w-md space-y-8 animate-slide-in-up">
          <!-- Header -->
          <div class="text-center space-y-2">
            <h2 :class="['text-3xl font-bold animate-fade-in', isDark ? 'text-white' : 'text-gray-900']">Join RootDrop</h2>
            <p :class="['text-sm animate-fade-in', isDark ? 'text-gray-400' : 'text-gray-500']" style="animation-delay: 0.1s;">Create your account to start sharing files securely.</p>
          </div>

          <!-- Register Form -->
          <form @submit.prevent="handleRegister" class="space-y-6">
            <!-- Name Field -->
            <div class="space-y-2 animate-fade-in" style="animation-delay: 0.2s;">
              <input
                v-model="username"
                type="text"
                required
                :class="[
                  'w-full px-4 py-3 border rounded-xl focus:outline-none focus:ring-2 transition-all duration-300 hover:border-gray-400 hover:shadow-sm hover:shadow-blue-500/10',
                  isDark 
                    ? 'bg-gray-700 border-gray-600 text-white focus:ring-blue-500 focus:border-blue-500 placeholder-gray-400' 
                    : 'bg-white border-gray-300 text-gray-900 focus:ring-blue-500 focus:border-blue-500'
                ]"
                placeholder="Choose a username"
              />
            </div>

            <!-- Email Field -->
            <div class="space-y-2 animate-fade-in" style="animation-delay: 0.3s;">
              <input
                v-model="email"
                type="email"
                required
                :class="[
                  'w-full px-4 py-3 border rounded-xl focus:outline-none focus:ring-2 transition-all duration-300 hover:border-gray-400 hover:shadow-sm hover:shadow-blue-500/10',
                  isDark 
                    ? 'bg-gray-700 border-gray-600 text-white focus:ring-blue-500 focus:border-blue-500 placeholder-gray-400' 
                    : 'bg-white border-gray-300 text-gray-900 focus:ring-blue-500 focus:border-blue-500'
                ]"
                placeholder="workmail@gmail.com"
              />
            </div>

            <!-- Password Field -->
            <div class="space-y-2 animate-fade-in" style="animation-delay: 0.4s;">
              <input
                v-model="password"
                :type="showPassword ? 'text' : 'password'"
                required
                minlength="6"
                :class="[
                  'w-full px-4 py-3 border rounded-xl focus:outline-none focus:ring-2 transition-all duration-300 hover:border-gray-400 hover:shadow-sm hover:shadow-blue-500/10',
                  isDark 
                    ? 'bg-gray-700 border-gray-600 text-white focus:ring-blue-500 focus:border-blue-500 placeholder-gray-400' 
                    : 'bg-white border-gray-300 text-gray-900 focus:ring-blue-500 focus:border-blue-500'
                ]"
                placeholder="••••••••"
              />
            </div>

            <!-- Confirm Password Field -->
            <div class="space-y-2 animate-fade-in" style="animation-delay: 0.5s;">
              <input
                v-model="confirmPassword"
                :type="showConfirmPassword ? 'text' : 'password'"
                required
                :class="[
                  'w-full px-4 py-3 border rounded-xl focus:outline-none focus:ring-2 transition-all duration-300 hover:border-gray-400 hover:shadow-sm hover:shadow-blue-500/10',
                  isDark 
                    ? 'bg-gray-700 border-gray-600 text-white focus:ring-blue-500 focus:border-blue-500 placeholder-gray-400' 
                    : 'bg-white border-gray-300 text-gray-900 focus:ring-blue-500 focus:border-blue-500'
                ]"
                placeholder="••••••••"
              />
            </div>

            <!-- Error Message -->
            <div v-if="errorMessage" :class="['p-4 border rounded-xl animate-fade-in shadow-sm', isDark ? 'bg-red-900/50 border-red-700 text-red-200' : 'bg-red-50 border-red-200 text-red-600']">
              <p class="text-sm">{{ errorMessage }}</p>
            </div>

            <!-- Submit Button -->
            <button
              type="submit"
              :disabled="isLoading"
              class="w-full bg-gradient-to-r from-blue-600 to-indigo-600 hover:from-blue-700 hover:to-indigo-700 disabled:from-gray-400 disabled:to-gray-500 text-white font-semibold py-3 px-6 rounded-xl transition-all duration-300 transform hover:scale-[1.02] hover:shadow-lg hover:shadow-blue-500/25 active:scale-[0.98] disabled:cursor-not-allowed disabled:transform-none animate-fade-in"
              style="animation-delay: 0.6s;"
            >
              <span v-if="isLoading">Creating Account...</span>
              <span v-else>Create Account</span>
            </button>
          </form>

          <!-- Login Link -->
          <div class="text-center animate-fade-in" style="animation-delay: 0.7s;">
            <p :class="isDark ? 'text-gray-400' : 'text-gray-600'">
              Already have an account?
              <button
                @click="switchToLogin"
                :class="['font-semibold ml-1 transition-all duration-200 hover:underline', isDark ? 'text-blue-400 hover:text-blue-300' : 'text-blue-600 hover:text-blue-700']"
              >
                Sign in
              </button>
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template><script setup lang="ts">
import { ref } from 'vue'
import { useAuth } from '../../composables/useAuth'
import { useTheme } from '../../composables/useTheme'
import { useRouter } from 'vue-router'
import IconDiscord from '~icons/mdi/discord'
import IconInstagram from '~icons/mdi/instagram'
import IconSpotify from '~icons/mdi/spotify'
import IconYoutube from '~icons/mdi/youtube'
import IconTwitter from '~icons/mdi/twitter'

const emit = defineEmits<{
  switchToLogin: []
}>()

const { register, isLoading } = useAuth()
const { isDark } = useTheme()
const router = useRouter()

const username = ref('')
const email = ref('')
const password = ref('')
const confirmPassword = ref('')
const errorMessage = ref('')
const showPassword = ref(false)
const showConfirmPassword = ref(false)

const handleRegister = async () => {
  errorMessage.value = ''

  if (!username.value || !email.value || !password.value || !confirmPassword.value) {
    errorMessage.value = 'Please fill in all fields'
    return
  }

  if (password.value !== confirmPassword.value) {
    errorMessage.value = 'Passwords do not match'
    return
  }

  if (password.value.length < 6) {
    errorMessage.value = 'Password must be at least 6 characters long'
    return
  }

  const result = await register(username.value, email.value, password.value)

  if (result.success) {
    router.push('/')
  } else {
    errorMessage.value = result.message
  }
}

const switchToLogin = () => {
  emit('switchToLogin')
}
</script>

<style scoped>
@keyframes spin-slow {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.animate-spin-slow {
  animation: spin-slow 8s linear infinite;
}

@keyframes fade-in {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes slide-in-up {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes float {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-10px); }
}

@keyframes icon-bounce {
  0% {
    opacity: 0;
    transform: scale(0.5) translateY(20px);
  }
  60% {
    opacity: 0.8;
    transform: scale(1.1) translateY(-5px);
  }
  100% {
    opacity: 0.8;
    transform: scale(1) translateY(0);
  }
}

.animate-fade-in {
  animation: fade-in 0.6s ease-out forwards;
  opacity: 0;
}

.animate-slide-in-up {
  animation: slide-in-up 0.8s ease-out forwards;
  opacity: 0;
}

.animate-float {
  animation: float 4s ease-in-out infinite;
}

.animate-icon-bounce {
  animation: icon-bounce 0.6s cubic-bezier(0.68, -0.55, 0.265, 1.55) forwards;
  opacity: 0;
}
</style>
