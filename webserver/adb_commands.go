package webserver

const PUSH string = "push"
const INSTALL string = "install"
const UNINSTALL string = "uninstall"
const SHELL string = "shell"
const PUSH_FILE_ARG string = "FILE /data/app/"
const INSTALL_FILE_ARG string = "FILE"

type AdbCommandInfo struct {
    Command string
    Arguments string
    FilePath string
}

type AdbResponse struct {
    Data CommandResult `json:"data"`  
    Version float64 `json:"version"`
}

type CommandResult struct {
    Message string `json:"message"`
    Version string `json:"version"`
    Output string  `json:"adb_output"`
}

func BuildAdbPushCommand(filePath string) AdbCommandInfo {
    return AdbCommandInfo{PUSH, PUSH_FILE_ARG, filePath}
}

func BuildAdbInstallCommand(apkPath string) AdbCommandInfo {
    return AdbCommandInfo{INSTALL, INSTALL_FILE_ARG, apkPath}
}

func BuildAdbUninstallCommand(packageName string) AdbCommandInfo {
    return AdbCommandInfo{UNINSTALL, packageName, ""}
}

func BuildAdbShellCommand(command string) AdbCommandInfo {
    return AdbCommandInfo{SHELL, command, ""}
}