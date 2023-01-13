// Code generated by goctl. DO NOT EDIT.
package types

type UserAddReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Realname string `json:"realname,optional"`
	Sex      int    `json:"sex"`
	Avatar   string `json:"avatar,optional"`
}

type UserEditReq struct {
	Id       int64  `json:"id"`
	Username string `json:"username,optional"`
	Password string `json:"password,optional"`
	Realname string `json:"realname,optional"`
	Sex      int    `json:"sex,optional"`
	Avatar   string `json:"avatar,optional"`
	Status   int    `json:"status,optional"`
}

type UserResp struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	Realname  string `json:"realname"`
	Password  string `json:"password"`
	Salt      string `json:"salt"`
	Sex       int    `json:"sex"`
	Avatar    string `json:"avatar"`
	Status    int    `json:"status"`
	DeletedAt int    `json:"deleted_at"`
	CreatedBy int64  `json:"created_by"`
	CreatedAt int64  `json:"created_at"`
	UpdatedBy int64  `json:"updated_by"`
	UpdatedAt int64  `json:"updated_at"`
}

type UserListResp struct {
	UserList []*UserResp `json:"userList"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type LoginUserResp struct {
	Id          int64    `json:"id"`
	Username    string   `json:"username"`
	Realname    string   `json:"realname"`
	Password    string   `json:"password"`
	Salt        string   `json:"salt"`
	Sex         int      `json:"sex"`
	Avatar      string   `json:"avatar"`
	Status      int      `json:"status"`
	DeletedAt   int      `json:"deleted_at"`
	CreatedBy   int64    `json:"created_by"`
	CreatedAt   int64    `json:"created_at"`
	UpdatedBy   int64    `json:"updated_by"`
	UpdatedAt   int64    `json:"updated_at"`
	Permissions []string `json:"permissions"`
}

type FormParamId struct {
	Id int64 `form:"id"`
}

type FormParamIds struct {
	Ids string `form:"ids"`
}

type RoleAddReq struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description,optional"`
}

type RoleEditReq struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description,optional"`
}

type RoleResp struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	CreatedBy   int64  `json:"created_by"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedBy   int64  `json:"updated_by"`
	UpdatedAt   int64  `json:"updated_at"`
}

type RoleListResp struct {
	RoleList []*RoleResp `json:"roleList"`
}

type MenuAddReq struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Sort        int    `json:"sort"`
	Icon        int    `json:"icon"`
	Description string `json:"description,optional"`
}

type MenuEditReq struct {
	Id          int64  `json:"id"`
	Url         string `json:"url"`
	Sort        int    `json:"sort"`
	Icon        int    `json:"icon"`
	Description string `json:"description,optional"`
}

type MenuResp struct {
	Id          int64       `json:"id"`
	ParentId    int64       `json:"parentId"`
	Name        string      `json:"name"`
	Url         string      `json:"url"`
	Sort        int         `json:"sort"`
	Icon        string      `json:"icon"`
	Description string      `json:"description"`
	CreatedBy   int64       `json:"createdBy"`
	CreatedAt   string      `json:"createdAt"`
	UpdatedBy   int64       `json:"updatedBy"`
	UpdatedAt   string      `json:"updatedAt"`
	Children    []*MenuResp `json:"children"`
}

type MenuListResp struct {
	MenuList []*MenuResp `json:"menuList"`
}

type UserMenuResp struct {
	Permissions []string    `json:"permissions"`
	Menus       []*MenuResp `json:"menus"`
}

type PermissionAddReq struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status int    `json:"sort"`
}

type PermissionEditReq struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status int    `json:"sort"`
}

type PermissionResp struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Status    int    `json:"sort"`
	CreatedBy int64  `json:"created_by"`
	CreatedAt int64  `json:"created_at"`
	UpdatedBy int64  `json:"updated_by"`
	UpdatedAt int64  `json:"updated_at"`
}

type PermissionListResp struct {
	PermissionList []*PermissionResp `json:"permissionList"`
}

type DictAddReq struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type DictEditReq struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type DictResp struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	CreatedBy   int64  `json:"created_by"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedBy   int64  `json:"updated_by"`
	UpdatedAt   int64  `json:"updated_at"`
}

type DictListResp struct {
	DictList []*DictResp `json:"DictList"`
}
