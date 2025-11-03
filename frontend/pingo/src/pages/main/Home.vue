<template>
    <div class="min-h-screen flex" :class="isDark ? 'bg-neutral-900 text-white' : 'bg-neutral-50 text-neutral-900'">
        
        <!-- Left Sidebar - Upload Form -->
        <div class="w-full lg:w-[420px] flex flex-col relative z-20" 
            :class="isDark ? 'bg-neutral-900' : 'bg-white'">
            
            <!-- Logo/Brand -->
            <div class="p-8">
                <h2 class="text-2xl font-semibold tracking-tight">PInGO</h2>
            </div>

            <!-- Upload Form Content -->
            <div class="flex-1 flex flex-col px-8 pb-8">
                <div class="flex-1 flex flex-col justify-center max-w-md">
                    
                    <!-- Heading -->
                    <div class="mb-8">
                        <h1 class="text-4xl font-semibold mb-3 tracking-tight" 
                            :class="isDark ? 'text-white' : 'text-neutral-900'">
                            Transfer files
                        </h1>
                        <p class="text-base" :class="isDark ? 'text-neutral-400' : 'text-neutral-600'">
                            Up to {{ formatFileSize(maxUploadSize) }}, free
                        </p>
                    </div>

                    <!-- Upload Section -->
                    <div v-if="!uploadComplete" class="animate-fade-in-up">
                        <!-- Request Files Button -->
                        <div class="mb-4 flex gap-3">
                            <button @click="triggerFileInput"
                                class="flex-1 px-6 py-4 rounded-lg font-medium transition-all duration-300 text-white bg-gradient-to-r from-blue-600 to-blue-500 hover:from-blue-700 hover:to-blue-600 shadow-lg shadow-blue-500/30 hover:shadow-xl hover:shadow-blue-500/40 hover:scale-[1.02] active:scale-[0.98]">
                                <span class="flex items-center justify-center gap-2">
                                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
                                    </svg>
                                    Add files
                                </span>
                            </button>
                        </div>

                        <!-- File Size Counter -->
                        <div v-if="selectedFiles.length > 0" class="mb-4 p-3 rounded-lg backdrop-blur-sm animate-slide-down"
                            :class="totalSize > maxUploadSize ? 'bg-red-500/10 border border-red-500/20' : isDark ? 'bg-neutral-800/50 border border-neutral-700/50' : 'bg-blue-50/50 border border-blue-200/50'">
                            <div class="flex items-center justify-between text-sm">
                                <span :class="totalSize > maxUploadSize ? 'text-red-500' : isDark ? 'text-neutral-300' : 'text-neutral-700'">
                                    <span class="font-semibold">{{ selectedFiles.length }}</span> file{{ selectedFiles.length !== 1 ? 's' : '' }} selected
                                </span>
                                <span class="font-mono" :class="totalSize > maxUploadSize ? 'text-red-500 font-bold' : isDark ? 'text-neutral-400' : 'text-neutral-600'">
                                    {{ formatFileSize(totalSize) }} / {{ formatFileSize(maxUploadSize) }}
                                </span>
                            </div>
                            <div class="mt-2 h-1.5 rounded-full overflow-hidden" :class="isDark ? 'bg-neutral-700' : 'bg-neutral-200'">
                                <div class="h-full transition-all duration-500 rounded-full"
                                    :class="totalSize > maxUploadSize ? 'bg-red-500' : 'bg-gradient-to-r from-blue-500 to-blue-400'"
                                    :style="{ width: `${Math.min((totalSize / maxUploadSize) * 100, 100)}%` }">
                                </div>
                            </div>
                        </div>

                        <!-- Hidden File Input -->
                        <input ref="fileInput" type="file" multiple @change="onFileChange" class="hidden" />

                        <!-- Drag and Drop Zone -->
                        <div v-if="selectedFiles.length === 0"
                            @drop.prevent="onDrop"
                            @dragover.prevent="isDragging = true"
                            @dragleave.prevent="isDragging = false"
                            class="relative mt-6 p-16 rounded-xl border-2 border-dashed cursor-pointer transition-all duration-300 group overflow-hidden"
                            :class="[
                                isDragging ? 'border-blue-500 bg-blue-500/10 scale-[1.02]' : isDark ? 'border-neutral-700 hover:border-neutral-600 hover:bg-neutral-800/30' : 'border-neutral-300 hover:border-blue-400 hover:bg-blue-50/30',
                            ]"
                            @click="triggerFileInput">
                            
                            <!-- Animated background effect -->
                            <div class="absolute inset-0 bg-gradient-to-br from-blue-500/5 via-transparent to-purple-500/5 opacity-0 group-hover:opacity-100 transition-opacity duration-500"></div>
                            
                            <!-- Upload icon with animation -->
                            <div class="relative flex flex-col items-center">
                                <div class="mb-4 p-4 rounded-2xl transition-all duration-300"
                                    :class="isDragging ? 'bg-blue-500/20 scale-110' : isDark ? 'bg-neutral-800' : 'bg-neutral-200 group-hover:bg-blue-100'">
                                    <svg class="w-10 h-10 transition-all duration-300"
                                        :class="isDragging ? 'text-blue-500 animate-bounce' : isDark ? 'text-neutral-400 group-hover:text-blue-500' : 'text-neutral-500 group-hover:text-blue-600'"
                                        fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                                            d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"></path>
                                    </svg>
                                </div>
                                <p class="text-center font-medium transition-colors duration-300" 
                                    :class="isDragging ? 'text-blue-500' : isDark ? 'text-neutral-300 group-hover:text-white' : 'text-neutral-700 group-hover:text-blue-600'">
                                    {{ isDragging ? 'Drop files here' : 'Drop files here' }}
                                </p>
                                <p class="text-sm text-center mt-2" :class="isDark ? 'text-neutral-500' : 'text-neutral-500'">
                                    or click to browse
                                </p>
                            </div>
                        </div>

                        <!-- File List with enhanced animations -->
                        <div v-if="selectedFiles.length > 0" class="mt-6 space-y-2 max-h-[300px] overflow-y-auto custom-scrollbar">
                            <div v-for="(file, index) in selectedFiles" :key="index"
                                class="group flex items-center gap-3 p-4 rounded-xl border transition-all duration-300 hover:scale-[1.01] animate-slide-in-left backdrop-blur-sm"
                                :style="{ animationDelay: `${index * 50}ms` }"
                                :class="isDark ? 'bg-neutral-800/80 border-neutral-700/50 hover:bg-neutral-800 hover:border-neutral-600 hover:shadow-lg hover:shadow-neutral-900/20' : 'bg-white/80 border-neutral-200 hover:bg-white hover:border-blue-300 hover:shadow-lg hover:shadow-blue-500/10'">
                                
                                <!-- File type icon with gradient -->
                                <div class="relative w-12 h-12 rounded-xl flex items-center justify-center text-xs font-bold uppercase flex-shrink-0 overflow-hidden group-hover:scale-110 transition-transform duration-300"
                                    :class="getFileTypeColor(file)">
                                    <div class="absolute inset-0 bg-gradient-to-br from-white/20 to-transparent"></div>
                                    <span class="relative z-10">{{ getFileExtension(file).slice(0, 3) }}</span>
                                </div>
                                
                                <!-- File info -->
                                <div class="flex-1 min-w-0">
                                    <p class="text-sm font-semibold truncate transition-colors duration-300" 
                                        :class="isDark ? 'text-neutral-100 group-hover:text-white' : 'text-neutral-900'">
                                        {{ file.name }}
                                    </p>
                                    <div class="flex items-center gap-2 mt-1">
                                        <p class="text-xs" :class="isDark ? 'text-neutral-500' : 'text-neutral-600'">
                                            {{ formatFileSize(file.size) }}
                                        </p>
                                        <span class="text-xs" :class="isDark ? 'text-neutral-600' : 'text-neutral-400'">•</span>
                                        <p class="text-xs" :class="isDark ? 'text-neutral-500' : 'text-neutral-600'">
                                            {{ getFileType(file) }}
                                        </p>
                                    </div>
                                </div>
                                
                                <!-- Remove button with animation -->
                                <button @click.stop="removeFile(index)" 
                                    class="w-9 h-9 rounded-lg flex items-center justify-center transition-all duration-300 opacity-0 group-hover:opacity-100 hover:scale-110 active:scale-95"
                                    :class="isDark ? 'hover:bg-red-500/20 text-neutral-400 hover:text-red-400' : 'hover:bg-red-50 text-neutral-500 hover:text-red-600'">
                                    <XMarkIcon class="w-5 h-5" />
                                </button>
                            </div>
                        </div>

                        <!-- Options -->
                        <div v-if="selectedFiles.length > 0" class="mt-6 space-y-3">
                            <!-- Email to -->
                            <div>
                                <label class="block text-sm mb-2" :class="isDark ? 'text-neutral-400' : 'text-neutral-600'">
                                    Email to
                                </label>
                                <input v-model="recipientEmail" type="email" placeholder="name@email.com"
                                    class="w-full px-4 py-3 rounded-lg border outline-none focus:ring-2 focus:ring-blue-500 transition-all bg-transparent text-sm"
                                    :class="isDark ? 'border-neutral-700 focus:border-transparent' : 'border-neutral-300 focus:border-transparent'" />
                            </div>

                            <!-- Message -->
                            <div>
                                <label class="block text-sm mb-2" :class="isDark ? 'text-neutral-400' : 'text-neutral-600'">
                                    Message
                                </label>
                                <textarea v-model="transferMessage" placeholder="Add a message..." rows="3"
                                    class="w-full px-4 py-3 rounded-lg border outline-none focus:ring-2 focus:ring-blue-500 transition-all resize-none bg-transparent text-sm"
                                    :class="isDark ? 'border-neutral-700 focus:border-transparent' : 'border-neutral-300 focus:border-transparent'"></textarea>
                            </div>

                            <!-- Password & Validity -->
                            <div class="grid grid-cols-2 gap-3">
                                <div>
                                    <label class="block text-sm mb-2" :class="isDark ? 'text-neutral-400' : 'text-neutral-600'">
                                        Password
                                    </label>
                                    <div class="relative">
                                        <input v-model="uploadPassword" :type="showPassword ? 'text' : 'password'" 
                                            placeholder="Optional"
                                            class="w-full px-4 py-3 rounded-lg border outline-none focus:ring-2 focus:ring-blue-500 transition-all bg-transparent text-sm pr-10"
                                            :class="isDark ? 'border-neutral-700 focus:border-transparent' : 'border-neutral-300 focus:border-transparent'" />
                                        <button @click="showPassword = !showPassword"
                                            class="absolute right-3 top-1/2 -translate-y-1/2 transition-opacity"
                                            :class="isDark ? 'text-neutral-500 hover:text-neutral-400' : 'text-neutral-400 hover:text-neutral-600'">
                                            <EyeIcon v-if="!showPassword" class="w-4 h-4" />
                                            <EyeSlashIcon v-else class="w-4 h-4" />
                                        </button>
                                    </div>
                                </div>
                                
                                <div>
                                    <label class="block text-sm mb-2" :class="isDark ? 'text-neutral-400' : 'text-neutral-600'">
                                        Expiry
                                    </label>
                                    <select v-model="selectedValidity" 
                                        class="w-full px-4 py-3 rounded-lg border outline-none focus:ring-2 focus:ring-blue-500 transition-all bg-transparent cursor-pointer text-sm"
                                        :class="isDark ? 'border-neutral-700 focus:border-transparent' : 'border-neutral-300 focus:border-transparent'">
                                        <option v-for="option in validityOptions" :key="option.value" :value="option.value"
                                            :class="isDark ? 'bg-neutral-900' : 'bg-white'">
                                            {{ option.label }}
                                        </option>
                                    </select>
                                </div>
                            </div>

                            <!-- Transfer Button with enhanced styling -->
                            <button @click="uploadFiles" :disabled="isUploading || totalSize > maxUploadSize"
                                class="relative w-full py-4 rounded-xl font-semibold transition-all duration-300 text-white overflow-hidden group disabled:opacity-50 disabled:cursor-not-allowed"
                                :class="totalSize > maxUploadSize ? 'bg-red-500 cursor-not-allowed' : isUploading ? 'bg-gradient-to-r from-blue-600 to-blue-500' : 'bg-gradient-to-r from-blue-600 to-blue-500 hover:from-blue-700 hover:to-blue-600 shadow-lg shadow-blue-500/30 hover:shadow-xl hover:shadow-blue-500/40 hover:scale-[1.02] active:scale-[0.98]'">
                                
                                <!-- Shimmer effect -->
                                <div v-if="!isUploading && totalSize <= maxUploadSize" class="absolute inset-0 opacity-0 group-hover:opacity-100 transition-opacity duration-500">
                                    <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/20 to-transparent animate-shimmer" style="background-size: 200% 100%;"></div>
                                </div>
                                
                                <span class="relative flex items-center justify-center gap-2">
                                    <svg v-if="isUploading" class="w-5 h-5 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                                        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                                    </svg>
                                    <svg v-else-if="totalSize > maxUploadSize" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"></path>
                                    </svg>
                                    <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
                                    </svg>
                                    {{ totalSize > maxUploadSize ? 'Files too large' : isUploading ? 'Transferring...' : 'Transfer' }}
                                </span>
                            </button>

                            <!-- Enhanced Progress -->
                            <div v-if="isUploading" class="mt-4 space-y-3 animate-slide-down">
                                <!-- Progress stats -->
                                <div class="flex justify-between items-center text-sm">
                                    <div class="flex items-center gap-2">
                                        <div class="w-2 h-2 rounded-full bg-blue-500 animate-pulse"></div>
                                        <span class="font-medium" :class="isDark ? 'text-neutral-300' : 'text-neutral-700'">
                                            Uploading...
                                        </span>
                                    </div>
                                    <span class="font-mono font-semibold" :class="isDark ? 'text-blue-400' : 'text-blue-600'">
                                        {{ progress }}%
                                    </span>
                                </div>

                                <!-- Progress bar with gradient -->
                                <div class="relative h-2 rounded-full overflow-hidden" :class="isDark ? 'bg-neutral-800' : 'bg-neutral-200'">
                                    <div class="absolute inset-0 bg-gradient-to-r from-blue-600 via-blue-500 to-blue-400 transition-all duration-300 animate-glow" 
                                        :style="{ width: `${progress}%` }">
                                        <!-- Shimmer effect on progress bar -->
                                        <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/30 to-transparent" 
                                            style="animation: shimmer 2s linear infinite; background-size: 200% 100%;"></div>
                                    </div>
                                </div>

                                <!-- Upload statistics -->
                                <div class="grid grid-cols-2 gap-3 text-xs">
                                    <div class="p-3 rounded-lg backdrop-blur-sm" :class="isDark ? 'bg-neutral-800/50' : 'bg-blue-50/50'">
                                        <div class="flex items-center gap-2 mb-1">
                                            <svg class="w-4 h-4" :class="isDark ? 'text-blue-400' : 'text-blue-600'" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
                                            </svg>
                                            <span :class="isDark ? 'text-neutral-500' : 'text-neutral-600'">Speed</span>
                                        </div>
                                        <p class="font-mono font-semibold" :class="isDark ? 'text-neutral-300' : 'text-neutral-900'">
                                            {{ uploadSpeed > 0 ? formatFileSize(uploadSpeed) + '/s' : 'Calculating...' }}
                                        </p>
                                    </div>
                                    
                                    <div class="p-3 rounded-lg backdrop-blur-sm" :class="isDark ? 'bg-neutral-800/50' : 'bg-blue-50/50'">
                                        <div class="flex items-center gap-2 mb-1">
                                            <svg class="w-4 h-4" :class="isDark ? 'text-blue-400' : 'text-blue-600'" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                                            </svg>
                                            <span :class="isDark ? 'text-neutral-500' : 'text-neutral-600'">Time left</span>
                                        </div>
                                        <p class="font-mono font-semibold" :class="isDark ? 'text-neutral-300' : 'text-neutral-900'">
                                            {{ estimatedTimeRemaining > 0 ? (estimatedTimeRemaining < 60 ? Math.ceil(estimatedTimeRemaining) + 's' : Math.ceil(estimatedTimeRemaining / 60) + 'm') : 'Calculating...' }}
                                        </p>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>

                    <!-- Success State -->
                    <div v-else class="flex flex-col items-center justify-center py-12">
                        <div class="w-16 h-16 rounded-full bg-green-500 flex items-center justify-center mb-4">
                            <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                            </svg>
                        </div>
                        
                        <h3 class="text-xl font-semibold mb-2">Transfer complete</h3>
                        <p class="text-sm mb-6" :class="isDark ? 'text-neutral-400' : 'text-neutral-600'">
                            Your files are ready to share
                        </p>

                        <div class="w-full space-y-3">
                            <div class="p-4 rounded-lg border flex items-center gap-3"
                                :class="isDark ? 'bg-neutral-800/50 border-neutral-700' : 'bg-neutral-100 border-neutral-300'">
                                <input :value="uploadLink" readonly 
                                    class="flex-1 bg-transparent outline-none text-sm truncate" />
                                <button @click="copyLink"
                                    class="px-4 py-2 rounded-lg text-sm font-medium transition-colors"
                                    :class="isDark ? 'bg-neutral-700 hover:bg-neutral-600' : 'bg-neutral-200 hover:bg-neutral-300'">
                                    {{ copied ? 'Copied!' : 'Copy' }}
                                </button>
                            </div>

                            <button @click="uploadNew"
                                class="w-full py-3 rounded-lg font-medium transition-colors text-white bg-blue-600 hover:bg-blue-700">
                                Transfer new files
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <!-- Right Side - Hero Image/Content -->
        <div class="hidden lg:flex flex-1 relative overflow-hidden"
            :class="isDark ? 'bg-neutral-800' : 'bg-neutral-200'">
            
            <!-- Placeholder gradient -->
            <div class="absolute inset-0 bg-gradient-to-br from-blue-500 via-purple-500 to-pink-500 opacity-20"></div>
            
            <!-- Optional: Add background image -->
            <div class="absolute inset-0 flex items-center justify-center">
                <div class="text-center p-12 max-w-lg">
                    <h2 class="text-5xl font-bold mb-4" :class="isDark ? 'text-white' : 'text-neutral-900'">
                        Simple, secure file sharing
                    </h2>
                    <p class="text-lg" :class="isDark ? 'text-neutral-400' : 'text-neutral-600'">
                        Transfer files up to {{ formatFileSize(maxUploadSize) }} with end-to-end encryption
                    </p>
                </div>
            </div>
        </div>

    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed } from "vue";
