package sts

type DatabaseAccount struct {
	MysqlUser, MysqlPassword, MysqlDatabase, MysqlRootPassword string
}

var default_dbaccount = DatabaseAccount{"testuser", "testpass", "testdb", "rootpass"}
