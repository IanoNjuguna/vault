math.randomseed(os.time())

-- Player Position (Middle of Screen)
local player = { x = 400, y = 300 }

local user_option = 0

-- Loop while user option is not 4
repeat

	-- Display a menu on the console
	print("+-------------------------------------+")
	print("  Welcome, ".. os.date())
	print("+-------------------------------------+")
	print("| 1. Generate random enemy position   |")
	print("| 2. Distance from enemy to player    |")
	print("| 3. Get angle from enemy to player   |")
	print("| 4. Exit                             |")
	print("+-------------------------------------+")

	-- Read the user option from console
	print("Please, select your option")
    local user_option = io.read("*n")

	-- Generate Enemy Position
	local enemy = { x = math.random(0, 800), y = math.random(0, 600) }

    -- Distance from enemy to player
	local distance_x = (enemy.x - player.x) * (enemy.x - player.x)
    local distance_y = (enemy.y - player.y) * (enemy.y - player.y)

    if user_option == 1 then
        -- Generate random enemy position
        print("New Enemy pos: (" .. enemy.x .. "," .. enemy.y .. ")")
    end

	if user_option == 2 then
        -- Get distance from enemy to player
        local distance = math.sqrt(distance_x + distance_y)

		print("Distance from enemy to player: ("..distance..")")
	end

    if user_option == 3 then
        -- Get angle from enemy to player
        local angle = math.deg(math.atan(distance_y, distance_x))
        print("Angle Between Enemy and Player: (" .. angle .. ")")
    end

until user_option == 4 or user_option >= 5 or user_option < 0

print("Goodbye")

