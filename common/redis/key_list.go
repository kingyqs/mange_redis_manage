//
//	redis List相关的所有操作
//

package redis

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//获取List value
func ListLRANGEALL(rc redis.Conn, keyname string) []interface{} {
	fmt.Println("执行redis : ", "LRANGE", keyname, 0, -1)
	res, err := redis.Values(rc.Do("LRANGE", keyname, 0, -1))
	if err != nil {
		fmt.Println("GET error", err.Error())
	}
	fmt.Println(res)
	return res
}

// LRANGE key start stop
// 返回列表 key 中指定区间内的元素，区间以偏移量 start 和 stop 指定。
func ListLRANGE(rc redis.Conn, keyname string, start, stop int64) (res []interface{}, err error) {
	fmt.Println("执行redis : ", "LRANGE", keyname, start, stop)
	res, err = redis.Values(rc.Do("LRANGE", keyname, start, stop))
	if err != nil {
		fmt.Println("GET error", err.Error())
		return
	}
	fmt.Println(res)
	return
}

//新创建list 将一个或多个值 value 插入到列表 key 的表头
func ListLPUSH(rc redis.Conn, keyname string, values []interface{}) error {
	args := redis.Args{}.Add(keyname)
	for _, value := range values {
		args = args.Add(value)
	}
	fmt.Println("执行redis : ", "LPUSH", args)
	res, err := rc.Do("LPUSH", args...)
	if err != nil {
		fmt.Println("GET error", err.Error())
		return err
	}
	fmt.Println(res)
	return nil
}

// RPUSH key value [value ...]
// 将一个或多个值 value 插入到列表 key 的表尾(最右边)。
// 如果有多个 value 值，那么各个 value 值按从左到右的顺序依次插入到表尾：比如对一个空列表 mylist 执行
// RPUSH mylist a b c ，得出的结果列表为 a b c ，等同于执行命令 RPUSH mylist a 、 RPUSH mylist b 、 RPUSH mylist c 。
//新创建List  将一个或多个值 value 插入到列表 key 的表尾(最右边)。
func ListRPUSH(rc redis.Conn, keyname string, values []interface{}) error {
	args := redis.Args{}.Add(keyname)
	for _, value := range values {
		args = args.Add(value)
	}
	fmt.Println("执行redis : ", "RPUSH", args)
	res, err := rc.Do("RPUSH", args)
	if err != nil {
		fmt.Println("GET error", err.Error())
		return err
	}
	fmt.Println(res)
	return nil
}

// BLPOP key [key ...] timeout
// BLPOP 是列表的阻塞式(blocking)弹出原语。
// 它是 LPOP 命令的阻塞版本，当给定列表内没有任何元素可供弹出的时候，连接将被 BLPOP 命令阻塞，直到等待超时或发现可弹出元素为止。
func ListBLPOP() {}

// BRPOP key [key ...] timeout
// BRPOP 是列表的阻塞式(blocking)弹出原语。
// 它是 RPOP 命令的阻塞版本，当给定列表内没有任何元素可供弹出的时候，连接将被 BRPOP 命令阻塞，直到等待超时或发现可弹出元素为止。
func ListBRPOP() {}

// BRPOPLPUSH source destination timeout
// BRPOPLPUSH 是 RPOPLPUSH 的阻塞版本，当给定列表 source 不为空时， BRPOPLPUSH 的表现和 RPOPLPUSH 一样。
// 当列表 source 为空时， BRPOPLPUSH 命令将阻塞连接，直到等待超时，或有另一个客户端对 source 执行 LPUSH 或 RPUSH 命令为止。
func ListBRPOPLPUSH() {}

// LINDEX key index
// 返回列表 key 中，下标为 index 的元素。
func ListLINDEX(rc redis.Conn, keyname string, index int64) (res string, err error) {
	fmt.Println("执行redis : ", "LINDEX", keyname, index)
	res, err = redis.String(rc.Do("LINDEX", keyname, index))
	if err != nil {
		fmt.Println("GET error", err.Error())
		return
	}
	fmt.Println(res)
	return
}

// LINSERT key BEFORE|AFTER pivot value
// 将值 value 插入到列表 key 当中，位于值 pivot 之前或之后。
// 当 pivot 不存在于列表 key 时，不执行任何操作。
// 当 key 不存在时， key 被视为空列表，不执行任何操作。
// 如果 key 不是列表类型，返回一个错误。
// direction : 方向 bool true:BEFORE(前)    false: AFTER(后)
func ListLINSERT(rc redis.Conn, direction bool, keyname, pivot, value string) (err error) {
	directionStr := "AFTER"
	if direction {
		directionStr = "BEFORE"
	}

	fmt.Println("执行redis : ", "LINSERT", keyname, directionStr, pivot, value)
	res, err := rc.Do("LINSERT", keyname, directionStr, pivot, value)
	if err != nil {
		fmt.Println("GET error", err.Error())
		return
	}
	fmt.Println(res)
	return
}

// LLEN key
// 返回列表 key 的长度。
// 如果 key 不存在，则 key 被解释为一个空列表，返回 0 .
func ListLLEN(rc redis.Conn, keyname string) (res int64, err error) {
	fmt.Println("执行redis : ", "LLEN", keyname)
	res, err = redis.Int64(rc.Do("LLEN", keyname))
	if err != nil {
		fmt.Println("GET error", err.Error())
		return
	}
	fmt.Println(res)
	return
}

