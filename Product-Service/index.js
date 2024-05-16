// index.js
const express = require('express');
const app = express();
const port = 3000;
const winston = require('winston');
const fs = require('fs');

const logDirectory = '/var/log';
if (!fs.existsSync(logDirectory)) {
    fs.mkdirSync(logDirectory);
}

const logger = winston.createLogger({
    level: 'info',
    format: winston.format.json(),
    transports: [
        new winston.transports.File({ filename: '/var/log/product.log' })
    ]
});

app.get('/', (req, res) => {
    logger.info('Hello from Product-Service get request');
    res.send('Hello from Product-Service');
});

app.listen(port, () => {
    logger.info('Executing Product-Service');
    console.log(`Product-Service listening at http://localhost:${port}`);
});
