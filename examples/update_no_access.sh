curl -v -X PUT http://localhost:8085/articles/11 \
  -H 'Content-Type: application/json' \
  -d '{"title":"Title New #?+5"}'