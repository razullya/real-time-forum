package models

type Notification struct {
	Username string
	Sender   string
	Type     int
	Message  string
	Checked  bool
}
