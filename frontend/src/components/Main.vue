<template>
  <div class="main-container">
    <div class="panel" style="margin-top: 1vw">
      <a-input-search v-model="srcPath" addon-before="文件路径" class="input" placeholder="" size="large" @search="onSearch">
        <a-button slot="enterButton" @click="getSrcPath" class="search-button">
          <a-icon type="folder" style="margin: auto;font-size: 2.25vw" />
        </a-button>
      </a-input-search>
    </div>
    <div class="panel">
      <a-input-search v-model="dstPath" addon-before="输出路径" class="input" placeholder="" size="large" @search="onSearch">
        <a-button slot="enterButton" @click="getDistPath" class="search-button">
          <a-icon type="folder" style="margin: auto;font-size: 2.25vw"/>
        </a-button>
      </a-input-search>
    </div>
    <div class="panel">
      <!--  TODO: 压制代码块样式    -->
      <a-textarea  placeholder="FFmpeg压制的代码"
                   size="large"
                   style="font-size: 2.25vw;padding: 0.75vw 1vw 0.75vw"
                   :auto-size="{ minRows: 4, maxRows: 4}"
      />
<!--      border-radius: 1vw;-->
    </div>
<!--    <div class="panel">-->
<!--      <div class="dragArea">-->
<!--        拖拽-->
<!--      </div>-->
<!--    </div>-->
    <div class="panel" style="">
      <a-progress
          size="large"
          stroke-linecap="round"
          :percent="percent"
          class="progressbar"
          strokeWidth="10"
          :stroke-color="{
                       from: '#108ee9',
                       to: '#87d068',
                     }"
      />
<!--      <a-space>-->
<!--      <span  style="right: 0">-->
<!--        <a-button size="large" class="btn" style="width: 10vw">-->
<!--          结束-->
<!--        </a-button>-->
<!--      </span>-->
      <a-button size="large" class="btn" style="">
        开始
      </a-button>
        <!--          <a-icon type="caret-right" />-->
        <!--          <a-icon type="pause" />-->
<!--      </a-space>-->
    </div>
  </div>
</template>

<script>
import Wails from '@wailsapp/runtime';

export default {
  name: "Main",
  data() {
    return {
      percent: 66.57,
    };
  },
  mounted() {
    Wails.Events.On("SetProcessBar", (percent) => {
      this.percent = percent;
    });
  },
  methods: {

  }
}
</script>

<style scoped>

/deep/ .ant-input-lg {
  height: 6vw;
  font-size: 2.25vw;
}

/deep/ .ant-input-group-addon {
  font-size: 2vw;
  height: 6vw;
  padding-left: 2vw;
  padding-right: 2vw;
}

.search-button{
  height: 6vw;
}

.main-container{
  /*padding: 50x;*/
  /*margin: 0px 8px 8px 8px;*/
}

.btn {
  float: right;
  width: 15vw;
  height: 5vw;
  font-size: 2vw;
  /*box-shadow: 0 0.25vw 2vw rgba(0,0,0,0.2);*/
}

.dragArea {
  height: 15vw;
  margin: auto;
  border: 1px solid rgba(0,0,0,0.4);
  border-radius: 0.5vw;
  padding: 1.25vw;
  font-size: 2.25vw;
  text-align: center;
  vertical-align: middle;
}

.progressbar{
  font-size: 2vw;
  margin-right: 5vw;
  margin-top: 1vw;
  /*width: 70vw;*/
  max-width: 75vw;
}

.panel{
  /*margin: 0.5vw 0.25vw 0;*/
  /*height: 10vw;*/
}

input {
  height: 12vw;font-size: 6vw;
  border-radius: 0.5vw;
}

</style>