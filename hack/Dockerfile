ARG IMAGE_REGISTRY
FROM ${IMAGE_REGISTRY}docker.io/library/golang:1.21

WORKDIR /workspace

RUN git clone --branch=v0.28.1 --depth=1 https://github.com/kubernetes/code-generator.git $GOPATH/src/k8s.io/code-generator
