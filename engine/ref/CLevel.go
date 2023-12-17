package ref

const (
	CLEVEL_ACTION_LIKE    int = 300
	CLEVEL_ACTION_DISLIKE int = 301
)

type CLevel struct {
	//Main
	Id              int
	Name            string
	Description     string
	Uid             int
	Password        string
	Version         int
	Length          int
	Difficulty      int
	DemonDifficulty int

	SuggestDifficulty    float64
	SuggestDifficultyCnt int

	//Level
	TrackId         int
	SongId          int
	VersionGame     int
	VersionBinary   int
	StringExtra     string
	StringSettings  string
	StringLevel     string
	StringLevelInfo string
	OrigId          int

	//Stats
	Objects        int
	StarsRequested int
	StarsGot       int
	Ucoins         int
	Coins          int
	Downloads      int
	Likes          int
	Reports        int

	//Params
	Is2p       bool
	IsVerified bool
	IsFeatured int
	ISHall     bool
	IsEpic     int
	IsUnlisted int
	IsLDM      bool

	//Dates
	UploadDate string
	UpdateDate string

	UnlockLevelObject bool
}

func (lvl *CLevel) Exists(lvlid int) bool {
	return true
}

func (lvl *CLevel) CountLevels() int {
	return 0
}

func (lvl *CLevel) LoadParams() {}

func (lvl *CLevel) PushParams() {}

func (lvl *CLevel) LoadDates() {}

func (lvl *CLevel) LoadLevel() {}

func (lvl *CLevel) LoadStats() {}

func (lvl *CLevel) LoadMain() {}

func (lvl *CLevel) LoadAll() {}

func (lvl *CLevel) LoadBase() {}

func (lvl *CLevel) IsOwnedBy(uid int) bool {
	return true
}

func (lvl *CLevel) CheckParams() bool {
	if len(lvl.Name) > 32 || len(lvl.Description) > 256 || len(lvl.Password) > 8 || lvl.Version < 1 || lvl.Version > 127 || lvl.TrackId < 0 || lvl.SongId < 0 || lvl.VersionGame < 0 {
		return false
	}
	if lvl.VersionBinary < 0 || len(lvl.StringLevel) < 16 || lvl.OrigId < 0 || (lvl.Objects < 100 && !lvl.UnlockLevelObject) || lvl.StarsRequested < 0 || lvl.StarsRequested > 10 || lvl.Ucoins < 0 || lvl.Ucoins > 3 {
		return false
	}
	return true
}

func (lvl *CLevel) DeleteLevel() {}

func (lvl *CLevel) UploadLevel() int {
	id := 1
	return id | -1
}

func (lvl *CLevel) UpdateLevel() int {
	id := 1
	return id | -1
}

func (lvl *CLevel) UpdateDescription(desc string) bool {
	return true
}

func (lvl *CLevel) DoSuggestDifficulty(diffx int) {}

func (lvl *CLevel) RateLevel(stars int) {}

func (lvl *CLevel) RateDemon(diff int) {}

func (lvl *CLevel) FeatureLevel(featured int) bool {
	return true
}

func (lvl *CLevel) EpicLevel(epic bool) {}

func (lvl *CLevel) LegendaryLevel(legend bool) {}

func (lvl *CLevel) GodlikeLevel(legend bool) {}

func (lvl *CLevel) LikeLevel(lvlid int, uid int, action int) bool {
	action = CLEVEL_ACTION_LIKE | CLEVEL_ACTION_DISLIKE
	return true
}

func (lvl *CLevel) VerifyCoins(verify bool) {}

func (lvl *CLevel) ReportLevel() {}

func (lvl *CLevel) RecalculateCPoints(uid int) {}

func (lvl *CLevel) SendReq(modUid int, stars int, isFeatured int) bool {
	return true
}
