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
              <el-form-item label="用户姓名" prop="realname">
                <searchInput v-model="searchForm.realname" placeholder="请输入用户姓名"></searchInput>
              </el-form-item>
            </el-form>
          </el-collapse-item>
        </el-collapse>
      </div>
      <template v-slot:table-wrap>
        <!-- <ul class="toolbar">
          <li>
            <el-button type="primary" icon="el-icon-plus" @click="existingUsers">已有用户</el-button>
          </li>
        </ul> -->
        <el-table border v-loading="isWorking.search" :data="tableData.list" :default-sort="{ prop: 'createTime', order: 'descending' }" stripe @selection-change="handleSelectionChange" @sort-change="handleSortChange">
          <el-table-column type="selection" fixed="left" align="center" width="55"></el-table-column>
          <el-table-column prop="username" label="用户账号" align="center" fixed="left" min-width="100px"></el-table-column>
          <el-table-column prop="realname" label="用户名称" align="center" fixed="left" min-width="100px"></el-table-column>
          <el-table-column prop="sex_dictText" label="性别" align="center" min-width="60px"></el-table-column>
          <el-table-column prop="orgCodeTxt" label="部门" align="center" min-width="60px"></el-table-column>
        </el-table>
        <pagination @size-change="handleSizeChange" @current-change="handlePageChange" :pagination="tableData.pagination"></pagination>
      </template>
    </TableLayout>
  </GlobalWindow>
</template>
<script>
import TableLayout from "@/layouts/TableLayout";
import GlobalWindow from "@/components/common/GlobalWindow";
import searchInput from "@/components/common/searchInput";
import BaseTable from "@/components/base/BaseTable";
import Pagination from "@/components/common/Pagination";
import { addSysUserRole } from "@/api/system/userRole";

export default {
  extends: BaseTable,
  components: { GlobalWindow, searchInput, TableLayout, Pagination },
  data() {
    return {
      visible: false,
      // isWorking: false,
      searchForm: {
        roleId: "",
        realname: "",
        username :""
      },
      activeNames: "1",
      ID: "",
    };
  },
  created() {
    this.config({
      module: "添加已有用户",
      api: "/system/user",
    });
  },
  methods: {
    open(row) {
      this.ID = row;
      this.visible = true;
      this.searchForm.username = '',
      this.searchForm.realname = "",
      this.search();
    },
    confirm() {
      let thisArr = this.tableData.selectedRows.map((row) => row.id);
      let thisStr = {};
      thisStr.userIdList = thisArr;
      thisStr.roleId = this.ID;
      addSysUserRole(thisStr)
        .then((res) => {
          if (res.success) {
            this.$notify({
              title: "成功",
              message: res.message,
              type: "success",
            });
          }
          this.$emit("success");
        })
        .finally(() => {
          this.visible = false;
        });
    },
  },
};
</script>
