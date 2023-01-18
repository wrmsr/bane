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
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/log"
)

func main() {
	addr := "irc.darwin.network:6697"

	cfg := tls.Config{}
	conn, err := tls.Dial("tcp", addr, &cfg)
	if err != nil {
		log.Fatal("TLS connection failed: " + err.Error())
	}
	defer log.OrError(conn.Close)

	//

	certChain := conn.ConnectionState().PeerCertificates
	for i, cert := range certChain {
		fmt.Println(i)
		fmt.Println("Issuer:", cert.Issuer)
		fmt.Println("Subject:", cert.Subject)
		fmt.Println("Version:", cert.Version)
		fmt.Println("NotAfter:", cert.NotAfter)
		fmt.Println("DNS names:", cert.DNSNames)
		fmt.Println("")
	}

	sendCmd := func(cmd string, message string) {
		command := []byte(fmt.Sprintf("%s %s\r\n", cmd, message))
		check.Must1(conn.Write(command))
	}

	username := "bane_test_hi_shiv"
	sendCmd("PASS", "smellyoulater")
	sendCmd("NICK", username)
	sendCmd("USER", fmt.Sprintf("%s * * :%s", username, username))

	getResponse := func() string {
		var b [512]byte
		check.Must1(conn.Read(b[:]))
		return string(b[:])
	}

	channel := "#test"

	joinChannel := func() {
		sendCmd("JOIN", channel)
	}

	joined := false
	for !joined {
		resp := getResponse()
		fmt.Println(resp)

		if strings.Contains(resp, "No Ident response") {
			sendCmd("NICK", username)
			sendCmd("USER", fmt.Sprintf("%s * * :%s", username, username))
		}

		// we"re accepted, now let"s join the channel!
		if strings.Contains(resp, "376") {
			joinChannel()
		}

		// username already in use? try to use username with _
		if strings.Contains(resp, "433") {
			username = "_" + username
			sendCmd("NICK", username)
			sendCmd("USER", fmt.Sprintf("%s * * :%s", username, username))
		}

		// if PING send PONG with name of the server
		if strings.Contains(resp, "PING") {
			sendCmd("PONG", ":"+strings.Split(resp, ":")[1])
		}

		// we"ve joined
		if strings.Contains(resp, "366") {
			joined = true
		}
	}

	fmt.Println("connected!")

	sendMessageToChannel := func(msg string) {
		cmd := fmt.Sprintf("PRIVMSG %s", channel)
		msg = ":" + msg
		sendCmd(cmd, msg)
	}

	printResponse := func() {
		resp := getResponse()
		if resp != "" {
			fmt.Printf("resp: %s\n", resp)
			// msg = utils.parse_message(resp)
			// print(f'msg: {msg}')
			// msg = resp.strip().split(':')
			// print('< {}> {}'.format(msg[1].split('!')[0], msg[2].strip()))
		}
	}

	//

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		text := check.Must1(reader.ReadString('\n'))
		text = strings.Replace(text, "\n", "", -1)

		if text == "/quit" {
			sendCmd("QUIT", "Good bye!")
			break
		}

		if "hi" == text {
			fmt.Println("hello, Yourself")
		}

		sendMessageToChannel(text)

		go printResponse()
	}
}

//

type Conn struct {
	c   net.Conn
	br  *bufio.Reader
	rfn func([]byte) error
}

func (c *Conn) run() error {
	var b [512]byte
	for {
		n, err := c.br.Read(b[:])
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}

		if err = c.rfn(b[:n]); err != nil {
			return err
		}
	}
}
