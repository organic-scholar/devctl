package common

import kotlinx.coroutines.delay

suspend fun <T> retry(
    initialDelayMillis: Long = 100,
    maxDelayMillis: Long = 1000,
    factor: Double = 2.0,
    block: suspend () -> T
): T {
    var currentDelay = initialDelayMillis
    try {
        return block()
    } catch (e: Exception) {
        delay(currentDelay)
        currentDelay = (currentDelay * factor).toLong().coerceAtMost(maxDelayMillis)
    }
    throw IllegalStateException("Unreachable")
}

fun exponential(
    initial: Long,
    max: Long,
    factor: Int
): Sequence<Long> = sequence {
    var currentDelay = initial
    while (true) {
        yield(currentDelay)
        currentDelay = (currentDelay * factor).toLong()
        if (currentDelay > max) {
            currentDelay = initial
        }
    }
}
