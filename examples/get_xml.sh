ARTICLE_ID=$1
curl -X GET http://localhost:8085/articles/${ARTICLE_ID} \
  -H 'Accept: text/xml'
