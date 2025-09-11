<template>
    <div class="min-h-screen relative overflow-hidden">
        <!-- Background Image -->
        <div class="absolute inset-0 z-0">
            <div 
                v-if="!isDark"
                class="w-full h-full bg-gradient-to-br from-blue-400 via-purple-500 to-pink-500"
                :style="{
                    backgroundImage: `url('data:image/svg+xml,${encodeURIComponent(lightBackground)}')`,
                    backgroundSize: 'cover',
                    backgroundPosition: 'center',
                    backgroundRepeat: 'no-repeat'
                }"
            ></div>
            <div 
                v-else
                class="w-full h-full bg-gradient-to-br from-gray-900 via-blue-900 to-purple-900"
                :style="{
                    backgroundImage: `url('data:image/svg+xml,${encodeURIComponent(darkBackground)}')`,
                    backgroundSize: 'cover',
                    backgroundPosition: 'center',
                    backgroundRepeat: 'no-repeat'
                }"
            ></div>
            <!-- Overlay -->
            <div class="absolute inset-0 bg-black/20"></div>
        </div>

        <!-- Header -->
        <header class="relative z-50 p-4 sm:p-6">
            <div class="flex justify-between items-center">
                <!-- Logo -->
                <div class="flex items-center space-x-3">
                    <div class="w-10 h-10 bg-white/20 backdrop-blur-md rounded-xl flex items-center justify-center border border-white/30">
                        <span class="text-white font-bold text-lg">P</span>
                    </div>
                    <span class="text-white text-xl font-semibold">PInGO Share</span>
                </div>

                <!-- Header Actions -->
                <div class="flex items-center space-x-4">
                    <button
                        @click="isDark = !isDark"
                        class="p-2 text-white/80 hover:text-white transition-colors rounded-lg hover:bg-white/10"
                    >
                        <svg v-if="isDark" class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                            <path d="M12 2.25a.75.75 0 01.75.75v2.25a.75.75 0 01-1.5 0V3a.75.75 0 01.75-.75zM7.5 12a4.5 4.5 0 119 0 4.5 4.5 0 01-9 0zM18.894 6.166a.75.75 0 00-1.06-1.06l-1.591 1.59a.75.75 0 101.06 1.061l1.591-1.59zM21.75 12a.75.75 0 01-.75.75h-2.25a.75.75 0 010-1.5H21a.75.75 0 01.75.75zM17.834 18.894a.75.75 0 001.06-1.06l-1.59-1.591a.75.75 0 10-1.061 1.06l1.59 1.591zM12 18a.75.75 0 01.75.75V21a.75.75 0 01-1.5 0v-2.25A.75.75 0 0112 18zM7.758 17.303a.75.75 0 00-1.061-1.06l-1.591 1.59a.75.75 0 001.06 1.061l1.591-1.59zM6 12a.75.75 0 01-.75.75H3a.75.75 0 010-1.5h2.25A.75.75 0 016 12zM6.697 7.757a.75.75 0 001.06-1.06l-1.59-1.591a.75.75 0 00-1.061 1.06l1.59 1.591z"/>
                        </svg>
                        <svg v-else class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                            <path fill-rule="evenodd" d="M9.528 1.718a.75.75 0 01.162.819A8.97 8.97 0 009 6a9 9 0 009 9 8.97 8.97 0 003.463-.69.75.75 0 01.981.98 10.503 10.503 0 01-9.694 6.46c-5.799 0-10.5-4.701-10.5-10.5 0-4.368 2.667-8.112 6.46-9.694a.75.75 0 01.818.162z" clip-rule="evenodd"/>
                        </svg>
                    </button>
                    
                    <div v-if="!isAuthenticated">
                        <router-link 
                            to="/auth" 
                            class="text-white/80 hover:text-white text-sm font-medium transition-colors px-4 py-2 rounded-lg hover:bg-white/10"
                        >
                            Sign in
                        </router-link>
                    </div>
                    
                    <div v-else>
                        <router-link 
                            to="/account" 
                            class="text-white/80 hover:text-white text-sm font-medium transition-colors px-4 py-2 rounded-lg hover:bg-white/10"
                        >
                            Account
                        </router-link>
                    </div>
                </div>
            </div>
        </header>

        <!-- Main Content -->
        <main class="relative z-40 flex items-center justify-center min-h-[calc(100vh-6rem)] p-4">
            
            <!-- Upload Success Modal -->
            <div 
                v-if="uploadSuccess && shareableLink" 
                class="w-full max-w-md mx-auto"
            >
                <div class="bg-white/95 dark:bg-gray-900/95 backdrop-blur-xl rounded-3xl shadow-2xl p-8 border border-white/20">
                    <!-- Success Icon -->
                    <div class="text-center mb-6">
                        <div class="w-16 h-16 bg-green-500 rounded-full flex items-center justify-center mx-auto mb-4">
                            <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                            </svg>
                        </div>
                        <h2 class="text-2xl font-bold text-gray-900 dark:text-white mb-2">
                            Files uploaded successfully! 🎉
                        </h2>
                        <p class="text-gray-600 dark:text-gray-300 text-sm">
                            Your files are ready to share
                        </p>
                    </div>

                    <!-- Share Link -->
                    <div class="space-y-4">
                        <div class="flex space-x-2">
                            <input
                                v-model="shareableLink"
                                readonly
                                class="flex-1 px-4 py-3 bg-gray-100 dark:bg-gray-800 border border-gray-200 dark:border-gray-700 rounded-xl text-sm font-mono"
                            />
                            <button
                                @click="copyToClipboard"
                                class="px-4 py-3 bg-blue-500 hover:bg-blue-600 text-white rounded-xl transition-colors whitespace-nowrap text-sm font-semibold"
                            >
                                {{ copied ? '✓' : 'Copy' }}
                            </button>
                        </div>
                        
                        <button
                            @click="resetUpload"
                            class="w-full px-6 py-3 bg-gray-100 dark:bg-gray-800 hover:bg-gray-200 dark:hover:bg-gray-700 text-gray-700 dark:text-gray-300 rounded-xl transition-colors font-medium"
                        >
                            Upload more files
                        </button>
                    </div>
                </div>
            </div>

            <!-- Main Upload Modal -->
            <div v-else class="w-full max-w-md mx-auto">
                <div class="bg-white/95 dark:bg-gray-900/95 backdrop-blur-xl rounded-3xl shadow-2xl p-8 border border-white/20">
                    
                    <!-- Header -->
                    <div class="text-center mb-8">
                        <div class="w-16 h-16 bg-blue-500 rounded-2xl flex items-center justify-center mx-auto mb-4">
                            <CloudArrowUpIcon class="w-8 h-8 text-white" />
                        </div>
                        <h1 class="text-2xl font-bold text-gray-900 dark:text-white mb-2">
                            Share files easily
                        </h1>
                        <p class="text-gray-600 dark:text-gray-300 text-sm">
                            Send files up to 100MB for free. No registration required.
                        </p>
                    </div>

                    <!-- Upload Area -->
                    <div class="space-y-6">
                        <!-- Drop Zone -->
                        <div
                            @drop="handleFileDrop"
                            @dragover="handleDragOver" 
                            @dragleave="handleDragLeave"
                            @click="triggerFileSelect"
                            :class="[
                                'border-2 border-dashed rounded-2xl p-8 text-center cursor-pointer transition-all duration-200',
                                isDragging 
                                    ? 'border-blue-500 bg-blue-50 dark:bg-blue-900/20' 
                                    : 'border-gray-300 dark:border-gray-600 hover:border-blue-400 dark:hover:border-blue-500'
                            ]"
                        >
                            <input
                                ref="fileInput"
                                type="file"
                                multiple
                                @change="handleFileSelect"
                                class="hidden"
                            />
                            
                            <div class="space-y-3">
                                <div class="w-12 h-12 bg-blue-100 dark:bg-blue-900/50 rounded-full flex items-center justify-center mx-auto">
                                    <CloudArrowUpIcon class="w-6 h-6 text-blue-600 dark:text-blue-400" />
                                </div>
                                <div>
                                    <p class="font-semibold text-gray-900 dark:text-white">
                                        Drop files here or click to browse
                                    </p>
                                    <p class="text-xs text-gray-500 dark:text-gray-400 mt-1">
                                        Maximum file size: 100MB
                                    </p>
                                </div>
                            </div>
                        </div>

                        <!-- Selected Files -->
                        <div v-if="selectedFiles.length > 0" class="space-y-4">
                            <!-- Files List -->
                            <div class="bg-gray-50 dark:bg-gray-800/50 rounded-xl p-4">
                                <div class="flex items-center justify-between mb-3">
                                    <span class="text-sm font-medium text-gray-900 dark:text-white">
                                        {{ selectedFiles.length }} file{{ selectedFiles.length === 1 ? '' : 's' }} selected
                                    </span>
                                    <button
                                        @click="clearFiles"
                                        class="text-xs text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200"
                                    >
                                        Clear all
                                    </button>
                                </div>
                                
                                <div class="space-y-2 max-h-32 overflow-y-auto">
                                    <div
                                        v-for="(file, index) in selectedFiles"
                                        :key="index"
                                        class="flex items-center justify-between p-2 bg-white dark:bg-gray-700 rounded-lg"
                                    >
                                        <div class="flex items-center space-x-3 flex-1 min-w-0">
                                            <img :src="getFileIconPath(file.name)" :alt="getFileIconAltText(file.name)" 
                                                 class="w-4 h-4 flex-shrink-0" />
                                            <div class="flex-1 min-w-0">
                                                <p class="text-sm font-medium text-gray-900 dark:text-white truncate">
                                                    {{ file.name }}
                                                </p>
                                                <p class="text-xs text-gray-500 dark:text-gray-400">
                                                    {{ formatFileSize(file.size) }}
                                                </p>
                                            </div>
                                        </div>
                                        <button
                                            @click="removeFile(index)"
                                            class="p-1 text-gray-400 hover:text-red-500 transition-colors"
                                        >
                                            <XMarkIcon class="w-4 h-4" />
                                        </button>
                                    </div>
                                </div>
                            </div>

                            <!-- Expiration Settings -->
                            <div class="space-y-3">
                                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">
                                    Link expires in
                                </label>
                                <select
                                    v-model="selectedValidity"
                                    class="w-full px-4 py-3 bg-gray-50 dark:bg-gray-800 border border-gray-200 dark:border-gray-700 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-sm"
                                >
                                    <option value="1h">1 hour</option>
                                    <option value="1d">1 day</option>
                                    <option value="7d">7 days</option>
                                    <option value="30d">30 days</option>
                                </select>
                            </div>

                            <!-- Upload Button -->
                            <button
                                @click="uploadFiles"
                                :disabled="uploading || selectedFiles.length === 0"
                                class="w-full bg-blue-500 hover:bg-blue-600 disabled:bg-gray-400 disabled:cursor-not-allowed text-white font-semibold py-3 px-6 rounded-xl transition-all duration-200 flex items-center justify-center"
                            >
                                <span v-if="uploading" class="flex items-center">
                                    <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" fill="none" viewBox="0 0 24 24">
                                        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                                        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                                    </svg>
                                    Uploading... {{ uploadProgress }}%
                                </span>
                                <span v-else>
                                    Upload {{ selectedFiles.length }} file{{ selectedFiles.length === 1 ? '' : 's' }}
                                </span>
                            </button>
                        </div>

                        <!-- Error Message -->
                        <div v-if="errorMessage" class="p-4 bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-xl">
                            <div class="flex items-center space-x-2">
                                <svg class="w-5 h-5 text-red-500" fill="currentColor" viewBox="0 0 24 24">
                                    <path fill-rule="evenodd" d="M2.25 12c0-5.385 4.365-9.75 9.75-9.75s9.75 4.365 9.75 9.75-4.365 9.75-9.75 9.75S2.25 17.385 2.25 12zM12 8.25a.75.75 0 01.75.75v3.75a.75.75 0 01-1.5 0V9a.75.75 0 01.75-.75zm0 8.25a.75.75 0 100-1.5.75.75 0 000 1.5z" clip-rule="evenodd"/>
                                </svg>
                                <span class="text-sm text-red-700 dark:text-red-300">{{ errorMessage }}</span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </main>

        <!-- Footer -->
        <footer class="relative z-30 p-6 text-center">
            <p class="text-white/70 text-sm">
                Files are stored securely and automatically deleted after expiration
            </p>
        </footer>
    </div>
