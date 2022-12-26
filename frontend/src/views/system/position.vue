<template>
  <TableLayout :permissions="['system:position:query']">
    <!-- 搜索表单 -->
    <div slot="search-form">
      <el-collapse v-model="activeNames" accordion>
        <el-collapse-item title="查询条件" name="1">
          <template slot="title">
            <div class="kj-title">
              <span style="margin: 0 10px;">查询条件</span>
              <section>
                <el-button type="primary" size="mini" icon="el-icon-search" @click.stop="search">搜索</el-button>
                <el-button @click.stop="reset" size="mini" icon="el-icon-refresh">重置</el-button>
              </section>
            </div>
          </template>
          <el-form ref="searchForm" :model="searchForm" label-width="100px" inline>
            <el-form-item label="职务编码" prop="code">
              <search-input v-model="searchForm.code" v-trim placeholder="请输入职务编码" @keypress.enter.native="search" />
            </el-form-item>
            <el-form-item label="职务名称" prop="name">
              <search-input v-model="searchForm.name" v-trim placeholder="请输入职务名称" @keypress.enter.native="search" />
            </el-form-item>
          </el-form>
        </el-collapse-item>
      </el-collapse>
    </div>
    <!-- 表格和分页 -->
    <template v-slot:table-wrap>
      <ul class="toolbar">
        <li>
          <el-button type="primary" @click="$refs.operaPositionWindow.open('新建')" icon="el-icon-plus">新建</el-button>
        </li>
        <li>
          <el-button @click="deleteByIdInBatch" icon="el-icon-delete">删除</el-button>
        </li>
      </ul>
      <el-table border v-loading="isWorking.search" :data="tableData.list" row-key="id" stripe default-expand-all @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" fixed="left"></el-table-column>
        <!-- <el-table-column label="#" fixed="left" min-width="60px">
          <template slot-scope="scope">
            <label>{{scope.$index+1}}</label>
          </template>
        </el-table-column> -->
        <el-table-column prop="code" label="职务编码" fixed="left" min-width="150px"></el-table-column>
        <el-table-column prop="name" label="职务名称" min-width="120px"></el-table-column>
        <el-table-column prop="postRank_dictText" label="职级" min-width="100px"></el-table-column>
        <el-table-column label="操作" min-width="100" fixed="right" align="center">
          <template slot-scope="{row}">
            <el-button style="margin-left: 0;!important" type="text" @click="$refs.operaPositionWindow.open('修改', row)" icon="el-icon-edit">编辑</el-button>
            <el-button style="margin-left: -15px;!important" type="text" @click="deleteById(row)" icon="el-icon-delete">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </template>
    <!-- 新建/修改 -->
    <OperaPositionWindow  ref="operaPositionWindow" @success="handlePageChange" />
  </TableLayout>
</template>

<script>
import TableLayout from "@/layouts/TableLayout";
import BaseTable from "@/components/base/BaseTable";
import OperaPositionWindow from "@/components/system/position/OperaPositionWindow";
export default {
  name: "SystemPosition",
  extends: BaseTable,
  components: { OperaPositionWindow, TableLayout },
  data() {
    return {
      // 搜索
      searchForm: {},
      activeNames: "1",
    };
  },
  methods: {},
  created() {
    this.config({
      module: "职位",
      api: "/system/position",
      "field.main": "name",
    });
    this.search();
  },
};
</script>
<style lang="scss" scoped>
.table-layout {
  /deep/ .table-content {
    margin-top: 0;
  }
}
</style>
