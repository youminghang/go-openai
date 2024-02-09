// Copyright 2022 Innkeeper youminghang(游铭杭) <youminghang123@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/youminghang/go-openai.

package main

import (
	"os"

	_ "go.uber.org/automaxprocs"

	"github.com/youminghang/go-openai/internal/openai"
)

// go程序的默入口函数（主函数）
func main() {
	command := openai.NewOpenAiCommmand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}

}
