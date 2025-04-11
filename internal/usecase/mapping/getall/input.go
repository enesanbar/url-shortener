package getall

import "strconv"

const (
	FirstPage       = 1
	DefaultPageSize = 16
	MaxPageSize     = 64
)

// Request data
type Request struct {
	Page      int64 `default:"1"`
	PageSize  int64 `maximum:"64" default:"16"`
	SortBy    string
	SortOrder string
	CodeQuery string
	URLQuery  string
}

func NewRequest(page string, pageSize string, sortBy string, sortOrder string, codeQuery, urlQuery string) *Request {
	p, err := strconv.Atoi(page)
	if err != nil {
		p = FirstPage
	}

	ps, err := strconv.Atoi(pageSize)
	if err != nil {
		ps = DefaultPageSize
	}

	if ps > MaxPageSize {
		ps = MaxPageSize
	}

	if sortOrder == "" {
		sortOrder = "desc"
	}

	return &Request{
		Page:      int64(p),
		PageSize:  int64(ps),
		SortBy:    sortBy,
		SortOrder: sortOrder,
		CodeQuery: codeQuery,
		URLQuery:  urlQuery,
	}
}
