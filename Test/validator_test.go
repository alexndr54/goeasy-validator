package Test

import (
	"fmt"
	"github.com/alexndr54/goeasy-validator/Validation"
	"github.com/alexndr54/goeasy-validator/Validation/Helper"
	"testing"
)

func TestNewValidatorByMethod(t *testing.T) {
	t.Run("Test validasi data pengguna", func(t *testing.T) {
		t.Log("--- Validasi Data Pengguna 1 (Valid) ---")
		userData1 := map[string]interface{}{
			"nama_depan": "Budi",
			"email":      "budi.santoso@example.com",
			"kata_sandi": "Password123",
		}

		rules1 := map[string]string{
			"nama_depan": "required|min:2",
			"email":      "required|email",
			"kata_sandi": "required|password",
		}

		validator1 := Validation.NewValidator(userData1, rules1)
		errors1 := validator1.Validate()

		if errors1.HasErrors() {
			t.Errorf("Kesalahan validasi:")
			for field, msgs := range errors1 {
				for _, msg := range msgs {
					t.Errorf("- %s: %s\n", field, msg)
				}
			}
		} else {
			t.Logf("Data pengguna 1 valid.")
		}
	})

	t.Run("Test validasi data pengguna 2 (Email Tidak Valid)", func(t *testing.T) {
		t.Log("\n--- Validasi Data Pengguna 2 (Tidak Valid) ---")
		userData2 := map[string]interface{}{
			"nama_depan": "a",             // Terlalu pendek
			"email":      "invalid-email", // Format salah
			"kata_sandi": "short",         // Terlalu pendek, tidak ada angka/huruf besar
			"umur":       10,              // Field tidak didefinisikan di rules, akan diabaikan
		}

		rules2 := map[string]string{
			"nama_depan": "required|min:2",
			"email":      "required|email",
			"kata_sandi": "required|password",
			"nomor_hp":   "required", // Field ini tidak ada di data, tapi required
		}

		validator2 := Validation.NewValidator(userData2, rules2)
		errors2 := validator2.Validate()

		if errors2.HasErrors() {
			t.Log("Kesalahan validasi:")
			for field, msgs := range errors2 {
				for _, msg := range msgs {
					fmt.Printf("- %s: %s\n", field, msg)
				}
			}
		} else {
			t.Log("Data pengguna 2 valid.")
		}
	})

	t.Run("Test validasi data pengguna 3 (Field Tidak Ada)", func(t *testing.T) {
		t.Log("\n--- Validasi Data Pengguna 3 (Input dari Struct) ---")
		type FormData struct {
			FirstName string
			EmailAddr string
			Password  string
		}
		dataFromForm := FormData{
			FirstName: "Jane",
			EmailAddr: "jane.doe@example.com",
			Password:  "SecurePassword123",
		}

		// Ubah struct menjadi map untuk validator
		dataMap := Helper.StructToMap(dataFromForm)

		rules3 := map[string]string{
			"firstname": "required|min:3", // Perhatikan nama field di map akan jadi lowercase
			"emailaddr": "required|email",
			"password":  "required|password",
		}

		validator3 := Validation.NewValidator(dataMap, rules3)
		errors3 := validator3.Validate()

		if errors3.HasErrors() {
			t.Error("Kesalahan validasi:")
			for field, msgs := range errors3 {
				for _, msg := range msgs {
					t.Errorf("- %s: %s\n", field, msg)
				}
			}
		} else {
			t.Log("Data pengguna 3 valid.")
		}

	})

	t.Run("Test validasi data pengguna 4 (Field Tidak Ada dan Tidak Required)", func(t *testing.T) {
		t.Log("\n--- Validasi Data Pengguna 4 (Password Lemah) ---")
		userData4 := map[string]interface{}{
			"nama_depan": "Andi",
			"email":      "andi@example.com",
			"kata_sandi": "lem4h", // Tidak memenuhi aturan password
		}

		rules4 := map[string]string{
			"nama_depan": "required|min:2",
			"email":      "required|email",
			"kata_sandi": "required|password",
		}

		validator4 := Validation.NewValidator(userData4, rules4)
		errors4 := validator4.Validate()

		if errors4.HasErrors() {
			t.Error("Kesalahan validasi:")
			for field, msgs := range errors4 {
				for _, msgg := range msgs {
					t.Errorf("- %s: %s\n", field, msgg)
				}
			}
		} else {
			t.Log("Data pengguna 4 valid.")
		}
	})
}

func TestRules(t *testing.T) {
	data := map[string]interface{}{
		"nama_lengkap":  "awgfjjedugfjeJLGF",
		"nama_lengkap2": "sad",
	}
	rules := map[string]string{
		"nama_lengkap":  "required|password|min:5",
		"nama_lengkap2": "required|password|min:5",
	}

	simple := Validation.NewSimpleValidator(data, rules)
	field, msg := Validation.NewSingleRuleValidator(data, rules)

	fmt.Println(simple)
	fmt.Println(field, msg)

}