</template>

<script setup lang="ts">
import { ref, computed } from "vue";
import { useAuth } from "../../composables/useAuth";
import { useTheme } from "../../composables/useTheme";
import { useIcons } from "../../composables/useIcons";
import { getApiUrl } from "../../utils/apiUtils";
import axios from "axios";
import {
    CloudArrowUpIcon,
    XMarkIcon,
} from "@heroicons/vue/24/outline";

const { isAuthenticated } = useAuth();
const { isDark } = useTheme();
const { getFileIcon, getFileIconAlt } = useIcons();

// Helper methods for file icons
const getFileIconPath = (filename: string) => {
    return getFileIcon(filename);
};

const getFileIconAltText = (filename: string) => {
    return getFileIconAlt(filename);
};

// Beautiful SVG backgrounds
const lightBackground = computed(() => `
<svg viewBox="0 0 1200 800" xmlns="http://www.w3.org/2000/svg">
  <defs>
    <linearGradient id="skyGrad" x1="0%" y1="0%" x2="0%" y2="100%">
      <stop offset="0%" style="stop-color:#87CEEB;stop-opacity:1" />
      <stop offset="100%" style="stop-color:#98D8E8;stop-opacity:1" />
    </linearGradient>
  </defs>
  
  <!-- Sky -->
  <rect width="1200" height="600" fill="url(#skyGrad)"/>
  
  <!-- Mountains -->
  <polygon points="0,400 200,200 400,300 600,150 800,250 1000,180 1200,220 1200,600 0,600" fill="#8FBC8F" opacity="0.8"/>
  <polygon points="100,500 300,300 500,350 700,250 900,320 1100,280 1200,300 1200,600 0,600" fill="#9ACD32" opacity="0.6"/>
  
  <!-- Clouds -->
  <ellipse cx="200" cy="100" rx="60" ry="30" fill="white" opacity="0.9"/>
  <ellipse cx="180" cy="90" rx="40" ry="25" fill="white" opacity="0.9"/>
  <ellipse cx="220" cy="90" rx="45" ry="20" fill="white" opacity="0.9"/>
  
  <ellipse cx="800" cy="120" rx="70" ry="35" fill="white" opacity="0.8"/>
  <ellipse cx="780" cy="110" rx="45" ry="28" fill="white" opacity="0.8"/>
  <ellipse cx="820" cy="110" rx="50" ry="22" fill="white" opacity="0.8"/>
  
  <!-- Ground -->
  <rect y="600" width="1200" height="200" fill="#90EE90"/>
</svg>
`);

