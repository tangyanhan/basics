/**
 * JS自身只有字符串数据类型，没有二进制数据类型。
 * Buffer 能够分配一块内存来存储和处理这些二进制数据, 也能够将其转换成不同编码输出
 */

var buf = Buffer.from('hellokitty', 'ascii');

console.log('Hex=', buf.toString('hex'));
console.log('Base64=', buf.toString('base64'));

// 创建

// 创建一个长度 10， 用0填充的Buffer（不指定用0填充）
const buf1 = Buffer.alloc(10);

// 创建一个长度10， 用0x1填充的Buffer
const buf2 = Buffer.alloc(10, 1);

// 创建一个长度为10， 未初始化的Buffer，速度快，但可能包含旧数据
// 需要使用fill()/write()重写
const buf3 = Buffer.allocUnsafe(10);

// 创建一个包含 [0x1, 0x2, 0x3] 的Buffer
const buf4 = Buffer.from([1, 2, 3]);

// 创建一个包含UTF-8 字节的Buffer
const buf5 = Buffer.from('tést');

// 创建一个包含Latin-1字节的Buffer
const buf6 = Buffer.from('tést', 'latin1');


// 写入

// 写入Buffer时，如果buffer空间不足，只会写入部分字符串
// buf.write(string[, offset[, length]][, encoding])
buf = Buffer.alloc(5);
len = buf.write('1234567890');
console.log('Written=', len, 'Content=', buf.toString('ascii'));
len = buf.write('12345', 2, 'ascii');
console.log('Written=', len, 'Content=', buf.toString('ascii'));

// 读取

console.log('Content from offset 2=', buf.toString('ascii', 2));
console.log('Content from 2-5', buf.toString('ascii', 2, 5));

// 将Buffer转换为JSON对象

console.log('JSON=\n', buf.toJSON());

// 缓冲区合并

var b1 = Buffer.from('hello');
var b2 = Buffer.from('kitty');
var b3 = Buffer.concat([b1, b2]);

console.log('Concat b3=', b3.toString());

// 缓冲区比较

var v1 = Buffer.from('ABC');
var v2 = Buffer.from('ABCD');

var result = v1.compare(v2);

// 比较字符串大小， <0 表明小， =0表示相等， >0 表示较大
console.log('Compare ', v1.toString(), v2.toString, 'Result=', result);

// 拷贝缓冲区
// buf.copy(targetBuffer[, targetStart[, sourceStart[, sourceEnd]]])
// 将v2拷贝到v1从2开始的位置
// 拷贝仍然不能突破原本大小
v2.copy(v1, 2);
console.log('After copy, v1=', v1.toString());

// 缓冲区切片
// buf.slice([start[, end]])
// 返回一个新的Buffer，与旧的缓冲区指向同一块内存，但开始和结束位置另外指定
// 实际返回的子集是 [start, end), 即不会包含end位置的数据
var sliceSource = Buffer.from('0123456789');

var s = sliceSource.slice(3, 7);
console.log('Slice =', s.toString(), 'Length=', s.length);
