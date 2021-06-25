package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
  // Fetch env vars
	connectionString := os.Getenv("CONNECTION_STRING")
	key := os.Getenv("SPACES_KEY")
	secret := os.Getenv("SPACES_SECRET")
	bucketEndpoint := os.Getenv("BUCKET_ENDPOINT")
	bucketRegion := os.Getenv("BUCKET_REGION")
	bucketName := os.Getenv("BUCKET_NAME")

	// Get a dump of the database
	println("Getting dump")
	cmd("mongodump", fmt.Sprintf("--uri=%s", connectionString), "--forceTableScan")

	// // Zip the dump
	t := time.Now()
	timestamp := t.Format("2006-01-02")	
	dumpName := fmt.Sprintf("%s.zip",timestamp)
	println("Zipping dump")
	cmd("zip", "-r", dumpName, "./dump")

	data, err := os.Open(dumpName)
	check(err)
	
	s3Config := &aws.Config{
			Credentials: credentials.NewStaticCredentials(key, secret, ""),
			Endpoint:    aws.String(bucketEndpoint),
			Region:      aws.String(bucketRegion),
	}

	newSession := session.New(s3Config)
	s3Client := s3.New(newSession)

	println("Uploading dump")
	object := s3.PutObjectInput{
    Bucket: aws.String(bucketName),
    Key:    aws.String("backups/" + dumpName),
    Body:   data,
    ACL:    aws.String("private"),
	}
	output, err := s3Client.PutObject(&object)
	check(err)

	println(output)

	// Cleanup
	cmd("rm", "-rf", "./dump")
	cmd("rm", dumpName)

	println(fmt.Sprintf("Dump uploaded to: backups/%s", dumpName))
}

func cmd(name string, args ...string) {
	cmd := exec.Command(name, args...)
	stdout, err := cmd.Output()
	check(err)

	println(string(stdout))
}

func check(err error) {
	if err != nil {
			log.Fatal(err)
	}
}