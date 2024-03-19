package devctl.auth

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

@Serializable
data class DeviceCode(
    @SerialName("device_code")
    val deviceCode: String,

    @SerialName("user_code")
    val userCode: String,

    @SerialName("verification_uri")
    val verificationURI: String,

    @SerialName("expires_in")
    val expiresIn: Long,

    val interval: Long,

    @SerialName("verification_uri_complete")
    val verificationURIComplete: String
)