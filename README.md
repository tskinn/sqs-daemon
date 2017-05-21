

Env Vars:

- ACCESS - aws access key; if not provided, sqsd falls back to local config or role for credentials
- SECRET - aws access secret;
- SQS_URL - url of sqs queue (required)
- POST_ENDPOINT - url to post messages to (required)
- POST_HOST - host to post messages to; defaults to "http://127.0.0.1:80"
- REGION - aws region; defaults to "us-east-1"
- CONTENT_TYPE - content type of post messages; defaults to "application/json"
- MAX_SLEEP - max time to sleep if no messages in queue in seconds; defaults to 300(5 minutes)
- CONNECTION_TIMEOUT - connection time of posts in seconds; defaults to 300
- CONNNECTIONS - number of messages to retrieve from queue; defaults to 1, max is 10
