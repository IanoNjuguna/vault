const circle = {
	radius: 1,

	draw: () => {
		console.log('draw');
	}
}

// for .. in loop

for (let key in circle)
	console.log(key, circle[key]);

// for .. of loop

for (let key of Object.keys(circle))
	// Returns a string array constaining all properties, keys and methods
	console.log(key);

for (let entry of Object.entries(circle))
	// Returns each key value pair as an array
	console.log(entry);

// conditional statement
if ('radius' in circle)
	console.log('yes');
else
	console.log('no');


