//
//	redis String相关的所有操作
//

package redis

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	_ "github.com/mangenotwork/mange_redis_manage/common/manlog"
)

//获取String value
func StringGet(rc redis.Conn, keyname string) string {
	fmt.Println("执行redis : ", "GET", keyname)
	res, err := redis.String(rc.Do("GET", keyname))
	if err != nil {
		fmt.Println("GET error", err.Error())
	}
	fmt.Println(res)
	return res
}

//新建String
func StringSET(rc redis.Conn, keyname string, value interface{}) error {
	fmt.Println("执行redis : ", "SET", keyname, value)
	res, err := rc.Do("SET", keyname, value)
	if err != nil {
		fmt.Println("SET error", err.Error())
		return err
	}
	fmt.Println(res)
	return nil
}

//新建String 含有时间
func StringSETEX(rc redis.Conn, keyname string, ttl int64, value interface{}) error {
	fmt.Println("执行redis : ", "SETEX", keyname, ttl, value)
	res, err := rc.Do("SETEX", keyname, ttl, value)
	if err != nil {
		fmt.Println("SET error", err.Error())
		return err
	}
	fmt.Println(res)
	return nil
}

// PSETEX key milliseconds value
// 这个命令和 SETEX 命令相似，但它以毫秒为单位设置 key 的生存时间，而不是像 SETEX 命令那样，以秒为单位。
func StringPSETEX(rc redis.Conn, keyname string, ttl int64, value interface{}) error {
	fmt.Println("执行redis : ", "PSETEX", keyname, ttl, value)
	res, err := rc.Do("PSETEX", keyname, ttl, value)
	if err != nil {
		fmt.Println("SET error", err.Error())
		return err
	}
	fmt.Println(res)
	return nil
}

// SETNX key value
// 将 key 的值设为 value ，当且仅当 key 不存在。
// 若给定的 key 已经存在，则 SETNX 不做任何动作。
func StringSETNX(rc redis.Conn, keyname string, value interface{}) error {
	fmt.Println("执行redis : ", "SETNX", keyname, value)
	res, err := rc.Do("SETNX", keyname, value)
	if err != nil {
		fmt.Println("SET error", err.Error())
		return err
	}
	fmt.Println(res)
	return nil
}

// SETRANGE key offset value
// 用 value 参数覆写(overwrite)给定 key 所储存的字符串值，从偏移量 offset 开始。
// 不存在的 key 当作空白字符串处理。
func StringSETRANGE(rc redis.Conn, keyname string, offset int64, value interface{}) error {
	fmt.Println("执行redis : ", "SETRANGE", keyname, offset, value)
	res, err := rc.Do("SETRANGE", keyname, offset, value)
	if err != nil {
		fmt.Println("SET error", err.Error())
		return err
	}
	fmt.Println(res)
	return nil
}

//APPEND key value
//如果 key 已经存在并且是一个字符串， APPEND 命令将 value 追加到 key 原来的值的末尾。
//如果 key 不存在， APPEND 就简单地将给定 key 设为 value ，就像执行 SET key value 一样。
func StringAPPEND(rc redis.Conn, keyname string, value interface{}) error {
	fmt.Println("执行redis : ", "APPEND", keyname, value)
	res, err := redis.String(rc.Do("APPEND", keyname, value))
	if err != nil {
		fmt.Println("SET error", err.Error())
		return err
	}
	fmt.Println(res)
	return nil
}

//SETBIT key offset value
//对 key 所储存的字符串值，设置或清除指定偏移量上的位(bit)。
//value : 位的设置或清除取决于 value 参数，可以是 0 也可以是 1 。
//注意 offset 不能太大，越大key越大
func StringSETBIT() {}

//BITCOUNT key [start] [end]
//计算给定字符串中，被设置为 1 的比特位的数量。
func StringBITCOUNT() {}

//GETBIT key offset
//对 key 所储存的字符串值，获取指定偏移量上的位(bit)。
//当 offset 比字符串值的长度大，或者 key 不存在时，返回 0 。
func StringGETBIT() {}

//BITOP operation destkey key [key ...]
//对一个或多个保存二进制位的字符串 key 进行位元操作，并将结果保存到 destkey 上。
//BITOP AND destkey key [key ...] ，对一个或多个 key 求逻辑并，并将结果保存到 destkey 。
// BITOP OR destkey key [key ...] ，对一个或多个 key 求逻辑或，并将结果保存到 destkey 。
// BITOP XOR destkey key [key ...] ，对一个或多个 key 求逻辑异或，并将结果保存到 destkey 。
// BITOP NOT destkey key ，对给定 key 求逻辑非，并将结果保存到 destkey 。
func StringBITOP() {}

// DECR key
// 将 key 中储存的数字值减一。
// 如果 key 不存在，那么 key 的值会先被初始化为 0 ，然后再执行 DECR 操作。
// 如果值包含错误的类型，或字符串类型的值不能表示为数字，那么返回一个错误。
func StringDECR(rc redis.Conn, keyname string) (res int64, err error) {
	fmt.Println("执行redis : ", "DECR", keyname)
	res, err = redis.Int64(rc.Do("DECR", keyname))
	if err != nil {
		fmt.Println("SET error", err.Error())
		res = -1
	}
	return
}

