package devctl.auth


import kotlinx.serialization.*
import java.time.Instant

@Serializable
data class Token(

    @SerialName("access_token")
    val accessToken: String,

    @SerialName("refresh_token")
    val refreshToken: String? = null,

    @SerialName("id_token")
    val idToken: String,

    @SerialName("token_type")
    val tokenType: String,

    @SerialName("expires_in")
    val expiresIn: Long,

    @SerialName("created_at")
    val createdAt: Long = Instant.now().epochSecond,

    val scope: String
)

fun Token.IsExpired(): Boolean {
    return Instant.now().epochSecond > expiresIn + createdAt
}

