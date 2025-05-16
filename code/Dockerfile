FROM golang:1.23.6-bookworm

WORKDIR /home/code

COPY . .

# Installs into /go/bin
# sqlc-gen-go-server does not work with sqlc 1.29.0 but does with 1.25.0
# RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.25.0
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.29.0
RUN go install github.com/fdietze/sqlc-gen-from-template@v1.0.0

# installed for sqlc-gen-go-server
RUN go install golang.org/x/tools/cmd/goimports@latest

RUN go install github.com/walterwanderley/sqlc-http@latest

# RUN chmod +777 /proc/sys/fs/protected_symlinks
# RUN echo 0 > /proc/sys/fs/protected_symlinks

# RUN git clone https://github.com/sqlc-dev/sqlc /home/code/sqlc
