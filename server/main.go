package main

import (
	"flag"
	"fmt"
	"github.com/DreamerLWJ/all_in_one_clean_arch/app/blog"
	"github.com/DreamerLWJ/all_in_one_clean_arch/app/shop"
	"github.com/DreamerLWJ/all_in_one_clean_arch/routers"
)

func main() {
	var cfgPath string
	flag.StringVar(&cfgPath, "config", "", "配置文件")

	routers.Include(shop.Routers, blog.Routers)
	r := routers.Init()

	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
