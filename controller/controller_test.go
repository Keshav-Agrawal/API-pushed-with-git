package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Keshav-Agrawal/mongoapi/datasource/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//
//func TestGetMyAllTask(t *testing.T) {
//	req, err := http.NewRequest("GET", "/api/tasks", nil)
//	if err != nil {
//		t.Fatalf("%v", err)
//	}
//	rec := httptest.NewRecorder()
//	GetMyAllTask(rec, req)
//	res := rec.Result()
//	defer res.Body.Close()
//	if res.StatusCode != http.StatusOK {
//		t.Errorf("%v", res.Status)
//	}
//	b, err := ioutil.ReadAll(res.Body)
//	str := string(b)
//	fmt.Println(str)
//	if err != nil {
//		t.Fatalf("%v", err)
//	}
//	//fmt.Println("b", b)
//	//c := client.Ping(context.Background(), nil)
//	//t.Log(c)
//}
//func TestInsertTask(t *testing.T) {
//	/*arg := model.Homework{
//		Task: "tom",
//		Done: false,
//	}*/
//	postBody, _ := json.Marshal(model.Homework{
//		Task: "tom",
//		Done: false,
//	})
//	responseBody := bytes.NewBuffer(postBody)
//	req, err := http.NewRequest("POST", "/api/task", responseBody)
//	if err != nil {
//		t.Fatalf("%v", err)
//	}
//	rec := httptest.NewRecorder()
//	CreateTask(rec, req)
//	res := rec.Result()
//	defer res.Body.Close()
//	if res.StatusCode != http.StatusOK {
//		t.Errorf("%v", res.Status)
//	}
//	b, err := ioutil.ReadAll(res.Body)
//	str := string(b)
//	fmt.Println(str)
//	if err != nil {
//		t.Fatalf("%v", err)
//	}
//	//fmt.Println("b", b)
//	//c := client.Ping(context.Background(), nil)
//	//t.Log(c)
//}
//
//func TestUpdateTask(t *testing.T) {
//
//	postBody, _ := json.Marshal(model.Homework{
//		Task: "1234",
//		Done: true,
//	})
//	client := &http.Client{}
//	json, err := json.Marshal(postBody)
//	if err != nil {
//		panic(err)
//	}
//	req, err := http.NewRequest(http.MethodPut, "http://localhost:4000/api/task/621f1031ddc5de508f36d0db", bytes.NewBuffer(json))
//	if err != nil {
//		panic(err)
//	}
//	req.Header.Set("Content-Type", "application/json; charset=utf-8")
//	resp, err := client.Do(req)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(resp.StatusCode)
//	/*MarkAsDone(resp, req)
//	res := rec.Result()
//	defer res.Body.Close()
//	if res.StatusCode != http.StatusOK {
//		t.Errorf("%v", res.Status)
//	}
//	b, err := ioutil.ReadAll(res.Body)
//	str := string(b)
//	fmt.Println(str)
//	if err != nil {
//		t.Fatalf("%v", err)
//	}
//	//fmt.Println("b", b)
//	//c := client.Ping(context.Background(), nil)
//	//t.Log(c)
//	*/
//}

func TestGetMyAllTask(t *testing.T) {
	isSuccess := false
	mockDB := mongo.NewMock(func() (interface{}, error) {
		t.Log("My custom mock function init")
		if isSuccess {
			return []primitive.M{
				map[string]interface{}{
					"1": 1,
				},
			}, nil
		}
		return nil, errors.New("custom error")
	})

	controller := NewHomeWorkService(mockDB)

	tests := []struct {
		name     string
		expected string
		success  bool
		code     int
	}{
		{
			name:     "failure case",
			expected: `"custom error"`,
			success:  false,
			code:     http.StatusInternalServerError,
		},
		{
			name:     "success case",
			expected: `"custom error"`,
			success:  true,
			code:     http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isSuccess = tt.success

			w := httptest.NewRecorder()
			//r := httptest.NewRequest("", "some", bytes.NewBuffer([]byte("")))
			controller.CreateTask(w, nil)

			if w.Code != tt.code {
				t.Errorf("Want: %+v, Got: %+v", tt.code, w.Code)
			}

			//if w.Body.String() != tt.expected {
			//	t.Errorf("Want: %+v, Got: %+v", tt.expected, w.Body.String())
			//}
		})
	}
}
func TestCreateTask(t *testing.T) {
	isSuccess := false
	mockDB := mongo.NewMock(func() (interface{}, error) {
		t.Log("My custom mock function init")
		if isSuccess {
			return []primitive.M{
				map[string]interface{}{
					"1": 1,
				},
			}, nil
		}
		return nil, errors.New("custom error")
	})

	controller := NewHomeWorkService(mockDB)

	tests := []struct {
		name     string
		expected string
		success  bool
		code     int
	}{
		{
			name:     "failure case",
			expected: `"custom error"`,
			success:  false,
			code:     http.StatusInternalServerError,
		},
		{
			name:     "success case",
			expected: `"custom error"`,
			success:  true,
			code:     http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isSuccess = tt.success

			w := httptest.NewRecorder()
			//r := httptest.NewRequest("", "some", bytes.NewBuffer([]byte("")))
			controller.GetMyAllTask(w, nil)

			if w.Code != tt.code {
				t.Errorf("Want: %+v, Got: %+v", tt.code, w.Code)
			}

			//if w.Body.String() != tt.expected {
			//	t.Errorf("Want: %+v, Got: %+v", tt.expected, w.Body.String())
			//}
		})
	}
}
