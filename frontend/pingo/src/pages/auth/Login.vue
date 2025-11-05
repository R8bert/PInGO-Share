<template>
  <div :class="['rounded-r-3xl min-h-screen flex items-center justify-center p-4', isDark ? 'bg-gray-900' : 'bg-gray-100']">
    <!-- Main Container -->
    <div :class="['w-full max-w-6xl rounded-3xl shadow-2xl flex min-h-[600px]', isDark ? 'bg-gray-800' : 'bg-white']">
      <!-- Left Section - Blue Gradient Panel -->
      <div class=" rounded-l-3xl rounded-r-3xl w-[45%] bg-gradient-to-br from-green-700 via-emerald-700 to-green-800 relative flex flex-col justify-between p-12 text-white">
        <!-- Subtle animated background elements -->
        <div class="absolute top-0 left-0 w-full h-full rounded-l-3xl rounded-r-3xl">
          <div class="absolute top-10 right-10 w-32 h-32 bg-white/5 rounded-full blur-xl animate-float"></div>
          <div class="absolute bottom-20 left-10 w-24 h-24 bg-emerald-400/10 rounded-full blur-lg animate-float" style="animation-delay: 1s;"></div>
          <div class="absolute top-1/2 right-1/4 w-16 h-16 bg-green-400/10 rounded-full blur-md animate-float" style="animation-delay: 2s;"></div>
        </div>

        <!-- Top Content -->
        <div class="space-y-6 relative z-10">
          <div class="text-sm font-light opacity-90 animate-fade-in">Share files securely</div>
          <h1 class="text-3xl font-bold leading-tight animate-fade-in" style="animation-delay: 0.2s;">
            Welcome to RootDrop
          </h1>
        </div>

        <!-- Icon Grid -->
        <div class="flex-1 flex items-center justify-center relative z-10 animate-fade-in p-8" style="animation-delay: 0.8s;">
          <div class="grid grid-cols-5 gap-4 max-w-lg">
            <img
              v-for="(icon, index) in icons"
              :key="icon"
              :src="icon"
              :style="{
                animationDelay: `${index * 0.1}s`,
                animationFillMode: 'forwards'
              }"
              class="w-12 h-12 opacity-70 hover:opacity-100 transition-all duration-300 hover:scale-125 transform animate-bounce-in animate-gentle-float animate-pulse-glow"
              alt="icon"
            />
          </div>
        </div>
      </div>

      <!-- Right Section - Login Form Panel -->
      <div :class="['w-[55%] flex items-center justify-center p-16 rounded-r-3xl', isDark ? 'bg-gray-800' : 'bg-white']">
        <div class="w-full max-w-md space-y-8 animate-slide-in-up">
          <!-- Header -->
          <div class="text-center space-y-2">
            <h2 :class="['text-3xl font-bold animate-fade-in', isDark ? 'text-white' : 'text-gray-900']">Sign In to RootDrop</h2>
            <p :class="['text-sm animate-fade-in', isDark ? 'text-gray-400' : 'text-gray-500']" style="animation-delay: 0.1s;">Access your secure file sharing account.</p>
          </div>

          <!-- Login Form -->
          <form @submit.prevent="handleLogin" class="space-y-6">
            <!-- Username or Email Field -->
            <div class="space-y-2 animate-fade-in" style="animation-delay: 0.2s;">
              <input
                v-model="usernameOrEmail"
                type="text"
                required
                :class="[
                  'w-full px-4 py-3 border rounded-xl focus:outline-none focus:ring-2 transition-all duration-300 hover:border-gray-400 hover:shadow-sm hover:shadow-green-500/10',
                  isDark
                    ? 'bg-gray-700 border-gray-600 text-white focus:ring-green-500 focus:border-green-500 placeholder-gray-400'
                    : 'bg-white border-gray-300 text-gray-900 focus:ring-green-500 focus:border-green-500'
                ]"
                placeholder="Username or Email"
              />
            </div>

            <!-- Password Field -->
            <div class="space-y-2 animate-fade-in" style="animation-delay: 0.3s;">
              <input
                v-model="password"
                :type="showPassword ? 'text' : 'password'"
                required
                :class="[
                  'w-full px-4 py-3 border rounded-xl focus:outline-none focus:ring-2 transition-all duration-300 hover:border-gray-400 hover:shadow-sm hover:shadow-green-500/10',
                  isDark
                    ? 'bg-gray-700 border-gray-600 text-white focus:ring-green-500 focus:border-green-500 placeholder-gray-400'
                    : 'bg-white border-gray-300 text-gray-900 focus:ring-green-500 focus:border-green-500'
                ]"
                placeholder="••••••••"
              />
            </div>

            <!-- Error Message -->
            <div v-if="errorMessage" :class="['p-4 border rounded-xl animate-fade-in shadow-sm', isDark ? 'bg-red-900/50 border-red-700 text-red-200' : 'bg-red-50 border-red-200 text-red-600']">
              <p>{{ errorMessage }}</p>
            </div>

            <!-- Submit Button -->
            <button
              type="submit"
              :disabled="isLoading"
              class="w-full bg-gradient-to-r from-green-600 to-emerald-600 hover:from-green-700 hover:to-emerald-700 disabled:from-gray-400 disabled:to-gray-500 text-white font-semibold py-3 px-6 rounded-xl transition-all duration-300 transform hover:scale-[1.02] hover:shadow-lg hover:shadow-green-500/25 active:scale-[0.98] disabled:cursor-not-allowed disabled:transform-none animate-fade-in"
              style="animation-delay: 0.4s;"
            >
              <span v-if="isLoading">Signing In...</span>
              <span v-else>Sign In</span>
            </button>
          </form>

          <!-- Register Link -->
          <div class="text-center animate-fade-in" style="animation-delay: 0.5s;">
            <p :class="isDark ? 'text-gray-400' : 'text-gray-600'">
              Don't have an account?
              <button
                @click="switchToRegister"
                :class="['font-semibold ml-1 transition-all duration-200 hover:underline', isDark ? 'text-green-400 hover:text-green-300' : 'text-green-600 hover:text-green-700']"
              >
                Sign up
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

