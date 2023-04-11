package delivery

type Error struct {
	Status int    `json:"status,omitempty"`
	Text   string `json:"text,omitempty"`
}

func (h *Handler) onError(text string, status int) Error {
	return Error{
		Status: status,
		Text:   text,
	}
}