// LPOP key
// 移除并返回列表 key 的头元素。
func ListLPOP(rc redis.Conn, keyname string) (res string, err error) {
	fmt.Println("执行redis : ", "LPOP", keyname)
	res, err = redis.String(rc.Do("LPOP", keyname))
	if err != nil {
		fmt.Println("GET error", err.Error())
		return
	}
	fmt.Println(res)
	return
}

// LPUSHX key value
// 将值 value 插入到列表 key 的表头，当且仅当 key 存在并且是一个列表。
// 和 LPUSH 命令相反，当 key 不存在时， LPUSHX 命令什么也不做。
func ListLPUSHX(rc redis.Conn, keyname string, value interface{}) (err error) {
	fmt.Println("执行redis : ", "LPUSHX", keyname, value)
	res, err := rc.Do("LPUSHX", keyname, value)
	if err != nil {
		fmt.Println("GET error", err.Error())
		return
	}
	fmt.Println(res)
	return
}

// LREM key count value
// 根据参数 count 的值，移除列表中与参数 value 相等的元素。
// count 的值可以是以下几种：
// count > 0 : 从表头开始向表尾搜索，移除与 value 相等的元素，数量为 count 。
// count < 0 : 从表尾开始向表头搜索，移除与 value 相等的元素，数量为 count 的绝对值。
// count = 0 : 移除表中所有与 value 相等的值。
func ListLREM(rc redis.Conn, keyname string, count int64, value interface{}) (err error) {
	fmt.Println("执行redis : ", "LREM", keyname, count, value)
	res, err := rc.Do("LREM", keyname, count, value)
	if err != nil {
		fmt.Println("GET error", err.Error())
		return
	}
	fmt.Println(res)
	return
}

// LSET key index value
// 将列表 key 下标为 index 的元素的值设置为 value 。
// 当 index 参数超出范围，或对一个空列表( key 不存在)进行 LSET 时，返回一个错误。
func ListLSET(rc redis.Conn, keyname string, index int64, value interface{}) (err error) {
	fmt.Println("执行redis : ", "LSET", keyname, index, value)
	res, err := rc.Do("LSET", keyname, index, value)
	if err != nil {
		fmt.Println("GET error", err.Error())
		return
	}
	fmt.Println(res)
	return
}

// LTRIM key start stop
// 对一个列表进行修剪(trim)，就是说，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除。
// 举个例子，执行命令 LTRIM list 0 2 ，表示只保留列表 list 的前三个元素，其余元素全部删除。
func ListLTRIM(rc redis.Conn, keyname string, start, stop int64) (err error) {
	fmt.Println("执行redis : ", "LTRIM", keyname, start, stop)
	res, err := rc.Do("LTRIM", keyname, start, stop)
	if err != nil {
		fmt.Println("GET error", err.Error())
		return
	}
	fmt.Println(res)
	return
}

// RPOP key
// 移除并返回列表 key 的尾元素。
func ListRPOP(rc redis.Conn, keyname string) (res string, err error) {
	fmt.Println("执行redis : ", "RPOP", keyname)
	res, err = redis.String(rc.Do("RPOP", keyname))
	if err != nil {
		fmt.Println("GET error", err.Error())
		return
	}
	fmt.Println(res)
	return
}

// RPOPLPUSH source destination
// 命令 RPOPLPUSH 在一个原子时间内，执行以下两个动作：
// 将列表 source 中的最后一个元素(尾元素)弹出，并返回给客户端。
// 将 source 弹出的元素插入到列表 destination ，作为 destination 列表的的头元素。
// 举个例子，你有两个列表 source 和 destination ， source 列表有元素 a, b, c ， destination
// 列表有元素 x, y, z ，执行 RPOPLPUSH source destination 之后， source 列表包含元素 a, b ，
// destination 列表包含元素 c, x, y, z ，并且元素 c 会被返回给客户端。
// 如果 source 不存在，值 nil 被返回，并且不执行其他动作。
// 如果 source 和 destination 相同，则列表中的表尾元素被移动到表头，并返回该元素，可以把这种特殊情况视作列表的旋转(rotation)操作。
func ListRPOPLPUSH(rc redis.Conn, keyname, destination string) (res string, err error) {
	fmt.Println("执行redis : ", "RPOPLPUSH", keyname, destination)
	res, err = redis.String(rc.Do("RPOPLPUSH", keyname, destination))
	if err != nil {
		fmt.Println("GET error", err.Error())
		return
	}
	fmt.Println(res)
	return
}

// RPUSHX key value
// 将值 value 插入到列表 key 的表尾，当且仅当 key 存在并且是一个列表。
func ListRPUSHX(rc redis.Conn, keyname string, value interface{}) (err error) {
	fmt.Println("执行redis : ", "RPUSHX", keyname, value)
	res, err := rc.Do("RPUSHX", keyname, value)
	if err != nil {
		fmt.Println("GET error", err.Error())
		return
	}
	fmt.Println(res)
	return
}
