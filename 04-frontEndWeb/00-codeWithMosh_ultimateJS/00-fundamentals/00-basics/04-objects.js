/*
 * Objects are a reference type (non-primitive)
 *
 */

let person = {
	name: "Iano",
	age: 27
};

console.log(person);

// DOT NOTATION
person.age = 19;

console.log(person.age);

// BRACKET NOTATION
person["name"] = "Njuguna";

console.log(person.name);

console.log(person);