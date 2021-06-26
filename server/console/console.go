package console

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

var Xonsole *map[string](*Console) = &map[string](*Console){}

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
	cmd := execCommand("server.jar", 2048, 6000)
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
	// something ill like to implement in my free time
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
	return
}
