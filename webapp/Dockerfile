FROM node:16.17.0-alpine AS build
WORKDIR /app
COPY package.json yarn.lock ./
COPY tsconfig.json ./
COPY vue.config.js ./
RUN npm install
COPY . .
RUN npm run build

FROM nginx:1.21.3-alpine AS prod-stage
COPY --from=build /app/dist /usr/share/nginx/html
EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]