import { useAuth } from "../../composables/useAuth";
import { useTheme } from "../../composables/useTheme";
import axios from "axios";
import {
    XMarkIcon,
    EyeIcon,
    EyeSlashIcon,
} from "@heroicons/vue/24/outline";

const { isAuthenticated } = useAuth();
const { isDark } = useTheme();

// Refs
const fileInput = ref<HTMLInputElement | null>(null);
const uploadSection = ref<HTMLElement | null>(null);
const { getSettings } = useAuth();
// State
const selectedFiles = ref<File[]>([]);
const isDragging = ref(false);
const isUploading = ref(false);
const progress = ref(0);
const uploadComplete = ref(false);

// New upload options
const uploadPassword = ref('');
const transferMessage = ref('');
const recipientEmail = ref('');
const showPassword = ref(false);

// Upload statistics
const uploadSpeed = ref(0);
const timeRemaining = ref('Calculating...');
const estimatedTimeRemaining = ref(0);
const startTime = ref(0);
const lastLoaded = ref(0);
const lastTime = ref(0);
const uploadedUrl = ref("");
const uploadLink = ref("");
const copied = ref(false);
const message = ref({ text: "", type: "success" as "success" | "error" });
const previewingFiles = ref<Set<number>>(new Set());
const previewUrls = ref<Map<number, string>>(new Map());
const textPreviews = ref<Map<number, string>>(new Map());
const maxUploadSize = ref<number>(104857600); // Default 100 MB

