<template>
  <div class="min-h-screen relative overflow-hidden" 
       :class="isDark ? 'bg-gray-900' : 'bg-gray-50'">
    
    <!-- Floating Orb Background -->
    <div class="fixed inset-0 pointer-events-none z-0">
      <div class="floating-orb floating-orb-1 bg-gradient-to-r from-blue-400/30 to-purple-500/30"></div>
      <div class="floating-orb floating-orb-2 bg-gradient-to-r from-green-400/30 to-blue-500/30"></div>
      <div class="floating-orb floating-orb-3 bg-gradient-to-r from-purple-400/30 to-pink-500/30"></div>
    </div>

    <!-- Main Content -->
    <main class="relative z-10 min-h-screen flex flex-col justify-center px-4 py-12">
      
      <!-- Loading State -->
      <div v-if="loading" class="fade-in-up text-center space-y-8" style="animation-delay: 0.1s;">
        <div class="relative mx-auto w-24 h-24">
          <div class="absolute inset-0 rounded-full">
            <div class="w-full h-full rounded-full border-4 border-transparent border-t-blue-500 animate-spin"></div>
            <div class="absolute inset-2 w-20 h-20 rounded-full border-4 border-transparent border-t-purple-500 animate-spin-reverse"></div>
            <div class="absolute inset-4 w-16 h-16 rounded-full border-4 border-transparent border-t-green-500 animate-spin-slow"></div>
          </div>
        </div>
        <div class="space-y-3">
          <h1 class="text-3xl sm:text-4xl font-bold bg-gradient-to-r from-blue-600 via-purple-600 to-green-600 bg-clip-text text-transparent animate-gradient-x">
            Preparing your files
          </h1>
          <p class="text-lg animate-fade-in-word" style="animation-delay: 0.3s;" :class="isDark ? 'text-gray-400' : 'text-gray-600'">
            Setting up your download experience
          </p>
        </div>
      </div>

      <!-- Content State -->
      <div v-else-if="files.length > 0" class="max-w-4xl mx-auto w-full space-y-8">
        
        <!-- Header with Status -->
        <div class="text-center space-y-6 fade-in-up" style="animation-delay: 0.1s;">
          <div class="inline-flex items-center gap-2 px-4 py-2 rounded-full backdrop-blur-md border"
               :class="[
                 isDark ? 'bg-white/5 border-white/10' : 'bg-white/80 border-gray-200/50',
                 'pulse-glow'
               ]">
            <div class="w-2 h-2 rounded-full bg-green-500 animate-pulse"></div>
            <span class="text-sm font-medium" :class="isDark ? 'text-gray-300' : 'text-gray-700'">
              {{ files.length }} file{{ files.length > 1 ? 's' : '' }} ready
            </span>
          </div>
          
          <h1 class="text-4xl sm:text-6xl font-bold bg-gradient-to-r from-blue-600 via-purple-600 to-green-600 bg-clip-text text-transparent animate-gradient-x">
            Download Files
          </h1>
          
          <p class="text-lg animate-fade-in-word" style="animation-delay: 0.3s;" :class="isDark ? 'text-gray-400' : 'text-gray-600'">
            Secure and instant file access
          </p>
        </div>

        <!-- Uploader Card -->
        <div v-if="uploader" class="fade-in-up" style="animation-delay: 0.2s;">
          <div class="backdrop-blur-md border rounded-3xl p-6 sm:p-8"
               :class="isDark ? 'bg-white/5 border-white/10' : 'bg-white/80 border-gray-200/50'">
            
            <div class="flex items-center gap-6">
              <!-- Avatar -->
              <div class="relative group">
                <div class="w-20 h-20 rounded-2xl overflow-hidden backdrop-blur-sm border-2"
                     :class="isDark ? 'border-white/20' : 'border-gray-200'">
                  <img v-if="uploader.avatar" 
                       :src="`http://localhost:8080${uploader.avatar}`" 
                       :alt="uploader.username"
                       class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-110" 
                       @error="handleAvatarError" />
                  <div v-else class="w-full h-full bg-gradient-to-br from-blue-500 via-purple-500 to-green-500 flex items-center justify-center">
                    <span class="text-white text-2xl font-bold">{{ uploader.username.charAt(0).toUpperCase() }}</span>
                  </div>
                </div>
                <!-- Online indicator -->
                <div class="absolute -bottom-1 -right-1 w-6 h-6 bg-green-500 rounded-full border-3"
                     :class="isDark ? 'border-gray-900' : 'border-white'">
                  <div class="w-full h-full bg-green-400 rounded-full animate-pulse"></div>
                </div>
              </div>
              
              <!-- User Info -->
              <div class="flex-1 min-w-0">
                <h3 class="text-2xl font-bold mb-1" :class="isDark ? 'text-white' : 'text-gray-900'">
                  {{ uploader.username }}
                </h3>
                <p class="text-lg opacity-80 mb-3" :class="isDark ? 'text-gray-400' : 'text-gray-600'">
                  {{ uploader.email }}
                </p>
                
                <!-- Stats -->
                <div class="flex flex-wrap gap-3">
                  <div class="px-3 py-1.5 rounded-xl backdrop-blur-sm text-sm font-medium border"
                       :class="isDark ? 'bg-green-500/20 border-green-500/30 text-green-400' : 'bg-green-100 border-green-200 text-green-700'">
                    <span>{{ formatTotalSize() }}</span>
                  </div>
                  
                  <div class="px-3 py-1.5 rounded-xl backdrop-blur-sm text-sm font-medium border"
                       :class="isDark ? 'bg-purple-500/20 border-purple-500/30 text-purple-400' : 'bg-purple-100 border-purple-200 text-purple-700'">
                    <span>{{ files.length }} file{{ files.length > 1 ? 's' : '' }}</span>
                  </div>
                  
                  <div v-if="uploader.expirationDate" 
                       class="px-3 py-1.5 rounded-xl backdrop-blur-sm text-sm font-medium border"
                       :class="formatExpirationDate(uploader.expirationDate).isExpired 
                         ? (isDark ? 'bg-red-500/20 border-red-500/30 text-red-400' : 'bg-red-100 border-red-200 text-red-700')
                         : (isDark ? 'bg-blue-500/20 border-blue-500/30 text-blue-400' : 'bg-blue-100 border-blue-200 text-blue-700')">
                    <span>{{ formatExpirationDate(uploader.expirationDate).text }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Download All Button -->
        <div class="text-center fade-in-up" style="animation-delay: 0.3s;">
          <button 
            @click="downloadAll"
            :disabled="downloadingAll"
            class="group relative px-8 py-4 bg-gradient-to-r from-blue-600 to-purple-600 text-white font-semibold text-lg rounded-2xl transition-all duration-300 hover:scale-105 hover:shadow-2xl hover:shadow-blue-500/25 disabled:opacity-50 disabled:hover:scale-100"
          >
            <span v-if="downloadingAll" class="flex items-center gap-3">
              <div class="w-5 h-5 border-2 border-white/30 border-t-white rounded-full animate-spin"></div>
              Downloading...
            </span>
            <span v-else class="flex items-center gap-3">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"></path>
              </svg>
              Download All Files
            </span>
          </button>
        </div>

        <!-- Files List -->
        <div class="space-y-4 fade-in-up" style="animation-delay: 0.4s;">
          <h2 class="text-2xl font-bold text-center mb-6" :class="isDark ? 'text-white' : 'text-gray-900'">
            Files ({{ files.length }})
          </h2>

          <transition-group name="file-slide" tag="div" class="space-y-3">
            <div v-for="(file, index) in files" :key="`file-${index}-${file.name}`" 
                 class="group backdrop-blur-md border rounded-2xl p-4 sm:p-6 transition-all duration-300 hover:scale-[1.02] hover:shadow-lg"
                 :class="isDark ? 'bg-white/5 border-white/10 hover:bg-white/10' : 'bg-white/80 border-gray-200/50 hover:bg-white'">
            
            <div class="flex flex-col sm:flex-row sm:items-center gap-4">
              <div class="flex items-center gap-4 flex-1 min-w-0">
                <!-- File Icon -->
                <div class="flex items-center justify-center flex-shrink-0">
                  <img 
                    :src="getFileIconPath(file.name)" 
                    :alt="getFileIconAltText(file.name)"
                    class="w-10 h-10 object-contain transition-transform duration-300 group-hover:scale-110"
                  />
                </div>
                
                <!-- File Info -->
                <div class="flex-1 min-w-0">
                  <h3 class="text-lg font-semibold truncate" :class="isDark ? 'text-white' : 'text-gray-900'" :title="file.name">
                    {{ file.name }}
                  </h3>
                  <p class="text-sm mt-1" :class="isDark ? 'text-gray-400' : 'text-gray-600'">
                    {{ formatFileSize(file.size) }}
                  </p>
                </div>
              </div>

              <!-- Actions -->
              <div class="flex items-center gap-3 flex-shrink-0 w-full sm:w-auto">
                <!-- Preview Button -->
                <button v-if="canPreview(file)" 
                        @click="togglePreview(index)"
                        class="p-3 rounded-xl transition-all duration-200 hover:scale-110 flex-1 sm:flex-none"
                        :class="previewingFiles.has(index) 
                          ? (isDark ? 'bg-red-500/20 text-red-400 border border-red-500/30' : 'bg-red-100 text-red-600 border border-red-200')
                          : (isDark ? 'bg-green-500/20 text-green-400 border border-green-500/30' : 'bg-green-100 text-green-600 border border-green-200')">
                  <svg v-if="!previewingFiles.has(index)" class="w-5 h-5 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                  </svg>
                  <svg v-else class="w-5 h-5 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.878 9.878L3 3m6.878 6.878L21 21"></path>
                  </svg>
                </button>

                <!-- Download Button -->
                <button @click="downloadFile(file, index)"
                        :disabled="downloadingStates[index]"
                        class="px-6 py-3 bg-gradient-to-r from-blue-600 to-purple-600 text-white font-semibold rounded-xl transition-all duration-200 hover:scale-110 disabled:opacity-50 flex-1 sm:flex-none">
                  <span v-if="downloadingStates[index]" class="flex items-center justify-center gap-2">
                    <div class="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin"></div>
                    <span class="hidden sm:inline">Downloading</span>
                    <span class="sm:hidden">...</span>
                  </span>
                  <span v-else class="flex items-center gap-2">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"></path>
                    </svg>
                    Download
                  </span>
                </button>
              </div>
            </div>

            <!-- Preview Section -->
            <transition name="preview-slide">
              <div v-if="previewingFiles.has(index)" class="mt-4 pt-4 border-t"
                   :class="isDark ? 'border-white/10' : 'border-gray-200'">
                
                <!-- Image preview -->
                <div v-if="['jpg', 'jpeg', 'png', 'gif', 'webp'].includes(getFileExtension(file.name))">
                  <img 
                    :src="`http://localhost:8080/file/${$route.params.id}/${file.name}?preview=true`"
                    :alt="file.name"
                    class="w-full max-w-md mx-auto rounded-xl shadow-lg"
                    @error="handlePreviewError">
                </div>

                <!-- Video preview -->
                <div v-else-if="getFileExtension(file.name) === 'mp4'">
                  <video 
                    :src="`http://localhost:8080/file/${$route.params.id}/${file.name}?preview=true`"
                    controls 
                    class="w-full max-w-md mx-auto rounded-xl shadow-lg">
                  </video>
                </div>

                <!-- Audio preview -->
                <div v-else-if="['mp3', 'wav', 'flac'].includes(getFileExtension(file.name))">
                  <audio 
                    :src="`http://localhost:8080/file/${$route.params.id}/${file.name}?preview=true`"
                    controls 
                    class="w-full max-w-md mx-auto">
                  </audio>
                </div>

                <!-- PDF preview -->
                <div v-else-if="getFileExtension(file.name) === 'pdf'">
                  <iframe 
                    :src="`http://localhost:8080/file/${$route.params.id}/${file.name}?preview=true`"
                    class="w-full max-w-md mx-auto h-64 rounded-xl border"
                    :class="isDark ? 'border-white/10' : 'border-gray-200'">
                  </iframe>
                </div>

                <!-- Other file types -->
                <div v-else class="text-center p-8">
                  <div class="w-16 h-16 bg-gradient-to-r from-gray-500 to-gray-600 rounded-xl flex items-center justify-center mx-auto mb-4">
                    <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                    </svg>
                  </div>
                  <p class="text-lg" :class="isDark ? 'text-gray-400' : 'text-gray-600'">
                    Preview not available for this file type
                  </p>
                </div>
              </div>
            </transition>
          </div>
          </transition-group>
        </div>
      </div>

      <!-- Error State -->
      <div v-else class="text-center space-y-8 fade-in-up" style="animation-delay: 0.1s;">
        <div class="relative mx-auto w-24 h-24">
          <div class="w-full h-full rounded-full bg-gradient-to-r from-red-500 to-orange-500 flex items-center justify-center">
            <svg class="w-12 h-12 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
            </svg>
          </div>
        </div>
        <div class="space-y-4">
          <h1 class="text-4xl sm:text-5xl font-bold" :class="isDark ? 'text-white' : 'text-gray-900'">
            Share not found
          </h1>
          <p class="text-xl" :class="isDark ? 'text-gray-400' : 'text-gray-600'">
            This share may have expired or been removed
          </p>
          <div class="pt-6">
            <button 
              @click="$router.push('/')"
              class="px-8 py-4 bg-gradient-to-r from-blue-600 to-purple-600 text-white font-semibold rounded-2xl transition-all duration-300 hover:scale-105"
            >
              Go Back Home
            </button>
          </div>
        </div>
      </div>
    </main>

    <!-- Success Message -->
    <div v-if="message" 
         class="fixed bottom-6 right-6 p-4 rounded-2xl shadow-2xl z-50 backdrop-blur-xl border transition-all duration-300"
         :class="message.type === 'success' 
           ? (isDark ? 'bg-green-500/20 border-green-500/30 text-green-400' : 'bg-green-100 border-green-200 text-green-700')
           : (isDark ? 'bg-red-500/20 border-red-500/30 text-red-400' : 'bg-red-100 border-red-200 text-red-700')">
      {{ message.text }}
    </div>

    <!-- GitHub Icon -->
    <a 
      href="https://github.com/R8bert/PInGO-Share" 
      target="_blank" 
      rel="noopener noreferrer" 
      class="fixed bottom-4 left-4 z-40 p-2 rounded-lg backdrop-blur-sm transition-all duration-300 hover:scale-110"
      :class="isDark ? 'text-gray-400 hover:text-white bg-white/5' : 'text-gray-600 hover:text-gray-900 bg-white/50'"
      title="View GitHub Repository"
    >
      <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
        <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
      </svg>
    </a>
  </div>
</template> 
                             :alt="uploader.username"
                             class="w-full h-full object-cover transition-transform duration-300 group-hover:scale-110" 
                             @error="handleAvatarError" />
                        <div v-else class="w-full h-full bg-gradient-to-br from-green-500 via-blue-500 to-purple-500 flex items-center justify-center">
                          <span class="text-white text-3xl font-bold">{{ uploader.username.charAt(0).toUpperCase() }}</span>
                        </div>
                      </div>
                      <!-- Online indicator -->
                      <div class="absolute -bottom-1 -right-1 w-7 h-7 bg-green-500 rounded-full border-4 shadow-lg"
                           :style="{ borderColor: isDark ? '#000000' : '#ffffff' }">
                        <div class="w-full h-full bg-green-400 rounded-full animate-pulse"></div>
                      </div>
                    </div>
                    
                    <!-- User Info -->
                    <div class="flex-1 min-w-0">
                      <h3 class="text-3xl font-bold mb-2 truncate"
                          :style="{ color: isDark ? '#ffffff' : '#000000' }">
                        <!-- {{ uploader.username }} -->
                          <DecryptedText
                            :text="uploader.username"
                            :speed="50"
                            :max-iterations="10"
                            :sequential="false"
                            reveal-direction="start"
                            :use-original-chars-only="false"
                            animate-on="view"
                            class-name="text-green-500"
                            encrypted-class-name="text-red-500"
                        />
                      </h3>
                      <p class="text-lg mb-4 truncate opacity-80"
                         :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">
                        {{ uploader.email }}
                      </p>
                      
                      <!-- Enhanced Stats -->
                      <div class="flex flex-wrap gap-3">
                        <div class="px-4 py-2 rounded-xl backdrop-blur-sm border"
                             :style="{
                               backgroundColor: isDark ? 'rgba(34, 197, 94, 0.2)' : 'rgba(34, 197, 94, 0.1)',
                               borderColor: 'rgba(34, 197, 94, 0.3)',
                               color: '#22c55e'
                             }">
                          <div class="flex items-center space-x-2">
                            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
                            </svg>
                            <span class="text-sm font-medium">{{ formatTotalSize() }}</span>
                          </div>
                        </div>
                        
                        <div class="px-4 py-2 rounded-xl backdrop-blur-sm border"
                             :style="{
                               backgroundColor: isDark ? 'rgba(147, 51, 234, 0.2)' : 'rgba(147, 51, 234, 0.1)',
                               borderColor: 'rgba(147, 51, 234, 0.3)',
                               color: '#9333ea'
                             }">
                          <div class="flex items-center space-x-2">
                            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
                            </svg>
                            <span class="text-sm font-medium">{{ files.length }} file{{ files.length > 1 ? 's' : '' }}</span>
                          </div>
                        </div>
                        
                        <div v-if="uploader.expirationDate" 
                             class="px-4 py-2 rounded-xl backdrop-blur-sm border"
                             :style="{
                               backgroundColor: formatExpirationDate(uploader.expirationDate).isExpired 
                                 ? (isDark ? 'rgba(239, 68, 68, 0.2)' : 'rgba(239, 68, 68, 0.1)')
                                 : (isDark ? 'rgba(59, 130, 246, 0.2)' : 'rgba(59, 130, 246, 0.1)'),
                               borderColor: formatExpirationDate(uploader.expirationDate).isExpired 
                                 ? 'rgba(239, 68, 68, 0.3)' 
                                 : 'rgba(59, 130, 246, 0.3)',
                               color: formatExpirationDate(uploader.expirationDate).isExpired ? '#ef4444' : '#3b82f6'
                             }">
                          <div class="flex items-center space-x-2">
                            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                            </svg>
                            <span class="text-sm font-medium">{{ formatExpirationDate(uploader.expirationDate).text }}</span>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- Download All Button -->
            <div class="text-center">
              <button 
                @click="downloadAll"
                :disabled="downloadingAll"
                :title="files.length > 1 ? `Download all ${files.length} files as ZIP` : 'Download file'"
                class="group relative px-12 py-6 bg-gradient-to-r from-green-600 to-blue-600 text-white font-bold text-xl rounded-3xl transition-all duration-300 hover:scale-105 hover:shadow-2xl hover:shadow-green-500/25 disabled:opacity-50 disabled:hover:scale-100"
              >
                <span v-if="downloadingAll" class="relative z-10 flex items-center">
                  <svg class="animate-spin -ml-1 mr-4 h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  Downloading...
                </span>
                <span v-else class="relative z-10 flex items-center">
                  <svg class="w-6 h-6 mr-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"></path>
                  </svg>
                  Download All Files
                </span>
                <div class="absolute inset-0 bg-gradient-to-r from-green-700 to-blue-700 rounded-3xl opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
              </button>
            </div>

            <!-- Files List -->
            <div class="max-w-4xl mx-auto">
              <div class="relative">
                <div class="absolute inset-0 bg-gradient-to-r from-blue-600/10 to-purple-600/10 rounded-3xl blur-3xl"></div>
                <div class="relative backdrop-blur-xl border rounded-3xl p-8"
                     :style="{
                       backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
                       borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)'
                     }">
                  
                  <h2 class="text-2xl font-bold mb-6"
                      :style="{ color: isDark ? '#ffffff' : '#000000' }">
                    Files ({{ files.length }})
                  </h2>

                  <div class="grid gap-4">
                    <transition-group name="file-list" tag="div" class="grid gap-3 sm:gap-4">
                      <div v-for="(file, index) in files" :key="`file-${index}-${file.name}`" 
                           class="group rounded-xl sm:rounded-2xl border p-4 sm:p-6 transition-all duration-300 hover:scale-[1.02] overflow-hidden"
                           :style="{
                             backgroundColor: isDark ? 'rgba(255,255,255,0.03)' : 'rgba(255,255,255,0.5)',
                             borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)'
                           }">
                      
                      <div class="flex flex-col sm:flex-row sm:items-center gap-4">
                        <div class="flex items-center space-x-3 sm:space-x-4 flex-1 min-w-0">
                          <!-- File Icon -->
                          <div class="flex items-center justify-center flex-shrink-0">
                            <img 
                              :src="getFileIconPath(file.name)" 
                              :alt="getFileIconAltText(file.name)"
                              class="w-8 h-8 sm:w-10 sm:h-10 object-contain"
                            />
                          </div>
                          
                          <!-- File Info -->
                          <div class="flex-1 min-w-0">
                            <h3 class="text-base sm:text-lg font-semibold truncate break-all"
                                :style="{ color: isDark ? '#ffffff' : '#000000' }"
                                :title="file.name">
                              {{ file.name }}
                            </h3>
                            <div class="flex items-center mt-1">
                              <span class="text-sm"
                                    :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">
                                {{ formatFileSize(file.size) }}
                              </span>
                            </div>
                          </div>
                        </div>

                        <!-- Actions -->
                        <div class="flex items-center gap-2 sm:gap-3 flex-shrink-0 w-full sm:w-auto">
                          <!-- Preview Button -->
                          <button v-if="canPreview(file)" 
                                  @click="togglePreview(index)"
                                  :title="previewingFiles.has(index) ? 'Hide preview' : 'Show preview'"
                                  class="p-2 sm:p-3 rounded-lg sm:rounded-xl transition-all duration-200 hover:scale-110 flex-1 sm:flex-none"
                                  :style="{
                                    backgroundColor: previewingFiles.has(index) 
                                      ? (isDark ? 'rgba(239, 68, 68, 0.2)' : 'rgba(239, 68, 68, 0.1)')
                                      : (isDark ? 'rgba(34, 197, 94, 0.2)' : 'rgba(34, 197, 94, 0.1)'),
                                    color: previewingFiles.has(index) ? '#ef4444' : '#22c55e'
                                  }">
                            <svg v-if="!previewingFiles.has(index)" class="w-4 h-4 sm:w-5 sm:h-5 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                            </svg>
                            <svg v-else class="w-4 h-4 sm:w-5 sm:h-5 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.878 9.878L3 3m6.878 6.878L21 21"></path>
                            </svg>
                          </button>

                          <!-- Download Button -->
                          <button @click="downloadFile(file, index)"
                                  :disabled="downloadingStates[index]"
                                  :title="`Download ${file.name}`"
                                  class="px-4 sm:px-6 py-2 sm:py-3 bg-gradient-to-r from-blue-600 to-purple-600 text-white font-semibold rounded-lg sm:rounded-xl transition-all duration-200 hover:scale-110 disabled:opacity-50 flex-1 sm:flex-none text-sm sm:text-base">
                            <span v-if="downloadingStates[index]" class="flex items-center justify-center">
                              <svg class="animate-spin -ml-1 mr-2 h-3 w-3 sm:h-4 sm:w-4" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                              </svg>
                              <span class="hidden sm:inline">Downloading</span>
                              <span class="sm:hidden">...</span>
                            </span>
                            <span v-else class="flex items-center">
                              <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"></path>
                              </svg>
                              Download
                            </span>
                          </button>
                        </div>
                      </div>

                      <!-- Preview Section -->
                      <div v-if="previewingFiles.has(index)" class="mt-6 pt-6 border-t"
                           :style="{ borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)' }">
                        
                        <!-- Image preview -->
                        <div v-if="['jpg', 'jpeg', 'png', 'gif', 'webp'].includes(getFileExtension(file.name))">
                          <img 
                            :src="`http://localhost:8080/file/${$route.params.id}/${file.name}?preview=true`"
                            :alt="file.name"
                            class="w-full max-w-md mx-auto rounded-xl shadow-lg"
                            @error="handlePreviewError">
                        </div>

                        <!-- Video preview -->
                        <div v-else-if="getFileExtension(file.name) === 'mp4'">
                          <video 
                            :src="`http://localhost:8080/file/${$route.params.id}/${file.name}?preview=true`"
                            controls 
                            class="w-full max-w-md mx-auto rounded-xl shadow-lg">
                          </video>
                        </div>

                        <!-- Audio preview -->
                        <div v-else-if="['mp3', 'wav', 'flac'].includes(getFileExtension(file.name))">
                          <audio 
                            :src="`http://localhost:8080/file/${$route.params.id}/${file.name}?preview=true`"
                            controls 
                            class="w-full max-w-md mx-auto">
                          </audio>
                        </div>

                        <!-- PDF preview -->
                        <div v-else-if="getFileExtension(file.name) === 'pdf'">
                          <iframe 
                            :src="`http://localhost:8080/file/${$route.params.id}/${file.name}?preview=true`"
                            class="w-full max-w-md mx-auto h-64 rounded-xl border"
                            :style="{ borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)' }">
                          </iframe>
                        </div>

                        <!-- Other file types -->
                        <div v-else class="text-center p-8">
                          <div class="w-16 h-16 bg-gradient-to-r from-gray-500 to-gray-600 rounded-xl flex items-center justify-center mx-auto mb-4">
                            <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                            </svg>
                          </div>
                          <p class="text-lg" :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">
                            Preview not available for this file type
                          </p>
                        </div>
                      </div>
                    </div>
                    </transition-group>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Error State -->
          <div v-else class="text-center space-y-8 animate-fade-in">
            <div class="relative mx-auto w-32 h-32">
              <div class="absolute inset-0 bg-gradient-to-r from-red-500 to-orange-500 rounded-3xl blur-2xl opacity-30"></div>
              <div class="relative bg-gradient-to-r from-red-500 to-orange-500 rounded-3xl p-8 flex items-center justify-center">
                <svg class="w-16 h-16 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
                </svg>
              </div>
            </div>
            <div class="space-y-4">
              <h1 class="text-4xl sm:text-5xl font-bold"
                  :style="{ color: isDark ? '#ffffff' : '#000000' }">
                Share not found
              </h1>
              <p class="text-xl"
                 :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">
                This share may have expired or been removed
              </p>
              <div class="pt-6">
                <button 
                  @click="$router.push('/')"
                  class="px-8 py-4 bg-gradient-to-r from-blue-600 to-purple-600 text-white font-semibold rounded-2xl transition-all duration-300 hover:scale-105"
                >
                  Go Back Home
                </button>
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- Footer -->
      <footer class="relative py-8 px-4 sm:px-6 lg:px-8 border-t"
              :style="{ borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)' }">
        <div class="max-w-7xl mx-auto text-center">
          <p :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">
            © 2025 PInGO Share. Secure file sharing made simple.
          </p>
        </div>
      </footer>
    </main>

    <!-- Success Message -->
    <div v-if="message" 
         class="fixed bottom-6 right-6 p-4 rounded-2xl shadow-2xl z-50 backdrop-blur-xl border"
         :style="{
           backgroundColor: message.type === 'success' 
             ? (isDark ? 'rgba(34, 197, 94, 0.2)' : 'rgba(34, 197, 94, 0.1)')
             : (isDark ? 'rgba(239, 68, 68, 0.2)' : 'rgba(239, 68, 68, 0.1)'),
           borderColor: message.type === 'success' ? '#22c55e' : '#ef4444',
           color: message.type === 'success' ? '#22c55e' : '#ef4444'
         }">
      {{ message.text }}
    </div>

    <!-- GitHub Icon - Bottom Left -->
    <a 
      href="https://github.com/R8bert/PInGO-Share" 
      target="_blank" 
      rel="noopener noreferrer" 
      class="fixed bottom-4 left-4 z-40"
      :style="{ color: isDark ? '#9ca3af' : '#6b7280' }"
      title="View GitHub Repository"
    >
      <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
        <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
      </svg>
    </a>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useTheme } from '../../composables/useTheme'
