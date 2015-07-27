/**
 * A solution to Euler problem 2 using streams.
 */
object Problem_2 {
    def main(args: Array[String]) {
        lazy val fibs: Stream[Int] = 0 #:: fibs.scanLeft(1)(_ + _)
        lazy val evenFibs = fibs filter (_ % 2 == 0)
        println(evenFibs.takeWhile(_ < 4000000).sum)
    }
}
