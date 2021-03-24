<template>
  <div id="app" class="container" :style="`transform: scale(${scale});`">
    <div class="opacity-container" :style="`opacity: ${opacity}`">
      <van-tabs
          type="line"
          animated
          line-height="1.5"
          line-width="50"
          background="#fefefe"
          color="rgb(64,158,255)"
      >
        <!--          color="#9370DB"-->
        <!--        title-active-color="pink"-->
        <van-tab title="主页" style="padding: 10px">
          <Main></Main>
        </van-tab>
        <van-tab title="设置" style="padding: 10px">
          <Setting></Setting>
        </van-tab>
        <van-tab title="关于" style="padding: 10px">
          <About></About>
        </van-tab>
      </van-tabs>
    </div>
<!--    <Main/>-->
  </div>
</template>

<script>
import Main from "@/components/Main";
import Setting from "@/components/Setting";
import About from "@/components/About";
// import Test from "@/components/Test";
import "./assets/css/main.css";

export default {
  name: "app",
  components: {
    Main,
    Setting,
    About
  },
  data() {
    return {
      scale: 1,
      opacity: 0
    };
  },
  mounted() {
    this.adaptScale()
    setTimeout(()=>{
      this.opacity = 1
    }, 300)
    window.onresize = () => {
      this.adaptScale()
    }
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
  height: 380px;
  width: 640px;
  position: fixed;
  /*background-color: black;*/
  background-color: #fefefe;
  transition: all 0.6s ease;
  transform-origin: left top;
}

.opacity-container {
  transition: opacity 0.3s ease;
}

.van-tabs__wrap {
  width: 34%;
  height: 36px !important;
}

.van-tab {
  font-size: 16px !important;
}
</style>