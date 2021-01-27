package main

type GetMeResult struct {
	Id                          int    `json:"id"`
	Is_bot                      bool   `json:"is_bot"`
	First_name                  string `json:"first_name"`
	Username                    string `json:"username"`
	Can_join_groups             bool   `json:"can_join_groups"`
	Can_read_all_group_messages bool   `json:"can_read_all_group_messages"`
	Supports_inline_queries     bool   `json:"supports_inline_queries"`
}

type GetMeResponse struct {
	Ok     bool        `json:"ok"`
	Result GetMeResult `json:"result"`
}

type Update struct {
	Update_id int     `json:"update_id"`
	Message   Message `json:"message"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	ChatID int `json:"id"`
}

type Respond struct {
	Update []Update `json:"result"`
}

type BotMessage struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}
