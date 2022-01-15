package stack

import (
	"bufio"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	sLinkedList := NewStackLinkedList()
	assert.Equal(t, true, sLinkedList.IsEmpty())
	assert.Equal(t, "", sLinkedList.Pop())
	sLinkedList.Push("a")
	sLinkedList.Push("b")
	sLinkedList.Push("c")
	result := sLinkedList.Pop()
	assert.Equal(t, "c", result)
	assert.Equal(t, false, sLinkedList.IsEmpty())

	sArray := NewStackLinkedList()
	assert.Equal(t, true, sArray.IsEmpty())
	assert.Equal(t, "", sArray.Pop())
	sArray.Push("a")
	sArray.Push("b")
	sArray.Push("c")
	result = sArray.Pop()
	assert.Equal(t, "c", result)
	assert.Equal(t, false, sArray.IsEmpty())
}

func TestArithmeticExpressionEvaluation(t *testing.T) {
	vals := NewStackLinkedList()
	ops := NewStackLinkedList()
	f, _ := os.Open("arithmetic.txt")
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)
	for s.Scan() {
		text := s.Text()
		switch text {
		// ignore left parenthesis
		case "(":
		case "+":
			ops.Push(text)
		case "*":
			ops.Push(text)
		case ")":
			num1, _ := strconv.Atoi(vals.Pop())
			num2, _ := strconv.Atoi(vals.Pop())
			op := ops.Pop()
			if op == "+" {
				vals.Push(strconv.Itoa(num1 + num2))
			} else if op == "*" {
				vals.Push(strconv.Itoa(num1 * num2))
			}
		default:
			vals.Push(text)
		}
	}
	assert.Equal(t, "101", vals.Pop())
}

func TestLinter(t *testing.T) {
	stack := NewStackLinkedList()
	f, _ := os.Open("linter.txt")
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)
	for s.Scan() {
		text := s.Text()
		switch text {
		case "(":
			stack.Push(text)
		case "[":
			stack.Push(text)
		case "{":
			stack.Push(text)
		case ")":
			el := stack.Pop()
			if el != "(" {
				t.Fatalf("expect to get ( but got %s", el)
			}
		case "]":
			el := stack.Pop()
			if el != "[" {
				t.Fatalf("expect to get [ but got %s", el)
			}
		case "}":
			el := stack.Pop()
			if el != "{" {
				t.Fatalf("expect to get { but got %s", el)
			}
		}
	}
	if !stack.IsEmpty() {
		t.Error("expect stack to be empty")
	}
}
