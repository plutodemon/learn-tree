let num = 10;

const pi = 3.1415926;

let user = {
    name: "Jane User",
    age: 25,
    isMan: false
}

console.log(num++);

// typeof 判断数据类型
alert(typeof typeof user);

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
    alert("Hello World!");
};
buttonTag.addEventListener("click", function () {
    alert("Hello World!111");
});
buttonTag.addEventListener("click", function () {
    alert("Hello World!222");
});