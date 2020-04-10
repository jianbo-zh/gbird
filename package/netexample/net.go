package netexample

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"time"
)

// TCPServer example
func TCPServer() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()
	fmt.Println("Server ready to read ...")
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("A client connected :" + tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}
}

func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println(" Disconnected : " + ipStr)
		conn.Close()
	}()

	reader := bufio.NewReader(conn)
	i := 0

	for {
		message, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		fmt.Println(string(message))
		time.Sleep(time.Second * 3)
		msg := time.Now().String() + conn.RemoteAddr().String() + " Server Say hello! \n"
		b := []byte(msg)
		conn.Write(b)
		i++
		if i > 10 {
			break
		}
	}
}

// TCPClient example
func TCPClient() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("Client connect error ! " + err.Error())
		return
	}
	defer conn.Close()
	fmt.Println(conn.LocalAddr().String() + " : Client connected!")
	onMessageReceived(conn)
}

func onMessageReceived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	b := []byte(conn.LocalAddr().String() + " Say hello to Server... \n")
	conn.Write(b)
	for {
		msg, err := reader.ReadString('\n')
		fmt.Println("ReadString")
		fmt.Println(msg)

		if err != nil || err == io.EOF {
			fmt.Println(err)
			break
		}
		time.Sleep(time.Second * 2)

		fmt.Println("writing...")

		b := []byte(conn.LocalAddr().String() + " write data to Server... \n")
		_, err = conn.Write(b)

		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

// GeneralServer example
func GeneralServer() {
	const recvLen = 10
	listener, err := net.Listen("tcp", "127.0.0.1:10005")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer conn.Close()

		for {
			data := make([]byte, recvLen)
			_, err = conn.Read(data)
			if err != nil {
				fmt.Println(err)
				break
			}
			strData := string(data)
			fmt.Println("Received:", strData)

			upper := strings.ToUpper(strData)
			_, err = conn.Write([]byte(upper))
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("Send:", upper)
		}
	}
}

// GeneralClient example
func GeneralClient() {
	const recvLen = 10
	conn, err := net.Dial("tcp", "127.0.0.1:10005")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		lineLen := len(line)
		n := 0
		for written := 0; written < lineLen; written += n {
			var toWrite string
			if lineLen-written > recvLen {
				toWrite = line[written : written+recvLen]
			} else {
				toWrite = line[written:]
			}
			n, err = conn.Write([]byte(toWrite))
			if err != nil {
				panic(err)
			}
			fmt.Println("Write:", toWrite)

			msg := make([]byte, recvLen)
			n, err = conn.Read(msg)
			if err != nil {
				panic(err)
			}
			fmt.Println("Response:", string(msg))
		}
	}
}

// UDPServer example
func UDPServer() {
	const recvLen = 10

	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10006")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer conn.Close()

	for {
		// Here must use make and give the lenth of buffer
		data := make([]byte, recvLen)
		_, rAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println(err)
			continue
		}

		strData := string(data)
		fmt.Println("Received:", strData)

		upper := strings.ToUpper(strData)
		_, err = conn.WriteToUDP([]byte(upper), rAddr)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Send:", upper)
	}
}

// UDPClient example
func UDPClient() {
	conn, err := net.Dial("udp", "127.0.0.1:10006")
	if err != nil {
		panic(err)
	}

	const recvLen = 10

	defer conn.Close()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()

		lineLen := len(line)

		n := 0
		for written := 0; written < lineLen; written += n {
			var toWrite string
			if lineLen-written > recvLen {
				toWrite = line[written : written+recvLen]
			} else {
				toWrite = line[written:]
			}

			n, err = conn.Write([]byte(toWrite))
			if err != nil {
				panic(err)
			}

			fmt.Println("Write:", toWrite)

			msg := make([]byte, recvLen)
			n, err = conn.Read(msg)
			if err != nil {
				panic(err)
			}

			fmt.Println("Response:", string(msg))
		}
	}
}

// UnixServer example
func UnixServer() {
	var unixAddr *net.UnixAddr
	unixAddr, _ = net.ResolveUnixAddr("unix", "/tmp/unix_sock")
	unixListener, _ := net.ListenUnix("unix", unixAddr)
	defer unixListener.Close()
	for {
		unixConn, err := unixListener.AcceptUnix()
		if err != nil {
			continue
		}

		fmt.Println("A client connected : " + unixConn.RemoteAddr().String())
		go unixPipe(unixConn)
	}
}

func unixPipe(conn *net.UnixConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected :" + ipStr)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		fmt.Println(string(message))
		msg := time.Now().String() + "\n"
		b := []byte(msg)
		conn.Write(b)
	}
}

// UnixClient example
func UnixClient() {
	var unixAddr *net.UnixAddr
	unixAddr, _ = net.ResolveUnixAddr("unix", "/tmp/unix_sock")

	conn, _ := net.DialUnix("unix", nil, unixAddr)
	defer conn.Close()
	fmt.Println("connected!")

	quitSignal := make(chan bool)

	go onMessageReceived2(conn, quitSignal)

	b := []byte("time\n")
	conn.Write(b)

	<-quitSignal
}

func onMessageReceived2(conn *net.UnixConn, quitSignal chan<- bool) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		fmt.Println(msg)
		if err != nil {
			quitSignal <- true
			break
		}
		time.Sleep(time.Second)
		b := []byte(msg)
		conn.Write(b)
	}
}

func testUnixAddr() string {
	f, err := ioutil.TempFile("", "go-nettest")
	if err != nil {
		panic(err)
	}
	addr := f.Name()
	f.Close()
	os.Remove(addr)
	return addr
}

// UnixgramServer example
func UnixgramServer() {
	var unixAddr *net.UnixAddr
	unixAddr, _ = net.ResolveUnixAddr("unixgram", "/tmp/unix_sock")
	conn, err := net.ListenUnixgram("unixgram", unixAddr)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		fmt.Println(string(message))
		msg := time.Now().String() + "\n"
		b := []byte(msg)
		// 此处貌似有问题
		conn.Write(b)
	}
}

// UnixgramClient example
func UnixgramClient() {
	var unixAddr *net.UnixAddr
	unixAddr, _ = net.ResolveUnixAddr("unixgram", "/tmp/unix_sock")

	conn, _ := net.DialUnix("unixgram", nil, unixAddr)
	defer conn.Close()
	fmt.Println("connected!")

	quitSignal := make(chan bool)

	go onMessageReceived3(conn, quitSignal)

	b := []byte("time\n")
	conn.Write(b)

	<-quitSignal
}

func onMessageReceived3(conn *net.UnixConn, quitSignal chan<- bool) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		fmt.Println(msg)
		if err != nil {
			quitSignal <- true
			break
		}
		time.Sleep(time.Second)
		b := []byte(msg)
		conn.Write(b)
	}
}
