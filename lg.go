package lg

import (
	"fmt"
	"log"
	"os"

	"github.com/krsanky/config"
)

var Log *log.Logger

func init() {
	log_file := config.Get("log_file")
	f, err := os.OpenFile(log_file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("error opening file:%s err:%s", log_file, err.Error()))
	}

	Log = log.New(f, ":", log.Lshortfile)
	Log.Println("opening logfile ...")
}
