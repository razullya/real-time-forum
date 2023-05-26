package models

type Notification struct {
	Username string `json:"username,omitempty"`
	Sender   string `json:"sender,omitempty"`
	Message  string `json:"message,omitempty"`
	Checked  bool   `json:"checked,omitempty"`
}
