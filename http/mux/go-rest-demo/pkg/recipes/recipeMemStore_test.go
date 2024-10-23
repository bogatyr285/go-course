package recipes

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func getHamCheeseToasties() Recipe {
	return Recipe{
		Name: "ham and cheese toastie",
		Ingredients: []Ingredient{
			{Name: "bread"},
			{Name: "ham"},
			{Name: "cheese"},
		},
	}
}

func TestMemStore_Add(t *testing.T) {
	type fields struct {
		list map[string]Recipe
	}
	type args struct {
		name   string
		recipe Recipe
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		wantLen int
	}{
		{
			name: "Add to empty map",
			fields: fields{
				map[string]Recipe{},
			},
			args: args{
				name:   "ham and cheese toastie",
				recipe: getHamCheeseToasties(),
			},
			wantLen: 1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MemStore{
				list: tt.fields.list,
			}
			err := m.Add(tt.args.name, tt.args.recipe)
			if !tt.wantErr {
				assert.NoError(t, err)
			}

			assert.Len(t, tt.fields.list, tt.wantLen)
		})
	}
}

func TestMemStore_Get(t *testing.T) {
	type fields struct {
		list map[string]Recipe
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Recipe
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Find Ham and cheese toasties",
			fields: fields{
				map[string]Recipe{
					"Ham and cheese toasties": getHamCheeseToasties(),
				},
			},
			args: args{
				name: "Ham and cheese toasties",
			},
			want:    getHamCheeseToasties(),
			wantErr: nil,
		},
		{
			name: "Not Found Ratatouille",
			fields: fields{
				map[string]Recipe{
					"Ham and cheese toasties": getHamCheeseToasties(),
				},
			},
			args: args{
				name: "Ratatouille",
			},
			want: Recipe{},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err == NotFoundErr
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MemStore{
				list: tt.fields.list,
			}
			got, err := m.Get(tt.args.name)
			if tt.wantErr != nil {
				if !tt.wantErr(t, err, fmt.Sprintf("Get(%v)", tt.args.name)) {
					require.Failf(t, "Invalid error message", "Got: %v", err.Error())
				}
			} else {
				assert.NoError(t, err)
			}

			assert.Equalf(t, tt.want, got, "Get(%v)", tt.args.name)
		})
	}
}

func TestMemStore_List(t *testing.T) {
	type fields struct {
		list map[string]Recipe
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]Recipe
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Simple list",
			fields: fields{
				map[string]Recipe{
					"Ham and cheese toasties": getHamCheeseToasties(),
				},
			},
			want: map[string]Recipe{
				"Ham and cheese toasties": getHamCheeseToasties(),
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MemStore{
				list: tt.fields.list,
			}
			got, err := m.List()
			if tt.wantErr != nil {
				if !tt.wantErr(t, err, fmt.Sprintf("List()")) {
					assert.Fail(t, "Invalid error")
				}
			} else {
				assert.NoError(t, err)
			}

			assert.Equalf(t, tt.want, got, "List()")
		})
	}
}

func TestMemStore_Remove(t *testing.T) {
	type fields struct {
		list map[string]Recipe
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
		wantLen int
	}{
		{
			name: "Empty list",
			fields: fields{
				map[string]Recipe{
					"Ham and cheese toasties": getHamCheeseToasties(),
				},
			},
			args: args{
				name: "Ham and cheese toasties",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MemStore{
				list: tt.fields.list,
			}

			err := m.Remove(tt.args.name)

			if tt.wantErr != nil {
				if !tt.wantErr(t, err, fmt.Sprintf("List()")) {
					assert.Fail(t, "Invalid error")
				}
			} else {
				assert.NoError(t, err)
			}

			assert.Len(t, m.list, tt.wantLen)
		})
	}
}

func TestMemStore_Update(t *testing.T) {
	type fields struct {
		list map[string]Recipe
	}
	type args struct {
		name   string
		recipe Recipe
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
		wantLen int
	}{
		{
			name: "Update butter to Ham and cheese",
			fields: fields{
				map[string]Recipe{
					"Ham and cheese toasties": getHamCheeseToasties(),
				},
			},
			args: args{
				name: "Ham and cheese toasties",
				recipe: Recipe{
					Name: "Ham and cheese toasties",
					Ingredients: []Ingredient{
						{Name: "bread"},
						{Name: "ham"},
						{Name: "cheese"},
						{Name: "butter"},
					},
				},
			},
			wantErr: nil,
			wantLen: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MemStore{
				list: tt.fields.list,
			}

			err := m.Update(tt.args.name, tt.args.recipe)
			if tt.wantErr != nil {
				if !tt.wantErr(t, err, fmt.Sprintf("List()")) {
					assert.Fail(t, "Invalid error")
				}
			} else {
				assert.NoError(t, err)
			}

			assert.Len(t, m.list, tt.wantLen)
		})
	}
}
