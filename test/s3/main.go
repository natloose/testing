package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)


var ( s3session *s3.S3 )

const (
	REGION = "eu-west-1"
)

func init() {
	 // By default NewSession will only load credentials from the shared credentials file (~/.aws/credentials).

	// Creating a Session without additional options will load credentials region, and profile loaded from the environment and shared
	// config automatically.

	// This will create a session 
	// sess, err := session.NewSession()

	// When creating Sessions optional aws.Config values can be passed in that will override the default, or loaded, config values the Session is being created with.
	// This allows you to provide additional, or case based, configuration as needed.

	// Create a Session with an optional aws.Config value, region eu-west-1
	// session.Must() is a helper function to ensure the Session is valid and there was no error when calling a NewSession function.
	// This helper is intended to be used in variable initialization to load the Session and configuration at startup.
	mySession := session.Must(session.NewSession())

	// Create a S3 client with additional configuration
	s3session = s3.New(mySession, aws.NewConfig().WithRegion(REGION))

}
func listBuckets() (resp *s3.ListBucketsOutput){
	resp, err := s3session.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		fmt.Println(err)
		return
	}
	return resp
}

func createBucket() (resp *s3.CreateBucketOutput) {
	bucketName := "create-bucket-golang-sdk"
	resp, err := s3session.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String(REGION),
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf(bucketName + " successfully created")
	return resp
}
func deleteBucket() (resp *s3.DeleteBucketOutput) {
	bucketName := "create-bucket-golang-sdk"
	resp, err := s3session.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(bucketName)})

	if err != nil {
		fmt.Println(err)
		return
	}


	fmt.Printf(bucketName + " bucket successfully deleted")
	return resp
}
func main() {
	fmt.Println(listBuckets())
	fmt.Println(createBucket())
	fmt.Println(listBuckets())
	fmt.Println(deleteBucket())
	fmt.Println(listBuckets())
}
