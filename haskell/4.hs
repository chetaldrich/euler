-- Problem 2
-- My current approach because I don't know how to Haskell
-- that well yet.

main :: IO ()
main = print result

isPalindrome :: String -> Bool
isPalindrome str = str == reverse str

result :: Integer
result = ( maximum
         . map read
         . filter isPalindrome
         . map show
         . (concat :: [[Integer]] -> [Integer])
         . map (\x -> map (x *) [100..1000])
         ) [100..1000]
