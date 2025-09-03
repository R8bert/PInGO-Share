<template>
  <div class="upload-container">
    <input type="file" @change="handleUpload" />
    <p v-if="link">Download link: <a :href="link" target="_blank">{{ link }}</a></p>
  </div>
</template>

<script lang="ts">
import { ref } from 'vue'
import axios from 'axios'
import { getApiUrl, getBaseUrl } from '../../../utils/apiUtils'

export default {
  setup() {
    const link = ref("")

    const handleUpload = async (event: Event) => {
      const target = event.target as HTMLInputElement
      if (!target.files || target.files.length === 0) return

      const formData = new FormData()
      formData.append('file', target.files[0])

      try {
        const response = await axios.post(getApiUrl('/upload'), formData)
        link.value = getBaseUrl() + response.data.download_url
      } catch (error) {
        console.error('Upload failed:', error)
      }
    }

    return { handleUpload, link }
  }
}
</script>

<style scoped>
.upload-container {
  padding: 20px;
  border: 2px dashed #ccc;
  text-align: center;
}
</style>
