<template>
  <div>
    <AdminSettings 
      :is-admin="isAdmin"
      @settings-updated="handleSettingsUpdated"
    />
  </div>
</template>

<script setup lang="ts">
import AdminSettings from '../../../components/AdminSettings.vue'

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
