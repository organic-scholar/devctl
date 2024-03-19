package devctl

import java.nio.file.Paths
import kotlin.io.path.Path


val DEVCTL_DIR = Paths.get(System.getenv("HOME"), ".devctl").toString()
val TOKEN_FILE = Paths.get(DEVCTL_DIR, "token")

