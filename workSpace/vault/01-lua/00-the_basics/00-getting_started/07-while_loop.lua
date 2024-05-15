-- Use a different seed every time we start
math.randomseed(os.time())

-- Set positions of the hero
local player_x = 400
local player_y = 300
local num_enemies = 1

while num_enemies < 501 do
	-- Get random enemy position for a 800 * 600 resolution
	local enemy_x = math.random(0, 800)
	local enemy_y = math.random(0, 600)

    if player_x == enemy_x and player_y == enemy_y then
        -- TODO: Error, enemy cannot be in the same pos
		print("Enemy and player position clashed!")
	else
		-- Display the two random values
		print("Enemy "..num_enemies..": ("..enemy_x..","..enemy_y..")")

		num_enemies = num_enemies + 1
		end
end

print("Thank you, goodbye!")

