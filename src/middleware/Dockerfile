FROM node:14-alpine

WORKDIR /app

COPY package.json .
COPY package-lock.json .
RUN npm install

COPY server.js .

EXPOSE 9000

CMD [ "node", "server.js" ]