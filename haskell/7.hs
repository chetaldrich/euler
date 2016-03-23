-- Problem 7

main :: IO ()
main = print (primes !! 10000)

-- using trial division
primes :: [Integer]
primes = 2 : [x | x <- [3..], isprime x]

isprime :: Integer -> Bool
isprime x = all (\p -> x `mod` p > 0) (factorsToTry x)
    where
        factorsToTry factors = takeWhile (\p -> p*p <= factors) primes
