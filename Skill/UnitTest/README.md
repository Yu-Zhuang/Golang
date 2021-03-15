# unit test

1. 在要測試的package中新增 _test.go檔案, 可為 name1_test.go or name2_test.go, 因為testing會辨識```_test.go```
2. 撰寫測試的包為 ```testing```
3. 測試的func必須為以 ***Test*** 開頭的名字, 且須引入```*testing.T```
4. func內容建議: 
```go
func TestFuncName(t *testing.T) {
  input := some []input
  
  expect := some []expect value
  
  for i := range expect {
		output := FuncName(input[i])
		if output != expect[i] {
			t.Errorf("FuncName(%v) = \"%v\", expect \"%v\"", input[i], output, expect[i])
		}
	}
}
```
5. 執行測試的command: ```go test $path/packageFolderName -v ```
