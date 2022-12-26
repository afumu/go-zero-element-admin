import request from '@/utils/request'

// 查询
export function fetchList (data) {
  return request.get('/sys/user/userRoleList', { params: data })
}

// // 查询所有
// export function fetchAll () {
//   return request.get('/system/role/all')
// }

// // 新建
// export function create (data) {
//   return request.post('/sys/role/add', data, {
//     trim: true
//   })
// }

// // 修改
// export function updateById (data) {
//   return request.put('/sys/role/edit', data, {
//     trim: true
//   })
// }

// 删除
export function deleteById (obj) {
  return request.delete(`/sys/user/deleteUserRole?roleId=${obj.roleId}&userId=${obj.userId}`)
}

// // 批量删除
export function deleteByIdInBatch (data) {
  return request.delete(`/sys/user/deleteUserRoleBatch?roleId=${data.roleId}&userIds=${data.userIds}`)
}

// 添加
export function addSysUserRole (data) {
  return request.post('/sys/user/addSysUserRole', data, {
    trim: true
  })
}
