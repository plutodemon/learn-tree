let num = 10;

const pi = 3.1415926;

let user = {
    name: "Jane User",
    age: 25,
    isMan: false
}

console.log(num++);

// typeof 判断数据类型
//alert(typeof typeof user);

if (typeof user === "object") {
    document.write("user: " + user.name + "<br>");
} else {
    document.write("user: " + user + "<br>");
}

switch (user.age) {
    case 25:
        document.write("user age: 25<br>");
        break;
    case 26:
        document.write("user age: 26<br>");
        break;
    default:
        document.write("user age: 999<br>");
}

// 三元运算符
user.isMan ?
    document.write("user men: " + user.isMan + "<br>") :
    document.write("user women: " + user.isMan + "<br>");

for (let i = 0; i < 5; i++) {
    console.log(i);
}

while (num < 15) {
    console.log(num++);
}

// 字符串

let str = "Hello World!";
let str1 = "aaa";
let str2 = "bbb";
console.log(str.length, str.charAt(str.length - 1),
    str1.concat(str2), str.substring(5),
    str.indexOf("o"), str.lastIndexOf("o"),
    str.trim(), str.trimStart());

let arr = [1, 2, 3, 4, 5];
for (let num in arr) {
    console.log(num);
}
for (let num of arr) {
    console.log(num);
}

// 函数
function add(a, b) {
    return a + b;
}

document.write("add: " + add("a", "b") + "<br>");

function rand(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

console.log(new Date());

// 是整个页面的根节点
console.log(document)
// 节点类型 9
console.log(document.nodeType)

// 读取元素
let divTag = document.getElementsByTagName("div")[0];
divTag.innerHTML = "Hello World!";

let divRoot = document.getElementById("root");
divRoot.innerHTML = "Hello Hello Hello World!";

let nav = document.querySelector(".nav");
let navs = document.querySelectorAll(".nav");

// 创建元素
let newDiv = document.createElement("div");
divRoot.appendChild(newDiv);

// 事件(其中还包含鼠标事件等等)
let buttonTag = document.getElementById("but");
buttonTag.onclick = function () {
    //alert("Hello World!");
};
buttonTag.addEventListener("click", function () {
    //alert("Hello World!111");
});
buttonTag.addEventListener("click", function () {
    //alert("Hello World!222");
});

// 对象解构赋值
const {name, age, isMan} = user;
const {log} = console;
log(name, age, isMan);
const {abs, random} = Math;
log(abs(-1), random());

// 字符串模板
let str3 = `Hello ${name}!`;
let addStr = "https://www.baidu.com"
let link = `<a href="${addStr}">百度</a>`;

log(str3.includes("Jane"), str3.startsWith("Hello"), str3.endsWith("!"),
    str3.repeat(3), str3.padStart(20, "a"), str3.padEnd(20, "a"), str3.at(-11));

// 数组扩展运算符
let arr1 = [1, 2, 3];
let arr2 = [4, 5, 6];
let arr3 = [...arr1, ...arr2];
// Array.from() 将类数组对象转换为数组对象 Array.of() 将一组值转换为数组对象

// 函数扩展 箭头函数
let add1 = (a, b) => a + b;
let add2 = (a, b) => {
    return a + b;
};

// 集合
let str4 = "abbaabdfx"
let set1 = new Set(str4);
console.log([...set1].join(""));
set1.add("x");
set1.has("x");
set1.clear()

// 异步
async function loadImageAsync(url) {
    return new Promise((resolve, reject) => {
        const image = new Image();
        image.src = url;
        image.onload = () => {
            resolve(image);
        };
        image.onerror = () => {
            reject(new Error("Could not load image at "));
        };
    });
}

const url = "https://www.bing.com/images/search?view=detailV2&ccid=VzhOTC3S&id=E9DB385D5FEDB9095FC5FCAB1BED3E3EBD7E2FBB&thid=OIP.VzhOTC3SVqdVV48AhF5grwHaFS&mediaurl=https%3a%2f%2fth.bing.com%2fth%2fid%2fR.57384e4c2dd256a755578f00845e60af%3frik%3duy9%252bvT4%252b7Rur%252fA%26riu%3dhttp%253a%252f%252fimg06file.tooopen.com%252fimages%252f20171224%252ftooopen_sy_231021357463.jpg%26ehk%3dwhpCWn%252byPBvtGi1%252boY1sEBq%252frEUaP6w2N5bnBQsLWdo%253d%26risl%3d%26pid%3dImgRaw%26r%3d0&exph=732&expw=1024&q=%e5%9b%be%e7%89%87&simid=608052299156357946&FORM=IRPRST&ck=89C5C3561E8CC288E7F093201B6D62C5&selectedIndex=0&itb=0"
const promise = loadImageAsync(url);
promise.then(
    function (data) {
        divRoot.appendChild(data);
    },
    error => {
        divRoot.innerHTML = "失败了！！！！";
    }
);

// 类
class Animal {
    constructor(name) {
        this.name = name;
    }

    static say() {
        return "Hello World!";
    }

    say() {
        return "My name is " + this.name;
    }
}

const animal = new Animal("Tom");
console.log(Animal.say());
console.log(animal.say());