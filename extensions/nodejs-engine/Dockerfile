FROM node

ENV NODE_ENV production

# Create app directory
WORKDIR /usr/src/app

# Install app dependencies
COPY package.json ./

RUN npm install --production --frozen-lockfile

COPY . .
COPY src/proto ./dist/proto

RUN npm run build

USER node
EXPOSE 23619

CMD [ "node", "dist/index.js" ]