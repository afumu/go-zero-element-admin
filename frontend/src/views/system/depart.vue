<template>
  <TableLayout :permissions="['system:department:query']">
    <!-- 表格和分页 -->
    <template v-slot:table-wrap>
      <ul class="toolbar" >
        <li><el-button type="primary" @click="$refs.OperaCompanyWindow.open('新建')" icon="el-icon-plus">新建</el-button></li>
      </ul>
      <el-table border v-loading="isWorking.search" :data="tableData.list" :tree-props="{children: 'children', hasChildren: 'hasChildren'}"
      row-key="id" stripe default-expand-all @selection-change="handleSelectionChange">
        <el-table-column type="selection" fixed="left" width="55"></el-table-column>
        <el-table-column prop="departName" label="部门名称" fixed="left" min-width="160px"></el-table-column>
        <el-table-column prop="orgCode" label="部门编码" fixed="left" min-width="100px"></el-table-column>
        <el-table-column prop="orgType" label="机构类型" fixed="left" min-width="100px">
          <template slot-scope="{row}">{{ getText(row)}}</template>
        </el-table-column>
        <el-table-column prop="mobile" label="联系电话" fixed="left" min-width="100px" />
        <el-table-column prop="address" label="地址" fixed="left" min-width="100px" />
        <el-table-column label="操作" align="center" min-width="160" fixed="right">
          <template slot-scope="{row}">
            <el-button  v-if="row.orgType==1" type="text" @click="$refs.OperaCompanyWindow.open('编辑公司', row)" icon="el-icon-edit">编辑</el-button>
            <el-button  v-else type="text" @click="$refs.operaDepartmentWindow.open('编辑部门', row)" icon="el-icon-edit">编辑</el-button>
            <el-button  type="text" @click="$refs.operaDepartmentWindow.open('添加下级', null, row)" icon="el-icon-edit">添加下级</el-button>
            <el-button  type="text" @click="deleteById(row)" icon="el-icon-delete">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </template>
    <!-- 新建/修改 -->
    <OperaDepartmentWindow ref="operaDepartmentWindow" @success="handlePageChange(tableData.pagination.pageIndex)" />

    <OperaCompanyWindow ref="OperaCompanyWindow" @success="handlePageChange(tableData.pagination.pageIndex)" />
  </TableLayout>
</template>

<script>
import TableLayout from "@/layouts/TableLayout";
import { fetchTree } from "@/api/system/department";
import BaseTable from "@/components/base/BaseTable";
import OperaDepartmentWindow from "@/components/system/department/OperaDepartmentWindow";
import OperaCompanyWindow from "@/components/system/department/OperaCompanyWindow";
export default {
  name: "SystemDepartment",
  extends: BaseTable,
  components: { OperaDepartmentWindow, TableLayout, OperaCompanyWindow },
  data() {
    return {
      // 搜索
      searchForm: {
      },
    };
  },
  methods: {
    // 查询数据
    handlePageChange() {
      this.tableData.list.splice(0, this.tableData.list.length);
      this.isWorking.search = true;
      fetchTree()
        .then((records) => {
          this.tableData.list = records;
        })
        .catch((e) => {
          this.$tip.apiFailed(e);
        })
        .finally(() => {
          this.isWorking.search = false;
        });
    },
    getText(row) {
      let text = this.$dict.getDictItemText(
        this.$consts.dicConstant.orgType,
        row.orgType
      );
      return text;
    },
  },
  created() {
    this.config({
      module: "部门",
      api: "/system/department",
      // 主字段
      "field.main": "departName",
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
