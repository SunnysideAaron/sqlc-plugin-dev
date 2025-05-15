FROM golang:1.23.6-bookworm

WORKDIR /home/code

COPY . .

# Installs into /go/bin
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.29.0
RUN go install github.com/fdietze/sqlc-gen-from-template@v1.0.0


# RUN chmod +777 /proc/sys/fs/protected_symlinks
# RUN echo 0 > /proc/sys/fs/protected_symlinks

# RUN git clone https://github.com/sqlc-dev/sqlc /home/code/sqlc
