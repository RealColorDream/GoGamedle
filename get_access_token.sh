#!/bin/bash

source ./.env

client_id=$CLIENT_ID
client_secret=$CLIENT_SECRET

response=$(curl -X POST 'https://id.twitch.tv/oauth2/token' \
     -d "client_id=$client_id" \
     -d "client_secret=$client_secret" \
     -d 'grant_type=client_credentials')

token=$(echo "$response" | jq -r '.access_token')

echo "ACCESS_TOKEN=$token" > .env.token

echo "Access token saved to .env.token file."

