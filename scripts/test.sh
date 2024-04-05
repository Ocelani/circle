#!/bin/bash
#
# This script is used to manual test the API
base_url="http://localhost:3000"
endpoint_route="/tb01"

# Post tb01
echo "Post tb01"
curl -i -X POST "${base_url}${endpoint_route}" -d '{"col_texto":"test"}'
