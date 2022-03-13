curl -v -X POST http://localhost:8085/articles \
  -H 'Content-Type: application/json' \
  -d '{"title":"Title New #3", "body": "Some body"}'