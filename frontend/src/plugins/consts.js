/**
 * 常量定义
 */

/**
 * 字典常量
 * @type {{applyResult: string}}
 */
const dicConstant = {
    sex: 'sex', // 性别
    msg_category: 'msg_category',
    send_status: 'send_status',
    userStatus: 'user_status',
    orgType: 'org_category',
    templateType: 'template_type',
}


/**
 * 文件相关
 * @type {{}}
 */
const fileConstant = {
    importUserPath: window._CONFIG['domianURL'] + '/sys/user/importExcel', // 导入用户路径（集成部署的时候，弄去除路径）
    exportUserPath: '/sys/user/exportXls', // 导出用户路径
    uploadTemplateFilePath: window._CONFIG['domianURL'] + '/system/sysFile/upload', // 上传文件
    exportApplyUserPath: '/business/applyUser/exportXls', // 导出用户路径

    importAssetPath: window._CONFIG['domianURL'] + '/base/asset/importExcel',
}

export default {
    dicConstant: dicConstant,
    fileConstant: fileConstant,
}
