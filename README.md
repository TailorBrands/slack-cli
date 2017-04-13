### slack-cli

A scratch container that sends messages to Slack

Mainly used as a way to send notifications from [Kubernetes Init Containers](https://kubernetes.io/docs/concepts/workloads/pods/init-containers/)

An example:
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: myapp-pod
  labels:
    app: myapp
  annotations:
    pod.beta.kubernetes.io/init-containers: '[
        {
            "name": "init-myapp",
            "image": "TailorBrands/slack-cli",
            "command": ["slack-cli", "-c", "alerts", "-u", "slack-cli", "--webhook-url", "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX
", "myapp is running!"]
        }
    ]'
spec:
  containers:
  - name: myapp-container
    image: busybox
    command: ['sh', '-c', 'echo The app is running! && sleep 3600']
```


## Usage:
```
  slack-cli [flags] [string to send]

Flags:
  -c, --channel string       slack channel to send the message to (default "alerts")
  -i, --icon-emoji string    user's icon emoji
  -u, --username string      slack username that sends out the message (default "slack-cli")
      --webhook-url string   slack webhook url
```
