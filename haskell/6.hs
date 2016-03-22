-- Problem 6

main :: IO ()
main = print result

result :: Integer
result = sum [x * x | x <- [1..100]] - sum [1..100] ^ (2 :: Integer)
