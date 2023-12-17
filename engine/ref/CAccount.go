package ref

const (
	CBAN_BAN           int = 44
	CBAN_UNBAN         int = 53
	CBLACKLIST_BLOCK   int = 62
	CBLACKLIST_UNBLOCK int = 71
	CFRIENDSHIP_ADD    int = 37
	CFRIENDSHIP_REMOVE int = 38
)

type CAccount struct {
	//Main/Auth
	Uid      int
	Uname    string
	Passhash string
	GjpHash  string
	Email    string
	Role_id  int
	IsBanned int

	//Stats
	Stars         int
	Diamonds      int
	Coins         int
	UCoins        int
	Demons        int
	CPoints       int
	Orbs          int
	Moons         int
	Special       int
	LvlsCompleted int

	//Technical
	RegDate    string
	AccessDate string
	LastIP     string
	GameVer    string

	//Social
	Blacklist     string
	FriendsCount  int
	FriendshipIds string

	//Vessels
	IconType       int
	ColorPrimary   int
	ColorSecondary int
	Cube           int
	Ship           int
	Ball           int
	Ufo            int
	Wave           int
	Robot          int
	Spider         int
	Swing          int
	Jetpack        int
	Trace          int
	Death          int

	//Chests
	ChestSmallCount int
	ChestSmallTime  int
	ChestBigCount   int
	ChestBigTime    int

	//Settings
	FrS     int
	CS      int
	MS      int
	Youtube string
	Twitch  string
	Twitter string
}

func (acc *CAccount) CountUsers() int {
	return 0
}

func (acc *CAccount) Exists(uid int) bool {
	return true
}

func (acc *CAccount) SearchUsers(sterm string) int {
	return 0
}

func (acc *CAccount) LoadSettings() {}

func (acc *CAccount) PushSettings() {}

func (acc *CAccount) LoadChests() {}

func (acc *CAccount) PushChests() {}

func (acc *CAccount) LoadVessels() {}

func (acc *CAccount) PushVessels() {}

func (acc *CAccount) LoadStats() {}

func (acc *CAccount) PushStats() {}

func (acc *CAccount) LoadTechnical() {}

func (acc *CAccount) LoadSocial() {}

func (acc *CAccount) LoadAll() {}

func (acc *CAccount) GetUIDByUname(uname string, autoSave bool) int {
	return 0
}

func (acc *CAccount) GetUnameByUID(uid int) string {
	return ""
}

func (acc *CAccount) UpdateIP(ip string) {}

func (acc *CAccount) UpdateGJP2(gjp2 string) {}

func (acc *CAccount) CountIPs(ip string) int {
	return 0
}

func (acc *CAccount) UpdateBlacklist(action int, uid int) {
	action = CBLACKLIST_BLOCK | CBLACKLIST_UNBLOCK
}

func (acc *CAccount) UpdateFriendships(action int, uid int) int {
	action = CFRIENDSHIP_ADD | CFRIENDSHIP_REMOVE
	return 1 | -1
}

func (acc *CAccount) GetShownIcon() int {
	return 0
}

func (acc *CAccount) GetLeaderboardRank() int {
	return 0
}

func (acc *CAccount) UpdateRole(role_id int) {}

func (acc *CAccount) GetRoleObj(fetchPrivs bool) Role {
	return Role{}
}

func (acc *CAccount) UpdateAccessTime() {}

func (acc *CAccount) BanUser(action int) {
	action = CBAN_BAN | CBAN_UNBAN
}

// Role
type Role struct {
	RoleName     string
	CommentColor string
	ModLevel     int
	Privs        map[string]int
}
