// index.js
const express = require('express');
const app = express();
const port = 3000;

app.get('/', (req, res) => {
    res.send('Hello from Product-Service');
});

app.listen(port, () => {
    console.log(`Product-Service listening at http://localhost:${port}`);
});
