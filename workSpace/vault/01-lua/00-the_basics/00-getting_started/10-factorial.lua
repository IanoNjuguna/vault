-- defines a factorial function
function Fact(n)
  if n == 0 then
    return 1
  else
    return n * Fact(n - 1)
  end
end

print("Enter a Number:")
local nmbr = io.read("*n")

print(Fact( nmbr ))