const validityOptions = ref([
    { value: "7days", label: "7 Days", description: "One week" },
    { value: "1month", label: "1 Month", description: "30 days" },
    { value: "6months", label: "6 Months", description: "Half year" },
    { value: "1year", label: "1 Year", description: "12 months" },
    { value: "never", label: "Never", description: "Permanent" },
]);
const selectedValidity = ref("7days");
const maxAllowedValidity = ref("7days");

// Computed properties
const totalSize = computed(() => {
    return selectedFiles.value.reduce((total, file) => total + file.size, 0);
});

// Helper methods
const getFileType = (file: File): string => {
    const ext = getFileExtension(file).toLowerCase();
    const types: { [key: string]: string } = {
        // Images
        jpg: 'Image', jpeg: 'Image', png: 'Image', gif: 'Image', webp: 'Image', svg: 'Image',
        // Documents
        pdf: 'Document', doc: 'Document', docx: 'Document', txt: 'Text', rtf: 'Document',
        // Spreadsheets
        xls: 'Spreadsheet', xlsx: 'Spreadsheet', csv: 'Spreadsheet',
        // Presentations
        ppt: 'Presentation', pptx: 'Presentation',
        // Archives
        zip: 'Archive', rar: 'Archive', '7z': 'Archive', tar: 'Archive', gz: 'Archive',
        // Videos
        mp4: 'Video', avi: 'Video', mov: 'Video', mkv: 'Video', webm: 'Video',
        // Audio
        mp3: 'Audio', wav: 'Audio', ogg: 'Audio', m4a: 'Audio',
        // Code
        js: 'Code', ts: 'Code', py: 'Code', java: 'Code', cpp: 'Code', html: 'Code', css: 'Code',
    };
    return types[ext] || 'File';
};

