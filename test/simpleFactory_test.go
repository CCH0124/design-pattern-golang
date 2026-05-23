package test

import (
	"designpattern/pkg/creational/factory/simple"
	"testing"
)

func TestCreatePizza(t *testing.T) {
	factory := simple.NewSimplePizzaFactory()
	// 測試案例
	tests := []struct {
		pizzaType string
		wantErr   bool
	}{
		{"nycheese", false},
		{"chicago", false},
		{"unknown", true},
	}

	for _, tt := range tests {
		pizza, err := factory.CreatePizza(tt.pizzaType)
		if tt.wantErr && err == nil {
			t.Errorf("expected error for type %s, got nil", tt.pizzaType)
		}
		if !tt.wantErr && err != nil {
			t.Errorf("unexpected error for type %s: %v", tt.pizzaType, err)
		}
		if !tt.wantErr && pizza == nil {
			t.Errorf("expected pizza for type %s, got nil", tt.pizzaType)
		}
	}
}
