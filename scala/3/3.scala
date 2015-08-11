/**
 * A solution to Euler Problem 3
 */
object Problem_3 {
  def main(args: Array[String]) = {
    var limit = 600851475143L
    lazy val naturals: Stream[Int] = Stream.cons(1, naturals.map(_ + 1))
    val primeFactors =
      naturals.drop(1).dropWhile(
        n => {while(limit % n == 0) {limit /= n}; limit > 1}
      );
    println(primeFactors.head)
  }
}
