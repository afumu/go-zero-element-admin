<template>
  <div>
    <!-- 搜索表单 -->
    <div class="table-search-form">
      <div class="form-wrap">
        <el-collapse v-model="activeNames" accordion>
          <el-collapse-item title="查询条件" name="1">
            <template slot="title">
              <div class="kj-title">
                <span style="margin: 0 10px">查询条件</span>
                <section>
                  <el-button type="primary" size="mini" icon="el-icon-search" @click.stop="search">搜索
                  </el-button>
                  <el-button @click.stop="reset" size="mini" icon="el-icon-refresh">重置
                  </el-button>
                </section>
              </div>
            </template>
            <el-form ref="searchForm" :model="searchForm" label-width="100px" inline>
              <el-form-item label="搜索日志" prop="keyWord">
                <el-input v-model="searchForm.keyWord" placeholder="请输入日志关键字" @keypress.enter.native="search"></el-input>
              </el-form-item>
              <el-form-item label="创建时间" prop="loginTime">
                <el-date-picker v-model="searchDateRange" type="datetimerange" range-separator="至" value-format="yyyy-MM-dd HH:mm:ss" start-placeholder="开始时间" end-placeholder="结束时间" @change="handleSearchTimeChange"></el-date-picker>
              </el-form-item>
            </el-form>
          </el-collapse-item>
        </el-collapse>
      </div>
    </div>

    <!-- 表格和分页 -->
    <div class="table-content">
      <div class="table-wrap">
        <!-- 表格和分页 -->
        <el-table border v-loading="isWorking.search" :data="tableData.list" stripe :default-sort="{prop: 'loginTime', order: 'descending'}" @sort-change="handleSortChange">
          <el-table-column prop="logContent" label="日志内容" min-width="110px"></el-table-column>
          <el-table-column prop="ip" label="IP" min-width="120px"></el-table-column>
          <el-table-column prop="logType_dictText" label="日志类型" min-width="100px"></el-table-column>
          <el-table-column prop="createTime" label="创建时间" min-width="160px" sortable="custom" sort-by="LOGIN_TIME"></el-table-column>
        </el-table>
        <pagination @size-change="handleSizeChange" @current-change="handlePageChange" :pagination="tableData.pagination"></pagination>
      </div>
    </div>
  </div>
</template>

<script>
import BaseTable from "@/components/base/BaseTable";
import TableLayout from "@/layouts/TableLayout";
import Pagination from "@/components/common/Pagination";
export default {
  name: "loginLog",
  extends: BaseTable,
  components: {
    TableLayout,
    Pagination,
  },
  data() {
    return {
      searchDateRange: [],
      // 搜索
      searchForm: {
        logType: "1",
        keyWord: null,
        order: "desc",
        column: "createTime",
        createTime_begin: null,
        createTime_end: null,
        operateType: null,
      },
      activeNames: "1",
    };
  },
  methods: {
    // 时间搜索范围变化
    handleSearchTimeChange(value) {
      if (value != null) {
        this.searchForm.createTime_begin = value[0];
        this.searchForm.createTime_end = value[1];
      } else {
        this.searchForm.createTime_begin = null;
        this.searchForm.createTime_end = null;
      }
      this.search();
    },
    // 搜索框重置
    reset() {
      this.$refs.searchForm.resetFields();
      this.searchDateRange = [];
      this.searchForm.createTime_begin = null;
      this.searchForm.createTime_end = null;
      this.search();
    },
  },
  created() {
    this.config({
      module: "日志",
      api: "/system/loginLog",
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
