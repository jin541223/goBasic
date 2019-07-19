package main

// 引入包
import (
	"fmt"
	"os"
	"simplemath"
	"strconv"
)

// 定义用于点程序使用指南函数
var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe commands are:\n\tadd\t计算两个数值相加\n\tsqrt\t计算一个非负数平方根")
}

func main() {
	// 获取命令行参数，程序名称本身为第一个参数；如：calc add 1 2 第一个参数是calc
	args := os.Args
	// 除程序本身外，至少需要传入两个其他参数，否则就退出
	if args == nil || len(args) < 3 {
		Usage()
		return
	}

	switch args[1] {
	// 加法
	case "add":
		// 至少包含四个参数
		if len(args) != 4 {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}
		// 待相加数值，并将类型转化为整型
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		// 参数出错，退出
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}
		// 从simplemath包引入 Add 方法进行加法运算
		ret := simplemath.Add(v1, v2)
		// 打印计算结果
		fmt.Println("Result: ", ret)
	// 平方根
	case "sqrt":
		// 至少需要三个参数
		if len(args) != 3 {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		// 待计算平方根的值，将类型转化为整型
		v, err := strconv.Atoi(args[2])
		// 参数错误
		if err != nil {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		// 从 simplemath 包引入Sqrt方法做平方根运算
		ret := simplemath.Sqrt(v)

		fmt.Println("Result: ", ret)
		// 未定义计算方法时，调用程序指南
	default:
		Usage()
	}
}
