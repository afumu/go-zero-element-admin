import request from '../../utils/request'

// 查询
export function fetchList(data) {
    return request.get('sys/docTemplate/list', {params: data})
}

// 删除
export function deleteById(id) {
    return request.delete('/sys/docTemplate/delete', {params: {'id': id}})
}

// 新建
export function create(data) {
    return request.post('/sys/docTemplate/add', data)
}

// 修改
export function updateById(data) {
    return request.put('/sys/docTemplate/edit', data, {
        trim: true
    })
}