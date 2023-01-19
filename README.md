# Nuntium

> nuntium -> news in latim

Telegram Bot written in Golang to send news from RSS Feeds

## Docker image

- [Repository in dockerhub](https://hub.docker.com/r/leonardobiffi/nuntium/tags)

```sh
docker pull leonardobiffi/nuntium:latest
```

## Environment Variables

| Name              | Description                           | Default     |
|-------------------|---------------------------------------|-------------|
| SCHEDULE_HOURS    | Perido of time to check feeds         | 1           |
| SKIP_NOTIFICATION | Define with will send to Telegram     | false       |
| CONFIG_FILE       | Filename Feed config file             | config.yml  |
| TELEGRAM_TOKEN    | Token create for Bot                  |   -         |
| TELEGRAM_CHAT_ID  | Chat id as receiver for our messages  |   -         |
