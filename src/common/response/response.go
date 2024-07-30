package respons

func Paginate(data any, limit int, page int) map[string]interface{} {
	message := "success"
	response := map[string]interface{}{
		"message": message,
		"data":    data,
		"meta": map[string]interface{}{
			"page":  page,
			"limit": limit,
		},
	}
	return response
}

func ResponseSuccess(data any) map[string]interface{} {
	response := map[string]interface{}{
		"message": "success",
		"data":    data,
		"meta":    map[string]interface{}{},
	}
	return response
}
