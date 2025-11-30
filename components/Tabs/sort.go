package tabs

import (
	"cmp"
	"slices"

	list "github.com/Keivan-sf/Bushuray-tui/components/List"
)

func (m *Model) sort_by_test_results() {
	active := m.Children[m.ActiveTap].Content
	primary_item_id := -1
	if active.Primary != -1 {
		primary_item_id = active.Items[active.Primary].ProfileId
	}
	sortFn := func(a, b list.ListItem) int {
		if a.TestResult > 0 && b.TestResult > 0 {
			return cmp.Compare(a.TestResult, b.TestResult)
		} else if (a.TestResult <= 0 && b.TestResult > 0) || a.TestResult == -1 {
			return 1
		} else {
			// a is either -2 or more than b
			return -1
		}
	}
	slices.SortFunc(active.Items, sortFn)

	if primary_item_id != -1 {
		for i, item := range active.Items {
			if item.ProfileId == primary_item_id {
				active.Primary = i
				break
			}
		}
	}

	m.Children[m.ActiveTap].Content = active

}
