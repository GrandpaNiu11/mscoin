FROM ubuntu:latest
LABEL authors="szr"

ENTRYPOINT ["top", "-b"]