package devctl.auth

import devctl.common.config
import httpClient
import io.ktor.client.call.*
import io.ktor.client.request.forms.*
import io.ktor.http.*

suspend fun requestDeviceCode(): DeviceCode {
    val parameters = Parameters.build {
        append("client_id", config.authClientId)
        append("scope", "openid")
    }
    val response = httpClient.submitForm("${config.authUrl}/device/code", parameters)
    return response.body<DeviceCode>()
}