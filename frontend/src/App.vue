<template>
  <div id="app" class="container" :style="`transform: scale(${scale});opacity: ${opacity}`">
    <Main/>
  </div>
</template>

<script>
import Main from "@/components/Main";
// import Setting from "@/components/Setting";
// import Test from "@/components/Test";
import "./assets/css/main.css";

export default {
  name: "app",
  components: {
    Main,
    // Setting,
    // Test
  },
  data() {
    return {
      scale: 1,
      opacity: 0
    };
  },
  mounted() {
    this.adaptScale()
    this.opacity = 1
  },
  methods: {
    adaptScale () {
      let h = document.body.clientWidth; //获得屏幕宽度
      this.scale = h/600;
    }
  }
};
</script>

<style>
.container {
  height: 360px;
  width: 600px;
  position: fixed;
  /*background-color: black;*/
  background-color: #fefefe;
  transition: opacity 0.3s ease;
  transform-origin: left top;
}
</style>