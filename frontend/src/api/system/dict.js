import request from '@/utils/request'

// 查询
export function fetchList (data) {
  return request.get('/sys/dict/list', { params : data })
}

// 新建
export function create (data) {
  return request.post('/sys/dict/add', data)
}

// 修改
export function updateById (data) {
  return request.put('/sys/dict/edit', data )
}

// 删除
export function deleteById (id) {
  return request.delete('/sys/dict/delete', { params : {'id': id} })
}
 
// 批量删除  字典没有批量删除功能
// export function deleteByIdInBatch (ids) {
//   return request.get('/system/dict/delete/batch', {
//     params: {
//       ids
//     }
//   })
// }

// 刷新缓存  
export function refleshCache () {
  return request.get('/sys/dict/refleshCache')
}
// 回收站
export function deleteList () {
  return request.get('/sys/dict/deleteList')
}
// 字典取回
export function deleteBack (id) {
  return request.put(`/sys/dict/back/${id}`)
}
// 彻底删除
export function deletePhysic (id) {
  return request.delete(`/sys/dict/deletePhysic/${id}`)
}
// 查询所有字典
export function getAll () {
  return request.get('/sys/dict/queryAllDictItems')
}
