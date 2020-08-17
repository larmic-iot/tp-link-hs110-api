package crypto

import (
	"reflect"
	"testing"
)

func TestEncryptor_encrypt(t *testing.T) {
	tests := []struct {
		name    string
		message string
		want    []int32
	}{
		{"info", "{\"system\":{\"get_sysinfo\":null}}", []int32{208, 242, 129, 248, 139, 255, 154, 247, 213, 239, 148, 182, 209, 180, 192, 159, 236, 149, 230, 143, 225, 135, 232, 202, 240, 158, 235, 135, 235, 150, 235}},
		{"on", "{\"system\":{\"set_relay_state\":{\"state\":1}}}}", []int32{208, 242, 129, 248, 139, 255, 154, 247, 213, 239, 148, 182, 197, 160, 212, 139, 249, 156, 240, 145, 232, 183, 196, 176, 209, 165, 192, 226, 216, 163, 129, 242, 134, 231, 147, 246, 212, 238, 223, 162, 223, 162, 223}},
		{"off", "{\"system\":{\"set_relay_state\":{\"state\":0}}}}", []int32{208, 242, 129, 248, 139, 255, 154, 247, 213, 239, 148, 182, 197, 160, 212, 139, 249, 156, 240, 145, 232, 183, 196, 176, 209, 165, 192, 226, 216, 163, 129, 242, 134, 231, 147, 246, 212, 238, 222, 163, 222, 163, 222}},
		{"emeter", "{\"emeter\":{\"get_realtime\":null}}", []int32{208, 242, 151, 250, 159, 235, 142, 252, 222, 228, 159, 189, 218, 191, 203, 148, 230, 131, 226, 142, 250, 147, 254, 155, 185, 131, 237, 152, 244, 152, 229, 152}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewEncryptor(false)
			if got := e.encrypt(tt.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
