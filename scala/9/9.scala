/**
 * Solution to Problem 9 in Scala
 */
object Problem_9 {
  def main(args: Array[String]): Unit = {
    val solution = (for (b <- 2 until 1000;
                         a <- 1 until b; 
                         c = 1000 - a - b
                         if  a * a + b * b == c * c) yield a * b * c).head;
    println(solution)
  }
}
