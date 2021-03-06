## INSTAGRAM-BACKEND-API
- This API is developed using go language with mongoDB as the database. 
- You can create user and post with the http calls listed down below.
- Also user details and post can be fetch corresponding to the gievn ID.
 
 #### :heavy_check_mark: Create an User
- Should be a POST request
- Use JSON request body
- URL should be ‘/users'

- ##### POST : http:localhost:9000/users
  ##### Body 
  ```
  {
      "Name": "XYZ ABC",
      "Email": "xyzabc@example.com",
      "Password": "xyzabc"
  }
  ```

#### :heavy_check_mark: Get a user using id
- Should be a GET request
- Id should be in the url parameter
- URL should be ‘/users/<id here>’
 - ##### To Get a user with given id say 61612bb9e11c71a9c06d5ea9: 
   ##### GET: :link: http://localhost:9000/users/61612bb9e11c71a9c06d5ea9

 
#### :heavy_check_mark: Create a Post
- Should be a POST request
- Use JSON request body
- URL should be ‘/post'
- ##### POST: http:localhost:9000/post
  ##### Body 
  ```
  {
      "Caption": "My first post",
      "Image_URL": "https://first_post.png",
      "UserId": "61612bb9e11c71a9c06d5ea9"
  }
  ```

#### :heavy_check_mark: Get a post using id
- Should be a GET request
- Id should be in the url parameter
- URL should be ‘/post/<id here>’
- ##### To get a post with given id say 61607b00e11c71fb8b4f433d:
  ##### GET: :link: http://localhost:9000/post/61607b00e11c71fb8b4f433d

#### :heavy_check_mark: List all posts of a user
- Should be a GET request
- URL should be ‘/posts/users/<Id here>'
- ##### To get all posts corresponding to a user with id say 61612bb9e11c71a9c06d5ea9:
  ##### GET: :link: http://localhost:9000/posts/users/61612bb9e11c71a9c06d5ea9

##### :heavy_check_mark: Passwords should be securely stored such they can't be reverse engineered

##### :heavy_check_mark: Make the server thread safe

##### :heavy_check_mark: Add pagination to the list endpoint with 2 posts per page
 - ##### To check pagination 
   ##### GET: :link: http://localhost:9000/posts/users/61612bb9e11c71a9c06d5ea9?page=2

 

###### :x: Add unit tests

##### :heavy_check_mark: The API should be developed using Go.

##### :heavy_check_mark: MongoDB should be used for storage.

##### :heavy_check_mark: Only packages/libraries listed [here](https://golang.org/pkg/) and [here](https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.4.0) can be used.


#### :heavy_check_mark: Users should have the following attributes
 - Id
 - Name
 - Email
 - Password

#### :heavy_check_mark: Posts should have the following Attributes. All fields are mandatory unless marked optional:
 - Id
 - Caption
 - Image URL
 - Posted Timestamp
