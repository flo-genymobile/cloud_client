package webserver

const PUSH string = "push"
const INSTALL string = "install"
const PUSH_FILE_ARG string = "FILE /data/app/"
const INSTALL_FILE_ARG string = "FILE"

type AdbCommandInfo struct {
    Command string
    Arguments string
    FilePath string
}

func BuildAdbPushCommand(filePath string) AdbCommandInfo {
    return AdbCommandInfo{PUSH, PUSH_FILE_ARG, filePath}
}

func BuildAdbInstallCommand(apkPath string) AdbCommandInfo {
    return AdbCommandInfo{INSTALL, INSTALL_FILE_ARG, apkPath}
}
