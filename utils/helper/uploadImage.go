package helper

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"strings"
	"warehouse/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func CheckFileExtension(filename string) (string, error) {
	extension := strings.ToLower(filename[strings.LastIndex(filename, ".")+1:])

	if extension != "jpg" && extension != "jpeg" && extension != "png" {
		return "", fmt.Errorf("forbidden file type")
	}

	return extension, nil
}

func CheckFileSize(size int64) error {
	if size == 0 {
		return fmt.Errorf("illegal file size")
	}
	if size > 1097152 {
		return fmt.Errorf("file size too big")
	}

	return nil
}

func UploadFileToS3(directory string, fileName string, contentType string, fileData multipart.File) (string, error) {

	// The session the S3 Uploader will use
	sess := config.GetSession()

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("AWS_BUCKET")),
		Key:         aws.String("/" + directory + "/" + fileName),
		Body:        fileData,
		ContentType: aws.String(contentType),
	})

	if err != nil {
		log.Print(err.Error())
		return "", fmt.Errorf("failed to upload file")
	}
	return result.Location, nil
}
