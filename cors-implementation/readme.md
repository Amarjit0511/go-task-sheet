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

