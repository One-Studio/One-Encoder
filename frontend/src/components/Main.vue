<template>
  <div >
    <div class="q-gutter-y-md" style="max-width: 600px">
      <q-card>
        <q-tabs
            v-model="tab"
            dense
            class="bg-grey-1"
            active-color="primary"
            indicator-color="primary"
            align="left"
            narrow-indicator
        >
          <q-tab name="main" label="主要" />
          <q-tab name="setting" label="设置" />
          <q-tab name="about" label="关于" />
        </q-tabs>

        <q-separator />

        <q-tab-panels v-model="tab" animated>
          <q-tab-panel name="main">
            <div>Mails</div>
            Lorem ipsum dolor sit amet consectetur adipisicing elit.
          </q-tab-panel>

          <q-tab-panel name="setting">
            <div class="text-h6">Alarms</div>
            Lorem ipsum dolor sit amet consectetur adipisicing elit.
          </q-tab-panel>

          <q-tab-panel name="about">
            <div class="text-h6">Movies</div>
            Lorem ipsum dolor sit amet consectetur adipisicing elit.
          </q-tab-panel>
        </q-tab-panels>
      </q-card>
    </div>
  </div>
</template>

<script>
import Wails from '@wailsapp/runtime';

export default {
  name: "Main",
  data() {
    return {
      tab: 'main',
      opacity: 0,
      srcPath: 'C:/Users/Purp1e/Videos/测试.mp4',
      dstPath: 'C:/Users/Purp1e/Desktop/测试One-Encoder.mp4',
      param: {
        ffmpeg: '-vcodec libx264 -crf 20 -preset slow',
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
    this.opacity = 1;
  },
  methods: {
    setVar () {
      window.backend.App.SetSrcPath(this.srcPath)
      window.backend.App.SetDstPath(this.dstPath)
      window.backend.App.SetParam(this.select, this.param[this.select])
    },
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
      this.setVar()
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
      this.status = true
      window.backend.App.StartEncoding(this.select).then(() => {
        this.status = false
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
.main-container{
  width: 100%;
  transition: all 1s ease;
}

</style>