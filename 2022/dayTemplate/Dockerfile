FROM golang:1.19-alpine as BASE
WORKDIR /src
COPY ./ .
RUN go build -buildvcs=false -o /bin/aoc

FROM scratch
COPY --from=BASE /bin/aoc /bin/aoc
COPY input.txt .
ENTRYPOINT [ "/bin/aoc" ]