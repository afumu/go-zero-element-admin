<template>
  <GlobalWindow :title="title" :visible.sync="visible" :confirm-working="isWorking" @confirm="confirm" @closeTrigger="$refs.form.resetFields()">
    <el-form :model="form" ref="form" :rules="rules">
      <el-form-item label="角色编码" prop="roleCode" required>
        <el-input v-model="form.roleCode" placeholder="请输入角色编码" v-trim maxlength="50" />
      </el-form-item>
      <el-form-item label="角色名称" prop="roleName" required>
        <el-input v-model="form.roleName" placeholder="请输入角色名称" v-trim maxlength="50" />
      </el-form-item>
      <el-form-item label="角色备注" prop="description">
        <el-input v-model="form.description" type="textarea" placeholder="请输入角色备注" :rows="3" v-trim maxlength="500" />
      </el-form-item>
    </el-form>
  </GlobalWindow>
</template>

<script>
import BaseOpera from "@/components/base/BaseOpera";
import GlobalWindow from "@/components/common/GlobalWindow";
export default {
  name: "OperaRoleWindow",
  extends: BaseOpera,
  components: { GlobalWindow },
  data() {
    return {
      // 原角色码
      originRoleCode: "",
      // 表单数据
      form: {
        id: undefined,
        roleCode: "",
        roleName: "",
        description: "",
      },
      // 验证规则
      rules: {
        roleCode: [{ required: true, message: "请输入角色编码" }],
        roleName: [{ required: true, message: "请输入角色名称" }],
      },
    };
  },
  methods: {
    open(title, target) {
      this.title = title;
      this.visible = true;
      this.form.id = "";
      this.form.roleCode = "";
      this.form.roleName = "";
      this.form.description = "";
      // 新建
      if (target == null) {
        this.$nextTick(() => {
          this.$refs.form.resetFields();
          this.form[this.configData["field.id"]] = undefined;
        });
        return;
      }
      // 编辑
      this.$nextTick(() => {
        this.originRoleCode = target.code;
        for (const key in this.form) {
          this.form[key] = target[key];
        }
      });
    },
    confirm() {
      if (
        this.form.id == null ||
        this.form.id === "" ||
        this.form.id === undefined
      ) {
        this.__confirmCreate();
        return;
      }
      if (this.originRoleCode === this.form.code) {
        this.__confirmEdit();
        return;
      }
      // 修改了角色编码
      this.$dialog
        .confirm(
          "检测到您修改了角色编码，角色编码修改后前后端均可能需要调整代码，确认修改吗？",
          "提示",
          {
            confirmButtonText: "确认修改",
            type: "warning",
          }
        )
        .then(() => {
          this.__confirmEdit();
        });
    },
  },
  created() {
    this.config({
      api: "/system/role",
    });
  },
};
</script>
