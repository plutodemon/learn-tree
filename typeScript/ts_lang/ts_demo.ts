// ts语法
let str: string = 'hello world';
let num: number = 1;

function rand(min: number, max: number): number {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

let flag: boolean | number;
// unknown 为类型安全的any
let un: unknown;
// 断言
str = un as string;
str = <string>un;

// object
let userObj: { name: string, age?: number };
userObj = { name: 'a' };