const darkBackground = computed(() => `
<svg viewBox="0 0 1200 800" xmlns="http://www.w3.org/2000/svg">
  <defs>
    <linearGradient id="nightSky" x1="0%" y1="0%" x2="0%" y2="100%">
      <stop offset="0%" style="stop-color:#1a1a2e;stop-opacity:1" />
      <stop offset="50%" style="stop-color:#16213e;stop-opacity:1" />
      <stop offset="100%" style="stop-color:#0f3460;stop-opacity:1" />
    </linearGradient>
  </defs>
  
  <!-- Night Sky -->
  <rect width="1200" height="600" fill="url(#nightSky)"/>
  
  <!-- Stars -->
  <circle cx="100" cy="50" r="2" fill="white" opacity="0.8"/>
  <circle cx="300" cy="80" r="1.5" fill="white" opacity="0.9"/>
  <circle cx="500" cy="40" r="1" fill="white" opacity="0.7"/>
  <circle cx="700" cy="90" r="2" fill="white" opacity="0.8"/>
  <circle cx="900" cy="60" r="1.5" fill="white" opacity="0.9"/>
  <circle cx="1100" cy="70" r="1" fill="white" opacity="0.7"/>
  
  <!-- Mountains -->
  <polygon points="0,400 200,200 400,300 600,150 800,250 1000,180 1200,220 1200,600 0,600" fill="#2c3e50" opacity="0.9"/>
  <polygon points="100,500 300,300 500,350 700,250 900,320 1100,280 1200,300 1200,600 0,600" fill="#34495e" opacity="0.7"/>
  
  <!-- Ground -->
  <rect y="600" width="1200" height="200" fill="#27ae60" opacity="0.3"/>
</svg>
`);

