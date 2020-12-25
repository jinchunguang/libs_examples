package main

import "go_mysql/dbs"

func main() {

	dbs.StructInsert()
	dbs.StructUpdate()
	dbs.StructQueryField()
	dbs.StructQueryAllField()
	dbs.StructDel()
	dbs.StructTx()
	dbs.RawQueryField()
	dbs.RawQueryAllField()
}
