/**
 * A solution to Euler Problem 4
 */
object Problem_4 {

    def isPalindrome(drome: String): Boolean = {
        if (drome.reverse == drome) true else false
    }

    def main(args: Array[String]) = {
        val solution = (100 until 1000)
          .view
          .flatMap(i => (100 until 1000).map(i * _))
          .filter(n => isPalindrome(n.toString))
          .max

        println(solution)
    }
}
