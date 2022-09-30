package update

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"errors"
	"io"
	"net/http"
	"os"
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
	reader, _ := zip.NewReader(bytes.NewReader(rsp), resp.ContentLength)
	file, err := reader.Open("tdlib.exe")
	if err != nil {
		return err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	err = os.WriteFile("tdlib.exe", data, 0666)
	if err != nil {
		return err
	}
	return nil
}
