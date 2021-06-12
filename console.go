package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/gorilla/mux"
)

type javaproc interface {
	Stdout() io.ReadCloser
	Stdin() io.WriteCloser
	Start() error
	Stop() error
}

type javaprocm struct {
	cmd *exec.Cmd
}

type Console struct {
	Cmd    *exec.Cmd
	stdout *bufio.Reader
	stdin  *bufio.Writer
	// Experemintal features may cause some issues
	tee       *io.Reader
	dupBuffer *bytes.Buffer
}

func (p *javaprocm) Stdout() (r io.ReadCloser) {
	r, _ = p.cmd.StdoutPipe()
	return
}

func (p *javaprocm) Stdin() (w io.WriteCloser) {
	w, _ = p.cmd.StdinPipe()
	return
}

func (p *javaprocm) Start() (err error) {
	err = p.cmd.Start()
	return
}

func (p *javaprocm) Stop() (err error) {
	err = p.cmd.Process.Kill()
	return
}

func execCommand(spath string, minHeap, maxHeap int) *javaprocm {
	minHeapFlag := fmt.Sprintf("-Xms%dM", minHeap)
	maxHeapFlag := fmt.Sprintf("-Xms%dM", maxHeap)
	command := exec.Command("java", minHeapFlag, maxHeapFlag, "-jar", spath, "nogui")
	// Should be like /createserver/{id} or maybe drop it
	command.Dir = "./serverfiles"
	return &javaprocm{cmd: command}
}

func InitConsole() *Console {
	cmd := execCommand("server.jar", 2048, 10240)
	console := &Console{Cmd: cmd.cmd}
	console.stdout = bufio.NewReader(cmd.Stdout())
	console.stdin = bufio.NewWriter(cmd.Stdin())
	return console
}

func (c *Console) ExecCommand(command string) (err error) {
	// Logging experimental
	file, err := os.OpenFile("serverlogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Println("[ " + time.Now().String() + "] " + "user /{USER} issued server command /{COMMAND}" + command + " on server /{server}")

	comm := fmt.Sprintf("%s\r\n", command)
	_, err = c.stdin.WriteString(comm)

	if err != nil {
		return
	}
	return c.stdin.Flush()
}

func (c *Console) ReadLine() (s string, err error) {
	s, err = c.stdout.ReadString('\n')
	ioReader := io.TeeReader(c.stdout, c.dupBuffer)
	c.tee = &ioReader
	return
}

// TODO : Depricate this
func (c *Console) ReadLineAndRespond() (s string, err error) {
	s, err = c.stdout.ReadString('\n')
	var x bytes.Buffer
	tee := io.TeeReader(c.stdout, &x)
	_ = tee
	return
}

func ServerCommand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	file, err := os.OpenFile("serverlogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Println("Recived a command from client")

	params := mux.Vars(r)
	command := params["command"]

	Xonsole.ExecCommand(command)

	/* for {
		if x, err := Xonsole.ReadLine(); err != io.EOF {
			_ = x //fmt.Println(x)
			fmt.Println(err)
		} else {
			fmt.Println("Shoud break now")
			break
		}
	} */

	resp := struct {
		Message string `json:"message"`
	}{Message: command}
	json.NewEncoder(w).Encode(resp)
	// w.Write([]byte("vusdohviosdhjfgiosdjflkdsjcmkl;dsaj"))
}
