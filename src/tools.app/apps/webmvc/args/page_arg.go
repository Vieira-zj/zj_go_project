package args

import (
	"fmt"
)

// PageArg 请求数据的公共属性
type PageArg struct {
	Pagefrom int    `json:"pagefrom" form:"pagefrom"`
	Pagesize int    `json:"pagesize" form:"pagesize"`
	Asc      string `json:"asc" form:"asc"`
	Desc     string `json:"desc" form:"desc"`
	// Kword    string `json:"kword" form:"kword"`
	// Stat     int       `json:"stat" form:"stat"`
	// Datefrom time.Time `json:"datafrom" form:"datafrom"`
	// Dateto   time.Time `json:"dateto" form:"dateto"`
	// Total int64 `json:"total" form:"total"`
}

// GetPageSize 获得分页大小
func (p *PageArg) GetPageSize() int {
	if p.Pagesize == 0 {
		return 100
	}
	return p.Pagesize
}

// GetPageFrom 获得分页当前第几页
func (p *PageArg) GetPageFrom() int {
	if p.Pagefrom < 0 {
		return 0
	}
	return p.Pagefrom
}

// GetOrderBy 获得排序 ID DESC, 前端传递参数 desc=排序字段 或者 asc=排序字段
func (p *PageArg) GetOrderBy() string {
	if len(p.Asc) > 0 {
		return fmt.Sprintf(" %s asc", p.Asc)
	} else if len(p.Desc) > 0 {
		return fmt.Sprintf(" %s desc", p.Desc)
	} else {
		return ""
	}
}
