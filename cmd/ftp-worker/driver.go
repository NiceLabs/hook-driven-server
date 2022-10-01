package main

import (
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/NiceLabs/hook-driven-server/utils"
	"goftp.io/server/core"
)

type HookDriver struct {
	Workdir   string
	ReadHook  string
	WriteHook string
	writeMaps map[string]io.ReaderFrom
}

func (h *HookDriver) Stat(path string) (core.FileInfo, error) {
	return DummyFileInfo(path), nil
}

func (h *HookDriver) ListDir(string, func(core.FileInfo) error) error {
	return errors.New("ftp-worker: list directory operation not permitted")
}

func (h *HookDriver) DeleteDir(string) error {
	return errors.New("ftp-worker: delete directory operation not permitted")
}

func (h *HookDriver) DeleteFile(string) error {
	return errors.New("ftp-worker: delete file operation not permitted")
}

func (h *HookDriver) Rename(string, string) error {
	return errors.New("ftp-worker: rename operation not permitted")
}

func (h *HookDriver) MakeDir(string) error {
	return errors.New("ftp-worker: make directory operation not permitted")
}

func (h *HookDriver) GetFile(path string, offset int64) (n int64, stdout io.ReadCloser, err error) {
	log.Printf("Read %q with %d offset\n", path, offset)
	if h.ReadHook == "" {
		err = errors.New("ftp-worker: read operation not permitted")
		return
	}
	cmd := exec.Command(h.ReadHook, path, strconv.FormatInt(offset, 64))
	utils.AddEnv(cmd, map[string]string{
		"FTP_ACTION":      "READ",
		"FTP_PATH":        path,
		"FTP_READ_OFFSET": strconv.FormatInt(offset, 64),
	})
	cmd.Dir = h.Workdir
	cmd.Stderr = os.Stderr
	if err = cmd.Start(); err != nil {
		return
	}
	stdout, err = cmd.StdoutPipe()
	return
}

func (h *HookDriver) PutFile(destPath string, data io.Reader, appendData bool) (n int64, err error) {
	if appendData {
		log.Printf("Write %q\n with append data", destPath)
	} else {
		log.Printf("Write %q\n", destPath)
	}
	if h.WriteHook == "" {
		err = errors.New("ftp-worker: write operation not permitted")
		return
	}
	if appendData && h.writeMaps[destPath] != nil {
		return h.writeMaps[destPath].ReadFrom(data)
	}
	cmd := exec.Command(h.WriteHook, destPath)
	utils.AddEnv(cmd, map[string]string{
		"FTP_ACTION": "WRITE",
		"FTP_PATH":   destPath,
	})
	cmd.Dir = h.Workdir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	var stdin bytes.Buffer
	n, err = stdin.ReadFrom(data)
	cmd.Stdin = &stdin
	h.writeMaps[destPath] = &stdin
	go func() {
		_ = cmd.Run()
		delete(h.writeMaps, destPath)
	}()
	return
}
