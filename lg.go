package lg

import (
	"log"
	"os"

	"github.com/krsanky/config"
)

var Log *log.Logger

func init() {
	f, err := os.OpenFile(config.Get("log_file"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("error opening file: logfile.txt")
	}
	//defer f.Close()

	//Log = log.New(f, "logger: ", log.Llongfile)
	Log = log.New(f, ":", log.Lshortfile)
	//log.SetOutput(f)

	Log.Println("opening logfile ...")
}
