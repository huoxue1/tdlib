//go:build !windows
// +build !windows

package update

import (
	"archive/tar"
	"bytes"
	"crypto/sha256"
	"errors"
	"io"
	"net/http"

	"github.com/klauspost/compress/gzip"
)

// update study_xxqg自我更新
func update(url string, sum []byte) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	wc := writeSumCounter{
		hash: sha256.New(),
	}
	rsp, err := io.ReadAll(io.TeeReader(resp.Body, &wc))
	if err != nil {
		return err
	}
	if !bytes.Equal(wc.hash.Sum(nil), sum) {
		return errors.New("文件已损坏")
	}
	gr, err := gzip.NewReader(bytes.NewReader(rsp))
	if err != nil {
		return err
	}
	tr := tar.NewReader(gr)
	for {
		header, err := tr.Next()
		if err != nil {
			return err
		}
		if header.Name == "tdlib" {
			data, err := io.ReadAll(tr)
			if err != nil {
				return err
			}
			err = os.WriteFile("tdlib", data, 0666)
			if err != nil {
				return err
			}
			return nil
		}
	}
}
