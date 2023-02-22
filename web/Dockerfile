FROM node:gallium

WORKDIR /app

COPY . ./

COPY package*.json ./

RUN npm install --omit=dev

CMD [ "npm" "run" "dev"]