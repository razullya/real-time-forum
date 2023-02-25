package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type socketReader struct {
	con  *websocket.Conn
	mode int
	name string
}

var savedSocketReader []*socketReader

func Socket(w http.ResponseWriter, r *http.Request) {
	log.Println("socket request")
	if savedSocketReader == nil {
		savedSocketReader = make([]*socketReader, 0)
	}
	defer func() {

		err := recover()
		if err != nil {
			log.Println(err)
		}
		r.Body.Close()
		
	}()

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("socket: err:", err)
		return
	}
	ptrSocketReader := &socketReader{
		con: conn,
	}
	savedSocketReader = append(savedSocketReader, ptrSocketReader)

	ptrSocketReader.startThread()

}
func (s *socketReader) startThread() {
	s.writeMsg("razullya", "hello my dear friend!")
	s.mode = 1
	go func() {

		defer func() {
			err := recover()
			if err != nil {
				log.Println()
			}
			log.Println("thread socketreader finish")
		}()
		for {
			s.read()
		}
	}()
}
func (s *socketReader) writeMsg(name, msg string) {
	s.con.WriteMessage(websocket.TextMessage, []byte("<b>"+name+": </b>"+msg))
}
func (s *socketReader) read() {
	_, b, err := s.con.ReadMessage()
	if err != nil {
		panic(err)
	}
	log.Println(s.name + " " + string(b))
	log.Println(s.mode)
	if s.mode == 1 {
		s.name = string(b)
		s.writeMsg("razullya", "Welcome"+s.name+", please write a message and we will broadcast it to other users!")
		s.mode = 2
		return
	}
	s.broadcast(string(b))
	log.Println(s.name + " " + string(b))
}
func (s *socketReader) broadcast(str string) {
	for _, g := range savedSocketReader {
		if g == s {
			continue
		}
		if g.mode == 1 {
			continue
		}
		g.writeMsg(s.name, str)

	}

}
