FROM node:14-alpine

WORKDIR /app

COPY package.json .
COPY package-lock.json .
RUN npm install

COPY views/ views/

COPY server.js .

EXPOSE 3000

CMD [ "node", "server.js" ]