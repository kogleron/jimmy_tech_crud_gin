curl -v -X POST http://localhost:8085/articles?token=auth_token_xxx \
  -H 'Content-Type: application/json' \
  -d '{"title":"Title New #??", "body": "Some body+"}'