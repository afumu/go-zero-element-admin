import request from '../../utils/request'

// 查询
export function fetchList(data) {
    return request.get('/sys/position/list', {params: data})
}
export function position_rank(data) {
    return request.get('/sys/dict/getDictItems/position_rank', {params: data})
}

// 删除
export function deleteById(id) {
    return request.delete('/sys/position/delete', {params: {'id': id}})
}

// 新建
export function create(data) {
    return request.post('/sys/position/add', data)
}

// 修改
export function updateById(data) {
    return request.put('/sys/position/edit', data, {
        trim: true
    })
}
