<template>
  <TableLayout :permissions="['system:dict:query']">
    <!-- 搜索表单 -->
    <div slot="search-form">
      <el-collapse v-model="activeNames" accordion>
        <el-collapse-item title="查询条件" name="1">
          <template slot="title">
            <div class="kj-title">
              <span style="margin: 0 10px">查询条件</span>
              <section>
                <el-button type="primary" size="mini" icon="el-icon-search" @click.stop="search">搜索</el-button>
                <el-button @click.stop="reset" size="mini" icon="el-icon-refresh">重置</el-button>
              </section>
            </div>
          </template>
          <el-form ref="searchForm" :model="searchForm" label-width="100px" inline>
            <el-form-item label="字典编码" prop="dictCode">
              <search-input v-model="searchForm.dictCode" placeholder="请输入字典编码" @keypress.enter.native="search" />
            </el-form-item>
            <el-form-item label="字典名称" prop="dictName">
              <search-input v-model="searchForm.dictName" placeholder="请输入字典名称" @keypress.enter.native="search" />
            </el-form-item>
          </el-form>
        </el-collapse-item>
      </el-collapse>
      <div style="height: 10px"></div>
    </div>
    <!-- 表格和分页 -->
    <template v-slot:table-wrap>
      <ul class="toolbar">
        <li>
          <el-button type="primary" @click="$refs.operaDictWindow.open('新建字典')" icon="el-icon-plus">新建</el-button>
        </li>
        <li>
          <el-button type="primary" @click="refleshCache" icon="el-icon-refresh">刷新缓存</el-button>
        </li>
        <li>
          <el-button type="primary" @click="$refs.dictDeleteList.open()" icon="el-icon-delete">回收站</el-button>
        </li>
      </ul>
      <el-table border v-loading="isWorking.search" :data="tableData.list" :default-sort="{ prop: 'createTime' }" stripe @sort-change="handleSortChange">
        <!-- <el-table-column label="#" fixed="left" min-width="50px" align="center">
          <template slot-scope="scope">
            <label>{{ scope.$index + 1 }}</label>
          </template>
        </el-table-column> -->
        <el-table-column prop="dictName" label="字典名称" fixed="left" min-width="100px"></el-table-column>
        <el-table-column prop="dictCode" label="字典编码" fixed="left" min-width="100px"></el-table-column>
        <el-table-column prop="description" label="描述" min-width="150px"></el-table-column>
        <el-table-column label="操作" min-width="150" fixed="right" align="center">
          <template slot-scope="{ row }">
            <el-button type="text" @click="$refs.operaDictWindow.open('编辑字典', row)" icon="el-icon-edit" style="margin-left: 0;!important">编辑</el-button>
            <el-button type="text" @click="$refs.dictDataManagerWindow.open('字典列表', row.id)" icon="el-icon-setting">字典配置</el-button>
            <el-button type="text" @click="deleteById(row)" icon="el-icon-delete">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <pagination @size-change="handleSizeChange" @current-change="handlePageChange" :pagination="tableData.pagination"></pagination>
    </template>
    <!-- 新建/修改 -->
    <OperaDictWindow ref="operaDictWindow" @success="handlePageChange(tableData.pagination.pageIndex)" />
    <!-- 字典配置 -->
    <DictDataManagerWindow ref="dictDataManagerWindow" />
    <!-- 回收站 -->
    <DictDeleteList ref="dictDeleteList" @success="search()"/>
  </TableLayout>
</template>

<script>
import Pagination from "@/components/common/Pagination";
import TableLayout from "@/layouts/TableLayout";
import BaseTable from "@/components/base/BaseTable";
import OperaDictWindow from "@/components/system/dict/OperaDictWindow";
import DictDataManagerWindow from "@/components/system/dict/DictDataManagerWindow";
import DictDeleteList from "@/components/system/dict/DictDeleteList";
export default {
  name: "SystemDict",
  extends: BaseTable,
  components: {
    DictDataManagerWindow,
    OperaDictWindow,
    TableLayout,
    Pagination,
    DictDeleteList,
  },
  data() {
    return {
      // 搜索
      searchForm: {
        dictCode: "",
        dictName: "",
        column: "createTime",
        order: "desc",
      },
      activeNames: "1",
    };
  },
  methods: {
    refleshCache: function () {
      this.api
        .refleshCache()
        .then(() => {
          this.$tip.apiSuccess("刷新缓存完成");
        })
        .catch((e) => {
          this.$tip.apiFailed(e);
        });
    },
    searchMethod() {},
  },
  created() {
    this.config({
      module: "字典",
      api: "/system/dict",
      sorts: [
        {
          property: "dict.CREATE_TIME",
          direction: "DESC",
        },
      ],
      "field.main": "dictName",
    });
    this.search();
  },
};
</script>
