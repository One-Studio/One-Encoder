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
      <a-radio-group size="large" v-model="select" @change="onChange">
        <a-radio-button value="ffmpeg" style="font-size: 2vw;height: 5vw;line-height: 5vw;">
          ffmpeg
        </a-radio-button>
        <a-radio-button value="x264" style="font-size: 2vw;height: 5vw;line-height: 5vw;">
          x264
        </a-radio-button>
        <a-radio-button value="x265" style="font-size: 2vw;height: 5vw;line-height: 5vw;">
          x265
        </a-radio-button>
      </a-radio-group>
    </div>
    <div class="panel">
      <!--  TODO: 压制代码块样式    -->
      <a-textarea  placeholder="压制代码"
                   size="large"
                   style="font-size: 2.25vw;padding: 0.75vw 1vw 0.75vw"
                   :auto-size="{ minRows: 4, maxRows: 4}"
                   v-model="param[select]"
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
          :percent="progress"
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
      <a-button size="large" class="btn" @click="onStart">
        <span v-if='status===false'>开始</span>
        <span v-else>结束</span>
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
      srcPath: '',
      dstPath: '',
      param: {
        ffmpeg: '',
        x264: '',
        x265: ''
      },
      select: 'ffmpeg',
      progress: 0,
      status: false  //是否正在执行
    };
  },
  mounted() {
    Wails.Events.On("SetProgess", (progress) => {
      this.progress = progress;
    });
    Wails.Events.On("NoticeSuccess", (msg) => {
      this.$message.success(msg, 5);
    });
    Wails.Events.On("NoticeError", (msg) => {
      this.$message.error(msg, 5);
    });
    Wails.Events.On("NoticeWarning", (msg) => {
      this.$message.warning(msg, 5);
    });
    //通知传参
    window.backend.App.SetVar();
    //检查更新
    this.checkUpdate();
  },
  methods: {
    getSrcPath () {
      window.backend.App.GetSrcPath().then(path => {
          this.srcPath = path
      });
    },
    getDistPath () {
      window.backend.App.GetDistPath().then(path => {
        this.dstPath = path
      });
    },
    onStart () {
      //debug
      // this.status = !this.status
      //根据情况决定开始/结束压制
      if (this.status === false) {
        this.StartEncoding()
      } else {
        this.QuitEncoding()
      }
    },
    StartEncoding () {
      window.backend.App.StartEncoding(this.select).then(() => {

      });
    },
    PauseEncoding () {
      window.backend.App.PauseEncoding().then(() => {

      });
    },
    QuitEncoding () {
      window.backend.App.QuitEncoding().then(() => {

      });
    }
  }
}
</script>

<style scoped>

/deep/ .ant-input-lg {
  height: 5vw;
  font-size: 2vw;
}

/deep/ .ant-input-group-addon {
  font-size: 2.25vw;
  height: 5vw;
  padding-left: 2vw;
  padding-right: 2vw;
}

.tool-select {
  font-size: 2.25vw;
  height: 10vw;
}


.search-button{
  height: 5vw;
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