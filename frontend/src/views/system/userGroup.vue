<template>
    <TableLayout :permissions="['system:userGroup:query']">
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
                        <el-form-item label="名称" prop="name">
                            <search-input v-model="searchForm.name" v-trim placeholder="请输入名称" @keypress.enter.native="search" />
                        </el-form-item>
                    </el-form>
                </el-collapse-item>
            </el-collapse>
        </div>
        <template v-slot:table-wrap>
            <ul class="toolbar">
                <li>
                    <el-button type="primary" @click="$refs.OperauserGroupWindow.open('新建')">新建</el-button>
                </li>
            </ul>
            <el-table  border v-loading="isWorking.search" :data="tableData.list" stripe @sort-change="handleSortChange" @selection-change="handleSelectionChange">
                <el-table-column align="center" prop="name" label="名称" fixed="left" min-width="110px"></el-table-column>
                <el-table-column align="center" prop="memo" label="备注" fixed="left" min-width="110px"></el-table-column>
                <el-table-column align="center" prop="createTime" label="创建时间" fixed="left" min-width="110px"></el-table-column>
                <el-table-column prop="number" label="操作 " fixed="left" align="center" min-width="110px">
                    <template slot-scope="{ row }">
                        <el-button type="text" @click="$refs.User.open('用户',row)">用户</el-button>
                        <el-button type="text" @click="$refs.OperauserGroupWindow.open('修改',row)" >修改</el-button>
                        <el-button type="text" @click="$refs.Manifest.open(row)" >授权</el-button>
                        <el-button type="text" @click="deleteById(row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <pagination @size-change="handleSizeChange" @current-change="handlePageChange" :pagination="tableData.pagination"></pagination>
        </template>
        <OperauserGroupWindow ref="OperauserGroupWindow" @success="handlePageChange(tableData.pagination.pageIndex)" ></OperauserGroupWindow>
        <User ref="User" @success="handlePageChange(tableData.pagination.pageIndex)"></User>
        <Manifest ref="Manifest" @success="handlePageChange(tableData.pagintion.pageIndex)"></Manifest>
    </TableLayout>
</template>
<script>
import TableLayout from "@/layouts/TableLayout";
import BaseTable from "@/components/base/BaseTable";
import Pagination from '@/components/common/Pagination.vue'
import OperauserGroupWindow from '@/components/system/userGroup/OperauserGroupWindow'
import User from '@/components/system/userGroup/user'
import Manifest from '@/components/system/userGroup/Manifest'
export default {
    extends: BaseTable,
    components: {TableLayout,Pagination ,OperauserGroupWindow,User,Manifest},
    data(){
        return{
            searchForm:{
                name:''
            },
            activeNames:'1',
            Visible:false
        }
    },
    created() {
    this.config({
      api: "/system/userGroup/userGroup",
      "field.main": "name",
    });
    this.search();
  },
  methods:{
      userGroup(){
          this.Visible = true
      }
  }
}
</script>
<style lang="scss" scoped>

</style>