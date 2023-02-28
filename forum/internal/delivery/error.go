package delivery

import "fmt"

func (h *Handler) onError(err string) string {
	return fmt.Sprintf(`{"error":"%s"}`, err)
}
