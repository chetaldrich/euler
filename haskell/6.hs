-- Problem 6

main :: IO ()
main = print result

result :: Integer
result = sum [1..100] ^ (2 :: Integer) - sum [x * x | x <- [1..100]]
