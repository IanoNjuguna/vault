const circle = {
	radius: 2,

	draw: () => {
		console.log(draw)
	}
};

// Can clone multiple objects onto a single object
const cloneOne = Object.assign({}, circle);

// Concise and readable
const cloneTwo = { ...circle };

console.log(cloneOne);
console.log(cloneTwo);

