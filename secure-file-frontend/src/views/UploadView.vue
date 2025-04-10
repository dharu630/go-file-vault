<template>
  <div id="app">
    <h1>Secure File Upload</h1>
    <input type="file" @change="handleFileChange" />
    <button @click="uploadFile">Upload</button>

    <div v-if="downloadLink">
      <p>File uploaded!</p>
      <a :href="downloadLink" target="_blank">Download File</a>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      file: null,
      downloadLink: '',
    }
  },
  methods: {
    handleFileChange(e) {
      this.file = e.target.files[0]
    },
    async uploadFile() {
      if (!this.file) return alert('Please select a file first')

      const formData = new FormData()
      formData.append('file', this.file)

      try {
        const res = await axios.post('http://localhost:3000/upload', formData)
        this.downloadLink = `http://localhost:3000/download/${res.data.filename}`
      } catch (err) {
        alert('Upload failed')
        console.error(err)
      }
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  margin-top: 100px;
}
</style>
