<template>
  <!-- 新建/修改 -->
  <GlobalWindow :title="title" :visible.sync="visible" :withFooter="withFooter" :confirm-working="isWorking" @confirm="addUserConfirm" @closeTrigger="closeTrigger">
    <el-form :model="form" ref="form" :rules="rules">
      <el-form-item label="用户账号" prop="username">
        <el-input :disabled="editFlag" v-model="form.username" placeholder="请输入账号" v-trim maxlength="50" />
      </el-form-item>

      <el-form-item v-if="form.id == null" label="登录密码" prop="password">
        <el-input v-model="form.password" type="password" placeholder="请输入登录密码" v-trim   maxlength="30" show-password />
      </el-form-item>

      <el-form-item v-if="form.id == null" label="确认密码" prop="confirmpassword">
        <el-input v-model="form.confirmpassword" type="password" placeholder="请重新输入登录密码" v-trim  maxlength="30" show-password />
      </el-form-item>

      <el-form-item label="用户姓名" prop="realname">
        <el-input v-model="form.realname" placeholder="请输入用户姓名" v-trim maxlength="50" />
      </el-form-item>

      <el-form-item label="工号" prop="workNo">
        <el-input v-model="form.workNo" placeholder="请输入工号" v-trim maxlength="50" />
      </el-form-item>

      <el-form-item label="角色权限" prop="selectedroles">
        <TreeSelect placeholder="请选择角色权限" v-model="form.selectedrolesArr" :data="allRole" :clearable="true" :append-to-body="false" :inline="true" :multiple="true" :flat="false" @input="$emit('input', $event)" />
      </el-form-item>

      <el-form-item label="部门分配" prop="selecteddeparts">
        <TreeSelect placeholder="请选择部门" v-model="form.selecteddeparts" :check-strictly="true" :data="allDepart" node-key="id" :clearable="true" :append-to-body="false" :inline="true" :multiple="false" :flat="false" @input="$emit('input', $event)" />
      </el-form-item>

      <el-form-item label="身份" prop="userIdentity">
        <el-radio-group v-model="form.userIdentity">
          <el-radio :label="1">普通用户</el-radio>
          <el-radio :label="2">上级</el-radio>
        </el-radio-group>
      </el-form-item>

      <el-form-item label="生日" prop="birthday">
        <el-date-picker  v-model="form.birthday"  value-format="yyyy-MM-dd"  type="date"  placeholder="请选择出生年月" />
      </el-form-item>

      <el-form-item label="性别" prop="sex">
        <el-select clearable v-model="form.sex" placeholder="请选择性别">
          <el-option label="男" :value="1"></el-option>
          <el-option label="女" :value="2"></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="手机号码" prop="phone">
        <el-input v-model="form.phone" placeholder="请输入手机号码" v-trim maxlength="50" />
      </el-form-item>

      <el-form-item label="邮箱" prop="email">
        <el-input v-model="form.email" placeholder="请输入邮箱" v-trim maxlength="50" />
      </el-form-item>

    </el-form>
  </GlobalWindow>
</template>

<script>
import BaseOpera from "@/components/base/BaseOpera";
import GlobalWindow from "@/components/common/GlobalWindow";
import PositionSelect from "@/components/common/PositionSelect";
import TreeSelect from "@/components/common/TreeSelect";
import request from '@/utils/request'

