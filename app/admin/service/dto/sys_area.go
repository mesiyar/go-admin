package dto


// SysAreaReq 功能获取请求参数
type SysAreaReq struct {
	ParAreaCode       string `form:"parAreaCode" search:"type:exact;column:par_area_code;table:sys_area" comment:""`     //路径
}


func (s *SysAreaReq) GetByAreaCode() interface{} {
	return *s
}


