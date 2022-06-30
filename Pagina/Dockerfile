FROM node:16

ADD . /app
WORKDIR /app
COPY package.json .
COPY package-lock.json .

RUN npm install

COPY . .
#COPY ./ ./

RUN npm run build 

EXPOSE 3000

CMD ["npm", "start"]

#CMD ["npm", "run", "start"]
