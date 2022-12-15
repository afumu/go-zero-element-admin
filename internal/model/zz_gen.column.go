package model

var SysUserColumn = struct {
	Id          string
	Username    string
	Realname    string
	Password    string
	Salt        string
	Sex         string
	Avatar      string
	Status      string
	Deleted     string
	CreatedBy   string
	CreatedTime string
	UpdatedBy   string
	UpdatedTime string
}{
	Id:          "id",
	Username:    "username",
	Realname:    "realname",
	Password:    "password",
	Salt:        "salt",
	Sex:         "sex",
	Avatar:      "avatar",
	Status:      "status",
	Deleted:     "deleted",
	CreatedBy:   "created_by",
	CreatedTime: "created_time",
	UpdatedBy:   "updated_by",
	UpdatedTime: "updated_time",
}
