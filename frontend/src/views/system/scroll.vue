<template>
  <TableLayout :permissions="['system:scroll:query']">
    <!-- 表格和分页 -->
    <template v-slot:table-wrap>
      <el-button icon="el-icon-plus" type="primary" @click="$refs.operaScrollWindow.open('新建')">新建</el-button>
      <el-table border v-loading="isWorking.search" :data="tableData.list" :default-sort="{ prop: 'createTime', order: 'descending' }" stripe>
        <el-table-column label="缩略图" fixed="left" min-width="150px">
          <template slot-scope="{ row }">
            <img :src="row.fileId_dictText" width="100" height="37.2" class="head_pic" />
          </template>
        </el-table-column>
        <el-table-column prop="memo" label="描述" fixed="left" min-width="150px"></el-table-column>
        <el-table-column prop="sort" label="排序" fixed="left" min-width="150px"></el-table-column>
        <el-table-column label="是否开启" fixed="left" align="center" min-width="150px">
          <template slot-scope="{ row }">
            <el-switch v-model="row.enable" active-text="开" :active-value="1" inactive-text="关" :inactive-value="0" @change="changeSwitch(row)"></el-switch>
          </template>
        </el-table-column>
        <el-table-column label="操作" min-width="80" fixed="right" align="center">
          <template slot-scope="{ row }">
            <el-button type="text" @click="deleteById(row)" icon="el-icon-delete">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <pagination @size-change="handleSizeChange" @current-change="handlePageChange" :pagination="tableData.pagination"></pagination>

      <OperaScrollWindow ref="operaScrollWindow" @success="handlePageChange(tableData.pagination.pageIndex)" />
    </template>
  </TableLayout>
</template>

<script>
import Pagination from "@/components/common/Pagination";
import TableLayout from "@/layouts/TableLayout";
import BaseTable from "@/components/base/BaseTable";
import { updateEnable } from "@/api/system/scroll";
import OperaScrollWindow from "@/components/system/scroll/OperaScrollWindow";

export default {
  name: "Scroll",
  extends: BaseTable,
  components: {
    OperaScrollWindow,
    TableLayout,
    Pagination,
  },
  data() {
    return {
      searchForm: {
        order: "desc",
        column: "createTime",
      },
    };
  },
  methods: {
    changeSwitch(row) {
      updateEnable({
        id: row.id,
        enable: row.enable,
      }).then((res) => {
        this.$notify({
          title: "成功",
          message: res.message,
          type: "success",
        });
      });
    },
  },
  created() {
    this.config({
      module: "轮播图",
      api: "/system/scroll",
      "field.id": "id",
      "field.main": "id",
      sorts: [
        {
          property: "LOGIN_TIME",
          direction: "DESC",
        },
      ],
    });
    this.search();
  },
};
</script>

<style scoped lang="scss">
@import "@/assets/style/variables.scss";
</style>
