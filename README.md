# Nuntium

> nuntium -> news in latim

Telegram Bot written in Golang to send news from RSS Feeds

## Docker image

- [Repository in dockerhub](https://hub.docker.com/r/leonardobiffi/nuntium/tags)

```sh
docker pull leonardobiffi/nuntium:latest
```

## Install Locally

Downloads the app based on your OS/arch and puts it in `/usr/local/bin`.
Needed [jq](https://stedolan.github.io/jq/download) instaled.

```sh
curl -fsSL https://raw.githubusercontent.com/leonardobiffi/nuntium/master/scripts/install.sh | sh
```

## Environment Variables

| Name              | Description                                 | Default     |
|-------------------|---------------------------------------------|-------------|
| SCHEDULE_HOURS    | Perido of time to check feeds               | 1           |
| SKIP_NOTIFICATION | Define with will send to Telegram           | false       |
| CONFIG_FILE       | Filename Feed config file                   | config.yml  |
| CONFIG_VALUE      | Load config Feed from environment variable  |   -         |
| TELEGRAM_TOKEN    | Token create for Bot                        |   -         |
| TELEGRAM_CHAT_ID  | Chat id as receiver for our messages        |   -         |

## Configuration Feed URLs

If config yml file not exist, the will try load from `CONFIG_VALUE` set in environments variables.
