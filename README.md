To Run: 

# To build and run the Docker container on Apple Silicon, use the following command:
docker build --platform linux/arm64 -t ebpf-user-agent .
docker run --privileged ebpf-user-agent