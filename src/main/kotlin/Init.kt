package devctl

import devctl.common.initConfig
import devctl.common.initS3Client

suspend fun init() {
    initConfig()
    initS3Client()
}