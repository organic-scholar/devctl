package devctl.common

import com.sksamuel.hoplite.ConfigLoaderBuilder
import com.sksamuel.hoplite.addResourceSource

data class Config(
    val s3Endpoint: String?,
    val authUrl: String,
    val authClientId: String
)

lateinit var config: Config

fun initConfig(){
    val buildType = System.getProperty("devctl.build.type") ?: "release"
    val builder = ConfigLoaderBuilder.default().addResourceSource("/application.yml")
    if(buildType == "debug") {
        builder.addResourceSource("/application.debug.yaml")
    }
    config = builder.build().loadConfigOrThrow<Config>()
}
