# mackerel-notification-alerts

![20151120105942_original](https://user-images.githubusercontent.com/5633085/56470980-da756500-6487-11e9-9b27-b93d46e75ef4.png)

This is a program to notify slack of mackerel's open alert.

## Used

https://github.com/mackerelio/mackerel-client-go

used mackerel-client-go.

## How to

```
$ go get github.com/ashwanthkumar/slack-go-webhook
$ go get github.com/mackerelio/mackerel-client-go
$ export MACKEREL_APIKEY="XXXXXXXX"
add slack webhook,cannnel
$ go run main.go
```

## Slack

![スクリーンショット_2019-05-06_20_09_50](https://user-images.githubusercontent.com/5633085/57221751-3cfb6300-703b-11e9-955b-61c468a920a3.jpg)

