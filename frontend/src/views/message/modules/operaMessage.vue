<template>
  <GlobalWindow :title="title" :visible.sync="visible" :confirm-working="isWorking" @confirm="confirm" @closeTrigger="$refs.form.resetFields()" :withFooter="!isEdit">
    <el-form :model="form" ref="form" :rules="rules">
      <el-form-item label="标题" prop="titile">
        <el-input v-model="form.titile" placeholder="请输入标题" v-trim maxlength="50" />
      </el-form-item>
      <el-form-item label="消息类型" prop="msgCategory">
        <el-select v-model="form.msgCategory" placeholder="请选择消息类型">
          <el-option v-for="item in msgCategorys" :key="item.value" :label="item.text" :value="item.value" />
        </el-select>
      </el-form-item>
      <el-form-item label="开始时间" prop="startTime">
        <el-date-picker v-model="form.startTime" type="datetime" @change="changeStart($event)" value-format="yyyy-MM-dd HH:mm:ss" format="yyyy-MM-dd HH:mm:ss" placeholder="请选择开始时间" />
      </el-form-item>
      <el-form-item label="结束时间" prop="endTime">
        <el-date-picker v-model="form.endTime" type="datetime" :picker-options="pickerOptions" value-format="yyyy-MM-dd HH:mm:ss" format="yyyy-MM-dd HH:mm:ss" placeholder="请选择结束时间" :disabled="disable" />
      </el-form-item>
      <el-form-item label="优先级" prop="priority">
        <el-select v-model="form.priority" placeholder="请选择">
          <el-option label="低" value="L" />
          <el-option label="中" value="M" />
          <el-option label="高" value="H" />
        </el-select>
      </el-form-item>
      <el-form-item label="通告类型" prop="msgType">
        <el-select v-model="form.msgType" placeholder="请选择">
          <el-option label="指定用户" value="USER" />
          <el-option label="全体用户" value="ALL" />
        </el-select>
      </el-form-item>
      <el-form-item label="指定用户" prop="msgType" v-if="form.msgType=='USER'">
        <el-input v-model="selectUsersText"  @click.native="$refs.selectUser.open(selectUsers)"></el-input>
        <!-- <el-select v-model="selectUsers" multiple value-key="id" placeholder="请选择用户" @click.native="$refs.selectUser.open(selectUsers)">
          <el-option v-for="item in selectUsers" :key="item.value" :label="item.label" :value="item.value"></el-option>
        </el-select> -->
      </el-form-item>
      <el-form-item label="摘要" prop="msgAbstract">
        <el-input v-model="form.msgAbstract" type="textarea" placeholder="请输入摘要" :rows="3" v-trim maxlength="500" />
      </el-form-item>
      <el-form-item label="内容" prop="msgContent">
        <quill-editor ref="text" v-model="form.msgContent" @focus="onEditorFocus($event)" :options="editorOption"></quill-editor>
      </el-form-item>
    </el-form>
    <selectUser ref="selectUser" @success="handleSelectChange" />
  </GlobalWindow>
</template>

<script>
import BaseOpera from '@/components/base/BaseOpera'
import GlobalWindow from '@/components/common/GlobalWindow'
import selectUser from './selectUser'
import { quillEditor } from 'vue-quill-editor'
import * as Quill from 'quill'
var fonts = [
  'SimSun',
  'NSimSun',
  'SimHei',
  'Microsoft-YaHei',
  'KaiTi',
  'FangSong',
  'Arial',
  'Times-New-Roman',
  'sans-serif'
]
var Font = Quill.import('formats/font')
Font.whitelist = fonts // 将字体加入到白名单
Quill.register(Font, true)

var sizes = [false, '16px', '18px', '20px', '22px', '26px', '28px', '30px']
var Size = Quill.import('formats/size')
Size.whitelist = sizes
export default {
  name: 'operaMessage',
  extends: BaseOpera,
  components: { GlobalWindow, selectUser, quillEditor },
  data () {
    return {
      title: null,
      visible: false,
      disable: true,
      // 表单数据
      form: {
        id: null,
        titile: '',
        msgCategory: '',
        startTime: null,
        endTime: null,
        priority: '',
        msgType: '',
        msgAbstract: '',
        msgContent: ''
      },
      msgCategorys: this.$dict.getAllDictItemByCode(
        this.$consts.dicConstant.msg_category
      ),
      selectUsers: [],
      selectUsersText: '',
      // 验证规则
      rules: {
        titile: [{ required: true, message: '请输入标题' }],
        msgCategory: [{ required: true, message: '请选择消息类型' }],
        startTime: [{ required: true, message: '请选择开始时间' }],
        endTime: [{ required: true, message: '请选择结束时间' }],
        msgType: [{ required: true, message: '请选择通告类型' }],
        msgAbstract: [{ required: true, message: '请输入摘要' }],
        msgContent: [{ required: true, message: '请输入摘要' }]
      },
      isEdit: false,
      // 富文本编辑器 参数
      editorOption: {
        placeholder: '请输入内容',
        modules: {
          toolbar: [
            ['bold', 'italic', 'underline', 'strike'], // toggled buttons
            ['blockquote', 'code-block'],
            [{ header: [1, 2, 3, 4, 5, 6, false] }],
            [{ size: sizes }], // 配置字号
            [{ color: [] }, { background: [] }], // 显示背景字体颜色
            [{ font: fonts }], // 显示字体选择
            [{ align: [] }], // 显示居中
            ['image', 'video']
          ]
        }
      },
      beginTime: Date.now(),
      pickerOptions: this.endPickerOptions(),
    }
  },
  methods: {
    open (title, target, isEdit) {
      this.title = title
      this.visible = true
      this.isEdit = isEdit
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
      if (target.userIds) {
        this.searchUser(target.userIds)
      }
    },
    onEditorFocus (event) {
      if (this.isEdit) {
        event.enable(false)
      } else {
        event.enable(true)
      }
    },
    handleSelectChange (users) {
      this.selectUsers = []
      const realName = []
      users.forEach((e) => {
        this.fullUser(e)
        realName.push(e.realname)
      })
      this.selectUsersText = realName.join(',')
    },
    fullUser (e) {
      const user = {
        value: e.id,
        label: e.username
      }
      this.selectUsers.push(user)
    },
    searchUser (userIds) {
      this.api.queryByIds(userIds).then((res) => {
        this.handleSelectChange(res)
      })
    },
    confirm () {
      if (this.selectUsers != []) {
        this.form.userIds = this.selectUsers.map((row) => row.value).join(',')
      }
      if (this.form.id == null || this.form.id === '') {
        this.__confirmCreate()
        return
      }
      this.__confirmEdit()
    },
    changeStart(event) {
      this.beginTime = new Date(event).getTime()
      this.disable = false
    },
    endPickerOptions(){
      let _this = this
      return {
        disabledDate(time) {
          return time.getTime() < _this.beginTime; //禁止选择以前的时间
        },
      } 
    },
  },
  created () {
    this.config({
      api: '/message/sendMessage'
    })
  }
}
</script>
