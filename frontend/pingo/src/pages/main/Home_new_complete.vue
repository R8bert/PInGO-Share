<template>

    <div class="min-h-screen flex items-center justify-center p-6" :class="isDark ? 'bg-black text-white' : 'bg-white text-black'">

        <div class="w-full max-w-2xl space-y-12">            <div class="text-center space-y-3">

                <h1 class="text-4xl font-light tracking-tight">Share files securely</h1>

                <p class="text-sm opacity-50">End-to-end encrypted • Open source • No tracking</p>            </div>



            <div ref="uploadSection" class="space-y-6">                <div v-if="!uploadComplete" @drop.prevent="onDrop" @dragover.prevent="isDragging = true" @dragleave.prevent="isDragging = false" @click="triggerFileInput" class="border-2 border-dashed rounded-lg p-12 text-center cursor-pointer transition-colors" :class="[isDragging ? 'border-opacity-100 bg-opacity-5' : 'border-opacity-20', isDark ? 'border-white hover:border-opacity-40' : 'border-black hover:border-opacity-40']">

                    <CloudArrowUpIcon class="w-12 h-12 mx-auto mb-4 opacity-30" />

                    <p class="text-lg font-light mb-2">{{ isDragging ? 'Drop files here' : 'Click to upload or drag files' }}</p>                    <p class="text-sm opacity-40">Max {{ formatFileSize(maxUploadSize) }} • Expires in {{ selectedValidity.replace(/(\d+)(\w+)/, '$1 $2') }}</p>

                </div>

                <div v-if="selectedFiles.length > 0 && !uploadComplete" class="space-y-2">

                    <div v-for="(file, index) in selectedFiles" :key="index" class="border rounded-lg p-4" :class="isDark ? 'border-white/10 bg-white/5' : 'border-black/10 bg-black/5'">

                        <div class="flex items-start justify-between">                            <div class="flex-1 min-w-0">

                                <p class="font-medium truncate">{{ file.name }}</p>

                                <p class="text-sm opacity-50 mt-1">{{ formatFileSize(file.size) }}</p>                                <button v-if="canPreview(file)" @click.stop="togglePreview(index)" class="text-sm opacity-50 hover:opacity-100 mt-2 transition-opacity">{{ previewingFiles.has(index) ? 'Hide preview' : 'Show preview' }}</button>

                            </div>

                            <button @click.stop="removeFile(index)" class="p-2 rounded-lg hover:bg-black/10 dark:hover:bg-white/10 transition-colors"><XMarkIcon class="w-5 h-5" /></button>                        </div>

                        <div v-if="previewingFiles.has(index)" class="mt-4 pt-4 border-t" :class="isDark ? 'border-white/10' : 'border-black/10'">

                            <img v-if="previewUrls.has(index) && /\.(jpg|jpeg|png|gif|webp|svg)$/i.test(file.name)" :src="previewUrls.get(index)" class="max-w-full h-auto rounded-lg" />                            <video v-else-if="previewUrls.has(index) && /\.(mp4|webm|ogg)$/i.test(file.name)" :src="previewUrls.get(index)" controls class="max-w-full h-auto rounded-lg" />

                            <audio v-else-if="previewUrls.has(index) && /\.(mp3|wav|ogg|m4a)$/i.test(file.name)" :src="previewUrls.get(index)" controls class="w-full" />

                            <pre v-else-if="textPreviews.has(index)" class="text-xs p-4 rounded-lg overflow-x-auto" :class="isDark ? 'bg-white/5' : 'bg-black/5'">{{ textPreviews.get(index) }}</pre>                        </div>

                    </div>

                    <button @click="triggerFileInput" class="w-full border border-dashed rounded-lg p-4 text-sm opacity-50 hover:opacity-100 transition-opacity" :class="isDark ? 'border-white/20' : 'border-black/20'">Add more files</button>                </div>



                <div v-if="selectedFiles.length > 0 && !uploadComplete" class="space-y-4">                    <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">

                        <div class="relative">

                            <input v-model="uploadPassword" :type="showPassword ? 'text' : 'password'" placeholder="Password (optional)" class="w-full px-4 py-3 rounded-lg border bg-transparent outline-none focus:border-opacity-100 transition-colors" :class="isDark ? 'border-white/20 focus:border-white/40' : 'border-black/20 focus:border-black/40'" />                            <button @click="showPassword = !showPassword" class="absolute right-3 top-1/2 -translate-y-1/2 opacity-50 hover:opacity-100"><EyeIcon v-if="!showPassword" class="w-5 h-5" /><EyeSlashIcon v-else class="w-5 h-5" /></button>

                        </div>

                        <select v-model="selectedValidity" class="w-full px-4 py-3 rounded-lg border bg-transparent outline-none focus:border-opacity-100 transition-colors appearance-none cursor-pointer" :class="isDark ? 'border-white/20 focus:border-white/40' : 'border-black/20 focus:border-black/40'"><option v-for="option in validityOptions" :key="option.value" :value="option.value">Expires: {{ option.label }}</option></select>                    </div>

                    <input v-model="recipientEmail" type="email" placeholder="Recipient email (optional)" class="w-full px-4 py-3 rounded-lg border bg-transparent outline-none focus:border-opacity-100 transition-colors" :class="isDark ? 'border-white/20 focus:border-white/40' : 'border-black/20 focus:border-black/40'" />

                    <textarea v-model="transferMessage" placeholder="Add a message (optional)" rows="3" class="w-full px-4 py-3 rounded-lg border bg-transparent outline-none focus:border-opacity-100 transition-colors resize-none" :class="isDark ? 'border-white/20 focus:border-white/40' : 'border-black/20 focus:border-black/40'"></textarea>                    <button @click="uploadFiles" :disabled="isUploading" class="w-full px-6 py-4 rounded-lg font-medium transition-opacity disabled:opacity-50" :class="isDark ? 'bg-white text-black hover:opacity-90' : 'bg-black text-white hover:opacity-90'">{{ isUploading ? 'Uploading...' : 'Upload' }}</button>

                    <div v-if="isUploading" class="space-y-3">

                        <div class="h-1 rounded-full overflow-hidden" :class="isDark ? 'bg-white/10' : 'bg-black/10'"><div class="h-full transition-all duration-300" :class="isDark ? 'bg-white' : 'bg-black'" :style="{ width: progress + '%' }"></div></div>                        <div class="flex justify-between text-sm opacity-50"><span>{{ uploadSpeed }}</span><span>{{ timeRemaining }}</span></div>

                    </div>

                </div>

                <div v-if="uploadComplete" class="text-center space-y-6 py-12">

                    <div class="w-16 h-16 mx-auto rounded-full flex items-center justify-center" :class="isDark ? 'bg-white/10' : 'bg-black/10'"><svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg></div>                    <div><h3 class="text-2xl font-light mb-2">Upload complete</h3><p class="text-sm opacity-50">Share this link with anyone</p></div>

                    <div class="flex gap-2"><input :value="uploadedUrl" readonly class="flex-1 px-4 py-3 rounded-lg border bg-transparent outline-none" :class="isDark ? 'border-white/20' : 'border-black/20'" /><button @click="copyLink" class="px-6 py-3 rounded-lg font-medium" :class="isDark ? 'bg-white text-black hover:opacity-90' : 'bg-black text-white hover:opacity-90'">Copy</button></div>

                    <button @click="uploadNew" class="text-sm opacity-50 hover:opacity-100 transition-opacity">Upload another file</button>                </div>



                <div v-if="message.text" class="p-4 rounded-lg text-center" :class="message.type === 'error' ? 'bg-red-500/10 text-red-500' : 'bg-green-500/10 text-green-500'">{{ message.text }}</div>            </div>



            <div ref="featuresSection" class="grid grid-cols-1 md:grid-cols-3 gap-6 pt-12">                <div class="text-center space-y-2"><div class="text-3xl mb-3">🔒</div><h3 class="font-medium">End-to-end encrypted</h3><p class="text-sm opacity-50">Your files are secure and private</p></div>

                <div class="text-center space-y-2"><div class="text-3xl mb-3">⚡</div><h3 class="font-medium">Lightning fast</h3><p class="text-sm opacity-50">Direct peer-to-peer transfer</p></div>

                <div class="text-center space-y-2"><div class="text-3xl mb-3">🌐</div><h3 class="font-medium">Open source</h3><p class="text-sm opacity-50">Transparent and community-driven</p></div>            </div>

        </div>

        <input ref="fileInput" type="file" multiple @change="onFileChange" class="hidden" />

    </div>