export default {
  name: "OperaUserWindow",
  extends: BaseOpera,
  components: {
    PositionSelect,
    GlobalWindow,
    TreeSelect,
  },
  data() {
    var confirmUsername = (rule, value, callback) => {
      // this.form.confirmpassword = value;
     if(!value){
          return callback(new Error("请输入用户账号"));
      } else if(value !== '') {
            var params = {
                tableName: 'sys_user',
                fieldName: 'username',
                fieldVal: value,
                dataId: '/sys/user/generateUserId'
            };
            request.get('/sys/duplicate/check',{params:params}).then((res) => {
                if (res.success) {
                    callback()
            } 
      }).catch(e =>{
           return callback(new Error("用户账号已存在"));
      })
      }
    };
    // 校验两次密码是否一致
    var confirmPassword = (rule, value, callback) => {
      this.form.confirmpassword = value;
      if (value) {
        if (value === this.form.password) {
          return callback();
        } else {
          return callback(new Error("两次输入密码不一致!"));
        }
      } else {
        return callback(new Error("请再次输入密码"));
      }
    };
    return {
      form: {
        id: null,
        username: null, // 用户账号
        password: null, // 登录密码
        confirmpassword: null, // 登录密码
        realname: null, // 用户姓名
        selectedroles: null,// 角色权限
        selectedrolesArr: null,// 角色权限（数组形式）
        selecteddeparts: null,// 部门分配
        // selecteddepartsArr: null,
        birthday: null,// 出生日期
        sex: null, // 性别
        userIdentity: 1,
        phone: null,
        email: null,
        workNo: null,
        orgCodeTxt:null

      },
      allUserIdentity: this.$dict.getAllDictItemByCode(
        this.$consts.dicConstant.userIdentity
      ),
      allSex: this.$dict.getAllDictItemByCode(this.$consts.dicConstant.sex),
      allFrontlineFlag: this.$dict.getAllDictItemByCode(this.$consts.dicConstant.frontlineFlag),
      allRole: [],
      allDepart: [],
      withFooter: true,// 是否显示footer
      allPosition: [],
      editFlag: false,
      // 验证规则
      rules: {
        username: [
          {
            trigger: "blur",
            required: true,
            // message:'用户名已存在',
            validator: confirmUsername,
          },
        ],
        realname: [
          {
            required: true,
            message: "请输入用户账号",
            trigger: "blur",
          },
        ],
        password: [
          {
            required: true,
            message: "请输入密码",
            trigger: "blur",
          },
        ],
        confirmpassword: [
          {
            trigger: "blur",
            required: true,
            message: "两次密码不一致",
            validator: confirmPassword,
          },
        ],
        workNo: [
          {
            required: true,
            message: "请输入工号",
            trigger: "blur",
          },
        ],
        selectedroles: [
          {
            required: true,
            message: "请选择角色权限",
            trigger: "blur",
          },
        ],
        selecteddeparts: [
          {
            required: true,
            message: "请选择部门",
            trigger: "blur",
          },
        ],
      }
    }
  },
  methods: {
    closeTrigger() {
      let form = this.$refs["form"];
      if (form) {
        form.resetFields();
        form.clearValidate();
      }
    },
    addUserConfirm() {
      if (this.form.selectedrolesArr) {
        this.form.selectedroles = this.form.selectedrolesArr.join(",");
      }
      this.confirm();
    },
    // 确认新建/修改
    confirm () {
      // 如果是编辑，不验证用户名
      if(this.form.id){
        this.rules.username = []
      }
      this.$refs.form.validate((valid) => {
        if (!valid) {
          return
        }
        if (this.form.id == null || this.form.id === '') {
          this.__confirmCreate()
          return
        }
        this.__confirmEdit()
      })

    },
    // 确认修改
    __confirmEdit () {
      // 调用修改接口
      this.isWorking = true
      this.api.updateById(this.form)
        .then(() => {
          this.visible = false
          this.$tip.apiSuccess('修改成功')
          this.$emit('success')
        })
        .catch(e => {
          this.$tip.apiFailed(e)
          this.$emit("fail",e)
        })
        .finally(() => {
          this.isWorking = false
        })
    },
    /**
     * @title 窗口标题
     * @target 编辑的用户对象
     */
    open(title, target) {
    //   console.log(target.selecteddeparts)
      this.title = title;
      this.visible = true;
    //   console.log(this.form)
      // 新建
      if (target == null) {
        this.$nextTick(() => {
          this.$refs.form.resetFields();
          this.form.id = null;
          this.form.departmentId = null;
          this.form.selectedrolesArr = [];
          this.form.typeArr = [];
        });
        return;
      }
      // 编辑
      this.$nextTick(() => {
        console.log(target)
        for (const key in this.form) {
          this.form[key] = target[key];

        }
      });
    },

    // 获取所有的角色
    queryAllRole() {
      this.api.queryAllRole().then((res) => {
        this.allRole = [];
        this.fillRoleData(this.allRole, res);
      });
    },
    //查询所有部门
    queryDepart() {
      this.api.queryAllDepart().then((res) => {
        this.allDepart = [];
        this.fillData(this.allDepart, res);
      });
    },
    //填充部门树
    fillData(list, pool) {
      for (const dept of pool) {
        const deptNode = {
          id: dept.value,
          label: dept.title,
        };
        list.push(deptNode);
        if (dept.children != null && dept.children.length > 0) {
          deptNode.children = [];
          this.fillData(deptNode.children, dept.children);
          if (deptNode.children.length === 0) {
            deptNode.children = undefined;
          }
        }
      }
    },

    // 填充角色能够多选
    fillRoleData(list, pool) {
      for (const role of pool) {
        const roleNode = {
          id: role.id,
          label: role.roleName,
        };
        list.push(roleNode);
      }
    },
  },
  created() {
    this.config({
      model:'user',
      api: '/system/user'
    })
    this.queryAllRole()
    this.queryDepart()
  }
}
</script>

<style lang="scss" scoped>
.global-window {
  /deep/ .el-date-editor {
    width: 100%;
  }

  // 表单头像处理
  /deep/ .form-item-avatar {
    .el-radio.is-bordered {
      height: auto;
      padding: 6px;
      margin: 0 10px 0 0;

      .el-radio__input {
        display: none;
      }

      .el-radio__label {
        padding: 0;
        width: 48px;
        display: block;

        img {
          width: 100%;
        }
      }
    }
  }
}
</style>