import { useIcons } from '../../composables/useIcons'
import axios from 'axios'

// Interfaces
interface FileInfo {
  name: string
  size: number
}

interface UploaderInfo {
  username: string
  email: string
  avatar?: string
  expirationDate?: string
}

interface Message {
  text: string
  type: 'success' | 'error'
}

const { isDark } = useTheme()
const { 
  getFileIcon, 
  getFileIconAlt, 
  getFileExtension 
} = useIcons()
const route = useRoute()

// Helper methods for file icons
const getFileIconPath = (filename: string) => {
  return getFileIcon(filename)
}

const getFileIconAltText = (filename: string) => {
  return getFileIconAlt(filename)
}

// State
const loading = ref(true)
const files = ref<FileInfo[]>([])
const uploader = ref<UploaderInfo | null>(null)
const downloadingAll = ref(false)
const downloadingStates = ref<Record<number, boolean>>({})
const message = ref<Message | null>(null)
const previewingFiles = ref<Set<number>>(new Set())

// Methods
const fetchFiles = async () => {
  try {
    const response = await axios.get(`/files/${route.params.id}`)
    files.value = response.data.files || []
    uploader.value = response.data.uploader || null
    loading.value = false
  } catch (error) {
    console.error('Error fetching files:', error)
    message.value = { text: 'Share not found', type: 'error' }
    loading.value = false
  }
}

