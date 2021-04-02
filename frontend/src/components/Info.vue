<template>
  <div class="font-color-overwrite">
    {{ mediaInfo }}
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
      this.mediaInfo = JSON.stringify(mediaInfo, null, "\t")
    })
  }
}
</script>

<style scoped>

</style>