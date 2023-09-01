#!/bin/bash

# Loop through user numbers from 1 to 200
for i in {1..200}; do
    # Generate username based on the loop index
    username="user$i"
    
    # Check if the user exists in the /etc/passwd file
    grep "^$username:" /etc/passwd
done
