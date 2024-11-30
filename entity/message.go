package entity

// Message represents a message sent by a client.
type Message struct {
	// Message content
	Content string `json:"content"`
	// Message sender name
	SenderName string `json:"sender"`
	// Message sent time
	SentTime string `json:"sentTime"`
}
