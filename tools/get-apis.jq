.docs.resources |
map_values(.
  | .methods[]
  | select(.name == "show")|
  .apis |
  map_values(. | .api_url)
)
