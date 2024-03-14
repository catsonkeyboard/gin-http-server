FROM alpine:latest

ENV SECRET 【use your secret 】
ENV PORT 【use your port】

WORKDIR /apps

COPY go_build_main_go_linux /apps/app
COPY public.pem /apps/public.pem
COPY private.pem /apps/private.pem

EXPOSE 【use your port】
ENTRYPOINT ["./app"]
