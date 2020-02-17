<template>
    <div class="notice-list">
        <div class="header">
            <el-button type="primary" @click="handleNewClick">新建</el-button>
        </div>

        <el-table :data="tableData" style="width: 100%">
            <el-table-column label="id" prop="id" width="50"></el-table-column>
            <el-table-column label="发布时间" prop="showTime" width="100"></el-table-column>
            <el-table-column label="过期时间" prop="outTime" width="100"></el-table-column>
            <el-table-column label="内容" prop="content"></el-table-column>
            <el-table-column label="状态" prop="status" width="100">
                <template slot-scope="scope" >
                    <span :class="{
                        'status-icon':true,
                        'has-publish':scope.row.status==='1'
                        ,'no-publish' :scope.row.status==='2' 
                        ,'has-out-time' :scope.row.status==='3' 
                    }">{{ statusOptions[scope.row.status] }}</span>
                </template>
            </el-table-column>
            <el-table-column label="操作" width="300">
                <template slot-scope="scope">
                    <div class="opteace">
                    <el-button size="mini"  @click="handleEdit(scope.$index, scope.row)">预览</el-button>
                    <el-button size="mini" v-if="scope.row.status==='2'"  @click="handleEdit(scope.$index, scope.row)">发布</el-button>
                    <el-button size="mini" v-if="scope.row.status==='2'" @click="handleEdit(scope.$index, scope.row)">修改</el-button>
                    <el-button size="mini" v-if="scope.row.status==='2'" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
                    </div>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script>

import { apiGetNoticeList } from './api.js'

import { dateFormat } from '../assets/js/until.js'

export default {
    name: 'notice-list',
    components: {},
    data() {
        return {
            statusOptions: {
                1: '已发布',
                2: '未发布',
                3: '已过期',
            },
            tableData: []
        }
    },
    mounted() {
        apiGetNoticeList().then(res=>{
            this.tableData=res.map(item=>({
                ...item,
                outTime:dateFormat(new Date(item.expire_at), 'MM-dd hh:mm'),
                showTime:dateFormat(new Date(item.time), 'MM-dd hh:mm'),
            }))
        })
    },
    methods: {
        handleNewClick(){
            this.$router.push('/notice-add')
        }
    },
    beforeDestroy(){
       
    }
}
</script>

<style lang="less" scoped>
.notice-list {
    width: 100%;
    height: 100%;

    .header {
        padding: 0 20px;
        height: 60px;
        line-height: 60px;
        text-align: right;
    }

    .status-icon{
        display: inline-block;
        height: 30px;
        padding:0 10px;
        line-height: 30px;
        border-radius:4px;
        color: #ffffff;
        &.has-publish{
            background-color: #409eff;
        }
        &.no-publish{
             background-color: #67c23a;
        }
        &.has-out-time{
            background-color: #e6a23c;
        }
    }

    .opteace{
        display:flex;
    }
}
</style>
