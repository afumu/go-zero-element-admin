import request from '@/utils/request'

// 查询
export function fetchList (data) {
  return request.get('/sys/user/list', { params: data })
}

// 获取所有的角色
export function queryAllRole () {
  return request.get('/sys/role/queryall')
}

// 获取所有的公司
export function queryAllCompany () {
  return request.get('/sys/company/getDropdownList')
}

// 根据公司Id获取党组织
export function queryPartyByCompanyId (data) {
  return request.get('/sys/party/queryIdTree', { params: data })
}

// 获取所有的职位
export function queryAllPosition (data) {
  return request.get('/sys/position/list', { params: data })
}

// 获取所有的职位
export function queryDepartByCompanyId (data) {
  return request.get('sys/sysDepart/queryIdTree', { params: data })
}

// 查询用户有哪些角色权限
export function queryUserRoleByUserId (data) {
  return request.get('/sys/user/queryUserRole', { params: data })
}

// 查询用户部门
export function queryUserDepartByUserId (data) {
  return request.get('/sys/user/userDepartList', { params: data })
}




// 新建
export function create (data) {
  return request.post('/sys/user/add', data, {
    trim: true
  })
}

// 修改
export function updateById (data) {
  return request.put('/sys/user/edit', data, {
    trim: true
  })
}

// 彻底删除用户，不做逻辑删除了
export function deleteById (id) {
  return request.delete(`/sys/user/delete?id=${id}`)
}

// 批量删除
export function deleteByIdInBatch (ids) {
  return request.delete('/sys/user/deleteBatch', {
    params: {
      ids
    }
  })
}

// 配置用户角色
export function createUserRole (data) {
  return request.post('/system/user/createUserRole', data)
}

// 重置密码
export function changePassword (data) {
  return request.put('/sys/user/changePassword', data)
}

// 冻结用户
export function frozenBatch (data) {
  return request.put('/sys/user/frozenBatch', data)
}

// 查询回收站的用户
export function recycleBin (data) {
  return request.get('/sys/user/recycleBin', data)
}

// 还原用户
export function putRecycleBin (data) {
  return request.put('/sys/user/putRecycleBin', data)
}

// 彻底删除用户
export function deleteRecycleBin (userIds) {
  return request.delete(`/sys/user/deleteRecycleBin?userIds=${userIds}`)
}


// 获取所有的部门
export function queryAllDepart () {
  return request.get('/sys/sysDepart/queryIdTree')
}