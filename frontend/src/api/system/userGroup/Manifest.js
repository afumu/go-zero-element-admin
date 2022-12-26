import request from '@/utils/request'

export function queryPermissionRule(data) {
    return request.get('/sys/permission/queryPermissionRule', {params: data})
}

export function permission_field(data) {
    return request.get('/sys/dict/getDictItems/permission_field', {params: data})
}

export function queryByField(data) {
    return request.get('/system/userGroup/queryByField', {params: data})
}

export function editPermissionRule(data) {
    return request.put('/sys/permission/editPermissionRule', data)
  }

// 删除
export function deleteId (id) {
    return request.delete('/sys/permission/deletePermissionRule', {params: {
      'id': id
    }
    })
  }

  // 新建
export function addPermissionRule (data) {
    return request.post('/sys/permission/addPermissionRule', data)
  }