</template>import { ref, onMounted, onUnmounted } from "vue";

import { useAuth } from "../../composables/useAuth";

import { useTheme } from "../../composables/useTheme";import { useUIColors } from "../../composables/useUIColors";

import axios from "axios";

import {    CloudArrowUpIcon,

    XMarkIcon,

    EyeIcon,    EyeSlashIcon,

} from "@heroicons/vue/24/outline";

import WebGLBackground from "../../components/WebGLBackground.vue";import RotatingText from "../../blocks/TextAnimations/RotatingText/RotatingText.vue";

import GlitchText from "../../blocks/TextAnimations/GlitchText/GlitchText.vue";

import ShinyText from "../../blocks/TextAnimations/ShinyText/ShinyText.vue";

const { isAuthenticated } = useAuth();

const { isDark } = useTheme();const {

    primaryGradient,

    buttonGradient,    hoverGradient,

    primaryColor,

    secondaryColor,    accentColor,

} = useUIColors();

// Rotating text logic

const rotatingWords = ref([

    "effortlessly",    "securely",

    "instantly",

    "freely",    "seamlessly",

]);

const currentWordIndex = ref(0);let rotationInterval: number | null = null;



const getWordStyle = (index: number) => {    const colors = [

        "#ef4444", // effortlessly - red

        "#10b981", // securely - emerald        "#8b5cf6", // instantly - purple

        "#3b82f6", // freely - blue

        "#ec4899", // seamlessly - pink    ];



    const isActive = index === currentWordIndex.value;

    return {

        color: colors[index] || "#ef4444",        textShadow: isActive

            ? `0 0 20px ${colors[index]}40, 0 0 40px ${colors[index]}20`

            : "none",        filter: isActive ? `drop-shadow(0 0 10px ${colors[index]}60)` : "none",

    };

};

const startRotation = () => {

    rotationInterval = window.setInterval(() => {        currentWordIndex.value =

            (currentWordIndex.value + 1) % rotatingWords.value.length;

    }, 2500);};



const stopRotation = () => {    if (rotationInterval) {

        clearInterval(rotationInterval);

        rotationInterval = null;    }

};

// Refs

const fileInput = ref<HTMLInputElement | null>(null);

const uploadSection = ref<HTMLElement | null>(null);const featuresSection = ref<HTMLElement | null>(null);

const { getSettings } = useAuth();

// Stateconst selectedFiles = ref<File[]>([]);

const isDragging = ref(false);

const isUploading = ref(false);const progress = ref(0);

const uploadComplete = ref(false);

// New upload options

const uploadPassword = ref('');

const transferMessage = ref('');const recipientEmail = ref('');

const showPassword = ref(false);

// Upload statistics

const uploadSpeed = ref('0 MB/s');

const timeRemaining = ref('Calculating...');const startTime = ref(0);

const lastLoaded = ref(0);

const lastTime = ref(0);const uploadedUrl = ref("");

const message = ref({ text: "", type: "success" as "success" | "error" });

const previewingFiles = ref<Set<number>>(new Set());const previewUrls = ref<Map<number, string>>(new Map());

const textPreviews = ref<Map<number, string>>(new Map());

const maxUploadSize = ref<number>(104857600); // Default 100 MB

const validityOptions = ref([

    { value: "7days", label: "7 Days", description: "One week" },    { value: "1month", label: "1 Month", description: "30 days" },

    { value: "6months", label: "6 Months", description: "Half year" },

    { value: "1year", label: "1 Year", description: "12 months" },    { value: "never", label: "Never", description: "Permanent" },

]);

const selectedValidity = ref("7days");const maxAllowedValidity = ref("7days");



// Methodsconst triggerFileInput = () => {

    fileInput.value?.click();

};

