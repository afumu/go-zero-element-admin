import request from '@/utils/request'

// 查询
export function fetchTree (data) {
  return request.get('/sys/sysDepart/queryTreeList',data)
}

// 查询部门用户
export function fetchUserList (data) {
  return request.post('/system/department/users', data)
}

// 新建
export function create (data) {
  return request.post('/sys/sysDepart/add', data)
}

// 修改
export function updateById (data) {
  return request.put('/sys/sysDepart/edit', data)
}

// 删除
export function deleteById (id) {
  return request.delete(`/sys/sysDepart/delete?id=${id}`)
}

// 批量删除
export function deleteByIdInBatch (ids) {
  return request.delete('/sys/sysDepart/deleteBatch', {
    params: {
      ids
    }
  })
}
