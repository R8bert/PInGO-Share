<template>
    <div class="min-h-screen relative overflow-hidden">
        <!-- Beautiful Dynamic Background -->
        <div 
            class="absolute inset-0 z-0"
            :style="backgroundStyle"
        ></div>

        <!-- Main Content -->
        <div class="relative z-40 flex flex-col lg:flex-row items-center justify-between min-h-screen px-6 lg:px-12 gap-8 pt-20">
            <!-- Left Side - Tagline -->
            <div class="flex-1 max-w-2xl text-center lg:text-left">
                <h1 class="text-4xl sm:text-5xl lg:text-7xl font-bold text-white mb-6 leading-tight">
                    We can turn your
                    <br />
                    <span class="italic font-light">dream project</span>
                    <br />
                    into <em class="italic text-yellow-300">reality</em>
                </h1>
            </div>

            <!-- Right Side - Upload Modal -->
            <div class="flex-shrink-0 w-full max-w-md lg:max-w-lg">
                <!-- Success State -->
                <div v-if="uploadSuccess && shareableLink" class="bg-white/70 backdrop-blur-xl rounded-2xl p-8 shadow-2xl animate-success-bounce">
                    <div class="text-center mb-6">
                        <div class="w-16 h-16 bg-green-500 rounded-full flex items-center justify-center mx-auto mb-4 animate-success-pulse">
                            <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                            </svg>
                        </div>
                        <h3 class="text-2xl font-semibold text-gray-900 mb-2 animate-fade-in">Files uploaded! 🎉</h3>
                        <p class="text-gray-600 animate-fade-in-delay">Your files are ready to share</p>
                    </div>
                    
                    <div class="space-y-4">
                        <div class="flex space-x-2">
                            <input
                                v-model="shareableLink"
                                readonly
                                class="flex-1 px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl text-sm font-mono"
                            />
                            <button
                                @click="copyToClipboard"
                                class="px-4 py-3 bg-blue-500 hover:bg-blue-600 text-white rounded-xl transition-colors text-sm font-semibold"
                            >
                                {{ copied ? '✓' : 'Copy' }}
                            </button>
                        </div>
                        <button
                            @click="resetUpload"
                            class="w-full px-6 py-3 bg-gray-100 hover:bg-gray-200 text-gray-700 rounded-xl transition-colors font-medium"
                        >
                            Upload more files
                        </button>
                    </div>
                </div>

                <!-- Upload Form -->
                <div v-else class="bg-white/60 backdrop-blur-xl rounded-3xl p-8 shadow-2xl">
                    <!-- Header -->
                    <div class="mb-8">
                        <h2 class="text-2xl font-bold text-gray-900 mb-2 animate-fade-in-up">Let's talk! 👋</h2>
                        
                        <div class="mb-6">
                            <p class="text-gray-600 text-sm mb-2">Mail us at</p>
                            <a href="mailto:hello@pingoshare.com" class="text-blue-600 font-semibold hover:text-blue-700 transition-colors">hello@pingoshare.com</a>
                        </div>
                        
                        <div class="flex items-center mb-6">
                            <div class="flex-1 border-t border-gray-200"></div>
                            <span class="px-4 text-gray-400 text-sm">OR</span>
                            <div class="flex-1 border-t border-gray-200"></div>
                        </div>
                        
                        <p class="text-gray-600 text-sm mb-4">Leave us a brief message</p>
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
                                'border-2 border-dashed rounded-xl p-8 text-center cursor-pointer transition-all duration-200',
                                isDragging 
                                    ? 'border-blue-500 bg-blue-50 animate-drag-hover' 
                                    : 'border-gray-300 hover:border-blue-400'
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
                                <CloudArrowUpIcon class="w-12 h-12 text-white-400 mx-auto animate-gentle-pulse" />
                                <div>
                                    <p class="font-semibold text-gray-900">
                                        Drop files here or click to browse
                                    </p>
                                </div>
                            </div>
                        </div>

                        <!-- Selected Files -->
                        <div v-if="selectedFiles.length > 0" class="space-y-4">
                            <div class="bg-gray-50 rounded-xl p-4">
                                <div class="flex items-center justify-between mb-3">
                                    <span class="text-sm font-medium text-gray-900">
                                        {{ selectedFiles.length }} file{{ selectedFiles.length === 1 ? '' : 's' }} selected
                                    </span>
                                    <button @click="clearFiles" class="text-xs text-gray-500 hover:text-gray-700">
                                        Clear all
                                    </button>
                                </div>
                                
                                <div class="space-y-2 max-h-32 overflow-y-auto">
                                    <div
                                        v-for="(file, index) in selectedFiles"
                                        :key="index"
                                        class="flex items-center justify-between p-2 bg-white rounded-lg"
                                    >
                                        <div class="flex items-center space-x-3 flex-1 min-w-0">
                                            <img :src="getFileIconPath(file.name)" :alt="getFileIconAltText(file.name)" class="w-4 h-4 flex-shrink-0" />
                                            <div class="flex-1 min-w-0">
                                                <p class="text-sm font-medium text-gray-900 truncate">{{ file.name }}</p>
                                                <p class="text-xs text-gray-500">{{ formatFileSize(file.size) }}</p>
                                            </div>
                                        </div>
                                        <button @click="removeFile(index)" class="p-1 text-gray-400 hover:text-red-500 transition-colors">
                                            <XMarkIcon class="w-4 h-4" />
                                        </button>
                                    </div>
                                </div>
                            </div>

                            <!-- Expiration Settings -->
                            <div class="space-y-3">
                                <label class="block text-sm font-medium text-gray-700">Link expires in</label>
                                <select v-model="selectedValidity" class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-sm">
                                    <option value="1h">1 hour</option>
                                    <option value="1d">1 day</option>
                                    <option value="7d">7 days</option>
                                    <option value="30d">30 days</option>
                                </select>
                            </div>

                            <!-- Upload and Preview Buttons -->
                            <div class="flex space-x-3">
                                <button
                                    @click="uploadFiles"
                                    :disabled="uploading || selectedFiles.length === 0"
                                    class="flex-1 bg-black hover:bg-gray-800 disabled:bg-gray-400 disabled:cursor-not-allowed text-white font-semibold py-4 px-6 rounded-xl transition-all duration-200 flex items-center justify-center"
                                >
                                    <span v-if="uploading" class="flex items-center">
                                        <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" fill="none" viewBox="0 0 24 24">
                                            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                                            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                                        </svg>
                                        Uploading... {{ uploadProgress }}%
                                    </span>
                                    <span v-else>Upload</span>
                                </button>
                                
                                <button
                                    @click="showPreview = true"
                                    :disabled="selectedFiles.length === 0"
                                    class="px-4 py-4 bg-gray-100 hover:bg-gray-200 disabled:bg-gray-50 disabled:cursor-not-allowed text-gray-700 rounded-xl transition-all duration-200 flex items-center justify-center"
                                    title="Preview files"
                                >
                                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                                    </svg>
                                </button>
                            </div>
                        </div>

                        <!-- Default Send Button -->
                        <button
                            v-if="selectedFiles.length === 0"
                            @click="triggerFileSelect"
                            class="w-full bg-black hover:bg-gray-800 text-white font-semibold py-4 px-6 rounded-xl transition-all duration-200"
                        >
                            Send a message
                        </button>

                        <!-- Error Message -->
                        <div v-if="errorMessage" class="p-4 bg-red-50 border border-red-200 rounded-xl">
                            <div class="flex items-center space-x-2">
                                <svg class="w-5 h-5 text-red-500" fill="currentColor" viewBox="0 0 24 24">
                                    <path fill-rule="evenodd" d="M2.25 12c0-5.385 4.365-9.75 9.75-9.75s9.75 4.365 9.75 9.75-4.365 9.75-9.75 9.75S2.25 17.385 2.25 12zM12 8.25a.75.75 0 01.75.75v3.75a.75.75 0 01-1.5 0V9a.75.75 0 01.75-.75zm0 8.25a.75.75 0 100-1.5.75.75 0 000 1.5z" clip-rule="evenodd"/>
                                </svg>
                                <span class="text-sm text-red-700">{{ errorMessage }}</span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <!-- Preview Modal -->
        <div v-if="showPreview" class="fixed inset-0 z-50 flex items-center justify-center bg-black/60 backdrop-blur-md animate-fade-in" @click="showPreview = false">
            <div class="bg-white/75 backdrop-blur-xl rounded-2xl p-6 max-w-2xl w-full mx-4 max-h-[80vh] overflow-y-auto animate-modal-appear" @click.stop>
                <div class="flex items-center justify-between mb-6">
                    <h3 class="text-xl font-semibold text-gray-900">File Preview</h3>
                    <button @click="showPreview = false" class="p-2 text-gray-400 hover:text-gray-600 transition-colors">
                        <XMarkIcon class="w-6 h-6" />
                    </button>
                </div>
                
                <div class="space-y-4 max-h-96 overflow-y-auto">
                    <div
                        v-for="(file, index) in selectedFiles"
                        :key="index"
                        class="flex items-start space-x-4 p-4 bg-gray-50/80 rounded-xl backdrop-blur-sm animate-slide-in-up"
                        :style="{ animationDelay: `${index * 100}ms` }"
                    >
                        <img :src="getFileIconPath(file.name)" :alt="getFileIconAltText(file.name)" class="w-8 h-8 flex-shrink-0 mt-1" />
                        <div class="flex-1 min-w-0">
                            <h4 class="font-medium text-gray-900 truncate mb-1">{{ file.name }}</h4>
                            <div class="text-sm text-gray-600 space-y-1">
                                <p>Size: {{ formatFileSize(file.size) }}</p>
                                <p>Type: {{ file.type || 'Unknown' }}</p>
                                <p>Last modified: {{ new Date(file.lastModified).toLocaleDateString() }}</p>
                            </div>
                        </div>
                        <button 
                            @click="removeFile(index)" 
                            class="p-2 text-red-400 hover:text-red-600 hover:bg-red-50 rounded-lg transition-all duration-200"
                            title="Remove file"
                        >
                            <XMarkIcon class="w-5 h-5" />
                        </button>
                    </div>
                </div>
                
                <div class="mt-6 flex justify-between items-center pt-4 border-t border-gray-200">
                    <span class="text-sm text-gray-600">{{ selectedFiles.length }} file{{ selectedFiles.length === 1 ? '' : 's' }} selected</span>
                    <button 
                        @click="showPreview = false" 
                        class="px-6 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-colors font-medium"
                    >
                        Done
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { useAuth } from "../../composables/useAuth";
import { useIcons } from "../../composables/useIcons";
import { useSettings } from "../../composables/useSettings";
import { getApiUrl, getAssetUrl } from "../../utils/apiUtils";
import axios from "axios";
import {
    CloudArrowUpIcon,
    XMarkIcon,
} from "@heroicons/vue/24/outline";

