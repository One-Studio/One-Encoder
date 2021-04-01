<template>
	<div id="app" class="container" :style="`transform: scale(${scale});`">
		<div class="opacity-container" :style="`opacity: ${opacity}`">
			<!--			<van-tabs-->
			<!--				type="line"-->
			<!--				animated-->
			<!--				line-height="1.5"-->
			<!--				line-width="50"-->
			<!--				background="#fefefe"-->
			<!--				color="rgb(64,158,255)"-->
			<!--			>-->
			<!--				&lt;!&ndash;          color="#9370DB"&ndash;&gt;-->
			<!--				&lt;!&ndash;        title-active-color="pink"&ndash;&gt;-->
			<!--				<van-tab title="主页" style="padding: 10px">-->
			<!--					<Main></Main>-->
			<!--				</van-tab>-->
			<!--				<van-tab title="设置" style="padding: 10px">-->
			<!--					<Setting></Setting>-->
			<!--				</van-tab>-->
			<!--				<van-tab title="关于" style="padding: 10px">-->
			<!--					<About></About>-->
			<!--				</van-tab>-->
			<!--			</van-tabs>-->
			<a-tabs default-active-key="1">
				<a-tab-pane key="1">
					<span slot="tab">
						主页
					</span>
					<Main></Main>
				</a-tab-pane>
				<a-tab-pane key="2">
					<span slot="tab">
						设置
					</span>
					<Setting></Setting>
				</a-tab-pane>
				<a-tab-pane key="3">
					<span slot="tab">
						关于
					</span>
					<About></About>
				</a-tab-pane>
			</a-tabs>
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
		setTimeout(() => {
			this.opacity = 1
		}, 300)
		window.onresize = () => {
			this.adaptScale()
		}
	},
	methods: {
		adaptScale() {
			let h = document.body.clientWidth; //获得屏幕宽度
			this.scale = h / 640;
		}
	}
};
</script>

<style>

</style>
