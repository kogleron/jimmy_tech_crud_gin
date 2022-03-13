ARTICLE_ID=$1

curl -v -X GET http://localhost:8085/articles/${ARTICLE_ID} \
  -H 'Accept: application/json'
