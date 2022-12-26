<template>
  <GlobalWindow :title="title" :visible.sync="visible" :confirm-working="isWorking" @confirm="confirm">
    <el-form :model="form" ref="form" :rules="rules">
      <div>
        <table>
          <tr>
            <td>
              <el-form-item label="首页轮播图" required>
                <el-input style="width:280px;margin:10px 0px 10px 0px" :placeholder='resMessage' v-model="resMessage" :disabled="true"></el-input>
              </el-form-item>
            </td>
            <td style="margin-left:50px">
              <el-upload :action="this.$consts.fileConstant.uploadTemplateFilePath" :headers="importHeaders" :show-file-list="false" :on-success="uploadSuccess">
                <el-button size="small" type="primary" style="height: 28px;margin:10px 0px 0px 0px">
                  点击上传
                </el-button>
              </el-upload>
            </td>
          </tr>
          <tr>
            <td>
              <template>
                <el-radio v-model="form.enable" label="1">开</el-radio>
                <el-radio v-model="form.enable" label="0">关</el-radio>
              </template>
            </td>
          </tr>
          <tr>
            <td>
              <el-form-item label="排序" prop="sort" required>
                <template>
                  <el-input v-model="form.sort" placeholder="请输入排序"></el-input>
                </template>
              </el-form-item>
            </td>
          </tr>
          <tr>
            <td>
              <el-form-item prop="memo" required label="备注">
                <template>
                  <el-input type="textarea" label="备注" :rows="2" placeholder="请输入内容" v-model="form.memo">
                  </el-input>
                </template>
              </el-form-item>
            </td>
          </tr>
        </table>
      </div>
    </el-form>
  </GlobalWindow>
</template>

<script>
import BaseOpera from "@/components/base/BaseOpera";
import GlobalWindow from "@/components/common/GlobalWindow";
import { mapState } from "@/store/index.js";
import Cookies from "js-cookie";

export default {
  name: "OperaScrollWindow",
  extends: BaseOpera,
  components: { GlobalWindow },
  data() {
    return {
      // 表单数据
      resMessage: "请上传轮播图",
      importHeaders: { "X-Access-Token": Cookies.get("X-Access-Token") },
      form: {
        id: null,
        fileId: null,
        enable: "1",
        memo: "",
        sort: "",
      },
      // 验证规则
      rules: {
        img: [{ required: true, message: "请上传文件" }],
        sort: [{ required: true, message: "请输入排序值" }],
        memo: [{ required: true, message: "请输入备注" }],
      },
    };
  },
  computed: {},
  methods: {
    open(title, target) {
      this.title = title;
      this.visible = true;
      this.resMessage = "请上传轮播图";
      this.form.fileId = null;
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
    uploadSuccess(res, file) {
      this.resMessage = file.name;
      this.form.fileId = res.message;
    },
  },
  created() {
    this.config({
      api: "/system/scroll",
    });
  },
};
</script>
<style scoped lang="scss">
</style>