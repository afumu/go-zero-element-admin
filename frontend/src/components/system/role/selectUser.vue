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
                <el-input v-model="searchForm.username" placeholder="请输入用户账号"></el-input>
              </el-form-item>
            </el-form>
          </el-collapse-item>
        </el-collapse>
      </div>
      <template v-slot:table-wrap>
        <ul class="toolbar">
          <li>
            <el-button type="primary" icon="el-icon-plus" @click="existingUsers">已有用户</el-button>
            <el-button @click="deleteByIdInBatch_all" icon="el-icon-delete">删除</el-button>
          </li>
        </ul>
        <el-table border v-loading="isWorking.search" :data="tableData.list" :default-sort="{ prop: 'createTime', order: 'descending' }" stripe @selection-change="handleSelectionChange" @sort-change="handleSortChange">
          <el-table-column type="selection" align="center" fixed="left" width="55"></el-table-column>
          <el-table-column prop="username" align="center" label="用户账号" fixed="left" min-width="100px"></el-table-column>
          <el-table-column prop="realname" align="center" label="用户名称" fixed="left" min-width="100px"></el-table-column>
          <el-table-column prop="status_dictText" align="center" label="状态" min-width="60px"></el-table-column>
          <el-table-column align="center" label="操作" min-width="100" fixed="right">
            <template slot-scope="{ row }">
              <el-button v-if="!row.fixed" type="text" @click="deleteById_row(row)" icon="el-icon-delete">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <pagination @size-change="handleSizeChange" @current-change="handlePageChange" :pagination="tableData.pagination"></pagination>
      </template>
    </TableLayout>

    <existing ref="existing" @success="handlePage"></existing>

  </GlobalWindow>
</template>
<script>
import TableLayout from "@/layouts/TableLayout";
import GlobalWindow from "@/components/common/GlobalWindow";
import searchInput from "@/components/common/searchInput";
import BaseTable from "@/components/base/BaseTable";
import Pagination from "@/components/common/Pagination";
import { deleteById, deleteByIdInBatch } from "@/api/system/userRole";

import existing from "@/components/system/role/existingUsers";

export default {
  extends: BaseTable,
  components: { GlobalWindow, searchInput, TableLayout, Pagination, existing },
  data() {
    return {
      visible: false,
      // isWorking: false,
      searchForm: {
        roleId: "",
      },
      activeNames: "1",
      ID: "",
    };
  },
  created() {
    this.config({
      module: "角色",
      api: "/system/userRole",
    });
  },
  methods: {
    open(row) {
      this.ID = row.id;
      this.visible = true;
      this.searchForm.roleId = row.id;
      this.search();
      // fetchMenuTree({}).then((records) => {});
    },
    confirm() {
      this.visible = false;
      this.$emit("success");
    },
    deleteByIdInBatch_all(childConfirm = true) {
      this.__checkApi();
      if (this.tableData.selectedRows.length === 0) {
        this.$tip.warning("请至少选择一条数据");
        return;
      }
      let message = `确认删除已选中的 ${this.tableData.selectedRows.length} 条${this.module}记录吗?`;
      if (childConfirm) {
        const containChildrenRows = [];
        for (const row of this.tableData.selectedRows) {
          if (row.children != null && row.children.length > 0) {
            containChildrenRows.push(row[this.configData["field.main"]]);
          }
        }
        if (containChildrenRows.length > 0) {
          message = `本次将删除${this.module}【${containChildrenRows.join(
            "、"
          )}】及其子${this.module}记录，确认删除吗？`;
        }
      }
      this.$dialog
        .deleteConfirm(message)
        .then(() => {
          // this.isWorking.delete = true
          deleteByIdInBatch({
            userIds: this.tableData.selectedRows.map((row) => row.id).join(","),
            roleId: this.ID,
          })
            .then(() => {
              this.$tip.apiSuccess("删除成功");
              this.__afterDelete(this.tableData.selectedRows.length);
            })
            .catch((e) => {
              this.$tip.apiFailed(e);
            })
            .finally(() => {
              this.isWorking.delete = false;
            });
        })
        .catch(() => {});
    },
    // 删除
    deleteById_row(row, childConfirm = true) {
      this.__checkApi();
      let message = `确认删除吗?`;
      if (childConfirm && row.children != null && row.children.length > 0) {
        message = `确认删除${this.module}【${
          row[this.configData["field.main"]]
        }】及其子${this.module}吗?`;
      }

      this.$dialog
        .deleteConfirm(message)
        .then(() => {
          deleteById({
            userId: row[this.configData["field.id"]],
            roleId: this.ID,
          })
            .then(() => {
              this.$tip.apiSuccess("删除成功");
              this.__afterDelete();
            })
            .catch((e) => {
              this.$tip.apiFailed(e);
            })
            .finally(() => {
              this.isWorking.delete = false;
            });
        })
        .catch(() => {});
    },
    // 已有用户
    existingUsers() {
      this.$refs.existing.open(this.ID);
    },
    handlePage() {
      this.search();
    },
  },
};
</script>
