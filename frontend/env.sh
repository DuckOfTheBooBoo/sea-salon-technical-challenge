#!/bin/sh
set -e

# Generate a file that contains the environment variable
echo "window.VITE_API_URL='${VITE_API_URL}';" > /usr/share/nginx/html/env-config.js

exec "$@"
