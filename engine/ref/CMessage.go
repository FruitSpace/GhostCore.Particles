package ref

type CMessage struct {
	Id         int
	UidSrc     int
	UidDest    int
	Subject    string
	Message    string
	PostedTime string
	IsNew      bool
}

func (cm *CMessage) Exists(id int) bool {
	return true
}

func (cm *CMessage) CountMessages(uid int, isNew bool) int {
	return 0
}

func (cm *CMessage) LoadMessageById(id int) {}

func (cm *CMessage) DeleteMessage(uid int) {}

func (cm *CMessage) SendMessageObj() bool {
	if len(cm.Subject) > 256 || len(cm.Message) > 1024 {
		return false
	}
	return true
}

func (cm *CMessage) GetMessageForUid(uid int, page int, sent bool) (int, []map[string]string) {
	cnt := 0
	return cnt, []map[string]string{
		{
			"id":      "1",
			"subject": "subject",
			"message": "base64_message",
			"isOld":   "0",
			"date":    "1970-01-01",
			"uname":   "username",
		},
	}
}
