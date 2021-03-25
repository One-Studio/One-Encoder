<template>
  <div>
    <el-input class="panel" :span="24" placeholder="请选择文件" v-model="input" @change="generateOutput">
      <template slot="prepend">输入</template>
      <el-button slot="append" icon="el-icon-folder"  @click="getInput"></el-button>
    </el-input>
    <el-input class="panel" placeholder="请选择文件" v-model="output">
      <template slot="prepend">输出</template>
      <el-button slot="append" icon="el-icon-folder" @click="getOutput"></el-button>
    </el-input>

  <div class="panel">
     <span style="float: left">
        <el-radio-group v-model="tool" size="small">
          <el-radio-button label="ffmpeg">ffmpeg</el-radio-button>
          <el-radio-button label="x264">x264</el-radio-button>
          <el-radio-button label="x265">x265</el-radio-button>
        </el-radio-group>
<!--        <a-radio-group size="normal" v-model="select">-->
<!--          <a-radio-button value="ffmpeg" style="">-->
<!--            ffmpeg-->
<!--          </a-radio-button>-->
<!--          <a-radio-button value="x264" style="">-->
<!--            x264-->
<!--          </a-radio-button>-->
<!--          <a-radio-button value="x265" style="">-->
<!--            x265-->
<!--          </a-radio-button>-->
<!--        </a-radio-group>-->
      </span>
    <span style="float: right;margin-right: 0">
        <el-radio-group v-model="preset"  size="small">
          <el-radio-button :label=0>预设1</el-radio-button>
          <el-radio-button :label=1>预设2</el-radio-button>
          <el-radio-button :label=2>预设3</el-radio-button>
        </el-radio-group>
      </span>
  </div>

  <el-input
      class="panel"
      type="textarea"
      :rows="5"
      placeholder="压制代码"
      v-model="param[tool][preset]">
  </el-input>

  <el-row class="panel"  type="flex">
    <el-col :span="2">
      <span style="float: left;margin: 9px;height: 40px;">
        <el-progress
            type="circle"
            :width="20"
            stroke-linecap="square"
            :text-inside="true"
            :stroke-width="2"
            :percentage="66.57"
        ></el-progress>
      </span>
    </el-col>
    <el-col :span="14">
      <span style="">
        <div style="margin-top: 8px;margin-left: -10px;height: 40px;font-size: 12px">
          {{perLog}}
        </div>
      </span>
    </el-col>
    <el-col :span="4">
      <span  style="float: right;margin-right: -16px">
        <el-button @click="pauseEncode">暂停</el-button>
      </span>
    </el-col>
    <el-col :span="4">
      <span  style="float: right;margin-right: 0">
        <el-button @click="encode">{{startOrQuit}}</el-button>
      </span>
    </el-col>
  </el-row>

<!--    <br>压制代码档位-->
<!--    <br>信息-->
  </div>
</template>

<script>
import Wails from '@wailsapp/runtime';

export default {
  name: "Main",
  data() {
    return {
      input: '',
      output: '',
      param: {
        ffmpeg: ['-vcodec libx264 -crf 20 -preset slow', '', ''],
        x264: ['', '', ''],
        x265: ['', '', ''],
      },
      tool: "ffmpeg",
      preset: 0,
      perLog: "这里是单行日志信息",
      progress: 0,
      paused: false,
      startOrQuit: "开始"
      // toolSelect: 0,
      // progress: 0,
      // status: false  //是否正在执行
    };
  },
  mounted() {
    // let dropbox = document.getElementById('drop_area');
    // dropbox.addEventListener("drop",this.eventDrop,false)
    // dropbox.addEventListener("dragleave",function (e) {
    //   e.stopPropagation();
    //   e.preventDefault();
    //   this.borderhover =  false;
    // })
    // dropbox.addEventListener("dragenter",function (e) {
    //   e.stopPropagation();
    //   e.preventDefault();
    //   this.borderhover =  true;
    // })
    // dropbox.addEventListener("dragover",function (e) {
    //   e.stopPropagation();
    //   e.preventDefault();
    //   this.borderhover =  true
    // })
    Wails.Events.On("SetProgess", (progress) => {
      this.progress = progress;
    });
    Wails.Events.On("SetPerLog", (perLog) => {
      this.perLog = perLog;
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
    window.backend.App.SetupBackend().then(error => {
      if (error !== "") {
        this.$message.error(error)
      }
    });
  },
  methods: {
    // eventDrop: function(e){
    //   this.borderhover = false
    //   e.stopPropagation();
    //   e.preventDefault();  //必填字段
    //   let efile = e.dataTransfer.files[0];
    //   console.log(efile.path);
      // let fileData = e.dataTransfer.files;
      // console.log(fileData);
      // this.uploadFile(fileData)
    // },
    // uploadFile: function (file){   //渲染上传文件
    //   console.log(file)
      // for (let i = 0; i !== file.length; i++) {
      //   let fileJson = {
      //     Url:'',
      //     progressStatl:0,
      //     fileText:'',
      //   };
      //   let video_type=file[i].type == "video/mp4" || file[i].type == "video/ogg";
      //   if(file[i].type.indexOf('image') === 0){  //如果是图片
      //     let fileurl = window.URL.createObjectURL(file[i]); //创建一个url连接,供src属性引用
      //     fileJson.Url = fileurl;
      //   }else if(video_type){
      //     alert("不支持此类型文件")
      //   }
      //   fileJson.fileText = file[i].name;
        // this.fileData.push(fileJson);
      // }
    // setVar () {
    //   window.backend.App.SetSrcPath(this.srcPath)
    //   window.backend.App.SetDstPath(this.dstPath)
    //   window.backend.App.SetParam(this.select, this.param[this.select])
    // },
    getInput () {
      window.backend.App.SelectFileTitle("选择输入文件").then(path => {
        if ( path.length !== 0 ) {
          this.input = path
        }
        //TODO 使用ffprobe获取输入文件参数

        this.generateOutput()
      });
    },
    getOutput () {
      window.backend.App.SelectSaveFileTitle("选择输出文件").then(path => {
        if ( path.length !== 0 ) {
          this.output = path
        }
      });
    },
    generateOutput () {
      //根据输入检测并自动指定输出
      window.backend.App.GenerateOutput(this.input).then(path => {
        if ( path.length !== 0 ) {
          this.output = path
        }
      })
    },
    // onStart () {
    //   this.setVar()
    //   //debug
    //   // this.status = !this.status
    //   //根据情况决定开始/结束压制
    //   if (this.status === false) {
    //     this.StartEncoding()
    //   } else {
    //     this.QuitEncoding()
    //   }
    // },
    encode () {
      if (this.startOrQuit === "开始") {
        this.startOrQuit = "结束"
        window.backend.App.StartEncode(this.input, this.output, this.param[this.tool][this.preset], this.tool).then((error) => {
          if ( error !== null) {
            this.$message.error(error, 5);
          }
          this.startOrQuit = "开始"
          this.paused = false
        });
      } else {
        Wails.Events.Emit("RealtimeSignal", 'q')
      }
    },
    pauseEncode () {
      if ( this.paused ) {
        Wails.Events.Emit("RealtimeSignal", 'r')
      } else {
        Wails.Events.Emit("RealtimeSignal", 'p')
      }
    },
  }
}
</script>

<style scoped>
.panel {
  width: 580px;
  padding-bottom: 10px;
  overflow: hidden;
  /*屏蔽滚动条*/
  -ms-overflow-style: none;
}

.progress {
  padding-bottom: 10px;
}
</style>