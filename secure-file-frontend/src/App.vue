<template>
  <div class="container">
    <h1>Secure File Upload</h1>
    <input type="file" @change="handleFileChange" />
    <button @click="uploadFile">Upload</button>
    <p v-if="message">{{ message }}</p>
    <a v-if="downloadUrl" :href="downloadUrl" download>Download File</a>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const selectedFile = ref(null)
const message = ref('')
const downloadUrl = ref('')

function handleFileChange(event) {
  selectedFile.value = event.target.files[0]
}

async function uploadFile() {
  if (!selectedFile.value) {
    message.value = 'Please select a file first.'
    return
  }

  const formData = new FormData()
  formData.append('file', selectedFile.value)

  try {
    const response = await fetch('http://localhost:3000/upload', {
      method: 'POST',
      body: formData,
    })

    const data = await response.json()
    if (response.ok) {
      message.value = 'File uploaded!'
      downloadUrl.value = `http://localhost:3000/download/${data.filename}`
    } else {
      message.value = data.message || 'Upload failed.'
    }
  } catch (error) {
    message.value = 'Error uploading file.'
  }
}
</script>

<style>
.container {
  max-width: 600px;
  margin: 80px auto;
  padding: 30px;
  border: 1px solid #ccc;
  border-radius: 12px;
  text-align: center;
  font-family: Arial, sans-serif;
}
input, button {
  margin: 10px;
}
</style>
