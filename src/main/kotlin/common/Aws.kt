package devctl.common

import aws.sdk.kotlin.runtime.auth.credentials.StsWebIdentityCredentialsProvider
import aws.sdk.kotlin.services.s3.S3Client
import aws.smithy.kotlin.runtime.net.url.Url

lateinit var s3Client: S3Client;

suspend fun initS3Client() {

//    StsWebIdentityCredentialsProvider.fromEnvironment(webIdentityTokenFilePath = "")

//    val url = config.s3Endpoint?.let { Url.parse(it) }
    s3Client = S3Client.fromEnvironment {
//        endpointUrl = url
    }
}