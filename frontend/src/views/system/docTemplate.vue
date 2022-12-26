<template>
    <TableLayout :permissions="['system:docTemplate:query']">
    <!-- 搜索表单 -->
    <div slot="search-form">
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
            <el-form-item label="文件名称" prop="docName">
              <search-input v-model="searchForm.docName" v-trim placeholder="请输入文件名称" @keypress.enter.native="search" />
            </el-form-item>
          </el-form>
        </el-collapse-item>
      </el-collapse>
    </div>
    <template v-slot:table-wrap>
      <ul class="toolbar">
        <li>
          <el-button icon="el-icon-plus" type="primary" @click="$refs.operaDocTemplateWindow.open('新建')">
            新建
          </el-button>
        </li>
      </ul>
      <el-table border v-loading="isWorking.search" :data="tableData.list" :default-sort="{ prop: 'createTime', order: 'descending' }" stripe @sort-change="handleSortChange">
        <el-table-column prop="docName" label="文档名称" fixed="left" min-width="150px"></el-table-column>
        <el-table-column label="操作" min-width="80" fixed="right" align="center">
          <template slot-scope="{ row }">
            <el-button style="margin-left: 0;!important" type="text" @click="$refs.operaDocTemplateWindow.open('编辑', row)" icon="el-icon-edit">编辑</el-button>
            <el-button style="margin-left: -15px;!important" type="text" @click="deleteById(row)" icon="el-icon-delete">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
        <pagination @size-change="handleSizeChange" @current-change="handlePageChange" :pagination="tableData.pagination"></pagination>
    </template>

    <template>
      <OperaDocTemplateWindow ref="operaDocTemplateWindow" @success="handlePageChange(tableData.pagination.pageIndex)" />
    </template>
  </TableLayout>
</template>

<script>
    import TableLayout from "@/layouts/TableLayout";
    import BaseTable from "@/components/base/BaseTable";
    import Pagination from "@/components/common/Pagination";
    import OperaDocTemplateWindow from "@/components/system/docTemplate/OperaDocTemplateWindow";

    export default {
        name: "docTemplate",
        extends: BaseTable,
        components: {
            TableLayout,
            Pagination,
            OperaDocTemplateWindow,
        },
        data() {
            return {
                // 搜索
                searchForm: {
                    id: null,
                    order: 'desc',
                    column: 'createTime',
                    docName: '',
                },
                activeNames: "1",
            };
        },
        created() {
            this.config({
                module: "模板",
                api: "/system/docTemplate",
                'field.id': 'id',
                'field.main': 'docName',
            })
            this.search()
        }
    }
</script>

<style scoped>

</style>
