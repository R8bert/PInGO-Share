<template>
  <div>
    <AdminSettings 
      :is-admin="isAdmin"
      @settings-updated="handleSettingsUpdated"
    />
  </div>
</template>

<script setup lang="ts">
import { useTheme } from '../../../composables/useTheme'
import AdminSettings from '../../../components/AdminSettings.vue'

const { isDark } = useTheme()

// Props
interface Props {
  isAdmin: boolean
}

defineProps<Props>()

// Emits
const emit = defineEmits<{
  settingsUpdated: []
}>()

// Event handlers
const handleSettingsUpdated = () => {
  emit('settingsUpdated')
}
</script>
