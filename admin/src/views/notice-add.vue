<template>
	<div class="notice-add">
		<div class="edit">
			<el-tabs v-model="activeName" @tab-click="handleTabClick">
				<el-tab-pane label="纯文本" name="text">
					<el-input
						type="textarea"
						placeholder="请输入内容"
						v-model="currentHtml"
						@input="handleValChange"
						rows="15"
						show-word-limit
					>
					</el-input>
				</el-tab-pane>
				<el-tab-pane label="富文本" name="richText">
					<div ref="editor"></div>
				</el-tab-pane>
			</el-tabs>
			<div class="out-time-input">
				<span>设置过期时间(单位:天)</span>
				<el-input
					class="input"
					v-model="outTime"
                    @input="handleOutTimeInput"
					placeholder="请输入内容"
				>
					<template slot="append">天</template>
				</el-input>
			</div>

			<div class="add-button">
				<el-button type="primary" @click="handleSubmit">提交</el-button>
			</div>

			<div class="edit-html">
				{{ currentHtml }}
			</div>
		</div>
		<div class="view">
			<noticeView :viewContent="viewContent" />
		</div>
	</div>
</template>

<script>
import noticeView from "../components/notice-view/index.vue";

const E = window.wangEditor;

let editor = null;

export default {
	name: "notice-add",
	components: {
		noticeView
	},
	data() {
		return {
			noticeList: [],
			viewContent: [],
			currentActiveKey: "",
			activeName: "text",
			outTime: 30,
			currentHtml: "",
			noticeDetail:{},
		};
	},
	mounted() {
		const isEdit = this.$route.query.isEdit||false;
		if(isEdit){
			this.noticeDetail=this.$localStorage.getStore('noticeDetail')
			this.currentHtml=this.noticeDetail.content||''
		}
	},
	methods: {
		handleTabClick(tag) {
			if (tag.name === "richText" && !editor) {
				this.$nextTick(() => {
					editor = new E(this.$refs.editor);
					editor.customConfig.menus = [
						"head", // 标题
						"bold", // 粗体
						"fontSize", // 字号
						"italic", // 斜体
						"underline", // 下划线
						"strikeThrough", // 删除线
						"foreColor", // 文字颜色
						"backColor", // 背景颜色
						"link", // 插入链接
						"list", // 列表
						"justify", // 对齐方式
						"emoticon", // 表情
						"image", // 插入图片
						"video", // 插入视频
						"code" // 插入代码
					];
				
					editor.customConfig.onchange = html => {
						console.log("=======", html);
						this.handleValChange(html);
					};
					editor.create();
					editor.txt.html(this.currentHtml)
				});
			}
		},

		handleValChange(val) {
			console.log(val);
			this.currentHtml = val;
			this.viewContent = [
				{
					time: new Date().getTime(),
					content: val
				}
			];
        },

        handleOutTimeInput(val){
            if(Number.isNaN(Number(val))){
                this.outTime=30
            }
        },
        
        handleSubmit(){
            if(this.currentHtml!==''&&this.outTime){
				console.log('==========提交')
				// 提交成功后
				this.$localStorage.setStore('noticeDetail',{})
            }
            else{
                this.$message.error('请补全信息')
            }
        }
	},
	beforeDestroy() {}
};
</script>


<style lang="less">
.notice-add {
	width: 100%;
	height: 100%;
	display: flex;
	padding: 20px;
	text-align: left;
	background-color: #ffffff;

	.edit {
		width: 100%;
		flex: auto;
		padding: 20px;
		border-right: 1px solid #ccc;

		.out-time-input {
			display: flex;
			height: 60px;
			align-items: center;
			.input {
				width: 200px;
				margin-left: 20px;
			}
		}

		.add-button {
			height: 60px;
			line-height: 60px;
			text-align: right;
		}

		.edit-html {
			width: 100%;
			height: 40%;
			border: 1px solid #ccc;
			text-align: left;
		}
	}

	.view {
		margin-left: 20px;
		flex: none;
	}
}
</style>
