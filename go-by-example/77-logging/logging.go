package main

import ( 
	"bytes"
	"fmt"
	"log"
	"os"
	"log/slog"
)

func main () {
	log.Println("standard logger")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	mylog := log.New(os.Stdout, "my:", log.Lshortfile)
	mylog.Println("from mylog")

	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)
	buflog.Println("hello")
	fmt.Print("From buflog:", buf.String())

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	myslog.Info("hellow again", "key", "val", "age", 25)
}
