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
              <el-form-item label="日志内容" prop="keyWord">
                <el-input v-model="searchForm.keyWord" placeholder="请输入搜索关键词" @keypress.enter.native="search"></el-input>
              </el-form-item>
              <el-form-item label="创建时间">
                <el-date-picker v-model="searchDateRange" type="datetimerange" range-separator="至" value-format="yyyy-MM-dd HH:mm:ss" start-placeholder="开始时间" end-placeholder="结束时间" @change="handleSearchTimeChange"></el-date-picker>
              </el-form-item>
              <el-form-item label="操作类型" prop="operateType">
                <el-select v-model="searchForm.operateType" clearable @change="search">
                  <el-option value="1" label="查询" />
                  <el-option value="2" label="添加" />
                  <el-option value="3" label="修改" />
                  <el-option value="4" label="删除" />
                  <el-option value="5" label="导入" />
                  <el-option value="6" label="导出" />
                </el-select>
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
        <el-table border v-loading="isWorking.search" :data="tableData.list" stripe :default-sort="{prop: 'operaTime', order: 'descending'}" @sort-change="handleSortChange">
          <el-table-column type="expand">
            <template slot-scope="props">
              <div><i class="el-icon-info"></i>{{ props.row.method }}</div>
              <div><i class="el-icon-info"></i>{{ props.row.requestParam }}</div>
            </template>
          </el-table-column>
          <el-table-column prop="logContent" label="日志内容" min-width="100px"></el-table-column>
          <el-table-column prop="userid" label="操作人ID" min-width="100px"></el-table-column>
          <el-table-column prop="username" label="操作人名称" min-width="100px"></el-table-column>
          <el-table-column prop="ip" label="IP" min-width="100px"></el-table-column>
          <el-table-column prop="costTime" label="耗时（毫秒）" min-width="100px"></el-table-column>
          <el-table-column prop="logType_dictText" label="日志类型" min-width="100px"></el-table-column>
          <el-table-column prop="operateType_dictText" label="操作类型" min-width="100px"></el-table-column>
          <el-table-column prop="createTime" label="创建时间" min-width="100px"></el-table-column>
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
  name: "systemLog",
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
        logType: "2",
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
