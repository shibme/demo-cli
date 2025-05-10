FROM scratch
ARG TARGETARCH
COPY ./dist/demo_linux_${TARGETARCH}*/ /
ENTRYPOINT [ "/demo" ]