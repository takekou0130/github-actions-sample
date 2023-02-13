payload()
{
cat <<EOF
{
  "text": "$MESSAGE"
}
EOF
}

echo payload
response=$(curl -s -X POST --data-urlencode "payload=$(payload)" "https://www.google.com")

echo $response