const onFileChange = (event: Event) => {

    const target = event.target as HTMLInputElement;    if (target.files) {

        selectedFiles.value = Array.from(target.files);

    }};



const onDrop = (event: DragEvent) => {    isDragging.value = false;

    if (event.dataTransfer?.files) {

        selectedFiles.value = Array.from(event.dataTransfer.files);    }

};

const removeFile = (index: number) => {

    // Clean up preview URLs and refs when removing file

    if (previewingFiles.value.has(index)) {        const url = previewUrls.value.get(index);

        if (url) URL.revokeObjectURL(url);

        previewUrls.value.delete(index);        textPreviews.value.delete(index);

        previewingFiles.value.delete(index);

    }

    selectedFiles.value.splice(index, 1);

    // Re-index remaining files

    const newPreviewingFiles = new Set<number>();

    const newPreviewUrls = new Map<number, string>();    const newTextPreviews = new Map<number, string>();



    previewingFiles.value.forEach((previewIndex) => {        if (previewIndex > index) {

            newPreviewingFiles.add(previewIndex - 1);

            const url = previewUrls.value.get(previewIndex);            if (url) newPreviewUrls.set(previewIndex - 1, url);

            const textPreview = textPreviews.value.get(previewIndex);

            if (textPreview) newTextPreviews.set(previewIndex - 1, textPreview);        } else if (previewIndex < index) {

            newPreviewingFiles.add(previewIndex);

            const url = previewUrls.value.get(previewIndex);            if (url) newPreviewUrls.set(previewIndex, url);

            const textPreview = textPreviews.value.get(previewIndex);

            if (textPreview) newTextPreviews.set(previewIndex, textPreview);        }

    });

    previewingFiles.value = newPreviewingFiles;

    previewUrls.value = newPreviewUrls;

    textPreviews.value = newTextPreviews;};



