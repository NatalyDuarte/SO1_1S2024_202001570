FROM node:14.21-alpine

WORKDIR /fronten

COPY package.json ./
COPY package-lock.json ./

RUN npm install

COPY . ./

EXPOSE 3000

CMD [ "npm", "start" ]