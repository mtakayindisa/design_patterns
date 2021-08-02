package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	instance = GetInstance()
	if instance == nil{
		t.Error("Expected singleton instance after initialization")
	}
}

func TestSingleton_AddOne(t *testing.T) {
	instance = GetInstance()
	instance.AddOne()
	count := instance.AddOne()
	if count != 2 {
		t.Error("Expected singleton instance to have count of 2 after calling AddOne")
	}
}