const getFileTypeColor = (file: File): string => {
    const type = getFileType(file);
    const colors: { [key: string]: string } = {
        'Image': isDark.value ? 'bg-gradient-to-br from-purple-600 to-pink-600 text-white' : 'bg-gradient-to-br from-purple-500 to-pink-500 text-white',
        'Document': isDark.value ? 'bg-gradient-to-br from-blue-600 to-cyan-600 text-white' : 'bg-gradient-to-br from-blue-500 to-cyan-500 text-white',
        'Spreadsheet': isDark.value ? 'bg-gradient-to-br from-green-600 to-emerald-600 text-white' : 'bg-gradient-to-br from-green-500 to-emerald-500 text-white',
        'Presentation': isDark.value ? 'bg-gradient-to-br from-orange-600 to-red-600 text-white' : 'bg-gradient-to-br from-orange-500 to-red-500 text-white',
        'Archive': isDark.value ? 'bg-gradient-to-br from-yellow-600 to-amber-600 text-white' : 'bg-gradient-to-br from-yellow-500 to-amber-500 text-white',
        'Video': isDark.value ? 'bg-gradient-to-br from-indigo-600 to-purple-600 text-white' : 'bg-gradient-to-br from-indigo-500 to-purple-500 text-white',
        'Audio': isDark.value ? 'bg-gradient-to-br from-pink-600 to-rose-600 text-white' : 'bg-gradient-to-br from-pink-500 to-rose-500 text-white',
        'Code': isDark.value ? 'bg-gradient-to-br from-slate-600 to-gray-600 text-white' : 'bg-gradient-to-br from-slate-500 to-gray-500 text-white',
        'Text': isDark.value ? 'bg-gradient-to-br from-teal-600 to-cyan-600 text-white' : 'bg-gradient-to-br from-teal-500 to-cyan-500 text-white',
    };
    return colors[type] || (isDark.value ? 'bg-gradient-to-br from-neutral-600 to-neutral-700 text-white' : 'bg-gradient-to-br from-neutral-400 to-neutral-500 text-white');
};

