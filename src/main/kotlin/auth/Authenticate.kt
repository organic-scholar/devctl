package devctl.auth

import common.exponential
import kotlinx.coroutines.delay


suspend fun authenticate(): Boolean {
    val deviceCode = requestDeviceCode()
    println(deviceCode)
    for (duration in exponential(500, 8000, 2)) {
        val token = try {
            requestAuthToken(deviceCode)
        } catch (exception: ResponseException) {
            println(exception.data)
            when(exception.data.error) {
                "authorization_pending", "slow_down" -> {
                    delay(5000)
                    continue
                }
                else ->  throw exception
            }
        }
        writeToken(token)
        break
    }
    return false
}

suspend fun readToken(){


}

