<template>
  <div id="app" class="container">
    <div class="card-container">
<!--      <a-tabs>-->
<!--        <a-tab-pane key="1" tab="主页">-->
          <Main/>
<!--        </a-tab-pane>-->
<!--        <a-tab-pane key="2" tab="设置">-->
<!--          <Setting/>-->
<!--        </a-tab-pane>-->
<!--        <a-tab-pane key="3" tab="测试">-->
<!--          <Test/>-->
<!--        </a-tab-pane>-->
        <!--        <p slot="tabBarExtraContent" style="margin-right: 10px ">One Encoder</p>-->
        <!--        <a-button slot="tabBarExtraContent">One Encoder</a-button>-->
<!--      </a-tabs>-->
    </div>
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
    return {};
  },
};
</script>

<style>
.container {
  height: 360px;
  width: 600px;
  position: fixed;
  /*background-color: black;*/
  background-color: #fefefe;
  /*width: 680px;*/
}

.card-container {
  overflow: hidden;
  padding: 4px 8px 4px;
  /*margin-top: -1vw;*/
}
</style>