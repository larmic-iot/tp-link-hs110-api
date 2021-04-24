package http

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
)

func TestErrorEncoder_Encode(t *testing.T) {
	type args struct {
		statusCode int
		message    string
	}
	tests := []struct {
		name           string
		args           args
		wantStatusCode int
		wantMessage    string
	}{
		{name: "status not found", args: args{http.StatusNotFound, "some message"}, wantStatusCode: http.StatusNotFound, wantMessage: "some message"},
		{name: "empty message", args: args{http.StatusOK, ""}, wantStatusCode: http.StatusOK, wantMessage: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			responseWriter := httptest.NewRecorder()
			NewErrorEncoder(responseWriter).Encode(tt.args.statusCode, tt.args.message)

			gotContentType := responseWriter.Header().Get("Content-Type")
			wantContentType := "application/json; charset=UTF-8"

			if !reflect.DeepEqual(gotContentType, wantContentType) {
				t.Errorf("Encode() [content-type] got = %v, want %v", gotContentType, wantContentType)
			}

			if responseWriter.Code != tt.wantStatusCode {
				t.Errorf("Encode() [status-code] got = %v, want %v", responseWriter.Code, tt.wantStatusCode)
			}

			gotMessage := string(responseWriter.Body.Bytes())
			wantMessage := "{\"code\":" + strconv.Itoa(responseWriter.Code) + ",\"message\":\"" + tt.wantMessage + "\"}\n"

			if gotMessage != wantMessage {
				t.Errorf("Encode() [message] got = %v, want %v", gotMessage, wantMessage)
			}
		})
	}
}