import defaultMultipleFiles from '../../assets/svg/icons/default_multiple_files.svg'
import fbreader from '../../assets/svg/icons/fbreader.svg'
import filesArchive from '../../assets/svg/icons/files_archive.svg'
import filesImage from '../../assets/svg/icons/files_image.svg'
import filesNoIcon from '../../assets/svg/icons/files_no_icon.svg'
import filesPdf from '../../assets/svg/icons/files_pdf.svg'
import filesUploaded from '../../assets/svg/icons/files_uploaded.svg'
import filesUploaded2 from '../../assets/svg/icons/files_uploaded2.svg'
import filesVideo from '../../assets/svg/icons/files_video.svg'
import filesVideo2 from '../../assets/svg/icons/files_video2.svg'
import folderRoot from '../../assets/svg/icons/folder-root.svg'
import folderTemplates from '../../assets/svg/icons/folder-templates.svg'
import inspector from '../../assets/svg/icons/io.github.nokse22.inspector.svg'
import minitext from '../../assets/svg/icons/io.github.nokse22.minitext.svg'
import multimediaPhotoViewer from '../../assets/svg/icons/multimedia-photo-viewer.svg'
import pptPptxFile from '../../assets/svg/icons/ppt_pptx_file.svg'
import totalFilesIcon from '../../assets/svg/icons/total_files_icon.svg'
import unknownFileFormats from '../../assets/svg/icons/unknown_file_formats.svg'
import wordDocxDoc from '../../assets/svg/icons/word_docx_doc.svg'
import wordDocxDoc2 from '../../assets/svg/icons/word_docx_doc2.svg'

const emit = defineEmits<{
  switchToRegister: []
}>()

const { login, isLoading } = useAuth()
const { isDark } = useTheme()
const router = useRouter()

const usernameOrEmail = ref('')
const password = ref('')
const errorMessage = ref('')
const showPassword = ref(false)

const icons = ref([
  defaultMultipleFiles,
  fbreader,
  filesArchive,
  filesImage,
  filesNoIcon,
  filesPdf,
  filesUploaded,
  filesUploaded2,
  filesVideo,
  filesVideo2,
  folderRoot,
  folderTemplates,
  inspector,
  minitext,
  multimediaPhotoViewer,
  pptPptxFile,
  totalFilesIcon,
  unknownFileFormats,
  wordDocxDoc,
  wordDocxDoc2,
])

const handleLogin = async () => {
  errorMessage.value = ''

  if (!usernameOrEmail.value || !password.value) {
    errorMessage.value = 'Please fill in all fields'
    return
  }

  const result = await login(usernameOrEmail.value, password.value)

  if (result.success) {
    router.push('/')
  } else {
    errorMessage.value = result.message
  }
}

const switchToRegister = () => {
  emit('switchToRegister')
}
</script>

<style scoped>
@keyframes bounce-in {
  0% {
    opacity: 0;
    transform: scale(0.3) translateY(50px);
  }
  50% {
    opacity: 0.8;
    transform: scale(1.1) translateY(-10px);
  }
  70% {
    opacity: 1;
    transform: scale(0.95) translateY(2px);
  }
  100% {
    opacity: 1;
    transform: scale(1) translateY(0);
  }
}

@keyframes gentle-float {
  0%, 100% {
    transform: translateY(0px) rotate(0deg);
  }
  25% {
    transform: translateY(-3px) rotate(0.5deg);
  }
  50% {
    transform: translateY(-6px) rotate(0deg);
  }
  75% {
    transform: translateY(-3px) rotate(-0.5deg);
  }
}

@keyframes pulse-glow {
  0%, 100% {
    opacity: 0.7;
    filter: brightness(1);
  }
  50% {
    opacity: 0.9;
    filter: brightness(1.1);
  }
}

.animate-bounce-in {
  animation: bounce-in 0.8s cubic-bezier(0.68, -0.55, 0.265, 1.55) forwards;
  opacity: 0;
}

.animate-gentle-float {
  animation: gentle-float 4s ease-in-out infinite;
}

.animate-pulse-glow {
  animation: pulse-glow 3s ease-in-out infinite;
}
</style>
