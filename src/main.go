package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var (
	cfg   Config
	queue SQS
)

// try to gracefully stop when signalled to stop
func watchSignal() {
	sigc := make(chan os.Signal, 1)

	signal.Notify(
		sigc,
		os.Interrupt,
		os.Kill,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGKILL,
		syscall.SIGQUIT,
	)

	go func() {
		s := <-sigc
		fmt.Println(s)
		// TODO log this
		time.Sleep(time.Second)
		os.Exit(0)
	}()
}

// Post message to POST_HOST + POST_ENDPOINT and delete from queue if successfull
// if unsuccessful do nothing and the message will be pulled from the queue again
// until the default amount set by the queue
func processMessage(message *sqs.Message, wg *sync.WaitGroup) {
	defer wg.Done()

	// Post to service
	client := http.Client{
		Timeout: cfg.ConnectionTimeout,
	}
	req, err := http.NewRequest("POST", cfg.PostHost+cfg.PostEndpoint, strings.NewReader(*message.Body))
	if err != nil {
		fmt.Printf("Error: %s processing %s", err, *message.ReceiptHandle)
		return
	}

	req.Close = true
	req.Header.Set("Content-Type", cfg.ContentType)

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error: %s, Post to service failed on MessageId: %s", err, *message.ReceiptHandle)
		return
	}
	if res.StatusCode != 200 {
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			body = []byte("")
		}
		fmt.Printf("Error Failed Message. StatusCode: %d, MessageId: %s, ResponseBody: %s", res.StatusCode, *message.ReceiptHandle, body)
		return
	}
	// delete message
	if err = queue.Complete(cfg.SQSURL, *message.ReceiptHandle); err != nil {
		fmt.Printf("Error: %s, Failed to delete MessageId: %s from SQS", err, *message.ReceiptHandle)
		return
	}
}

func backoff(dur time.Duration) time.Duration {
	newWait := dur.Nanoseconds() * 2
	if newWait > cfg.MaxSleep.Nanoseconds() || newWait < 1 {
		return cfg.MaxSleep
	}
	return time.Duration(newWait)
}

// Run the main loop in which messages are processed
func work() {
	dur := time.Millisecond * 10

	wg := &sync.WaitGroup{}
	creds, err := queue.sqsService.Config.Credentials.Get()
	if err == nil {
		fmt.Printf("Using aws provider: %s\n", creds.ProviderName)
	}

	for {
		time.Sleep(dur)
		dur = backoff(dur)
		// get next message batch
		messages, err := queue.NextMessages(cfg.SQSURL)
		if err != nil {
			// TODO log or something
			fmt.Printf("Error: %s, Failed to get next messaged in queue\n", err)
			if strings.Contains(err.Error(), "NoCredentialProviders") {
				os.Exit(1)
			}
			continue
		}

		// No messages
		if len(messages) == 0 {
			fmt.Println("No Messages. Backing off ", dur)
			continue
		}

		fmt.Printf("Processing %d messages...\n", len(messages))

		// Process messages
		for _, message := range messages {
			wg.Add(1)
			go processMessage(message, wg)
		}
		wg.Wait()

		// TODO log success or something
		dur = time.Millisecond * 10
	}
}

func main() {
	initConfig()

	sesh := session.New()

	awsConfig := aws.Config{}
	awsConfig.WithRegion(cfg.Region)
	awsConfig.MaxRetries = aws.Int(10)

	queue = SQS{
		cfg: cfg,
	}
	queue.sqsService = sqs.New(sesh, &awsConfig)
	watchSignal()
	work()
}
