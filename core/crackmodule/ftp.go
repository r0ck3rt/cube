package crackmodule

import (
	"cube/config"
)

type FtpCrack struct {
	*Crack
}

func (ftpCrack FtpCrack) SetName() string {
	return "ftp"
}

func (ftpCrack *FtpCrack) SetPort() string {
	return "21"
}

func (ftpCrack FtpCrack) SetAuthUser() []string {
	return []string{"anonymous", "ftp", "admin", "www", "web", "root", "db", "wwwroot", "data"}
}

func (ftpCrack FtpCrack) SetAuthPass() []string {
	return config.PASSWORDS
}

func (ftpCrack FtpCrack) IsLoad() bool {
	return true
}

func (ftpCrack FtpCrack) IsMutex() bool {
	return false
}

func (ftpCrack FtpCrack) IsTcp() bool {
	return true
}

func (ftpCrack FtpCrack) Exec() (result CrackResult) {
	result = CrackResult{Crack: *ftpCrack.Crack, Result: "", Err: nil}

	return result
}

func init() {
	AddCrackKeys("ftp")
}