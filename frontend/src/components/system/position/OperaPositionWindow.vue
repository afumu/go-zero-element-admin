<template>
  <GlobalWindow :title="title" :visible.sync="visible" :confirm-working="isWorking" @confirm="confirm" @closeTrigger="$refs.form.resetFields()">
    <el-form :model="form" ref="form" :rules="rules">
      <el-form-item label="岗位编码" prop="code" required>
        <el-input v-model="form.code" placeholder="请输入岗位编码" v-trim maxlength="50" />
      </el-form-item>
      <el-form-item label="岗位名称" prop="name" required>
        <el-input v-model="form.name" placeholder="请输入岗位名称" v-trim maxlength="50" />
      </el-form-item>
      <el-form-item label="职级" prop="postRank" required>
        <el-select clearable v-model="form.postRank" placeholder="请选择">
          <el-option :value="item.value" v-for="item,index in dictOptions" :key="index" :label="item.text">
          </el-option>
        </el-select>
      </el-form-item>
    </el-form>
  </GlobalWindow>
</template>

<script>
import BaseOpera from "@/components/base/BaseOpera";
import GlobalWindow from "@/components/common/GlobalWindow";
export default {
  name: "OperaPositionWindow",
  extends: BaseOpera,
  components: { GlobalWindow },
  data() {
    return {
      // 需排除选择的岗位ID
      excludePositionId: null,
      // 表单数据
      form: {
        id: null,
        code: "",
        name: "",
        postRank: "",
      },
      // 验证规则
      rules: {
        code: [{ required: true, message: "请输入岗位编码" }],
        name: [{ required: true, message: "请输入岗位名称" }],
        postRank: [{ required: true, message: "请选择职级" }],
      },
      //字典数据
      dictOptions: [],
    };
  },
  methods: {
    /**
     * @title 窗口标题
     * @target 编辑的岗位对象
     * @parent 新建时的上级岗位对象
     * @positionList 岗位列表
     */
    open(title, target, parent) {
      this.title = title;
      this.visible = true;
      this.form.id = "";
      this.form.code = "";
      this.form.name = "";
      this.form.postRank = "";

      // 新建
      if (target == null) {
        this.excludePositionId = null;
        this.$nextTick(() => {
          this.$refs.form.resetFields();
          this.form.id = null;
          this.form.parentId = parent == null ? null : parent.id;
        });
        return;
      }
      // 编辑
      this.$nextTick(() => {
        this.excludePositionId = target.id;
        for (const key in this.form) {
          this.form[key] = target[key];
        }
      });
    },
  },
  // 确认新建
  __confirmCreate () {
    this.$refs.form.validate((valid) => {
      if (!valid) {
        return
      }
      // 调用新建接口
      this.isWorking = true
      this.api.create(this.form)
          .then(() => {
            this.visible = false
            this.$tip.apiSuccess('新建成功')
            this.$emit('success')
          })
          .catch(e => {
            this.$tip.apiFailed("编码重复")
            this.$emit('fail',e)
          })
          .finally(() => {
            this.isWorking = false
          })
    })
  },
  created() {
    this.dictOptions = this.$dict.getAllDictItemByCode(
      this.$consts.dicConstant.positionRank
    );
    this.config({
      api: "/system/position",
    });
  },
};
</script>
