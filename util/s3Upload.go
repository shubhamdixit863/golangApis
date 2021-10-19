package util

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

var (
	AWS_S3_REGION         string
	AWS_S3_BUCKET         string
	AWS_ACCESS_KEY_ID     string
	AWS_SECRET_ACCESS_KEY string
)

func init() {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	AWS_S3_REGION = os.Getenv("AWS_S3_REGION")
	AWS_S3_BUCKET = os.Getenv("AWS_S3_BUCKET")
	AWS_ACCESS_KEY_ID = os.Getenv("AWS_ACCESS_KEY_ID")
	AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")

}

func UploadS3(uploadFileDir string) string {

	session, err := session.NewSession(&aws.Config{Region: aws.String(AWS_S3_REGION),
		Credentials: credentials.NewStaticCredentials(
			AWS_ACCESS_KEY_ID,
			AWS_SECRET_ACCESS_KEY,
			"", // a token will be created when the session it's used.
		),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Upload Files
	filePath, err := uploadFile(session, uploadFileDir)
	if err != nil {
		log.Fatal(err)
	}

	return filePath
}

func uploadFile(session *session.Session, uploadFileDir string) (string, error) {
	var filepath string

	upFile, err := os.Open(uploadFileDir)
	if err != nil {
		return filepath, err
	}
	defer upFile.Close()

	upFileInfo, _ := upFile.Stat()
	var fileSize int64 = upFileInfo.Size()
	fileBuffer := make([]byte, fileSize)
	upFile.Read(fileBuffer)

	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(AWS_S3_BUCKET),
		Key:           aws.String(uploadFileDir),
		ACL:           aws.String("public-read"),
		Body:          bytes.NewReader(fileBuffer),
		ContentLength: aws.Int64(fileSize),
		ContentType:   aws.String(http.DetectContentType(fileBuffer)),
		//ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	if err != nil {
		return filepath, err

	}
	filepath = "https://" + AWS_S3_BUCKET + "." + "s3-" + AWS_S3_REGION + ".amazonaws.com/" + uploadFileDir

	return filepath, nil
}
