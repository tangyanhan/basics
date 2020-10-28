/**
 * 管道，runoob上的代码压缩和解压缩无法一起执行，主要是因为stream的写入读取过程都是异步的，
 * 压缩和解压缩管道一起执行，就会造成解压缩尝试去解压一个还没写完的文件，从而报错。
 */
var fs = require('fs');
var zlib = require('zlib');

const fileName = 'setup.md.gz';

fs.createReadStream('SETUP.md').pipe(zlib.createGzip()).pipe(fs.createWriteStream(fileName).on('finish', function() {
    console.log('write finished');
    fs.createReadStream(fileName).
    pipe(zlib.createGunzip()).
    pipe(fs.createWriteStream('uncompressed.txt')).on('error', function(err) {
        console.log('uncompress failed:', err.stack);
    });
    console.log('File uncompressed');
}));

console.log('File compressed');


