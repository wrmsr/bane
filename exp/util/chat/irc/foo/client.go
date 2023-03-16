/*
https://github.com/golang/go/issues/15735
https://stackoverflow.com/a/19992525

TODO:
  - integration chat iface
  - https://github.com/Rapptz/discord.py
  - https://github.com/slackapi/python-slackclient

https://tools.ietf.org/html/rfc2812
https://ircv3.net/irc/
https://modern.ircdocs.horse/

https://github.com/slingamn/irctest/
https://github.com/thobbs/pure-sasl

Host:     irc.darwin.network
Port:     6697
SSL/TLS:  true
Password: smellyoulater
Channel:  #darwin
*/
package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"sync"

	"github.com/wrmsr/bane/pkg/util/check"
	eu "github.com/wrmsr/bane/pkg/util/errors"
	"github.com/wrmsr/bane/pkg/util/log"
)

//

type IrcClientConfig struct {
	Addr string

	// FIXME: app-layer

	Pass string

	Nick string
	Chan string
}

func DefaultIrcClientConfig() IrcClientConfig {
	return IrcClientConfig{
		Addr: "irc.darwin.network:6697",

		// FIXME:

		Pass: "smellyoulater",

		Nick: "bane_test_hi_shiv",
		Chan: "#test",
	}
}

//

type Cmd struct {
	C, A string
}

func (c Cmd) Bytes() []byte {
	return []byte(fmt.Sprintf("%s %s\r\n", c.C, c.A))
}

//

type Handler = func(buf []byte) error

type IrcClient interface {
	Close() error

	AddHandler(h Handler)
	Connect() error
	Send(buf []byte) error
	Error() error
}

//

type IrcClientImpl struct {
	cfg IrcClientConfig

	conn net.Conn

	br *bufio.Reader

	hs []func(buf []byte) error

	mtx sync.Mutex

	closeCh chan struct{}

	readCh    chan []byte
	readErrCh chan error

	writeCh chan []byte

	running bool
	err     error
}

var _ IrcClient = &IrcClientImpl{}

func NewIrcClient(cfg IrcClientConfig) IrcClient {
	return &IrcClientImpl{
		cfg: cfg,

		closeCh: make(chan struct{}),

		readCh:    make(chan []byte),
		readErrCh: make(chan error),

		writeCh: make(chan []byte),
	}
}

func (c *IrcClientImpl) Close() (err error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.br = nil

	if c.conn != nil {
		err = eu.Append(err, c.conn.Close())
		c.conn = nil
	}

	return
}

func (c *IrcClientImpl) AddHandler(h Handler) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.hs = append(c.hs, h)
}

func (c *IrcClientImpl) Connect() (err error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	check.Nil(c.conn)

	tlsCfg := tls.Config{}
	conn, err := tls.Dial("tcp", c.cfg.Addr, &tlsCfg)
	if err != nil {
		return
	}

	c.conn = conn
	c.br = bufio.NewReader(conn)
	c.running = true

	go c.handleLoop()
	go c.readLoop()
	go c.writeLoop()

	//defer func() {
	//	if err != nil {
	//		err = eu.Append(err, c.Close())
	//	}
	//}()

	return
}

func (c *IrcClientImpl) Error() error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	return c.err
}

func (c *IrcClientImpl) setError(err error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

}

func (c *IrcClientImpl) handleLoop() {
	for {
		select {
		case <-c.closeCh:
			break
		case buf := <-c.readCh:
			for _, h := range c.hs {
				if err := h(buf); err != nil {

				}
			}
		case err := <-c.readErrCh:
			panic(err) // FIXME
		}
	}
}

func (c *IrcClientImpl) readLoop() {
	for {
		select {
		case <-c.closeCh:
			break
		}
	}
}

func (c *IrcClientImpl) writeLoop() {
	for {
		select {
		case <-c.closeCh:
			break
		}
	}
}

func (c *IrcClientImpl) Send(buf []byte) error {
	_, err := c.conn.Write(buf) // FIXME: lol
	return err
}

func (c *IrcClientImpl) read() (string, error) {
	b, err := c.br.ReadBytes('\n')
	if err != nil {
		return "", err
	}
	return string(b), nil
}

//

func SetNick(nick string) []Cmd {
	return []Cmd{
		{C: "NICK", A: nick},
		{C: "USER", A: fmt.Sprintf("%s * * :%s", nick, nick)},
	}
}

//

func main() {
	c := NewIrcClient(DefaultIrcClientConfig())

	check.Must(c.Connect())
	defer log.OrError(c.Close)

	//

	//check.Must(c.sendCmd(setNick(c.cfg.Nick)...))
	//
	//joined := false
	//for !joined {
	//	resp := check.Must1(c.read())
	//	fmt.Println(resp)
	//
	//	if strings.Contains(resp, "No Ident response") {
	//		check.Must(c.sendCmd(setNick(c.cfg.Nick)...))
	//	}
	//
	//	// we"re accepted, now let"s join the channel!
	//	if strings.Contains(resp, "376") {
	//		check.Must(c.sendCmd(Cmd{"JOIN", c.cfg.Chan}))
	//	}
	//
	//	// username already in use? try to use username with _
	//	if strings.Contains(resp, "433") {
	//		c.cfg.Nick = "_" + c.cfg.Nick
	//		check.Must(c.sendCmd(setNick(c.cfg.Nick)...))
	//	}
	//
	//	// if PING send PONG with name of the server
	//	if strings.Contains(resp, "PING") {
	//		check.Must(c.sendCmd(Cmd{"PONG", ":" + strings.Split(resp, ":")[1]}))
	//	}
	//
	//	// we"ve joined
	//	if strings.Contains(resp, "366") {
	//		joined = true
	//	}
	//}

	fmt.Println("connected!")

	//sendMessageToChannel := func(msg string) {
	//	cmd := fmt.Sprintf("PRIVMSG %s", channel)
	//	msg = ":" + msg
	//	sendCmd(cmd, msg)
	//}
	//
	//printResponse := func() {
	//	resp := getResponse()
	//	if resp != "" {
	//		fmt.Printf("resp: %s\n", resp)
	//		// msg = utils.parse_message(resp)
	//		// print(f'msg: {msg}')
	//		// msg = resp.strip().split(':')
	//		// print('< {}> {}'.format(msg[1].split('!')[0], msg[2].strip()))
	//	}
	//}
	//
	////
	//
	//reader := bufio.NewReader(os.Stdin)
	//
	//for {
	//	fmt.Print("-> ")
	//	text := check.Must1(reader.ReadString('\n'))
	//	text = strings.Replace(text, "\n", "", -1)
	//
	//	if text == "/quit" {
	//		sendCmd("QUIT", "Good bye!")
	//		break
	//	}
	//
	//	if "hi" == text {
	//		fmt.Println("hello, Yourself")
	//	}
	//
	//	sendMessageToChannel(text)
	//
	//	go printResponse()
	//}
}

//

//type Conn struct {
//	c   net.Conn
//	br  *bufio.Reader
//	rfn func([]byte) error
//}
//
//func (c *Conn) run() error {
//	var b [512]byte
//	for {
//		n, err := c.br.Read(b[:])
//		if err != nil {
//			if errors.Is(err, io.EOF) {
//				return nil
//			}
//			return err
//		}
//
//		if err = c.rfn(b[:n]); err != nil {
//			return err
//		}
//	}
//}
