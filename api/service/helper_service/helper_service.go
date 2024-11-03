package HelperService

func IntPtr(i int) *int {
	return &i
}

func StrPtr(s string) *string {
	return &s
}

func BoolPtr(b bool) *bool {
	return &b
}

func IntBoardArrayPtr(a [8][8]int) *[8][8]int {
	return &a
}

func IntValue(p *int) interface{} {
	if p != nil {
		return *p
	}
	return nil
}

func BoolValue(p *bool) interface{} {
	if p != nil {
		return *p
	}
	return nil
}
