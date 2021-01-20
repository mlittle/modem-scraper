#FROM alpine:3.10.2
FROM ubuntu:20.04

COPY modem-scraper ./modem-scraper

VOLUME [ "/config" ]

COPY ./scraper_config.yaml /config/config.yaml

ARG name=modem-scraper
ARG network=influx_network

#ENTRYPOINT [ "./modem-scraper", "-config", "/config/config.yaml" ]
#ENTRYPOINT [ "ls", "-ltra", "/config/"]
ENTRYPOINT [ "./modem-scraper", "-config", "/config/config.yaml" ]

