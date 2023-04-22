FROM node:18-slim

WORKDIR /app

RUN mkdir -p /app/frontend
RUN mkdir -p /app/swagger

COPY ./frontend /app/frontend
COPY ./swagger /app/swagger

WORKDIR /app/frontend
RUN npm install
RUN npm run generate
RUN npm run build

CMD [ "npm", "start" ]
