ARTICLE_ID=$1

curl -v -X PUT http://localhost:8085/articles/${ARTICLE_ID}?token=auth_token_xxx \
  -H 'Content-Type: text/xml' \
  -H 'Accept: text/xml' \
  -d '<data><Title>Title New #?+XML+1556</Title></data>'
