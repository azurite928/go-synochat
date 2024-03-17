package synochat

type Message struct {
	Text    string `json:"text"`
	FileUrl string `json:"file_url"`
}
