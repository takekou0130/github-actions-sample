payload()
{
cat <<EOF
{
  "text": "$MESSAGE"
}
EOF
}

response=$(curl -s -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer ${GITHUB_TOKEN}" \
  -H "X-GitHub-Api-Version: 2022-11-28" \
  "https://api.github.com/repos/${GITHUB_REPOSITORY}/issues/6" \
  -X PATCH \
  -d "$(payload)")

echo $response
