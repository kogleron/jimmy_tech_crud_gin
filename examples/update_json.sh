ARTICLE_ID=$1

curl -v -X PUT http://localhost:8085/articles/${ARTICLE_ID}?token=auth_token_xxx \
  -H 'Content-Type: application/json' \
  -d '{"title":"Title New #?+6"}'