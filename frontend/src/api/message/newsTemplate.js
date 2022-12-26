import request from '@/utils/request'
// 选择用户 - 新建通知

// 查询用户
export function fetchList(data) {
    return request.get('/message/sysMessageTemplate/list', {
        params: data
    })
}
export function template_type(data) {
    return request.get('/sys/dict/getDictItems/template_type', {
        params: data
    })
}

// 新建
export function create(data) {
    return request.post('/message/sysMessageTemplate/add', data, {
        trim: true
    })
}

// 修改
export function updateById(data) {
    return request.put('/message/sysMessageTemplate/edit', data, {
        trim: true
    })
}

// 删除
export function deleteById(id) {
    return request.delete('/message/sysMessageTemplate/delete', {
        params: {
            'id': id
        }
    })
}