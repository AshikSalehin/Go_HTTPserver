# Go_HTTPserver
Here I created a HTTP server script using GO Language. This includes public and private end points, GET and POST methods, JSON and Form Data Operations, HTML file serving.

----------------------------------------------------------------------Explanation of Endpoints:----------------------------------------------------------------------

Public Endpoints:
  /public/json:
        Method: GET
        Response: JSON with a message.
        Purpose: A simple public endpoint that returns a JSON response.
  /public/form:
        Method: GET
        Response: HTML form.
        Purpose: To display a form where the user can submit data.
  /public/form-post:
        Method: POST
        Response: Text response after form submission.
        Purpose: Handles the POST request from the form submission.
        
Private Endpoints:
  /private/json:
        Method: GET
        Response: Protected JSON response.
        Purpose: A private endpoint protected by basic authentication.
  /private/post:
        Method: POST
        Request: JSON payload.
        Response: JSON response confirming receipt of data.
        Purpose: A private endpoint that handles POST requests and processes JSON data.

----------------------------------------------------------------------Testing the Server:--------------------------------------------------------------------------------------------------

Run the Go server:      go run main.go

For public JSON (/public/json), access it via:      http://localhost:8080/public/json

For public HTML form (/public/form), and submit the form to trigger a POST request to /public/form-post; access it via:      http://localhost:8080/public/form

For private JSON (/private/json), use Basic Authentication (username: admin, password: password). You can use a tool like Postman or curl to send requests:      curl -u admin:password http://localhost:8080/private/json

For private POST (/private/post), send a JSON POST request with Basic Authentication:      curl -u admin:password -X POST -d '{"data":"hello"}' -H "Content-Type: application/json" http://localhost:8080/private/post

-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
