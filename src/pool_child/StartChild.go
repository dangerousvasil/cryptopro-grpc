package pool_child

import (
	"cryptopro-jsonrpc/src/rwCloser"
	"errors"
	"io"
	"os"
	"os/exec"
	"syscall"
)

var SERVICE_HI_MSG = `SERVICE_HI_MSG`

// StartChild стартуем потомка
func StartChild(binFile string) (*exec.Cmd, error) {

	// запускаем потомка
	var cmd = exec.Command(binFile, []string{`child`}...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Pdeathsig: syscall.SIGINT,
		//Ptrace:    true,
		Setpgid: true,
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}
	go io.Copy(os.Stderr, stderr)

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	// ждем приглашение от потомка что он запустился
	hi, _, err := rwCloser.BufferReader(stdout)
	if err != nil {
		return nil, err
	}
	if SERVICE_HI_MSG != string(hi)[:len(SERVICE_HI_MSG)] {
		return nil, errors.New("unexpected hi message " + string(hi))
	}
	//log.Println(`greeting accepted`)

	go io.Copy(os.Stdout, stdout)

	return cmd, err
}
