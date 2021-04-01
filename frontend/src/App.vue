<template>
	<div id="app" class="container" :style="`transform: scale(${scale});`">
		<div class="opacity-container" :style="`opacity: ${opacity}`">
						<van-tabs
							type="line"
							animated
							line-height="2"
							line-width="60"
							background="#fefefe"
              duration="0.4"
							color="rgb(64,158,255)"
						>
							<!--          color="#9370DB"-->
							<!--        title-active-color="pink"-->
							<van-tab title="压制" class="tab-container"><Main></Main></van-tab>
							<van-tab title="信息" class="tab-container"><Info></Info></van-tab>
							<van-tab title="设置" class="tab-container"><Setting></Setting></van-tab>
              <van-tab title="关于" class="tab-container"><About></About></van-tab>
						</van-tabs>
<!--			<a-tabs default-active-key="1" style="margin-top: -6px">-->
<!--				<a-tab-pane key="1">-->
<!--					<span slot="tab">-->
<!--						压制-->
<!--					</span>-->
<!--					<Main></Main>-->
<!--				</a-tab-pane>-->
<!--				<a-tab-pane key="2">-->
<!--					<span slot="tab">-->
<!--            MediaInfo-->
<!--					</span>-->
<!--					<Setting></Setting>-->
<!--				</a-tab-pane>-->
<!--				<a-tab-pane key="3">-->
<!--					<span slot="tab">-->
<!--						关于-->
<!--					</span>-->
<!--					<About></About>-->
<!--				</a-tab-pane>-->
<!--        <div slot="tabBarExtraContent" style="margin:auto">-->
<!--          <a-icon type="setting" style="" />-->
<!--        </div>-->
<!--			</a-tabs>-->
		</div>
		<!--    <Main/>-->
	</div>
</template>

<script>
import Main from "@/components/Main";
import Info from "@/components/Info";
import Setting from "@/components/Setting";
import About from "@/components/About";
// import Test from "@/components/Test";
import "./assets/css/main.css";

export default {
	name: "app",
	components: {
		Main,
    Info,
		Setting,
		About
	},
	data() {
		return {
			scale: 1,
			opacity: 0,
      appWidth: 660
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
			this.scale = h / this.appWidth;  //应用目标宽度
		}
	}
};
</script>

<style>

</style>
