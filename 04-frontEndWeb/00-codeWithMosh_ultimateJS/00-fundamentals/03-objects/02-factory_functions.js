// Declare objects using factory functions syntax
function createCircle(radius) {
	return {
		radius,

		draw: function () {
			console.log("draw");
		}
	};
}

circleOne = createCircle(1);
circleTwo = createCircle(2);

console.log("Circle One: ", circleOne, "Circle Two: ", circleTwo);

