<template>
  <div class="hw-container">
    114514
    <h1>{{message}}</h1>
    <a-button @click="getMessage">按钮</a-button>
    <h1>{{errorMessage}}</h1>

    <a-space size="middle">
      <a-progress stroke-linecap="square" :percent="percent" class="progressbar"/>
      <a-button type="primary" @dblclick="popup" class="ant-btn">
        Popup
      </a-button>
      <a-button type="primary" @click="add">
        增加进度
      </a-button>
      <a-button type="primary" @click="minus">
        减小进度
      </a-button>
    </a-space>
  </div>
</template>

<!--suppress JSUnresolvedVariable -->
<script>
//使用Wails要导入
//import Wails from '@wailsapp/runtime';
export default {
  data() {
    return {
      percent: 0,
      message: " ",
      errorMessage: ""
    };
  },
  mounted() {
    // Wails.Events.On("error", (message, number) => {
    //   let result = number * 1.0001;
    //   this.errorMessage = `${message}: ${result}`;
    //   setTimeout(() => {
    //     this.errorMessage = "";
    //   }, 3000);
    // });
  },
  methods: {
    //测试消息
    getMessage: function() {
      let self = this;
      //调用Go的方法返回结果并复制
      window.backend.App.SayHello().then(result => {
        self.message = result;
      });
      //发送wails信息->Go?
      window.wails.Events.Emit("error", "这是一条错误信息！")
    },
    //弹出确认框
    popup: function () {
      this.$confirm({
        title: '确认框标题',
        content: '确认框内容',
        onOk() {
          return new Promise((resolve, reject) => {
            setTimeout(Math.random() > 0.5 ? resolve : reject, 1000);
          }).catch(() => console.log('Oops errors!'));
        },
        onCancel() {},
      });
    },
    //进度条进度+1
    add: function () {
      if (this.percent < 100) {
        this.percent += 1;
      }
    },
    minus: function () {
      if (this.percent > 0) {
        this.percent -= 1;
      }
    }
  }
};
</script>

<style scoped>
.hw-container {
  text-align: center;
  height: 100%;
  /*background: #ffffff;*/
  margin: 0;
  padding: 0;
}
h1 {
  color: #131313;
  margin-top: 2em;
  position: relative;
  min-height: 5rem;
  width: 100%;
}
a:hover {
  font-size: 1.7em;
  background-color: blue;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
  transition: 500ms;
}
a {
  font-size: 1.7em;
  background-color: #121212;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
}
.progressbar {
  width: 300px;
  /*margin: 48px;*/
}
</style>
