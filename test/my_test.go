package test

import (
	"gos7"
	"log"
	"os"
	"testing"
	"time"
)

func TestReadSomeValue(t *testing.T) {

	handler := gos7.NewTCPClientHandler("192.168.1.111:102", 0, 2)
	handler.Timeout = 200 * time.Second
	handler.IdleTimeout = 200 * time.Second
	handler.Logger = log.New(os.Stdout, "tcp: ", log.LstdFlags)
	// Connect manually so that multiple requests are handled in one connection session
	err := handler.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	defer handler.Close()
	//init client
	client := gos7.NewClient(handler)
	address := 5628
	start := 0
	size := 2
	//buffer := make([]byte, 255)
	//value := 100
	//AGWriteDB to address DB2710 with value 100, start from position 8 with size = 2 (for an integer)
	//var helper gos7.Helper
	//helper.SetValueAt(buffer, 0, value)
	//err := client.AGWriteDB(address, start, size, buffer)
	//log.Print(err)

	buf := make([]byte, 255)
	//AGReadDB to address DB2710, start from position 8 with size = 2
	err = client.AGReadDB(address, start, size, buf)
	log.Println(err)

	var s7 gos7.Helper
	var result int16
	s7.GetValueAt(buf, 0, &result)

	log.Printf("油位:%d\n", result)
}
