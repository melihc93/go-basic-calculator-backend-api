package endpointhandlers

import (
	"encoding/json"
	"fmt"
	"go-basic-calculator-backend-api/config"
	"go-basic-calculator-backend-api/src/calculator"
	"go-basic-calculator-backend-api/src/errorhandlers"
	"net/http"
	"strings"
)

type responseObject struct {
	Result string `json:"result"`
	User   string `json:"user"`
}

type errorObject struct {
	Message string `json:"message"`
}

var responseObjects responseObject
var errorObjects errorObject

func DivResponse(w http.ResponseWriter, r *http.Request) {
	divHeaders := r.Header
	valueOfUsername, isUsernameExist := divHeaders["Username"]
	valueOfPassword, isPasswordExist := divHeaders["Password"]
	query := strings.Split(r.URL.Query().Get("params"), ",")
	if isUsernameExist {
		fmt.Printf("username exist in request headers with the value of %s\n", valueOfUsername)
		if isPasswordExist {
			fmt.Printf("password exist in request header with the value of %s\n", valueOfPassword)
			if valueOfUsername[config.HEADER_INDEX] == config.AUTHORISED_USERNAME && valueOfPassword[config.HEADER_INDEX] == config.AUTHORISED_USERNAME_PASSWORD {
				fmt.Printf("Successfully logged in with %s id\n", valueOfUsername[config.HEADER_INDEX])
				if len(query) > config.MAXIMUM_PARAM_LENGTH || len(query) == config.MINIMUM_PARAM_LENGTH {
					fmt.Printf("Query param length should be greater than 0 and less than 5, current length is: %d\n", len(query))
					errorObjects.Message = errorhandlers.QueryParamOutOfRange()
					w.WriteHeader(config.HTML_NOT_AVAILABLE)
					w.Header().Set("Access-Control-Allow-Origin", "*")
					json.NewEncoder(w).Encode(errorObjects)
				} else {
					fmt.Println("Division response object creating...")
					fmt.Printf("Param object: %s\n type: %T\n its length: %d\n", query, query, len(query))
					responseObjects.User = valueOfUsername[config.HEADER_INDEX]
					responseObjects.Result = calculator.Divider(query)
					w.Header().Set("Access-Control-Allow-Origin", "*")
					json.NewEncoder(w).Encode(responseObjects)
				}
			} else {
				fmt.Println("Invalid username or password\n")
				errorObjects.Message = errorhandlers.InvalidUsernameOrPassword()
				w.WriteHeader(config.HTML_UNAUTHORIZED_RESPONSE)
				w.Header().Set("Access-Control-Allow-Origin", "*")
				json.NewEncoder(w).Encode(errorObjects)
			}
		} else {
			fmt.Println("Password not exist\n")
			errorObjects.Message = errorhandlers.PasswordNotExist()
			w.WriteHeader(config.HTML_BAD_REQUEST)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			json.NewEncoder(w).Encode(errorObjects)
		}
	} else {
		fmt.Println("Username not exist!\n")
		errorObjects.Message = errorhandlers.UsernameNotExist()
		w.WriteHeader(config.HTML_BAD_REQUEST)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(errorObjects)
	}
}

func AddResponse(w http.ResponseWriter, r *http.Request) {
	divHeaders := r.Header
	valueOfUsername, isUsernameExist := divHeaders["Username"]
	valueOfPassword, isPasswordExist := divHeaders["Password"]
	query := strings.Split(r.URL.Query().Get("params"), ",")
	if isUsernameExist {
		fmt.Printf("username exist in request headers with the value of %s\n", valueOfUsername)
		if isPasswordExist {
			fmt.Printf("password exist in request header with the value of %s\n", valueOfPassword)
			if valueOfUsername[config.HEADER_INDEX] == config.AUTHORISED_USERNAME && valueOfPassword[config.HEADER_INDEX] == config.AUTHORISED_USERNAME_PASSWORD {
				fmt.Printf("Successfully logged in with %s id\n", valueOfUsername[config.HEADER_INDEX])
				if len(query) > config.MAXIMUM_PARAM_LENGTH || len(query) == config.MINIMUM_PARAM_LENGTH {
					fmt.Printf("Query param length should be greater than 0 and less than 5, current length is: %d\n", len(query))
					errorObjects.Message = errorhandlers.QueryParamOutOfRange()
					w.WriteHeader(config.HTML_NOT_AVAILABLE)
					w.Header().Set("Access-Control-Allow-Origin", "*")
					json.NewEncoder(w).Encode(errorObjects)
				} else {
					fmt.Println("Adding response object creating...")
					fmt.Printf("Param object: %s\n type: %T\n its length: %d\n", query, query, len(query))
					responseObjects.User = valueOfUsername[config.HEADER_INDEX]
					responseObjects.Result = calculator.Adder(query)
					w.Header().Set("Access-Control-Allow-Origin", "*")
					json.NewEncoder(w).Encode(responseObjects)
				}
			} else {
				fmt.Println("Invalid username or password\n")
				errorObjects.Message = errorhandlers.InvalidUsernameOrPassword()
				w.WriteHeader(config.HTML_UNAUTHORIZED_RESPONSE)
				w.Header().Set("Access-Control-Allow-Origin", "*")
				json.NewEncoder(w).Encode(errorObjects)
			}
		} else {
			fmt.Println("Password not exist\n")
			errorObjects.Message = errorhandlers.PasswordNotExist()
			w.WriteHeader(config.HTML_BAD_REQUEST)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			json.NewEncoder(w).Encode(errorObjects)
		}
	} else {
		fmt.Println("Username not exist!\n")
		errorObjects.Message = errorhandlers.UsernameNotExist()
		w.WriteHeader(config.HTML_BAD_REQUEST)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(errorObjects)
	}
}

