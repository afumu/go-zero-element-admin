import request from '@/utils/request'

// 查询
export function fetchList (data) {
  return request.get('/sys/role/list', { params: data })
}

// 查询所有
export function fetchAll () {
  return request.get('/system/role/all')
}

// 新建
export function create (data) {
  return request.post('/sys/role/add', data, {
    trim: true
  })
}

// 修改
export function updateById (data) {
  return request.put('/sys/role/edit', data, {
    trim: true
  })
}

// 删除
export function deleteById (id) {
  return request.delete(`/sys/role/delete?id=${id}`)
}

// 批量删除
export function deleteByIdInBatch (ids) {
  return request.delete('/sys/user/deleteUserRoleBatch', data)
}

// 配置权限
export function createRolePermission (data) {
  return request.post('/system/role/createRolePermission', data)
}

// 配置菜单
export function createRoleMenu (data) {
  return request.post('/sys/permission/saveRolePermission', data)
}
