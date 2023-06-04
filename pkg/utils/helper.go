package utils

import (
	"clean-arch-template/pkg/response"
	"math"
	"strconv"
	"unicode"
)

func OffsetCalculator(page int, pageSize int) int {
	if page <= 0 {
		return 0
	}
	return (page - 1) * pageSize
}

func LimitCalculator(page int, pageSize int) int {
	if page <= 0 {
		return 1 * pageSize
	}
	return page * pageSize
}

func MaxPageCalculator(count int, pageSize int) float64 {
	return math.Ceil(float64(count) / float64(pageSize))
}

func FromPaginationCalculator(from int) int {
	if from <= 0 {
		return 1
	}
	return from
}

func ToPaginationCalculator(to int, lastPage int) int {
	if to > lastPage {
		return lastPage
	}
	return to
}

func CurrentPageCalculator(currentPage int) int {
	if currentPage <= 0 {
		return 1
	}
	return currentPage
}

func GeneratorPaginationResponse(data interface{}, totalData int, currentPage int, perPage int, baseUrl string) *response.PaginationResponse {

	lastPage := strconv.Itoa(int(MaxPageCalculator(int(totalData), perPage)))

	return &response.PaginationResponse{
		Data: data,
		Meta: &response.MetaResponse{
			CurrentPage: CurrentPageCalculator(currentPage),
			From:        FromPaginationCalculator(currentPage - 1),
			LastPage:    int(MaxPageCalculator(int(totalData), perPage)),
			PerPage:     perPage,
			To:          ToPaginationCalculator(currentPage+1, int(MaxPageCalculator(int(totalData), perPage))),
			Total:       int(totalData),
		},
		Links: &response.LinkPaginationResponse{
			First: baseUrl + "?page=1",
			Last:  baseUrl + "?page=" + lastPage,
			Prev:  PrevLinkCalculator(baseUrl, currentPage),
			Next:  NextLinkCalculator(baseUrl, currentPage, int(MaxPageCalculator(int(totalData), perPage))),
		},
		Success: true,
		Message: "Data loaded",
	}
}

func PrevLinkCalculator(baseUrl string, page int) string {
	page = page - 1

	if page <= 0 {
		return ""
	}

	pageString := strconv.Itoa(page)

	return baseUrl + "?page=" + pageString
}

func NextLinkCalculator(baseUrl string, page int, lastPage int) string {
	page = page + 1

	if page > lastPage {
		return ""
	}

	pageString := strconv.Itoa(page + 1)

	return baseUrl + "?page=" + pageString
}

func CapitalizeWord(word string) string {
	var output []rune //create an output slice
	isWord := true
	for _, val := range word {
		if isWord && unicode.IsLetter(val) { //check if character is a letter convert the first character to upper case
			output = append(output, unicode.ToUpper(val))
			isWord = false
		} else if !unicode.IsLetter(val) {
			isWord = true
			output = append(output, val)
		} else {
			output = append(output, val)
		}
	}

	return string(output)
}
