package response

type Notfound struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type InternalServer struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type PaginationResponse struct {
	Data    interface{}             `json:"data"`
	Meta    *MetaResponse           `json:"meta"`
	Links   *LinkPaginationResponse `json:"links"`
	Success bool                    `json:"success"`
	Message string                  `json:"message"`
}

type MetaResponse struct {
	CurrentPage int `json:"current_page"`
	From        int `json:"from"`
	LastPage    int `json:"last_page"`
	PerPage     int `json:"per_page"`
	To          int `json:"to"`
	Total       int `json:"total"`
}

type LinkPaginationResponse struct {
	First string
	Last  string
	Prev  string
	Next  string
}
