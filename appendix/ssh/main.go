package main

import (
	"fmt"
	"os/exec"
	"unsafe"

	"io/ioutil"
	"time"

	"golang.org/x/crypto/ssh"
)

const (
	privateKeyPath = ""
)

func main() {
	out, err := exec.Command("ssh-keyscan", "github.com").Output()
	if err != nil {
		panic(err)
	}

	key, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		panic(err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		panic(err)
	}

	hostKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(*(*string)(unsafe.Pointer(&out))))
	if err != nil {
		panic(err)
	}

	config := &ssh.ClientConfig{
		User:            "git",
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		Timeout:         5 * time.Second,
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	conn, err := ssh.Dial("tcp", "github.com:22", config)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	if err := session.Shell(); err != nil {
		panic(err)
	}
	if err := session.Wait(); err != nil {
		sshError, ok := err.(*ssh.ExitError)
		if ok {
			fmt.Println("ssh error:", sshError.Error())
		} else {
			panic(err)
		}
	}
}
