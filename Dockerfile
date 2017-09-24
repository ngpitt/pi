FROM scratch
COPY pi /
EXPOSE 80
ENTRYPOINT ["/pi"]
