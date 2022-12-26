<template>
  <GlobalWindow :title="title" :visible.sync="visible" :confirm-working="isWorking" @confirm="confirm" @closeTrigger="$refs.form.resetFields()">
    <el-form :model="form" ref="form" :rules="rules">
      <el-form-item label="名称" prop="name" >
        <el-input v-model="form.name" placeholder="请输入角色名称" v-trim maxlength="50" />
      </el-form-item>
      <el-form-item label="描述" prop="memo">
        <el-input v-model="form.memo" type="textarea" placeholder="请输入角色备注" :rows="3" v-trim maxlength="500" />
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
      // 表单数据
      form: {
        id: null,
        name:'',
        memo:''
      },
      // 验证规则
      rules: {
        name: [{ required: true, message: "请输入名称" }],
      },
    };
  },
  methods: {
  },
  created() {
    this.config({
      api: "/system/userGroup/userGroup",
    });
  },
};
</script>
