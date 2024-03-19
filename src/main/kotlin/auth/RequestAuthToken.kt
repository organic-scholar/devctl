package devctl.auth

import devctl.common.config
import httpClient
import io.ktor.client.call.*
import io.ktor.client.request.forms.*
import io.ktor.client.statement.*
import io.ktor.http.*

suspend fun requestAuthToken(deviceCodeResponse: DeviceCode): Token {
    val parameters = Parameters.build {
        append("client_id", config.authClientId)
        append("grant_type", "urn:ietf:params:oauth:grant-type:device_code")
        append("device_code", deviceCodeResponse.deviceCode)
    }
    val response = httpClient.submitForm("${config.authUrl}/token", parameters)
    if (response.status.isSuccess()) {
        println(response.bodyAsText())
        return response.body<Token>()
    }
    throw ResponseException(response.body<Error>())
}