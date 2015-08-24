/**
 * Solution to Problem 10 in Scala.
 */
object Problem_10 {
  def main(args: Array[String]): Unit = {
    lazy val primes: Stream[Int] = 2 #:: Stream.from(3).filter(i =>
      primes.takeWhile{j => j * j <= i}.forall{ k => i % k > 0});

    val solution : Long = primes.view.takeWhile(_ < 2000000).foldLeft(0L)(_ + _)
    println(solution)
  }
}
