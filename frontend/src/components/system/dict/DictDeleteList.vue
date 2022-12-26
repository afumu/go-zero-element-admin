<template>
    <el-dialog
    title="回收站"
    :visible.sync="visible"
    width="70%">
    <el-table
        border
        v-loading="isWorking.search"
        :data="deleteList"
        :default-sort = "{prop: 'createTime'}"
        stripe
      >
        <el-table-column label="#" fixed="left" min-width="50px" align="center">
          <template slot-scope="scope">
            <label>{{scope.$index+1}}</label>
          </template>
        </el-table-column>
        <el-table-column prop="dictName" label="字典名称" fixed="left" min-width="100px"></el-table-column>
        <el-table-column prop="dictCode" label="字典编码" fixed="left" min-width="100px"></el-table-column>
        <el-table-column prop="description" label="描述" min-width="150px"></el-table-column>
        <el-table-column
          label="操作"
          min-width="100"
          fixed="right"
          align="center"
        >
          <template slot-scope="{row}">
            <el-button type="text" @click="deleteBack(row.id)" icon="el-icon-refresh-right" >字典取回</el-button>
            <el-button type="text" @click="deletePhysic(row.id)" icon="el-icon-delete" >彻底删除</el-button>
          </template>
        </el-table-column>
      </el-table>
</el-dialog>
</template>

<script>
import BaseOpera from '@/components/base/BaseOpera'
export default {
    name: 'DictDeleteList',
    extends: BaseOpera,
    data() {
        return {
            visible : false,
            deleteList: []
        }
    },
    methods:{
        // 查询回收站的列表
        search () {
            this.api.deleteList()
                .then((data) => {
                this.deleteList = data
                })
                .catch(e => {
                this.$tip.apiFailed(e)
                })
        },
        // 字典取回
        deleteBack (id) {
            this.api.deleteBack(id)
                .then(() => {
                    this.$tip.apiSuccess('成功恢复')
                    this.search()
                    this.$emit('success')
                })
                .catch(e => {
                    this.$tip.apiFailed(e)
                })

        },
        // 彻底删除
        deletePhysic (id) {
            this.api.deletePhysic(id)
                .then(() => {
                    this.$tip.apiSuccess('成功删除')
                    this.search()
                })
                .catch(e => {
                    this.$tip.apiFailed(e)
                })
        },
        open () {
            this.visible = true
            this.search()
        }
    },
    //生命周期 - 创建完成（访问当前this实例）
    created() {
        this.config({
            api: '/system/dict'
        })
    }
}
</script>
<style scoped lang="scss">
/* @import url(); 引入css类 */
</style>
