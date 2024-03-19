# Run instructions:

## To build and run the Docker container on Apple Silicon, use the following command:
```docker build --platform linux/arm64 -t ebpf-user-agent .```

## To execute the MVP run the following command:
```docker run --privileged ebpf-user-agent```

Note: The `--privileged` flag gives the container full access to the host, which can be a security risk. Only use it if you trust the application and understand the implications.