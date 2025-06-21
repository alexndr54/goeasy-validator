## CARA MENGGUNAKAN

### 1. Install

```bash
go get github.com/alexndr54/goeasy-validator
```

### 2. Contoh Pengunaan
  1. import
```bash
import ("github.com/alexndr54/goeasy-validator/Validation"
	"github.com/alexndr54/goeasy-validator/Validation/Helper"
)
```

  2. Map And Map
```bash 
      data := map[string]interface{}{
		"what": 32323,
	}
	rules := map[string]string{
		"what": "only_letters",
	}

	result := Validation.NewSimpleValidator(data, rules)
	if result != nil {
		t.Error("Expected no errors, but got:", result)
	} else {
		t.Log("Validation passed, no errors found.")
	}

```

atau
```bash
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
			fmt.Println("Kesalahan validasi:")
			for field, msgs := range errors1 {
				for _, msg := range msgs {
					t.Errorf("- %s: %s\n", field, msg)
				}
			}
		} else {
		  fmt.Println("Data pengguna 1 valid.")
		}
 ```

  3. Rules
<table>
<thead>
<tr>
<th>Rules</th>
<th>Contoh Dan Keterangan</th>
</tr>
</thead>
<tbody>
<tr>
<td> require </td>
<td>map[string]string{ "nama": "required" }, Wajib di isi</td>
</tr>

<tr>
<td> min:int </td>
<td>map[string]string{ "nama": "min:5" }, Ubah angka 5 dengan jumlah minimum</td>
</tr>

<tr>
<td> max:int </td>
<td>map[string]string{ "nama": "max:10" }, Ubah angka 10 dengan jumlah maksimal</td>
</tr>

<tr>
<td> only_letters </td>
<td>map[string]string{ "nama": "only_letters" }, Hanya boleh huruf besar & kecil</td>
</tr>

<tr>
<td> password </td>
<td>map[string]string{ "nama": "password" }, Yang di input wajib minimal 8 huruf dan mengandung huruf kecil, huruf besar, dan angka</td>
</tr>


<tr>
<td> email </td>
<td>map[string]string{ "nama": "email" }, Yang di input harus dalam format email</td>
</tr>
</tbody>
</table>
