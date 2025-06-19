package anglerfish

import "testing"

func TestWeightedRandomPick(t *testing.T) {
	tests := []struct {
		name    string
		items   []WeightedItem[string]
		want    string
		wantErr bool
	}{
		{
			name:    "empty list",
			items:   []WeightedItem[string]{},
			want:    "",
			wantErr: true,
		},
		{
			name: "single item",
			items: []WeightedItem[string]{
				{Item: "item1", Weight: 10},
			},
			want:    "item1",
			wantErr: false,
		},
		{
			name: "multiple items with weights",
			items: []WeightedItem[string]{
				{Item: "item1", Weight: 1},
				{Item: "item2", Weight: 2},
				{Item: "item3", Weight: 3},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WeightedRandomPick(tt.items)
			if (err != nil) != tt.wantErr {
				t.Errorf("WeightedRandomPick() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == "" {
				t.Errorf("WeightedRandomPick() got = %v, want non-empty string", got)
			}
		})
	}
}
