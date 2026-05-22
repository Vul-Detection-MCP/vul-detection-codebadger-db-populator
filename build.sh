#!/bin/zsh

docker build -t biggujo/vul-detection-codebadger-db-populator:1.0.0 .
docker push biggujo/vul-detection-codebadger-db-populator:1.0.0
