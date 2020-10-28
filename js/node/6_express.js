/**
 * Express 样例， express是一个web框架
 * 需要先运行 cnpm install express --save
*/

var express = require('express');
var app = express();

app.get('/', function(req, res) {
    res.send('<html><body><h2>Hello World</h2></body></html>');
});

var server = app.listen(8081, function() {
    var host = server.address().address
    var port = server.address().port
    console.log('Listening on http://%s:%s', host, port);
});
