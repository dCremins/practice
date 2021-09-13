FROM node:12-alpine

RUN mkdir -p /usr/src/app

WORKDIR /usr/src/app

COPY /src ./src
COPY package*.json ./
COPY babel.config.json ./

RUN npm install
RUN npx babel ./src --out-dir ./dist --copy-files

EXPOSE 3000

CMD ["npm", "run", "start"]