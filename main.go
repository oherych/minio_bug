package main

import (
	"bytes"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/minio/minio-go/pkg/encrypt"

	"github.com/minio/minio-go"
)

var (
	letterRunes = string("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s3Client    *minio.Client
	rsaKey      *encrypt.CBCSecureMaterials
)

const (
	enpoint      = "localhost:9000"
	accessKey    = "BH0K3LEZSZX2KFM53LLS"
	accessSecret = "Pzfn+HrTbw+oPO8Tz5NFnj/1RbWSjH1qQ+cqCJE6"
	secure       = false
	bucket       = "some-test"
	location     = "eu-central-1"
	testfile     = "test_file"
	server       = ":8000"
)

func main() {
	initMinioClient()
	initHTTPServer()
}

func handlerPage(writer http.ResponseWriter, request *http.Request) {
	defer rec(writer)

	b, err := os.Open("index.html")
	if err != nil {
		log.Error(err)
		return
	}

	defer b.Close()

	writer.WriteHeader(http.StatusOK)
	if _, err := io.Copy(writer, b); err != nil {
		log.Error(err)
		return
	}
}

func handlerFile(writer http.ResponseWriter, request *http.Request) {
	defer rec(writer)

	obj, err := s3Client.GetObject(bucket, testfile, minio.GetObjectOptions{Materials: rsaKey})
	if err != nil {
		log.Error(err)
		return
	}

	// defer obj.Close()
	defer func() {
		obj.Close()
	}()

	_, err = obj.Stat()
	if err != nil {
		log.Error(err)
		return
	}

	b, err := ioutil.ReadAll(obj)
	if err != nil {
		log.Error(err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(fmt.Sprintln(len(b))))
}

func initHTTPServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerPage)
	mux.HandleFunc("/file", handlerFile)

	err := http.ListenAndServe(server, mux)
	if err != nil {
		panic(err)
	}
}

func initMinioClient() {
	rsaPrivateKey, err := ioutil.ReadFile("ssh/storage")
	if err != nil {
		panic(err)
	}

	rsaPublicKey, err := ioutil.ReadFile("ssh/storage.pub")
	if err != nil {
		panic(err)
	}

	rsaKey, err = createMaterials(rsaPrivateKey, rsaPublicKey)
	if err != nil {
		panic(err)
	}

	s3Client, err = minio.New(enpoint, accessKey, accessSecret, secure)
	if err != nil {
		panic(err)
	}

	if ok, err := s3Client.BucketExists(bucket); !ok {
		if err != nil {
			panic(err)
		}

		err = s3Client.MakeBucket(bucket, location)
		if err != nil {
			panic(err)
		}
	}

	reader := bytes.NewReader(testFile(10000))

	_, err = s3Client.PutObject(bucket, testfile, reader, reader.Size(), minio.PutObjectOptions{EncryptMaterials: rsaKey})
	if err != nil {
		panic(err)
	}
}

func testFile(size int) []byte {
	data := make([]byte, size)
	for i := 0; i < size; i++ {
		data[i] = letterRunes[i%len(letterRunes)]
	}
	return data
}

func createMaterials(privateKey, publicKey []byte) (*encrypt.CBCSecureMaterials, error) {
	privateKey = keyNormalize(privateKey)
	publicKey = keyNormalize(publicKey)

	key, err := encrypt.NewAsymmetricKey(privateKey, publicKey)
	if err != nil {
		return nil, err
	}

	return encrypt.NewCBCSecureMaterials(key)
}

func keyNormalize(data []byte) []byte {
	block, _ := pem.Decode(data)
	if block == nil {
		return data
	}

	return block.Bytes
}

func rec(_ http.ResponseWriter) {
	if r := recover(); r != nil {
		log.Error(r)
	}
}