const togglePreview = (index: number) => {    if (previewingFiles.value.has(index)) {

        // Hide preview

        previewingFiles.value.delete(index);        const url = previewUrls.value.get(index);

        if (url) {

            URL.revokeObjectURL(url);            previewUrls.value.delete(index);

        }

        textPreviews.value.delete(index);    } else {

        // Show preview

        previewingFiles.value.add(index);        const file = selectedFiles.value[index];

        const fileExtension = getFileExtension(file);

        if (["txt", "md", "json", "csv", "xml"].includes(fileExtension)) {

            // Handle text file preview

            const reader = new FileReader();            reader.onload = (e) => {

                const content = e.target?.result as string;

                const truncatedContent =                    content.length > 1000

                        ? content.substring(0, 1000) + "\n\n... (truncated)"

                        : content;                textPreviews.value.set(index, truncatedContent);

            };

            reader.readAsText(file);        } else {

            // Handle other file types (images, videos, audio)

            createPreviewUrl(file, index);        }

    }

};

const createPreviewUrl = (file: File, index: number) => {

    const url = URL.createObjectURL(file);    previewUrls.value.set(index, url);

};

const canPreview = (file: File): boolean => {

    const ext = getFileExtension(file);

    return [        "mp4",

        "pdf",

        "jpg",        "jpeg",

        "png",

        "gif",        "webp",

        "txt",

        "md",        "json",

        "csv",

        "xml",        "torrent",

        "mp3",

        "wav",        "flac",

    ].includes(ext);

};

