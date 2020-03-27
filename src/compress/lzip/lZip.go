package lzip

import (
	"archive/zip"
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
)

/*
 *	Descryption: compress zip file
 *	params:
 *		zipName: zip file name.
 *		file:	 compress file name.
 *		body:	 file content.
 */
func lZip(zipName string, file string, body string) error {
	buf := new(bytes.Buffer)

	w := zip.NewWriter(buf)

	var files = []struct {
		Name, Body string
	}{
		{file, body},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
			return err
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	err := w.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}

	f, err := os.OpenFile(zipName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, errFile := buf.WriteTo(f)
	if errFile != nil {
		log.Fatal(errFile)
		return err
	}
	return nil
}

/*
 *	Descryption: unCompress zip file
 *	params:
 *		zipName: zip file name.
 *		targetPath:	 target path.
 */
func lUnZip(zipName string, targetPath string) error {
	reader, err := zip.OpenReader(zipName)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(targetPath, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(targetPath, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}

/*
 *	Descryption: judge zip file
 *	params:
 *		zipName: zip file name.
 *		target:	 target path.
 */
func lIsZip(zipPath string) bool {
	f, err := os.Open(zipPath)
	if err != nil {
		return false
	}
	defer f.Close()

	buf := make([]byte, 4)
	if n, err := f.Read(buf); err != nil || n < 4 {
		return false
	}

	return bytes.Equal(buf, []byte("PK\x03\x04"))
}
