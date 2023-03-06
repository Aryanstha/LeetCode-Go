func reverse(x int) int {
    sign := 1
    if x < 0 {
        sign = -1
        x *= -1
    }

    // reverse the digits
    reversed := 0
    for x > 0 {
        reversed = reversed*10 + x%10
        x /= 10
    }

    // apply the sign to the reversed number
    reversed *= sign

    // check for overflow
    if reversed < math.MinInt32 || reversed > math.MaxInt32 {
        return 0
    }

    return reversed
}
