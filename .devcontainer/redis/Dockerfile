FROM redis:latest

LABEL maintainer="Mahmoud Zalt <mahmoud@zalt.me>"

## For security settings uncomment, make the dir, copy configs, and also start with the configs, to use it
#RUN mkdir -p /usr/local/etc/redis
#COPY redis.configs /usr/local/etc/redis/redis.configs

VOLUME /data

EXPOSE 6379

#CMD ["redis-server", "/usr/local/etc/redis/redis.configs"]
CMD ["redis-server"]