//DECRBY key decrement
//将 key 所储存的值减去减量 decrement 。
func StringDECRBY(rc redis.Conn, keyname, decrement string) (res int64, err error) {
	fmt.Println("执行redis : ", "DECRBY", keyname, decrement)
	res, err = redis.Int64(rc.Do("DECRBY", keyname, decrement))
	if err != nil {
		fmt.Println("SET error", err.Error())
		res = -1
	}
	return
}

//GETRANGE key start end
//返回 key 中字符串值的子字符串，字符串的截取范围由 start 和 end 两个偏移量决定(包括 start 和 end 在内)。
func StringGETRANGE(rc redis.Conn, keyname string, start, end int64) (res string, err error) {
	fmt.Println("执行redis : ", "GETRANGE", keyname, start, end)
	res, err = redis.String(rc.Do("GETRANGE", keyname, start, end))
	if err != nil {
		fmt.Println("SET error", err.Error())
		res = ""
	}
	return
}

//GETSET key value
// 将给定 key 的值设为 value ，并返回 key 的旧值(old value)。
// 当 key 存在但不是字符串类型时，返回一个错误。
func StringGETSET(rc redis.Conn, keyname string, value interface{}) (res string, err error) {
	fmt.Println("执行redis : ", "GETSET", keyname, value)
	res, err = redis.String(rc.Do("GETSET", keyname, value))
	if err != nil {
		fmt.Println("SET error", err.Error())
		res = ""
	}
	return
}

// INCR key
// 将 key 中储存的数字值增一。
// 如果 key 不存在，那么 key 的值会先被初始化为 0 ，然后再执行 INCR 操作。
// 如果值包含错误的类型，或字符串类型的值不能表示为数字，那么返回一个错误。
func StringINCR(rc redis.Conn, keyname string) (res int64, err error) {
	fmt.Println("执行redis : ", "INCR", keyname)
	res, err = redis.Int64(rc.Do("INCR", keyname))
	if err != nil {
		fmt.Println("SET error", err.Error())
		res = -1
	}
	return
}

// INCRBY key increment
// 将 key 所储存的值加上增量 increment 。
func StringINCRBY(rc redis.Conn, keyname, increment string) (res int64, err error) {
	fmt.Println("执行redis : ", "INCRBY", keyname, increment)
	res, err = redis.Int64(rc.Do("INCRBY", keyname, increment))
	if err != nil {
		fmt.Println("SET error", err.Error())
		res = -1
	}
	return
}

//INCRBYFLOAT key increment
//为 key 中所储存的值加上浮点数增量 increment 。
func StringINCRBYFLOAT(rc redis.Conn, keyname, increment float64) (res float64, err error) {
	fmt.Println("执行redis : ", "INCRBYFLOAT", keyname, increment)
	res, err = redis.Float64(rc.Do("INCRBYFLOAT", keyname, increment))
	if err != nil {
		fmt.Println("SET error", err.Error())
		res = -1
	}
	return
}

// MGET key [key ...]
// 返回所有(一个或多个)给定 key 的值。
// 如果给定的 key 里面，有某个 key 不存在，那么这个 key 返回特殊值 nil 。因此，该命令永不失败。
func StringMGET(rc redis.Conn, keynames []interface{}) (res []string, err error) {
	args := redis.Args{}
	for _, value := range keynames {
		args = args.Add(value)
	}
	fmt.Println("执行redis : ", "MGET", args)
	res, err = redis.Strings(rc.Do("MGET", args))
	if err != nil {
		fmt.Println("SET error", err.Error())
	}
	return
}

// MSET key value [key value ...]
// 同时设置一个或多个 key-value 对。
// 如果某个给定 key 已经存在，那么 MSET 会用新值覆盖原来的旧值，如果这不是你所希望的效果，请考虑使用 MSETNX 命令：它只会在所有给定 key 都不存在的情况下进行设置操作。
// MSET 是一个原子性(atomic)操作，所有给定 key 都会在同一时间内被设置，某些给定 key 被更新而另一些给定 key 没有改变的情况，不可能发生。
func StringMSET(rc redis.Conn, datas []interface{}) (err error) {
	args := redis.Args{}
	for _, value := range datas {
		args = args.Add(value)
	}
	fmt.Println("执行redis : ", "MSET", args)
	res, err := rc.Do("MSET", args)
	if err != nil {
		fmt.Println("SET error", err.Error())
	}
	fmt.Println(res)
	return
}

// MSETNX key value [key value ...]
// 同时设置一个或多个 key-value 对，当且仅当所有给定 key 都不存在。
// 即使只有一个给定 key 已存在， MSETNX 也会拒绝执行所有给定 key 的设置操作。
// MSETNX 是原子性的，因此它可以用作设置多个不同 key 表示不同字段(field)的唯一性逻辑对象(unique logic object)，所有字段要么全被设置，要么全不被设置。
func StringMSETNX(rc redis.Conn, datas []interface{}) (err error) {
	args := redis.Args{}
	for _, value := range datas {
		args = args.Add(value)
	}
	fmt.Println("执行redis : ", "MSETNX", args)
	res, err := rc.Do("MSETNX", args)
	if err != nil {
		fmt.Println("SET error", err.Error())
	}
	fmt.Println(res)
	return
}

// STRLEN key
// 返回 key 所储存的字符串值的长度。
// 当 key 储存的不是字符串值时，返回一个错误。
func StringSTRLEN() {}
