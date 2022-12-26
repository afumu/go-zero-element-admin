<template>
  <GlobalWindow
    :title="title"
    :visible.sync="visible"
    :confirm-working="isWorking"
    @confirm="confirm"
    @closeTrigger="$refs.form.resetFields()"
  >
    <el-form :model="form" ref="form" :rules="rules">
      <el-form-item label="上级部门" prop="parentId">
        <PartySelect v-if="visible" ref="partySelect" placeholder="请选择上级党组织" v-model="form.parentId" :exclude-id="excludeDeptId" :inline="false"/>
      </el-form-item>
      <el-form-item label="党组织名称" prop="name" required>
        <el-input v-model="form.name" placeholder="请输入党组织名称" v-trim maxlength="50"/>
      </el-form-item>
      <el-form-item label="党组织类型" prop="type" required>
        <el-select clearable v-model="form.type" placeholder="请选择" >
          <el-option :value="item.value" v-for="item,index in dictOptions" :key="index" :label="item.text">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="联系电话" prop="phone">
        <el-input v-model="form.phone" placeholder="请输入联系电话" v-trim maxlength="11"/>
      </el-form-item>
      <el-form-item label="部门邮箱" prop="email">
        <el-input v-model="form.email" placeholder="请输入部门邮箱" v-trim maxlength="200"/>
      </el-form-item>
    </el-form>
  </GlobalWindow>
</template>

<script>
import BaseOpera from '@/components/base/BaseOpera'
import GlobalWindow from '@/components/common/GlobalWindow'
import PartySelect from '@/components/common/PartySelect'
import { checkMobile, checkEmail } from '@/utils/form'
export default {
  name: 'OperaPartyWindow',
  extends: BaseOpera,
  components: { PartySelect, GlobalWindow },
  data () {
    return {
      // 需排除选择的部门ID
      excludeDeptId: null,
      // 表单数据
      form: {
        id: null,
        parentId: null,
        name: '',
        type: '',
        phone: '',
        email: ''
      },
      dictOptions:[],
      // 验证规则
      rules: {
        name: [
          { required: true, message: '请输入党组织名称' }
        ],
        type: [
          { required: true, message: '请输入党组织类型' }
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
    this.dictOptions=this.$dict.getAllDictItemByCode(this.$consts.dicConstant.partyType)
    this.config({
      api: '/system/party'
    })
  }
}
</script>
