# What is CORS?
CORS is a security feature implemented by web browsers to restrict cross-origin requests by default. This is done to prevent unauthorized access to resources and protect user data. CORS is a set of HTTP headers and mechanisms that allow servers to indicate which domains are allowed to make cross-origin requests and which types of requests are permitted.

# CORS in action:
After running the code files:
1. Go to a any Directory, be it Downloads, Desktop, or any other directory.
2. Make a directory named <b>cors-test</b> and cd into the directory.
```
mkdir cors-test
```
```
cd cors-test
```
3. Open a text editor and paste the code of index.html and script.js in their respective files and make sure that they are in the same directory.
4. Now open the terminal and in the same directory type:
```
python -m http.server
```
5. This will open a page on the web browser where you will get a click button. That button is set to getting all the albums from the database.
6. If you will click on the button nothing will be visible as it is.
7. On the same page press <b>cmd option I</b>
8. For Windows and linux user press <b> Ctrl option I</b>
9. After the developer console opens up clcik on the button and select inspect and then go to console.
It will show like:
![Screenshot 2023-07-23 at 8 57 13 PM](https://github.com/Amarjit0511/go-task-sheet/assets/54772122/3fe13993-ebe2-4071-92cf-737e1d19b237)

Now to confirm if CORS is properly implemented, go to the network section and click on any of the request you made and then click on headers.
You will see something like:
![Screenshot 2023-07-23 at 9 00 10 PM](https://github.com/Amarjit0511/go-task-sheet/assets/54772122/d7a2484a-a1e0-4da2-9e8a-337941ae4ce8)
Note the response header: it should be something like <b> Access-Control-Allow-Origin:* </b>

![Screenshot 2023-07-23 at 9 50 12 PM](https://github.com/Amarjit0511/go-task-sheet/assets/54772122/b391ac80-ae9a-4c0b-84fa-37ecf2d42436)


# Working:
This CORS allows web application running on one domain to make requests to GO server running on another domain. Without CORS today's modern web browsers normally restricts these types of cross-origin requests for security reasons.

1. Our Go server is running on localhost:8443 and is handling endpoints for GET, POST, PUT, DELETE.
2. Now we also have a simple web app that is running on port 8000.
3. In my Go code, I have enabled CORS on all endpoints and set it to allow all origins(AllowAllOrigins=true) thereby allowing any domain to access the server's resources. Any doamin can now fetch the data from the server, etc.
4. On clicking the fetch button on the web application, the fetchAlbums() is triggered.
5. Inside the fetchAlbum() a cross origin request is made using fetch() functionwhich points to localhost:8443/albums/get
6. In the VS Code images above an OPTION request might be visible. It is basically a preflight request that takes before any request to the GO server cnfirming if the server allows this request.
7. Hence after all this GET or other request can be made.

