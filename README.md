# gcache
set v1 1
set v2 1.1
set v3 "123"
incr v1 #自增1
decr v1 #自减1

get v1 #获取值

type v1 #获取类型（int/float/string）

backup filename #(备份)
recovery filename #（恢复备份）
