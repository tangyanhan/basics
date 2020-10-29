// NodeJS 可以通过回调(callback)来实现异步编程，回调函数可以通过异步形式调用避免阻塞
var fs = require('fs');

// 阻塞方式调用
var data = fs.readFileSync('SETUP.md');
console.log("File content:\n", data.toString());

// 非阻塞方式调用
fs.readFile('SETUP.md', function(err, data) {
    if (err) return console.error(err);
    console.log('File content async:\n', data.toString());
});
