package errx

// OK 成功返回
const OK uint32 = 200

// 全局错误码 前3位代表业务,后三位代表具体功能

const ServerCommonError uint32 = 100001
const RequestParamError uint32 = 100002
const TokenExpireError uint32 = 100003
const TokenGenerateError uint32 = 100004
const DbError uint32 = 100005
const DbUpdateAffectedError uint32 = 100006
const CopyerError uint32 = 100007
