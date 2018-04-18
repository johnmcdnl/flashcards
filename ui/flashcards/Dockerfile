ARG node_version=latest

FROM node:${node_version}

WORKDIR /usr/src/flashcards

COPY package.json .
RUN yarn install

COPY . .

EXPOSE 3000

CMD [ "yarn", "start" ]
