package main
//testing changes

import (
//   "fmt"
   "github.com/gorilla/websocket"
   "log"
   "flag"
   "net/http"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

//var upgrader = websocket.Upgrader{} // use default options
var	upgrader  = websocket.Upgrader{
        CheckOrigin: myOrigin, 
	}

func myOrigin(r *http.Request) bool {
    //accept anything for now
    return true
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}


func main() {
	flag.Parse()
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}


