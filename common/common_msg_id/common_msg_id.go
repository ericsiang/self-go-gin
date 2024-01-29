package common_msg_id

type MsgId uint32

const (
	Success MsgId = iota
	Fail
	Token_Expires
	Token_Invalid
	No_Content
	Duplicate_Entry
	Rule_Not_Allow
)
