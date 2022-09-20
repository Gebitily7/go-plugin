package base

import (
	"log"
	"os"
	"syscall"
)

func getWd()(path string) {
	// 获取当前所在目录
	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}

	log.Println(path)

	return
}

func Handle()  {
	// 打开进程可以看到的 / 目录
	RealRoot, err := os.Open("/")

	defer func(RealRoot *os.File) {
		err := RealRoot.Close()
		if err != nil {
			log.Println(err)
		}
	}(RealRoot)

	if err != nil {
		log.Fatalf("[ Error ] - /: %v\n", err)
	}

	path := getWd()

	// chroot 到当前目录
	err = syscall.Chroot(path)
	if err != nil {
		log.Fatalf("[ Error ] - chroot: %v\n", err)
	}

	getWd()

	// 切换到 / 目录
	err = RealRoot.Chdir()
	if err != nil {
		log.Fatalf("[ Error ] - chdir(): %v", err)
	}
	getWd()

	// chroot 到当前目录
	err = syscall.Chroot(".")
	if err != nil {
		log.Fatalf("[ Error ] - chroot back: %v", err)
	}
	getWd()
}
