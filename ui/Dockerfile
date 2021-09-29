FROM node:14-alpine

WORKDIR /ui

# COPY rollup.config.js ./
COPY package*.json ./

RUN npm install

COPY ./src ./src
COPY ./ ./
# COPY ./public ./public

RUN npm run build

EXPOSE 3000

ENV HOST=0.0.0.0

CMD [ "npm", "run", "dev" ]