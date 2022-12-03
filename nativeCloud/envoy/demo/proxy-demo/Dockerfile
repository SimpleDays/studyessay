FROM envoyproxy/envoy-alpine:v1.9.1

ADD baidu-proxy-envoy.yaml /etc

CMD /usr/local/bin/envoy -c /etc/baidu-proxy-envoy.yaml