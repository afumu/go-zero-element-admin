<template>
  <TableLayout :permissions="['system:addressList:query']">
    <div slot="table-wrap" class="content">
      <div class="left">
        <el-input placeholder="输入部门名称查询..." size="100%" v-model="filterText"> </el-input>
        <el-tree class="filter-tree" :data="treeList" :props="defaultProps" default-expand-all :filter-node-method="filterNode" ref="tree" @node-click="handleNodeClick"> </el-tree>
      </div>
      <div class="right">
        <ul>
          <el-form :model="searchForm" ref="form" label-width="100px" inline>
            <el-form-item label="姓名：" prop="realname">
              <el-input placeholder="请输入姓名查询" v-model="searchForm.realname" />
            </el-form-item>
            <el-form-item label="工号：" prop="workNo">
              <el-input v-model="searchForm.workNo" placeholder="请输入工号查询" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" size="mini" icon="el-icon-search" @click.stop="search">搜索</el-button>
            </el-form-item>
            <el-form-item>
              <el-button @click.stop="reset" size="mini" icon="el-icon-refresh">重置 </el-button>
            </el-form-item>
          </el-form>
        </ul>
        <el-table border v-loading="isWorking.search" :data="tableData.list" :default-sort="{ prop: 'createTime', order: 'descending' }" stripe @selection-change="handleSelectionChange" @sort-change="handleSortChange">
          <el-table-column align="center" label="#" fixed="left" type="index" width="55"></el-table-column>
          <el-table-column align="center" prop="realname" label="姓名" fixed="left" min-width="80"></el-table-column>
          <el-table-column align="center" prop="workNo" label="工号" fixed="left" min-width="80"></el-table-column>
          <el-table-column align="center" prop="departName" label="部门" fixed="left" min-width="100"></el-table-column>
          <el-table-column align="center" prop="post" label="用户名称" fixed="left" min-width="100"></el-table-column>
          <el-table-column align="center" prop="telephone" label="用户名称" fixed="left" min-width="100"></el-table-column>
          <el-table-column align="center" prop="email" label="用户名称" fixed="left" min-width="100"></el-table-column>
        </el-table>
      </div>
    </div>
  </TableLayout>
</template>
<script>
import { queryTreeList } from "@/api/system/addressList";
import TableLayout from "@/layouts/TableLayout";
import BaseTable from "../../components/base/BaseTable";
export default {
  extends: BaseTable,
  components: { TableLayout },
  data() {
    return {
      treeList: [],
      filterText: "",
      defaultProps: {
        children: "children",
        label: "departName",
      },
      searchForm: {
        realname: "",
        workNo: "",
        orgCode: '',
      },
    };
  },
  created() {
    this.config({
      module: "通讯录",
      api: "/system/addressList",
      "field.main": "departName",
    });
    queryTreeList().then((res) => {
      this.treeList = res;
    });
    this.search();
  },
  watch: {
    filterText(val) {
      this.$refs.tree.filter(val);
    },
  },

  methods: {
    reset() {
      this.$refs.form.resetFields();
      this.searchForm.orgCode = ''
      this.search();
    },
    queryTreeData(keyword) {
      this.commonRequestThen(
        queryDepartTreeList({
          departName: keyword ? keyword : undefined,
        })
      );
    },
    filterNode(value, data) {
      if (!value) return true;
      return data.departName.indexOf(value) !== -1;
    },
    handleSearch(value) {
      if (value) {
        this.commonRequestThen(searchByKeywords({ keyWord: value }));
      } else {
        this.queryTreeData();
      }
    },
    resetReportCondition() {
      (this.searchForm = {
        realname: "",
        workno: "",
      }),
        this.reset();
    },
    handleTreeSelect(selectedKeys, event) {
      if (selectedKeys.length > 0 && this.selectedKeys[0] !== selectedKeys[0]) {
        this.selectedKeys = [selectedKeys[0]];
        let orgCode = event.node.dataRef.orgCode;
        this.emitInput(orgCode);
      }
    },

    emitInput(orgCode) {
      this.$emit("input", orgCode);
    },

    commonRequestThen(promise) {
      this.loading = true;
      promise
        .then((res) => {
          if (res.success) {
            this.treeDataSource = res.result;
          } else {
            this.$message.warn(res.message);
            console.error("组织机构查询失败:", res);
          }
        })
        .finally(() => {
          this.loading = false;
          this.cardLoading = false;
        });
    },
    handleNodeClick(data) {
        this.searchForm.orgCode = data.orgCode 
        this.search()
    },
  },
};
</script>
<style lang="scss" scoped>
.content {
  display: flex;
  .left {
    flex: 1;
    background-color: #fff;
    min-height: 300px;
    margin-right: 10px;
    padding: 30px;
    border: 1px solid #eee;
    .el-input__icon {
      background-color: #13c2c2;
      color: #fff;
      margin-right: -5px;
      border-top-right-radius: 3px;
      border-bottom-right-radius: 3px;
      padding: 0 5px;
    }
  }
  .right {
    flex: 3;
    // background-color: #000;
    height: 100px;
    margin-left: 10px;
  }
}
</style>