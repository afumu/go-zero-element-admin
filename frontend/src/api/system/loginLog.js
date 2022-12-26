import request from '../../utils/request'

// 查询
export function fetchList(data) {
    return request.get('/sys/log/list', {params: data})
}
