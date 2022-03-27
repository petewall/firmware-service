FROM node:17.8-alpine
WORKDIR /app

ENV NODE_ENV=production
ENV PORT=5000
ENV FIRMWARE_DB_HOST=
ENV FIRMWARE_DB_PORT=
ENV FIRMWARE_DB_NAME=
ENV FIRMWARE_DB_USERNAME=
ENV FIRMWARE_DB_PASSWORD=

COPY ["*.js", "package.json", "yarn.lock", "./"]

RUN yarn install --production

CMD [ "node", "index.js" ]
