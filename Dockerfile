FROM golang:1.11

LABEL "name"="Annotate GitHub Release"
LABEL "maintainer"="Christoph Blecker <admin@toph.ca>"
LABEL "version"="0.0.1"

LABEL "com.github.actions.name"="Annotate GitHub Release"
LABEL "com.github.actions.description"="Annotate a release event with a specific body"
LABEL "com.github.actions.icon"="edit"
LABEL "com.github.actions.color"="green"

COPY main.go go.mod go.sum /

ENTRYPOINT ["go"]
CMD ["run", "/main.go"]
