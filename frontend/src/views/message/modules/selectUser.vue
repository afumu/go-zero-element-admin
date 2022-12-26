<template>
  <GlobalWindow class="menu-config-dialog" :visible.sync="visible" width="776px" title="用户" @confirm="confirm">
    <TableLayout :withBreadcrumb="false">
      <div slot="search-form">
        <el-collapse v-model="activeNames" accordion>
          <el-collapse-item title="查询条件" name="1">
            <template slot="title">
              <div class="kj-title">
                <span style="margin: 0 10px">查询条件</span>
                <section>
                  <el-button type="primary" size="mini" icon="el-icon-search" @click.stop="search">搜索</el-button>
                  <el-button @click.stop="reset" size="mini" icon="el-icon-search">重置</el-button>
                </section>
              </div>
            </template>
            <el-form ref="searchForm" :model="searchForm" label-width="100px" inline>
              <el-form-item label="用户账号" prop="username">
                <searchInput v-model="searchForm.username" placeholder="请输入用户账号"></searchInput>
              </el-form-item>
            </el-form>
          </el-collapse-item>
        </el-collapse>
      </div>
      <template v-slot:table-wrap>
        <el-table border v-loading="isWorking.search" :data="tableData.list" ref="userLists" stripe @selection-change="handleSelectionChange">
          <el-table-column type="selection" align="center" fixed="left" width="55"></el-table-column>
          <el-table-column prop="username" align="center" label="用户账号" fixed="left" min-width="100px"></el-table-column>
          <el-table-column prop="realname" align="center" label="用户名称" fixed="left" min-width="100px"></el-table-column>
          <el-table-column prop="status_dictText" align="center" label="状态" min-width="60px"></el-table-column>
        </el-table>
        <pagination @size-change="handleSizeChange" @current-change="handlePageChange" :pagination="tableData.pagination"></pagination>
      </template>
    </TableLayout>
  </GlobalWindow>
</template>
<script>
import TableLayout from '@/layouts/TableLayout'
import GlobalWindow from '@/components/common/GlobalWindow'
import searchInput from '@/components/common/searchInput'
import BaseTable from '@/components/base/BaseTable'
import Pagination from '@/components/common/Pagination'

export default {
  extends: BaseTable,
  components: { GlobalWindow, searchInput, TableLayout, Pagination },
  data () {
    return {
      visible: false,
      // isWorking: false,
      searchForm: {
        username: ''
      },
      activeNames: '1',
      users: []
    }
  },

  methods: {
    open (users) {
      this.users = users
      this.visible = true
    },
    confirm () {
      this.visible = false
      const data = this.tableData.selectedRows
      this.$emit('success', data)
    },
    selectUser () {
      const users = this.users
      if (users) {
        for (let i = 0; i < users.length; i++) {
          for (let j = 0; j < this.tableData.list.length; j++) {
            if (users[i].value === this.tableData.list[j].id) {
              this.$refs.userLists.toggleRowSelection(this.tableData.list[j])
            }
          }
        }
      }
    }
  },
  updated () {
    if (this.$refs.userLists) {
      this.selectUser()
    }
  },
  created () {
    this.config({
      module: '消息发送-选择用户',
      api: '/message/selectUser'
    })
    this.search()
  }
}
</script>
