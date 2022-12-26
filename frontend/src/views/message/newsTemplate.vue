<template>
    <TableLayout>
        <div slot="search-form">
            <el-collapse v-model="activeNames" accordion>
                <el-collapse-item name="1" title="查询条件">
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
                        <el-form-item label="模板CODE" prop="templateCode">
                            <search-input v-model="searchForm.templateCode"></search-input>
                        </el-form-item>
                        <el-form-item label="模板标题" prop="templateName">
                            <search-input v-model="searchForm.templateName"></search-input>
                        </el-form-item>
                        <el-form-item label="模板类型" prop="templateType">
                            <el-select clearable v-model="searchForm.templateType">
                                <el-option v-for="item in templateTypeList" :key="item.text" :label="item.text" :value="item.value"></el-option>
                            </el-select>
                        </el-form-item>
                    </el-form>
                </el-collapse-item>
            </el-collapse>
        </div>
        <template v-slot:table-wrap>
            <ul class="toolbar">
                <li>
                    <el-button type="primary" @click="$refs.OperanewsTemplateWindow.open('新建')" icon="el-icon-plus">新建</el-button>
                </li>
            </ul>
            <el-table  border v-loading="isWorking.search" :data="tableData.list" stripe>
                <el-table-column prop="templateCode" label="模板CODE " fixed="left" align="center" min-width="100px"></el-table-column>
                <el-table-column prop="templateName" label="模板标题" fixed="left" align="center" min-width="100px"></el-table-column>
                <el-table-column prop="templateContent" label="模板内容 " fixed="left" align="center" min-width="100px"></el-table-column>
                <el-table-column prop="templateType_dictText" label="模板类型 " fixed="left" align="center" min-width="100px"></el-table-column>
                <el-table-column label="操作 " fixed="left" align="center" min-width="70px">
                    <template slot-scope="{row}">
                        <el-button style="margin-left: 15px;!important" type="text" @click="$refs.OperanewsTemplateWindow.open('编辑', row)" icon="el-icon-edit">编辑</el-button>
                        <el-button style="margin-left: 15px;!important" type="text" @click="deleteById(row)" icon="el-icon-delete">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <pagination @size-change="handleSizeChange" @current-change="handlePageChange" :pagination="tableData.pagination"></pagination>
        </template>
        <OperanewsTemplateWindow ref="OperanewsTemplateWindow" @success="handlePageChange(tableData.pagination.pageIndex)"> </OperanewsTemplateWindow>
    </TableLayout>
</template>
<script>
import BaseTable from '@/components/base/BaseTable.vue'
import Pagination from '@/components/common/Pagination.vue'
import TableLayout from "@/layouts/TableLayout";
// import OperanewsTemplateWindow from './components/OperanewsTemplateWindow.vue'
import OperanewsTemplateWindow from './modules/OperanewsTemplateWindow.vue'
export default {
    extends:BaseTable,
    components:{
        Pagination,
        TableLayout,
        OperanewsTemplateWindow
    },
    data(){
        return{
            activeNames:'1',
            searchForm:{
                templateCode:'',
                templateName:'',
                templateType:'',
            },
            templateTypeList:this.$dict.getAllDictItemByCode(this.$consts.dicConstant.templateType)
        }
    },
    created(){
        this.config({
            api: "/message/newsTemplate",
            "field.main": "templateName",
        })
        this.search()
    },
}
</script>
<style lang="scss" scoped>

</style>