func SumResponse(w http.ResponseWriter, r *http.Request) {
	divHeaders := r.Header
	valueOfUsername, isUsernameExist := divHeaders["Username"]
	valueOfPassword, isPasswordExist := divHeaders["Password"]
	query := strings.Split(r.URL.Query().Get("params"), ",")
	if isUsernameExist {
		fmt.Printf("username exist in request headers with the value of %s\n", valueOfUsername)
		if isPasswordExist {
			fmt.Printf("password exist in request header with the value of %s\n", valueOfPassword)
			if valueOfUsername[config.HEADER_INDEX] == config.AUTHORISED_USERNAME && valueOfPassword[config.HEADER_INDEX] == config.AUTHORISED_USERNAME_PASSWORD {
				fmt.Printf("Successfully logged in with %s id\n", valueOfUsername[config.HEADER_INDEX])
				if len(query) > config.MAXIMUM_PARAM_LENGTH || len(query) == config.MINIMUM_PARAM_LENGTH {
					fmt.Printf("Query param length should be greater than 0 and less than 5, current length is: %d\n", len(query))
					errorObjects.Message = errorhandlers.QueryParamOutOfRange()
					w.WriteHeader(config.HTML_NOT_AVAILABLE)
					w.Header().Set("Access-Control-Allow-Origin", "*")
					json.NewEncoder(w).Encode(errorObjects)
				} else {
					fmt.Println("Division response object creating...")
					fmt.Printf("Param object: %s\n type: %T\n its length: %d\n", query, query, len(query))
					responseObjects.User = valueOfUsername[config.HEADER_INDEX]
					responseObjects.Result = calculator.Summer(query)
					w.Header().Set("Access-Control-Allow-Origin", "*")
					json.NewEncoder(w).Encode(responseObjects)
				}
			} else {
				fmt.Println("Invalid username or password\n")
				errorObjects.Message = errorhandlers.InvalidUsernameOrPassword()
				w.WriteHeader(config.HTML_UNAUTHORIZED_RESPONSE)
				w.Header().Set("Access-Control-Allow-Origin", "*")
				json.NewEncoder(w).Encode(errorObjects)
			}
		} else {
			fmt.Println("Password not exist\n")
			errorObjects.Message = errorhandlers.PasswordNotExist()
			w.WriteHeader(config.HTML_BAD_REQUEST)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			json.NewEncoder(w).Encode(errorObjects)
		}
	} else {
		fmt.Println("Username not exist!\n")
		errorObjects.Message = errorhandlers.UsernameNotExist()
		w.WriteHeader(config.HTML_BAD_REQUEST)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(errorObjects)
	}
}

