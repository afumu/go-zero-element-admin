<template>
  <GlobalWindow
    :title="title"
    :visible.sync="visible"
    :confirm-working="isWorking"
    @confirm="confirm"
    @closeTrigger="$refs.form.resetFields()"
  >
    <el-form :model="form" ref="form" :rules="rules">
      <el-form-item label="公司名称" prop="departName" required>
        <el-input v-model="form.departName" placeholder="请输入公司名称" v-trim maxlength="50"/>
      </el-form-item>
      <el-form-item label="类型" prop="orgType" required>
        <el-radio-group v-model="form.orgType">
          <el-radio :label="'1'">公司</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="联系电话" prop="mobile">
        <el-input v-model="form.mobile" placeholder="请输入联系电话" v-trim maxlength="11"/>
      </el-form-item>
      <el-form-item label="地址" prop="address">
        <el-input v-model="form.address" placeholder="请输入公司地址" v-trim maxlength="200"/>
      </el-form-item>
      <el-form-item label="排序" prop="departOrder">
        <el-input v-model="form.departOrder" type="number" placeholder="请输入排序值" v-trim />
      </el-form-item>
    </el-form>
  </GlobalWindow>
</template>

<script>
import BaseOpera from '@/components/base/BaseOpera'
import GlobalWindow from '@/components/common/GlobalWindow'
import { checkMobile, checkEmail } from '@/utils/form'
export default {
  name: 'OperaCompanyWindow',
  extends: BaseOpera,
  components: { GlobalWindow },
  data () {
    return {
      // 需排除选择的部门ID
      excludeDeptId: null,
      // 表单数据
      form: {
        id: null,
        orgType: '1',
        orgCategory: 1,
        departName: '',
        mobile: '',
        departOrder: 0,
        address: null,
      },
      dictOptions:[],
      // 验证规则
      rules: {
        departName: [
          { required: true, message: '请输入公司名称' }
        ],
        phone: [
          { validator: checkMobile }
        ],
        email: [
          { validator: checkEmail }
        ]
      }
    }
  },
  methods: {
    /**
     * @title 窗口标题
     * @target 编辑的部门对象
     * @parent 新建时的上级部门对象
     * @departmentList 部门列表
     */
    open (title, target, parent) {
      this.title = title
      this.visible = true
      // 新建
      if (target == null) {
        this.excludeDeptId = null
        this.$nextTick(() => {
          this.$refs.form.resetFields()
          this.form.id = null
          this.form.parentId = parent == null ? null : parent.id
        })
        return
      }
      // 编辑
      this.$nextTick(() => {
        this.excludeDeptId = target.id
        for (const key in this.form) {
          this.form[key] = target[key]
        }
      })
    }
  },
  created () {
    this.dictOptions=this.$dict.getAllDictItemByCode(this.$consts.dicConstant.orgType)
    this.config({
      api: '/system/department'
    })
  }
}
</script>
