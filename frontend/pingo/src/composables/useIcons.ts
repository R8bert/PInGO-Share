import { computed } from 'vue'

// Define file type to icon mapping
const FILE_TYPE_ICONS = {
  // Document files
  pdf: '/src/assets/svg/icons/files_pdf.svg',
  doc: '/src/assets/svg/icons/word_docx_doc.svg',
  docx: '/src/assets/svg/icons/word_docx_doc.svg',
  ppt: '/src/assets/svg/icons/ppt_pptx_file.svg',
  pptx: '/src/assets/svg/icons/ppt_pptx_file.svg',
  
  // Image files
  jpg: '/src/assets/svg/icons/files_image.svg',
  jpeg: '/src/assets/svg/icons/files_image.svg',
  png: '/src/assets/svg/icons/files_image.svg',
  gif: '/src/assets/svg/icons/files_image.svg',
  svg: '/src/assets/svg/icons/files_image.svg',
  webp: '/src/assets/svg/icons/files_image.svg',
  
  // Video files
  mp4: '/src/assets/svg/icons/files_video.svg',
  mov: '/src/assets/svg/icons/files_video.svg',
  avi: '/src/assets/svg/icons/files_video.svg',
  mkv: '/src/assets/svg/icons/files_video.svg',
  wmv: '/src/assets/svg/icons/files_video.svg',
  flv: '/src/assets/svg/icons/files_video.svg',
  webm: '/src/assets/svg/icons/files_video.svg',
  
  // Audio files
  mp3: '/src/assets/svg/icons/files_audio.svg',
  wav: '/src/assets/svg/icons/files_audio.svg',
  flac: '/src/assets/svg/icons/files_audio.svg',
  aac: '/src/assets/svg/icons/files_audio.svg',
  ogg: '/src/assets/svg/icons/files_audio.svg',
  
  // Archive files
  zip: '/src/assets/svg/icons/files_archive.svg',
  rar: '/src/assets/svg/icons/files_archive.svg',
  '7z': '/src/assets/svg/icons/files_archive.svg',
  tar: '/src/assets/svg/icons/files_archive.svg',
  gz: '/src/assets/svg/icons/files_archive.svg',
  
  // Text files
  txt: '/src/assets/svg/icons/files_text.svg',
  md: '/src/assets/svg/icons/files_text.svg',
  rtf: '/src/assets/svg/icons/files_text.svg',
  
  // Spreadsheet files
  xls: '/src/assets/svg/icons/files_excel.svg',
  xlsx: '/src/assets/svg/icons/files_excel.svg',
  csv: '/src/assets/svg/icons/files_excel.svg',
  
  // Code files
  js: '/src/assets/svg/icons/files_code.svg',
  ts: '/src/assets/svg/icons/files_code.svg',
  html: '/src/assets/svg/icons/files_code.svg',
  css: '/src/assets/svg/icons/files_code.svg',
  json: '/src/assets/svg/icons/files_code.svg',
  xml: '/src/assets/svg/icons/files_code.svg',
  vue: '/src/assets/svg/icons/files_code.svg',
  py: '/src/assets/svg/icons/files_code.svg',
  java: '/src/assets/svg/icons/files_code.svg',
  cpp: '/src/assets/svg/icons/files_code.svg',
  c: '/src/assets/svg/icons/files_code.svg',
  php: '/src/assets/svg/icons/files_code.svg',
  rb: '/src/assets/svg/icons/files_code.svg',
  go: '/src/assets/svg/icons/files_code.svg',
} as const

// System icons for UI components
const SYSTEM_ICONS = {
  // Upload/file management
  filesUploaded: '/src/assets/svg/icons/files_uploaded.svg',
  filesUploaded2: '/src/assets/svg/icons/files_uploaded2.svg',
  totalFiles: '/src/assets/svg/icons/total_files_icon.svg',
  multipleFiles: '/src/assets/svg/icons/default_multiple_files.svg',
  fileFolder: '/src/assets/svg/icons/file-folder.png',
  
  // File type icons (for missing SVGs, we'll fall back to default)
  filesImage: '/src/assets/svg/icons/files_image.svg',
  filesAudio: '/src/assets/svg/icons/files_audio.svg',
  filesArchive: '/src/assets/svg/icons/files_archive.svg',
  filesText: '/src/assets/svg/icons/files_text.svg',
  filesExcel: '/src/assets/svg/icons/files_excel.svg',
  filesCode: '/src/assets/svg/icons/files_code.svg',
  
  // File operations
  download: '/src/assets/svg/icons/download.svg',
  upload: '/src/assets/svg/icons/upload.svg',
  share: '/src/assets/svg/icons/share.svg',
  copy: '/src/assets/svg/icons/copy.svg',
  delete: '/src/assets/svg/icons/delete.svg',
  
  // Status icons
  success: '/src/assets/svg/icons/success.svg',
  error: '/src/assets/svg/icons/error.svg',
  warning: '/src/assets/svg/icons/warning.svg',
  info: '/src/assets/svg/icons/info.svg',
  
  // Navigation
  home: '/src/assets/svg/icons/home.svg',
  settings: '/src/assets/svg/icons/settings.svg',
  account: '/src/assets/svg/icons/account.svg',
  
  // Default fallback
  default: '/src/assets/svg/icons/files_no_icon.svg',
} as const

