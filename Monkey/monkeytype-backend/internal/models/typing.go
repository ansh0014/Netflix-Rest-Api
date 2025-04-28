type Typing struct {
    UserID    string    `json:"user_id"`
    Text      string    `json:"text"`
    Timestamp time.Time `json:"timestamp"`
}