package common

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

type Paginate struct{ page, perpage, min, max int }

func NewPaginate[T string | int | float64](page, perPage T, minmax ...int) *Paginate {

	const MAX = 100
	const MIN = 10

	var (
		pageInt, perPageInt int
	)

	switch any(*new(T)).(type) {
	case string:
		_page, _ := any(page).(string)
		_perpage, _ := any(perPage).(string)

		pageInt, _ = strconv.Atoi(_page)
		perPageInt, _ = strconv.Atoi(_perpage)

	case int:
		pageInt, _ = any(page).(int)
		perPageInt, _ = any(perPage).(int)

	case float64:
		_page, _ := any(page).(float64)
		_perpage, _ := any(perPage).(float64)

		pageInt = int(_page)
		perPageInt = int(_perpage)

	}

	var min, max int
	{
		if len(minmax) == 2 {
			min = minmax[0]
			max = minmax[1]
		} else {
			min = MIN
			max = MAX
		}
	}

	return &Paginate{
		page:    transformPage(pageInt),
		perpage: transformPerPage(perPageInt, min, max),
		min:     min,
		max:     max,
	}
}

func NewPaginateEchoWithContext(c echo.Context, minmax ...int) *Paginate {
	query := c.QueryParam
	return NewPaginate(query("page"), query("perpage"), minmax...)
}

func (p *Paginate) Page() int {
	return p.page
}

func (p *Paginate) PerPage() int {
	return p.perpage
}

func (p *Paginate) Offset() int {
	return p.perpage * (p.page - 1)
}

func (p *Paginate) Limit() int {
	return p.perpage
}

func transformPerPage(perPage, min, max int) int {

	if perPage < 1 {
		perPage = min
	} else if perPage > max {
		perPage = max
	}

	return perPage
}

func transformPage(page int) int {

	if page < 1 {
		page = 1
	}

	return page
}
