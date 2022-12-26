<template>
  <div class="main-header">
    <div class="header">
      <h2>
        <i class="el-icon-s-unfold" v-if="menuData.collapse" @click="switchCollapseMenu(null)"></i>
        <i class="el-icon-s-fold" v-else @click="switchCollapseMenu(null)"></i>
        {{ title }}
      </h2>
      <div class="user">
        <header-notice />
        <el-dropdown trigger="hover">
          <span class="el-dropdown-link kj-dropdown-link">
            欢迎您，{{ userInfo | displayName
            }}<i class="el-icon-arrow-down el-icon--right"></i>
          </span>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item @click.native="changePwd"><i class="el-icon-s-tools"></i>修改密码</el-dropdown-item>
            <el-dropdown-item @click.native="emptyCache"><i class="el-icon-refresh"></i>清理缓存</el-dropdown-item>
            <el-dropdown-item @click.native="logout">
              <i class="el-icon-switch-button"></i>退出登录
            </el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </div>
    </div>
    <!-- 修改密码 -->
    <GlobalWindow title="修改密码" :visible.sync="visible.changePwd" @confirm="confirmChangePwd" @close="visible.changePwd = false">
      <el-form :model="changePwdData.form" ref="changePwdDataForm" :rules="changePwdData.rules">
        <el-form-item label="原始密码" prop="oldpassword" required>
          <el-input v-model="changePwdData.form.oldpassword" type="password" placeholder="请输入原始密码" maxlength="30" show-password></el-input>
        </el-form-item>
        <el-form-item label="新密码" prop="password" required>
          <el-input v-model="changePwdData.form.password" type="password" placeholder="请输入新密码" maxlength="30" show-password></el-input>
        </el-form-item>
        <el-form-item label="确认新密码" prop="confirmpassword" required>
          <el-input v-model="changePwdData.form.confirmpassword" type="password" placeholder="请再次输入新密码" maxlength="30" show-password></el-input>
        </el-form-item>
      </el-form>
    </GlobalWindow>
  </div>
</template>

<script>
import { mapState, mapMutations } from "vuex";
import GlobalWindow from "./GlobalWindow";
import { logout, updatePwd, refleshCache } from "@/api/system/common";
import headerNotice from "./HeaderNotice"
// import { getAction } from "@/api/biz/searchPartyGroup";

export default {
  name: "Header",
  components: { GlobalWindow, headerNotice, },
  data() {
    return {
      visible: {
        // 修改密码
        changePwd: false,
      },
      isWorking: {
        // 修改密码
        changePwd: false,
      },
      username: "bob", // 用户名
      // 修改密码弹框
      changePwdData: {
        form: {
          oldpassword: "",
          password: "",
          confirmpassword: "",
        },
        rules: {
          oldpassword: [{ required: true, message: "请输入原始密码" }],
          password: [{ required: true, message: "请输入新密码" }],
          confirmpassword: [{ required: true, message: "请再次输入新密码" }],
        },
      },
    };
  },
  computed: {
    ...mapState(["menuData", "userInfo", "messageNum"]),
    title() {
      return this.$route.meta.title;
    },
  },
  filters: {
    // 展示名称
    displayName(userInfo) {
      if (userInfo == null) {
        return "";
      }
      if (userInfo.realname != null && userInfo.realname.trim().length > 0) {
        return userInfo.realname;
      }
      return userInfo.username;
    },
  },
  created() {
    // getAction("/business/notice/noticeDetailCount", {}).then((res) => {
    //   this.$store.commit("SET_MESSAGENUM", res);
    // });
    this.$store.commit("SET_MESSAGENUM", 0);
  },
  methods: {
    ...mapMutations(["setUserInfo", "switchCollapseMenu"]),
    // 修改密码
    changePwd() {
      this.visible.changePwd = true;
      this.$nextTick(() => {
        this.$refs.changePwdDataForm.resetFields();
      });
    },
    // 确定修改密码
    confirmChangePwd() {
      if (this.isWorking.changePwd) {
        return;
      }
      this.$refs.changePwdDataForm.validate((valid) => {
        if (!valid) {
          return;
        }
        // 验证两次密码输入是否一致
        if (
          this.changePwdData.form.password !==
          this.changePwdData.form.confirmpassword
        ) {
          this.$tip.warning("两次密码输入不一致");
          return;
        }
        // 执行修改
        this.isWorking.changePwd = true;
        updatePwd({
          oldpassword: this.changePwdData.form.oldpassword,
          password: this.changePwdData.form.password,
          confirmpassword: this.changePwdData.form.confirmpassword,
          username: this.$store.state.userInfo.username,
        })
          .then(() => {
            this.$tip.apiSuccess("修改成功");
            this.visible.changePwd = false;
          })
          .catch((e) => {
            this.$tip.apiFailed(e);
          })
          .finally(() => {
            this.isWorking.changePwd = false;
          });
      });
    },
    // 退出登录
    logout() {
      logout()
        .then(() => {
          this.$router.push({ name: "login" });
          this.setUserInfo(null);
          this.$tip.apiSuccess("退出成功");
        })
        .catch((e) => {
          this.$tip.apiFailed(e);
        });
    },
    emptyCache() {
      refleshCache().then(() => {
        this.$tip.apiSuccess("刷新缓存成功");
      });
    },
    toNext() {
      this.$router.push({
        path: "/message/userMessage",
      });
    },
  },
};
</script>

<style scoped lang="scss">
@import "@/assets/style/variables.scss";
.header {
  overflow: hidden;
  padding: 0 25px;
  background: #fff;
  height: 100%;
  display: flex;
  h2 {
    width: 50%;
    flex-shrink: 0;
    line-height: $header-height;
    font-size: 19px;
    font-weight: 600;
    color: #606263;
    display: inline;
    & > i {
      font-size: 20px;
      margin-right: 12px;
    }
  }
  .user {
    // margin-top: 15px;
    width: 50%;
    flex-shrink: 0;
    text-align: right;
    .el-dropdown {
      top: 2px;
    }
    img {
      width: 32px;
      position: relative;
      top: 10px;
      margin-right: 10px;
    }

    .kj-icon {
      padding: 15px 12px 15px 12px;
    }

    .kj-icon:hover {
      background: #f7f8f9;
      cursor: pointer;
      border-radius: 4px;
    }
  }
}

// 下拉菜单框
.el-dropdown-menu {
  width: 140px;
  .el-dropdown-menu__item:hover {
    background: #e3edfb;
    color: $primary-color;
  }
}

.el-dropdown-menu--mini .el-dropdown-menu__item {
  line-height: 36px !important;
}

.kj-dropdown-link {
  display: inline-flex;
  height: 56px;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  padding: 0 10px 0 15px;
}

.kj-dropdown-link:hover {
  background: #f1f3f4;
}
</style>
