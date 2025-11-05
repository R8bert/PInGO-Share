<template>
    <div class="min-h-screen flex relative overflow-hidden" :class="isDark ? 'bg-neutral-950 text-white' : 'bg-neutral-50 text-neutral-900'">

        <!-- Left Side - Form Container -->
        <div class="w-full lg:flex-1 flex items-center justify-center p-8 lg:p-12">
            
            <!-- Form Card -->
            <div class="w-full max-w-xl p-8 rounded-3xl shadow-2xl backdrop-blur-sm border border-white/5 transition-all duration-500 hover:shadow-3xl" 
                :class="isDark ? 'bg-neutral-900/95' : 'bg-white/95'">
                
                <!-- Compact Logo/Brand -->
                <div class="mb-8 animate-fade-in">
                    <h2 class="text-lg font-semibold tracking-tight transition-colors duration-300" 
                        :class="isDark ? 'text-white' : 'text-neutral-900'">
                        File sharing 
                    </h2>
                    <p class="text-xs mt-1 transition-colors duration-300" :class="isDark ? 'text-neutral-500' : 'text-neutral-500'">
                        by rootdrop
                    </p>
                </div>

                <!-- Request Files Header -->
                <div class="mb-6 text-center animate-fade-in" style="animation-delay: 0.1s">
                    <h3 class="text-sm font-medium mb-1 transition-colors duration-300" :class="isDark ? 'text-white' : 'text-neutral-900'">
                        Request files
                    </h3>
                </div>

                <!-- Upload Section -->
                <div v-if="!uploadComplete" class="space-y-4">
                    <!-- Add Files Buttons -->
                    <div class="space-y-3">
                            <button @click="triggerFileInput"
                                class="group w-full px-4 py-3 rounded-lg font-medium transition-all duration-300 text-white bg-blue-600 hover:bg-blue-700 text-sm flex items-center justify-center gap-2 hover:shadow-lg hover:shadow-blue-600/30 hover:scale-[1.02] active:scale-[0.98]">
                                <svg class="w-4 h-4 transition-transform duration-300 group-hover:rotate-90" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
                                </svg>
                                Add files
                            </button>

                            <button @click="triggerFileInput"
                                class="group w-full px-4 py-3 rounded-lg font-medium transition-all duration-300 text-white border border-neutral-700 hover:bg-neutral-800 hover:border-neutral-600 text-sm flex items-center justify-center gap-2 hover:shadow-md hover:scale-[1.01] active:scale-[0.99]">
                                <svg class="w-4 h-4 transition-transform duration-300 group-hover:scale-110" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z"></path>
                                </svg>
                                Add folder
                            </button>
                        </div>

                        <!-- Hidden File Input -->
                        <input ref="fileInput" type="file" multiple @change="onFileChange" class="hidden" />

                        <!-- Drag and Drop Zone - more compact -->
                        <div v-if="selectedFiles.length === 0"
                            @drop.prevent="onDrop"
                            @dragover.prevent="isDragging = true"
                            @dragleave.prevent="isDragging = false"
                            class="relative mt-4 p-8 rounded-lg border border-dashed cursor-pointer transition-all duration-300 group"
                            :class="[
                                isDragging ? 'border-blue-500 bg-blue-500/10 scale-[1.02]' : 'border-neutral-700 hover:border-neutral-600 hover:bg-neutral-800/30',
                            ]"
                            @click="triggerFileInput">
                            <div class="text-center">
                                <p class="text-sm text-neutral-400 transition-all duration-300 group-hover:text-neutral-300">
                                    or drop files here
                                </p>
                            </div>
                        </div>

                        <!-- File List with enhanced styling -->
                        <div v-if="selectedFiles.length > 0" class="mt-6 space-y-2 animate-slide-up" style="animation-delay: 0.2s">
                            <!-- Total size indicator -->
                            <div class="flex items-center justify-between mb-3 p-3 rounded-lg backdrop-blur-sm transition-all duration-300 hover:scale-[1.01]" 
                                :class="isDark ? 'bg-neutral-800/30' : 'bg-neutral-100'">
                                <span class="text-sm font-medium">{{ selectedFiles.length }} file{{ selectedFiles.length > 1 ? 's' : '' }}</span>
                                <span class="text-sm transition-colors duration-300" :class="isDark ? 'text-neutral-400' : 'text-neutral-600'">
                                    {{ formatFileSize(selectedFiles.reduce((sum, f) => sum + f.size, 0)) }} total
                                </span>
                            </div>

                            <div v-for="(file, index) in selectedFiles" :key="index"
                                class="group flex items-center gap-3 p-3 rounded-xl transition-all duration-300 animate-scale-in"
                                :style="{ animationDelay: `${0.3 + index * 0.05}s` }"
                                :class="isDark ? 'bg-neutral-800/50 hover:bg-neutral-800 hover:shadow-lg' : 'bg-neutral-100 hover:bg-neutral-200 hover:shadow-md'">
                                
                                <div class="relative">
                                    <div class="w-12 h-12 rounded-xl flex items-center justify-center text-xs font-bold uppercase flex-shrink-0 transition-transform group-hover:scale-110"
                                        :class="isDark ? 'bg-gradient-to-br from-neutral-700 to-neutral-600' : 'bg-gradient-to-br from-neutral-300 to-neutral-400'">
                                        {{ getFileExtension(file).slice(0, 3) }}
                                    </div>
                                    <!-- File type icon overlay -->
                                    <div class="absolute -bottom-1 -right-1 w-5 h-5 rounded-full flex items-center justify-center shadow-lg"
                                        :class="getFileTypeColor(file)">
                                        <svg class="w-3 h-3 text-white" fill="currentColor" viewBox="0 0 20 20">
                                            <path d="M4 3a2 2 0 100 4h12a2 2 0 100-4H4z"></path>
                                            <path fill-rule="evenodd" d="M3 8h14v7a2 2 0 01-2 2H5a2 2 0 01-2-2V8zm5 3a1 1 0 011-1h2a1 1 0 110 2H9a1 1 0 01-1-1z" clip-rule="evenodd"></path>
                                        </svg>
                                    </div>
                                </div>
                                
                                <div class="flex-1 min-w-0">
                                    <p class="text-sm font-semibold truncate group-hover:text-blue-600 transition-colors">{{ file.name }}</p>
                                    <div class="flex items-center gap-2 mt-1">
                                        <p class="text-xs" :class="isDark ? 'text-neutral-500' : 'text-neutral-600'">
                                            {{ formatFileSize(file.size) }}
                                        </p>
                                        <span class="text-xs px-2 py-0.5 rounded-full font-medium"
                                            :class="isDark ? 'bg-neutral-700 text-neutral-300' : 'bg-neutral-200 text-neutral-700'">
                                            {{ getFileExtension(file).toUpperCase() }}
                                        </span>
                                    </div>
                                </div>
                                
                                <button @click.stop="removeFile(index)" 
                                    class="w-9 h-9 rounded-lg flex items-center justify-center transition-all opacity-0 group-hover:opacity-100 hover:scale-110"
                                    :class="isDark ? 'hover:bg-red-500/20 text-neutral-400 hover:text-red-400' : 'hover:bg-red-500/10 text-neutral-600 hover:text-red-600'">
                                    <IconXMark class="w-5 h-5" />
                                </button>
                            </div>
                        </div>

                        <!-- Options - Compact Style -->
                        <div v-if="selectedFiles.length > 0" class="mt-6 space-y-3">
                            <!-- Email input -->
                            <div class="group">
                                <label class="text-xs text-neutral-500 mb-1.5 block transition-colors duration-300 group-hover:text-neutral-400">Email to</label>
                                <input v-model="recipientEmail" type="email" placeholder="friend@email.com"
                                    class="w-full px-3 py-2 rounded-lg border bg-transparent outline-none focus:ring-2 focus:ring-blue-500/50 focus:border-blue-500 transition-all duration-300 text-sm text-white placeholder-neutral-600"
                                    :class="'border-neutral-700 hover:border-neutral-600'" />
                            </div>

                            <!-- Title input -->
                            <div class="group">
                                <label class="text-xs text-neutral-500 mb-1.5 block transition-colors duration-300 group-hover:text-neutral-400">Title</label>
                                <input type="text" placeholder="Untitled"
                                    class="w-full px-3 py-2 rounded-lg border bg-transparent outline-none focus:ring-2 focus:ring-blue-500/50 focus:border-blue-500 transition-all duration-300 text-sm text-white placeholder-neutral-600"
                                    :class="'border-neutral-700 hover:border-neutral-600'" />
                            </div>

                            <!-- Message textarea -->
                            <div class="group">
                                <label class="text-xs text-neutral-500 mb-1.5 block transition-colors duration-300 group-hover:text-neutral-400">Message</label>
                                <textarea v-model="transferMessage" placeholder="Add a note..." rows="2"
                                    class="w-full px-3 py-2 rounded-lg border bg-transparent outline-none focus:ring-2 focus:ring-blue-500/50 focus:border-blue-500 transition-all duration-300 resize-none text-sm text-white placeholder-neutral-600"
                                    :class="'border-neutral-700 hover:border-neutral-600'"></textarea>
                            </div>

                            <!-- Expiry select -->
                            <div class="flex items-center gap-2 group">
                                <svg class="w-4 h-4 text-neutral-500 transition-all duration-300 group-hover:text-neutral-400 group-hover:rotate-180" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                                </svg>
                                <select v-model="selectedValidity" 
                                    class="flex-1 px-3 py-2 rounded-lg border bg-transparent cursor-pointer text-sm text-white outline-none focus:ring-2 focus:ring-blue-500/50 transition-all duration-300 hover:border-neutral-600"
                                    :class="'border-neutral-700'">
                                    <option v-for="option in validityOptions" :key="option.value" :value="option.value"
                                        class="bg-neutral-900">
                                        {{ option.label }}
                                    </option>
                                </select>
                            </div>

                            <!-- More options link -->
                            <button @click="showAdvanced = !showAdvanced" 
                                class="text-xs text-blue-500 hover:text-blue-400 transition-all duration-300 flex items-center gap-1 hover:gap-2">
                                <svg class="w-3 h-3 transition-transform duration-300" :class="{ 'rotate-45': showAdvanced }" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
                                </svg>
                                More options
                            </button>

                            <!-- Transfer Button -->
                            <button @click="uploadFiles" :disabled="isUploading"
                                class="group w-full py-3 rounded-lg font-medium transition-all duration-300 text-white mt-4 hover:shadow-lg hover:scale-[1.02] active:scale-[0.98]"
                                :class="isUploading ? 'bg-blue-500 cursor-not-allowed' : 'bg-blue-600 hover:bg-blue-700 hover:shadow-blue-600/30'">
                                <span class="flex items-center justify-center gap-2">
                                    <svg v-if="isUploading" class="w-4 h-4 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
                                    </svg>
                                    <svg v-else class="w-4 h-4 transition-transform duration-300 group-hover:-translate-y-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4-4m0 0L8 8m4-4v12"></path>
                                    </svg>
                                    {{ isUploading ? 'Transferring...' : 'Transfer' }}
                                </span>
                            </button>

                            <!-- Progress with enhanced styling -->
                            <transition
                                enter-active-class="transition-all duration-300"
                                enter-from-class="opacity-0 translate-y-4"
                                enter-to-class="opacity-100 translate-y-0">
                                <div v-if="isUploading" class="space-y-3 p-4 rounded-xl" 
                                    :class="isDark ? 'bg-neutral-800/50' : 'bg-neutral-100'">
                                    <div class="flex justify-between text-sm font-medium">
                                        <span class="flex items-center gap-2">
                                            <span class="w-2 h-2 rounded-full bg-blue-500 animate-pulse"></span>
                                            Uploading
                                        </span>
                                        <span class="text-blue-600">{{ progress }}%</span>
                                    </div>
                                    <div class="relative h-2 rounded-full overflow-hidden" :class="isDark ? 'bg-neutral-700' : 'bg-neutral-200'">
                                        <div class="absolute inset-0 bg-gradient-to-r from-blue-500 to-blue-600 transition-all duration-300 rounded-full" 
                                            :style="{ width: `${progress}%` }">
                                            <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/30 to-transparent animate-shimmer"></div>
                                        </div>
                                    </div>
                                    <div class="flex justify-between text-xs" :class="isDark ? 'text-neutral-400' : 'text-neutral-600'">
                                        <span v-if="uploadSpeed > 0">{{ formatFileSize(uploadSpeed) }}/s</span>
                                        <span v-if="estimatedTimeRemaining > 0">
                                            {{ Math.ceil(estimatedTimeRemaining / 60) }} min remaining
                                        </span>
                                    </div>
                                </div>
                            </transition>
                        </div>
                    </div>

                    <!-- Success State with enhanced styling -->
                    <transition
                        enter-active-class="transition-all duration-500"
                        enter-from-class="opacity-0 scale-95"
                        enter-to-class="opacity-100 scale-100">
                        <div v-if="uploadComplete" class="py-12 text-center space-y-6">
                            <!-- Success Icon with animation -->
                            <div class="inline-flex">
                                <div class="relative">
                                    <div class="w-20 h-20 rounded-full bg-gradient-to-br from-green-500 to-green-600 flex items-center justify-center shadow-xl animate-bounce-once">
                                        <svg class="w-10 h-10 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7"></path>
                                        </svg>
                                    </div>
                                    <div class="absolute inset-0 rounded-full bg-green-500 animate-ping opacity-75"></div>
                                </div>
                            </div>
                            
                            <div>
                                <h3 class="text-2xl font-bold mb-2">Transfer complete!</h3>
                                <p class="text-base" :class="isDark ? 'text-neutral-400' : 'text-neutral-600'">
                                    Your files are ready to share
                                </p>
                            </div>

                            <div class="space-y-3 max-w-md mx-auto">
                                <div class="group p-4 rounded-xl border-2 flex items-center gap-3 transition-all duration-200 hover:shadow-lg"
                                    :class="isDark ? 'bg-neutral-800/50 border-neutral-700 hover:border-neutral-600' : 'bg-neutral-100 border-neutral-300 hover:border-neutral-400'">
                                    <input :value="uploadLink" readonly 
                                        class="flex-1 bg-transparent outline-none text-sm font-mono truncate" />
                                    <button @click="copyLink"
                                        class="px-5 py-2.5 rounded-lg text-sm font-semibold transition-all duration-200"
                                        :class="copied 
                                            ? 'bg-green-500 text-white' 
                                            : isDark ? 'bg-blue-600 text-white hover:bg-blue-700' : 'bg-blue-600 text-white hover:bg-blue-700'">
                                        <span class="flex items-center gap-2">
                                            <svg v-if="copied" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                                            </svg>
                                            <svg v-else class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
                                            </svg>
                                            {{ copied ? 'Copied!' : 'Copy link' }}
                                        </span>
                                    </button>
                                </div>

                                <button @click="uploadNew"
                                    class="w-full py-3.5 rounded-xl font-semibold transition-all duration-200 flex items-center justify-center gap-2"
                                    :class="isDark ? 'bg-neutral-800 hover:bg-neutral-700 text-white' : 'bg-neutral-200 hover:bg-neutral-300 text-neutral-900'">
                                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
                                    </svg>
                                    Transfer new files
                                </button>
                            </div>
                        </div>
                    </transition>
                
            </div>
            <!-- End Form Card -->
        </div>
        <!-- End Left Side -->

        <!-- Right Side - Hero Image -->
        <div class="hidden lg:flex flex-1 relative overflow-hidden">
            
            <!-- Hero Image/Content -->
            <div class="absolute inset-0 bg-gradient-to-br from-blue-900 via-blue-800 to-purple-900 transition-all duration-700">
                <!-- Subtle animated gradient overlay -->
                <div class="absolute inset-0 bg-gradient-to-tr from-purple-900/20 via-transparent to-blue-900/20 animate-pulse-slow"></div>
                
                <!-- You can replace this with an actual image -->
                <div class="absolute inset-0 flex items-end justify-start p-12 animate-fade-in">
                    <div class="max-w-2xl">
                        <h2 class="text-6xl font-serif italic text-white mb-4 transition-all duration-500 hover:scale-105 origin-left">
                            Digital consciousness
                        </h2>
                        <p class="text-lg text-white/90 transition-opacity duration-300 hover:text-white">
                            Inés Sieulle's experimental film explores our relationship<br/>
                            with artificial intelligence
                        </p>
                        <button class="mt-6 px-6 py-2 border border-white/30 rounded-full text-white text-sm hover:bg-white/10 transition-all duration-300 hover:border-white/50 hover:shadow-lg hover:shadow-white/10">
                            Experience our award-winning wallpapers
                        </button>
                    </div>
                </div>
            </div>
        </div>

    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from "vue";
