package handlers

type DefaultResponse struct {
	Message string `json:"message"`
}

func NewDefaultResponse(message string) DefaultResponse {
	return DefaultResponse{
		Message: message,
	}
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Error: message,
	}
}

type LyricsResponse struct {
	Lyrics []string `json:"lyrics"`
}
