<template>
    <div class="sw-login">
        <div class="sw-login_content">
            <h3>欢迎登陆🍓后台管理系统</h3>
            <el-form
                size="medium"
                :label-position="'right'"
                label-width="100px"
                ref='form'
                :rules="rules"
                :model="form">
                <el-form-item label="用户名：" prop="username">
                    <el-input v-model="form.username"></el-input>
                </el-form-item>
                <el-form-item label="密码：" prop="password">
                    <el-input v-model="form.password" type="password"></el-input>
                </el-form-item>
            </el-form>
            <div class="button">
                <el-button @click="submit()" type="primary">登录</el-button>
            </div>
        </div>
    </div>
</template>

<script>
import { apiLogin } from './api.js'

export default {
    name: 'login',
    data() {
        return {
            form: {
                username: '',
                password: ''
            },
            rules: {
                username: [
                    { required: true, message: '请输入用户名', trigger: 'blur' },
                ],
                password: [
                    { required: true, message: '请输入密码', trigger: 'blur' },
                ],
            }
        }
    },
    
    mounted() {
    
    },
    methods: {
        /**
         * 提交
         */
        submit() {
            this.$refs.form.validate((res) => {
                if (res){
                    apiLogin(this.form).then((result) => {
                        this.$localStorage.setStore('token', result.token)
                        this.$router.push('/notice-list')
                        console.log('登录成功')
                    })
                }
            })
        }
    }
}
</script>

<style lang="less" scoped>
.sw-login {
    width: 100vw;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    overflow: hidden;
    background: url('http://blog.mrabit.com/bing/today') no-repeat center center;

    .sw-login_content {
        max-width: 440px;
        width: calc(100% - 40px);
        padding: 24px;
        background-color: #fff;
        box-shadow: 0 2px 6px rgba(0,0,0,.2);
        min-width: 320px;
        overflow: hidden;

        .button{
            text-align:right;
        }
    }
}
</style>
