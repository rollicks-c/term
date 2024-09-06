package num

import "strconv"

func ParseInt(exp string) (int, error) {
	val, err := strconv.Atoi(exp)
	if err != nil {
		return 0, err
	}
	return val, nil
}
