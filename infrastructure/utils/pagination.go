package utils

import (
	"fmt"
	"math"
	"strings"

	"github.com/felixa1996/go-plate/domain"
)

type Pagination struct {
	PerPage     int
	CurrentPage int
	TotalPage   int
	Total       int64
	Offset      int
	Sort        string
	SortList    map[string]string
}

func (p *Pagination) ToMeta() domain.MetaPagination {
	p.getPageCount()
	p.getDefaultSort()
	p.getOffset()

	return domain.MetaPagination{
		PerPage:     p.PerPage,
		CurrentPage: p.CurrentPage,
		TotalPage:   p.TotalPage,
		Total:       p.Total,
		Offset:      p.Offset,
		Sort:        p.Sort,
	}
}

func (p *Pagination) getOffset() {
	p.Offset = (p.CurrentPage - 1) * p.PerPage
}

func (p *Pagination) getPageCount() {
	if p.PerPage < 1 {
		p.PerPage = 1
	} else {
		pageCountFloat := math.Ceil(float64(p.Total) / float64(p.PerPage))
		p.TotalPage = int(pageCountFloat)
	}
}

func (p *Pagination) getDefaultSort() {
	sort := p.getSortBySortList()
	isSortText := "ASC"
	isSortDesc := strings.Contains(p.Sort, "-")
	if isSortDesc {
		isSortText = "DESC"
	}
	p.Sort = fmt.Sprintf("%s %s", sort, isSortText)
}

func (p *Pagination) getSortBySortList() string {
	defaultSort := strings.ReplaceAll(p.Sort, "-", "")
	if len(p.SortList) > 0 {
		return p.SortList[defaultSort]
	}
	return defaultSort
}

type SortList struct {
}
