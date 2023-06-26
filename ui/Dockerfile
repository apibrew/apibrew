# stage1 as builder
FROM node:20-alpine as builder

RUN npm install -g npm@latest

WORKDIR /ui

# copy the package.json to install dependencies
COPY package.json ./

# Install the dependencies and make the folder
RUN npm install

COPY . .

RUN npm install -w @apibrew/core-lib 
RUN npm install -w @apibrew/documentation
RUN npm install -w @apibrew/app-designer
RUN npm install -w @apibrew/core-ui
RUN npm install -w @apibrew/app
RUN npm install --save-dev typescript

# Build the project and copy the files
RUN npm run build -w @apibrew/core-lib
RUN npm run build -w @apibrew/documentation
RUN npm run build -w @apibrew/app-designer
RUN npm run build -w @apibrew/core-ui
RUN npm run build -w @apibrew/app

FROM nginx:alpine

#!/bin/sh

COPY ./nginx.conf /etc/nginx/nginx.conf

## Remove default nginx index page
RUN rm -rf /usr/share/nginx/html/*

# Copy from the stahg 1
COPY --from=builder /ui/app/dist /usr/share/nginx/html

EXPOSE 80

ENTRYPOINT ["nginx", "-g", "daemon off;"]