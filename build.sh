#!/bin/zsh

docker build --platform linux/amd64,linux/arm64 -t biggujo/vul-detection-codebadger-db-populator:1.0.2 .
docker push biggujo/vul-detection-codebadger-db-populator:1.0.2
