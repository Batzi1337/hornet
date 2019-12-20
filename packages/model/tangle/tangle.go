package tangle

import "github.com/gohornet/hornet/packages/database"

func ConfigureDatabases(directory string) {
	database.Settings(directory)
	configureHealthDatabase()
	configureTransactionStorage()
	configureBundleDatabase()
	configureTransactionHashesForAddressDatabase()
	configureApproversDatabase()
	configureMilestoneDatabase()
	configureLedgerDatabase()
	configureSnapshotDatabase()
	configureSpentAddressesDatabase()
	configureTransactionHashesForAddressDatabase()
	configureUnconfirmedTransactionsDatabase()
}

func LoadInitialValuesFromDatabase() {
	loadSnapshotInfo()
	loadSolidEntryPoints()
}
