const prpl = require('prpl-server');
const express = require('express');
const app = express()
const port = 8080

app.get('/*', prpl.makeHandler('./build', require('./build/polymer.json')));

app.listen(port, () => console.log(`Server listening on  ${port}!`))