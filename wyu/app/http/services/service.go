package services

import (
	"math"
	"strconv"
	"wyu/modules"
)

type Services interface {
	Paginator(strPage string, strSize string) map[string]interface{}
}

const(
	defaultPage int = 1
	defaultSize int = 10
)

var (

)

type Service struct {
	Cache modules.Rd
}

func NewService() *Service {
	return &Service{
		Cache: modules.InstanceRedis(),
	}
}

func (srv *Service) Paginator(page int, pageNums int, pageSize ...int) map[string]interface{} {
	var nums int = pageNums
	var size int
	var prePage int
	var sufPage int

	if len(pageSize) > 0 {
		size = pageSize[0]
	} else {
		size = defaultSize
	}

	var totalPage int = int(math.Ceil(float64(nums) / float64(size)))

	if page > totalPage {
		page = totalPage
	}

	if page <= 0 {
		page = 1
	}

	var pages []int

	switch {
	case page >= totalPage-5 && totalPage > 5:
		start := totalPage-5+1
		prePage = page-1
		sufPage = int(math.Min(float64(totalPage), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start+i
		}

	case page >= 3 && totalPage > 5:
		start := page-3+1
		pages = make([]int, 5)
		prePage = page-3
		for i, _ := range pages {
			pages[i] = start + i
		}

		prePage = page-1
		sufPage = page+1

	default:
		pages = make([]int, int(math.Min(5, float64(totalPage))))
		for i, _ := range pages {
			pages[i] = i + 1
		}

		prePage = int(math.Max(float64(1), float64(page-1)))
		sufPage = page+1
	}

	paginator := map[string]interface{}{
		"pages": pages,
		"total": totalPage,
		"pre_page": prePage,
		"suf_page": sufPage,
		"cur_page": page,
	}

	return paginator
}

func (srv *Service) PaginatorParams(strPage string, strSize string) (int, int) {
	page, err := strconv.Atoi(strPage)
	if err != nil {
		page = defaultPage
	}

	size, err := strconv.Atoi(strSize)
	if err != nil {
		size = defaultSize
	}

	return page, size
}
