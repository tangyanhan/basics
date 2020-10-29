/**
 * Stream 是一个抽象接口，主要有四种类型：
 * Readable: 可读操作
 * Writable: 可写操作
 * Duplex: 可读可写操作
 * Transform: 操作被写入数据，然后读出结果
 * 
 * 所有的Stream对象都是EventEmitter的实例， 常用的事件有：
 * data: 有数据可读时触发
 * end: 没有更多数据可读时触发
 * error: 接收和写入过程中发生错误时触发
 * finish: 所有数据已被写入底层系统时触发
 */

var fs = require('fs');
var data = '';

var readStream = fs.createReadStream('SETUP.md');

readStream.setEncoding('utf8');
readStream.on('data', function(chunk) {
data += chunk;
});

readStream.on('end', function() {
    console.log(data);
});

readStream.on('error', function(err) {
    console.log(err.stack);
});

console.log('DONE!');

var writeStream = fs.createWriteStream('dummy');
writeStream.write('hello world', 'utf8');

writeStream.end();

writeStream.on('finish', function() {
    console.log('Write finished');
});

writeStream.on('error', function(err) {
    console.log('Failed to write:', err.stack);
});

console.log('Write complete!');

// 管道流
var srcRead = fs.createReadStream('SETUP.md');
var dstWrite = fs.createWriteStream('SETUP.md.bak');

srcRead.pipe(dstWrite);

console.log('Pipe complete!');
