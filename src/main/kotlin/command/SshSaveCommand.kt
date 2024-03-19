package devctl.command

import com.github.ajalt.clikt.core.CliktCommand
import com.github.ajalt.clikt.core.PrintMessage
import java.io.ByteArrayOutputStream
import java.nio.file.Files
import java.nio.file.Paths
import java.util.zip.ZipEntry
import java.util.zip.ZipOutputStream
import kotlin.io.path.*
import aws.sdk.kotlin.services.s3.model.PutObjectRequest
import devctl.auth.authenticate
import devctl.common.s3Client
import kotlinx.coroutines.runBlocking

class SshSaveCommand : CliktCommand(name ="save", help = "Save ssh config") {

    override fun run(): Unit = runBlocking {
        authenticate()

    }
    @OptIn(ExperimentalPathApi::class)
    fun _run() {
        val sshPath = Paths.get(System.getenv("HOME"), ".ssh")
        if (!sshPath.exists()) {
            throw PrintMessage("${sshPath.pathString} doesn't exists", 1)
        }
        if (!sshPath.isDirectory()) {
            throw PrintMessage("${sshPath.pathString} is not a directory", 1)
        }
        val byteArrayOutputStream = ByteArrayOutputStream()
        val zipOutputStream = ZipOutputStream(byteArrayOutputStream)
        sshPath.walk(PathWalkOption.INCLUDE_DIRECTORIES).forEach {
            if (it.isRegularFile()) {
                val zipEntry = ZipEntry(it.relativeTo(sshPath).pathString)
                zipOutputStream.putNextEntry(zipEntry)
                Files.copy(it, zipOutputStream)
                zipOutputStream.closeEntry()
            }
        }
        zipOutputStream.close()
    }
    suspend fun uploadToS3() {
        s3Client.putObject(PutObjectRequest {

        })
    }
}