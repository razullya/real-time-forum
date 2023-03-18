package delivery

type Error struct {
	Status int
	Text   string
}

func (h *Handler) onError(text string, status int) Error {
	return Error{
		Status: status,
		Text:   text,
	}
}
