package main

type PendingTransaction struct {
	from string
	to   string
	amount uint64
	age_seconds uint64
}

type PendingTransactions struct {
	entries  []PendingTransaction
	capacity uint64
}

func (transactions* PendingTransactions)selectTransaction() uint64{
	max_stake_index:= uint64(0)
	max_stake := uint64(0)
	for i := uint64(0); i < transactions.capacity; i++ {
		// age prevents big staked amounts from hogging the network
		stake := transactions.entries[i].amount * transactions.entries[i].age_seconds
		if max_stake < stake{
			max_stake = stake
			max_stake_index = i
		}
	}
	return max_stake_index
}