package gcache

const (
	/**
	 * gconnect
	 */
	DELIMITER   = '\000' //分割符号4
	IO_BUFF_END = '\n'
	/**
	 * gsever & gclient
	 */
	SEVER_ADDRESS            = "127.0.0.1:4396" //服务器监听端口
	CLIENT_CONNCET_WAIT_TIME = 2                //客户端连接等待时间
	INIT_GDATA_NUM           = 100              //数据库开始时的数据个数
	MAX_CONNECT              = 100              //最大连接数
	/*
	 *字符串常量
	 */
	STR_FAIL = "fail"
	STR_SUCC = "success"

	/*
	 *gdata
	 */
	TYPE_INT    = 1
	TYPE_FLOAT  = 2
	TYPE_STRING = 3

	/*
	 *request command
	 */
	SET      = 1
	GET      = 2
	DELETE   = 3
	TYPE     = 4
	INCR     = 5
	DECR     = 6
	BACKUP   = 7
	RECOVERY = 8
	/*
	 *respone command
	 */
	RESPONSE_SUCC  = 255
	RESPONSE_ERR   = 254
	RESPONSE_VALUE = 253

	/*
	 * file folder
	 */
	BACKUP_FOLDER = "backup"
)
