package result

const (
	Success = 200
	Error   = 500
)

func GetErrMap() map[uint]string {
	return map[uint]string{
		Success: "请求成功",
		Error:   "请求失败",
	}
}