// State
const selectedFiles = ref<File[]>([]);
const isDragging = ref(false);
const uploading = ref(false);
const uploadProgress = ref(0);
const uploadSuccess = ref(false);
const shareableLink = ref("");
const errorMessage = ref("");
const copied = ref(false);
const selectedValidity = ref("7d");
const fileInput = ref<HTMLInputElement | null>(null);

// Methods
const triggerFileSelect = () => {
    fileInput.value?.click();
};

const handleFileSelect = (event: Event) => {
    const target = event.target as HTMLInputElement;
    if (target.files) {
        selectedFiles.value = Array.from(target.files);
        uploadSuccess.value = false;
        errorMessage.value = "";
    }
};

const handleFileDrop = (event: DragEvent) => {
    event.preventDefault();
    isDragging.value = false;
    
    if (event.dataTransfer?.files) {
        selectedFiles.value = Array.from(event.dataTransfer.files);
        uploadSuccess.value = false;
        errorMessage.value = "";
    }
};

const handleDragOver = (event: DragEvent) => {
    event.preventDefault();
    isDragging.value = true;
};

const handleDragLeave = (event: DragEvent) => {
    event.preventDefault();
    isDragging.value = false;
};

const removeFile = (index: number) => {
    selectedFiles.value.splice(index, 1);
    if (selectedFiles.value.length === 0) {
        uploadSuccess.value = false;
        errorMessage.value = "";
    }
};

