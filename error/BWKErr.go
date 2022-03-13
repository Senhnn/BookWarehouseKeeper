package BWKErr

type ErrorCode = int32

// 错误码
const (
	SUCCESS = iota
	UPDATE_MYSQL_ERR
)

var ErrMap = map[int32]string{
	SUCCESS: "Grpc call success",
}
