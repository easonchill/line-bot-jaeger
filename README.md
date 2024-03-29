# line-bot-jaeger （杰哥）

This line-bot is a demonstration project for CINNOX, showcasing a practical implementation of a chatbot integrated with LINE Messenger.

- Chat with AI
This project features an interactive chat interface that allows users to converse with an AI

## LineID：@496ohpkr

I am running this bot server on GCP, so you can directly add this bot for testing purposes.

## [DEMO VIDEO](https://youtu.be/B0Tn8Y_zmDw) <= Click me

## Tools Used
- Golang with Gin framework
- LINE Messaging API
- MongoDB

## Setup configuration
Due to security concerns, the `.env` and `config.yml` files have been added to the `.gitignore` to prevent the exposure of sensitive keys.

To configure your application, please follow these steps： 

1. Remove the `.test` extension from `config.yml.test` and `.env.test`:
```sh
mv config.yml.test config.yml
mv .env.test .env
```
2. And set **your** line _token_ and _channel Secret_ in `config.yaml`:
```yaml
Line:
  ChannelSecret: YOUR LINE TOKEN
  Token: YOUR LINE SECRET
```


## Setup mongoDB container（use docker-compose）
```sh
make up
```
if you want to delete MongoDB container
```sh
make down
```
## Unit Test
```sh
make test
```
## Local build server
```sh
make build

./line-bot-jaeger
```

## API Doc

### Router

| Method | URL                           | Describe                                       |
| --- | --- | ---------------------------------------------- |
| POST | {{HOST}}/callback     | For line webhook                                   |
| POST | {{HOST}}/sendMessage         | Send message to user                           |
| GET | {{HOST}}/message/{{userid}} | Get a Single User's Message                               |
| GET |  {{HOST}}/message/{{userid}}  | Get All Messages from All Users                               |

### Send massage

**Request**

URL: `{{HOST}}/sendMessage`

Method: `POST`

**Path Variables**

| Parameter    |  Type  | Required | Describe |
| - | :-: | :-: | :-|
| userid     	  | string|Required |Userid in line|
| text     	  | string| Required |Message you want to send to the user|

### Get single user message

**Request**

URL: `{{HOST}}/message/{{userid}}`

Method: `GET`

**Path Variables**

| Parameter    |  Type  | Required | Describe |
| - | :-: | :-: | :-|
| userid     	  | string|Required | line user id |
