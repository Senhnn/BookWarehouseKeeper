package BWKErr

type ErrorCode = int32

// 错误码
const (
	SUCCESS       = iota // 成功
	PARAM_INVALID        // 参数无效

	MYSQL_ERR_UPDATE_BOOK            // 更新mysql数据失败
	MYSQL_ERR_BOOK_NOT_FOUND         // 书籍不存在
	MYSQL_ERR_BOOK_HAS_EXIST         // 书籍已经存在
	MYSQL_ERR_CREATE_BOOK_OR_BOOKNUM // 创建失败
	MYSQL_ERR_REMOVE_BOOK            // 删除book表数据失败
	MYSQL_ERR_REMOVE_BOOKNUM         // 删除booknum表数据失败
	MYSQL_ERR_GET_ALL_BOOK           // 查询所有书籍失败
	MYSQL_GET_AUTHOR_BOOKS_FAIL      // 查询一个作者的所有书籍失败

	REDIS_ERR_DEL_DATA        // 删除数据失败
	REDIS_GET_BOOK_CACHE_FAIL // redis获取书缓存失败
	REDIS_SET_BOOK_CACHE_FAIL // redis设置书缓存失败
	REDIS_DEL_BOOK_CACHE_FALI // redis删除书缓存失败

	JSON_MARSHAL_ERROR   // JSON序列化失败
	JSON_UNMARSHAL_ERROR // JSON反序列化失败

	BOOK_NUM_NOT_ENOUGH // 书籍数量不足
)

var ErrMap = map[int32]string{
	SUCCESS:                          "Grpc call success",
	MYSQL_ERR_UPDATE_BOOK:            "更新书籍失败",
	MYSQL_ERR_BOOK_NOT_FOUND:         "书籍不存在",
	MYSQL_ERR_BOOK_HAS_EXIST:         "书籍已经存在",
	MYSQL_ERR_CREATE_BOOK_OR_BOOKNUM: "创建书籍失败",
	MYSQL_ERR_REMOVE_BOOK:            "删除书籍失败",
	MYSQL_ERR_REMOVE_BOOKNUM:         "删除书籍数量失败",
	MYSQL_ERR_GET_ALL_BOOK:           "查询所有书籍失败",
	REDIS_ERR_DEL_DATA:               "Redis删除失败",
	JSON_MARSHAL_ERROR:               "Json序列化失败",
	BOOK_NUM_NOT_ENOUGH:              "书籍数量不足",
}
