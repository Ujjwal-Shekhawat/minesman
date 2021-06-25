package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
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

// Routes releated
type Ts struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var resBody Ts

	if err := decoder.Decode(&resBody); err != nil {
		log.Fatal(err.Error())
	}

	resp := struct {
		Username string `json:"username"`
		Token    string `json:"token"`
	}{
		Username: "kamisama",
		Token:    "lmao_success_boi",
	}
	if resBody.Username == "kamisama" && resBody.Password == "kamisama" {
		json.NewEncoder(w).Encode(resp)
		return
	} else {
		// no token for ya
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(struct{}{})
	}
}

func AuthConsole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Hitted Auth Endpoint")
	if r.Header.Get("auth-token") == "lmao_success_boi" {
		w.WriteHeader(200)
		resp := struct {
			Username string `json:"username"`
			Token    string `json:"token"`
			IsAuth   bool   `json:"isAuth"`
		}{
			Username: "kamisama",
			Token:    "lmao_success_boi",
			IsAuth:   true,
		}
		json.NewEncoder(w).Encode(resp)
		return
	} else {
		w.WriteHeader(405)
		json.NewEncoder(w).Encode(struct {
			IsAuth bool `json:"isAuth"`
		}{IsAuth: false})
	}
}