import { useAuth } from "../../composables/useAuth";
import { useTheme } from "../../composables/useTheme";
import { getApiUrl } from "../../utils/apiUtils";
import axios from "axios";
import IconXMark from '~icons/solar/close-circle-bold-duotone'

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
const showAdvanced = ref(false);

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

const getFileTypeColor = (file: File): string => {
    const ext = getFileExtension(file);
    const imageExts = ['jpg', 'jpeg', 'png', 'gif', 'webp', 'svg', 'bmp', 'ico'];
    const videoExts = ['mp4', 'webm', 'ogg', 'avi', 'mov', 'mkv'];
    const audioExts = ['mp3', 'wav', 'ogg', 'm4a', 'flac', 'aac'];
    const docExts = ['pdf', 'doc', 'docx', 'txt', 'md', 'rtf'];
    const codeExts = ['js', 'ts', 'jsx', 'tsx', 'py', 'java', 'cpp', 'c', 'html', 'css', 'json', 'xml'];
    const archiveExts = ['zip', 'rar', '7z', 'tar', 'gz'];
    
    if (imageExts.includes(ext)) return 'bg-purple-500';
    if (videoExts.includes(ext)) return 'bg-red-500';
    if (audioExts.includes(ext)) return 'bg-green-500';
    if (docExts.includes(ext)) return 'bg-blue-500';
    if (codeExts.includes(ext)) return 'bg-yellow-500';
    if (archiveExts.includes(ext)) return 'bg-orange-500';
    return 'bg-gray-500';
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
            "http://localhost:8080/api/upload",
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
            const fullUrl = getApiUrl(`download/${uploadId}`);
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
/* Custom Scrollbar */
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

/* Fade in animation */
@keyframes fade-in {
    from {
        opacity: 0;
        transform: translateY(20px);
    }
    to {
        opacity: 1;
        transform: translateY(0);
    }
}

/* Slide up animation */
@keyframes slide-up {
    from {
        opacity: 0;
        transform: translateY(30px);
    }
    to {
        opacity: 1;
        transform: translateY(0);
    }
}

/* Scale in animation */
@keyframes scale-in {
    from {
        opacity: 0;
        transform: scale(0.92);
    }
    to {
        opacity: 1;
        transform: scale(1);
    }
}

/* Floating animations */
@keyframes float-slow {
    0%, 100% {
        transform: translate(0, 0);
    }
    50% {
        transform: translate(40px, -40px);
    }
}

@keyframes float-delayed {
    0%, 100% {
        transform: translate(0, 0);
    }
    50% {
        transform: translate(-30px, 30px);
    }
}

/* Bounce once animation */
@keyframes bounce-once {
    0%, 100% {
        transform: scale(1);
    }
    50% {
        transform: scale(1.1);
    }
}

/* Shimmer effect */
@keyframes shimmer {
    0% {
        background-position: -200% 0;
    }
    100% {
        background-position: 200% 0;
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

/* Ping animation */
@keyframes ping {
    0% {
        transform: scale(1);
        opacity: 1;
    }
    75%, 100% {
        transform: scale(2);
        opacity: 0;
    }
}

/* Slow pulse animation */
@keyframes pulse-slow {
    0%, 100% {
        opacity: 0.3;
    }
    50% {
        opacity: 0.6;
    }
}

/* Animation classes */
.animate-fade-in {
    animation: fade-in 0.6s ease-out both;
}

.animate-slide-up {
    animation: slide-up 0.8s ease-out both;
}

.animate-scale-in {
    animation: scale-in 0.6s ease-out both;
}

.animate-float-slow {
    animation: float-slow 25s ease-in-out infinite;
}

.animate-float-delayed {
    animation: float-delayed 30s ease-in-out infinite;
}

.animate-bounce-once {
    animation: bounce-once 0.6s ease-out;
}

.animate-shimmer {
    background: linear-gradient(90deg, transparent, rgba(255,255,255,0.3), transparent);
    background-size: 200% 100%;
    animation: shimmer 2s infinite;
}

.animate-spin {
    animation: spin 1s linear infinite;
}

.animate-ping {
    animation: ping 1s cubic-bezier(0, 0, 0.2, 1) infinite;
}

.animate-pulse-slow {
    animation: pulse-slow 4s ease-in-out infinite;
}

/* Rotate utility */
.rotate-180 {
    transform: rotate(180deg);
}
</style>