const downloadFile = async (file: FileInfo, index: number) => {
  downloadingStates.value[index] = true
  
  try {
    const response = await axios.get(
      `http://localhost:8080/file/${route.params.id}/${file.name}`,
      { responseType: 'blob' }
    )
    
    const url = window.URL.createObjectURL(new Blob([response.data]))
    const link = document.createElement('a')
    link.href = url
    link.setAttribute('download', file.name)
    document.body.appendChild(link)
    link.click()
    link.remove()
    window.URL.revokeObjectURL(url)
    
    showMessage('File downloaded successfully!', 'success')
  } catch (error) {
    console.error('Download error:', error)
    showMessage('Download failed. Please try again.', 'error')
  } finally {
    downloadingStates.value[index] = false
  }
}

const downloadAll = async () => {
  downloadingAll.value = true
  
  try {
    const response = await axios.get(
      `http://localhost:8080/download/${route.params.id}`,
      { responseType: 'blob' }
    )
    
    const url = window.URL.createObjectURL(new Blob([response.data]))
    const link = document.createElement('a')
    link.href = url
    link.setAttribute('download', `share-${route.params.id}.zip`)
    document.body.appendChild(link)
    link.click()
    link.remove()
    window.URL.revokeObjectURL(url)
    
    showMessage('All files downloaded successfully!', 'success')
  } catch (error) {
    console.error('Download all error:', error)
    showMessage('Download failed. Please try again.', 'error')
  } finally {
    downloadingAll.value = false
  }
}