const { isAuthenticated } = useAuth();
const { getFileIcon, getFileIconAlt } = useIcons();
const { settings, fetchSettings } = useSettings();

// Helper methods for file icons
const getFileIconPath = (filename: string) => {
    return getFileIcon(filename);
};

const getFileIconAltText = (filename: string) => {
    return getFileIconAlt(filename);
};

// Dynamic background from backend settings
const backgroundImage = computed(() => {
    if (settings.value?.backgroundImage) {
        return getAssetUrl(settings.value.backgroundImage);
    }
    // Fallback to a simple gradient if no background is set
    return null;
});

const backgroundStyle = computed(() => {
    if (backgroundImage.value) {
        return {
            backgroundImage: `url('${backgroundImage.value}')`,
            backgroundSize: 'cover',
            backgroundPosition: 'center',
            backgroundRepeat: 'no-repeat'
        };
    }
    // Fallback gradient
    return {
        background: 'linear-gradient(135deg, rgba(59, 130, 246, 0.9), rgba(147, 51, 234, 0.9), rgba(236, 72, 153, 0.9))'
    };
});

// Fetch settings on component mount
onMounted(() => {
    fetchSettings();
});

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
const showPreview = ref(false);

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

/* Loading spinner animation */
@keyframes spin {
    to {
        transform: rotate(360deg);
    }
}

