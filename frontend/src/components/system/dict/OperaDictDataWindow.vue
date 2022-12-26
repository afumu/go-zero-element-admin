<template>
  <GlobalWindow
    :title="title"
    :visible.sync="visible"
    :confirm-working="isWorking.create"
    @confirm="confirm"
    @closeTrigger="$refs.form.resetFields()"
  >
    <el-form :model="form" ref="form" :rules="rules" class="content">
      <el-form-item label="名称" prop="itemText" required>
        <el-input v-model="form.itemText" placeholder="请输入名称" v-trim maxlength="50"/>
      </el-form-item>
      <el-form-item label="数据值" prop="itemValue" required>
        <el-input v-model="form.itemValue" placeholder="请输入数据值" v-trim maxlength="50"/>
      </el-form-item>
      <el-form-item label="描述" prop="description">
        <el-input v-model="form.description" placeholder="请输入描述信息" v-trim maxlength="50"/>
      </el-form-item>
      <el-form-item label="状态" prop="status" required class="form-item-status">
        <el-switch class="switch" v-model="form.status" :value="1" :active-value="1" :inactive-value="0"
        active-text = "启用" inactive-text="禁用">
        </el-switch>
      </el-form-item>
    </el-form>
  </GlobalWindow>
</template>

<script>
import BaseOpera from '@/components/base/BaseOpera'
import GlobalWindow from '@/components/common/GlobalWindow'
export default {
  name: 'OperaDictDataWindow',
  extends: BaseOpera,
  components: { GlobalWindow },
  data () {
    return {
      visible: false,
      // 表单数据
      form: {
        id: null,
        dictId: '',
        itemValue: '',
        itemText: '',
        description: '',
        status: 1,
        sortOrder: ""
      },
      // 验证规则
      rules: {
        itemText: [
          { required: true, message: '请输入名称' }
        ],
        itemValue: [
          { required: true, message: '请输入数据值' }
        ]
      }
    }
  },
  methods: {
    // confirm(){
    //   console.log(this.form)
    // },
    /**
     * @title 窗口标题
     * @dict 所属字典ID
     * @target 编辑的字典数据对象
     */
    open (title, dictId, target) {
      this.title = title
      this.visible = true
      // 新建
      if (target == null) {
        this.$nextTick(() => {
          this.$refs.form.resetFields()
          this.form.id = null
          this.form.dictId = dictId
        })
        return
      }
      // 编辑
      this.$nextTick(() => {
        for (const key in this.form) {
          this.form[key] = target[key]
        }
      })
      // console.log(this.form)
    }
  },
  created () {
    this.config({
      api: '/system/dictItem'
    })
  }
}
</script>

<style lang="scss" scoped>
/* switch按钮样式 */
.switch .el-switch__label {
    position: absolute;
    display: none;
    color: #fff !important;
}
/*打开时文字位置设置*/
.switch .el-switch__label--right {
    z-index: 1;
}
/* 调整打开时文字的显示位子 */
.switch .el-switch__label--right span{
    margin-right: 5px;
}
/*关闭时文字位置设置*/
.switch .el-switch__label--left {
    z-index: 1;
}
/* 调整关闭时文字的显示位子 */
.switch .el-switch__label--left span{
    margin-left: 20px;
}
/*显示文字*/
.switch .el-switch__label.is-active {
    display: block;
}
/* 调整按钮的宽度 */
.switch.el-switch .el-switch__core,
.el-switch .el-switch__label {
     width: 55px !important;
     margin: 0;
}
</style>
