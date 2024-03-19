package devctl.auth

import kotlinx.serialization.*

@Serializable
data class Error (
    val error: String,

    @SerialName("error_description")
    val errorDescription: String
)

class ResponseException(val data: Error) : Exception(data.errorDescription)

