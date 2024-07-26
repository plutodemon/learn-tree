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
// 任意属性
let userObj2;
userObj2 = { name: 'a', sex: 'man' };
let userObj3;
userObj3 = { name: 'a', age: 1 };
// 函数类型声明
let fn;
fn = (a, b) => a + b;
// 数组 元组 枚举
let arr = [1, 2, 3];
let arr2 = [1, 2, 3];
let arr3 = ['a', 1];
var Color;
(function (Color) {
    Color[Color["Red"] = 0] = "Red";
    Color[Color["Green"] = 1] = "Green";
    Color[Color["Blue"] = 2] = "Blue";
})(Color || (Color = {}));
let t;
t = 1;
