package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/websocket"

	"github.com/jamescooper/neuralRockets/model"
	"github.com/jamescooper/neuralRockets/rocket"
)

var root = flag.String("root", ".", "file system path")

func main() {
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/ws", websocket.Handler(handler))

	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Printf("ListenAndServe: %v", err)
	}
}

func handler(ws *websocket.Conn) {
	fmt.Println("Connection")
	fmt.Println(ws.RemoteAddr().String())
	rockets := rocket.InitRockets()
	flightTime := 0
	for ws.IsServerConn() {
		if websocket.Message.Send(ws, string(positions(rockets))) != nil {
			break
		}
		time.Sleep(10 * time.Millisecond)
		moveRockets(rockets)
		flightTime++
		if flightTime >= rocket.MaxFlightTime {
			flightTime = 0
			rocket.Breed(rockets)
		}
	}
	fmt.Println("dicsonnect")
}

func moveRockets(rs []*rocket.Rocket) {
	var err error
	for _, r := range rs {
		err = r.Update()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}
}

func positions(rs []*rocket.Rocket) []byte {
	numOfRockets := len(rs)
	rocketPositions := make([]model.Pos, numOfRockets)
	for index, r := range rs {
		rocketPositions[index] = r.BuildMapping()
	}

	posJSON, err := json.Marshal(rocketPositions)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return posJSON
}
