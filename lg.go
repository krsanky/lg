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
	fmt.Printf("log_file:%s\n", log_file)
	f, err := os.OpenFile(log_file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("error opening file:%s"))
	}
	//defer f.Close()

	//Log = log.New(f, "logger: ", log.Llongfile)
	Log = log.New(f, ":", log.Lshortfile)
	//log.SetOutput(f)

	Log.Println("opening logfile ...")
}
