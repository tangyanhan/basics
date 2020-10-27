// 创建服务器
// 通过 node fileName 运行
var http = require('http');

http.createServer(function (request, response) {
    response.writeHead(200, {'Content-Type': 'text/plain'});
    response.end('Hello World\n');
}).listen(8888);

console.log('Server is running at http://127.0.0.1:8888/');
