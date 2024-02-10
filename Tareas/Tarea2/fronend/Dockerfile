FROM node:20.5.0-alpine

WORKDIR /front

COPY package.json ./
COPY package-lock.json ./

RUN npm install

COPY . ./

EXPOSE 3000

CMD [ "npm", "start" ]