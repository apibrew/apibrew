FROM node:20-alpine

ENV NODE_ENV production

# Create app directory
WORKDIR /usr/src/app

# Install app dependencies
COPY package.json .
COPY yarn.lock .
COPY tsconfig.json .

RUN yarn

COPY . .

RUN yarn build

USER node
EXPOSE 23619

CMD [ "node", "dist/index.js" ]