package ref

type CQuests struct{}

func (cq *CQuests) Exists(cType int) bool {
	cType = -1 | -2 | -3
	return true
}

func (cq *CQuests) GetDaily() (int, int) {
	var id, lvlId int
	return id, lvlId
}

func (cq *CQuests) GetWeekly() (int, int) {
	var id, lvlId int
	return id + 100001, lvlId
}

func (cq *CQuests) GetEvent() (int, int) {
	var id, lvlId int
	return id, lvlId
}

func (cq *CQuests) PushLevel(lvlId, cType int) int {
	return 0
}
