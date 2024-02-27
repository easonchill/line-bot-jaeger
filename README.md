# line-bot-jaeger （杰哥）

This line-bot is a demonstration project for CINNOX, showcasing a practical implementation of a chatbot integrated with LINE Messenger.

- Chat with AI
This project features an interactive chat interface that allows users to converse with an AI

- LineID：@496ohpkr

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
