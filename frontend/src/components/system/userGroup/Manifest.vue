<template>
    <el-dialog title="授权" :visible.sync="visible">
        <el-form :ref="searchForm" :model="searchForm" label-width="100px" inline>
            <el-form-item label="规则名称" prop="ruleName">
                <search-input v-model="searchForm.ruleName" @keypress.enter.native="search"></search-input>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" icon="el-icon-search" @click.stop="search">查询</el-button>
                <el-button type="primary" @click.stop="reset" icon="el-icon-refresh">重置</el-button>
                <el-button type="primary" @click.stop="$refs.OperaddRuleWindow.open(searchForm.permissionId)" >新增</el-button>
            </el-form-item>
        </el-form>
        <el-table border v-loading="isWorking.search" :data="tableData.list" stripe>
            <el-table-column type="index" fixed="left" align="center"></el-table-column>
            <el-table-column label="规则名称" prop="ruleName" fixed="left" align="center"></el-table-column>
            <el-table-column label="规则字段" prop="ruleColumn_dictText" fixed="left" align="center"></el-table-column>
            <el-table-column label="规则值" prop="ruleValueText" fixed="left" align="center"></el-table-column>
            <el-table-column label="操作" fixed="left" align="center">
                <template slot-scope="{ row }">
                    <el-button type="text" @click="$refs.OperaddRuleWindow.open( '编辑',row )">编辑</el-button>
                    <el-button type="text" @click="deleteId(row.id)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
        <OperaddRuleWindow ref="OperaddRuleWindow" @success="search()"/>
    </el-dialog>
</template>
<script>
import TableLayout from "@/layouts/TableLayout";
import BaseTable from "@/components/base/BaseTable";
import Pagination from '@/components/common/Pagination.vue';
import OperaddRuleWindow from '@/components/system/userGroup/OperaddRuleWindow'
export default {
    extends: BaseTable,
    components: {TableLayout,Pagination,OperaddRuleWindow},
    data(){
        return{
            visible:false,
            searchForm:{
                permissionId:'',
                ruleName:'',
                ruleColumn: '',
            }
        }
    },
    methods:{
        open(type){
            this.$nextTick(() =>{
                this.searchForm.permissionId = type.id
                this.search()
            })
            this.visible = true
        },
        search(){
            this.api.queryPermissionRule(this.searchForm).then((res) => {
                    this.tableData.list = res
                })
        },
        deleteId(id){
            this.api.deleteId(id).then(res =>{
                this.$tip.success(res.message)
                this.search()
            })
        },
        reset(){
            this.searchForm.ruleColumn = '';
            this.searchForm.ruleName = '';
            this.search()
        }
    },
    created(){
        this.config({
            api:'/system/userGroup/Manifest'
        })
    }
}
</script>