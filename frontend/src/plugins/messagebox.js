import { MessageBox } from 'element-ui'

export default {
  ...MessageBox,
  // 删除二次确认
  deleteConfirm (message) {
    return MessageBox.confirm(message, '删除提醒', {
      confirmButtonText: '确认删除',
      cancelButtonText: '取消',
      type: 'warning'
    })
  },
  // 禁用二次确认
  disableConfirm (message) {
    return MessageBox.confirm(message, '禁用提醒', {
      confirmButtonText: '确认禁用',
      cancelButtonText: '取消',
      type: 'warning'
    })
  },
  // 导出二次确认
  exportConfirm (message) {
    return MessageBox.confirm(message, '导出提醒', {
      confirmButtonText: '确认导出',
      cancelButtonText: '取消',
      type: 'warning'
    })
  },
  // 确认冻结
    frozenBatchConfirm (message,status) {
        if(status == 2){
            return MessageBox.confirm(message, '提醒', {
                confirmButtonText: '确认冻结',
                cancelButtonText: '取消',
                type: 'warning'
            })
        }
        else{
            return MessageBox.confirm(message, '提醒', {
                confirmButtonText: '确认解冻',
                cancelButtonText: '取消',
                type: 'warning'
            })
        }
    
    },
}
