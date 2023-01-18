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
package irc

import (
	"bufio"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/dev/debug"
	"github.com/wrmsr/bane/pkg/util/log"
	"github.com/wrmsr/bane/pkg/util/sync/pools"
)

/*
class IrcSimpleClient:

    def __init__(self, username, channel, server='irc.darwin.network', port=6697):
        super().__init__()

        self.username = username
        self.server = server
        self.port = port
        self.channel = channel

    def connect(self):
        ssl_context = ssl.create_default_context()
        rawconn = socket.create_connection((self.server, self.port)).__enter__()
        self.conn = ssl_context.wrap_socket(rawconn, server_hostname=self.server).__enter__()

    def get_response(self):
        return self.conn.recv(512).decode('utf-8')

    def send_cmd(self, cmd, message):
        command = '{} {}\r\n'.format(cmd, message).encode('utf-8')
        self.conn.send(command)

    def send_message_to_channel(self, message):
        command = 'PRIVMSG {}'.format(self.channel)
        message = ':' + message
        self.send_cmd(command, message)

    def join_channel(self):
        cmd = 'JOIN'
        channel = self.channel
        self.send_cmd(cmd, channel)


@pytest.mark.skip()
def test_irc():
    username = 'omnibus_irc_test'
    channel = '#test'

    def print_response():
        resp = client.get_response()
        if resp:
            print(f'resp: {resp}')
            msg = utils.parse_message(resp)
            print(f'msg: {msg}')
            # msg = resp.strip().split(':')
            # print('< {}> {}'.format(msg[1].split('!')[0], msg[2].strip()))

    cmd = ''
    joined = False
    client = IrcSimpleClient(username, channel)
    client.connect()

    client.send_cmd('PASS', 'smellyoulater')

    client.send_cmd('NICK', username)
    client.send_cmd('USER', '{} * * :{}'.format(username, username))

    while not joined:
        resp = client.get_response()
        print(resp.strip())

        if 'No Ident response' in resp:
            client.send_cmd('NICK', username)
            client.send_cmd('USER', '{} * * :{}'.format(username, username))

        # we're accepted, now let's join the channel!
        if '376' in resp:
            client.join_channel()

        # username already in use? try to use username with _
        if '433' in resp:
            username = '_' + username
            client.send_cmd('NICK', username)
            client.send_cmd('USER', '{} * * :{}'.format(username, username))

        # if PING send PONG with name of the server
        if 'PING' in resp:
            client.send_cmd('PONG', ':' + resp.split(':')[1])

        # we've joined
        if '366' in resp:
            joined = True

    while cmd != '/quit':
        cmd = input('< {}> '.format(username)).strip()
        if cmd == '/quit':
            client.send_cmd('QUIT', 'Good bye!')
        client.send_message_to_channel(cmd)

        # socket conn.receive blocks the program until a response is received
        # to prevent blocking program execution, receive should be threaded
        response_thread = threading.Thread(target=print_response)
        response_thread.daemon = True
        response_thread.start()
*/

func TestClient(t *testing.T) {
	if !debug.IsDebuggerAttached(context.Background()) {
		t.Skip("debugger not attached")
	}

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
	sendCmd("QUIT", "Good bye!")

	//

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		text := check.Must1(reader.ReadString('\n'))
		text = strings.Replace(text, "\n", "", -1)

		if text == "/quit" {
			break
		}

		if "hi" == text {
			fmt.Println("hello, Yourself")
		}

		/*
		   cmd = input('< {}> '.format(username)).strip()
		   if cmd == '/quit':
		       client.send_cmd('QUIT', 'Good bye!')
		   client.send_message_to_channel(cmd)

		   # socket conn.receive blocks the program until a response is received
		   # to prevent blocking program execution, receive should be threaded
		   response_thread = threading.Thread(target=print_response)
		   response_thread.daemon = True
		   response_thread.start()
		*/
	}
}

//

type PooledBufioWriter struct {
	u io.Writer
	p pools.Pool[*bufio.Writer]
	b *bufio.Writer
}

var _ io.Writer = &PooledBufioWriter{}

func (pw *PooledBufioWriter) Pooled() *bufio.Writer {
	if pw.p == nil {
		return nil
	}
	return pw.b
}

func (pw *PooledBufioWriter) Get() *bufio.Writer {
	if pw.b == nil {
		if pw.p != nil {
			pw.b = pw.p.Get()
			pw.b.Reset(pw.u)
		} else {
			pw.b = bufio.NewWriter(pw.u)
		}
	}
	return pw.b
}

func (pw *PooledBufioWriter) Write(p []byte) (n int, err error) {
	panic("implement me")
}

func (pw *PooledBufioWriter) Flush() error {
	if pw.b == nil {
		return nil
	}

	if pw.p != nil {
		err := pw.b.Flush()
		pw.b.Reset(io.Discard)
		pw.p.Put(pw.b)
		pw.b = nil
		return err

	} else {
		if pw.b.Buffered() < 1 {
			return nil
		}
		return pw.b.Flush()
	}
}
