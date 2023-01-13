import request from '@/utils/request'

// 查询
export function fetchTree (data) {
  return request.get('/sys/menu/list', data)
}

// 新建
export function create (data) {
  return request.post('/sys/menu/add', data)
}

// 修改
export function updateById (data) {
  return request.put('/sys/menu/edit', data)
}


// 删除
export function deleteById (id) {
  return request.delete(`/sys/menu/removeById?id=${id}`)
}

// 批量删除
export function deleteByIdInBatch (ids) {
  return request.get('/sys/menu/removeByIds', {
    params: {
      ids
    }
  })
}

// 修改状态
export function updateStatus (data) {
  return request.post('/system/menu/updateStatus', data)
}


// 查询菜单树
export function fetchMenuTree () {
  return request.get('/sys/role/queryTreeList')
}

// 排序
export function sort (data) {
  return request.post('/system/menu/updateSort', data)
}

// 查询当前菜单
export function queryRolePermission (data) {
  return request.get('/sys/permission/queryRolePermission', { params: data })
}
