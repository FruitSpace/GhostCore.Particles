package ref

type CFriendship struct{}

func (cf *CFriendship) IsAlreadyFriend(uid_dest int, uid int) bool {
	return true
}

func (cf *CFriendship) IsAlreadySentFriend(uid_dest int, uid int) bool {
	return true
}

func (cf *CFriendship) CountFriendRequests(uid int, new bool) int {
	return 0
}

func (cf *CFriendship) GetFriendRequests(uid int, page int, sent bool) (int, []map[string]string) {
	//count, users
	return 0, []map[string]string{
		{
			"id":            "1",
			"comment":       "base64 text",
			"uid":           "1",
			"uname":         "username",
			"isNew":         "1",
			"special":       "0",
			"iconType":      "0",
			"clr_primary":   "0",
			"clr_secondary": "0",
			"iconId":        "0",
			"date":          "1970-01-01",
		},
	}
}

func (cf *CFriendship) GetFriendRequestsCount(uid int, sent bool) int {
	return 0
}

func (cf *CFriendship) DeleteFriendship(uid int, uid_dest int) {}

func (cf *CFriendship) GetFriendshipId(uid int, uid_dest int) int {
	return 0
}

func (cf *CFriendship) GetFriendByFID(id int) (int, int) {
	return 1, 2
}

func (cf *CFriendship) GetAccFriends(acc CAccount) []int {
	return []int{}
}

func (cf *CFriendship) ReadFriendRequest(id int) {}

func (cf *CFriendship) RequestFriend(uid int, uidDest int, comment string) int {
	return 1 | -1
}

func (cf *CFriendship) AcceptFriendRequest(id int, uid int) int {
	return 1 | -1
}

func (cf *CFriendship) RejectFriendRequestById(id int, uid int) int {
	return 1 | -1
}

func (cf *CFriendship) RejectFriendRequestByUid(uid int, uid_dest int, isSender bool) {}