const formatFileSize = (bytes: number): string => {

    if (bytes === 0) return "0 Bytes";    const k = 1024;

    const sizes = ["Bytes", "KB", "MB", "GB"];

    const i = Math.floor(Math.log(bytes) / Math.log(k));    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + " " + sizes[i];

};

const getFileExtension = (file: File): string => {

    return file.name.split(".").pop()?.toLowerCase() || "";

};

const uploadFiles = async () => {

    if (selectedFiles.value.length === 0 || isUploading.value) return;

    isUploading.value = true;

    progress.value = 0;    startTime.value = Date.now();

    lastLoaded.value = 0;

    lastTime.value = Date.now();    uploadSpeed.value = '0 MB/s';

    timeRemaining.value = 'Calculating...';

    try {

        const formData = new FormData();

        selectedFiles.value.forEach((file) => {            formData.append("files", file);

        });

        // Add upload options

        if (uploadPassword.value) {

            formData.append("password", uploadPassword.value);        }

        if (transferMessage.value) {

            formData.append("message", transferMessage.value);        }

        if (recipientEmail.value) {

            formData.append("recipient_email", recipientEmail.value);        }

        if (selectedValidity.value) {

            formData.append("validity_hours", selectedValidity.value.toString());        }



        const response = await axios.post(            "http://localhost:8080/upload",

            formData,

            {                headers: {

                    "Content-Type": "multipart/form-data",

                    ...(isAuthenticated.value &&                    localStorage.getItem("auth_token")

                        ? {

                              Authorization: `Bearer ${localStorage.getItem(                                  "auth_token"

                              )}`,

                          }                        : {}),

                },

                withCredentials: true,                onUploadProgress: (progressEvent) => {

                    if (progressEvent.total) {

                        progress.value = Math.round(                            (progressEvent.loaded * 100) / progressEvent.total

                        );

                        // Calculate upload speed and time remaining

                        const currentTime = Date.now();

                        const timeElapsed = (currentTime - lastTime.value) / 1000; // seconds                        

                        if (timeElapsed > 0.5) { // Update every 0.5 seconds

                            const bytesUploaded = progressEvent.loaded - lastLoaded.value;                            const speed = bytesUploaded / timeElapsed; // bytes per second

                            const speedMB = (speed / (1024 * 1024)).toFixed(2);

                            uploadSpeed.value = `${speedMB} MB/s`;

                            const remainingBytes = progressEvent.total - progressEvent.loaded;

                            const remainingSeconds = remainingBytes / speed;                            

                            if (remainingSeconds < 60) {

                                timeRemaining.value = `${Math.ceil(remainingSeconds)}s remaining`;                            } else if (remainingSeconds < 3600) {

                                timeRemaining.value = `${Math.ceil(remainingSeconds / 60)}m remaining`;

                            } else {                                timeRemaining.value = `${Math.ceil(remainingSeconds / 3600)}h remaining`;

                            }

                            lastLoaded.value = progressEvent.loaded;

                            lastTime.value = currentTime;

                        }                    }

                },

            }        );



        if (response.data.download_url) {            progress.value = 100;

            uploadComplete.value = true;

            // Extract upload ID from download_url and create full URL

            const uploadId = response.data.download_url.split("/").pop();

            uploadedUrl.value = `${window.location.origin}/download/${uploadId}`;

            message.value = {

                text: "Files uploaded successfully!",                type: "success",

            };

            // Scroll to center the upload success section

            setTimeout(() => {

                uploadSection.value?.scrollIntoView({                    behavior: "smooth",

                    block: "center",

                });            }, 100);



            // Clean up preview URLs            previewUrls.value.forEach((url) => URL.revokeObjectURL(url));

            previewUrls.value.clear();

            textPreviews.value.clear();            previewingFiles.value.clear();

        }

    } catch (error) {        console.error("Upload error:", error);

        message.value = {

            text: "Upload failed. Please try again.",            type: "error",

        };

        progress.value = 0;    } finally {

        isUploading.value = false;

        // Clear message after 5 seconds instead of 3

        setTimeout(() => {

            message.value = { text: "", type: "success" };            if (!isUploading.value) progress.value = 0;

        }, 5000);

    }};



