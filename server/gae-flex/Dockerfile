FROM node:8
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install --only=production
COPY . .
RUN ls -al
CMD [ "node", "server.js" ]