// Methods
const triggerFileInput = () => {
    fileInput.value?.click();
};

const onFileChange = (event: Event) => {
    const target = event.target as HTMLInputElement;
    if (target.files) {
        selectedFiles.value = Array.from(target.files);
    }
};

const onDrop = (event: DragEvent) => {
    isDragging.value = false;
    if (event.dataTransfer?.files) {
        selectedFiles.value = Array.from(event.dataTransfer.files);
    }
};

const removeFile = (index: number) => {
    // Clean up preview URLs and refs when removing file
    if (previewingFiles.value.has(index)) {
        const url = previewUrls.value.get(index);
        if (url) URL.revokeObjectURL(url);
        previewUrls.value.delete(index);
        textPreviews.value.delete(index);
        previewingFiles.value.delete(index);
    }

    selectedFiles.value.splice(index, 1);

    // Re-index remaining files
    const newPreviewingFiles = new Set<number>();
    const newPreviewUrls = new Map<number, string>();
    const newTextPreviews = new Map<number, string>();

    previewingFiles.value.forEach((previewIndex) => {
        if (previewIndex > index) {
            newPreviewingFiles.add(previewIndex - 1);
            const url = previewUrls.value.get(previewIndex);
            if (url) newPreviewUrls.set(previewIndex - 1, url);
            const textPreview = textPreviews.value.get(previewIndex);
            if (textPreview) newTextPreviews.set(previewIndex - 1, textPreview);
        } else if (previewIndex < index) {
            newPreviewingFiles.add(previewIndex);
            const url = previewUrls.value.get(previewIndex);
            if (url) newPreviewUrls.set(previewIndex, url);
            const textPreview = textPreviews.value.get(previewIndex);
            if (textPreview) newTextPreviews.set(previewIndex, textPreview);
        }
    });

    previewingFiles.value = newPreviewingFiles;
    previewUrls.value = newPreviewUrls;
    textPreviews.value = newTextPreviews;
};

