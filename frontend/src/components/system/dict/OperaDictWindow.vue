<template>
  <GlobalWindow
    :title="title"
    :visible.sync="visible"
    :confirm-working="isWorking"
    @confirm="confirm"
    @closeTrigger="$refs.form.resetFields()"
  >
    <el-form :model="form" ref="form" :rules="rules">
      <el-form-item label="字典名称" prop="dictName" >
        <el-input v-model="form.dictName" placeholder="请输入字典名称" v-trim maxlength="50"/>
      </el-form-item>
      <el-form-item label="字典编码" prop="dictCode" >
        <el-input v-model="form.dictCode" placeholder="请输入字典编码" v-trim maxlength="50"/>
      </el-form-item>
      <el-form-item label="描述" prop="description">
        <el-input v-model="form.description" type="textarea" placeholder="请输入描述信息" :rows="3" v-trim maxlength="500"/>
      </el-form-item>
    </el-form>
  </GlobalWindow>
</template>

<script>
import BaseOpera from '@/components/base/BaseOpera'
import GlobalWindow from '@/components/common/GlobalWindow'
export default {
  name: 'OperaDictWindow',
  extends: BaseOpera,
  components: { GlobalWindow },
  data () {
    return {
      // 表单数据
      form: {
        id: null,
        dictCode: '',
        dictName: '',
        description: ''
      },
      // 验证规则
      rules: {
        dictCode: [
          { required: true, message: '请输入字典编码' }
        ],
        dictName: [
          { required: true, message: '请输入字典名称' }
        ],
        description: [
          { required: true, message: '请输入描述信息' }
        ]
      }
    }
  },
  methods: {
    open (title, target) {
      this.title = title
      this.visible = true
      // 新建
      if (target == null) {
        this.$nextTick(() => {
          this.$refs.form.resetFields()
          this.form[this.configData['field.id']] = undefined
        })
        return
      }
      // 编辑
      this.$nextTick(() => {
        this.originRoleCode = target.code
        for (const key in this.form) {
          this.form[key] = target[key]
        }
      })
    }
  },
  created () {
    this.config({
      api: '/system/dict'
    })
  }
}
</script>
