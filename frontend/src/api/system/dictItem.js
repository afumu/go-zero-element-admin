import request from '@/utils/request'

// 查询
export function fetchList (data) {
  return request.get('/sys/dictItem/list', { params : data })
}

// 新建
export function create (data) {
  return request.post('/sys/dictItem/add', data)
}

// 修改
export function updateById (data) {
  return request.put('/sys/dictItem/edit', data)
}

// 删除
export function deleteById (id) {
  return request.delete('/sys/dictItem/delete', { params : {'id': id} })
}

// // 批量删除 字典数据没有批量删除功能
// export function deleteByIdInBatch (ids) {
//   return request.get('/system/dictData/delete/batch', {
//     params: {
//       ids
//     }
//   })
// }
