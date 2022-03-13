ARTICLE_ID=$1
curl -v -X DELETE http://localhost:8085/articles/${ARTICLE_ID}?token=auth_token_xxx