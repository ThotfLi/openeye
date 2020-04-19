package session

type Error struct{
	Err string
}

func (e Error) Error() string {
	return e.Err
}

var (
SESSION_EXPIRATION = Error{Err:"Session 已过期"}
SESSION_ISNOTFIND = Error{Err:"Session 在数据库中找不到"}
SESSION_ISEXISTENCE = Error{Err:"Session 已存在于SesManger中"}

)