// Action buttons for upload successconst uploadNew = () => {

    uploadComplete.value = false;

    uploadedUrl.value = "";    selectedFiles.value = [];

    progress.value = 0;

    message.value = { text: "", type: "success" };

    // Clean up any remaining preview URLs

    previewUrls.value.forEach((url) => URL.revokeObjectURL(url));    previewUrls.value.clear();

    textPreviews.value.clear();

    previewingFiles.value.clear();};



const copyLink = async () => {    try {

        await navigator.clipboard.writeText(uploadedUrl.value);

        message.value = { text: "Link copied to clipboard!", type: "success" };        setTimeout(() => {

            message.value = { text: "", type: "success" };

        }, 3000);    } catch (err) {

        message.value = { text: "Failed to copy link", type: "error" };

        setTimeout(() => {            message.value = { text: "", type: "success" };

        }, 3000);

    }};



const openLink = () => {    window.open(uploadedUrl.value, "_blank");

};

const scrollToUpload = () => {

    uploadSection.value?.scrollIntoView({ behavior: "smooth" });

};

const scrollToFeatures = () => {

    featuresSection.value?.scrollIntoView({ behavior: "smooth" });};



const loadSettings = async () => {    try {

        const settings = await getSettings();

        maxUploadSize.value = settings.maxUploadSize || 104857600;        maxAllowedValidity.value = settings.maxValidity || "7days";



        // Filter validity options based on max allowed        const validityOrder = ["7days", "1month", "6months", "1year", "never"];

        const maxIndex = validityOrder.indexOf(maxAllowedValidity.value);

        if (maxIndex !== -1) {

            const allowedOptions = validityOrder.slice(0, maxIndex + 1);

            validityOptions.value = validityOptions.value.filter((option) =>                allowedOptions.includes(option.value)

            );

        }

        // Set default validity to first available option

        if (validityOptions.value.length > 0) {            selectedValidity.value = validityOptions.value[0].value;

        }

    } catch (error) {        console.error("Error loading settings:", error);

    }

};

onMounted(() => {

    // Scroll to top smoothly when page loads    loadSettings();

    window.scrollTo({ top: 0, behavior: "smooth" });

    // Start the rotating text animation

    startRotation();

});

onUnmounted(() => {

    // Clean up the rotation interval    stopRotation();

});

</script>

<style scoped>

/* Smooth scrolling for the entire page */html {

    scroll-behavior: smooth;

}

@keyframes fade-in {

    from {        opacity: 0;

        transform: translateY(20px);

    }    to {

        opacity: 1;

        transform: translateY(0);    }

}

@keyframes fade-in-delay {

    from {

        opacity: 0;        transform: translateY(20px);

    }

    to {        opacity: 1;

        transform: translateY(0);

    }}



@keyframes fade-in-delay-2 {    from {

        opacity: 0;

        transform: translateY(20px);    }

    to {

        opacity: 1;        transform: translateY(0);

    }

}



.animate-gradient {  background-size: 200% 200%;

  animation: gradient-move 3s linear infinite;

}



.animate-fade-in {    animation: fade-in 0.8s ease-out;

}

.animate-fade-in-delay {

    animation: fade-in-delay 0.8s ease-out 0.2s both;

}

.animate-fade-in-delay-2 {

    animation: fade-in-delay-2 0.8s ease-out 0.4s both;}



@keyframes liquid-glow {

    0%,

    100% {        filter: drop-shadow(0 0 10px #ef444460) drop-shadow(0 0 20px #ef444440);

    }

    50% {        filter: drop-shadow(0 0 15px #ef444480) drop-shadow(0 0 30px #ef444460);

    }

}



@keyframes gradient-move {  0% { background-position: 0% 50%; }

  100% { background-position: 100% 50%; }

}



</style><style scoped>
html { scroll-behavior: smooth; }
</style>



