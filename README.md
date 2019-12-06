# dmz_ai.go
[![Go Report Card](https://goreportcard.com/badge/github.com/Sw-Saturn/dmz_ai.go?style=flat-square)](https://goreportcard.com/report/github.com/Sw-Saturn/dmz_ai.go?style=flat-square)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Sw-Saturn/dmz_ai.go?style=for-the-badge)
![Azure DevOps builds](https://img.shields.io/azure-devops/build/e8222/451756ff-78e6-4f82-96f2-cef50ad6fc84/3?label=Azure%20Pipelines&style=for-the-badge)
![Docker Pulls](https://img.shields.io/docker/pulls/swsaturn/dmz_ai.go?style=for-the-badge)
![GitHub](https://img.shields.io/github/license/Sw-Saturn/dmz_ai.go?style=for-the-badge)

## Introduction
自分のTwitterアカウントの発言をベースにして発言させるbot

Twitterの発言を形態素解析して，マルコフ連鎖によって言語化し，TwitterAPIから投稿する．

## Installing

Use the go Modules to install dependencies.

```bash
go mod download
```

## Deployment
Use the Azure AppService to deploy.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
This project is licensed under the Apache License - see the LICENSE.md file for details.
