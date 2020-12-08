# For running elastic and kibana
## lkdjhas
### xcajkda
dasdad

> block quote

    podman run -d --pod=goelkpod  --name elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.9.3

csadas

    podman run -d --pod=goelkpod --name kibana -p 5601:5601 kibana:7.9.3