const togglePreview = (index: number) => {
    if (previewingFiles.value.has(index)) {
        // Hide preview
        previewingFiles.value.delete(index);
        const url = previewUrls.value.get(index);
        if (url) {
            URL.revokeObjectURL(url);
            previewUrls.value.delete(index);
        }
        textPreviews.value.delete(index);
    } else {
        // Show preview
        previewingFiles.value.add(index);
        const file = selectedFiles.value[index];
        const fileExtension = getFileExtension(file);

        if (["txt", "md", "json", "csv", "xml"].includes(fileExtension)) {
            // Handle text file preview
            const reader = new FileReader();
            reader.onload = (e) => {
                const content = e.target?.result as string;
                const truncatedContent =
                    content.length > 1000
                        ? content.substring(0, 1000) + "\n\n... (truncated)"
                        : content;
                textPreviews.value.set(index, truncatedContent);
            };
            reader.readAsText(file);
        } else {
            // Handle other file types (images, videos, audio)
            createPreviewUrl(file, index);
        }
    }
};

const createPreviewUrl = (file: File, index: number) => {
    const url = URL.createObjectURL(file);
    previewUrls.value.set(index, url);
};

const canPreview = (file: File): boolean => {
    const ext = getFileExtension(file);
    return [
        "mp4",
        "pdf",
        "jpg",
        "jpeg",
        "png",
        "gif",
        "webp",
        "txt",
        "md",
        "json",
        "csv",
        "xml",
        "torrent",
        "mp3",
        "wav",
        "flac",
    ].includes(ext);
};

