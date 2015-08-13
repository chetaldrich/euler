/**
 * Solution to Euler 7
 */
object Problem_7 {
  def main(args: Array[String]): Unit = {
    lazy val primes: Stream[Int] = 2 #:: Stream.from(3).filter(i =>
      primes.takeWhile{j => j * j <= i}.forall{ k => i % k > 0});
    val solution = primes(10000)
    println(solution)
    }
}
