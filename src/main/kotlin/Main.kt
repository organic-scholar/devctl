package devctl

import com.github.ajalt.clikt.core.subcommands
import devctl.command.MainCommand
import devctl.command.SshCommand
import devctl.command.SshSaveCommand


suspend fun main(args: Array<String>) {
    init()
    val mainCommand = MainCommand()
    val sshCommand = SshCommand()
    mainCommand.subcommands(sshCommand)
    sshCommand.subcommands(SshSaveCommand())
    mainCommand.main(args)
}



