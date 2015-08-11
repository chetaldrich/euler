/**
 * A solution to Euler problem 6.
 */
object Problem_6 {
    def main(args: Array[String]) = {
      val solution = math.pow((1 to 100).sum, 2) - (1 to 100).map(i => i * i).sum
      println(solution)
    }
}
