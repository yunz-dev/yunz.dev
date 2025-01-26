#!/bin/bash
# Start both processes

go run main.go &
sudo cloudflared service install "$TUNNEL_TOKEN" &

# Wait for background processes to finish
wait
