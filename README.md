docker-index
============

Docker index は Docker registry に保存されているコンテナを一覧表示するツールです。

Docker index shows Docker registry container images. 

- [Docker Hub](https://registry.hub.docker.com/u/ringo/docker-index/)
- [Git Hub](https://github.com/ringohub/docker-index)

## Usage

### Build container

```
docker build --rm -t ringohub/docker-index .
```

### Run

#### Environment

- `DOCKER_REGISTRY_URL`: Docker Registry URL
- `DOCKER_REPOSITORY_PATH`: Docker Repository Path(default: `/var/registry/repos`)

```
docker run -v /var/registry/repositories:/var/tmp/repos -e DOCKER_REGISTRY_URL=foo.example.com -p 9000:9000 -ti ringohub/docker-index:latest
```

## Screen Shot

![docker-index](http://i.gyazo.com/fe2de79c089269253eea079e19840765.png)
