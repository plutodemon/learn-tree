"use strict";
// ts语法
let str = 'hello world';
let num = 1;
function rand(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}
let flag;
// unknown 为类型安全的any
let un;
// 断言
str = un;
str = un;
// object
let userObj;
userObj = { name: 'a' };