.animate-spin {
    animation: spin 1s linear infinite;
}

/* Dynamic idle animations - subtle and calming */

@keyframes gentle-pulse {
    0%, 100% {
        opacity: 1;
        transform: scale(1);
    }
    50% {
        opacity: 0.9;
        transform: scale(1.01);
    }
}

@keyframes bounce-subtle {
    0%, 100% {
        transform: translateY(0);
    }
    50% {
        transform: translateY(-1px);
    }
}

@keyframes fade-in-up {
    from {
        opacity: 0;
        transform: translateY(20px);
    }
    to {
        opacity: 1;
        transform: translateY(0);
    }
}

@keyframes fade-in {
    from {
        opacity: 0;
    }
    to {
        opacity: 1;
    }
}

@keyframes fade-in-delay {
    0% {
        opacity: 0;
        transform: translateY(10px);
    }
    20% {
        opacity: 0;
        transform: translateY(10px);
    }
    100% {
        opacity: 1;
        transform: translateY(0);
    }
}

@keyframes success-bounce {
    0% {
        opacity: 0;
        transform: scale(0.8) translateY(20px);
    }
    60% {
        opacity: 1;
        transform: scale(1.05) translateY(0);
    }
    80% {
        transform: scale(0.98);
    }
    100% {
        transform: scale(1);
    }
}