const formatFileSize = (bytes: number): string => {
    if (bytes === 0) return "0 Bytes";
    const k = 1024;
    const sizes = ["Bytes", "KB", "MB", "GB"];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + " " + sizes[i];
};

const getFileExtension = (file: File): string => {
    return file.name.split(".").pop()?.toLowerCase() || "";
};

const uploadFiles = async () => {
    if (selectedFiles.value.length === 0 || isUploading.value) return;

    isUploading.value = true;
    progress.value = 0;
    startTime.value = Date.now();
    lastLoaded.value = 0;
    lastTime.value = Date.now();
    uploadSpeed.value = 0;
    estimatedTimeRemaining.value = 0;
    timeRemaining.value = 'Calculating...';

    try {
        const formData = new FormData();
        selectedFiles.value.forEach((file) => {
            formData.append("files", file);
        });

        // Add upload options
        if (uploadPassword.value) {
            formData.append("password", uploadPassword.value);
        }
        if (transferMessage.value) {
            formData.append("message", transferMessage.value);
        }
        if (recipientEmail.value) {
            formData.append("recipient_email", recipientEmail.value);
        }
        if (selectedValidity.value) {
            formData.append("validity_hours", selectedValidity.value.toString());
        }

        const response = await axios.post(
            "http://localhost:8080/upload",
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
                        progress.value = Math.round(
                            (progressEvent.loaded * 100) / progressEvent.total
                        );

                        // Calculate upload speed and time remaining
                        const currentTime = Date.now();
                        const timeElapsed = (currentTime - lastTime.value) / 1000; // seconds
                        
                        if (timeElapsed > 0.5) { // Update every 0.5 seconds
                            const bytesUploaded = progressEvent.loaded - lastLoaded.value;
                            const speed = bytesUploaded / timeElapsed; // bytes per second
                            uploadSpeed.value = speed;

                            const remainingBytes = progressEvent.total - progressEvent.loaded;
                            const remainingSeconds = remainingBytes / speed;
                            estimatedTimeRemaining.value = remainingSeconds;
                            
                            if (remainingSeconds < 60) {
                                timeRemaining.value = `${Math.ceil(remainingSeconds)}s remaining`;
                            } else if (remainingSeconds < 3600) {
                                timeRemaining.value = `${Math.ceil(remainingSeconds / 60)}m remaining`;
                            } else {
                                timeRemaining.value = `${Math.ceil(remainingSeconds / 3600)}h remaining`;
                            }

                            lastLoaded.value = progressEvent.loaded;
                            lastTime.value = currentTime;
                        }
                    }
                },
            }
        );

        if (response.data.download_url) {
            progress.value = 100;
            uploadComplete.value = true;

            // Extract upload ID from download_url and create full URL
            const uploadId = response.data.download_url.split("/").pop();
            const fullUrl = `${window.location.origin}/download/${uploadId}`;
            uploadedUrl.value = fullUrl;
            uploadLink.value = fullUrl;

            message.value = {
                text: "Files uploaded successfully!",
                type: "success",
            };

            // Scroll to center the upload success section
            setTimeout(() => {
                uploadSection.value?.scrollIntoView({
                    behavior: "smooth",
                    block: "center",
                });
            }, 100);

            // Clean up preview URLs
            previewUrls.value.forEach((url) => URL.revokeObjectURL(url));
            previewUrls.value.clear();
            textPreviews.value.clear();
            previewingFiles.value.clear();
        }
    } catch (error) {
        console.error("Upload error:", error);
        message.value = {
            text: "Upload failed. Please try again.",
            type: "error",
        };
        progress.value = 0;
    } finally {
        isUploading.value = false;

        // Clear message after 5 seconds instead of 3
        setTimeout(() => {
            message.value = { text: "", type: "success" };
            if (!isUploading.value) progress.value = 0;
        }, 5000);
    }
};

// Action buttons for upload success
const uploadNew = () => {
    uploadComplete.value = false;
    uploadedUrl.value = "";
    selectedFiles.value = [];
    progress.value = 0;
    message.value = { text: "", type: "success" };

    // Clean up any remaining preview URLs
    previewUrls.value.forEach((url) => URL.revokeObjectURL(url));
    previewUrls.value.clear();
    textPreviews.value.clear();
    previewingFiles.value.clear();
};

