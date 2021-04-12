package math

import ("testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 3)
	if sum == 3{
		t.Log("the res is ok")
	} else {
		t.Fatal("is wrong")
	}

	sum = Add(3, 4)

	if sum == 3{
		t.Log("the res is ok")
	}else {
		t.Fatal("worng")
	}
}