const togglePreview = (index: number) => {
  if (previewingFiles.value.has(index)) {
    previewingFiles.value.delete(index)
  } else {
    previewingFiles.value.add(index)
  }
}

const canPreview = (file: FileInfo): boolean => {
  const ext = getFileExtension(file.name)
  return ['jpg', 'jpeg', 'png', 'gif', 'webp', 'mp4', 'mp3', 'wav', 'flac', 'pdf'].includes(ext)
}

const formatFileSize = (bytes: number): string => {
  if (bytes === 0) return '0 Bytes'
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

const formatTotalSize = (): string => {
  const total = files.value.reduce((sum, file) => sum + file.size, 0)
  return formatFileSize(total)
}

const formatExpirationDate = (dateString: string) => {
  const date = new Date(dateString)
  const now = new Date()
  const isExpired = date < now
  
  if (isExpired) {
    return { text: 'Expired', isExpired: true }
  }
  
  const diffTime = date.getTime() - now.getTime()
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
  
  if (diffDays === 1) return { text: 'Expires today', isExpired: false }
  if (diffDays <= 7) return { text: `Expires in ${diffDays} days`, isExpired: false }
  
  return { text: `Expires ${date.toLocaleDateString()}`, isExpired: false }
}

const handleAvatarError = () => {
  if (uploader.value) {
    uploader.value.avatar = undefined
  }
}

<style scoped>
/* Floating orbs animation */
.floating-orb {
  position: absolute;
  border-radius: 9999px;
  filter: blur(3rem);
  animation: float 20s ease-in-out infinite;
}

.floating-orb-1 {
  width: 24rem;
  height: 24rem;
  top: -12rem;
  left: -12rem;
  animation-delay: 0s;
}

.floating-orb-2 {
  width: 20rem;
  height: 20rem;
  top: 50%;
  right: -10rem;
  animation-delay: -7s;
}

.floating-orb-3 {
  width: 18rem;
  height: 18rem;
  bottom: -9rem;
  left: 33.333333%;
  animation-delay: -14s;
}

@keyframes float {
  0%, 100% { 
    transform: translateY(0px) translateX(0px) scale(1);
  }
  33% { 
    transform: translateY(-30px) translateX(20px) scale(1.1);
  }
  66% { 
    transform: translateY(20px) translateX(-15px) scale(0.9);
  }
}

/* Fade in animations */
.fade-in-up {
  animation: fadeInUp 0.8s ease-out forwards;
  opacity: 0;
  transform: translateY(30px);
}

@keyframes fadeInUp {
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-fade-in-word {
  animation: fadeInWord 0.6s ease-out forwards;
  opacity: 0;
  transform: translateY(20px);
}

@keyframes fadeInWord {
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Gradient text animation */
.animate-gradient-x {
  background-size: 200% 200%;
  animation: gradientX 3s ease infinite;
}

@keyframes gradientX {
  0%, 100% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
}

/* Custom animations */
.animate-spin-reverse {
  animation: spin 2s linear infinite reverse;
}

.animate-spin-slow {
  animation: spin 3s linear infinite;
}

.animate-pulse-slow {
  animation: pulse 3s ease-in-out infinite;
}

.pulse-glow {
  animation: pulseGlow 2s ease-in-out infinite;
}

@keyframes pulseGlow {
  0%, 100% {
    opacity: 1;
    filter: drop-shadow(0 0 10px currentColor);
  }
  50% {
    opacity: 0.7;
    filter: drop-shadow(0 0 20px currentColor);
  }
}

/* Transition groups */
.file-slide-enter-active,
.file-slide-leave-active {
  transition: all 0.5s ease;
}

.file-slide-enter-from {
  opacity: 0;
  transform: translateX(-30px);
}

.file-slide-leave-to {
  opacity: 0;
  transform: translateX(30px);
}

.preview-slide-enter-active,
.preview-slide-leave-active {
  transition: all 0.4s ease;
}

.preview-slide-enter-from,
.preview-slide-leave-to {
  opacity: 0;
  max-height: 0;
  transform: translateY(-10px);
}
</style>
</script>

const handlePreviewError = () => {
  console.log('Preview error occurred')
}

const showMessage = (text: string, type: 'success' | 'error') => {
  message.value = { text, type }
  setTimeout(() => {
    message.value = null
  }, 3000)
}

onMounted(() => {
  fetchFiles()
})
</script>

<style scoped>
@keyframes fade-in {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}

.animate-fade-in {
  animation: fade-in 0.8s ease-out;
}

/* File list animations */
.file-list-enter-active,
.file-list-leave-active {
  transition: all 0.3s ease;
}

.file-list-enter-from,
.file-list-leave-to {
  opacity: 0;
  transform: translateX(-30px);
}

.file-list-move {
  transition: transform 0.3s ease;
}



@font-face {
  font-family: 'HaloHandletter-Local';
  src: url('/HaloHandletter.otf') format('opentype');
  font-weight: normal;
  font-style: normal;
  font-display: swap;
}

.animated-text-container {
  font-family: 'Crima-Local', 'HaloHandletter-Local', 'HaloHandletter', 'Comic Sans MS', cursive !important;
  letter-spacing: 0.065em;
  background: linear-gradient(
    90deg,
    #ff6b6b 0%,
    #4ecdc4 12.5%,
    #45b7d1 25%,
    #96ceb4 37.5%,
    #feca57 50%,
    #ff9ff3 62.5%,
    #54a0ff 75%,
    #5f27cd 87.5%,
    #ff6b6b 100%
  );
  background-size: 800% 100%;
  -webkit-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;
  animation: liquid-rainbow 6s cubic-bezier(0.4, 0, 0.2, 1) infinite;
}

@keyframes liquid-rainbow {
  0% {
    background-position: 0% 50%;
  }
  25% {
    background-position: 100% 50%;
  }
  50% {
    background-position: 200% 50%;
  }
  75% {
    background-position: 300% 50%;
  }
  100% {
    background-position: 400% 50%;
  }
}

.animated-text-container * {
  font-family: inherit !important;
}
</style>
