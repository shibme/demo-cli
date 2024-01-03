FROM scratch
COPY demo-cli /
ENTRYPOINT ["/demo-cli"]