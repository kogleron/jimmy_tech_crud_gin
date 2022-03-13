curl -X POST http://localhost:8085/articles?token=auth_token_xxx \
  -H 'Accept: text/xml' \
  -H 'Content-Type: text/xml' \
  -d '<data><Title>New from xml #15:45</Title><Body>New super body</Body></data>'
