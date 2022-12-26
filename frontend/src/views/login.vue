<template>
    <div class="wrap">
        <div class="introduce">
            <img src="@/assets/images/logo01.png" alt="">
            <p>{{title}}</p>
            <span>Material management system</span>
        </div>
        <div class="login">
            <h1>欢迎登陆</h1>
            <span class="xiahua"></span>
            <div class="info-input">
                <el-input v-model="username" placeholder="请输入用户名" prefix-icon="el-icon-user-solid" maxlength="50" v-trim/>
                <el-input v-model="password" placeholder="请输入密码" type="password" prefix-icon="eva-icon-password" maxlength="30" show-password />
                <span :loading="loading" @click="login" @keypress.enter="login"></span>
            </div>
        </div>
    </div>
</template>

<script>
import { mapMutations } from "vuex";
import {loginByPassword } from "@/api/system/common";
import Cookies from 'js-cookie'

export default {
    name: "Login",
    data() {
        return {
            loading: false,
            username: "",
            password: "",
            // 验证码
            captcha: {
                loading: false,
                value: "",
                uuid: "",
                uri: "",
            },
            title:window._CONFIG['systemTitle']
        };
    },
    methods: {
        ...mapMutations(["setUserInfo"]),
        // 登录
        login() {
            if (this.loading) {
                return;
            }
            if (!this.__check()) {
                return;
            }
            this.loading = true;
            loginByPassword({
                username: this.username.trim(),
                password: this.password,
            })
                .then((res) => {
                  console.info('res.accessToken',res.accessToken)
                  localStorage.setItem("Authorization",res.accessToken)
                  // window.location.href = process.env.VUE_APP_CONTEXT_PATH;
                })
                .catch((e) => {
                  this.$tip.apiFailed(e);
                })
                .finally(() => {
                    this.loading = false;
                });
        },
        __check() {
            if (this.username.trim() === "") {
                this.$tip.error("请输入用户名");
                return false;
            }
            if (this.password === "") {
                this.$tip.error("请输入密码");
                return false;
            }
            return true;
        },
    },
    created() {
        // this.refreshCaptcha()
    },
    mounted(){
        window.addEventListener('keyup',(e) =>{
    if (e && e.keyCode == 13) {
        // console.log(111);
        this.login()
    }
})
    }
};
</script>

<style scoped lang="scss">
@import "@/assets/style/variables.scss";
$input-gap: 30px;
html,
body{
    overflow: auto;
}
.wrap {
    width: 100%;
    height: 100vh;
    background-image: url("../assets/images/beijing.png");
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-clip: content-box;
    background-position: center;
    position: relative;
    overflow: auto;
    // 左边介绍
    .introduce {
        width: 100%;
        // height: 100%;
        box-sizing: border-box;
        color: #fff;
        text-align: center;
        margin-top: 10%;
        p {
            font-size: 48px;
            //   font-style: italic;
            // text-shadow: 0 0 8px #0a4ebc, 0 0 5px #0a4ebc, 0 0 8px #0a4ebc,
            //     0 0 10px #0a4ebc, 0 0 15px #0a4ebc, 0 0 40px #0a4ebc,
            //     0 0 50px #0a4ebc, 0 0 75px #0a4ebc;
            margin-top: 50px;
            font-weight: 300;
            color: #fafafa;
            margin:0
        }
        span {
            font-size: 24px;
            font-weight: 300;
            margin: 0;
            // margin: 10px 70px;
        }
        img{
            width: 82px;
            height: 78px;
        }
    }
    // 右边登录
    .login {
        height: 450px;
        width: 80%;
        min-width: 460px;
        flex-shrink: 0;
        text-align: center;
        background-size: 100% 100%;
        padding: 0 80px;
        box-sizing: border-box;
        margin: 7vh auto;
        h1 {
            font-size: 28px;
            font-weight: 900;
            color: #fff;
            line-height: 20px;
        }
        .info-input {
            width: 65%;
            margin: 0 auto;
            margin-top: $input-gap;
            margin-bottom: 60px;
            display: flex;
            align-items: center;
            /deep/ .el-input {
                margin-top: 30px;
                width: 340px;
                margin: 0 10px 0 65px;
                .el-input__inner {
                    height: 50px;
                    background: rgba(255,255,255,0.3);
                    background-size: 100% 100%;
                    border: none;
                    color: #fff;
                    font-size: 18px;
                }
            }
            span{
                padding: 20px 24px ;
                background: url('../assets/images/login.png') no-repeat;
                background-size: 100%;
            }
        }
        .xiahua{
            display: inline-block;
            width: 2vw;
            height: 2px;
            background-color: #fff;

        }
        .el-button {
            height: 50px;
            width: 100%;
            color: #fff;
            font-size: 16px;
            border: none;
            background: linear-gradient(
                130deg,
                $primary-color + 20 0%,
                $primary-color - 20 100%
            );
        }
    }
}
</style>
