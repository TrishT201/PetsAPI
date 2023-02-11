/*
response: After the client creates and sends a request
	and the server matches the path in the request to
	the correct route, the server then return an
	HTTP Response back to the client.
*/

package responses

// the structure represents the API response
type PetResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
