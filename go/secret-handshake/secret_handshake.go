package secret

var secrets = []struct {
	bit  uint
	move string
}{
	{1, "wink"},
	{2, "double blink"},
	{4, "close your eyes"},
	{8, "jump"},
}

// Handshake outputs a secret handshake determined by set bits in code
func Handshake(code uint) []string {
	var hs = []string{}
	for _, secret := range secrets {
		if code&secret.bit == secret.bit {
			hs = append(hs, secret.move)
		}
	}

	if code&16 == 16 {
		last := len(hs) - 1
		for i := 0; i < len(hs)/2; i++ {
			hs[i], hs[last-i] = hs[last-i], hs[i]
		}
	}

	return hs
}
