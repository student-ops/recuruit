FROM nginx
COPY ./app/html/. /usr/share/nginx/html
COPY ./nginx/nginx.conf  /etc/nginx/nginx.conf