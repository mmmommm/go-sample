package database

import (
	"github.com/google/wire"
	"github.com/mmmommm/go-sample/layered-architecture/database/mysql"
)

var Set = wire.NewSet(
	mysql.ProvideMysqlClient,
)