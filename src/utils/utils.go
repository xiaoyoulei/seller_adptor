package utils

import (
	"context"
	"crypto/sha1"
	"fmt"
	//	"github.com/op/go-logging"
	"io"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func Init() {

}

func GenSearchid(imei string) (searchid string) {

	var tmp string
	tmp = imei
	tmp += time.Now().String()
	tmp += strconv.Itoa(rand.Int())
	log.Println(tmp)
	sha1_t := sha1.New()
	io.WriteString(sha1_t, tmp)
	searchid = fmt.Sprintf("%x", sha1_t.Sum(nil))
	return
}
func GenClickUrl(clk string, adinfo *context.AdInfo) (curl string) {
	return
}
