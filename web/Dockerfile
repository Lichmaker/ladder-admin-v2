FROM node:16

WORKDIR /la_web/
COPY . .

RUN npm config set registry https://registry.npm.taobao.org --global 
RUN npm install && npm run build

FROM nginx:alpine
LABEL MAINTAINER="lich.wu2014@gmail.com"

COPY .docker-compose/nginx/conf.d/my.conf /etc/nginx/conf.d/my.conf
COPY --from=0 /la_web/dist /usr/share/nginx/html
RUN cat /etc/nginx/nginx.conf
RUN cat /etc/nginx/conf.d/my.conf
RUN ls -al /usr/share/nginx/html
