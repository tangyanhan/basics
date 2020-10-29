// NodeJS 拥有自己的事件库，事件可以通过字符串定义，发出和接收处理
var events = require('events');

var eventEmitter = new events.EventEmitter();

eventEmitter.on('start', function() {
    console.log('Received event - start');
    eventEmitter.emit('run');
});

eventEmitter.on('run', function() {
    console.log('Received event - run');
});

eventEmitter.emit('start');