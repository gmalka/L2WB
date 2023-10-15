package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"sync"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	cur, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	host, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}
	UserNameAndHost := strings.Builder{}
	UserNameAndHost.WriteString(cur.Name)
	UserNameAndHost.WriteString("@")
	UserNameAndHost.WriteString(host)
	UserNameAndHost.WriteString("  ~ $ ")

	for {
		fmt.Print(UserNameAndHost.String())
		command, err := in.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		err = Execution(command)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
	}
}

func Execution(src string) (err error) {
	src = strings.Trim(src, "\n")
	commands := strings.Split(src, "|")
	cmds := make([]*exec.Cmd, len(commands))

	for k, command := range commands {
		command = strings.Trim(command, " ")
		args := strings.Split(command, " ")
		if args[0] == "cd" {
			if len(args) < 2 {
				err = os.Chdir("/Users")
				if err != nil {
					return err
				}
				continue
			}
			os.Chdir(args[1])
			return
		}
		if args[0] == "\\quit" {
			os.Exit(0)
		}
		cmd := exec.Command(args[0], args[1:]...)
		cmds[k] = cmd
	}
	var (
		i  int
		wg sync.WaitGroup
	)
	for i = 0; i < len(cmds)-1; i++ {
		r, w := io.Pipe()
		cmds[i].Stdout = w
		cmds[i].Stderr = os.Stderr
		cmds[i+1].Stdin = r
		wg.Add(1)
		go func(i int, w *io.PipeWriter) {
			defer w.Close()
			err = cmds[i].Run()
			wg.Done()
			if err != nil {
				fmt.Println(err)
			}
		}(i, w)
	}
	cmds[i].Stdout = os.Stdout
	cmds[i].Stderr = os.Stderr
	wg.Add(1)
	go func(i int) {
		cmds[i].Run()
		wg.Done()
		if err != nil {
			fmt.Println(err)
		}
	}(i)

	wg.Wait()
	return err
}