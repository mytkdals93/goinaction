package main

import (
	"log"
	"os"

	_ "github.com/mytkdals93/goinaction/chap2/sample/matchers"
	"github.com/mytkdals93/goinaction/chap2/sample/search"
)

func init() {
	// 표준 출력으로 로그를 출력하도록 변경한다.
	// 기본적으로 로그출 출력기는 표준 오류(stderr, stdard error)로 출력하도록 설정되어 있음
	log.SetOutput(os.Stdout)
}
func main() {
	search.Run("Knocking on our door")
}
