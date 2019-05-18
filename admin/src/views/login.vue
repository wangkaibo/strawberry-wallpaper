<template>
<div class="sw-login">
    <div class="sw-login_content">
        <el-form 
            size="medium" 
            :label-position="'right'" 
            label-width="60px" 
            :model="form">
            <el-form-item label="用户名">
                <el-input v-model="form.name"></el-input>
            </el-form-item>
            <el-form-item label="密码">
                <el-input v-model="form.password"></el-input>
            </el-form-item>
            <el-form-item>
                <el-button @click="submit()">登录</el-button>
            </el-form-item>
        </el-form>
    </div>
</div>
</template>

<script>
import { apiLogin } from './api.js';

const aa = {
    kk: '1234'
}
export default {
    name: 'login',
    data() {
        return {
            form: {
                name: '',
                password: ''
            },
        };
    },
    
    mounted() {
        console.log(aa)
    },
    methods: {
    /**
     * 提交
     */
        submit() {
            apiLogin(this.form).then((result) => {
                console.log('登录成功');
            });
        }
    }
};
</script>

<style lang="scss" scoped>
</style>
