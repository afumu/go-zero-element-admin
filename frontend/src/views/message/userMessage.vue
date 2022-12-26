<template>
  <div>
    <TableLayout>
      <!-- 搜索表单 -->
      <div slot="search-form">
        <el-collapse v-model="activeNames" accordion>
          <el-collapse-item title="查询条件" name="1">
            <template slot="title">
              <div class="kj-title">
                <span style="margin: 0 10px">查询条件</span>
                <section>
                  <el-button type="primary" size="mini" icon="el-icon-search" @click.stop="search">搜索</el-button>
                  <el-button @click.stop="reset" size="mini" icon="el-icon-refresh">重置</el-button>
                </section>
              </div>
            </template>
            <el-form ref="searchForm" :model="searchForm" label-width="100px" inline>
              <el-form-item label="标题" prop="titile">
                <el-input v-model="searchForm.titile" v-trim placeholder="请输入标题" @keypress.enter.native="search" />
              </el-form-item>
              <el-form-item label="发布人" prop="sender">
                <el-input v-model="searchForm.sender" v-trim placeholder="请输入发布人" @keypress.enter.native="search" />
              </el-form-item>
            </el-form>
          </el-collapse-item>
        </el-collapse>
      </div>
      <!-- 表格和分页 -->
      <template v-slot:table-wrap>
        <ul class="toolbar">
          <li>
            <el-button type="primary" @click="handleReadAll" icon="el-icon-circle-check">全部已读</el-button>
          </li>
        </ul>
        <el-table border v-loading="isWorking.search" :data="tableData.list" :default-sort="{ prop: 'createTime', order: 'descending' }" stripe @sort-change="handleSortChange">
          <el-table-column prop="titile" label="标题" fixed="left" min-width="100px"></el-table-column>
          <el-table-column prop="sender" label="发布人" fixed="left" min-width="50px"></el-table-column>
          <el-table-column prop="msgCategory" label="消息类型" fixed="left" min-width="50px">
            <template slot-scope="{ row }">
              {{ row.msgCategory == 1 ? '通知公告' : '系统消息'}}
            </template>
          </el-table-column>
          <el-table-column prop="sendTime" label="发布时间" fixed="left" min-width="100px"></el-table-column>
          <el-table-column prop="priority" label="优先级" fixed="left" min-width="50px">
            <template slot-scope="{ row }">
              {{ row.priority === 'L' ? '低' : row.priority === 'M' ? '中' : row.priority === 'H' ? '高' : ''}}
            </template>
          </el-table-column>
          <el-table-column prop="readFlag" label="阅读状态" fixed="left" min-width="50px">
            <template slot-scope="{ row }">
              {{ row.readFlag == 1 ? '已读' : '未读'}}
              </template>
          </el-table-column>
          <el-table-column label="操作" min-width="50px" fixed="right" align="center">
            <template slot-scope="{ row }">
              <el-button style="margin-left: -15px;!important" type="text" @click="showAnnouncement(row)" icon="el-icon-view">查看</el-button>
            </template>
          </el-table-column>
        </el-table>
        <pagination @size-change="handleSizeChange" @current-change="handlePageChange" :pagination="tableData.pagination"></pagination>
      </template>
    </TableLayout>
    <el-dialog
      title="通知消息"
      :visible.sync="visible"
      width="50%">
      <h2 style="color: rgba(0,0,0,.85);">{{message.titile}}</h2>
      <p style="color: rgba(0,0,0,.45);">发布人：{{message.sender}}  发布时间：{{message.sendTime}}</p>
      <el-divider></el-divider>

      <div v-html="message.msgContent" class="ql-editor" style="padding: 12px 0px"></div>

      <span slot="footer" class="dialog-footer">
        <el-button @click="visible = false">关 闭</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import Pagination from '@/components/common/Pagination'
import TableLayout from '@/layouts/TableLayout'
import BaseTable from '@/components/base/BaseTable'
import { readAll } from '@/api/message/usermessage'

export default {
  name: 'UserMessage',
  extends: BaseTable,
  components: {
    TableLayout,
    Pagination
  },
  data () {
    return {
      // 搜索
      searchForm: {
        sender: null,
        titile: null
      },
      activeNames: '1',
      url: {
        list: '/business/notice/list',
        editCementSend: '/business/notice/updateNoticeDetail',
        readAllMsg: '/business/notice/readAll'
      },
      visible: false,
      messInfo: '',
      category_MultiString: [],
      message: {}
    }
  },
  watch: {
    category_MultiString (val) {
      if (val) {
        this.searchForm.category_MultiString = val.join(',')
      } else {
        this.searchForm.category_MultiString = null
      }
    }
  },
  methods: {
    handleReadAll () {
      this.$confirm('是否全部标注已读?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          this.__checkApi()
          readAll()
            .then(() => {
              this.$tip.apiSuccess('操作成功')
              this.search()
            })
            .catch(() => {
              this.$tip.apiSuccess('操作失败')
              this.search()
            })
        })
        .catch((e) => {
          this.$message({
            type: 'info',
            message: e
          })
        })
    },
    showAnnouncement (record) {
      this.message = record
      console.log(record)
      if(record.readFlag == 0 ){
        this.api.readNotic(record.anntId)
        .then(()=>{
          this.search()
        })
      }
      this.visible = true
    }
  },
  created () {
    this.config({
      module: '我的消息',
      api: '/message/usermessage',
      'field.main': 'title',
    })

    const thisRouter = this.$route.query.category
    if (thisRouter) {
      this.category_MultiString = thisRouter.split(',')
      this.searchForm.readFlag = 0
      this.search()
    } else {
      this.search()
    }
  }
}
</script>
