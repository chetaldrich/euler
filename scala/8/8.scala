/**
 * Solution for Euler Problem 8
 */
import scala.io._

object Problem_8 {
  def main(args: Array[String]): Unit = {
    val number = Source.fromFile("number.txt")
    val numberString = number
      .getLines
      .foldLeft("")((b, a) => b + a)
    val solution = numberString
      .view
      .map(_.asDigit)
      .map(_.toLong)
      .sliding(13)
      .map(_.product)
      .max
    println(solution)
  }
}
