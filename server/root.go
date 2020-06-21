package server

import (
	"log"
	"net/http"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
	"github.com/markbates/pkger"
	"github.com/spf13/viper"
//  "github.com/arrase/multi-repo-workspace/cli/actions"
)

type Message struct {
	Topic   string                 `json:"topic"`
	Payload map[string]interface{} `json:"payload"`
}

type WSServer struct {
	Upgrader websocket.Upgrader
	Conn     *websocket.Conn
}

func (ws *WSServer) Run() {
	ws.Upgrader = websocket.Upgrader{
		ReadBufferSize:    4096,
		WriteBufferSize:   4096,
		EnableCompression: true,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

  viper.WatchConfig()
  viper.OnConfigChange(ws.ConfigCallback)

	http.Handle("/", http.FileServer(pkger.Dir("/public")))
	http.HandleFunc("/ws", ws.serve)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func (ws *WSServer) ConfigCallback(e fsnotify.Event) {
	if e.Op.String() == "WRITE" {
		ws.Conn.WriteJSON(Message{
			Topic:   "config-set",
			Payload: viper.AllSettings(),
		})
	}
}

func (ws *WSServer) serve(w http.ResponseWriter, r *http.Request) {
	ws.Conn, _ = ws.Upgrader.Upgrade(w, r, nil)
	defer ws.Conn.Close()

	for {
		var m Message
		var r = Message{}

		err := ws.Conn.ReadJSON(&m)
		if err != nil {
			log.Println("read:", err)
			break
		}

		log.Println(m.Topic)
		switch topic := m.Topic; topic {
		case "config-get":
			r.Topic = "config-set"
			r.Payload = viper.AllSettings()
		default:
			r.Topic = "error"
		}

		err = ws.Conn.WriteJSON(r)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
