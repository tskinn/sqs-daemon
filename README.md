## SQS-Daemon
Another implementation of the sqs daemon used in AWS ElasticBeanstalk worker environment and is inspired by many similar projects like https://github.com/mozart-analytics/sqsd. This implementation is written in golang and is currently used to fuel workers running in docker on AWS ECS. See the example task-definition.json for reference on how this works.

![Sqsd architecture](http://docs.aws.amazon.com/elasticbeanstalk/latest/dg/images/aeb-messageflow-worker.png)

## Configuration:
| Env Variable | Description |
| --- | --- |
| ACCESS | aws access key; if not provided, sqs-daemon falls back to local config or role for credentials |
| SECRET | aws access secret; if not provided, sqs-daemon falls back to local config or role for credentials |
| SQS_URL | url of sqs queue (required) |
| POST_ENDPOINT | url to post messages to (required) |
| POST_HOST | host to post messages to; defaults to "http://127.0.0.1:80" |
| REGION | aws region; defaults to "us-east-1" |
| CONTENT_TYPE | content type of post messages; defaults to "application/json" |
| MAX_SLEEP | max time to sleep if no messages in queue in seconds; defaults to 300(5 minutes) |
| CONNECTION_TIMEOUT | time to wait for response from worker; defaults to 300 |
| CONNNECTIONS | number of messages to retrieve from queue at a time; defaults to 1, max is 10 |

## Build

### Binary
`go get -u && GOOS=linux go build -o bin/linux/sqs-daemon . `

### Docker
``` 
docker build -t repo-name/sqs-daemon:tag .
docker push repo-name/sqs-daemon:tag
```
