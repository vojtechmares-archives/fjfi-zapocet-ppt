FROM node:12 as builder

WORKDIR /srv/app

RUN ["npm", "ci"]
RUN ["npm", "run", "build"]


FROM nginx:alpine

RUN apk add --no-cache ca-certificates

WORKDIR /var/www/html

COPY --from=builder /src/app/dist/* .

