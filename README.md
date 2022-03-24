# cf-firestore-template

## Architecture
Please follow the below format to name custom cloud functions.
```
.
├── README.md
├── .env.yaml
├── .gcloudignore
├── cmd
│   └── main.go
├── deploy.sh
├── init.go
├── cf_{funcName}_{trigger}.go
├── cf_{funcName}_{trigger}_test.go
├── ...
├── go.mod
├── go.sum
└── vendor
```

## Quick Start

## Deployment
- Local deploy
```shell
sh ./deploy.sh
```