FROM golang:bullseye

RUN set -xe \
  && apt-get update \
  && apt-get install -y \
    libauthen-ntlm-perl libcgi-pm-perl libcrypt-openssl-rsa-perl libdata-uniqid-perl libencode-imaputf7-perl \
    libfile-copy-recursive-perl libfile-tail-perl libio-compress-perl libio-socket-ssl-perl libio-socket-inet6-perl \
    libio-tee-perl libhtml-parser-perl libjson-webtoken-perl libmail-imapclient-perl libparse-recdescent-perl \
    libmodule-scandeps-perl libpar-packer-perl libproc-processtable-perl libreadonly-perl libregexp-common-perl \
    libsys-meminfo-perl libterm-readkey-perl libtest-mockobject-perl libtest-pod-perl libunicode-string-perl \
    liburi-perl libwww-perl procps wget make cpanminus lsof ncat openssl ca-certificates \
  && rm -rf /var/lib/apt/lists/* \
  && cpanm IO::Socket::SSL

RUN set -xe \
  && wget -N -P /usr/local/bin --no-check-certificate https://imapsync.lamiral.info/imapsync \
  && chmod +x /usr/local/bin/imapsync \
  && /usr/local/bin/imapsync --testslive
  #&& /usr/local/bin/imapsync --tests

WORKDIR /app

RUN go install github.com/githubnemo/CompileDaemon@latest

CMD [ \
    "CompileDaemon", \
    #"--include", ".env", \
    "--build", "go build -ldflags=-s -o mailbox-sync", \
    "--command", "./mailbox-sync ./mailboxes.json" \
]