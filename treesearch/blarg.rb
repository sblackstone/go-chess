def snoob(x)
  smallest = x & ~x
  ripple = x + smallest
  ones = x ^ ripple
  one = (ones >> 2 ) / smallest
  return ripple | ones
end



first = 1

while true
  x = first
  y = snoob(x)
  break if y > x
  puts y.to_s(2)
  x = y
end
