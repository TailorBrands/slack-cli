FROM centurylink/ca-certs
ADD slack-cli /
ENTRYPOINT ["/slack-cli"]
