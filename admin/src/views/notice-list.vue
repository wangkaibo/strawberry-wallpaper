<template>
	<div class="notice-list">
		<div class="content">
			<transition name="fade" mode="out-in">
				<div class="view" v-show="viewContent.length > 0">
					<!-- <div class="view-header">
						<i
							class="el-icon-close"
							@click="handleCloseView"
						></i>
					</div> -->
					<noticeView :viewContent="viewContent" />
				</div>
			</transition>

			<div class="table">
				<div class="header">
					<el-button type="primary" @click="handleNewClick"
						>新建</el-button
					>
				</div>
				<el-table :data="tableData" style="width: 100%">
					<el-table-column
						label="id"
						prop="id"
						width="50"
					></el-table-column>
					<el-table-column
						label="发布时间"
						prop="showTime"
						width="140"
					></el-table-column>
					<el-table-column
						label="过期时间"
						prop="outTime"
						width="140"
					></el-table-column>
					<el-table-column
						label="内容"
						prop="content"
					></el-table-column>
					<el-table-column label="状态" prop="status" width="100">
						<template slot-scope="scope">
							<span
								:class="{
									'status-icon': true,
									'has-publish': scope.row.status === '1',
									'no-publish': scope.row.status === '2',
									'has-out-time': scope.row.status === '3'
								}"
								>{{ statusOptions[scope.row.status] }}</span
							>
						</template>
					</el-table-column>
					<el-table-column label="操作" width="300">
						<template slot-scope="scope">
							<div class="opteace">
								<el-button
									size="mini"
									v-if="(currentViewNoticeId && currentViewNoticeId !== scope.row.id) || !currentViewNoticeId "
									@click="handleView(scope.row)"
									>预览</el-button
								>
								<el-button
									size="mini"
									v-if="currentViewNoticeId === scope.row.id "
									@click="handleCloseView(scope.row)"
									>关闭预览</el-button
								>
								<el-button
									size="mini"
									v-if="scope.row.status === '2'"
									@click="handlePublist(scope.row, scope.$index)"
									>发布</el-button
								>
								<el-button
									size="mini"
									v-if="scope.row.status === '2'"
									@click="handleEdit(scope.row)"
									>修改</el-button
								>
								<el-button
									size="mini"
									v-if="scope.row.status === '2'||scope.row.status === '3'"
									type="danger"
									@click="
										handleDelete(scope.row, scope.$index)
									"
									>删除</el-button
								>
							</div>
						</template>
					</el-table-column>
				</el-table>
			</div>
		</div>
	</div>
</template>

<script>
import { apiGetNoticeList, apiDeleteNotice, apiPublicNotice } from './api.js'

import { dateFormat } from '../assets/js/until.js'
import noticeView from '../components/notice-view/index.vue'

export default {
    name: 'notice-list',
    components: { noticeView },
    data() {
        return {
            statusOptions: {
                1: '已发布',
                2: '未发布',
                3: '已过期'
            },
            tableData: [],
            viewContent: [],
            currentViewNoticeId: '',
        }
    },
    mounted() {
        apiGetNoticeList().then((res) => {
            this.tableData = res.map(item => ({
                ...item,
                outTime: dateFormat(new Date(item.expire_at), 'yyyy-MM-dd hh:mm'),
                showTime: item.time !== -6795364578871 ? dateFormat(new Date(item.time), 'yyyy-MM-dd hh:mm') : '--'
            }))
        })
    },
    methods: {
        handleNewClick() {
            this.$router.push('/notice-add')
        },
        handleView(data) {
            this.viewContent = [
                {
                    time: data.time,
                    content: data.content
                }
            ]
            this.currentViewNoticeId = data.id
        },

        handleCloseView() {
            this.viewContent = []
            this.currentViewNoticeId = ''
        },

        handlePublist(data, index) {
            apiPublicNotice(data.id, { status: 1 }).then(() => {
                this.tableData[index].status = '1'
                this.tableData[index].showTime = dateFormat(new Date(), 'yyyy-MM-dd hh:mm')
                this.$message.success('发布成功')
            })
        },

        handleEdit(data) {
            this.$localStorage.setStore('noticeDetail', data)
            this.$router.push({
                path: '/notice-add',
                query: {
                    isEdit: true,
                }
            })
        },

        handleDelete(data, index) {
            this.$confirm('确定删除该条公告吗?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {
                this.$message.success('删除成功')
                apiDeleteNotice(data.id).then(() => {
                    this.tableData.splice(index, 1)
                })
            })
        }
    },
    beforeDestroy() {}
}
</script>

<style lang="less" scoped>
.notice-list {
	width: 100%;
	height: 100%;
	display: flex;
	position: relative;
	background: url('http://blog.mrabit.com/bing/today') no-repeat center center;

	.content {
		width: 100%;
		height: 100%;
		background-color: rgba(0,0,0,0.5);

		.view {
			position: absolute;
			top: 0;
			left: 0;
			width: 320px;
			// height: 100%;
			text-align: left;
			flex: none;
			z-index: 99;
			// background-color: #222222;
			overflow: hidden;
			.view-header {
				padding: 0 10px;
				height: 40px;
				line-height: 40px;
				text-align: right;
				color: #ffffff;
			}
		}
		.table {
			width: 100%;
			padding: 0 20px;
			flex: auto;
			.header {
				// padding: 0 20px;
				height: 60px;
				line-height: 60px;
				text-align: right;
			}
			.status-icon {
				display: inline-block;
				height: 30px;
				padding: 0 10px;
				line-height: 30px;
				border-radius: 4px;
				color: #ffffff;
				&.has-publish {
					background-color: #409eff;
				}
				&.no-publish {
					background-color: #67c23a;
				}
				&.has-out-time {
					background-color: #e6a23c;
				}
			}

			.opteace {
				display: flex;
			}
		}
	}
}
</style>

<style>
.fade-enter-active {
	animation: bounce-in 0.3s;
}
.fade-leave-active {
	animation: bounce-in 0.3s reverse;
}
@keyframes bounce-in {
	0% {
		width: 0;
	}
	100% {
		width: 320px;
	}
}
</style>
