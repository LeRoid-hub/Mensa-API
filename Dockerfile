FROM node:latest as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm i
COPY . .

FROM build-stage as production

ENV NODE_PATH=./build

RUN npm run build
