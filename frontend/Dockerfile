FROM node:17 as build
WORKDIR /build
COPY package.json yarn.lock ./
RUN yarn
COPY . .
RUN yarn run static

FROM nginx
COPY --from=build /build/out /usr/share/nginx/html
