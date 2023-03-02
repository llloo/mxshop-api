package response

import (
	"fmt"
	"time"
)

type JsonDateFormat uint64

func (j JsonDateFormat) MarshalJSON() ([]byte, error) {
	if j == 0 {
		return []byte("\"\""), nil
	}
	var d = fmt.Sprintf("\"%s\"", time.Unix(int64(j), 0).Format("2006-01-01"))
	return []byte(d), nil
}

type UserResponse struct {
	Id       uint32         `json:"id"`
	Mobile   string         `json:"mobile"`
	NickName string         `json:"nickName"`
	Avatar   string         `json:"avatar"`
	Birthday JsonDateFormat `json:"birthday"`
	Gender   uint32         `json:"gender"`
	Role     uint32         `json:"role"`
}
