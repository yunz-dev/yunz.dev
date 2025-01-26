FROM golang:1.20 AS builder

WORKDIR /

COPY start.sh /usr/local/bin/start.sh
RUN chmod +x /usr/local/bin/start.sh

RUN apt-get update && apt-get install -y curl sudo && \
    curl -L --output cloudflared.deb https://github.com/cloudflare/cloudflared/releases/latest/download/cloudflared-linux-amd64.deb && \
    dpkg -i cloudflared.deb && \
    rm cloudflared.deb && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

COPY . /

EXPOSE 8000

CMD ["/usr/local/bin/start.sh"]
