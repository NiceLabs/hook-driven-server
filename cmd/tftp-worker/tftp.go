package main

import (
	"errors"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/NiceLabs/hook-driven-server/utils"
	"github.com/pin/tftp/v3"
)

type TFTPConfigure struct {
	Workdir   string `json:"workdir"`
	ReadHook  string `json:"hook_read"`
	WriteHook string `json:"hook_write"`
}

func (t *TFTPConfigure) NewServer() *tftp.Server {
	return tftp.NewServer(t.onRead, t.onWrite)
}

func (t *TFTPConfigure) onRead(filename string, reader io.ReaderFrom) (err error) {
	log.Printf("Read %q\n", filename)
	if t.ReadHook == "" {
		return errors.New("tftp-worker: read operation not permitted")
	}
	cmd := exec.Command(t.ReadHook, filename)
	cmd.Dir = t.Workdir
	utils.AddEnv(cmd, map[string]string{
		"TFTP_ACTION":    "READ",
		"TFTP_READ_FILE": filename,
	})
	cmd.Stderr = os.Stderr
	go func() {
		if stdout, err := cmd.StdoutPipe(); err == nil {
			_, _ = reader.ReadFrom(stdout)
		}
	}()
	return cmd.Run()
}

func (t *TFTPConfigure) onWrite(filename string, writer io.WriterTo) (err error) {
	log.Printf("Write %q\n", filename)
	if t.WriteHook == "" {
		return errors.New("tftp-worker: write operation not permitted")
	}
	cmd := exec.Command(t.WriteHook, filename)
	cmd.Dir = t.Workdir
	utils.AddEnv(cmd, map[string]string{
		"TFTP_ACTION":     "WRITE",
		"TFTP_WRITE_FILE": filename,
	})
	cmd.Stderr = os.Stderr
	go func() {
		if stdin, err := cmd.StdinPipe(); err == nil {
			_, _ = writer.WriteTo(stdin)
		}
	}()
	return cmd.Run()
}
