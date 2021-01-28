package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./pi_point_100000000")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	b2 := make([]byte, 2)
	_, _ = f.Read(b2)

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println(err)
		return
	}

	for i, b := range bytes {
		if b == 48 { // 48 '0'
			continue
		}

		ibs := []byte(strconv.Itoa(i + 1))
		if ibs[0] != b {
			continue
		}
		ibsLen := len(ibs)
		subBytes := bytes[i : ibsLen+i]

		for k, ib := range ibs {
			if ib != subBytes[k] {
				break
			}
			if k+1 == ibsLen { //都匹配
				fmt.Println(i + 1)
			}
		}

	}
}
