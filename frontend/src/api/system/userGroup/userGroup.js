import request from '@/utils/request'

// 查询
export function fetchList(data) {
    return request.get('/system/userGroup/list', {params: data})
}

// 新建
export function create(data) {
    return request.post('/system/userGroup/add',  data)
}

// 修改
export function updateById(data) {
    return request.put('/system/userGroup/edit', data)
}

// 删除
export function deleteById(id) {
    return request.delete('/system/userGroup/delete', {params: {'id': id}})
}
