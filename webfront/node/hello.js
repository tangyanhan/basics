/**
 * NodeJS 一个文件就是一个模块
*/

function Hello() {
    var name;
    this.getName = function() {
        return this.name;
    };
    this.setName = function(name) {
        this.name = name;
    };
}

module.exports = Hello;
