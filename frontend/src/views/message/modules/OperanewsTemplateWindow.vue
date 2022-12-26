<template>
    <GlobalWindow :title="title" :visible.sync="visible" :confirm-working="isWorking" @confirm="confirm" @closeTrigger="$refs.form.resetFields()">
        <el-form ref="form" :model="form" :rules="rules">
            <el-form-item label="模板CODE" prop="templateCode">
                <el-input v-model="form.templateCode" placeholder="请选择模板CODE" v-trim maxlength="50" />
            </el-form-item>
            <el-form-item label="模板标题" prop="templateName">
                <el-input v-model="form.templateName" placeholder="请输入模板标题" v-trim maxlength="50" />
            </el-form-item>
            <el-form-item label="模板内容" prop="templateContent">
                <el-input type="textarea" v-model="form.templateContent" placeholder="请输入模板内容" v-trim maxlength="200" :rows="5" />
            </el-form-item>
            <el-form-item label="模板类型" prop="templateType">
                <el-select clearable v-model="form.templateType" placeholder="请选择模板类型">
                    <el-option v-for="item in templateTypeList" :key="item.text" :label="item.text" :value="item.value"></el-option>
                </el-select>
            </el-form-item>
        </el-form>
    </GlobalWindow>
</template>
<script>
import BaseOpera from "@/components/base/BaseOpera";
import GlobalWindow from "@/components/common/GlobalWindow";
export default {
    extends: BaseOpera,
    components: { GlobalWindow },
    data(){
        return{
            form:{
                id:null,
                templateCode:'',
                templateName:'',
                templateContent:'',
                templateType:'',
            },
            rules:{
                templateCode:[{required:true,message:'请选择模板CODE'}],
                templateName:[{required:true,message:'请输入模板标题'}],
                templateContent:[{required:true,message:'请输入模板内容'}],
                templateType:[{required:true,message:'请选择模板类型'}],
            },
            templateTypeList: this.$dict.getAllDictItemByCode(this.$consts.dicConstant.templateType)
        }
    },
    created(){
        this.config({
            api: "/message/newsTemplate",
        })
    }
}
</script>