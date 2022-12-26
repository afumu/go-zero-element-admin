import request from '@/utils/request'

// 查询
export function fetchList (data) {
  return request.get('sys/quartzJob/list', { params : data })
}

// 新建
export function create (data) {
  return request.post('/sys/quartzJob/add', data)
}

// 修改
export function updateById (data) {
  return request.put('/sys/quartzJob/edit', data )
}

// 删除
export function deleteById (id) {
  return request.delete('/sys/quartzJob/delete', { params : {'id': id} })
}
 
// 批量删除  字典没有批量删除功能
export function deleteByIdInBatch (ids) {
  return request.get('/system/quartzJob/deleteBatch', {
    params: {
      ids
    }
  })
}
 
// 启动
export function resume (className) {
  return request.get('/sys/quartzJob/resume', { params : {'jobClassName': className} })
}

// 暂停
export function pause (className) {
  return request.get('/sys/quartzJob/pause', { params : {'jobClassName': className} })
}

// 立即执行
export function execute (id) {
  return request.get('/sys/quartzJob/execute', { params : {'id': id} })
}

