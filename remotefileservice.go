package main

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type RemoteFileService struct {
	Key          string
	Secret       string
	HostBase     string
	UrlGenerator string
}

func (p *RemoteFileService) getClient() (*minio.Client, error) {
	return minio.New(p.HostBase, &minio.Options{
		Creds:  credentials.NewStaticV4(p.Key, p.Secret, ""),
		Secure: true,
	})
}

func NewRemoteFile(key string, secret string, hostbase string, urlgenerator string) *RemoteFileService {

	return &RemoteFileService{Key: key,
		Secret:       secret,
		HostBase:     hostbase,
		UrlGenerator: urlgenerator}

}

func (r *RemoteFileService) Upload(bucket string, filename string, contentType string, data []byte) string {
	client, _ := r.getClient()

	ctx := context.Background()
	bucketName := bucket
	objectName := filename
	//filePath := ""
	d := bytes.NewReader(data)
	info, _ := client.PutObject(ctx, bucketName, objectName,
		d, int64(len(data)), minio.PutObjectOptions{ContentType: contentType})
	fmt.Println("Successfully uploaded bytes: ", info)
	url := strings.Replace(r.UrlGenerator, "<BUCKETNAME>", bucketName, -1)
	url = strings.Replace(url, "<FILENAME>", filename, -1)

	return url
}

// func (r *RemoteFileService) ReadRemoteFile(bucket string) []byte, error {
// 	client := r.getClient()

// 	return nil, nil
// }
