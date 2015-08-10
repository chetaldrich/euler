/**
 * A brute force solution to Problem 1 of Project Euler.
 * author: Chet Aldrich
 */
object Problem_1 {
  def main(args: Array[String]) {
    val solution = (1 until 1000)
      .view
      .filter(x => x % 3 == 0 || x % 5 == 0)
      .sum

    println(solution)
  }
}
