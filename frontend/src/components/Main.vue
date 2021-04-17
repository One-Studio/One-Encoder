<template>
	<q-page>
		<q-input dense hide-bottom-space outlined v-model="input" label="请选择输入" class="q-mb-md">
			<template v-slot:append>
				<q-separator vertical />
				<q-avatar class="q-ml-sm" icon="far fa-folder-open" @click="getInput" />
			</template>
		</q-input>

		<input type="file" name="filename" id="open" style="display:none"/>

		<q-input dense hide-bottom-space outlined v-model="output" label="请选择输出" class="q-mb-md">
			<template v-slot:append>
				<q-separator vertical />
				<q-avatar class="q-ml-sm" icon="far fa-folder-open" @click="getOutput" />
			</template>
		</q-input>



		<div class="panel q-mb-md">
			<a-row type="flex" justify="space-between">
				<a-radio-group button-style="outline" size="large" v-model="tool">
					<a-radio-button value="ffmpeg">ffmpeg</a-radio-button>
					<a-radio-button value="x264">x264</a-radio-button>
					<a-radio-button value="x265">x265</a-radio-button>
				</a-radio-group>
				<a-radio-group button-style="outline" size="large" v-model="preset">
					<a-radio-button value="0">预设1</a-radio-button>
					<a-radio-button value="1">预设2</a-radio-button>
					<a-radio-button value="2">预设3</a-radio-button>
				</a-radio-group>
			</a-row>
		</div>

		<q-input
			v-model="param[tool][preset]"
			filled
			type="textarea"
			class="q-mb-md"
		/>

		<div class="col-12 row justify-between">
			<div class="col-10 text-left">
				<q-knob
					v-model="progress"
					track-color="grey-3"
					size="26px"
					color="primary"
					class="q-mr-md"
				/>
				<span class="q-mr-md">{{ progress }}%</span>
				<span>{{ perLog }}</span>
			</div>
			<div class="col-2 text-right">
				<q-btn outline text-color="primary" label="暂停" />
				<q-btn class="q-ml-sm" outline text-color="primary" label="开始" />
			</div>
		</div>
	</q-page>
</template>

<script>
import Wails from '@wailsapp/runtime'

export default {
	name: 'Main',
	data() {
		// this.$i18n.locale = 'en';
		return {
			input: '',
			output: '',
			param: {
				ffmpeg: ['-vcodec libx264 -crf 20 -preset slow', '-c:v prores', ''],
				x264: ['--crf 20 --preset slow', '', ''],
				x265: ['', '', ''],
			},
			tool: 'ffmpeg',
			preset: '0',
			perLog: '这里是单行日志信息',
			progress: 66.57,
			paused: false,
			startOrQuit: '开始',
			// toolSelect: 0,
			// progress: 0,
			// status: false  //是否正在执行
			progressFinished: true,
			progressDisplayNum: false,
			file: null,
		}
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
		Wails.Events.On('SetProgess', (progress) => {
			this.progress = progress
		})
		Wails.Events.On('SetPerLog', (perLog) => {
			this.perLog = perLog
		})
		Wails.Events.On('NoticeSuccess', (msg) => {
			this.$message.success(msg, 5)
		})
		Wails.Events.On('NoticeError', (msg) => {
			this.$message.error(msg, 5)
		})
		Wails.Events.On('NoticeWarning', (msg) => {
			this.$message.warning(msg, 5)
		})
		//通知传参
		window.backend.App.SetupBackend().then((error) => {
			if (error !== '') {
				this.$message.error(error)
			}
		})
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
		getInput() {
			// console.log("选择输入文件debug")
			// document.getElementById('open').click()
			// console.log('document.getElementById(\'open\').files', document.getElementById('open').files)
			window.backend.App.SelectFileTitle('选择输入文件').then((path) => {
				if (path.length !== 0) {
					this.input = path
					this.generateOutput()
				}
				//TODO 使用ffprobe获取输入文件参数

			})
		},
		getOutput() {
			window.backend.App.SelectSaveFileTitle('选择输出文件').then((path) => {
				if (path.length !== 0) {
					this.output = path
				}
			})
		},
		generateOutput() {
			//根据输入检测并自动指定输出
			window.backend.App.GenerateOutput(this.input).then((path) => {
				if (path.length !== 0) {
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
		encode() {
			if (this.startOrQuit === '开始') {
				this.startOrQuit = '结束'
				window.backend.App.StartEncode(
					this.input,
					this.output,
					this.param[this.tool][this.preset],
					this.tool
				).then((error) => {
					if (error !== null) {
						this.$message.error(error, 5)
					}
					this.startOrQuit = '开始'
					this.paused = false
					this.progressFinished = false
				})
			} else {
				Wails.Events.Emit('RealtimeSignal', 'q')
			}
		},
		pauseEncode() {
			if (this.paused) {
				Wails.Events.Emit('RealtimeSignal', 'r')
			} else {
				Wails.Events.Emit('RealtimeSignal', 'p')
			}
		},
	},
}
</script>
