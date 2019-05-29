FROM golang:1

LABEL "com.github.actions.name"="Mention Notifier"
LABEL "com.github.actions.description"="Mention Notifier"
LABEL "com.github.actions.icon"="at-sign"
LABEL "com.github.actions.color"="red"
LABEL "repository"="https://github.com/lowply/mention-notifier"
LABEL "homepage"="https://github.com/lowply/mention-notifier"
LABEL "maintainer"="Sho Mizutani <lowply@github.com>"

WORKDIR /go/src
COPY src .
RUN GO111MODULE=on go build -o /go/bin/main

ADD entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
