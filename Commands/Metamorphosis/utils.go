package main

type MsgU struct {
	Id        string `json:"id"`
	LangCode  string `json:"lang_code"`
	Content   string `json:"content"`
	ChangeLog string `json:"change_log"`
}
