package devctl.auth

import devctl.DEVCTL_DIR
import devctl.TOKEN_FILE
import kotlinx.serialization.json.Json
import kotlinx.serialization.encodeToString
import kotlin.io.path.Path
import kotlin.io.path.createDirectories
import java.nio.file.Paths
import kotlin.io.path.writeText


fun writeToken(token: Token) {
    val json = Json.encodeToString(token)
    Paths.get(DEVCTL_DIR, "token.json").writeText(json)
    Paths.get(DEVCTL_DIR, "token").writeText(token.idToken)
}