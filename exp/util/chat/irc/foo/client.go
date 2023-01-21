/*
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
	"strings"

	"github.com/wrmsr/bane/pkg/util/check"
	eu "github.com/wrmsr/bane/pkg/util/errors"
	"github.com/wrmsr/bane/pkg/util/log"
)

//

type IrcClientConfig struct {
	Addr string
	Pass string

	Nick string
	Chan string
}

func DefaultIrcClientConfig() IrcClientConfig {
	return IrcClientConfig{
		Addr: "irc.darwin.network:6697",
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
	Send(cmds ...Cmd) error
}

//

type IrcClientImpl struct {
	cfg IrcClientConfig

	conn net.Conn

	br *bufio.Reader

	hs []func(buf []byte) error
}

var _ IrcClient = &IrcClientImpl{}

func NewIrcClient(cfg IrcClientConfig) IrcClient {
	return &IrcClientImpl{
		cfg: cfg,
	}
}

func (c *IrcClientImpl) Close() (err error) {
	c.br = nil

	if c.conn != nil {
		err = eu.Append(err, c.conn.Close())
		c.conn = nil
	}

	return
}

func (c *IrcClientImpl) AddHandler(h Handler) {
	c.hs = append(c.hs, h)
}

func (c *IrcClientImpl) Connect() (err error) {
	check.Nil(c.conn)

	tlscfg := tls.Config{}
	conn, err := tls.Dial("tcp", c.cfg.Addr, &tlscfg)
	if err != nil {
		return
	}

	c.conn = conn
	c.br = bufio.NewReader(conn)

	defer func() {
		if err != nil {
			err = eu.Append(err, c.Close())
		}
	}()

	if err = c.Send(Cmd{"pass", c.cfg.Pass}); err != nil {
		return
	}

	return
}

func (c *IrcClientImpl) Send(cmds ...Cmd) error {
	for _, cmd := range cmds {
		if _, err := c.conn.Write(cmd.Bytes()); err != nil {
			return err
		}
	}
	return nil
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

	check.Must(c.sendCmd(setNick(c.cfg.Nick)...))

	joined := false
	for !joined {
		resp := check.Must1(c.read())
		fmt.Println(resp)

		if strings.Contains(resp, "No Ident response") {
			check.Must(c.sendCmd(setNick(c.cfg.Nick)...))
		}

		// we"re accepted, now let"s join the channel!
		if strings.Contains(resp, "376") {
			check.Must(c.sendCmd(Cmd{"JOIN", c.cfg.Chan}))
		}

		// username already in use? try to use username with _
		if strings.Contains(resp, "433") {
			c.cfg.Nick = "_" + c.cfg.Nick
			check.Must(c.sendCmd(setNick(c.cfg.Nick)...))
		}

		// if PING send PONG with name of the server
		if strings.Contains(resp, "PING") {
			check.Must(c.sendCmd(Cmd{"PONG", ":" + strings.Split(resp, ":")[1]}))
		}

		// we"ve joined
		if strings.Contains(resp, "366") {
			joined = true
		}
	}

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
