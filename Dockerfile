FROM scratch
COPY client_cli_grpc client_cli_grpc
COPY . .
ENTRYPOINT ["/client_cli_grpc"]