@keyframes success-pulse {
    0% {
        transform: scale(1);
        box-shadow: 0 0 0 0 rgba(34, 197, 94, 0.7);
    }
    70% {
        transform: scale(1.05);
        box-shadow: 0 0 0 10px rgba(34, 197, 94, 0);
    }
    100% {
        transform: scale(1);
        box-shadow: 0 0 0 0 rgba(34, 197, 94, 0);
    }
}

@keyframes drag-hover {
    0%, 100% {
        transform: scale(1) rotate(0deg);
    }
    50% {
        transform: scale(1.02) rotate(1deg);
    }
}

@keyframes modal-appear {
    from {
        opacity: 0;
        transform: scale(0.9) translateY(-50px);
    }
    to {
        opacity: 1;
        transform: scale(1) translateY(0);
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

/* Gentle pulse animation for icons */
.animate-gentle-pulse {
    animation: gentle-pulse 3s ease-in-out infinite;
}

.animate-bounce-subtle {
    animation: bounce-subtle 2s ease-in-out infinite;
    animation-delay: 1s;
}

.animate-fade-in-up {
    animation: fade-in-up 0.6s ease-out;
}

.animate-fade-in {
    animation: fade-in 0.8s ease-out;
}

.animate-fade-in-delay {
    animation: fade-in-delay 1.2s ease-out;
}

.animate-success-bounce {
    animation: success-bounce 0.6s cubic-bezier(0.68, -0.55, 0.265, 1.55);
}

.animate-success-pulse {
    animation: success-pulse 2s infinite;
}

.animate-drag-hover {
    animation: drag-hover 0.3s ease-in-out;
}

.animate-modal-appear {
    animation: modal-appear 0.3s cubic-bezier(0.34, 1.56, 0.64, 1);
}

.animate-slide-in-up {
    animation: slide-in-up 0.5s ease-out forwards;
    opacity: 0;
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

.backdrop-blur-sm {
    backdrop-filter: blur(4px);
}

/* Focus styles for accessibility */
button:focus,
input:focus,
select:focus {
    outline: 2px solid #3b82f6;
    outline-offset: 2px;
}

/* Hover enhancements */
button:hover:not(:disabled) {
    transform: translateY(-2px);
}

.group:hover .animate-bounce-subtle {
    animation-duration: 1s;
}

/* Responsive text scaling */
@media (max-width: 640px) {
    h1 {
        font-size: 2.5rem;
        line-height: 1.2;
    }
}

@media (max-width: 480px) {
    h1 {
        font-size: 2rem;
        line-height: 1.2;
    }
}

/* Enhanced interaction states */
.cursor-pointer:hover {
    transform: scale(1.01);
}

/* Stagger animation delays for file list */
.animate-slide-in-up:nth-child(1) { animation-delay: 0ms; }
.animate-slide-in-up:nth-child(2) { animation-delay: 100ms; }
.animate-slide-in-up:nth-child(3) { animation-delay: 200ms; }
.animate-slide-in-up:nth-child(4) { animation-delay: 300ms; }
.animate-slide-in-up:nth-child(5) { animation-delay: 400ms; }
</style>
