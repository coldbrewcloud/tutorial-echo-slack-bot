package slack

type APIResponse struct {
	OK      bool   `json:"ok"`
	Error   string `json:"error,omitempty"`
	Warning string `json:"warning,omitempty"`
}

type RTMStartResponse struct {
	APIResponse
	URL  string `json:"url,omitempty"`
	Self struct {
		ID string `json:"id"`
	} `json:"self,omitempty"`
}

type RTMMessage struct {
	Type  string `json:"type"`
	Error struct {
		Code    int    `json:"code,omitempty"`
		Message string `json:"msg,omitempty"`
	} `json:"error,omitempty"`
	Channel string `json:"channel,omitempty"`
	User    string `json:"user,omitempty"`
	Text    string `json:"text,omitempty"`
}
