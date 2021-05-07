package main

import "testing"

func TestTransaction_whenTypical(t *testing.T) {
	expected := uint64(0)

	entries := []PendingTransaction{}
	entries = append(entries,
		PendingTransaction{
			"A",
			"B",
			1,
			2,
		})

	transactions := PendingTransactions{entries, uint64(len(entries))}
	actual := transactions.selectTransaction()

	if actual != expected {
		t.Errorf("No Match: %d != %d", actual, expected)
	}
}

func TestMultipleIndependentTransactions_whenTypical(t *testing.T) {
	expected := uint64(1)

	entries := []PendingTransaction{}
	entries = append(entries,
		PendingTransaction{
			"A",
			"B",
			1,
			0,
		})
	entries = append(entries,
		PendingTransaction{
			"C",
			"D",
			1,
			2,
		})

	transactions := PendingTransactions{entries, uint64(len(entries))}
	actual := transactions.selectTransaction()

	if actual != expected {
		t.Errorf("No Match: %d != %d", actual, expected)
	}
}
