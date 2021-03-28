<template>
	<div style="padding: 0 10px">
		<br>包含各个工具的设置
		<!--    <a-card style="text-align: center;height: 150px;margin: 4px 10px 14px;" :hoverable="true">-->
		<!--      <h2 style="font-size: xx-large">One Encoder</h2>-->
		<!--      <p>版本：0.0.1</p>-->
		<!--    </a-card>-->
		<!--    <div class="panel">-->
		<!--      <a-input-search addon-before="FFmpeg API" allow-clear placeholder="" size="large" @search="onSearch">-->
		<!--        <a-button slot="enterButton">-->
		<!--          <a-icon type="question-circle" />-->
		<!--        </a-button>-->
		<!--      </a-input-search>-->
		<!--    </div>-->
		<!--    <div class="panel">-->
		<!--      <a-input-search addon-before="FFmpeg 位置" allow-clear placeholder="" size="large" @search="onSearch">-->
		<!--        <a-button slot="enterButton">-->
		<!--          <a-icon type="question-circle" />-->
		<!--        </a-button>-->
		<!--      </a-input-search>-->
		<!--    </div>-->
		<!--    <div class="panel">-->
		<!--      <a-input addon-before="匹配正则表达式" size="large">-->
		<!--      </a-input>-->
		<!--    </div>-->
		<!--    <div class="panel" style="">-->
		<!--      <a-input-search addon-before="FFmpeg 版本号" allow-clear placeholder="" size="large" @search="onSearch">-->
		<!--        <a-button slot="enterButton">-->
		<!--          检查更新-->
		<!--          <a-icon type="reload" />-->
		<!--        </a-button>-->
		<!--      </a-input-search>-->
		<!--    </div>-->
		<!--    <div class="panel" style="">-->
		<!--      <a-row type="flex" justify="space-around" align="middle">-->
		<!--        <a-col :span="4">-->
		<!--          <div style="">缩放百分比：</div>-->
		<!--        </a-col>-->
		<!--        <a-col :span="8">-->
		<!--          <a-slider v-model="scale" value="scale" :min="100"-->
		<!--                    :max="200" :step="25" :tipFormatter="null"-->
		<!--                    style="width: 90%;margin: auto"/>-->
		<!--        </a-col>-->
		<!--        <a-col :span="2" >-->
		<!--          {{scale}}%-->
		<!--        </a-col>-->
		<!--        <a-col :span="4" :offset="0">-->
		<!--          <a-button size="large" shape="" style="width: 100%">浏览日志</a-button>-->
		<!--        </a-col>-->
		<!--        <a-col :span="4" :offset="0">-->
		<!--          <a-button size="large" shape="" style="width: 105%">程序位置</a-button>-->
		<!--        </a-col>-->
		<!--      </a-row>-->
		<!--      <a-button></a-button>-->
		<!--    </div>-->
	</div>
</template>

<script>
export default {
	name: "Setting",
	data() {
		return {
			scale: 100,
		};
	},
}
</script>

<style scoped>

</style>
