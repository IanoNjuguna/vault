-- Use a different seed every time we start
math.randomseed(os.time())

-- FOR LOOP
for count = 1, 500 do
	-- Get random enemy position for a 800 * 600 resolution
	local enemy_x = math.random(0, 800)
	local enemy_y = math.random(0, 600)

	-- Display the two random values
	print("Enemy ("..count.."): ("..enemy_x..","..enemy_y..")")
end

print("Thank you, goodbye!")

