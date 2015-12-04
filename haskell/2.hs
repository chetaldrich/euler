-- Problem 2

fibs :: [Integer]
fibs = 0 : 1 : zipWith (+) fibs (tail fibs)

main :: IO ()
main = print $ sum [x | x <- takeWhile (<= 4000000) fibs, even x]
