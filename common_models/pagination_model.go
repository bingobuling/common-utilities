//author xinbing
//time 2018/10/11 15:07
//
package common_models

type PageModel struct {
	List       interface{} `json:"list" form:"list"`
	Pagination *Pagination `json:"pagination" form:"pagination"`
}

type Pagination struct {
	CurrentPage int `json:"current_page" form:"current_page"`
	PageSize    int `json:"page_size" form:"page_size"`
	LastPage    int `json:"last_page"`
	Total       int `json:"total" form:"total"`
}

func BuildPagination(currentPage, pageSize, totalCount int) *Pagination {
	p := &Pagination{
		CurrentPage: currentPage,
		PageSize:    pageSize,
		Total:       totalCount,
	}
	p.LastPage = p.GetLastPage()
	return p
}

// 计算偏移量、起始下标
func (p *Pagination) Offset() int {
	return (p.GetCurrentPage() - 1) * p.GetPageSize()
}

// 计算当前页的容量
func (p *Pagination) CalCurrCapacity(total int) int {
	z := total - p.Offset()
	if z <= 0 {
		return 0
	} else if z >= p.GetPageSize() {
		return p.GetPageSize()
	}
	return z
}

func (p *Pagination) GetCurrentPage() int {
	if p.CurrentPage <= 0 {
		p.CurrentPage = 1
	}
	return p.CurrentPage
}

func (p *Pagination) GetPageSize() int {
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	return p.PageSize
}

func (p *Pagination) GetLastPage() int {
	if p.Total <= 0 {
		return 1
	}
	return (p.Total + p.GetPageSize()) / p.GetPageSize()
}
