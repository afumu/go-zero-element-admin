<template>
  <GlobalWindow
    :title="title"
    width="40%"
    @confirm="confirm"
    :visible.sync="visible"
  >
    <TableLayout :with-breadcrumb="false">
      <!-- 表格和分页 -->
      <template v-slot:table-wrap>
        <ul class="toolbar">
          <li><el-button type="primary" @click="$refs.operaDictDataWindow.open('新增', searchForm.dictId)" icon="el-icon-plus">新建</el-button></li>
        </ul>
        <el-table
            border
          v-loading="isWorking.search"
          :data="tableData.list"
          stripe
        >
          <el-table-column prop="itemText" label="名称" min-width="80px"></el-table-column>
          <el-table-column prop="itemValue" label="数据值" min-width="50px"></el-table-column>
          <el-table-column
            label="操作"
            min-width="60px"
            fixed="right"
            align="center"
          >
            <template slot-scope="{row}">
              <el-button type="text" @click="$refs.operaDictDataWindow.open('修改', searchForm.dictId, row)" icon="el-icon-edit">编辑</el-button>
              <el-button type="text" @click="deleteById(row)" icon="el-icon-delete">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <pagination
          @size-change="handleSizeChange"
          @current-change="handlePageChange"
          :pagination="tableData.pagination"
        ></pagination>
      </template>
      <!-- 新建/修改 -->
      <OperaDictDataWindow ref="operaDictDataWindow" @success="handlePageChange(tableData.pagination.pageIndex)"/>
    </TableLayout>
  </GlobalWindow>
</template>

<script>
import BaseTable from '@/components/base/BaseTable'
import Pagination from '@/components/common/Pagination'
import GlobalWindow from '@/components/common/GlobalWindow'
import TableLayout from '@/layouts/TableLayout'
import OperaDictDataWindow from './OperaDictDataWindow'
export default {
  name: 'DictDataManagerWindow',
  extends: BaseTable,
  components: { OperaDictDataWindow, TableLayout, GlobalWindow, Pagination },
  data () {
    return {
      visible: false,
      title:'',
      searchForm: {
        // 字典ID
        dictId: null
      },
      configData: {
        'field.id': 'id',
        'field.main': 'itemText'
      }
    }
  },
  methods: {
    confirm(){
      this.visible=false
    },
    // 打开数据管理
    open (title,dictId) {
      this.title = title
      this.searchForm.dictId = dictId
      this.visible = true
      this.search()
    }
  },
  created () {
    this.config({
      api: '/system/dictItem',
      'field.main': 'itemText'
    })
  }
}
</script>

<style scoped lang="scss">
/deep/ .window__body {
  .table-content {
    padding: 0;
    .table-wrap {
      padding-top: 0;
    }
  }
}
</style>
