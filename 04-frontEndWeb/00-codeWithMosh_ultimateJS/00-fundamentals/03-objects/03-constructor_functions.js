// declare objects using constructor functions syntax

class Circle {
	constructor(radius) {
		this.radius = radius;
		this.draw = function () {
			console.log("draw");
		};
	}
}

const circleOne = new Circle(1)

console.log(circleOne);
