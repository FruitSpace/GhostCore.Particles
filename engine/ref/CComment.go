package ref

type CComment struct {
	Id         int
	Uid        int
	Likes      int
	PostedTime string
	Comment    string

	LvlId   int
	Percent int
	IsSpam  bool
}

func (cc *CComment) ExistsLevelComment(id int) bool {
	return true
}

func (cc *CComment) ExistsAccComment(id int) bool {
	return true
}

func (cc *CComment) CountAccComments(uid int) int {
	return 0
}

func (cc *CComment) CountLevelComments(lvlId int) int {
	return 0
}

func (cc *CComment) CountCommentHistory(uid int) int {
	return 0
}

func (cc *CComment) LoadAccComment() {}

func (cc *CComment) LoadLevelComment() {}

func (cc *CComment) GetAllAccComments(uid int, page int) []CComment {
	return []CComment{}
}

// sortMode: true=likes, false=postedTime
func (cc *CComment) GetAllLevelComments(lvlId int, page int, sortMode bool) []CComment {
	return []CComment{}
}

func (cc *CComment) GetAllCommentsHistory(uid int, page int, sortMode bool) []CComment {
	return []CComment{}
}

func (cc *CComment) PostAccComment() bool {
	return true
}

func (cc *CComment) PostLevelComment() bool {
	return true
}

func (cc *CComment) DeleteAccComment(id int, uid int) {}

func (cc *CComment) DeleteLevelComment(id int, uid int) {}

func (cc *CComment) DeleteOwnerLevelComment(id int, lvlId int) {}

func (cc *CComment) LikeAccComment(id int, uid int, actionLike bool) bool {
	return true
}

func (cc *CComment) LikeLevelComment(id int, uid int, actionLike bool) bool {
	return true
}
