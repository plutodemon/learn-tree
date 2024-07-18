"use strict";
function greeter(person) {
    if (person === "Jane User") {
        return "Hello, Jane User";
    }
    return "Hello, " + person;
}
let user = "Jane";
console.log(greeter(user));
