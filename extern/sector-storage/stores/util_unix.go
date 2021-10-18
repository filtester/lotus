package stores

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/mitchellh/go-homedir"
	"golang.org/x/xerrors"
)

func move(from, to string) error {
	from, err := homedir.Expand(from)
	if err != nil {
		return xerrors.Errorf("move: expanding from: %w", err)
	}

	to, err = homedir.Expand(to)
	if err != nil {
		return xerrors.Errorf("move: expanding to: %w", err)
	}

	if filepath.Base(from) != filepath.Base(to) {
		return xerrors.Errorf("move: base names must match ('%s' != '%s')", filepath.Base(from), filepath.Base(to))
	}

	log.Debugw("move sector data", "from", from, "to", to)

	toDir := filepath.Dir(to)

	// `mv` has decades of experience in moving files quickly; don't pretend we
	//  can do better

	var errOut bytes.Buffer

	var cmd *exec.Cmd
	if runtime.GOOS == "darwin" {
		if err := os.MkdirAll(toDir, 0777); err != nil {
			return xerrors.Errorf("failed exec MkdirAll: %s", err)
		}

		cmd = exec.Command("/usr/bin/env", "mv", from, toDir) // nolint
	} else {
		cmd = exec.Command("/usr/bin/env", "mv", "-t", toDir, from) // nolint
	}

	cmd.Stderr = &errOut
	if err := cmd.Run(); err != nil {
		return xerrors.Errorf("exec mv (stderr: %s): %w", strings.TrimSpace(errOut.String()), err)
	}

	return nil
}

func MvFromLocal(url, dest string) error {
	prefixStr := "remote/"
	filename := url[strings.Index(url, prefixStr)+len(prefixStr):]
	param := fmt.Sprintf("du -m /mnt/lotus*/worker*/lotusworker/%v", filename)

	cmd := exec.Command("sh", "-c", param)
	out, err := cmd.CombinedOutput()

	log.Debugf("MvFromLocal,du out: %s ", out)
	if err != nil {
		return xerrors.Errorf("MvFromLocal du error, url:%s, %s, %w", url, param, err)
	}

	filesStr := string(out)
	files := strings.Split(filesStr, "\n")

	if len(files) != 2 {
		return xerrors.Errorf("MvFromLocal file count not right, url:%s, count:%v", url, len(files))
	}
	rows := strings.Split(files[0], "\t")
	fileSize, err := strconv.ParseInt(rows[0], 10, 64)
	if fileSize <= 0 {
		return xerrors.Errorf("MvFromLocal file size not right, url:%s, count:%v", url, fileSize)
	}
	srcFile := rows[1]

	log.Infof("MvFromLocal,move %s -> %s ", srcFile, out)

	cmdMv := exec.Command("mv", srcFile, dest)
	out, errMv := cmdMv.CombinedOutput()
	if errMv != nil {
		return xerrors.Errorf("MvFromLocal mv failed, url:%s,%w", url, errMv)
	}
	return nil
}
