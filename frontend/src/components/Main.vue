<template>
  <div class="main-container">
    <a-form-model class="main-form" layout='horizontal' :label-col={span:4} :wrapper-col={span:17}>
      <a-form-model-item label="文件路径">
        <a-input class="test" style="" v-model="srcPath" allow-clear placeholder=""  @search="onSearch">
        </a-input>
        <a-button slot="enterButton" @click="getSrcPath" style="height: 6vw">
          <a-icon type="folder" />
        </a-button>
      </a-form-model-item>
      <a-form-model-item label="输出路径">
        <a-input-search v-model="dstPath" addon-before="输出路径" allow-clear placeholder="" size="large" @search="onSearch">
          <a-button slot="enterButton" @click="getDistPath">
            <a-icon type="folder" />
          </a-button>
        </a-input-search>
      </a-form-model-item>
      <a-form-model-item label="压制代码">
        <a-textarea  placeholder="输入FFmpeg压制的代码"
                     size="large"
                     style="font-size: 2.5vw;height: 20vw"
                     :auto-size="{ minRows: 1, maxRows: 1}"
                     addon-before="文件路径"
        />
      </a-form-model-item>
      <a-form-model-item label="上传文件">
        <a-upload-dragger
            name="file"
            :multiple="true"
            action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
            @change="handleDragFiles"
        >
          <a-icon type="file" style="font-size: xx-large;margin: 10px;height: 5vw"/>
          <p class="ant-upload-text">
            选择或拖拽想要压制的文件到此处
          </p>
        </a-upload-dragger>
      </a-form-model-item>
    </a-form-model>

    <div class="panel">
      <a-row type="flex" justify="space-around" align="middle">
        <a-col :span="16">
          <a-progress size="large" stroke-linecap="square" :percent="percent" class="progressbar"/>
        </a-col>
        <a-col :span="3" :offset="0">
            <a-button size="large" style="width: 100%">
              暂停
            </a-button>
        </a-col>
        <a-col :span="4" :offset="0">
            <a-button size="large" style="width: 104%">
              开始
            </a-button>
        </a-col>
      </a-row>
    </div>
  </div>
</template>

<script>
import Wails from '@wailsapp/runtime';

export default {
  name: "Main",
  data() {
    return {
      percent: 0,
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

* {
  font-size: 2.75vw;
}


.main-container{
  /*padding: 50x;*/
  /*margin: 0px 8px 8px 8px;*/
}

.progressbar{
  /*width: 400px;*/
  margin: auto;
}

.panel{
  /*height: 10vw;*/
}

input {
  height: 12vw;font-size: 6vw;
}

</style>

<style>
  .main-form{
    font-size: 10vw;
  }
</style>