package randexample

import (
	"fmt"
	"math/rand"
	"os"
	"text/tabwriter"
)

// Rand example
func Rand() {
	// 创造并设置生成器.
	// 通常应该使用非固定种子, 如time.Now().UnixNano().
	// 使用固定的种子挥在每次运行中产生相同的输出
	r := rand.New(rand.NewSource(99))

	//这里的tabwriter帮助我们生成对齐的输出。
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	defer w.Flush()
	show := func(name string, v1, v2, v3 interface{}) {
		fmt.Fprintf(w, "%s\t%v\t%v\t%v\n", name, v1, v2, v3)
	}

	// Float32 和Float64 的值在 [0, 1)之中.
	show("Float32", r.Float32(), r.Float32(), r.Float32())
	show("Float64", r.Float64(), r.Float64(), r.Float64())

	// ExpFloat64值的平均值为1，但是呈指数衰减。
	show("ExpFloat64", r.ExpFloat64(), r.ExpFloat64(), r.ExpFloat64())

	// NormFloat64值的平均值为0，标准差为1。
	show("NormFloat64", r.NormFloat64(), r.NormFloat64(), r.NormFloat64())

	// Int31，Int63和Uint32生成给定宽度的值。
	// Int方法（未显示）类似于Int31或Int63
	//取决于'int'的大小。
	show("Int31", r.Int31(), r.Int31(), r.Int31())
	show("Int63", r.Int63(), r.Int63(), r.Int63())
	show("Uint32", r.Uint32(), r.Uint32(), r.Uint32())

	// Intn，Int31n和Int63n将其输出限制为<n。
	//他们比使用r.Int（）％n更谨慎。
	show("Intn(10)", r.Intn(10), r.Intn(10), r.Intn(10))
	show("Int31n(10)", r.Int31n(10), r.Int31n(10), r.Int31n(10))
	show("Int63n(10)", r.Int63n(10), r.Int63n(10), r.Int63n(10))

	// Perm生成数字[0，n]的随机排列。
	show("Perm", r.Perm(5), r.Perm(5), r.Perm(5))
}
