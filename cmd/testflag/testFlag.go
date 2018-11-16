package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	h bool
	t, T bool
	q    *bool
	s string
	test string
)

/**
 * -flag // 只支持bool类型, 有则true，无则false， 如本例子， -h 出现 则为true
 * -flag=x // -h=true -s=test
 * -flag x // -h true -s test2
 * 注意： -h --h 都可以，但是 ---h 无法识别
 */
func main() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&t, "t", false, "test configuration and exit")
	flag.BoolVar(&T, "T", false, "test configuration, dump it and exit")
	q = flag.Bool("q", false, "suppress non-error messages during configuration testing")
	flag.StringVar(&s, "s", "", "send `signal` to a master process: stop, quit, reopen, reload")
	flag.StringVar(&test, "federation-namespace", "hehe", "test hehe")
	flag.Usage = usage

	// 解析参数
	flag.Parse()

	if h {
		flag.Usage()
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `nginx version: nginx/1.10.0
Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]

Options:
`)
	flag.PrintDefaults()
}