const copyLink = async () => {
    try {
        await navigator.clipboard.writeText(uploadLink.value);
        copied.value = true;
        setTimeout(() => {
            copied.value = false;
        }, 2000);
    } catch (err) {
        message.value = { text: "Failed to copy link", type: "error" };
        setTimeout(() => {
            message.value = { text: "", type: "success" };
        }, 3000);
    }
};

const loadSettings = async () => {
    try {
        const settings = await getSettings();
        maxUploadSize.value = settings.maxUploadSize || 104857600;
        maxAllowedValidity.value = settings.maxValidity || "7days";

        // Filter validity options based on max allowed
        const validityOrder = ["7days", "1month", "6months", "1year", "never"];
        const maxIndex = validityOrder.indexOf(maxAllowedValidity.value);

        if (maxIndex !== -1) {
            const allowedOptions = validityOrder.slice(0, maxIndex + 1);
            validityOptions.value = validityOptions.value.filter((option) =>
                allowedOptions.includes(option.value)
            );
        }

        // Set default validity to first available option
        if (validityOptions.value.length > 0) {
            selectedValidity.value = validityOptions.value[0].value;
        }
    } catch (error) {
        console.error("Error loading settings:", error);
    }
};

onMounted(() => {
    // Scroll to top smoothly when page loads
    loadSettings();
    window.scrollTo({ top: 0, behavior: "smooth" });
});

onUnmounted(() => {
    // Clean up preview URLs
    previewUrls.value.forEach((url) => URL.revokeObjectURL(url));
});
</script>

<style scoped>
/* Custom scrollbar */
.custom-scrollbar::-webkit-scrollbar {
    width: 6px;
}

.custom-scrollbar::-webkit-scrollbar-track {
    background: transparent;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
    background: rgba(156, 163, 175, 0.3);
    border-radius: 3px;
}

.custom-scrollbar::-webkit-scrollbar-thumb:hover {
    background: rgba(156, 163, 175, 0.5);
}

/* Smooth scrolling */
html {
    scroll-behavior: smooth;
}

/* Fade in up animation */
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

/* Slide down animation */
@keyframes slide-down {
    from {
        opacity: 0;
        transform: translateY(-10px);
    }
    to {
        opacity: 1;
        transform: translateY(0);
    }
}

/* Slide in from left */
@keyframes slide-in-left {
    from {
        opacity: 0;
        transform: translateX(-20px);
    }
    to {
        opacity: 1;
        transform: translateX(0);
    }
}

/* Scale in animation */
@keyframes scale-in {
    from {
        opacity: 0;
        transform: scale(0.95);
    }
    to {
        opacity: 1;
        transform: scale(1);
    }
}

/* Pulse animation */
@keyframes pulse {
    0%, 100% {
        opacity: 1;
    }
    50% {
        opacity: 0.5;
    }
}

/* Shimmer effect */
@keyframes shimmer {
    0% {
        background-position: -1000px 0;
    }
    100% {
        background-position: 1000px 0;
    }
}

/* Bounce animation */
@keyframes bounce {
    0%, 100% {
        transform: translateY(0);
    }
    50% {
        transform: translateY(-10px);
    }
}

/* Spin animation */
@keyframes spin {
    from {
        transform: rotate(0deg);
    }
    to {
        transform: rotate(360deg);
    }
}

/* Float animation */
@keyframes float {
    0%, 100% {
        transform: translateY(0px);
    }
    50% {
        transform: translateY(-10px);
    }
}

/* Glow animation */
@keyframes glow {
    0%, 100% {
        box-shadow: 0 0 5px rgba(59, 130, 246, 0.5);
    }
    50% {
        box-shadow: 0 0 20px rgba(59, 130, 246, 0.8), 0 0 30px rgba(59, 130, 246, 0.6);
    }
}

/* Animation classes */
.animate-fade-in-up {
    animation: fade-in-up 0.6s ease-out;
}

.animate-slide-down {
    animation: slide-down 0.4s ease-out;
}

.animate-slide-in-left {
    animation: slide-in-left 0.5s ease-out;
    animation-fill-mode: both;
}

.animate-scale-in {
    animation: scale-in 0.5s ease-out;
}

.animate-pulse {
    animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}

.animate-bounce {
    animation: bounce 1s infinite;
}

.animate-spin {
    animation: spin 1s linear infinite;
}

.animate-float {
    animation: float 3s ease-in-out infinite;
}

.animate-glow {
    animation: glow 2s ease-in-out infinite;
}
</style>