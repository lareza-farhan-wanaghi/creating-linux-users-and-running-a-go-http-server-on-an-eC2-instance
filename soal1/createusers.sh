#!/bin/bash

# Loop through creating 200 users
for i in {1..200}; do
    # Define the username based on the loop index
    username="user$i"
    
    # Calculate the UID based on the loop index
    uid=$((1500 + i - 1))
    
    # Generate a password using the UID and username
    password="$uid$username"
    
    # Add a new user with the specified UID, home directory, and hashed password
    sudo useradd -m -u $uid -p $(openssl passwd -1 $password) $username
    
    # Prepare and send a POST request to the webhook URL
    curl -X POST -H 'Content-Type: application/json' -H "Authorization: Basic $(echo -n $username:$password | base64)" -d "{\"username\":\"$username\",\"password\":\"$password\"}" <YOUR WEBHOOK URL>
    
    # Print a completion message for the current user iteration
    echo "completed $i"
done
