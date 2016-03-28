-- Problem 1

main :: IO ()
main = print (filterSum 999)

filterSum :: Int -> Int
filterSum n = sum [ x | x <- [1 .. n], x `mod` 3 == 0 || x `mod` 5 == 0]
