package serializers

import (
	"testing"

	"github.com/RenZorRUS/todo-backend/src/internal/adapters/errs"
	"github.com/RenZorRUS/todo-backend/src/internal/core/ports/serializers"
	"github.com/RenZorRUS/todo-backend/src/internal/tests"
	"github.com/RenZorRUS/todo-backend/src/internal/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type (
	User struct {
		Name string
		Age  int
	}

	UserWithMetaTags struct {
		Name     string `json:"name"`
		Age      int    `json:"years"`
		Password string `json:"-"`
		Email    string `json:"email"`
		Phone    string `json:"phone,omitempty"`
	}
)

func TestNewJSONSerializer(t *testing.T) {
	t.Parallel()

	t.Run("Should return JSON serializer without errors", func(t *testing.T) {
		t.Parallel()

		log := mocks.NewLogger()
		json, err := NewJSONSerializer(log)

		require.NoError(t, err)
		assert.NotNil(t, json)
		assert.Implements(t, (*serializers.JSONSerializer)(nil), json)
	})

	t.Run("Should return error if logger is nil", func(t *testing.T) {
		t.Parallel()

		json, err := NewJSONSerializer(nil)

		require.Error(t, err)
		assert.Nil(t, json)
		assert.Equal(t, errs.ErrLogNotSpecified, err)
	})
}

func TestJSONSerializer_Marshal_Check_Successful(t *testing.T) {
	t.Parallel()

	log := mocks.NewLogger()
	json, _ := NewJSONSerializer(log)

	cases := []tests.TestCase[any, string]{
		{
			Name: "Marshaling of a simple struct",
			Input: User{
				Name: "John Doe",
				Age:  30,
			},
			Expected: `{"Name":"John Doe","Age":30}`,
		},
		{
			Name: "Marshaling of a simple struct with meta tags",
			Input: UserWithMetaTags{
				Name:     "John Doe",
				Age:      30,
				Password: "secret",
				Email:    "p6e5s@example.com",
				Phone:    "",
			},
			Expected: `{"name":"John Doe","years":30,"email":"p6e5s@example.com"}`,
		},
		{
			Name: "Marshaling of a map",
			Input: map[string]any{
				"key1": "value1",
				"key2": 42,
			},
			Expected: `{"key1":"value1","key2":42}`,
		},
		{
			Name:     "Marshaling of a slice",
			Input:    []string{"apple", "banana", "cherry"},
			Expected: `["apple","banana","cherry"]`,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			result, err := json.Marshal(test.Input)

			require.NoError(t, err)
			assert.NotNil(t, result)
			assert.JSONEq(t, test.Expected, string(result))
		})
	}
}

func TestJSONSerializer_Marshal_Fail(t *testing.T) {
	t.Parallel()

	cases := []tests.TestCase[any, any]{
		{
			Name:     "Should return error if input has unsupported type (channel)",
			Input:    make(chan struct{}),
			Expected: struct{}{},
		},
		{
			Name:     "Should return error if input is nil",
			Input:    nil,
			Expected: struct{}{},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			log := mocks.NewLogger()
			log.On("Error", mock.Anything, mock.Anything).Once()
			json, _ := NewJSONSerializer(log)

			result, err := json.Marshal(test.Input)

			require.Error(t, err)
			assert.Nil(t, result)
			log.AssertExpectations(t)
		})
	}
}

func TestJSONSerializer_Unmarshal_Check_Successful(t *testing.T) {
	t.Parallel()

	log := mocks.NewLogger()
	json, _ := NewJSONSerializer(log)

	cases := []tests.TestCase[[]byte, any]{
		{
			Name:  "Unmarshaling of a simple struct",
			Input: []byte(`{"Name":"John Doe","Age":30}`),
			Expected: &User{
				Name: "John Doe",
				Age:  30,
			},
		},
		{
			Name:  "Unmarshaling of a simple struct with meta tags",
			Input: []byte(`{"name":"John Doe","years":30,"email":"p6e5s@example.com"}`),
			Expected: &UserWithMetaTags{
				Name:     "John Doe",
				Age:      30,
				Password: "",
				Email:    "p6e5s@example.com",
				Phone:    "",
			},
		},
		{
			Name:  "Unmarshaling of a map",
			Input: []byte(`{"key1":"value1","key2":42}`),
			Expected: &map[string]any{
				"key1": "value1",
				"key2": 42.0,
			},
		},
		{
			Name:  "Unmarshaling of a slice",
			Input: []byte(`["apple","banana","cherry"]`),
			Expected: &[]string{
				"apple",
				"banana",
				"cherry",
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			runUnmarshalValueTest(t, json, test.Input, test.Expected)
		})
	}
}

func runUnmarshalValueTest(
	t *testing.T,
	json serializers.JSONSerializer,
	input []byte,
	expected any,
) {
	t.Helper()

	var result any

	switch expected.(type) {
	case *User:
		result = new(User)
	case *UserWithMetaTags:
		result = new(UserWithMetaTags)
	case *map[string]any:
		result = new(map[string]any)
	case *[]string:
		result = new([]string)
	default:
		t.Fatalf("Unsupported type: %T", expected)
	}

	err := json.Unmarshal(input, result)

	require.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestJSONSerializer_Unmarshal_Fail(t *testing.T) {
	t.Parallel()

	cases := []tests.TestCase[[]byte, any]{
		{
			Name:     "Should return error if input is nil",
			Input:    nil,
			Expected: struct{}{},
		},
		{
			Name:     "Should return error if input is not a valid JSON",
			Input:    []byte(`{"Name":"John Doe","Age":30`),
			Expected: struct{}{},
		},
		{
			Name:     "Should return error if input is empty",
			Input:    []byte(``),
			Expected: struct{}{},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			runUnmarshalErrorTest(t, test.Input)
		})
	}
}

func runUnmarshalErrorTest(
	t *testing.T,
	input []byte,
) {
	t.Helper()

	var result any

	log := mocks.NewLogger()
	log.On("Error", mock.Anything, mock.Anything).Once()
	json, _ := NewJSONSerializer(log)

	err := json.Unmarshal(input, result)

	require.Error(t, err)
	assert.Nil(t, result)
	log.AssertExpectations(t)
}
