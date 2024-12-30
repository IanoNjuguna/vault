// Primitive Types are passed by value,
// i.e: when you clone a primitive type the value is passed onto the new variable.

// Reference Types are passed by Reference,
// i.e: when you clone a reference type the value passed onto the new variable
// is the memory address of the value in the original variable.


// PRIMITIVE TYPES
let x = 10;
let y = x;

// Doesn't change the value of y
x = 20;

console.log(x);
console.log(y);

// REFERENCE TYPES
let new_obj = {
	z: 40
};

let reference_obj = new_obj;

// Will change value of z in both objects
new_obj.z = 75;

console.log(new_obj);
console.log(reference_obj);