export type FileExtension = keyof typeof FILE_TYPE_ICONS
export type SystemIcon = keyof typeof SYSTEM_ICONS

export function useIcons() {
  /**
   * Get file extension from filename
   */
  const getFileExtension = (filename: string): string => {
    const parts = filename.split('.')
    return parts.length > 1 ? parts.pop()?.toLowerCase() || '' : ''
  }

  /**
   * Get icon path for a specific file based on its extension
   */
  const getFileIcon = (filename: string): string => {
    const extension = getFileExtension(filename) as FileExtension
    return FILE_TYPE_ICONS[extension] || SYSTEM_ICONS.default
  }

  /**
   * Get icon alt text for a file
   */
  const getFileIconAlt = (filename: string): string => {
    const extension = getFileExtension(filename)
    
    // Special cases for better alt text
    const altTextMap: Record<string, string> = {
      pdf: 'PDF document',
      doc: 'Word document',
      docx: 'Word document',
      ppt: 'PowerPoint presentation',
      pptx: 'PowerPoint presentation',
      xls: 'Excel spreadsheet',
      xlsx: 'Excel spreadsheet',
      jpg: 'Image file',
      jpeg: 'Image file',
      png: 'Image file',
      gif: 'Animated image',
      mp4: 'Video file',
      mov: 'Video file',
      mp3: 'Audio file',
      wav: 'Audio file',
      zip: 'Archive file',
      txt: 'Text file',
      js: 'JavaScript file',
      ts: 'TypeScript file',
      vue: 'Vue component',
      html: 'HTML file',
      css: 'CSS file',
    }
    
    return altTextMap[extension] || `${extension.toUpperCase()} file` || 'File'
  }

  /**
   * Get icon for upload based on files count and types
   */
  const getUploadIcon = (files: string[]): string => {
    // If multiple files, use multiple files icon
    if (files.length > 1) {
      return SYSTEM_ICONS.multipleFiles
    }
    
    // For single file, use specific file type icon
    return getFileIcon(files[0])
  }

  /**
   * Get icon alt text for upload
   */
  const getUploadIconAlt = (files: string[]): string => {
    if (files.length > 1) {
      return `${files.length} files`
    }
    
    return getFileIconAlt(files[0])
  }

  /**
   * Get system icon path
   */
  const getSystemIcon = (iconName: SystemIcon): string => {
    return SYSTEM_ICONS[iconName]
  }

  /**
   * Get file type category for styling purposes
   */
  const getFileCategory = (filename: string): string => {
    const extension = getFileExtension(filename)
    
    if (['jpg', 'jpeg', 'png', 'gif', 'svg', 'webp'].includes(extension)) {
      return 'image'
    }
    if (['mp4', 'mov', 'avi', 'mkv', 'wmv', 'flv', 'webm'].includes(extension)) {
      return 'video'
    }
    if (['mp3', 'wav', 'flac', 'aac', 'ogg'].includes(extension)) {
      return 'audio'
    }
    if (['pdf', 'doc', 'docx', 'ppt', 'pptx', 'xls', 'xlsx'].includes(extension)) {
      return 'document'
    }
    if (['zip', 'rar', '7z', 'tar', 'gz'].includes(extension)) {
      return 'archive'
    }
    if (['js', 'ts', 'html', 'css', 'json', 'xml', 'vue', 'py', 'java', 'cpp', 'c', 'php', 'rb', 'go'].includes(extension)) {
      return 'code'
    }
    if (['txt', 'md', 'rtf'].includes(extension)) {
      return 'text'
    }
    
    return 'other'
  }

  /**
   * Check if file is an image
   */
  const isImageFile = (filename: string): boolean => {
    return getFileCategory(filename) === 'image'
  }

  /**
   * Check if file is a video
   */
  const isVideoFile = (filename: string): boolean => {
    return getFileCategory(filename) === 'video'
  }

  /**
   * Check if file is an audio file
   */
  const isAudioFile = (filename: string): boolean => {
    return getFileCategory(filename) === 'audio'
  }

  /**
   * Get file size color class based on size
   */
  const getFileSizeColor = (bytes: number, isDark: boolean): string => {
    const mb = bytes / (1024 * 1024)
    
    if (mb < 1) {
      return isDark ? 'text-green-400' : 'text-green-600'
    } else if (mb < 10) {
      return isDark ? 'text-yellow-400' : 'text-yellow-600'
    } else if (mb < 100) {
      return isDark ? 'text-orange-400' : 'text-orange-600'
    } else {
      return isDark ? 'text-red-400' : 'text-red-600'
    }
  }

  return {
    // Core functions
    getFileExtension,
    getFileIcon,
    getFileIconAlt,
    getUploadIcon,
    getUploadIconAlt,
    getSystemIcon,
    
    // Category functions
    getFileCategory,
    isImageFile,
    isVideoFile,
    isAudioFile,
    
    // Utility functions
    getFileSizeColor,
    
    // Constants for external use
    FILE_TYPE_ICONS: computed(() => FILE_TYPE_ICONS),
    SYSTEM_ICONS: computed(() => SYSTEM_ICONS),
  }
}
