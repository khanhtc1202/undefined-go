package sshsession

import (
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"os"
	"time"
)

type runningSession struct {
	remoteSession *ssh.Session
}

var instance *runningSession

func Session() *runningSession {
	if instance == nil {
		instance = createSession()
		instance.remoteSession.Stdin = os.Stdin
		instance.remoteSession.Stdout = os.Stdout
		instance.remoteSession.Stderr = os.Stderr
	}
	return instance
}

func getClient() *ssh.ClientConfig {
	return nil
}

func createConnection() *ssh.Client {
	clientConfig := &ssh.ClientConfig{
		User: "root", // User
		Auth: []ssh.AuthMethod{
			ssh.Password("uzabase"), // Pass or Key pem file
		},
		Timeout:         100 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	endpoint := fmt.Sprintf("%s:%d", "speranza-dev.uzabase.lan", 22)
	conn, err := ssh.Dial("tcp", endpoint, clientConfig)
	if err != nil {
		panic(fmt.Sprintf("Error on create connection to %v : %v", endpoint, err))
	}
	return conn
}

func createSession() *runningSession {
	session, err := createConnection().NewSession()
	if err != nil {
		panic(fmt.Sprintf("Error on create new session : %v", err))
	}
	return &runningSession{
		remoteSession: session,
	}
}

func (r *runningSession) Run(command string) error {
	err := r.remoteSession.Run(command)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		if ee, ok := err.(*ssh.ExitError); ok {
			return errors.Wrap(err, fmt.Sprintf("Error on remote session return error code %v", ee.ExitStatus()))
		}
		return errors.New("Error on running session, maybe by connection error")
	}
	return nil
}

func main() {
	newSession := Session()
	newSession.Run("pacman")
}
