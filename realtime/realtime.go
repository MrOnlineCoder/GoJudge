package realtime

import (
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/gorilla/mux"

	"gojudge/api/utils"
)

type Event struct {
	Type string `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type Client struct {
	Ws *websocket.Conn
	WriteChan chan Event
	CloseChan chan bool
}

type Hub struct {
	clients map[*Client]bool
	broadcast chan Event
}

func (h *Hub) Init() {
	h.broadcast = make(chan Event);
	h.clients = make(map[*Client]bool);
}

func (h *Hub) Register(cl *Client) {
	h.clients[cl] = true;
}

func (h *Hub) Unregister(cl *Client) {
	delete(h.clients, cl);
	cl.CloseChan <- true;
}

func (h *Hub) Broadcast(e Event) {
	for c, _ := range h.clients {
		c.WriteChan<- e;
	}
}

var hub Hub;

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
  WriteBufferSize: 1024,
}

func clientServe(client *Client) {
	for {
		select {
			case ev := <-client.WriteChan:
					err := client.Ws.WriteJSON(ev);
					if err != nil {
						hub.Unregister(client);
						break;
					}
			case sig := <-client.CloseChan:
					if sig {
						break;
					}
		}
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		utils.SendError(w, "Websocket Upgrade failed.");
		return;
	}

	client := &Client{
		Ws: ws,
		WriteChan: make(chan Event),
		CloseChan: make(chan bool),
	};

	hub.Register(client);

	go clientServe(client);
}

func EmitSubmissionUpdate(id int, verdict string, tests int) {
	hub.Broadcast(Event{
		Type: "submission_update",
		Data: map[string]interface{} {
			"submission_id": id,
			"verdict": verdict,
			"passed_tests": tests,
		},
	});
}

func InitRealtimeAPI(rt *mux.Router) {
	hub.Init();
	rt.HandleFunc("/ws", wsHandler);
}