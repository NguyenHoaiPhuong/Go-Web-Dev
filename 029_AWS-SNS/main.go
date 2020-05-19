package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func listSNSTopics() {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	// sess := session.Must(session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// }))
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
	})
	if err != nil {
		fmt.Println("NewSession error:", err)
		return
	}

	svc := sns.New(sess)

	result, err := svc.ListTopics(nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, t := range result.Topics {
		fmt.Println(*t.TopicArn)
	}
}

func listSubscriptions() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
	})
	if err != nil {
		fmt.Println("NewSession error:", err)
		return
	}

	svc := sns.New(sess)

	result, err := svc.ListSubscriptions(nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, s := range result.Subscriptions {
		fmt.Println(*s.SubscriptionArn)
		fmt.Println("  " + *s.TopicArn)
		fmt.Println("")
	}
}

func sendSNSTopicToAllSubcribers() {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
	})

	if err != nil {
		fmt.Println("NewSession error:", err)
		return
	}

	client := sns.New(sess)
	input := &sns.PublishInput{
		Message:  aws.String("Hello world!"),
		TopicArn: aws.String("arn:aws:sns:ap-southeast-1:786598396932:Tokoin-staging"),
	}

	result, err := client.Publish(input)
	if err != nil {
		fmt.Println("Publish error:", err)
		return
	}

	fmt.Println(result)
}

func sendSMSToPhone() {
	fmt.Println("creating session")
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
	}))
	fmt.Println("session created")

	svc := sns.New(sess)
	fmt.Println("service created")

	params := &sns.PublishInput{
		Message:     aws.String("testing 123"),
		PhoneNumber: aws.String("+84372141603"),
	}
	resp, err := svc.Publish(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

// usage:
// go run sns_publish_to_topic.go
func main() {
	// listSNSTopics()
	// listSubscriptions()
	sendSMSToPhone()
}
