package connection

import (
	"bufio"
	servernotifs "bushuray-tui/lib/ServerNotifs"
	sharedtypes "bushuray-tui/shared_types"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type ConnectionHandler struct {
	host    string
	port    int
	conn    net.Conn
	writeMu sync.Mutex
}

func (ch *ConnectionHandler) Init(host string, port int) {
	ch.host = host
	ch.port = port
}

func (ch *ConnectionHandler) GetConnection() error {
	address := net.JoinHostPort(ch.host, strconv.Itoa(ch.port))

	conn, err := net.DialTimeout("tcp", address, 10*time.Second)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %w", address, err)
	}

	ch.conn = conn
	fmt.Printf("Connected to %s\n", address)

	return nil
}

func (ch *ConnectionHandler) HandleConnection(p *tea.Program) error {
	defer func() {
		p.Quit()
	}()

	if ch.conn == nil {
		return fmt.Errorf("no active connection - call GetConnection() first")
	}

	for {
		lengthBuf := make([]byte, 4)
		reader := bufio.NewReader(ch.conn)

		_, err := io.ReadFull(reader, lengthBuf)

		if err != nil {
			if err != io.EOF {
				log.Printf("Failed to read length , %v", err)
			}
			return err
		}

		length := binary.BigEndian.Uint32(lengthBuf)
		if length == 0 || length > 100*1024*1024 {
			log.Printf("Invalid length %d", length)
			return err
		}

		payload := make([]byte, length)

		_, err = io.ReadFull(reader, payload)

		if err != nil {
			log.Printf("Failed to read the payload %v", err)
			return err
		}

		log.Println("\n-----\n", string(payload))

		var raw_tcp_message sharedtypes.TcpMessage

		if err := json.Unmarshal(payload, &raw_tcp_message); err != nil {
			log.Printf("Invalid JSON: %v", err)
			return err
		}

		servernotifs.HandleNotification(raw_tcp_message)
	}
}

func (ch *ConnectionHandler) Send(msg []byte) error {
	ch.writeMu.Lock()
	defer ch.writeMu.Unlock()

	if ch.conn == nil {
		return fmt.Errorf("no active connection")
	}

	length := make([]byte, 4)
	binary.BigEndian.PutUint32(length, uint32(len(msg)))

	_, err := ch.conn.Write(length)
	if err != nil {
		log.Fatalf("Error sending length %d %v\n", length, err)
	}
	_, err = ch.conn.Write(msg)
	if err != nil {
		log.Fatalf("Error sending %s %v\n", msg, err)
	}

	return nil
}

func (ch *ConnectionHandler) Close() error {
	if ch.conn != nil {
		err := ch.conn.Close()
		ch.conn = nil
		return err
	}
	return nil
}
