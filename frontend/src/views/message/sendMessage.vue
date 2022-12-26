<!-- 定时任务列表 -->
<template>
  <table-layout>
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
          </el-form>
        </el-collapse-item>
      </el-collapse>
      <!-- <div style="height: 10px"></div> -->
    </div>
    <template v-slot:table-wrap>
      <ul class="toolbar">
        <li>
          <el-button type="primary" @click="$refs.operaMessage.open('新建')" icon="el-icon-plus">新建</el-button>
        </li>
      </ul>
      <el-table  border v-loading="isWorking.search" :data="tableData.list" stripe @sort-change="handleSortChange" @selection-change="handleSelectionChange">
        <el-table-column prop="titile" label="标题" fixed="left" min-width="100px"></el-table-column>
        <el-table-column prop="msgCategory_dictText" label="消息类型" min-width="50px"></el-table-column>
        <el-table-column prop="sender" label="发布人" min-width="50px"></el-table-column>
        <el-table-column prop="priority" label="优先级" min-width="50px">
          <template slot-scope="{ row }">
            {{ row.priority == 'L' ? '低' : row.priority == 'M' ? '中' : row.priority == 'H' ? '高' : ''}}
          </template>
        </el-table-column>
        <el-table-column prop="msgType" label="通告对象" min-width="50px">
          <template slot-scope="{ row }">
            {{ row.msgType == 'USER' ? '指定用户' : row.msgType == 'ALL' ? '全体用户' : ''}}
          </template>
        </el-table-column>
        <el-table-column prop="sendStatus" label="发布状态" min-width="50px">
          <template slot-scope="{ row }">
            {{ row.sendStatus == 0 ? '未发布' : row.sendStatus == 1 ? '已发布' : row.sendStatus == 2 ? '已撤销' : ''}}
          </template>
        </el-table-column>
        <el-table-column prop="sendTime" label="发布时间" min-width="100px"></el-table-column>
        <el-table-column prop="cancelTime" label="撤销时间" min-width="100px"></el-table-column>
        <el-table-column label="操作" min-width="80" fixed="right" align="center">
          <template slot-scope="{ row }">
            <el-button type="text" v-if="row.sendStatus == 0" @click="$refs.operaMessage.open('编辑', row)">编辑</el-button>
            <el-dropdown>
              <span class="el-dropdown-link">
                <el-button type="text" style="padding-right: 0px !important">更多</el-button>
                <i class="el-icon-arrow-down el-icon--right"></i>
              </span>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item v-if="row.sendStatus == 0">
                  <el-button type="text" @click="release(row)">发布</el-button>
                </el-dropdown-item>
                <el-dropdown-item v-if="row.sendStatus == 1">
                  <el-button type="text" @click="reovke(row)">撤销</el-button>
                </el-dropdown-item>
                <el-dropdown-item>
                  <el-button type="text" @click="$refs.operaMessage.open('查看', row, true)" >查看</el-button>
                </el-dropdown-item>
                <el-dropdown-item v-if="row.sendStatus == 0 || row.sendStatus == 2">
                  <el-button type="text" @click="deleteById(row)">删除</el-button>
                </el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </template>
        </el-table-column>
      </el-table>
      <pagination @size-change="handleSizeChange" @current-change="handlePageChange" :pagination="tableData.pagination"></pagination>
    </template>
    <!-- 新建/修改 -->
    <operaMessage ref="operaMessage" @success="handlePageChange(tableData.pagination.pageIndex)" />
  </table-layout>
</template>

<script>
import BaseTable from '@/components/base/BaseTable'
import Pagination from '@/components/common/Pagination'
import TableLayout from '@/layouts/TableLayout'
import operaMessage from './modules/operaMessage'
export default {
  name: 'QuzrtzJobList',
  extends: BaseTable,
  components: {
    TableLayout,
    Pagination,
    operaMessage
  },
  data () {
    return {
      activeNames: '1',
      searchForm: {
        titile: null,
      }
    }
  },
  methods: {
    // 发布
    release (row) {
      const message = '是否发布选中的消息?'
      this.$confirm(message, '确认发布', {
        distinguishCancelAndClose: true,
        confirmButtonText: '确定',
        cancelButtonText: '取消'
      })
        .then(() => {
          this.api
            .release(row.id)
            .then(() => {
              this.$tip.apiSuccess('发布成功')
              this.search()
            })
        })
        .catch(() => {
          this.$tip.apiSuccess('已取消')
        })
    },

    // 撤销
    reovke (row) {
      const message = '是否撤销选中的消息?'
      this.$confirm(message, '确认撤销', {
        distinguishCancelAndClose: true,
        confirmButtonText: '确定',
        cancelButtonText: '取消'
      })
        .then(() => {
          this.api
            .reovke(row.id)
            .then(() => {
              this.$tip.apiSuccess('撤销成功')
              this.search()
            })
        })
        .catch(() => {
          this.$tip.apiSuccess('已取消')
        })
    }
  },
  // 生命周期 - 创建完成（访问当前this实例）
  created () {
    this.config({
      module: '消息发送',
      api: '/message/sendMessage',
      'field.main': 'titile'
    })
    this.search()
  }
}
</script>

<style scoped>
.el-dropdown-link {
  margin-left: 10px;
  cursor: pointer;
  font-size: 12px;
}
</style>
