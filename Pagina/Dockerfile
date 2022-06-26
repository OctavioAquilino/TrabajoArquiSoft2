FROM node:16

ADD . /Pagina
WORKDIR /Pagina
COPY package.json ./
COPY package-lock.json ./

RUN npm install

COPY ./ ./

#RUN npm run build 

EXPOSE 3000

#CMD ["npm", "start"]

CMD ["npm", "run", "start"]