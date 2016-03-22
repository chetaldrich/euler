-- Problem 5

main :: IO ()
main = print result

result :: Integer
result = foldl lcm 1 [1..20]
