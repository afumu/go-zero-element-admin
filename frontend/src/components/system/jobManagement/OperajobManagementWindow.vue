<template>
  <GlobalWindow
    :title="title"
    :visible.sync="visible"
    :confirm-working="isWorking"
    @confirm="confirm"
    @closeTrigger="$refs.form.resetFields()"
  >
    <el-form :model="form" ref="form" :rules="rules">
      <el-form-item label="职务编码" prop="code" required>
        <el-input v-model="form.code" placeholder="请输入职务编码" v-trim />
      </el-form-item>
      <el-form-item label="职务名称" prop="name">
        <el-input v-model="form.name" placeholder="请输入职务名称" v-trim />
      </el-form-item>
      <el-form-item label="职级" prop="postRank">
        <el-select clearable v-model="form.postRank"  placeholder="请选择职级" >
            <el-option v-for="item in postRankList" :key="item.text" :label="item.text" :value="item.value"></el-option>
        </el-select>
      </el-form-item>
    </el-form>
  </GlobalWindow>
</template>

<script>
import BaseOpera from '@/components/base/BaseOpera'
import GlobalWindow from '@/components/common/GlobalWindow'
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
        code:'',
        name:'',
        postRank:'',
      },
      postRankList:[],
      dictOptions:[],
      // 验证规则
      rules: {
        code: [
          { required: true, message: '请输入职务编码' }
        ],
        name: [
          { required: true, message:'请输入职务名称' }
        ],
        postRank: [
          { required: true,message:'请选择职级' }
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
    },
  },
  created () {
    this.dictOptions=this.$dict.getAllDictItemByCode(this.$consts.dicConstant.orgType)
    this.config({
      api: '/system/jobManagement'
    })
    this.api.position_rank()
        .then(res =>{
            this.postRankList = res
        })
  }
}
</script>
