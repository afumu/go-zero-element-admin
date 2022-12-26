<template>
  <GlobalWindow :title="title" :visible.sync="visible" :confirm-working="isWorking" @confirm="confirm" @closeTrigger="$refs.form.resetFields()">
    <el-form :model="form" ref="form" :rules="rules" class="outer-form">
      <el-form-item label="任务类名" prop="jobClassName" required class="item-label">
        <el-input v-model="form.jobClassName" placeholder="请输入任务类名" class="inline-input" maxlength="50" />
      </el-form-item>
      <el-form-item label="cron表达式" prop="cronExpression" class="item-label" required>
        <!-- vue-cron 组件 -->
        <el-popover v-model="cronPopover">
          <el-input slot="reference" @click="cronPopover=true" v-model="form.cronExpression" placeholder="请输入定时策略" />
          <cron @change="setCorn" @close="cronPopover=false" i18n="cn"></cron>
        </el-popover>

      </el-form-item>
      <el-form-item label="参数" prop="parameter" class="item-label">
        <el-input v-model="form.parameter" placeholder="请输入参数" :rows="3" type="textarea" class="inline-input" v-trim maxlength="50" />
      </el-form-item>
      <el-form-item label="描述" prop="description" class="item-label">
        <el-input v-model="form.description" placeholder="请输入描述信息" :rows="5" type="textarea" class="inline-input" v-trim maxlength="50" />
      </el-form-item>
      <el-form-item label="状态" class="item-label" prop="status">
        <el-radio-group v-model="form.status">
          <el-radio-button label="-1">停止</el-radio-button>
          <el-radio-button label="0">正常</el-radio-button>
        </el-radio-group>
      </el-form-item>
    </el-form>
  </GlobalWindow>
</template>

<script>
import BaseOpera from "@/components/base/BaseOpera";
import GlobalWindow from "@/components/common/GlobalWindow";
import { cron } from "vue-cron";
export default {
  name: "OperaQuartzJobWindow",
  extends: BaseOpera,
  components: { GlobalWindow, cron },
  data() {
    return {
      cronPopover: false,
      visible: false,
      title: "",
      // 表单数据
      form: {
        id: null,
        jobClassName: "",
        cronExpression: "* * * * * ? *",
        parameter: "",
        description: "",
        status: 0,
      },
      // dictOptions:[0,1],
      // 验证规则
      rules: {
        jobClassName: [{ required: true, message: "请输入任务类名" }],
        cronExpression: [{ required: true, message: "请输入cron表达式" }],
      },
    };
  },
  methods: {
    open(title, target) {
      this.title = title;
      this.visible = true;
      // 新建
      if (target == null) {
        this.$nextTick(() => {
          this.$refs.form.resetFields();
          this.form[this.configData["field.id"]] = null;
        });
        return;
      }
      // 编辑
      this.$nextTick(() => {
        for (const key in this.form) {
          this.form[key] = target[key];
        }
      });
    },
    confirm() {
      if (this.form.id == null || this.form.id === "") {
        this.__confirmCreate();
      } else {
        this.__confirmEdit();
      }
    },
    setCorn(val) {
      this.form.cronExpression = val;
    },
  },
  created() {
    // 定时任务状态 从字典获得
    // this.dictOptions = this.$dict.getAllDictItemByCode(this.$consts.dicConstant.companyProperty)
    this.config({
      api: "/system/quartzJob",
    });
  },
};
</script>
<style lang="scss" scoped>
.outer-form {
  min-width: 500px;
  padding: 0 24px;
  .item-label > label {
    text-align: right;
    width: 25%;
  }
}
.inline-input {
  position: relative;
  box-sizing: border-box;
  min-width: 300px;
  width: 450px;
}
</style>
