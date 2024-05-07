FROM node:latest

WORKDIR /app

COPY . .

RUN yarn install
RUN yarn prebuild
RUN yarn build

RUN rm -rf ./src

CMD ["yarn", "start"]