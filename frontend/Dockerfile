FROM node:slim as build-stage
WORKDIR /app

ARG API_URL
ENV VITE_API_URL=$API_URL

COPY package*.json ./
RUN npm install
COPY ./ .
RUN npm run build

FROM nginx as production-stage
WORKDIR /app

COPY --from=build-stage /app/dist /app
COPY nginx.conf /etc/nginx/nginx.conf
