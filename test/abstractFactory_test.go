package test

import (
	"designpattern/pkg/creational/factory/abstract"
	"strings"
	"testing"
)

func TestNYPizzaStore(t *testing.T) {
	store := abstract.NewNYPizzaStore()

	// Test cheese pizza
	pizza := store.OrderPizza("cheese")
	if pizza == nil {
		t.Fatal("expected NY cheese pizza, got nil")
	}
	if pizza.GetName() != "New York Style Cheese Pizza" {
		t.Errorf("expected pizza name 'New York Style Cheese Pizza', got '%s'", pizza.GetName())
	}

	details := pizza.String()
	if !strings.Contains(details, "Thin Crust Dough") {
		t.Errorf("expected ingredients to contain 'Thin Crust Dough', got details:\n%s", details)
	}
	if !strings.Contains(details, "Marinara Sauce") {
		t.Errorf("expected ingredients to contain 'Marinara Sauce', got details:\n%s", details)
	}

	// Test unsupported type
	invalidPizza := store.OrderPizza("invalid_type")
	if invalidPizza != nil {
		t.Errorf("expected nil for invalid pizza type, got %v", invalidPizza)
	}
}

func TestChicagoPizzaStore(t *testing.T) {
	store := abstract.NewChicagoPizzaStore()

	// Test veggie pizza
	pizza := store.OrderPizza("veggie")
	if pizza == nil {
		t.Fatal("expected Chicago veggie pizza, got nil")
	}
	if pizza.GetName() != "Chicago Style Veggie Pizza" {
		t.Errorf("expected pizza name 'Chicago Style Veggie Pizza', got '%s'", pizza.GetName())
	}

	details := pizza.String()
	if !strings.Contains(details, "Thick Crust Dough") {
		t.Errorf("expected ingredients to contain 'Thick Crust Dough', got details:\n%s", details)
	}
	if !strings.Contains(details, "Black Olives") {
		t.Errorf("expected ingredients to contain 'Black Olives', got details:\n%s", details)
	}
}