func SubResponse(w http.ResponseWriter, r *http.Request) {
	divHeaders := r.Header
	valueOfUsername, isUsernameExist := divHeaders["Username"]
	valueOfPassword, isPasswordExist := divHeaders["Password"]
	query := strings.Split(r.URL.Query().Get("params"), ",")
	if isUsernameExist {
		fmt.Printf("username exist in request headers with the value of %s\n", valueOfUsername)
		if isPasswordExist {
			fmt.Printf("password exist in request header with the value of %s\n", valueOfPassword)
			if valueOfUsername[config.HEADER_INDEX] == config.AUTHORISED_USERNAME && valueOfPassword[config.HEADER_INDEX] == config.AUTHORISED_USERNAME_PASSWORD {
				fmt.Printf("Successfully logged in with %s id\n", valueOfUsername[config.HEADER_INDEX])
				if len(query) > config.MAXIMUM_PARAM_LENGTH || len(query) == config.MINIMUM_PARAM_LENGTH {
					fmt.Printf("Query param length should be greater than 0 and less than 5, current length is: %d\n", len(query))
					errorObjects.Message = errorhandlers.QueryParamOutOfRange()
					w.WriteHeader(config.HTML_NOT_AVAILABLE)
					w.Header().Set("Access-Control-Allow-Origin", "*")
					json.NewEncoder(w).Encode(errorObjects)
				} else {
					fmt.Println("Division response object creating...")
					fmt.Printf("Param object: %s\n type: %T\n its length: %d\n", query, query, len(query))
					responseObjects.User = valueOfUsername[config.HEADER_INDEX]
					responseObjects.Result = calculator.Substracter(query)
					w.Header().Set("Access-Control-Allow-Origin", "*")
					json.NewEncoder(w).Encode(responseObjects)
				}
			} else {
				fmt.Println("Invalid username or password\n")
				errorObjects.Message = errorhandlers.InvalidUsernameOrPassword()
				w.WriteHeader(config.HTML_UNAUTHORIZED_RESPONSE)
				w.Header().Set("Access-Control-Allow-Origin", "*")
				json.NewEncoder(w).Encode(errorObjects)
			}
		} else {
			fmt.Println("Password not exist\n")
			errorObjects.Message = errorhandlers.PasswordNotExist()
			w.WriteHeader(config.HTML_BAD_REQUEST)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			json.NewEncoder(w).Encode(errorObjects)
		}
	} else {
		fmt.Println("Username not exist!\n")
		errorObjects.Message = errorhandlers.UsernameNotExist()
		w.WriteHeader(config.HTML_BAD_REQUEST)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(errorObjects)
	}
}

func MultResponse(w http.ResponseWriter, r *http.Request) {
	divHeaders := r.Header
	valueOfUsername, isUsernameExist := divHeaders["Username"]
	valueOfPassword, isPasswordExist := divHeaders["Password"]
	query := strings.Split(r.URL.Query().Get("params"), ",")
	if isUsernameExist {
		fmt.Printf("username exist in request headers with the value of %s\n", valueOfUsername)
		if isPasswordExist {
			fmt.Printf("password exist in request header with the value of %s\n", valueOfPassword)
			if valueOfUsername[config.HEADER_INDEX] == config.AUTHORISED_USERNAME && valueOfPassword[config.HEADER_INDEX] == config.AUTHORISED_USERNAME_PASSWORD {
				fmt.Printf("Successfully logged in with %s id\n", valueOfUsername[config.HEADER_INDEX])
				if len(query) > config.MAXIMUM_PARAM_LENGTH || len(query) == config.MINIMUM_PARAM_LENGTH {
					fmt.Printf("Query param length should be greater than 0 and less than 5, current length is: %d\n", len(query))
					errorObjects.Message = errorhandlers.QueryParamOutOfRange()
					w.WriteHeader(config.HTML_NOT_AVAILABLE)
					w.Header().Set("Access-Control-Allow-Origin", "*")
					json.NewEncoder(w).Encode(errorObjects)
				} else {
					fmt.Println("Division response object creating...")
					fmt.Printf("Param object: %s\n type: %T\n its length: %d\n", query, query, len(query))
					responseObjects.User = valueOfUsername[config.HEADER_INDEX]
					responseObjects.Result = calculator.Multiplier(query)
					w.Header().Set("Access-Control-Allow-Origin", "*")
					json.NewEncoder(w).Encode(responseObjects)
				}
			} else {
				fmt.Println("Invalid username or password\n")
				errorObjects.Message = errorhandlers.InvalidUsernameOrPassword()
				w.WriteHeader(config.HTML_UNAUTHORIZED_RESPONSE)
				w.Header().Set("Access-Control-Allow-Origin", "*")
				json.NewEncoder(w).Encode(errorObjects)
			}
		} else {
			fmt.Println("Password not exist\n")
			errorObjects.Message = errorhandlers.PasswordNotExist()
			w.WriteHeader(config.HTML_BAD_REQUEST)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			json.NewEncoder(w).Encode(errorObjects)
		}
	} else {
		fmt.Println("Username not exist!\n")
		errorObjects.Message = errorhandlers.UsernameNotExist()
		w.WriteHeader(config.HTML_BAD_REQUEST)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(errorObjects)
	}
}
