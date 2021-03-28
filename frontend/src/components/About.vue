<template>
	<div style="padding: 0 10px">
		<br>介绍开发者
		<br>用到的技术栈
		<br>项目地址
		<br>开源协议
		<!--    <el-button @click="test">测试</el-button>-->
		<h2>{{ testLog }}</h2>
	</div>
</template>

<script>
export default {
	name: "About",
	data() {
		return {
			testLog: 'none'
		};
	},
	methods: {
		test() {
			window.backend.App.Test("str1", "str2", "str3").then(data => {
				this.testLog = data
			})
		}
	}
}
</script>

<style scoped>

</style>
