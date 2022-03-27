#!/usr/local/bin/node

const http = require('http')

const config = {
  port: process.env.PORT || 5000,
  db: {
    host: process.env.FIRMWARE_DB_HOST || 'localhost',
    port: process.env.FIRMWARE_DB_PORT || 27017,
    name: process.env.FIRMWARE_DB_NAME || 'firmware-service',
    username: process.env.FIRMWARE_DB_USERNAME || 'mongoadmin',
    password: process.env.FIRMWARE_DB_PASSWORD || 'secret'
  }
};

(
  async function main() {
    const app = await require('./server.js')(config)
    http.createServer(app).listen(config.port)    
  }
)()
