<template>
  <div class="font-color-overwrite" style="-ms-overflow-style: auto">
    <pre>{{ mediaInfo }}</pre>
  </div>
</template>

<script>
import Wails from "@wailsapp/runtime";

export default {
  name: "Info",
  data() {
    return {
      mediaInfo: ''
    }
  },
  mounted() {
    Wails.Events.On('SetMediaInfo', (mediaInfo) => {
      const json = JSON.parse(mediaInfo)
      this.mediaInfo = JSON.stringify(json, null, "\t")
    })
  }
}
</script>

<style scoped>

</style>