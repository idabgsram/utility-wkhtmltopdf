# Simple wkhtmltopdf Utility

1. Install docker engine or docker desktop

2. Make sure docker CLI is properly installed

3. Type : `docker pull surnet/alpine-wkhtmltopdf:3.17.0-0.12.6-full`

4. Install golang https://go.dev/doc/install

5. Execute : `go test -timeout 1h -run ^Test_BuildPDF$ guzram/utility/wkhtmltopdf -v -count=1`

6. Or Run `func Test_BuildPDF(t *testing.T) {` in the build_pdf_test.go