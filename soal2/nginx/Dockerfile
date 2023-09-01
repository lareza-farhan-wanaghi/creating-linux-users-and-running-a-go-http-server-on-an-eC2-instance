# Use the official nginx image as the base image
FROM nginx:latest

# Copy the custom nginx configuration file to the appropriate location
COPY nginx.conf /etc/nginx/nginx.conf

# Copy the SSL certificates and keys to the nginx SSL directory
COPY ssl /etc/nginx/ssl
