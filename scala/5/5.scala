/**
 * A solution to Euler problem 5.
 */
object Problem_5 {
  def gcd(a: Long, b: Long): Long = (a, b) match {
    case (a, b) if b == a => a
    case (a, b) if b > a  => gcd(b, a)
    case (a, 0)           => a
    case _                => gcd(b, a % b)
  }

  def lcm(a: Long, b: Long): Long = ((a * b) / gcd(a, b))

  def main(args: Array[String]) {
    val solution = (3 to 20).foldLeft(lcm(1, 2))((b, a) => lcm(a, b))
    println(solution)
  }
}
