package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func xray() {

	fi, err := os.Open("./test.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		xray_scan(string(a))
		fmt.Println(string(a))
		if c == io.EOF {
			break
		}
	}
}

func xray_scan(url string) {
	outHtml := fmt.Sprintf("xray_scan_%d.html", time.Now().UnixNano())
	command := exec.Command("./xray_windows_amd64.exe", "webscan", "--poc", "/pocs/*", "--basic-crawler", url, "--html-output", outHtml)
	outinfo := bytes.Buffer{}
	command.Stdout = &outinfo
	err := command.Start()
	if err != nil {
		fmt.Println(err.Error())
	}
	if err = command.Wait(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(command.ProcessState.Pid())
		fmt.Println(command.ProcessState.Sys().(syscall.WaitStatus).ExitCode)
		fmt.Println(outinfo.String())
	}
}
