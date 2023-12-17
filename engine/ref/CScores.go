package ref

const (
	CSCORE_TYPE_FRIENDS int = 400
	CSCORE_TYPE_TOP     int = 401
	CSCORE_TYPE_WEEK    int = 402
)

type CScores struct {
	Id         int
	Uid        int
	LvlId      int
	PostedTime string
	Percent    int
	Ranking    int
	Attempts   int
	Coins      int
}

func (cs *CScores) ScoreExistsByUid(uid int, lvlId int) bool {
	return true
}

func (cs *CScores) LoadScoreById() {}

func (cs *CScores) GetScoresForLevelId(lvlId int, types int, acc CAccount) []CScores {
	types = CSCORE_TYPE_FRIENDS | CSCORE_TYPE_TOP | CSCORE_TYPE_WEEK
	return []CScores{}
}

func (cs *CScores) UpdateLevelScore() {}

func (cs *CScores) UploadLevelScore() {}