const clearFiles = () => {
    selectedFiles.value = [];
    uploadSuccess.value = false;
    errorMessage.value = "";
};

const resetUpload = () => {
    selectedFiles.value = [];
    uploadSuccess.value = false;
    shareableLink.value = "";
    errorMessage.value = "";
    uploadProgress.value = 0;
};

const formatFileSize = (bytes: number): string => {
    if (bytes === 0) return "0 Bytes";
    const k = 1024;
    const sizes = ["Bytes", "KB", "MB", "GB"];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + " " + sizes[i];
};

const uploadFiles = async () => {
    if (selectedFiles.value.length === 0 || uploading.value) return;

    uploading.value = true;
    uploadProgress.value = 0;
    errorMessage.value = "";

    try {
        const formData = new FormData();
        selectedFiles.value.forEach((file) => {
            formData.append("files", file);
        });
        
        formData.append("validity", selectedValidity.value);

        const response = await axios.post(
            getApiUrl("/upload"),
            formData,
            {
                headers: {
                    "Content-Type": "multipart/form-data",
                    ...(isAuthenticated.value &&
                    localStorage.getItem("auth_token")
                        ? {
                              Authorization: `Bearer ${localStorage.getItem(
                                  "auth_token"
                              )}`,
                          }
                        : {}),
                },
                withCredentials: true,
                onUploadProgress: (progressEvent) => {
                    if (progressEvent.total) {
                        uploadProgress.value = Math.round(
                            (progressEvent.loaded * 100) / progressEvent.total
                        );
                    }
                },
            }
        );

        if (response.data.download_url) {
            uploadProgress.value = 100;
            uploadSuccess.value = true;

            const uploadId = response.data.download_url.split("/").pop();
            shareableLink.value = `${window.location.origin}/download/${uploadId}`;
        }
    } catch (error: any) {
        console.error("Upload error:", error);
        if (error.response?.status === 413) {
            errorMessage.value = "File too large. Maximum file size is 100MB.";
        } else {
            errorMessage.value = error.response?.data?.error || "Upload failed. Please try again.";
        }
        uploadProgress.value = 0;
    } finally {
        uploading.value = false;
    }
};

const copyToClipboard = async () => {
    try {
        await navigator.clipboard.writeText(shareableLink.value);
        copied.value = true;
        setTimeout(() => {
            copied.value = false;
        }, 2000);
    } catch (err) {
        console.error('Failed to copy:', err);
    }
};
</script>

<style scoped>
/* Custom scrollbar for file list */
.max-h-32::-webkit-scrollbar {
    width: 6px;
}

.max-h-32::-webkit-scrollbar-track {
    background: transparent;
}

.max-h-32::-webkit-scrollbar-thumb {
    background: #cbd5e0;
    border-radius: 3px;
}

.max-h-32::-webkit-scrollbar-thumb:hover {
    background: #a0aec0;
}

.dark .max-h-32::-webkit-scrollbar-thumb {
    background: #4a5568;
}

.dark .max-h-32::-webkit-scrollbar-thumb:hover {
    background: #2d3748;
}

/* Loading spinner animation */
@keyframes spin {
    to {
        transform: rotate(360deg);
    }
}

.animate-spin {
    animation: spin 1s linear infinite;
}

/* Smooth transitions */
.transition-all {
    transition-property: all;
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
    transition-duration: 200ms;
}

.transition-colors {
    transition-property: color, background-color, border-color;
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
    transition-duration: 200ms;
}

/* Glass morphism effect */
.backdrop-blur-xl {
    backdrop-filter: blur(20px);
}

.backdrop-blur-md {
    backdrop-filter: blur(12px);
}

/* Focus styles for accessibility */
button:focus,
input:focus,
select:focus {
    outline: 2px solid #3b82f6;
    outline-offset: 2px;
}
</style>
